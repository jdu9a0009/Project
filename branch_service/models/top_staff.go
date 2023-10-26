package models

type TopStaff struct {
	Name       string  `json:"name"`
	Branch     string  `json:"branch"`
	Total_Sum  float64 `json:"total_sum"`
	Staff_type string  `json:"staff_type"`
}
type TopStaffRequest struct {
	FromDate string `json:"from_date"`
	Todate   string `json:"to_date"`
}

type TopStaffResponse struct {
	TopStaff []*TopStaff
}
