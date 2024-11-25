package model

type Bank struct {
	Customer []Customer `json:"customers"`
	History  []History  `json:"histories"`
}
