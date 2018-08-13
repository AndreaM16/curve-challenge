package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/andream16/curve-challenge/api/handler"
	"github.com/andream16/curve-challenge/internal/configuration"
	internalPsql "github.com/andream16/curve-challenge/internal/psql"
	pkgPsql "github.com/andream16/curve-challenge/pkg/psql"
)

func main() {

	cfg, cfgErr := configuration.Get()
	if cfgErr != nil {
		log.Fatal(cfgErr)
	}

	svc, svcErr := pkgPsql.New(cfg)
	if svcErr != nil {
		log.Fatal(cfgErr)
	}

	if createTablesErr := internalPsql.CreateTables(svc); createTablesErr != nil {
		log.Fatal(createTablesErr)
	}

	r := mux.NewRouter()

	// Create User
	r.HandleFunc("/api/users", handler.CreateUser(svc)).Methods("POST")
	// Get Account aka both balances for easiness
	r.HandleFunc("/api/accounts", handler.GetPaymentAccount(svc)).Methods("GET")
	// Top up an account, this does not write a transaction since handling an external source would be quite tricky
	r.HandleFunc("/api/top-up", handler.TopUp(svc)).Methods("POST")
	// Performs a payment
	r.HandleFunc("/api/pay", handler.Pay(svc)).Methods("POST")

	http.Handle("/", r)

	fmt.Printf("Listening on port %v . . .", cfg.Server.Port)

	log.Fatal(http.ListenAndServe(cfg.Server.Port, r))

}
