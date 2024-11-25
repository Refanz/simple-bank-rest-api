package main

import (
	"bank-rest-api/constant"
	"bank-rest-api/controller"
	"bank-rest-api/middleware"
	"bank-rest-api/repository"
	"bank-rest-api/service"
	"bank-rest-api/util"
	"net/http"
)

var (
	bankRepo     = repository.BankRepository{FilePath: constant.FilePath}
	customerRepo = repository.CustomerRepository{BankRepository: bankRepo}
	historyRepo  = repository.HistoryRepository{BankRepository: bankRepo}

	customerService = service.CustomerService{CustomerRepository: &customerRepo}
	authService     = service.AuthService{CustomerService: &customerService}
	historyService  = service.HistoryService{HistoryRepository: &historyRepo}

	customerController = controller.CustomerController{CustomerService: &customerService, HistoryService: &historyService}
	authController     = controller.AuthController{AuthService: &authService, HistoryService: &historyService}
	bankController     = controller.BankController{BankService: &service.BankService{}, HistoryService: &historyService}

	authMiddleware    = middleware.AuthMiddleware{CustomerService: &customerService}
	historyMiddleware = middleware.HistoryMiddleware{HistoryService: &historyService}
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

func transfer(w http.ResponseWriter, r *http.Request) {
	util.CreateResponse(w, bankController.Transfer(r))
}

func main() {

	mux := http.NewServeMux()

	getAllCustomersHandler := http.HandlerFunc(getAllCustomers)
	loginHandler := http.HandlerFunc(login)
	logoutHandler := http.HandlerFunc(logout)
	transferHandler := http.HandlerFunc(transfer)

	mux.Handle(constant.CustomerApi, middleware.LoggingMiddleware(authMiddleware.AuthMiddleware(getAllCustomersHandler)))
	mux.Handle(constant.LoginApi, middleware.LoggingMiddleware(historyMiddleware.HistoryMiddleware(loginHandler)))
	mux.Handle(constant.LogoutApi, middleware.LoggingMiddleware(authMiddleware.AuthMiddleware(logoutHandler)))
	mux.Handle(constant.PaymentApi, middleware.LoggingMiddleware(authMiddleware.AuthMiddleware(transferHandler)))

	println(constant.ServerRunningMessage)
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}
}
