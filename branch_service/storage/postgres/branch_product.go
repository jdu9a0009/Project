package postgres

import (
	"branch_service/pkg/helper"
	"context"
	"database/sql"
	"fmt"
	"time"

	branch_service "branch_service/genproto"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type bproductRepo struct {
	db *pgxpool.Pool
}

func NewBranchProduct(db *pgxpool.Pool) *bproductRepo {
	return &bproductRepo{
		db: db,
	}
}

func (b *bproductRepo) CreateBproduct(c context.Context, req *branch_service.CreateBranchProductRequest) (string, error) {

	query := `
		INSERT INTO "bproducts"(
			"product_id", 
			"branch_id", 
			"count", 
			"created_at")
		VALUES ($1, $2, $3, now())
	`
	_, err := b.db.Exec(context.Background(), query,
		req.ProductId,
		req.BranchId,
		req.Count,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create branch product: %w", err)
	}

	return req.ProductId, nil
}

func (b *bproductRepo) GetBproduct(c context.Context, req *branch_service.IdRequest) (resp *branch_service.BranchProduct, err error) {
	query := `
				SELECT 
				   product_id,
					branch_id, 
					count, 
					created_at, 
					updated_at 
				FROM bproducts 
				WHERE product_id=$1`

	var (
		createdAt time.Time
		updatedAt sql.NullTime
	)

	bproduct := branch_service.BranchProduct{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&bproduct.ProductId,
		&bproduct.BranchId,
		&bproduct.Count,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("branch Product not found")
		}
		return nil, fmt.Errorf("failed to get branch product: %w", err)
	}

	bproduct.CreatedAt = createdAt.Format(time.RFC3339)
	if updatedAt.Valid {
		bproduct.UpdatedAt = updatedAt.Time.Format(time.RFC3339)
	}

	return &bproduct, nil
}

func (b *bproductRepo) UpdateBproduct(c context.Context, req *branch_service.UpdateBranchProductRequest) (string, error) {
	query := `
		UPDATE bproducts 
		SET 
			product_id = $1, 
			branch_id = $2, 
			count = $3, 
			updated_at = NOW() 
		WHERE product_id = $4 RETURNING product_id
	`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.ProductId,
		req.BranchId,
		req.Count,
		req.ProductId,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update branch product: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("branch product with ID %s not found", req.ProductId)
	}

	return fmt.Sprintf("branch product with ID %s updated", req.ProductId), nil
}

func (b *bproductRepo) GetAllBproduct(c context.Context, req *branch_service.GetAllBranchProductRequest) (*branch_service.GetAllBranchProductResponse, error) {
	var (
		resp   branch_service.GetAllBranchProductResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND product_id ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM bproducts WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = b.db.QueryRow(c, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `
		SELECT 
			product_id, 
			branch_id, 
			count, 
			created_at, 
			updated_at 
		FROM bproducts 
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

	createdAt := time.Time{}
	updatedAt := time.Time{}
	for rows.Next() {
		var bproduct branch_service.BranchProduct

		err = rows.Scan(
			&bproduct.ProductId,
			&bproduct.BranchId,
			&bproduct.Count,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning branch product err: %w", err)
		}

		bproduct.CreatedAt = createdAt.Format(time.RFC3339)
		bproduct.UpdatedAt = updatedAt.Format(time.RFC3339)

		resp.BranchProducts = append(resp.BranchProducts, &bproduct)
	}

	return &resp, nil
}

func (b *bproductRepo) DeleteBproduct(c context.Context, req *branch_service.IdRequest) (resp string, err error) {
	query := `
			DELETE 
				FROM bproducts 
			WHERE product_id = $1 RETURNING product_id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Id,
	)
	if err != nil {
		return "", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("branch product with ID %s not found", req.Id)

	}

	return fmt.Sprintf("branch  product with ID %s deleted", req.Id), nil
}
