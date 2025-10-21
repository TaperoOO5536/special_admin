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

type ItemServiceHandler struct {
	itemService *service.ItemService
}

func NewItemServiceHandler(itemService *service.ItemService) *ItemServiceHandler {
	return &ItemServiceHandler{ itemService: itemService }
}

func ItemToGetItemInfoResponse(item *models.Item) (*pb.GetItemInfoResponse) {
	pbPictures := make([]*pb.PictureInfo, 0, len(item.Pictures))
	for _, picture := range item.Pictures {
		pbPicture := &pb.PictureInfo{
			Id:       picture.ID.String(),
			Picture:  picture.Path,
			MimeType: picture.MimeType,
		}
		pbPictures = append(pbPictures, pbPicture)
	}
	
	response := &pb.GetItemInfoResponse{
		Id:            item.ID.String(),
		Title:         item.Title,
		Description:   item.Description,
		Price:         int32(item.Price),
		LittlePicture: &pb.LittlePictureInfo{
			Picture:  item.LittlePicture,
			MimeType: item.MimeType,
		},
		Pictures: pbPictures,
	}
	
	return response
}

func (h *ItemServiceHandler) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.GetItemInfoResponse, error) {
	if req.Title == "" {
		err := status.Error(codes.InvalidArgument, "title is required")
		return nil, err
	}
	if req.Description == "" {
		err := status.Error(codes.InvalidArgument, "title is required")
		return nil, err
	}
	if req.LittlePicture == nil {
		err := status.Error(codes.InvalidArgument, "little picture is required")
		return nil, err
	}
	if req.Price == 0 {
		err := status.Error(codes.InvalidArgument, "price is required and must be > 0")
		return nil, err
	}

	itemID := uuid.New()

	item := &models.Item{
		ID:            itemID,
		Title:         req.Title,
		Description:   req.Description,
		Price:         int64(req.Price),
		LittlePicture: req.LittlePicture.Picture,
		MimeType:      req.LittlePicture.MimeType,
	}

	createdItem, err := h.itemService.CreateItem(ctx, item)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create item: %v", err)
	}
	
	return ItemToGetItemInfoResponse(createdItem), nil
}

func (h *ItemServiceHandler) GetItemInfo(ctx context.Context, req *pb.GetItemInfoRequest) (*pb.GetItemInfoResponse, error) {
	if req.Id == "" {
		err := status.Error(codes.InvalidArgument, "item id is required")
		return nil, err
	}	

	itemID, err := uuid.Parse(req.Id)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid item id")
		return nil, err
	}

	item, err := h.itemService.GetItemInfo(ctx, itemID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get item: %v", err)
	}

	return ItemToGetItemInfoResponse(item), nil
}

func (h *ItemServiceHandler) GetItems(ctx context.Context, req *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
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
	
	paginatedItems, err := h.itemService.GetItems(ctx, pagination)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get items: %v", err)
	}

	response := &pb.GetItemsResponse{
		Items: make([]*pb.ItemInfoForList, 0, len(paginatedItems.Items)),
		Total:   int32(paginatedItems.TotalCount),
		Page:    int32(paginatedItems.Page),
		PerPage: int32(paginatedItems.PerPage),
	}	

	for _, item := range paginatedItems.Items {
		pbItem := &pb.ItemInfoForList{
			Id:            item.ID.String(),
			Title:         item.Title,
			Price:         int32(item.Price),
			LittlePicture: &pb.LittlePictureInfo{
				Picture:  item.LittlePicture,
				MimeType: item.MimeType,
			},
		}
		response.Items = append(response.Items, pbItem)
	}

	return response, nil
}
	
func (h *ItemServiceHandler) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.GetItemInfoResponse, error) {
	if req.Id == "" {
		err := status.Error(codes.InvalidArgument, "item id is required")
		return nil, err
	}
	if req.Title == nil && req.Description == nil && req.Price == 0 &&
		req.LittlePicture == nil {
		err := status.Error(codes.InvalidArgument, "at least one argument is required")
		return nil, err
	}

	itemID, err := uuid.Parse(req.Id)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid item id")
		return nil, err
	}

	var isPriceUpdated bool
	item := &models.Item{
		ID:            itemID,
	}
	if req.Title != nil {
		item.Title = *req.Title
	}
	if req.Description != nil {
		item.Description = *req.Description
	}
	if req.LittlePicture != nil {
		item.LittlePicture = req.LittlePicture.Picture
		item.MimeType = req.LittlePicture.MimeType
	}
	if req.Price != 0 {
		item.Price = int64(req.Price)
		isPriceUpdated = true
	}

	updatedItem, err := h.itemService.UpdateItem(ctx, item, isPriceUpdated)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to update item: %v", err)
	}
	
	return ItemToGetItemInfoResponse(updatedItem), nil
}

func (h *ItemServiceHandler) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*emptypb.Empty, error) {
	if req.Id == "" {
		err := status.Error(codes.InvalidArgument, "item id is required")
		return nil, err
	}

	itemID, err := uuid.Parse(req.Id)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid item id")
		return nil, err
	}

	err = h.itemService.DeleteItem(ctx, itemID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}