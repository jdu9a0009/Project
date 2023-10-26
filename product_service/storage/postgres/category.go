package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"product_service/pkg/helper"
	"time"

	product_service "product_service/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategory(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (b *categoryRepo) CreateCategory(c context.Context, req *product_service.CreateCategoriesRequest) (string, error) {
	id := uuid.NewString()

	query := `
		  INSERT INTO "categories"(
			"id",
			"name",
			"created_at")
		  VALUES ($1, $2, NOW())`

	var parentID interface{}
	if req.ParentId != "" {
		query = `
		  INSERT INTO "categories"(
			"id",
			"name",
			"parent_id",
			"created_at")
		  VALUES ($1, $2, $3, NOW())`
		parentID = req.ParentId
	}

	_, err := b.db.Exec(context.Background(), query,
		id,
		req.Name,
		parentID,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create Category: %w", err)
	}

	return id, nil
}

func (b *categoryRepo) GetCategory(c context.Context, req *product_service.IdRequest) (resp *product_service.Categories, err error) {
	query := `
				SELECT  
					"id", 
					"name", 
					"parent_id", 
					"created_at",
					"updated_at"
				FROM categories 
				WHERE id=$1`

	var (
		createdAt time.Time
		updatedAt sql.NullTime
		parentID  sql.NullString
	)

	category := product_service.Categories{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&category.Id,
		&category.Name,
		&parentID,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("Category not found")
		}
		return nil, fmt.Errorf("failed to get Category: %w", err)
	}
	category.CreatedAt = createdAt.Format(time.RFC3339)
	category.UpdatedAt = updatedAt.Time.Format(time.RFC3339)

	if parentID.Valid {
		category.ParentId = parentID.String
	}

	return &category, nil
}

func (b *categoryRepo) UpdateCategory(c context.Context, req *product_service.UpdateCategoriesRequest) (string, error) {

	query := `
				UPDATE categories 
				SET  
					name = $1, 
					parent_id = $2, 
					updated_at = NOW() 
				WHERE id = $3 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Name,
		req.ParentId,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update category: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("category with ID:  %s not found", req.Id)
	}

	return fmt.Sprintf("category with ID %s updated", req.Id), nil
}

func (b *categoryRepo) GetAllCategory(c context.Context, req *product_service.GetAllCategoriesRequest) (*product_service.GetAllCategoriesResponse, error) {
	var (
		resp   product_service.GetAllCategoriesResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM categories WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = b.db.QueryRow(c, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `
		SELECT 
		"id", 
		"name", 
		"parent_id",
		"created_at",
		"updated_at"
		FROM categories 
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
	parentID := sql.NullString{}
	for rows.Next() {
		var category product_service.Categories

		err = rows.Scan(
			&category.Id,
			&category.Name,
			&parentID,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning category err: %w", err)
		}

		category.CreatedAt = createdAt.Format(time.RFC3339)
		category.UpdatedAt = updatedAt.Format(time.RFC3339)

		if parentID.Valid {
			category.ParentId = parentID.String
		}

		resp.Categories = append(resp.Categories, &category)
	}

	return &resp, nil
}

func (b *categoryRepo) DeleteCategory(c context.Context, req *product_service.IdRequest) (resp string, err error) {
	query := `
			DELETE 
				FROM categories 
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
		return "", fmt.Errorf("category with ID %s not found", req.Id)

	}

	return fmt.Sprintf("category with ID %s deleted", req.Id), nil
}
