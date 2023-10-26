package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"staff_service/pkg/helper"
	"time"

	staff_service "staff_service/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type staffTarifRepo struct {
	db *pgxpool.Pool
}

func NewStaffTarif(db *pgxpool.Pool) *staffTarifRepo {
	return &staffTarifRepo{
		db: db,
	}
}

func (b *staffTarifRepo) CreateStaffTarif(c context.Context, req *staff_service.CreateStaffTarifRequest) (string, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO "stafftarifs"(
			"id", 
			"name", 
			"type", 
			"amountforcash", 
			"amountforcard",
			"created_at",
			"updated_at")
		VALUES ($1, $2, $3, $4, $5,NOW(), NOW())
	`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Type,
		req.AmountForCash,
		req.AmountForCard,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create staff Tarif: %w", err)
	}

	return id, nil
}
func (b *staffTarifRepo) GetStaffTarif(c context.Context, req *staff_service.IdRequest) (resp *staff_service.StaffTarif, err error) {
	query := `
				SELECT  
					"id", 
					"name", 
					"type", 
					"amountforcash", 
					"amountforcard",
					"created_at",
					"updated_at"
				FROM stafftarifs 
				WHERE id=$1`

	var (
		createdAt time.Time
		updatedAt sql.NullTime
	)

	tarif := staff_service.StaffTarif{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&tarif.Id,
		&tarif.Name,
		&tarif.Type,
		&tarif.AmountForCash,
		&tarif.AmountForCard,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("staffTarif not found")
		}
		return nil, fmt.Errorf("failed to get staff Tarif: %w", err)
	}
	tarif.CreatedAt = createdAt.Format(time.RFC3339)
	tarif.UpdatedAt = updatedAt.Time.Format(time.RFC3339)

	return &tarif, nil
}

func (b *staffTarifRepo) UpdateStaffTarif(c context.Context, req *staff_service.UpdateStaffTarifRequest) (string, error) {

	query := `
				UPDATE stafftarifs 
				SET  
					name = $1, 
					type = $2, 
					amountforcash = $3, 
					amountforcard=$4,
					updated_at = NOW() 
				WHERE id = $5 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Name,
		req.Type,
		req.AmountForCash,
		req.AmountForCard,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update staff tarif: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("staff with ID:  %s not found", req.Id)
	}

	return fmt.Sprintf("staff with ID %s updated", req.Id), nil
}

func (b *staffTarifRepo) GetAllStaffTarif(c context.Context, req *staff_service.GetAllStaffTarifRequest) (*staff_service.GetAllStaffTarifResponse, error) {
	var (
		resp   staff_service.GetAllStaffTarifResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM stafftarifs WHERE true ` + filter

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
		"type", 
		"amountforcash", 
		"amountforcard",
		"created_at",
		"updated_at"
		FROM stafftarifs 
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
		var tarif staff_service.StaffTarif

		err = rows.Scan(
			&tarif.Id,
			&tarif.Name,
			&tarif.Type,
			&tarif.AmountForCash,
			&tarif.AmountForCard,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning staff err: %w", err)
		}

		tarif.CreatedAt = createdAt.Format(time.RFC3339)
		tarif.UpdatedAt = updatedAt.Format(time.RFC3339)

		resp.StaffTarifs = append(resp.StaffTarifs, &tarif)
	}

	return &resp, nil
}
func (b *staffTarifRepo) DeleteStaffTarif(c context.Context, req *staff_service.IdRequest) (resp string, err error) {
	query := `
			DELETE 
				FROM stafftarifs 
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
		return "", fmt.Errorf("staff with ID %s not found", req.Id)

	}

	return fmt.Sprintf("staff with ID %s deleted", req.Id), nil
}
