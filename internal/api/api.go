package api

import (
	"context"

	pb "github.com/TaperoOO5536/special_admin/pkg/proto/v1"
)

type Handler struct {
	pb.UnimplementedSpecialAdminServiceServer
	userHandler      *UserServiceHandler
}

func NewHandler(
	userHandler      *UserServiceHandler,
) *Handler {
	return &Handler{
		userHandler:      userHandler,
	}
}


//users

func (h *Handler) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	return h.userHandler.GetUsers(ctx, req)
}
