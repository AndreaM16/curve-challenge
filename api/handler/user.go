package handler

import (
	"net/http"

	"github.com/andream16/curve-challenge/api/middleware"
	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
)

// CreateUser creates a new user
func CreateUser(svc *psql.PSQL) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var user model.User

		unmarshalErr := UnmarshalBody(r, &user)
		if unmarshalErr != nil {

			w.WriteHeader(http.StatusBadRequest)

			resp := NewResponse(InvalidParameters)

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return
		}

		err := middleware.CreateUser(svc, &user)
		if err != nil {

			w.WriteHeader(http.StatusInternalServerError)

			resp := NewResponse(err.Error())

			b, _ := resp.JsonMarshal()

			w.Write(b)

			return

		}

		w.WriteHeader(http.StatusCreated)

		return

	}
}
