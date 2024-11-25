package service

import (
	"bank-rest-api/model"
	"bank-rest-api/repository"
	"bank-rest-api/util"
	"time"
)

type HistoryService struct {
	*repository.HistoryRepository
}

func (r *HistoryService) CreateHistory(historyRequest model.HistoryRequest) {
	r.HistoryRepository.AddHistory(model.History{
		Id:          util.GetCurrentTimeInMillis(),
		Description: historyRequest.Description,
		Timestamp:   time.Now(),
	})
}
