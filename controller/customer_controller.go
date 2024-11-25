package controller

import (
	"bank-rest-api/constant"
	"bank-rest-api/model"
	"bank-rest-api/service"
	"net/http"
)

type CustomerController struct {
	*service.CustomerService
	*service.HistoryService
}

func (c *CustomerController) GetAllCustomers() model.Response {
	customers := c.CustomerService.GetAllCustomers()

	c.CreateHistory(model.HistoryRequest{Description: constant.GetAllCustomersSuccessMessage})

	return model.Response{
		Status:  http.StatusOK,
		Message: constant.GetAllCustomersSuccessMessage,
		Data:    customers,
	}
}
