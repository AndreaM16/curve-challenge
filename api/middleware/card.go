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
