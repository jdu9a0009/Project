package service

import (
	"context"
	"sale_service/config"
	sale_service "sale_service/genproto"
	"sale_service/pkg/logger"
	"sale_service/storage"
)

type BpTransactionService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	sale_service.UnimplementedBpTransactionServiceServer
}

func NewBpTransactionService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *BpTransactionService {
	return &BpTransactionService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *BpTransactionService) Create(ctx context.Context, req *sale_service.CreateBpTransactionRequest) (*sale_service.IdResponse, error) {

	id, err := b.storage.BpTransaction().CreateBpTransaction(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.IdResponse{Id: id}, nil
}

func (b *BpTransactionService) Get(ctx context.Context, req *sale_service.IdRequest) (*sale_service.GetBpTransactionResponse, error) {
	transaction, err := b.storage.BpTransaction().GetBpTransaction(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.GetBpTransactionResponse{Bpransaction: transaction}, nil
}

func (b *BpTransactionService) GetAll(ctx context.Context, req *sale_service.GetAllBpTransactionRequest) (*sale_service.GetAllBpTransactionResponse, error) {
	transactions, err := b.storage.BpTransaction().GetAllBpTransaction(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return &sale_service.GetAllBpTransactionResponse{Bpransactions: transactions.Bpransactions,
		Count: transactions.Count}, nil
}

func (s *BpTransactionService) Update(ctx context.Context, req *sale_service.UpdateBpTransactionRequest) (*sale_service.Response, error) {
	resp, err := s.storage.BpTransaction().UpdateBpTransaction(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}

func (s *BpTransactionService) Delete(ctx context.Context, req *sale_service.IdRequest) (*sale_service.Response, error) {
	resp, err := s.storage.BpTransaction().DeleteBpTransaction(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &sale_service.Response{Message: resp}, nil
}
