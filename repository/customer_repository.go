package repository

import (
	"bank-rest-api/model"
)

type CustomerRepository struct {
	BankRepository
}

func (r *CustomerRepository) ReadCustomers() []model.Customer {
	return r.BankRepository.ReadBankDataFromJsonFile().Customer
}
