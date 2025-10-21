package api

import (
	"context"

	pb "github.com/TaperoOO5536/special_admin/pkg/proto/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	pb.UnimplementedSpecialAdminServiceServer
	pb.UnimplementedAdminAuthServiceServer
	userHandler         *UserServiceHandler
	eventHandler        *EventServiceHandler
	eventPictureHandler *EventPictureServiceHandler
	itemHandler         *ItemServiceHandler
	itemPictureHandler  *ItemPictureServiceHandler
	orderHandler        *OrderServiceHandler
	authHandler         *AuthServiceHandler
}

func NewHandler(
	userHandler *UserServiceHandler,
	eventHandler *EventServiceHandler,
	eventPictureHandler *EventPictureServiceHandler,
	itemServiceHandler *ItemServiceHandler,
	itemPictureHandler *ItemPictureServiceHandler,
	orderHandler *OrderServiceHandler,
	authHandler *AuthServiceHandler,
) *Handler {
	return &Handler{
		userHandler:         userHandler,
		eventHandler:        eventHandler,
		eventPictureHandler: eventPictureHandler,
		itemHandler:         itemServiceHandler,
		itemPictureHandler:  itemPictureHandler,
		orderHandler:        orderHandler,
		authHandler:         authHandler,
	}
}

//users

func (h *Handler) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	return h.userHandler.GetUsers(ctx, req)
}

//events

func (h *Handler) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.GetEventInfoResponse, error) {
	return h.eventHandler.CreateEvent(ctx, req)
}

func (h *Handler) GetEventInfo(ctx context.Context, req *pb.GetEventInfoRequest) (*pb.GetEventInfoResponse, error) {
	return h.eventHandler.GetEventInfo(ctx, req)
}

func (h *Handler) GetEvents(ctx context.Context, req *pb.GetEventsRequest) (*pb.GetEventsResponse, error) {
	return h.eventHandler.GetEvents(ctx, req)
}

func (h *Handler) UpdateEvent(ctx context.Context, req *pb.UpdateEventRequest) (*pb.GetEventInfoResponse, error) {
	return h.eventHandler.UpdateEvent(ctx, req)
}

func (h *Handler) DeleteEvent(ctx context.Context, req *pb.DeleteEventRequest) (*emptypb.Empty, error) {
	return h.eventHandler.DeleteEvent(ctx, req)
}

//eventPictures

func (h *Handler) CreateEventPicture(ctx context.Context, req *pb.CreatePictureRequest) (*pb.GetEventInfoResponse, error) {
	return h.eventPictureHandler.CreateEventPicture(ctx, req)
}

func (h *Handler) DeleteEventPicture(ctx context.Context, req *pb.DeletePictureRequest) (*emptypb.Empty, error) {
	return h.eventPictureHandler.DeleteEventPicture(ctx, req)
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

//itemPictures

func (h *Handler) CreateItemPicture(ctx context.Context, req *pb.CreatePictureRequest) (*pb.GetItemInfoResponse, error) {
	return h.itemPictureHandler.CreateItemPicture(ctx, req)
}

func (h *Handler) DeleteItemPicture(ctx context.Context, req *pb.DeletePictureRequest) (*emptypb.Empty, error) {
	return h.itemPictureHandler.DeleteItemPicture(ctx, req)
}

//orders

func (h *Handler) GetOrderInfo(ctx context.Context, req *pb.GetOrderInfoRequest) (*pb.GetOrderInfoResponse, error) {
	return h.orderHandler.GetOrderInfo(ctx, req)
}

func (h *Handler) GetOrders(ctx context.Context, req *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	return h.orderHandler.GetOrders(ctx, req)
}

func (h *Handler) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.GetOrderInfoResponse, error) {
	return h.orderHandler.UpdateOrder(ctx, req)
}

func (h *Handler) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*emptypb.Empty, error) {
	return h.orderHandler.DeleteOrder(ctx, req)
}

//auth

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return h.authHandler.Login(ctx, req)
}

func (h *Handler) RefreshToken(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	return h.authHandler.RefreshToken(ctx, req)
}

func (h *Handler) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return h.authHandler.Logout(ctx, req)
}
