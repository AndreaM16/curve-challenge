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

	r.HandleFunc("/api/users", handler.CreateUser(svc)).Methods("POST")

	http.Handle("/", r)

	fmt.Printf("Listening on port %v . . .", cfg.Server.Port)

	log.Fatal(http.ListenAndServe(cfg.Server.Port, r))

}
