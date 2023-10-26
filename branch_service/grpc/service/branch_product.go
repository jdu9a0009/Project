package service

import (
	"branch_service/config"
	branch_service "branch_service/genproto"
	"branch_service/pkg/logger"
	"branch_service/storage"
	"context"
)

type BranchProductService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	branch_service.UnimplementedBranchProductServiceServer
}

func NewBranchProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *BranchProductService {
	return &BranchProductService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *BranchProductService) Create(ctx context.Context, req *branch_service.CreateBranchProductRequest) (*branch_service.IdResponse, error) {
	id, err := b.storage.Bproduct().CreateBproduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &branch_service.IdResponse{Id: id}, nil
}

func (b *BranchProductService) Get(ctx context.Context, req *branch_service.IdRequest) (*branch_service.GetBranchProductResponse, error) {
	branchProduct, err := b.storage.Bproduct().GetBproduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &branch_service.GetBranchProductResponse{BranchProduct: branchProduct}, nil
}

func (b *BranchProductService) GetAll(ctx context.Context, req *branch_service.GetAllBranchProductRequest) (*branch_service.GetAllBranchProductResponse, error) {
	branchProducts, err := b.storage.Bproduct().GetAllBproduct(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return &branch_service.GetAllBranchProductResponse{BranchProducts: branchProducts.BranchProducts,
		Count: branchProducts.Count}, nil
}

func (s *BranchProductService) Update(ctx context.Context, req *branch_service.UpdateBranchProductRequest) (*branch_service.Response, error) {
	resp, err := s.storage.Bproduct().UpdateBproduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &branch_service.Response{Message: resp}, nil
}

func (s *BranchProductService) Delete(ctx context.Context, req *branch_service.IdRequest) (*branch_service.Response, error) {
	resp, err := s.storage.Bproduct().DeleteBproduct(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &branch_service.Response{Message: resp}, nil
}
