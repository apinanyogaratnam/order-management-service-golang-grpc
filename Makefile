proto:
	protoc -I protos/ --go_out=plugins=grpc:orders --go_opt=paths=source_relative --go_opt=Morder.proto=github.com/apinanyogaratnam/order-management-service-golang-grpc/orders order.proto
