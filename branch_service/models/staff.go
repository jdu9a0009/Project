package models

type StaffType string

const (
	Cashier       StaffType = "cashier"
	ShopAssistant StaffType = "shop_assistant"
)

type CreateStaff struct {
	BranchId  string    `json:"branch_id"`
	TariffId  string    `json:"tariff_id"`
	StaffType StaffType `json:"staff_type"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
}

type Staff struct {
	Id        string    `json:"id"`
	BranchId  string    `json:"branch_id"`
	TariffId  string    `json:"tariff_id"`
	TypeId    StaffType `json:"type_id"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	// Login    string
	// Password string
	// Phone    string
}

type ExistsReq struct {
	Phone string `json:"phone"`
}
type UpdateBalanceRequest struct {
	SaleId          string
	TransactionType string
	SourceType      string
	Cashier         StaffIdAmount
	ShopAssisstant  StaffIdAmount
	Text            string
}

type StaffIdAmount struct {
	StaffId string
	Amount  float32
}

type StaffTop struct {
	BranchId string    `json:"branch_id"`
	TypeId   StaffType `json:"type_id"`
	Name     string    `json:"name"`
	Money    int       `json:"money"`
}

type ChangeBalance struct {
	Id      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type GetAllStaffRequest struct {
	Page        int       `json:"page"`
	Limit       int       `json:"limit"`
	Type        StaffType `json:"type"`
	Name        string    `json:"name"`
	BalanceFrom float64   `json:"balance_from"`
	BalanceTo   float64   `json:"balance_to"`
}

type GetAllStaff struct {
	Staffes []Staff `json:"staffes"`
	Count   int     `json:"count"`
}
type ChangePassword struct {
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}
type ReqNewPassword struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}
