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

type IventServiceHandler struct {
	iventService *service.IventService
}

func NewIventServiceHandler(iventService *service.IventService) *IventServiceHandler {
	return &IventServiceHandler{ iventService: iventService }
}

func IventToGetIventInfoResponse(ivent *models.Ivent) (*pb.GetIventInfoResponse) {
	pbPictures := make([]*pb.PictureInfo, 0, len(ivent.Pictures))
	for _, picture := range ivent.Pictures {
		pbPicture := &pb.PictureInfo{
			Id:       picture.ID.String(),
			Picture:  picture.Path,
			MimeType: picture.MimeType,
		}
		pbPictures = append(pbPictures, pbPicture)
	}

	pbUserIvents := make([]*pb.UserIventInfo, 0, len(ivent.UserIvents))
	for _, userIvent := range ivent.UserIvents {
		pbUserIvent := &pb.UserIventInfo{
			Id: userIvent.  ID.String(),
			UserNickname:   userIvent.User.Nickname,
			NumberOfGuests: userIvent.NumberOfGuests,
		}
		pbUserIvents = append(pbUserIvents, pbUserIvent)
	}
	
	response := &pb.GetIventInfoResponse{
		Id:            ivent.ID.String(),
		Title:         ivent.Title,
		Description:   ivent.Description,
		Datetime:      timestamppb.New(ivent.DateTime),
		Price:         ivent.Price,
		TotalSeats:    ivent.TotalSeats,
		OccupiedSeats: ivent.OccupiedSeats,
		LittlePicture: &pb.LittlePictureInfo{
			Picture: ivent.LittlePicture,
			MimeType: ivent.MimeType,
		},
		Pictures: pbPictures,
		UserIvents: pbUserIvents,
	}
	
	return response
}

func (h *IventServiceHandler) CreateIvent(ctx context.Context, req *pb.CreateIventRequest) (*pb.GetIventInfoResponse, error) {
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

	iventID := uuid.New()

	ivent := &models.Ivent{
		ID:            iventID,
		Title:         req.Title,
		Description:   req.Description,
		DateTime:      req.Datetime.AsTime(),
		Price:         req.Price.Value,
		TotalSeats:    req.TotalSeats,
		OccupiedSeats: req.OccupiedSeats.Value,
		LittlePicture: req.LittlePicture.Picture,
		MimeType:      req.LittlePicture.MimeType,
	}

	createdIvent, err := h.iventService.CreateIvent(ctx, ivent)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create ivent: %v", err)
	}
	
	return IventToGetIventInfoResponse(createdIvent), nil
}

func (h *IventServiceHandler) GetIventInfo(ctx context.Context, req *pb.GetIventInfoRequest) (*pb.GetIventInfoResponse, error) {

	if req.Id == "" {
		err := status.Error(codes.InvalidArgument, "ivent id is required")
		return nil, err
	}	

	IventID, err := uuid.Parse(req.Id)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid ivent id")
		return nil, err
	}

	ivent, err := h.iventService.GetIventInfo(ctx, IventID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get ivent: %v", err)
	}

	return IventToGetIventInfoResponse(ivent), nil
}

func (h *IventServiceHandler) GetIvents(ctx context.Context, req *pb.GetIventsRequest) (*pb.GetIventsResponse, error) {
	ivents, err := h.iventService.GetIvents(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get ivents: %v", err)
	}

	pbIvents := make([]*pb.IventInfoForList, 0, len(ivents))
	for _, ivent := range ivents {
		pbIvent := &pb.IventInfoForList{
			Id:            ivent.ID.String(),
			Title:         ivent.Title,
			Datetime:      timestamppb.New(ivent.DateTime),
			Price:         ivent.Price,
			TotalSeats:    ivent.TotalSeats,
			OccupiedSeats: ivent.OccupiedSeats,
			LittlePicture: &pb.LittlePictureInfo{
				Picture:  ivent.LittlePicture,
				MimeType: ivent.MimeType,
			},
		}
		pbIvents = append(pbIvents, pbIvent)
	}

	return &pb.GetIventsResponse{
		Ivents: pbIvents,
	}, nil
}
	
func (h *IventServiceHandler) UpdateIvent(ctx context.Context, req *pb.UpdateIventRequest) (*pb.GetIventInfoResponse, error) {
	if req.Id == "" {
		err := status.Error(codes.InvalidArgument, "ivent id is required")
		return nil, err
	}
	if req.Title == nil && req.Description == nil && req.Datetime == nil &&
		req.Price == nil && req.TotalSeats == nil && req.OccupiedSeats == nil &&
		req.LittlePicture == nil {
		err := status.Error(codes.InvalidArgument, "at least one argument is required")
		return nil, err
	}

	iventID, err := uuid.Parse(req.Id)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid ivent id")
		return nil, err
	}

	var isPriceUpdated bool
	var isOccupiedSeats bool
	ivent := &models.Ivent{
		ID:            iventID,
	}
	if req.Title != nil {
		ivent.Title = *req.Title
	}
	if req.Datetime != nil {
		ivent.DateTime = req.Datetime.AsTime()
	}
	if req.Description != nil {
		ivent.Description = *req.Description
	}
	if req.LittlePicture != nil {
		ivent.LittlePicture = req.LittlePicture.Picture
		ivent.MimeType = req.LittlePicture.MimeType
	}
	if req.TotalSeats != nil {
		ivent.TotalSeats = *req.TotalSeats
	}
	if req.Price != nil {
		ivent.Price = req.Price.Value
		isPriceUpdated = true
	}
	if req.OccupiedSeats != nil {
		ivent.OccupiedSeats = req.OccupiedSeats.Value
		isOccupiedSeats = true
	}

	updatedIvent, err := h.iventService.UpdateIvent(ctx, ivent, isPriceUpdated, isOccupiedSeats)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to update ivent: %v", err)
	}
	
	return IventToGetIventInfoResponse(updatedIvent), nil
}

func (h *IventServiceHandler) DeleteIvent(ctx context.Context, req *pb.DeleteIventRequest) (*emptypb.Empty, error) {
	if req.Id == "" {
		err := status.Error(codes.InvalidArgument, "ivent id is required")
		return nil, err
	}

	IventID, err := uuid.Parse(req.Id)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid ivent id")
		return nil, err
	}

	err = h.iventService.DeleteIvent(ctx, IventID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}