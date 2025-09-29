package api

import (
	"context"

	pb "github.com/TaperoOO5536/special_admin/pkg/proto/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	pb.UnimplementedSpecialAdminServiceServer
	userHandler         *UserServiceHandler
	iventHadler         *IventServiceHandler
	iventPictureHandler *IventPictureServiceHandler
}

func NewHandler(
	userHandler         *UserServiceHandler,
	iventHandler        *IventServiceHandler,
	iventPictureHandler *IventPictureServiceHandler,
) *Handler {
	return &Handler{
		userHandler:         userHandler,
		iventHadler:         iventHandler,
		iventPictureHandler: iventPictureHandler,
	}
}


//users

func (h *Handler) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	return h.userHandler.GetUsers(ctx, req)
}

//ivents

func (h *Handler) CreateIvent(ctx context.Context, req *pb.CreateIventRequest) (*pb.GetIventInfoResponse, error) {
	return h.iventHadler.CreateIvent(ctx, req)
}

func (h *Handler) GetIventInfo(ctx context.Context, req *pb.GetIventInfoRequest) (*pb.GetIventInfoResponse, error) {
	return h.iventHadler.GetIventInfo(ctx, req)
}

func (h *Handler) GetIvents(ctx context.Context, req *pb.GetIventsRequest) (*pb.GetIventsResponse, error) {
	return h.iventHadler.GetIvents(ctx, req)
}

func (h *Handler) UpdateIvent(ctx context.Context, req *pb.UpdateIventRequest) (*pb.GetIventInfoResponse, error) {
	return h.iventHadler.UpdateIvent(ctx, req)
}

func (h *Handler) DeleteIvent(ctx context.Context, req *pb.DeleteIventRequest) (*emptypb.Empty, error) {
	return h.iventHadler.DeleteIvent(ctx, req)
}

//iventPictures

func (h *Handler) CreateIventPicture(ctx context.Context, req *pb.CreatePictureRequest) (*pb.GetIventInfoResponse, error) {
	return h.iventPictureHandler.CreateIventPicture(ctx, req)
}

func (h *Handler) DeleteIventPicture(ctx context.Context, req *pb.DeletePictureRequest) (*emptypb.Empty, error) {
	return h.iventPictureHandler.DeleteIventPicture(ctx, req)
}