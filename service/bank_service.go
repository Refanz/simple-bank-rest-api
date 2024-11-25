package service

import "bank-rest-api/model"

type BankService struct{}

func (s *BankService) Transfer(paymentRequest model.Payment) *model.Payment {
	return &model.Payment{
		Amount: paymentRequest.Amount,
	}
}
