package middleware

import (
	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
)

// CreateMerchant creates a new merchant
func CreateMerchant(svc *psql.PSQL, merchant *model.Merchant) (*model.Merchant, error) {

	newMerchant := model.NewMerchant(merchant.Name, merchant.Location)

	query := `INSERT INTO merchants (ID,name,location,balance) VALUES ($1, $2, $3, $4)`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, insertError := stmt.Exec(
		newMerchant.ID,
		newMerchant.Name,
		newMerchant.Location,
		newMerchant.Balance,
	)
	if insertError != nil {
		return nil, insertError
	}

	return newMerchant, nil

}
