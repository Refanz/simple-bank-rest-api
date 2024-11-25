package controller

import (
	"bank-rest-api/constant"
	"bank-rest-api/model"
	"bank-rest-api/service"
	"bank-rest-api/util"
	"encoding/json"
	"net/http"
)

type AuthController struct {
	*service.AuthService
	*service.HistoryService
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) model.Response {

	var loginRequest model.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		c.CreateHistory(model.HistoryRequest{Description: err.Error()})

		return model.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	err = c.AuthService.Login(loginRequest)

	if err != nil {
		c.CreateHistory(model.HistoryRequest{Description: err.Error()})

		return model.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	var accessToken = util.EncodeCredentials(loginRequest.Username + ":" + loginRequest.Password)

	loginResponse := model.LoginResponse{AccessToken: accessToken}

	util.SetCookie(w, accessToken)

	c.CreateHistory(model.HistoryRequest{Description: loginRequest.Username + " successfully logged in"})

	return model.Response{
		Status:  http.StatusOK,
		Message: constant.LoginSuccessMessage,
		Data:    loginResponse,
	}
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) model.Response {
	err := c.AuthService.Logout(w, r)

	if err != nil {
		c.CreateHistory(model.HistoryRequest{Description: err.Error()})

		return model.Response{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		}
	}

	c.CreateHistory(model.HistoryRequest{Description: constant.LogoutSuccessMessage})

	return model.Response{
		Status:  http.StatusOK,
		Message: constant.LogoutSuccessMessage,
	}
}
