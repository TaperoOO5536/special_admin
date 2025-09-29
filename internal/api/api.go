package api

import (
	"context"

	pb "github.com/TaperoOO5536/special_admin/pkg/proto/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	pb.UnimplementedSpecialAdminServiceServer
	userHandler         *UserServiceHandler
	iventHandler         *IventServiceHandler
	iventPictureHandler *IventPictureServiceHandler
	itemHandler  *ItemServiceHandler
}

func NewHandler(
	userHandler         *UserServiceHandler,
	iventHandler        *IventServiceHandler,
	iventPictureHandler *IventPictureServiceHandler,
	itemServiceHandler  *ItemServiceHandler,
) *Handler {
	return &Handler{
		userHandler:         userHandler,
		iventHandler:         iventHandler,
		iventPictureHandler: iventPictureHandler,
		itemHandler:  itemServiceHandler,
	}
}


//users

func (h *Handler) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	return h.userHandler.GetUsers(ctx, req)
}

//ivents

func (h *Handler) CreateIvent(ctx context.Context, req *pb.CreateIventRequest) (*pb.GetIventInfoResponse, error) {
	return h.iventHandler.CreateIvent(ctx, req)
}

func (h *Handler) GetIventInfo(ctx context.Context, req *pb.GetIventInfoRequest) (*pb.GetIventInfoResponse, error) {
	return h.iventHandler.GetIventInfo(ctx, req)
}

func (h *Handler) GetIvents(ctx context.Context, req *pb.GetIventsRequest) (*pb.GetIventsResponse, error) {
	return h.iventHandler.GetIvents(ctx, req)
}

func (h *Handler) UpdateIvent(ctx context.Context, req *pb.UpdateIventRequest) (*pb.GetIventInfoResponse, error) {
	return h.iventHandler.UpdateIvent(ctx, req)
}

func (h *Handler) DeleteIvent(ctx context.Context, req *pb.DeleteIventRequest) (*emptypb.Empty, error) {
	return h.iventHandler.DeleteIvent(ctx, req)
}

//iventPictures

func (h *Handler) CreateIventPicture(ctx context.Context, req *pb.CreatePictureRequest) (*pb.GetIventInfoResponse, error) {
	return h.iventPictureHandler.CreateIventPicture(ctx, req)
}

func (h *Handler) DeleteIventPicture(ctx context.Context, req *pb.DeletePictureRequest) (*emptypb.Empty, error) {
	return h.iventPictureHandler.DeleteIventPicture(ctx, req)
}

//items

func (h *Handler) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.GetItemInfoResponse, error) {
	return h.itemHandler.CreateItem(ctx, req)
}

func (h *Handler) GetItemInfo(ctx context.Context, req *pb.GetItemInfoRequest) (*pb.GetItemInfoResponse, error) {
	return h.itemHandler.GetItemInfo(ctx, req)
}

func (h *Handler) GetItems(ctx context.Context, req *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
	return h.itemHandler.GetItems(ctx, req)
}

func (h *Handler) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.GetItemInfoResponse, error) {
	return h.itemHandler.UpdateItem(ctx, req)
}

func (h *Handler) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*emptypb.Empty, error) {
	return h.itemHandler.DeleteItem(ctx, req)
}