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

type saleRepo struct {
	db *pgxpool.Pool
}

func NewSale(db *pgxpool.Pool) *saleRepo {
	return &saleRepo{
		db: db,
	}
}

func (b *saleRepo) CreateSale(c context.Context, req *sale_service.CreateSaleRequest) (string, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO "sales"(
			"id", 
			"branch_id", 
			"shop_assistant_id", 
			"cashier_id", 
			"price",
			"payment_type", 
			"client_name")
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.Branch_Id,
		req.ShopAssistant_Id,
		req.CashierId,
		req.Price,
		req.PaymentType,
		req.ClientName,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create Sale: %w", err)
	}

	return id, nil
}
func (b *saleRepo) GetSale(c context.Context, req *sale_service.IdRequest) (resp *sale_service.Sale, err error) {
	query := `
		SELECT 
			"id",
			"branch_id",
			"shop_assistant_id",
			"cashier_id",
			"price",
			"payment_type",
			"status",
			"client_name",
			"created_at",
			"updated_at"
		FROM sales
		WHERE id=$1`

	sale := sale_service.Sale{}
	var (
		createdAt time.Time
		updatedAt sql.NullTime
	)
	err = b.db.QueryRow(c, query, req.Id).Scan(
		&sale.Id,
		&sale.Branch_Id,
		&sale.ShopAssistant_Id,
		&sale.CashierId,
		&sale.Price,
		&sale.PaymentType,
		&sale.Status,
		&sale.ClientName,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("sale not found")
		}
		return nil, fmt.Errorf("failed to get sale: %w", err)
	}

	// Assign values from sql.NullString to the fields

	sale.CreatedAt = createdAt.Format(time.RFC3339)
	sale.UpdatedAt = updatedAt.Time.Format(time.RFC3339)

	return &sale, nil
}

func (b *saleRepo) UpdateSale(c context.Context, req *sale_service.UpdateSaleRequest) (string, error) {

	query := `
				UPDATE sales 
				SET  
				       "branch_id"=$1, 
				       "shop_assistant_id"=$2, 
				       "cashier_id"=$3, 
				       "payment_type"=$4, 
				       "client_name"=$5,
					updated_at = NOW() 
				WHERE id = $6 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Branch_Id,
		req.ShopAssistant_Id,
		req.CashierId,
		req.PaymentType,
		req.ClientName,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update sale: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("staff with ID %s not found", req.Id)
	}

	return fmt.Sprintf("sale with ID %s updated", req.Id), nil
}

func (b *saleRepo) GetAllSale(c context.Context, req *sale_service.GetAllSaleRequest) (*sale_service.GetAllSaleResponse, error) {
	var (
		resp   sale_service.GetAllSaleResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	// Search filter
	if req.Search != "" {
		filter += " AND (client_name ILIKE '%' || :search || '%' OR status ILIKE '%' || :search || '%') "
		params["search"] = req.Search
	}

	// Additional filters
	if req.Branch_Id != "" {
		filter += " AND branch_id = :branchId "
		params["branchId"] = req.Branch_Id
	}
	if req.PaymentType != "" {
		filter += " AND payment_type = :payment_type "
		params["payment_type"] = req.PaymentType
	}
	if req.ShopAssistant_Id != "" {
		filter += " AND shop_assisstant_id = :shopAssistantId "
		params["shopAssistantId"] = req.ShopAssistant_Id
	}
	if req.CashierId != "" {
		filter += " AND cashier_id = :cashierId "
		params["cashierId"] = req.CashierId
	}

	if req.CreatedAtFrom != "" {
		filter += " AND created_at >= :createdAtFrom "
		params["createdAtFrom"] = req.CreatedAtFrom
	}
	if req.CreatedAtTo != "" {
		filter += " AND created_at <= :createdAtTo "
		params["createdAtTo"] = req.CreatedAtTo
	}

	countQuery := `SELECT count(1) FROM sales WHERE true ` + filter

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
				"shop_assistant_id", 
				"cashier_id", 
				"payment_type", 
				"status",
				"client_name",

			    created_at, 
			    updated_at 
		FROM sales 
		WHERE true` + filter

	// Sorting by price
	query += " ORDER BY price "

	query += " LIMIT :limit OFFSET :offset"
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
	payment_type := sql.NullString{}
	status := sql.NullString{}

	for rows.Next() {
		var sale sale_service.Sale

		err = rows.Scan(
			&sale.Id,
			&sale.Branch_Id,
			&sale.ShopAssistant_Id,
			&sale.CashierId,
			&payment_type,
			&status,
			&sale.ClientName,
			// float64(sale.Price),
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning sale err: %w", err)
		}

		if createdAt.Valid {
			sale.CreatedAt = createdAt.String
		}
		if updatedAt.Valid {
			sale.UpdatedAt = updatedAt.String
		}
		if payment_type.Valid {
			sale.PaymentType = payment_type.String
		}
		if status.Valid {
			sale.Status = status.String
		}

		resp.Sales = append(resp.Sales, &sale)
	}

	return &resp, nil
}

func (b *saleRepo) DeleteSale(c context.Context, req *sale_service.IdRequest) (resp string, err error) {
	query := `
			DELETE 
				FROM sales 
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
		return "", fmt.Errorf("sale with ID %s not found", req.Id)

	}

	return fmt.Sprintf("sale with ID %s deleted", req.Id), nil
}
