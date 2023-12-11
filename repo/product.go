package repo

import (
	"context"
	"kevin/gogrpc/pb"
)

type ProductModel struct {
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price,omitempty"`
	Id          string  `json:"id,omitempty"`
	Image       string  `json:"image,omitempty"`
}
type ProductImpl struct {
	pb.UnimplementedProductServiceServer
}

func NewProductImpl() *ProductImpl {
	return &ProductImpl{}
}
func (p *ProductImpl) CreateProduct(c context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	//logic
	product := &ProductModel{
		Id:          "1",
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Image:       req.Image,
	}

	return &pb.CreateProductResponse{
		Id: product.Id,
	}, nil
}
func (p *ProductImpl) GetProduct(c context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	//logic
	product := &ProductModel{
		Id:          "1",
		Name:        "Product 1",
		Description: "Product 1 Description",
		Price:       100,
		Image:       "https://www.google.com",
	}

	return &pb.GetProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Image:       product.Image,
	}, nil
}
