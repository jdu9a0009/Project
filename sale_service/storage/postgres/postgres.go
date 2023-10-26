package postgres

import (
	"context"
	"fmt"
	"sale_service/config"
	"sale_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type strg struct {
	db               *pgxpool.Pool
	sale             *saleRepo
	staffTransaction *staffTransactionRepo
	bpTransaction    *BpTransactionRepo
	saleProduct      *saleProductRepo
}

func NewStorage(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDatabase,
		),
	)

	if err != nil {
		fmt.Println("ParseConfig:", err.Error())
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		fmt.Println("ConnectConfig:", err.Error())
		return nil, err
	}

	return &strg{
		db: pool,
	}, nil
}

func (d *strg) Sale() storage.SaleI {
	if d.sale == nil {
		d.sale = NewSale(d.db)
	}
	return d.sale
}

func (d *strg) StaffTransaction() storage.StaffTransactionI {
	if d.staffTransaction == nil {
		d.staffTransaction = NewStaffTransaction(d.db)
	}
	return d.staffTransaction
}

func (d *strg) BpTransaction() storage.BpTransactionI {
	if d.bpTransaction == nil {
		d.bpTransaction = NewBpTransaction(d.db)
	}
	return d.bpTransaction
}

func (d *strg) SaleProduct() storage.SaleProductI {
	if d.saleProduct == nil {
		d.saleProduct = NewSaleProduct(d.db)
	}
	return d.saleProduct
}
