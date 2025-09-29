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

type IventPictureServiceHandler struct {
	iventPictureService *service.IventPictureService
}

func NewIventPictureServiceHandler(iventPictureService *service.IventPictureService) *IventPictureServiceHandler {
	return &IventPictureServiceHandler{ iventPictureService: iventPictureService }
}

func (h *IventPictureServiceHandler) CreateIventPicture(ctx context.Context, req *pb.CreatePictureRequest) (*pb.GetIventInfoResponse, error) {
	if req.Picture.Picture == nil {
		err := status.Error(codes.InvalidArgument, "picture is required")
		return nil, err
	}
	if req.Picture.MimeType == "" {
		err := status.Error(codes.InvalidArgument, "mimetype is required")
		return nil, err
	}
	if req.ParentId == "" {
		err := status.Error(codes.InvalidArgument, "ivent id is required")
		return nil, err
	}

	iventPictureID := uuid.New()
	IventID, err := uuid.Parse(req.ParentId)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid ivent id")
		return nil, err
	}

	iventPicture := &models.IventPicture{
		ID:       iventPictureID,
		Path:     req.Picture.Picture,
		MimeType: req.Picture.MimeType,
		IventID:  IventID,
	}

	ivent, err := h.iventPictureService.CreateIventPicture(ctx, iventPicture)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create ivent picture: %v", err)
	}
	
	return IventToGetIventInfoResponse(ivent), nil
}

func (h *IventPictureServiceHandler) DeleteIventPicture(ctx context.Context, req *pb.DeletePictureRequest) (*emptypb.Empty, error) {
	if req.Id == "" {
		err := status.Error(codes.InvalidArgument, "ivent picture id is required")
		return nil, err
	}

	IventPictureID, err := uuid.Parse(req.Id)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid ivent picture id")
		return nil, err
	}

	err = h.iventPictureService.DeleteIventPicture(ctx, IventPictureID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}