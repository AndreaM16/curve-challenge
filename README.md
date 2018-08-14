[![Go Report Card](https://goreportcard.com/badge/github.com/AndreaM16/curve-challenge)](https://goreportcard.com/report/github.com/AndreaM16/curve-challenge)
[![Apache V2 License](http://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/andream16/aws-sdk-go-bindings/blob/master/LICENSE.txt)

# curve-challenge

Small project for simulating payments via credit cards.

## Actors

There are two main entities, users and merchants. 
### A user can:
 - Create a new card
 - TopUp one of his cards
 - Pay a merchant
 
### A merchant can:
  - Capture an authorized amount of money from a user's card
  - Refund a user for a maximum of the captured amount
  - Revert a given amount of the original authorized one
 
### Extra

A transaction system has been implemented. In the latter every action performed is reported in the format like follows
```
// Transaction embeds all transaction information
type Transaction struct {
	// ID is the transaction unique identifier
	ID string `json:"ID"`
	// Receiver is receiver's unique identifier
	Receiver string `json:"receiver"`
	// Sender is receiver's unique identifier
	Sender string `json:"sender"`
	// Amount is the transaction amount
	Amount float64 `json:"amount"`
	// Date is the TS of the action
	Date string `json:"date"`
	// Type is the transaction type
	Type string `json:"type"`
}
```
Where transaction type can vary in:
```
const (

	// EXTERNAL is used to set a default sender for Top Ups
	EXTERNAL = "c9e35256-e831-49c8-8471-164e17a66e31"
	// TOPUP is used for top up actions
	TOPUP = "TOPUP"
	// PAYMENT is used for payment actions
	PAYMENT = "PAYMENT"
	// CAPTURE is used for capture actions
	CAPTURE = "CAPTURE"
	// REFUND is used for refund actions
	REFUND = "REFUND"
	// REVERT is used for revert actions
	REVERT = "REVERT"
)
```

For the TopUps has been generated a default external Merchant since I haven't covered any external actor.

Through such transaction system is possible for the user to know their available and marked balances, which payments they performed and to whom, the location where they spent such money (since each merchant has a sample location).
 
## What's missing

 - Any user can perform any action by having the right card, user, authorization and merchant IDs. No authentication has been implemented for the purposes of this project. 
 - Full test Coverage
 - Better error handling expecially for the postgres part and for Refund/Revert actions
 - Better code reuse
 
## What you need to run the project

 - go
 - postgresql (create a table called `curve`)
 - docker & docker-compose (optional)
 
## Work in local:
 - Install [dep](https://github.com/golang/dep)
 - run `dep ensure`
 - Edit `internal/configuration/configuration.json` by changing `"ENVIRONMENT" : "production"` to `"ENVIRONMENT" : "development"`
 - run tests with `go test ./...`

N.B. Make sure to start `postgresql` and create a table called `curve` with default settings.

## Deploy with docker
 - Install Docker and Docker-Compose
 - Make sure that you have `"ENVIRONMENT" : "production"` in `internal/configuration/configuration.json`
 - run `docker-compose` up
 
N.B. make sure you have port `8000` and `5432` available.

## Endpoints

### Create User:

`POST /api/users` -> `201` with an empty body
 
### Create Merchant:

`POST /api/merchants` -> `201` with a body like 
```
{
	"name" : "merchant100",
	"location" : "london"
}
```

### Create a card:

`POST /api/cards` -> `201` with a body like 
```
{
	"owner" : "0ca41fa2-4fee-4da8-ae2c-e3bea5ae56af",
	"name" : "mastercard3"
}
```
Where `owner` must be an already existing user.

### TopUp a card:

`POST /api/transactions/top-up` -> `201` with a body like 
```
{
	"card" : "8868e147-cfe4-465d-a365-93f7cafd07f7",
	"amount" : 100.0
}
```
Where `card` must be an already existing card.

### Payment:

`POST /api/transactions/payment` -> `201` with a body like 
```
{
	"sender" : "0ca41fa2-4fee-4da8-ae2c-e3bea5ae56af",
	"receiver" : "86c07840-aa72-456a-8bdb-4da6ae622564",
	"card" : "8868e147-cfe4-465d-a365-93f7cafd07f7",
	"amount" : 100.0
}
```
Where `card` must be an already existing card and owned by `sender` and `receiver` must be an existing `merchant`.

### Capture:

`POST /api/transactions/capture` -> `201` with a body like 
```
{
	"merchant" : "18b228c9-bf03-41e3-b7d6-4e875668a03b",
	"authorization" : "56ae9103-e91a-45df-8086-f28c7b98e1e3",
	"amount" : 10.0
}
```
Where `authorization` must be an already existing authorization and `merchant` an existing merchant.

### Refund:

`POST /api/transactions/refund` -> `201` with a body like 
```
{
	"amount" : 8.0,
	"authorization" : "56ae9103-e91a-45df-8086-f28c7b98e1e3"
}
```
Where `authorization` must be an already existing authorization and `amount` the amount to be refunded.

### Revert:

`POST /api/transactions/revert` -> `201` with a body like 
```
{
	"amount" : 8.0,
	"authorization" : "56ae9103-e91a-45df-8086-f28c7b98e1e3"
}
```
Where `authorization` must be an already existing authorization and `amount` the amount to be refunded.

## Postman collection:

Postman collection to perform example calls is available [here](https://www.getpostman.com/collections/a1fab754cc06045e95cb).
