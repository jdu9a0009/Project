package storage

import (
	"context"
	staff_service "staff_service/genproto"
	"time"
)

type StorageI interface {
	Staff() StaffI
	StaffTarif() StaffTarifI
}
type CacheI interface {
	Cache() RedisI
}

type RedisI interface {
	Create(ctx context.Context, key string, obj interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, response interface{}) (bool, error)
	Delete(ctx context.Context, key string) error
}

type StaffI interface {
	CreateStaff(context.Context, *staff_service.CreateStaffRequest) (string, error)
	GetStaff(context.Context, *staff_service.IdRequest) (*staff_service.Staff, error)
	GetAllStaff(context.Context, *staff_service.GetAllStaffRequest) (*staff_service.GetAllStaffResponse, error)
	UpdateStaff(context.Context, *staff_service.UpdateStaffRequest) (string, error)
	DeleteStaff(context.Context, *staff_service.IdRequest) (string, error)
}

type StaffTarifI interface {
	CreateStaffTarif(context.Context, *staff_service.CreateStaffTarifRequest) (string, error)
	GetStaffTarif(context.Context, *staff_service.IdRequest) (*staff_service.StaffTarif, error)
	GetAllStaffTarif(context.Context, *staff_service.GetAllStaffTarifRequest) (*staff_service.GetAllStaffTarifResponse, error)
	UpdateStaffTarif(context.Context, *staff_service.UpdateStaffTarifRequest) (string, error)
	DeleteStaffTarif(context.Context, *staff_service.IdRequest) (string, error)
}
