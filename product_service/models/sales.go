package models

import "time"

type CreateSale struct {
	ClientName        string
	Branch_id         string
	Shop_asisstant_id string
	Cashier_id        string
	Price             float64
	Payment_Type      string //1 for card ,2 for cash
	Status            string //1 for succes ,2 for  cancel
}

type Sales struct {
	Id                string
	ClientName        string
	Branch_id         string
	Shop_asisstent_id string
	Cashier_id        string
	Price             float64
	Payment_Type      string
	Status            string
	CreatedAt         time.Time
	Updated_at        time.Time
}

type GetAllSaleRequest struct {
	Page        int
	Limit       int
	Client_name string
}

type GetAllSaleResponse struct {
	Sales []Sales
	Count int
}

type SaleTopBranch struct {
	Day         string
	BranchId    string
	SalesAmount float64
}

type SaleCountSumBranch struct {
	BranchId    string
	Count       int
	SalesAmount float64
}
