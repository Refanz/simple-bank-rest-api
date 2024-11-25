package service

import (
	"bank-rest-api/constant"
	"bank-rest-api/model"
	"bank-rest-api/util"
	"errors"
	"net/http"
)

type AuthService struct {
	*CustomerService
}

func (s *AuthService) Login(loginRequest model.LoginRequest) error {
	customer := s.CustomerService.GetCustomerByUsername(model.Customer{Username: loginRequest.Username})

	if customer == nil {
		return errors.New(constant.UserNotFoundErrorMessage)
	}

	if customer.Password != loginRequest.Password {
		return errors.New(constant.CredentialsErrorMessage)
	}

	return nil
}

func (s *AuthService) Logout(w http.ResponseWriter, r *http.Request) error {
	cookie := util.GetCookie(r)

	if cookie == nil {
		return errors.New(constant.CookieErrorMessage)
	}

	util.DeleteCookie(w)

	return nil
}
