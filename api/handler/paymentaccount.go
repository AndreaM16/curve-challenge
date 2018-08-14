package handler

import (
	"net/http"
	"encoding/json"

	"github.com/andream16/curve-challenge/pkg/psql"
	"github.com/andream16/curve-challenge/api/middleware"
)

// GetCard gets user available balance
func GetCard(svc *psql.PSQL) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		userID := r.URL.Query().Get("user")

		if len(userID) == 0 {

			w.WriteHeader(http.StatusBadRequest)

			resp := NewResponse(MissingParameters)

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return

		}

		account, err := middleware.GetCard(svc, userID)
		if err != nil {

			w.WriteHeader(http.StatusInternalServerError)

			resp := NewResponse(err.Error())

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return

		}

		w.WriteHeader(http.StatusOK)

		b, _ := json.Marshal(account)

		w.Write(b)

		return

	}
}