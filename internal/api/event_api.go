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
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventServiceHandler struct {
	eventService *service.EventService
}

func NewEventServiceHandler(eventService *service.EventService) *EventServiceHandler {
	return &EventServiceHandler{ eventService: eventService }
}

func EventToGetEventInfoResponse(event *models.Event) (*pb.GetEventInfoResponse) {
	pbPictures := make([]*pb.PictureInfo, 0, len(event.Pictures))
	for _, picture := range event.Pictures {
		pbPicture := &pb.PictureInfo{
			Id:       picture.ID.String(),
			Picture:  picture.Path,
			MimeType: picture.MimeType,
		}
		pbPictures = append(pbPictures, pbPicture)
	}

	pbUserEvents := make([]*pb.UserEventInfo, 0, len(event.UserEvents))
	for _, userEvent := range event.UserEvents {
		pbUserEvent := &pb.UserEventInfo{
			Id: userEvent.  ID.String(),
			UserNickname:   userEvent.User.Nickname,
			NumberOfGuests: userEvent.NumberOfGuests,
		}
		pbUserEvents = append(pbUserEvents, pbUserEvent)
	}
	
	response := &pb.GetEventInfoResponse{
		Id:            event.ID.String(),
		Title:         event.Title,
		Description:   event.Description,
		Datetime:      timestamppb.New(event.DateTime),
		Price:         event.Price,
		TotalSeats:    event.TotalSeats,
		OccupiedSeats: event.OccupiedSeats,
		LittlePicture: &pb.LittlePictureInfo{
			Picture: event.LittlePicture,
			MimeType: event.MimeType,
		},
		Pictures: pbPictures,
		UserEvents: pbUserEvents,
	}
	
	return response
}

func (h *EventServiceHandler) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.GetEventInfoResponse, error) {
	if req.Title == "" {
		err := status.Error(codes.InvalidArgument, "title is required")
		return nil, err
	}
	if req.Description == "" {
		err := status.Error(codes.InvalidArgument, "title is required")
		return nil, err
	}
	if req.Datetime == nil {
		err := status.Error(codes.InvalidArgument, "datetime is required")
		return nil, err
	}
	if req.TotalSeats == 0 {
		err := status.Error(codes.InvalidArgument, "total seats is required")
		return nil, err
	}
	if req.OccupiedSeats == nil {
		err := status.Error(codes.InvalidArgument, "occupied seats is required")
		return nil, err
	}
	if req.LittlePicture == nil {
		err := status.Error(codes.InvalidArgument, "little picture is required")
		return nil, err
	}
	if req.Price == nil {
		err := status.Error(codes.InvalidArgument, "price is required")
		return nil, err
	}

	eventID := uuid.New()

	event := &models.Event{
		ID:            eventID,
		Title:         req.Title,
		Description:   req.Description,
		DateTime:      req.Datetime.AsTime(),
		Price:         req.Price.Value,
		TotalSeats:    req.TotalSeats,
		OccupiedSeats: req.OccupiedSeats.Value,
		LittlePicture: req.LittlePicture.Picture,
		MimeType:      req.LittlePicture.MimeType,
	}

	createdEvent, err := h.eventService.CreateEvent(ctx, event)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create event: %v", err)
	}
	
	return EventToGetEventInfoResponse(createdEvent), nil
}

func (h *EventServiceHandler) GetEventInfo(ctx context.Context, req *pb.GetEventInfoRequest) (*pb.GetEventInfoResponse, error) {

	if req.Id == "" {
		err := status.Error(codes.InvalidArgument, "event id is required")
		return nil, err
	}	

	EventID, err := uuid.Parse(req.Id)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid event id")
		return nil, err
	}

	event, err := h.eventService.GetEventInfo(ctx, EventID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get event: %v", err)
	}

	return EventToGetEventInfoResponse(event), nil
}

func (h *EventServiceHandler) GetEvents(ctx context.Context, req *pb.GetEventsRequest) (*pb.GetEventsResponse, error) {
	events, err := h.eventService.GetEvents(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get events: %v", err)
	}

	pbEvents := make([]*pb.EventInfoForList, 0, len(events))
	for _, event := range events {
		pbEvent := &pb.EventInfoForList{
			Id:            event.ID.String(),
			Title:         event.Title,
			Datetime:      timestamppb.New(event.DateTime),
			Price:         event.Price,
			TotalSeats:    event.TotalSeats,
			OccupiedSeats: event.OccupiedSeats,
			LittlePicture: &pb.LittlePictureInfo{
				Picture:  event.LittlePicture,
				MimeType: event.MimeType,
			},
		}
		pbEvents = append(pbEvents, pbEvent)
	}

	return &pb.GetEventsResponse{
		Events: pbEvents,
	}, nil
}
	
func (h *EventServiceHandler) UpdateEvent(ctx context.Context, req *pb.UpdateEventRequest) (*pb.GetEventInfoResponse, error) {
	if req.Id == "" {
		err := status.Error(codes.InvalidArgument, "event id is required")
		return nil, err
	}
	if req.Title == nil && req.Description == nil && req.Datetime == nil &&
		req.Price == nil && req.TotalSeats == nil && req.OccupiedSeats == nil &&
		req.LittlePicture == nil {
		err := status.Error(codes.InvalidArgument, "at least one argument is required")
		return nil, err
	}

	eventID, err := uuid.Parse(req.Id)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid event id")
		return nil, err
	}

	var isPriceUpdated bool
	var isOccupiedSeats bool
	event := &models.Event{
		ID:            eventID,
	}
	if req.Title != nil {
		event.Title = *req.Title
	}
	if req.Datetime != nil {
		event.DateTime = req.Datetime.AsTime()
	}
	if req.Description != nil {
		event.Description = *req.Description
	}
	if req.LittlePicture != nil {
		event.LittlePicture = req.LittlePicture.Picture
		event.MimeType = req.LittlePicture.MimeType
	}
	if req.TotalSeats != nil {
		event.TotalSeats = *req.TotalSeats
	}
	if req.Price != nil {
		event.Price = req.Price.Value
		isPriceUpdated = true
	}
	if req.OccupiedSeats != nil {
		event.OccupiedSeats = req.OccupiedSeats.Value
		isOccupiedSeats = true
	}

	updatedEvent, err := h.eventService.UpdateEvent(ctx, event, isPriceUpdated, isOccupiedSeats)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to update event: %v", err)
	}
	
	return EventToGetEventInfoResponse(updatedEvent), nil
}

func (h *EventServiceHandler) DeleteEvent(ctx context.Context, req *pb.DeleteEventRequest) (*emptypb.Empty, error) {
	if req.Id == "" {
		err := status.Error(codes.InvalidArgument, "event id is required")
		return nil, err
	}

	EventID, err := uuid.Parse(req.Id)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid event id")
		return nil, err
	}

	err = h.eventService.DeleteEvent(ctx, EventID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}