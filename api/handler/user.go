package handler

import (
	"net/http"

	"github.com/andream16/curve-challenge/api/middleware"
	"github.com/andream16/curve-challenge/pkg/psql"
)

// CreateUser creates a new user
func CreateUser(svc *psql.PSQL) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		user, err := middleware.CreateUser(svc)
		if err != nil {
			HandleError(w, err)
			return
		}

		CreatedResponseWithBody(w, user)
		return

	}
}
