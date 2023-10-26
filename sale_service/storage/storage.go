package storage

import (
	"context"
	sale_service "sale_service/genproto"
	"time"
)

type StorageI interface {
	Sale() SaleI
	StaffTransaction() StaffTransactionI
	BpTransaction() BpTransactionI
	SaleProduct() SaleProductI
}
type CacheI interface {
	Cache() RedisI
}

type RedisI interface {
	Create(ctx context.Context, key string, obj interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, response interface{}) (bool, error)
	Delete(ctx context.Context, key string) error
}

type SaleI interface {
	CreateSale(context.Context, *sale_service.CreateSaleRequest) (string, error)
	GetSale(context.Context, *sale_service.IdRequest) (*sale_service.Sale, error)
	GetAllSale(context.Context, *sale_service.GetAllSaleRequest) (*sale_service.GetAllSaleResponse, error)
	UpdateSale(context.Context, *sale_service.UpdateSaleRequest) (string, error)
	DeleteSale(context.Context, *sale_service.IdRequest) (string, error)
}

type StaffTransactionI interface {
	CreateStaffTransaction(context.Context, *sale_service.CreateStaffTransactionRequest) (string, error)
	GetStaffTransaction(context.Context, *sale_service.IdRequest) (*sale_service.StaffTransaction, error)
	GetAllStaffTransaction(context.Context, *sale_service.GetAllStaffTransactionRequest) (*sale_service.GetAllStaffTransactionResponse, error)
	UpdateStaffTransaction(context.Context, *sale_service.UpdateStaffTransactionRequest) (string, error)
	DeleteStaffTransaction(context.Context, *sale_service.IdRequest) (string, error)
}

type BpTransactionI interface {
	CreateBpTransaction(context.Context, *sale_service.CreateBpTransactionRequest) (string, error)
	GetBpTransaction(context.Context, *sale_service.IdRequest) (*sale_service.BpTransaction, error)
	GetAllBpTransaction(context.Context, *sale_service.GetAllBpTransactionRequest) (*sale_service.GetAllBpTransactionResponse, error)
	UpdateBpTransaction(context.Context, *sale_service.UpdateBpTransactionRequest) (string, error)
	DeleteBpTransaction(context.Context, *sale_service.IdRequest) (string, error)
}
type SaleProductI interface {
	CreateSaleProduct(context.Context, *sale_service.CreateSaleProductRequest) (string, error)
	GetSaleProduct(context.Context, *sale_service.IdRequest) (*sale_service.SaleProduct, error)
	GetAllSaleProduct(context.Context, *sale_service.GetAllSaleProductRequest) (*sale_service.GetAllSaleProductResponse, error)
	UpdateSaleProduct(context.Context, *sale_service.UpdateSaleProductRequest) (string, error)
	DeleteSaleProduct(context.Context, *sale_service.IdRequest) (string, error)
}
