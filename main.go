package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/apinanyogaratnam/order-management-service-golang-grpc/orders"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()

	orderService := orders.Server{}
	orders.RegisterOrderServiceServer(grpcServer, &orderService)

	log.Println("Starting gRPC listener on port 9000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
