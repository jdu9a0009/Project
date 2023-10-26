package models

type RegisterRequest struct {
}

type LoginRequest struct {
	Login    string
	Password string
}
type LoginRes struct {
	Token string `json:"token"`
}

type RequestByLogin struct {
	Login string
}
