package middleware

import (
	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
)

// CreateUser creates a new user
func CreateUser(svc *psql.PSQL) (*model.User, error) {

	user := model.NewUser()

	query := `INSERT INTO users (ID) VALUES ($1)`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, insertError := stmt.Exec(
		user.ID,
	)
	if insertError != nil {
		return nil, insertError
	}

	return user, nil

}
