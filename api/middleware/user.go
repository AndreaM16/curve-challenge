package middleware

import (
	"github.com/andream16/curve-challenge/api/model"
	internalError "github.com/andream16/curve-challenge/internal/error"
	"github.com/andream16/curve-challenge/pkg/psql"
)

const (
	MerchantType = "merchant"
	ClassicType  = "classic"
)

func CreateUser(svc *psql.PSQL, user *model.User) error {

	userLocation := user.Location
	userType := user.Type

	if len(userType) == 0 {
		return internalError.Format(UserType, ErrMissingRequiredParameter)
	}
	if userType != MerchantType && userType != ClassicType {
		return internalError.Format(UserType, ErrBadUserType)
	}
	if len(userLocation) == 0 && userType == MerchantType {
		return internalError.Format(UserLocation, ErrMissingRequiredParameter)
	}

	account, accountErr := CreateCard(svc)
	if accountErr != nil {
		return accountErr
	}

	userEntry, userEntryErr := model.NewUser(*account, userLocation, userType)
	if userEntryErr != nil {
		return userEntryErr
	}

	query := `INSERT INTO users (ID, type, payment_account, location) VALUES ($1, $2, $3, $4)`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, insertError := stmt.Exec(
		userEntry.ID,
		userEntry.Type,
		userEntry.Card,
		userEntry.Location,
	)
	if insertError != nil {
		return insertError
	}

	return nil

}

// GetUser retrieves a user from the database
func GetUser(svc *psql.PSQL, userID string) (*model.User, error) {

	if len(userID) == 0 {
		return nil, internalError.Format(UserID, ErrMissingRequiredParameter)
	}

	var user model.User

	query := `SELECT ID,payment_account,type,location FROM users WHERE ID = $1`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	selectErr := stmt.QueryRow(userID).Scan(&user.ID, &user.Card, &user.Type, &user.Location)
	if selectErr != nil {
		return nil, selectErr
	}

	return &user, nil

}
