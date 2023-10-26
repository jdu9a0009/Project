package service

import (
	"context"
	"fmt"
	"product_service/config"
	product_service "product_service/genproto"
	"product_service/pkg/logger"
	"product_service/storage"
)

type ProductService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	product_service.UnimplementedProductServiceServer
}

func NewProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *ProductService {
	return &ProductService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *ProductService) Create(ctx context.Context, req *product_service.CreateProductRequest) (*product_service.IdResponse, error) {
	fmt.Println("service create")

	id, err := b.storage.Product().CreateProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &product_service.IdResponse{Id: id}, nil
}

func (b *ProductService) Get(ctx context.Context, req *product_service.IdRequest) (*product_service.GetProductResponse, error) {
	staff, err := b.storage.Product().GetProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &product_service.GetProductResponse{Product: staff}, nil
}

func (b *ProductService) GetAll(ctx context.Context, req *product_service.GetAllProductRequest) (*product_service.GetAllProductResponse, error) {
	staffes, err := b.storage.Product().GetAllProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return &product_service.GetAllProductResponse{Products: staffes.Products,
		Count: staffes.Count}, nil
}

func (s *ProductService) Update(ctx context.Context, req *product_service.UpdateProductRequest) (*product_service.Response, error) {
	resp, err := s.storage.Product().UpdateProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &product_service.Response{Message: resp}, nil
}

func (s *ProductService) Delete(ctx context.Context, req *product_service.IdRequest) (*product_service.Response, error) {
	resp, err := s.storage.Product().DeleteProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &product_service.Response{Message: resp}, nil
}
