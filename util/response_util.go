package util

import (
	"bank-rest-api/model"
	"encoding/json"
	"net/http"
)

func CreateResponse(w http.ResponseWriter, r model.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)

	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		return
	}
}
