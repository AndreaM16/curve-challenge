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

	account, accountErr := CreatePaymentAccount(svc)
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
		userEntry.PaymentAccount,
		userEntry.Location,
	)
	if insertError != nil {
		return insertError
	}

	return nil

}
