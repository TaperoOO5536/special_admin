package api

import (
	"context"

	"github.com/TaperoOO5536/special_admin/internal/models"
	"github.com/TaperoOO5536/special_admin/internal/service"
	pb "github.com/TaperoOO5536/special_admin/pkg/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceHandler struct {
	userService *service.UserService
}

func NewUserServiceHandler(userService *service.UserService) *UserServiceHandler {
	return &UserServiceHandler{ userService: userService}
}

func (h *UserServiceHandler) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	pagination := models.Pagination{}	

	if req.Page == 0 {
		pagination.Page = 1
	} else {
		pagination.Page = int(req.Page)
	}
	if req.PerPage == 0 {
		pagination.PerPage = 1
	} else {
		pagination.PerPage = int(req.PerPage)
	}
	
	paginatedUsers, err := h.userService.GetUsers(ctx, pagination)
	if err != nil {
		err := status.Error(codes.Internal, err.Error())
		return nil, err
	}

	response := &pb.GetUsersResponse{
		Users: make([]*pb.UserInfo, 0, len(paginatedUsers.Users)),
		Total:   int32(paginatedUsers.TotalCount),
		Page:    int32(paginatedUsers.Page),
		PerPage: int32(paginatedUsers.PerPage),
	}
	
	for _, user := range paginatedUsers.Users {
		pbUser := &pb.UserInfo{
			Name:        user.Name,
			Surname:     user.Surname,
			Nickname:    user.Nickname,
			PhoneNumber: user.PhoneNumber,
		}
		response.Users = append(response.Users, pbUser)
	}

	return response, nil
}