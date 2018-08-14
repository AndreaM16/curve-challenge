package middleware

import (
	"github.com/andream16/curve-challenge/api/model"
	"github.com/andream16/curve-challenge/pkg/psql"
)

// CreateCard creates a new card for a given user
func CreateCard(svc *psql.PSQL, card *model.Card) (*model.Card, error) {

	newCard := model.NewCard(card.Owner, card.Name)

	query := `INSERT INTO cards (ID,name,owner,available_balance,marked_balance) VALUES ($1, $2, $3, $4, $5)`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, insertError := stmt.Exec(
		newCard.ID,
		newCard.Name,
		newCard.Owner,
		newCard.AvailableBalance,
		newCard.MarkedBalance,
	)
	if insertError != nil {
		return nil, insertError
	}

	return newCard, nil

}

// GetCard gets a card given its ID
func GetCard(svc *psql.PSQL, ID string) (*model.Card, error) {

	query := `SELECT ID,name,owner,available_balance,marked_balance FROM cards WHERE ID = $1`

	var card model.Card

	stmt, err := svc.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	selectErr := stmt.QueryRow(ID).Scan(&card.ID, &card.Name, &card.Owner, &card.AvailableBalance, &card.MarkedBalance)
	if selectErr != nil {
		return nil, selectErr
	}

	return &card, nil

}

// UpdateCard updates a card
func UpdateCard(svc *psql.PSQL, card *model.Card) (*model.Card, error) {

	query := `UPDATE cards SET name=$2, owner=$3, available_balance=$4, marked_balance=$5 WHERE ID = $1`

	stmt, err := svc.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, updateErr := stmt.Exec(&card.ID, &card.Name, &card.Owner, &card.AvailableBalance, &card.MarkedBalance)
	if updateErr != nil {
		return nil, updateErr
	}

	return card, nil

}
