package repository

import (
	model2 "bank-rest-api/internal/model"
	"encoding/json"
	"os"
)

type CustomerRepository struct {
	FilePath string
}

func (r *CustomerRepository) ReadCustomerDataFromJsonFile() []model2.Customer {
	data, err := os.ReadFile(r.FilePath)

	if err != nil {
		panic(err)
	}

	var bank model2.Bank
	err2 := json.Unmarshal(data, &bank)

	if err2 != nil {
		panic(err2)
	}

	return bank.Customer
}
