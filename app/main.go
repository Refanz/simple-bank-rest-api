package main

import (
	"bank-rest-api/controller"
	"bank-rest-api/middleware"
	"bank-rest-api/repository"
	"bank-rest-api/service"
	"bank-rest-api/util"
	"net/http"
)

var (
	customerRepo       = repository.CustomerRepository{FilePath: "../data/data.json"}
	customerService    = service.CustomerService{CustomerRepository: &customerRepo}
	customerController = controller.CustomerController{CustomerService: customerService}

	authService    = service.AuthService{CustomerService: &customerService}
	authController = controller.AuthController{AuthService: &authService}

	authMiddleware = middleware.AuthMiddleware{CustomerService: &customerService}
)

func login(w http.ResponseWriter, r *http.Request) {
	util.CreateResponse(w, authController.Login(r))
}

func getCustomerByUsername(w http.ResponseWriter, r *http.Request) {
	util.CreateResponse(w, customerController.GetCustomerByUsername(r))
}

func main() {

	mux := http.NewServeMux()

	getCustomerByUsernameHandler := http.HandlerFunc(getCustomerByUsername)
	loginHandler := http.HandlerFunc(login)

	mux.Handle("/customers", authMiddleware.AuthMiddleware(getCustomerByUsernameHandler))
	mux.Handle("/login", loginHandler)

	println("Server is running..")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}
}
