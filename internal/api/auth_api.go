package api

import (
	"context"

	"github.com/TaperoOO5536/special_admin/internal/service"
	pb "github.com/TaperoOO5536/special_admin/pkg/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServiceHandler struct {
	authService *service.AuthService
}

func NewAuthServiceHandler(authService *service.AuthService) *AuthServiceHandler {
	return &AuthServiceHandler{ authService: authService }
}

func (h *AuthServiceHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Login == "" {
		err := status.Error(codes.InvalidArgument, "login is required")
		return nil, err
	}
	if req.Password == "" {
		err := status.Error(codes.InvalidArgument, "password is required")
		return nil, err
	}

	tokens, err := h.authService.Login(ctx, req.Login, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to login: %v", err)
	}

	return &pb.LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		Message:      "login successful",
	}, nil
}

func (h *AuthServiceHandler) RefreshToken(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	if req.GrantType == "" {
		return nil, status.Error(codes.Unauthenticated, "missing refresh token")
	}

	token, err := h.authService.RefreshToken(ctx, req.GrantType)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to refresh: %v", err)
	}
	
	return &pb.RefreshResponse{
		AccessToken: token,
		Message:     "token refreshed",
	}, nil
}

func (h *AuthServiceHandler) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
  if req.GrantType == "" {
    return nil, status.Error(codes.Unauthenticated, "missing refresh token")
  }

	err := h.authService.Logout(ctx, req.GrantType)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to logout: %v", err)
	}

	return &pb.LogoutResponse{
		Message: "logged out successfully",
	}, nil
}

