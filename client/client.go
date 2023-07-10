package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/apinanyogaratnam/order-management-service-golang-grpc/orders"
)

func CreateOrder(c orders.OrderServiceClient) *orders.CreateOrderResponse {
	createOrderRequestMessage := orders.CreateOrderRequest{
		CartItems: []*orders.CartItem{
			{
				ProductId: 1,
				Quantity: 1,
				UserId: 1,
			},
			{
				ProductId: 2,
				Quantity: 2,
				UserId: 1,
			},
			{
				ProductId: 3,
				Quantity: 3,
				UserId: 1,
			},
		},
	}

	createOrderResponseMessage, err := c.CreateOrder(context.Background(), &createOrderRequestMessage)
	if err != nil {
		log.Fatalln("Failed to create order:", err)
	}

	log.Println("Order created:", createOrderResponseMessage)

	return createOrderResponseMessage
}

func GetOrder(c orders.OrderServiceClient, orderId uint32) *orders.GetOrderResponse{
	getOrderRequestMessage := orders.GetOrderRequest{
		OrderId: orderId,
	}

	getOrderResponseMessage, err := c.GetOrder(context.Background(), &getOrderRequestMessage)
	if err != nil {
		log.Fatalln("Failed to get order:", err)
	}

	log.Println("Order retrieved:", getOrderResponseMessage)

	return getOrderResponseMessage
}

func UpdateOrder(c orders.OrderServiceClient, orderId uint32) *orders.UpdateOrderResponse {
	updateOrderRequestMessage := orders.UpdateOrderRequest{
		OrderId: orderId,
		Status: orders.ORDER_STATUS_SHIPPED,
	}

	updateOrderResponseMessage, err := c.UpdateOrder(context.Background(), &updateOrderRequestMessage)
	if err != nil {
		log.Fatalln("Failed to update order:", err)
	}

	log.Println("Order updated:", updateOrderResponseMessage)

	return updateOrderResponseMessage
}

func DeleteOrder(c orders.OrderServiceClient, orderId uint32) *orders.DeleteOrderResponse {
	deleteOrderRequestMessage := orders.DeleteOrderRequest{
		OrderId: orderId,
	}

	deleteOrderResponseMessage, err := c.DeleteOrder(context.Background(), &deleteOrderRequestMessage)
	if err != nil {
		log.Fatalln("Failed to delete order:", err)
	}

	log.Println("Order deleted:", deleteOrderResponseMessage)

	return deleteOrderResponseMessage
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	defer conn.Close()

	c := orders.NewOrderServiceClient(conn)

	createdOrderResponse := CreateOrder(c)
	orderResponse := GetOrder(c, createdOrderResponse.OrderId)
	updatedOrderResponse := UpdateOrder(c, orderResponse.OrderId)
	DeleteOrder(c, updatedOrderResponse.OrderId)
}
