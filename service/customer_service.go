package service

import (
	"bank-rest-api/internal/model"
	"bank-rest-api/internal/repository"
	"encoding/json"
	"net/http"
)

type CustomerService struct {
	CustomerRepository *repository.CustomerRepository
}

func (s *CustomerService) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := s.CustomerRepository.ReadCustomerDataFromJsonFile()

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		panic(err)
	}
}

func (s *CustomerService) GetCustomerByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := s.CustomerRepository.ReadCustomerDataFromJsonFile()

	var customerRequest model.Customer

	err := json.NewDecoder(r.Body).Decode(&customerRequest)

	if err != nil {
		panic(err)
	}

	for _, customer := range data {
		if customer.Username == customerRequest.Username {
			err = json.NewEncoder(w).Encode(customer)

			if err != nil {
				panic(err)
			}

			break
		}
	}
}
