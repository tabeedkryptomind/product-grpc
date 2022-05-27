package products_services

import (
	"context"
	"fmt"
	"grpc-product/models"
	"grpc-product/product_pb"
)

type Server struct {
	product_pb.UnimplementedProductServiceServer
}

func (*Server) CreateProduct(ctx context.Context, in *product_pb.ProductRequest) (*product_pb.ProductResponse, error) {
	product_name := in.GetProduct().GetProductName()
	tag := in.GetProduct().GetTag()
	price := in.GetProduct().GetPrice()
	pro := models.Product{
		ProductName: product_name,
		Tag:         tag,
		Price:       float32(price),
	}
	fmt.Println(pro)
	return &product_pb.ProductResponse{
		Result: "product created",
	}, nil

}

func (*Server) GetProduct(ctx context.Context, in *product_pb.ReceiveProduct) (*product_pb.Product, error) {
	pro := in.GetProductName()
	if pro == " " {
		fmt.Println("empty req")
		return &product_pb.Product{}, nil
	}
	return &product_pb.Product{
		ProductName: "milk",
		Tag:         "oreal",
		Price:       3,
	}, nil

}
