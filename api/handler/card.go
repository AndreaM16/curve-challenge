package handler

import (
	"net/http"

	"github.com/andream16/curve-challenge/api/middleware"
	"github.com/andream16/curve-challenge/pkg/psql"
	"github.com/andream16/curve-challenge/api/model"
)

// CreateCard creates a new card for a user
func CreateCard(svc *psql.PSQL) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var card model.Card

		unmarshalErr := UnmarshalBody(r, &card)
		if unmarshalErr != nil {
			HandleError(w, unmarshalErr)
			return
		}

		out, err := middleware.CreateCard(svc, &card)
		if err != nil {
			HandleError(w, err)
			return
		}

		CreatedResponse(w, out)
		return

	}
}
