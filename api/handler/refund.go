package handler

import (
	"net/http"

	"github.com/andream16/curve-challenge/pkg/psql"
	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/api/middleware"
)

// Refund lets a merchant refund a customer
func Refund(svc *psql.PSQL) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var tx model.Transaction

		unmarshalErr := UnmarshalBody(r, &tx)
		if unmarshalErr != nil {

			w.WriteHeader(http.StatusBadRequest)

			resp := NewResponse(InvalidParameters)

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return
		}

		sender, senderErr := middleware.GetUser(svc, tx.Sender)
		if senderErr != nil {

			w.WriteHeader(http.StatusInternalServerError)

			resp := NewResponse(senderErr.Error())

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return

		}

		if sender.Type != middleware.MerchantType {
			w.WriteHeader(http.StatusInternalServerError)

			resp := NewResponse(ForbiddenAction)

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return
		}

		senderAccount, senderAccountErr := middleware.GetPaymentAccount(svc, tx.Sender)
		if senderAccountErr != nil {

			w.WriteHeader(http.StatusInternalServerError)

			resp := NewResponse(senderAccountErr.Error())

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return

		}

		if senderAccount.AvailableBalance < tx.Amount {

			w.WriteHeader(http.StatusBadRequest)

			resp := NewResponse(InsufficientFounds)

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return

		}

		receiver, receiverErr := middleware.GetUser(svc, tx.Receiver)
		if receiverErr != nil {

			w.WriteHeader(http.StatusInternalServerError)

			resp := NewResponse(receiverErr.Error())

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return

		}

		if receiver.Type != middleware.ClassicType {

			w.WriteHeader(http.StatusBadRequest)

			resp := NewResponse(ForbiddenAction)

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return

		}

		

		newTx, newTxErr := model.NewTransaction(tx.Receiver, tx.Sender, tx.Type, tx.Amount)
		if newTxErr != nil {

			w.WriteHeader(http.StatusBadRequest)

			resp := NewResponse(newTxErr.Error())

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return

		}

		createTxErr := middleware.CreateTransaction(svc, newTx)
		if createTxErr != nil {

			w.WriteHeader(http.StatusBadRequest)

			resp := NewResponse(createTxErr.Error())

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return

		}

	}
}
