package controller

import (
	"bank-rest-api/constant"
	"bank-rest-api/model"
	"bank-rest-api/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type BankController struct {
	*service.BankService
	*service.HistoryService
}

func (c *BankController) Transfer(r *http.Request) model.Response {
	var paymentRequest model.Payment

	err := json.NewDecoder(r.Body).Decode(&paymentRequest)

	if err != nil {

		c.CreateHistory(model.HistoryRequest{Description: err.Error()})

		return model.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	c.CreateHistory(model.HistoryRequest{Description: constant.TransferSuccessMessage})

	return model.Response{
		Status:  http.StatusOK,
		Message: constant.TransferSuccessWithAmountMessage + " " + strconv.Itoa(paymentRequest.Amount),
		Data:    nil,
	}
}
