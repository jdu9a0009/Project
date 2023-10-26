package storage

import (
	"context"
	product_service "product_service/genproto"
	"time"
)

type StorageI interface {
	Product() ProductI
	Category() CategoryI
}
type CacheI interface {
	Cache() RedisI
}

type RedisI interface {
	Create(ctx context.Context, key string, obj interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, response interface{}) (bool, error)
	Delete(ctx context.Context, key string) error
}

type ProductI interface {
	CreateProduct(context.Context, *product_service.CreateProductRequest) (string, error)
	GetProduct(context.Context, *product_service.IdRequest) (*product_service.Product, error)
	GetAllProduct(context.Context, *product_service.GetAllProductRequest) (*product_service.GetAllProductResponse, error)
	UpdateProduct(context.Context, *product_service.UpdateProductRequest) (string, error)
	DeleteProduct(context.Context, *product_service.IdRequest) (string, error)
}

type CategoryI interface {
	CreateCategory(context.Context, *product_service.CreateCategoriesRequest) (string, error)
	GetCategory(context.Context, *product_service.IdRequest) (*product_service.Categories, error)
	GetAllCategory(context.Context, *product_service.GetAllCategoriesRequest) (*product_service.GetAllCategoriesResponse, error)
	UpdateCategory(context.Context, *product_service.UpdateCategoriesRequest) (string, error)
	DeleteCategory(context.Context, *product_service.IdRequest) (string, error)
}
