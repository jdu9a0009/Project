package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"sale_service/pkg/helper"
	"time"

	sale_service "sale_service/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type staffTransactionRepo struct {
	db *pgxpool.Pool
}

func NewStaffTransaction(db *pgxpool.Pool) *staffTransactionRepo {
	return &staffTransactionRepo{
		db: db,
	}
}

func (b *staffTransactionRepo) CreateStaffTransaction(c context.Context, req *sale_service.CreateStaffTransactionRequest) (string, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO "stafftransactions"(
			"id", 
			"sale_id", 
			"staff_id", 
			"type", 
			"source_type",
			"amount",
			"about_text"
			"created_at",
			"updated_at")
		VALUES ($1, $2, $3, $4, $5,$6,$7, NOW(), NOW())
	`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.SaleId,
		req.StaffId,
		req.Type,
		req.SourceType,
		req.Amount,
		req.AboutText,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create sale Transaction: %w", err)
	}

	return id, nil
}
func (b *staffTransactionRepo) GetStaffTransaction(c context.Context, req *sale_service.IdRequest) (resp *sale_service.StaffTransaction, err error) {
	query := `
				SELECT  
					"id", 
					"sale_id", 
					"staff_id", 
					"type", 
					"source_type",
					"amount",
					"about_text"
					"created_at",
					"updated_at"
				FROM stafftransactions 
				WHERE id=$1`

	var (
		createdAt time.Time
		updatedAt sql.NullTime
	)

	Stransaction := sale_service.StaffTransaction{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&Stransaction.Id,
		&Stransaction.SaleId,
		&Stransaction.StaffId,
		&Stransaction.Type,
		&Stransaction.SourceType,
		&Stransaction.Amount,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("staffTransaction not found")
		}
		return nil, fmt.Errorf("failed to get staff Transaction: %w", err)
	}
	Stransaction.CreatedAt = createdAt.Format(time.RFC3339)
	Stransaction.UpdatedAt = updatedAt.Time.Format(time.RFC3339)

	return &Stransaction, nil
}

func (b *staffTransactionRepo) UpdateStaffTransaction(c context.Context, req *sale_service.UpdateStaffTransactionRequest) (string, error) {

	query := `
				UPDATE stafftransactions 
				SET  
					sale_id = $1, 
					staff_id = $2, 
					type = $3, 
					source_type=$4,
					amount=$5,
					about_text =$6,
					updated_at = NOW() 
				WHERE id = $5 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.SaleId,
		req.StaffId,
		req.Type,
		req.SourceType,
		req.Amount,
		req.AboutText,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update staff transaction: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("staff transaction with ID:  %s not found", req.Id)
	}

	return fmt.Sprintf("staff transaction with ID %s updated", req.Id), nil
}

func (b *staffTransactionRepo) GetAllStaffTransaction(c context.Context, req *sale_service.GetAllStaffTransactionRequest) (*sale_service.GetAllStaffTransactionResponse, error) {
	var (
		resp   sale_service.GetAllStaffTransactionResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Type != "" {
		filter += " AND type ILIKE '%' || :search || '%' "
		params["search"] = req.Type
	}

	if req.SaleId != "" {
		filter += ` AND sale_id = :sale_id`
		params["sale_id"] = req.SaleId
	}

	if req.StaffId != "" {
		filter += ` AND sale_id = :staff_id`
		params["staff_id"] = req.StaffId
	}

	if req.Amount != 0 {
		filter += ` AND amount = :amount`
		params["amount"] = req.Amount
	}

	countQuery := `SELECT count(1) FROM stafftransactions WHERE true ` + filter

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
		"sale_id", 
		"staff_id", 
		"type", 
		"source_type",
		"amount",
		"about_text"
		"created_at",
		"updated_at"
		FROM stafftransactions 
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
		var Stransaction sale_service.StaffTransaction

		err = rows.Scan(
			&Stransaction.Id,
			&Stransaction.SaleId,
			&Stransaction.StaffId,
			&Stransaction.Type,
			&Stransaction.SourceType,
			&Stransaction.Amount,
			&Stransaction.AboutText,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning staff err: %w", err)
		}

		Stransaction.CreatedAt = createdAt.Format(time.RFC3339)
		Stransaction.UpdatedAt = updatedAt.Format(time.RFC3339)

		resp.Stransactions = append(resp.Stransactions, &Stransaction)
	}

	return &resp, nil
}

func (b *staffTransactionRepo) DeleteStaffTransaction(c context.Context, req *sale_service.IdRequest) (resp string, err error) {
	query := `
			DELETE 
				FROM stafftransactions 
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
		return "", fmt.Errorf("staff transactions with ID %s not found", req.Id)

	}

	return fmt.Sprintf("staff transactions with ID %s deleted", req.Id), nil
}
