package model

type Bank struct {
	Customer []Customer `json:"customers"`
	Merchant []Merchant `json:"merchants"`
	History  []History  `json:"histories"`
}
