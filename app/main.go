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
	customerRepo    = repository.CustomerRepository{FilePath: "../data/data.json"}
	customerService = service.CustomerService{CustomerRepository: &customerRepo}
	authService     = service.AuthService{CustomerService: &customerService}

	customerController = controller.CustomerController{CustomerService: customerService}
	authController     = controller.AuthController{AuthService: &authService}

	authMiddleware = middleware.AuthMiddleware{CustomerService: &customerService}
)

func login(w http.ResponseWriter, r *http.Request) {
	util.CreateResponse(w, authController.Login(w, r))
}

func logout(w http.ResponseWriter, r *http.Request) {
	util.CreateResponse(w, authController.Logout(w, r))
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	util.CreateResponse(w, customerController.GetAllCustomers())
}

func main() {

	mux := http.NewServeMux()

	getAllCustomersHandler := http.HandlerFunc(getAllCustomers)
	loginHandler := http.HandlerFunc(login)
	logoutHandler := http.HandlerFunc(logout)

	mux.Handle("/customers", authMiddleware.AuthMiddleware(getAllCustomersHandler))
	mux.Handle("/login", loginHandler)
	mux.Handle("/logout", authMiddleware.AuthMiddleware(logoutHandler))

	println("Server is running..")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}
}
