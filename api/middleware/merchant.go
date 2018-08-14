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

// GetMerchant gets a merchant given its ID
func GetMerchant(svc *psql.PSQL, ID string) (*model.Merchant, error) {

	query := `SELECT ID,name,location,balance FROM merchants WHERE ID = $1`

	var merchant model.Merchant

	stmt, err := svc.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	selectErr := stmt.QueryRow(ID).Scan(&merchant.ID, &merchant.Name, &merchant.Location, &merchant.Balance)
	if selectErr != nil {
		return nil, selectErr
	}

	return &merchant, nil

}

// UpdateMerchant updates a merchant
func UpdateMerchant(svc *psql.PSQL, merchant *model.Merchant) error {

	query := `UPDATE merchants SET name=$2, location=$3, balance=$4 WHERE ID = $1`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, updateErr := stmt.Exec(&merchant.ID, &merchant.Name, &merchant.Location, &merchant.Balance)
	if updateErr != nil {
		return updateErr
	}

	return nil

}
