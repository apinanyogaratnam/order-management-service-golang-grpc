package orders

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {}

func (*Server) CreateOrder(ctx context.Context, message *CreateOrderRequest) (*CreateOrderResponse, error) {
	log.Println("Cart items:", message.CartItems)
	orderId := 123
	return &CreateOrderResponse{OrderId: uint32(orderId)}, nil
}

func (*Server) GetOrder(ctx context.Context, message *GetOrderRequest) (*GetOrderResponse, error) {
	log.Println("Order ID:", message.OrderId)
	cartItems := []*CartItem{}
	cartItems = append(cartItems, &CartItem{
		ProductId: 1,
		Quantity: 1,
		UserId: 1,
	})
	return &GetOrderResponse{
		OrderId: message.OrderId,
		UserId: 1,
		CartItems: cartItems,
		TotalPrice: 100,
		Status: ORDER_STATUS_CREATED,
	}, nil
}

func (*Server) UpdateOrder(ctx context.Context, message *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	log.Println("Order ID:", message.OrderId)
	return &UpdateOrderResponse{
		OrderId: message.OrderId,
		Status: message.Status,
	}, nil
}

func (*Server) DeleteOrder(ctx context.Context, message *DeleteOrderRequest) (*DeleteOrderResponse, error) {
	log.Println("Order ID:", message.OrderId)
	return &DeleteOrderResponse{
		OrderId: message.OrderId,
	}, nil
}
