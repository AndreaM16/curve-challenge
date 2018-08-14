package handler

import (
	"net/http"

	"github.com/andream16/curve-challenge/api/middleware"
	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
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

		out, err := middleware.TopUp(svc, &topUp)
		if err != nil {
			HandleError(w, err)
			return
		}

		CreatedResponseWithBody(w, out)

		return

	}
}

// Pay allows a user to send money to a merchant
func Pay(svc *psql.PSQL) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var payment model.Payment

		unmarshalErr := UnmarshalBody(r, &payment)
		if unmarshalErr != nil {
			HandleError(w, unmarshalErr)
			return
		}

		err := middleware.Pay(svc, &payment)
		if err != nil {
			HandleError(w, err)
			return
		}

		CreatedResponse(w)

		return

	}
}

// Capture allows a merchant to capture money by moving money from user's marked balance and merchant's available one
func Capture(svc *psql.PSQL) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var capture model.Capture

		unmarshalErr := UnmarshalBody(r, &capture)
		if unmarshalErr != nil {
			HandleError(w, unmarshalErr)
			return
		}

		err := middleware.Capture(svc, &capture)
		if err != nil {
			HandleError(w, err)
			return
		}

		CreatedResponse(w)

		return

	}
}

// Refund allows a merchant refund a variable amount to the user
func Refund(svc *psql.PSQL) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var refund model.Refund

		unmarshalErr := UnmarshalBody(r, &refund)
		if unmarshalErr != nil {
			HandleError(w, unmarshalErr)
			return
		}

		err := middleware.Refund(svc, &refund)
		if err != nil {
			HandleError(w, err)
			return
		}

		CreatedResponse(w)

		return

	}
}

// Revert allows a merchant reversing a variable amount
func Revert(svc *psql.PSQL) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var reverse model.Revert

		unmarshalErr := UnmarshalBody(r, &reverse)
		if unmarshalErr != nil {
			HandleError(w, unmarshalErr)
			return
		}

		err := middleware.Revert(svc, &reverse)
		if err != nil {
			HandleError(w, err)
			return
		}

		CreatedResponse(w)

		return

	}
}
