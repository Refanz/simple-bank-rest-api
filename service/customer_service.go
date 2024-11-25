package service

import (
	"bank-rest-api/model"
	"bank-rest-api/repository"
)

type CustomerService struct {
	*repository.CustomerRepository
}

func (s *CustomerService) GetAllCustomers() *[]model.CustomerResponse {
	data := s.CustomerRepository.ReadCustomers()

	var customers []model.CustomerResponse

	for _, customer := range data {
		customers = append(customers, model.CustomerResponse{
			Id:       customer.Id,
			Username: customer.Username,
		})
	}

	return &customers
}

func (s *CustomerService) GetCustomerByUsername(customerRequest model.Customer) *model.Customer {
	data := s.CustomerRepository.ReadCustomers()

	for _, customer := range data {
		if customer.Username == customerRequest.Username {
			return &customer
		}
	}

	return nil
}
