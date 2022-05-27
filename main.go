package main

import (
	"google.golang.org/grpc"
	"grpc-product/product_pb"
	"grpc-product/products_services"
	"log"
	"net"
)

func main() {
	opts := []grpc.ServerOption{}
	rpcServer := grpc.NewServer(opts...)
	product_pb.RegisterProductServiceServer(rpcServer, &products_services.Server{})
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot listen to the address  %v", err)
	}
	if err := rpcServer.Serve(listen); err != nil {
		log.Fatalf("cannot server the request %v", err)
	}
}
