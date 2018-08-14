package handler

import (
	"net/http"

	"github.com/andream16/curve-challenge/pkg/psql"
	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/api/middleware"
)

// Pay lets a user pay a merchant
func Pay(svc *psql.PSQL) func(w http.ResponseWriter, r *http.Request) {
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

		if sender.Type != middleware.ClassicType {
			w.WriteHeader(http.StatusInternalServerError)

			resp := NewResponse(ForbiddenAction)

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return
		}

		senderAccount, senderAccountErr := middleware.GetCard(svc, tx.Sender)
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

		if receiver.Type != middleware.MerchantType {

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

		newAuth, newAuthErr := model.NewAuthorization(newTx.ID, tx.Amount)
		if newAuthErr != nil {

			w.WriteHeader(http.StatusBadRequest)

			resp := NewResponse(newAuthErr.Error())

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return

		}

		newAuthErr = middleware.CreateAuthorization(svc, newAuth)
		if newAuthErr != nil {

			w.WriteHeader(http.StatusBadRequest)

			resp := NewResponse(newAuthErr.Error())

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return

		}

		senderAccount = senderAccount.SetAvailableBalance(senderAccount.AvailableBalance - tx.Amount).SetMarkedBalance(senderAccount.AvailableBalance + tx.Amount)

		updateErr := middleware.UpdateCard(svc, senderAccount)
		if updateErr != nil {

			w.WriteHeader(http.StatusBadRequest)

			resp := NewResponse(updateErr.Error())

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return

		}

		w.WriteHeader(http.StatusCreated)

		return

	}
}