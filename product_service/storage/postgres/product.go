package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"product_service/pkg/helper"

	product_service "product_service/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProduct(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (b *productRepo) CreateProduct(c context.Context, req *product_service.CreateProductRequest) (string, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO "products"(
			"id", 
			"name", 
			"price", 
			"barcode", 
			"created_at",
		     "category_id")
		VALUES ($1, $2, $3, $4, NOW(), $5)
	`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Price,
		req.Barcode,
		req.CategoryId,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create Product: %w", err)
	}

	return id, nil
}
func (b *productRepo) GetProduct(c context.Context, req *product_service.IdRequest) (*product_service.Product, error) {
	query := `
		SELECT 
			id, 
			"name", 
			"price", 
			"barcode", 
			"created_at",
			"updated_at",
			"category_id"
		FROM products 
		WHERE id=$1`

	var (
		createdAt  sql.NullString
		updatedAt  sql.NullString
		categoryID sql.NullString
		product    product_service.Product
	)

	err := b.db.QueryRow(c, query, req.Id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Barcode,
		&categoryID,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	if categoryID.Valid {
		product.CategoryId = categoryID.String
	}

	if createdAt.Valid {
		product.CreatedAt = createdAt.String
	}

	if updatedAt.Valid {
		product.UpdatedAt = updatedAt.String
	}

	return &product, nil
}

func (b *productRepo) UpdateProduct(c context.Context, req *product_service.UpdateProductRequest) (string, error) {

	query := `
				UPDATE products 
				SET  
					name = $1, 
					price = $2, 
					barcode = $3, 
					category_id =$4,
					updated_at = NOW() 
				WHERE id = $5 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Name,
		req.Price,
		req.Barcode,
		req.CategoryId,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update product: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("product with ID %s not found", req.Id)
	}

	return fmt.Sprintf("product with ID %s updated", req.Id), nil
}

func (b *productRepo) GetAllProduct(c context.Context, req *product_service.GetAllProductRequest) (*product_service.GetAllProductResponse, error) {
	var (
		resp   product_service.GetAllProductResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND (name ILIKE '%' || :search || '%' OR barcode ILIKE '%' || :search || '%') "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM products WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = b.db.QueryRow(c, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `
		SELECT 
		      id, 
			  "name", 
			  "price", 
			  "barcode", 
			  "category_id",
			  "created_at",
			  "updated_at"
		FROM products 
		WHERE true` + filter

	query += " ORDER BY created_at DESC LIMIT :limit OFFSET :offset"
	params["limit"] = req.Limit
	params["offset"] = req.Offset

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := b.db.Query(c, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}
	defer rows.Close()

	createdAt := sql.NullString{}
	updatedAt := sql.NullString{}
	for rows.Next() {
		var product product_service.Product

		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Barcode,
			&product.CategoryId,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning product err: %w", err)
		}

		product.CreatedAt = createdAt.String
		product.UpdatedAt = updatedAt.String

		resp.Products = append(resp.Products, &product)
	}

	return &resp, nil
}
func (b *productRepo) DeleteProduct(c context.Context, req *product_service.IdRequest) (resp string, err error) {
	query := `
			DELETE 
				FROM products 
			WHERE id = $1 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Id,
	)
	if err != nil {
		return "", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("product with ID %s not found", req.Id)

	}

	return fmt.Sprintf("product with ID %s deleted", req.Id), nil
}
