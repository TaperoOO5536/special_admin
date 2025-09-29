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

type OrderServiceHandler struct {
	orderService *service.OrderService
}

func NewOrderServiceHandler(orderService *service.OrderService) *OrderServiceHandler {
	return &OrderServiceHandler{ orderService: orderService }
}

func OrderToGetOrderInfoResponse(order *models.Order) (*pb.GetOrderInfoResponse) {
	pbOrderItems := make([]*pb.OrderItemInfoForList, 0, len(order.OrderItems))
	for _, orderItem := range order.OrderItems {
		pbOrderItem := &pb.OrderItemInfoForList{
			Id:            orderItem.ID.String(),
			ItemId:        orderItem.Item.ID.String(),
			Title:         orderItem.Item.Title,
			Price:         orderItem.Item.Price,
			Quantity:      orderItem.Quantity,
			LittlePicture: &pb.LittlePictureInfo{
				Picture: orderItem.Item.LittlePicture,
				MimeType: orderItem.Item.MimeType,
			},
		}
		pbOrderItems = append(pbOrderItems, pbOrderItem)
	}

	response := &pb.GetOrderInfoResponse{
		Id:             order.ID.String(),
		Number:         order.Number,
		UserNickname:   order.User.Nickname,
		FormDate:       timestamppb.New(order.FormDate),
		CompletionDate: timestamppb.New(order.CompletionDate),
		Comment:        order.Comment,
		Status:         order.Status,
		OrderAmount:    order.OrderAmount,
		Items:          pbOrderItems,
	}
	
	return response
}

func (h *OrderServiceHandler) GetOrderInfo(ctx context.Context, req *pb.GetOrderInfoRequest) (*pb.GetOrderInfoResponse, error) {
	if req.Id == "" {
		err := status.Error(codes.InvalidArgument, "order id is required")
		return nil, err
	}	

	OrderID, err := uuid.Parse(req.Id)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid order id")
		return nil, err
	}

	order, err := h.orderService.GetOrderInfo(ctx, OrderID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get order: %v", err)
	}

	return OrderToGetOrderInfoResponse(order), nil
}

func (h *OrderServiceHandler) GetOrders(ctx context.Context, req *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	orders, err := h.orderService.GetOrders(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get orders: %v", err)
	}

	pbOrders := make([]*pb.OrderInfoForList, 0, len(orders))
	for _, order := range orders {
		pbOrder := &pb.OrderInfoForList{
			Id:             order.ID.String(),
			Number:         order.Number,
			CompletionDate: timestamppb.New(order.CompletionDate),
			Status:         order.Status,
			OrderAmount:    order.OrderAmount,
		}
		pbOrders = append(pbOrders, pbOrder)
	}

	return &pb.GetOrdersResponse{
		Orders: pbOrders,
	}, nil
}

func (h *OrderServiceHandler) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.GetOrderInfoResponse, error) {
	if req.Id == "" {
		err := status.Error(codes.InvalidArgument, "order id is required")
		return nil, err
	}
	if req.Status == nil && req.Comment == nil{
		err := status.Error(codes.InvalidArgument, "at least one argument is required")
		return nil, err
	}

	orderID, err := uuid.Parse(req.Id)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid order id")
		return nil, err
	}

	order := &models.Order{
		ID:            orderID,
	}
	if req.Status != nil {
		order.Status = *req.Status
	}
	if req.Comment != nil {
		order.Comment = *req.Comment
	}

	updatedOrder, err := h.orderService.UpdateOrder(ctx, order)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to update order: %v", err)
	}
	
	return OrderToGetOrderInfoResponse(updatedOrder), nil
}

func (h *OrderServiceHandler) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*emptypb.Empty, error) {
	if req.Id == "" {
		err := status.Error(codes.InvalidArgument, "order id is required")
		return nil, err
	}

	OrderID, err := uuid.Parse(req.Id)
	if err != nil {
		err := status.Error(codes.InvalidArgument, "invalid order id")
		return nil, err
	}

	err = h.orderService.DeleteOrder(ctx, OrderID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}