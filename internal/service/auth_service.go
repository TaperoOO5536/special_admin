package service

import (
	"context"
	"errors"
	"time"

	"github.com/TaperoOO5536/special_admin/internal/repository"
	"github.com/TaperoOO5536/special_admin/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrFailedGenerateTokens = errors.New("failed to generate tokens")
	ErrFailedSavingRefreshToken = errors.New("failed to save refresh token")
	ErrInvalidAccessToken = errors.New("invalid access token")
    ErrInvalidRefreshToken = errors.New("invalid refresh token")
    ErrLogoutFailed = errors.New("logout failed")
)

type AuthService struct {
	adminAuthRepo repository.AdminAuthRepository
	Jwt           *jwt.JWTManager
	JwtSecret     string
}

func NewAuthService(adminAuthRepo repository.AdminAuthRepository, jwt *jwt.JWTManager, jwtSecret string) *AuthService {
	return &AuthService{
		adminAuthRepo: adminAuthRepo,
		Jwt:           jwt,
		JwtSecret:     jwtSecret,
	}
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func (s *AuthService) Login(ctx context.Context, login string, password string) (*Tokens, error) {
    admin, err := s.adminAuthRepo.GetAdmin(ctx, login)
    if err != nil || bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(password)) != nil {
        return nil, ErrInvalidCredentials
    }

    accessToken, refreshToken, err := s.Jwt.GenerateTokenPair(admin.Login, s.JwtSecret)
    if err != nil {
        return nil, ErrFailedGenerateTokens
    }

    if err := s.adminAuthRepo.UpdateRefreshToken(ctx, jwt.HashToken(refreshToken), 
        time.Now().Add(s.Jwt.RefreshDuration)); err != nil {
        return nil, ErrFailedSavingRefreshToken
    }

    return &Tokens{
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
    }, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, token string) (string, error) {
    admin, err := s.adminAuthRepo.GetRefreshToken(ctx, jwt.HashToken(token))
    if err != nil {
        return "", ErrInvalidRefreshToken
    }

    accessToken, err := s.Jwt.GenerateAccessToken(admin.Login, s.JwtSecret)
    if err != nil {
        return "", ErrFailedGenerateTokens
    }

    return accessToken, nil
}

func (s *AuthService) Logout(ctx context.Context, token string) (error) {    
    if err := s.adminAuthRepo.ClearRefreshToken(ctx); err != nil {
        return ErrLogoutFailed
    }
    return nil
}