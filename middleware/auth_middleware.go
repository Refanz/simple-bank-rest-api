package middleware

import (
	"bank-rest-api/constant"
	"bank-rest-api/model"
	"bank-rest-api/service"
	"bank-rest-api/util"
	"net/http"
)

type AuthMiddleware struct {
	CustomerService *service.CustomerService
}

func (m *AuthMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok {
			util.CreateResponse(w, model.Response{
				Status:  http.StatusUnauthorized,
				Message: constant.AuthorizationRequiredMessage,
			})
			return
		}

		customer := m.CustomerService.GetCustomerByUsername(model.Customer{Username: username})

		if customer == nil {
			util.CreateResponse(w, model.Response{
				Status:  http.StatusBadRequest,
				Message: constant.CredentialsErrorMessage,
			})

			return
		}

		isValid := (username == customer.Username) && (password == customer.Password)
		if !isValid {
			util.CreateResponse(w, model.Response{
				Status:  http.StatusBadRequest,
				Message: constant.CredentialsErrorMessage,
			})

			return
		}

		next.ServeHTTP(w, r)
	})
}
