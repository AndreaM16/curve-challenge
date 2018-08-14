package handler

import (
	"net/http"

	"github.com/andream16/curve-challenge/api/middleware"
	"github.com/andream16/curve-challenge/pkg/psql"
	"github.com/andream16/curve-challenge/api/model"
)

// CreateMerchant creates a new merchant
func CreateMerchant(svc *psql.PSQL) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var merchant model.Merchant

		unmarshalErr := UnmarshalBody(r, &merchant)
		if unmarshalErr != nil {
			HandleError(w, unmarshalErr)
			return
		}

		out, err := middleware.CreateMerchant(svc, &merchant)
		if err != nil {
			HandleError(w, err)
			return
		}

		CreatedResponse(w, out)

		return

	}
}
