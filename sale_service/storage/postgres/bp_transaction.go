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

type BpTransactionRepo struct {
	db *pgxpool.Pool
}

func NewBpTransaction(db *pgxpool.Pool) *BpTransactionRepo {
	return &BpTransactionRepo{
		db: db,
	}
}

func (b *BpTransactionRepo) CreateBpTransaction(c context.Context, req *sale_service.CreateBpTransactionRequest) (string, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO "bptransactions"(
			"id", 
			"branch_id", 
			"staff_id", 
			"product_id", 
			"price",
			"type",
			"quantity",
			"created_at")
		VALUES ($1, $2, $3, $4, $5,$6,$7, NOW())
	`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.BranchId,
		req.StaffId,
		req.ProductId,
		req.Price,
		req.Type,
		req.Quantity,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create BpTransaction: %w", err)
	}

	return id, nil
}
func (b *BpTransactionRepo) GetBpTransaction(c context.Context, req *sale_service.IdRequest) (resp *sale_service.BpTransaction, err error) {
	query := `
				SELECT  
					"id", 
					"branch_id", 
					"staff_id", 
					"product_id", 
					"price",
					"type",
					"quantity",
					"created_at",
					"updated_at"
				FROM bptransactions 
				WHERE id=$1`

	var (
		createdAt time.Time
		updatedAt sql.NullTime
	)

	Bptransaction := sale_service.BpTransaction{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&Bptransaction.Id,
		&Bptransaction.BranchId,
		&Bptransaction.StaffId,
		&Bptransaction.ProductId,
		&Bptransaction.Price,
		&Bptransaction.Type,
		&Bptransaction.Quantity,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("BpTransaction not found")
		}
		return nil, fmt.Errorf("failed to get BpTransaction: %w", err)
	}
	Bptransaction.CreatedAt = createdAt.Format(time.RFC3339)
	Bptransaction.UpdatedAt = updatedAt.Time.Format(time.RFC3339)

	return &Bptransaction, nil
}

func (b *BpTransactionRepo) UpdateBpTransaction(c context.Context, req *sale_service.UpdateBpTransactionRequest) (string, error) {

	query := `
				UPDATE bptransactions 
				SET  
					branch_id = $1, 
					staff_id = $2, 
					product_id = $3, 
					price=$4,
					type=$5,
					quantity =$6,
					updated_at = NOW() 
				WHERE id = $7 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.BranchId,
		req.StaffId,
		req.ProductId,
		req.Price,
		req.Type,
		req.Quantity,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update BpTransaction: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("BpTransaction with ID:  %s not found", req.Id)
	}

	return fmt.Sprintf("BpTransaction with ID %s updated", req.Id), nil
}

func (b *BpTransactionRepo) GetAllBpTransaction(c context.Context, req *sale_service.GetAllBpTransactionRequest) (*sale_service.GetAllBpTransactionResponse, error) {
	var (
		resp   sale_service.GetAllBpTransactionResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND type ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM bptransactions WHERE true ` + filter

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
		"branch_id", 
		"staff_id", 
		"product_id", 
		"price",
		"type",
		"quantity",
		"created_at",
		"updated_at"
		FROM bptransactions 
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
		var BpTransaction sale_service.BpTransaction

		err = rows.Scan(
			&BpTransaction.Id,
			&BpTransaction.BranchId,
			&BpTransaction.ProductId,
			&BpTransaction.Price,
			&BpTransaction.Type,
			&BpTransaction.Quantity,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning staff err: %w", err)
		}

		BpTransaction.CreatedAt = createdAt.Format(time.RFC3339)
		BpTransaction.UpdatedAt = updatedAt.Format(time.RFC3339)

		resp.Bpransactions = append(resp.Bpransactions, &BpTransaction)
	}

	return &resp, nil
}
func (b *BpTransactionRepo) DeleteBpTransaction(c context.Context, req *sale_service.IdRequest) (resp string, err error) {
	query := `
			DELETE 
				FROM bptransactions 
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
