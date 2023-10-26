package service

import (
	"context"
	"staff_service/config"
	staff_service "staff_service/genproto"
	"staff_service/pkg/logger"
	"staff_service/storage"
)

type StaffTarifService struct {
	cfg     config.Config
	log     logger.LoggerI
	storage storage.StorageI
	staff_service.UnimplementedStaffTarifServiceServer
}

func NewStaffTarifTarifService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *StaffTarifService {
	return &StaffTarifService{
		cfg:     cfg,
		log:     log,
		storage: strg,
	}
}

func (b *StaffTarifService) Create(ctx context.Context, req *staff_service.CreateStaffTarifRequest) (*staff_service.IdResponse, error) {

	id, err := b.storage.StaffTarif().CreateStaffTarif(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &staff_service.IdResponse{Id: id}, nil
}

func (b *StaffTarifService) Get(ctx context.Context, req *staff_service.IdRequest) (*staff_service.GetStaffTarifResponse, error) {
	tarif, err := b.storage.StaffTarif().GetStaffTarif(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &staff_service.GetStaffTarifResponse{Tarif: tarif}, nil
}

func (b *StaffTarifService) GetAll(ctx context.Context, req *staff_service.GetAllStaffTarifRequest) (*staff_service.GetAllStaffTarifResponse, error) {
	staffes, err := b.storage.StaffTarif().GetAllStaffTarif(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return &staff_service.GetAllStaffTarifResponse{StaffTarifs: staffes.StaffTarifs,
		Count: staffes.Count}, nil
}

func (s *StaffTarifService) Update(ctx context.Context, req *staff_service.UpdateStaffTarifRequest) (*staff_service.Response, error) {
	resp, err := s.storage.StaffTarif().UpdateStaffTarif(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &staff_service.Response{Message: resp}, nil
}

func (s *StaffTarifService) Delete(ctx context.Context, req *staff_service.IdRequest) (*staff_service.Response, error) {
	resp, err := s.storage.StaffTarif().DeleteStaffTarif(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return &staff_service.Response{Message: resp}, nil
}
