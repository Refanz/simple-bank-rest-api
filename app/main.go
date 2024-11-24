package main

import (
	"bank-rest-api/controller"
	"bank-rest-api/repository"
	"bank-rest-api/service"
	"net/http"
)

func main() {
	var customerRepo = repository.CustomerRepository{FilePath: "../data/data.json"}
	var customerService = service.CustomerService{CustomerRepository: &customerRepo}
	var customerController = controller.CustomerController{CustomerService: customerService}

	mux := http.NewServeMux()

	mux.HandleFunc("/customers", customerController.GetCustomerByUsername)

	println("Server is running..")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}
}
