package handler

import (
	"net/http"

	"github.com/andream16/curve-challenge/pkg/psql"
	"github.com/andream16/curve-challenge/api/middleware"
	"github.com/andream16/curve-challenge/api/model"
	"encoding/json"
)

// TopUp adds an amount of money to an user's payment account
func TopUp(svc *psql.PSQL) func(w http.ResponseWriter, r *http.Request) {
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

		var topUp model.TopUp

		unmarshalErr := UnmarshalBody(r, &topUp)
		if unmarshalErr != nil {

			w.WriteHeader(http.StatusBadRequest)

			resp := NewResponse(InvalidParameters)

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

		account = account.SetAvailableBalance(account.AvailableBalance + topUp.Amount)

		updateErr := middleware.UpdateCard(svc, account)
		if updateErr != nil {

			w.WriteHeader(http.StatusInternalServerError)

			resp := NewResponse(updateErr.Error())

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
