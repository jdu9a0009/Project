package models

import "time"

type CreateTransaction struct {
	Type        string //(withdraw,topup)
	Amount      int
	Source_type string // Sales,Bonus
	Text        string
	Sale_id     string
	Staff_id    string
}

type Transaction struct {
	Id          string
	Type        string //withdraw,topup
	Amount      int
	Source_type string
	Text        string
	Sale_id     string
	Staff_id    string
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetAllTransactionRequest struct {
	Page  int
	Limit int
	Text  string
}

type GetAllTransactionResponse struct {
	Transaction []Transaction
	Count       int
}

type TopWorkerRequest struct {
	Type     string
	FromDate string
	ToDate   string
}

type TopWorkerRespond struct {
	Staff []TopWorker
}
type TopWorker struct {
	BranchName string
	StaffName  string
	StaffType  string
	EarnedSum  int
}
