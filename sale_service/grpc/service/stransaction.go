package service

import (
	"context"
	"sale_service/config"
	sale_service "sale_service/genproto"
	"sale_service/pkg/logger"
	"sale_service/storage"
)

type StaffTransactionService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	sale_service.UnimplementedStaffTransactionServiceServer
}

func NewStaffTransactionService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *StaffTransactionService {
	return &StaffTransactionService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *StaffTransactionService) Create(ctx context.Context, req *sale_service.CreateStaffTransactionRequest) (*sale_service.IdResponse, error) {

	id, err := b.storage.StaffTransaction().CreateStaffTransaction(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.IdResponse{Id: id}, nil
}

func (b *StaffTransactionService) Get(ctx context.Context, req *sale_service.IdRequest) (*sale_service.GetStaffTransactionResponse, error) {
	transaction, err := b.storage.StaffTransaction().GetStaffTransaction(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.GetStaffTransactionResponse{STransaction: transaction}, nil
}

func (b *StaffTransactionService) GetAll(ctx context.Context, req *sale_service.GetAllStaffTransactionRequest) (*sale_service.GetAllStaffTransactionResponse, error) {
	transactions, err := b.storage.StaffTransaction().GetAllStaffTransaction(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return &sale_service.GetAllStaffTransactionResponse{Stransactions: transactions.Stransactions,
		Count: transactions.Count}, nil
}

func (s *StaffTransactionService) Update(ctx context.Context, req *sale_service.UpdateStaffTransactionRequest) (*sale_service.Response, error) {
	resp, err := s.storage.StaffTransaction().UpdateStaffTransaction(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}

func (s *StaffTransactionService) Delete(ctx context.Context, req *sale_service.IdRequest) (*sale_service.Response, error) {
	resp, err := s.storage.StaffTransaction().DeleteStaffTransaction(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}
