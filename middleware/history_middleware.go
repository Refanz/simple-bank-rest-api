package middleware

import (
	"bank-rest-api/model"
	"bank-rest-api/service"
	"net/http"
)

type HistoryMiddleware struct {
	*service.HistoryService
}

func (m *HistoryMiddleware) HistoryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		m.CreateHistory(model.HistoryRequest{
			Description: r.URL.Path,
		})

		next.ServeHTTP(w, r)
	})
}
