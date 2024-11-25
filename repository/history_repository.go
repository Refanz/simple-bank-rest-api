package repository

import (
	"bank-rest-api/model"
	"encoding/json"
	"fmt"
	"os"
)

type HistoryRepository struct {
	BankRepository
}

func (r *HistoryRepository) AddHistory(history model.History) {
	data := r.BankRepository.ReadBankDataFromJsonFile()

	data.History = append(data.History, history)

	updatedHistories, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = os.WriteFile(r.BankRepository.FilePath, updatedHistories, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
