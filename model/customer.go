package model

type Customer struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomerResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}
