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

type EventPictureServiceHandler struct {
	eventPictureService *service.EventPictureService
}

func NewEventPictureServiceHandler(eventPictureService *service.EventPictureService) *EventPictureServiceHandler {
	return &EventPictureServiceHandler{ eventPictureService: eventPictureService }
}

func (h *EventPictureServiceHandler) CreateEventPicture(ctx context.Context, req *pb.CreatePictureRequest) (*pb.GetEventInfoResponse, error) {
	if req.Picture.Picture == nil {
		err := status.Error(codes.InvalidArgument, "picture is required")
		return nil, err
	}
	if req.Picture.MimeType == "" {
		err := status.Error(codes.InvalidArgument, "mimetype is required")
		return nil, err
	}
	if req.ParentId == "" {
		err := status.Error(codes.InvalidArgument, "event id is required")
		return nil, err
	}

	eventPictureID := uuid.New()
	EventID, err := uuid.Parse(req.ParentId)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid event id")
		return nil, err
	}

	eventPicture := &models.EventPicture{
		ID:       eventPictureID,
		Path:     req.Picture.Picture,
		MimeType: req.Picture.MimeType,
		EventID:  EventID,
	}

	event, err := h.eventPictureService.CreateEventPicture(ctx, eventPicture)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create event picture: %v", err)
	}
	
	return EventToGetEventInfoResponse(event), nil
}

func (h *EventPictureServiceHandler) DeleteEventPicture(ctx context.Context, req *pb.DeletePictureRequest) (*emptypb.Empty, error) {
	if req.Id == "" {
		err := status.Error(codes.InvalidArgument, "event picture id is required")
		return nil, err
	}

	EventPictureID, err := uuid.Parse(req.Id)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid event picture id")
		return nil, err
	}

	err = h.eventPictureService.DeleteEventPicture(ctx, EventPictureID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}