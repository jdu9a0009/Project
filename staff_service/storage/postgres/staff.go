package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"staff_service/pkg/helper"

	staff_service "staff_service/genproto"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type staffRepo struct {
	db *pgxpool.Pool
}

func NewStaff(db *pgxpool.Pool) *staffRepo {
	return &staffRepo{
		db: db,
	}
}

func (b *staffRepo) CreateStaff(c context.Context, req *staff_service.CreateStaffRequest) (string, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO "staffs"(
			"id", 
			"branch_id", 
			"tariff_id", 
			"type_id", 
			"name", 
			"login",
			"password",
			"phone",
			"created_at",
			"updated_at")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
	`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.BranchId,
		req.TariffId,
		req.StaffType,
		req.Name,
		req.Login,
		req.Password,
		req.Phone,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create staff: %w", err)
	}

	return id, nil
}
func (b *staffRepo) GetStaff(c context.Context, req *staff_service.IdRequest) (resp *staff_service.Staff, err error) {
	query := `
				SELECT 
					id, 
					"branch_id", 
					"tariff_id", 
					"type_id", 
					"name", 
					"login",
					"password",
					"phone",
					"created_at", 
					"updated_at" 
				FROM staffs 
				WHERE id=$1`

	var (
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	staff := staff_service.Staff{}
	err = b.db.QueryRow(context.Background(), query, req.Id).Scan(
		&staff.Id,
		&staff.BranchId,
		&staff.TariffId,
		&staff.StaffType,
		&staff.Name,
		&staff.Login,
		&staff.Password,
		&staff.Phone,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("staff not found")
		}
		return nil, fmt.Errorf("failed to get staff: %w", err)
	}
	staff.CreatedAt = createdAt.String
	staff.UpdatedAt = updatedAt.String

	return &staff, nil
}

func (b *staffRepo) UpdateStaff(c context.Context, req *staff_service.UpdateStaffRequest) (string, error) {

	query := `
				UPDATE staffs 
				SET  
					branch_id = $1, 
					tariff_id = $2, 
					type_id = $3, 
					name=$4,
					login=$5,
					password=$6,
					phone=$7,
					balance=$8,
					updated_at = NOW() 
				WHERE id = $9 RETURNING id`

	result, err := b.db.Exec(
		context.Background(),
		query,
		req.BranchId,
		req.TariffId,
		req.StaffType,
		req.Name,
		req.Login,
		req.Password,
		req.Phone,
		req.Balance,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("failed to update staff: %w", err)
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("staff with ID %s not found", req.Id)
	}

	return fmt.Sprintf("staff with ID %s updated", req.Id), nil
}

func (b *staffRepo) GetAllStaff(c context.Context, req *staff_service.GetAllStaffRequest) (*staff_service.GetAllStaffResponse, error) {
	var (
		resp   staff_service.GetAllStaffResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND (name ILIKE '%' || :search || '%' OR type_id ILIKE '%' || :search || '%' OR branch_id ILIKE '%' || :search || '%')"
		params["search"] = req.Search
	}

	if req.BalanceFrom != 0 {
		filter += " AND balance >= :balanceFrom "
		params["balanceFrom"] = req.BalanceFrom
	}

	if req.BalanceTo != 0 {
		filter += " AND balance <= :balanceTo "
		params["balanceTo"] = req.BalanceTo
	}

	countQuery := `SELECT count(1) FROM staffs WHERE true ` + filter

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
			branch_id, 
			tariff_id, 
			type_id, 
			name, 
			login,
			password,
			phone,
			balance,
			created_at, 
			updated_at 
		FROM staffs 
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
		var staff staff_service.Staff

		err = rows.Scan(
			&staff.Id,
			&staff.BranchId,
			&staff.TariffId,
			&staff.StaffType,
			&staff.Name,
			&staff.Login,
			&staff.Password,
			&staff.Phone,
			&staff.Balance,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning staff err: %w", err)
		}

		staff.CreatedAt = createdAt.String
		staff.UpdatedAt = updatedAt.String

		resp.Staffs = append(resp.Staffs, &staff)
	}

	return &resp, nil
}
func (b *staffRepo) DeleteStaff(c context.Context, req *staff_service.IdRequest) (resp string, err error) {
	query := `
			DELETE 
				FROM staffs 
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

// func (b *staffRepo) GetTopStaff(c context.Context, req *staff_service.GetTopStaffRequest) (*staff_service.GetTopStaffResponse, error) {
// 	var (
// 		resp   staff_service.GetTopStaffResponse
// 		err    error
// 		query  string
// 		params []interface{}
// 	)

// 	// Determine the staff type
// 	switch req.Type {
// 	case staff_service.StaffType_ST_Cashier:
// 		query = `
// 			SELECT
// 				s.name,
// 				b.name AS branch,
// 				SUM(sa.price) AS earned_sum
// 			FROM sales sa
// 			INNER JOIN staffs s ON s.id = sa.cashier_id
// 			INNER JOIN branches b ON b.id = sa.branch_id
// 			WHERE sa.created_at >= $1 AND sa.created_at <= $2
// 			GROUP BY s.name, b.name
// 			ORDER BY earned_sum DESC
// 			LIMIT $3`
// 	case staff_service.StaffType_ST_ShopAssistant:
// 		query = `
// 			SELECT
// 				s.name,
// 				b.name AS branch,
// 				SUM(sa.price) AS earned_sum
// 			FROM sales sa
// 			INNER JOIN staffs s ON s.id = sa.shop_assistant_id
// 			INNER JOIN branches b ON b.id = sa.branch_id
// 			WHERE sa.created_at >= $1 AND sa.created_at <= $2
// 			GROUP BY s.name, b.name
// 			ORDER BY earned_sum DESC
// 			LIMIT $3`
// 	default:
// 		return nil, fmt.Errorf("invalid staff type")
// 	}

// 	params = append(params, req.StartDate, req.EndDate, req.Limit)

// 	rows, err := b.db.Query(c, query, params...)
// 	if err != nil {
// 		return nil, fmt.Errorf("error while executing query: %w", err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var (
// 			name      string
// 			branch    string
// 			earnedSum int64
// 			topStaff  staff_service.TopStaff
// 		)

// 		err = rows.Scan(&name, &branch, &earnedSum)
// 		if err != nil {
// 			return nil, fmt.Errorf("error while scanning row: %w", err)
// 		}

// 		topStaff.Name = name
// 		topStaff.Branch = branch
// 		topStaff.EarnedSum = earnedSum

// 		resp.TopStaff = append(resp.TopStaff, &topStaff)
// 	}

// 	return &resp, nil
// }
