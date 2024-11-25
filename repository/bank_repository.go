package repository

import (
	"bank-rest-api/model"
	"encoding/json"
	"os"
)

type BankRepository struct {
	FilePath string
}

func (r *BankRepository) ReadBankDataFromJsonFile() model.Bank {
	data, err := os.ReadFile(r.FilePath)

	if err != nil {
		panic(err)
	}

	var bank model.Bank
	err = json.Unmarshal(data, &bank)

	if err != nil {
		panic(err)
	}

	return bank
}
