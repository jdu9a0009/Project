package service

import (
	"context"
	"fmt"
	"sale_service/config"
	sale_service "sale_service/genproto"
	"sale_service/pkg/logger"
	"sale_service/storage"
)

type SaleProductService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	sale_service.UnimplementedSaleProductServiceServer
}

func NewSaleProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *SaleProductService {
	return &SaleProductService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *SaleProductService) Create(ctx context.Context, req *sale_service.CreateSaleProductRequest) (*sale_service.IdResponse, error) {
	fmt.Println("service create")

	id, err := b.storage.SaleProduct().CreateSaleProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.IdResponse{Id: id}, nil
}

func (b *SaleProductService) Get(ctx context.Context, req *sale_service.IdRequest) (*sale_service.GetSaleProductResponse, error) {
	saleProduct, err := b.storage.SaleProduct().GetSaleProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.GetSaleProductResponse{SaleProduct: saleProduct}, nil
}

func (b *SaleProductService) GetAll(ctx context.Context, req *sale_service.GetAllSaleProductRequest) (*sale_service.GetAllSaleProductResponse, error) {
	saleProducts, err := b.storage.SaleProduct().GetAllSaleProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return &sale_service.GetAllSaleProductResponse{SaleProducts: saleProducts.SaleProducts,
		Count: saleProducts.Count}, nil
}

func (s *SaleProductService) Update(ctx context.Context, req *sale_service.UpdateSaleProductRequest) (*sale_service.Response, error) {
	resp, err := s.storage.SaleProduct().UpdateSaleProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}

func (s *SaleProductService) Delete(ctx context.Context, req *sale_service.IdRequest) (*sale_service.Response, error) {
	resp, err := s.storage.SaleProduct().DeleteSaleProduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}
