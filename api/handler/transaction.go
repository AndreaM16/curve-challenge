package handler

import (
	"github.com/andream16/curve-challenge/api/middleware"
	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
	"net/http"
)

// TopUp adds money to an user's card
func TopUp(svc *psql.PSQL) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var topUp model.TopUp

		unmarshalErr := UnmarshalBody(r, &topUp)
		if unmarshalErr != nil {
			HandleError(w, unmarshalErr)
			return
		}

		out, err := middleware.TopUp(svc, topUp)
		if err != nil {
			HandleError(w, err)
			return
		}

		CreatedResponse(w, out)

		return

	}
}
