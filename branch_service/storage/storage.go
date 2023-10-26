package storage

import (
	sale_service "branch_service/genproto"
	"context"
	"time"
)

type StorageI interface {
	Branch() BranchI
	Bproduct() BproductI
}
type CacheI interface {
	Cache() RedisI
}

type RedisI interface {
	Create(ctx context.Context, key string, obj interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, response interface{}) (bool, error)
	Delete(ctx context.Context, key string) error
}

type BproductI interface {
	CreateBproduct(context.Context, *sale_service.CreateBranchProductRequest) (string, error)
	GetBproduct(context.Context, *sale_service.IdRequest) (*sale_service.BranchProduct, error)
	GetAllBproduct(context.Context, *sale_service.GetAllBranchProductRequest) (*sale_service.GetAllBranchProductResponse, error)
	UpdateBproduct(context.Context, *sale_service.UpdateBranchProductRequest) (string, error)
	DeleteBproduct(context.Context, *sale_service.IdRequest) (string, error)
}

type BranchI interface {
	CreateBranch(context.Context, *sale_service.CreateBranchRequest) (string, error)
	GetBranch(context.Context, *sale_service.IdRequest) (*sale_service.Branch, error)
	GetAllBranch(context.Context, *sale_service.GetAllBranchRequest) (*sale_service.GetAllBranchResponse, error)
	UpdateBranch(context.Context, *sale_service.UpdateBranchRequest) (string, error)
	DeleteBranch(context.Context, *sale_service.IdRequest) (string, error)
}
