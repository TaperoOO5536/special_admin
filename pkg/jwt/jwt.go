package jwt

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    Login string   `json:"login"`
    jwt.RegisteredClaims
}

type JWTManager struct {
    AccessDuration  time.Duration
    RefreshDuration time.Duration
}

func NewJWTManager(accessDur, refreshDur string) (*JWTManager) {
    ad, _ := time.ParseDuration(accessDur)
    rd, _ := time.ParseDuration(refreshDur)
    return &JWTManager{AccessDuration: ad, RefreshDuration: rd}
}

func (m *JWTManager) GenerateTokenPair(login, secret string) (accessToken, refreshToken string, err error) {
    accessToken, err = m.GenerateAccessToken(login, secret)
    if err != nil {
        return
    }
    refreshToken = generateRandomString(64)
    return accessToken, refreshToken, nil
}

func (m *JWTManager) GenerateAccessToken(login, secret string) (string, error) {
    accessClaims := Claims{
        Login: login,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.AccessDuration)),
        },
    }
    access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
    accessToken, err := access.SignedString([]byte(secret))
    if err != nil {
        return "", err
    }

    return accessToken, nil
}

func (m *JWTManager) ValidateToken(tokenStr, secret string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })
    if err != nil {
        return nil, err
    }
    claims, ok := token.Claims.(*Claims)
    if !ok {
        return nil, err
    }
    return claims, nil
}

func generateRandomString(length int) string {
    bytes := make([]byte, length)
    rand.Read(bytes)
    return hex.EncodeToString(bytes)[:length]
}

func HashToken(token string) string {
    h := sha256.Sum256([]byte(token))
    return hex.EncodeToString(h[:])
}