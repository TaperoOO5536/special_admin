package api

import (
	"context"

	"github.com/TaperoOO5536/special_admin/internal/models"
	"github.com/TaperoOO5536/special_admin/internal/service"
	pb "github.com/TaperoOO5536/special_admin/pkg/proto/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ItemPictureServiceHandler struct {
	itemPictureService *service.ItemPictureService
}

func NewItemPictureServiceHandler(itemPictureService *service.ItemPictureService) *ItemPictureServiceHandler {
	return &ItemPictureServiceHandler{ itemPictureService: itemPictureService }
}

func (h *ItemPictureServiceHandler) CreateItemPicture(ctx context.Context, req *pb.CreatePictureRequest) (*pb.GetItemInfoResponse, error) {
	if req.Picture.Picture == nil {
		err := status.Error(codes.InvalidArgument, "picture is required")
		return nil, err
	}
	if req.Picture.MimeType == "" {
		err := status.Error(codes.InvalidArgument, "mimetype is required")
		return nil, err
	}
	if req.ParentId == "" {
		err := status.Error(codes.InvalidArgument, "item id is required")
		return nil, err
	}

	itemPictureID := uuid.New()
	ItemID, err := uuid.Parse(req.ParentId)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid item id")
		return nil, err
	}

	itemPicture := &models.ItemPicture{
		ID:       itemPictureID,
		Path:     req.Picture.Picture,
		MimeType: req.Picture.MimeType,
		ItemID:  ItemID,
	}

	item, err := h.itemPictureService.CreateItemPicture(ctx, itemPicture)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create item picture: %v", err)
	}
	
	return ItemToGetItemInfoResponse(item), nil
}

func (h *ItemPictureServiceHandler) DeleteItemPicture(ctx context.Context, req *pb.DeletePictureRequest) (*emptypb.Empty, error) {
	if req.Id == "" {
		err := status.Error(codes.InvalidArgument, "item picture id is required")
		return nil, err
	}

	ItemPictureID, err := uuid.Parse(req.Id)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid item picture id")
		return nil, err
	}

	err = h.itemPictureService.DeleteItemPicture(ctx, ItemPictureID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}