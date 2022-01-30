package main

import (
	"fmt"
	"github.com/MollenAR/internOzonFintech/cmd/server"
	"github.com/MollenAR/internOzonFintech/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start")

	psqlCredentials, err := configs.PostgresConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	psqldb, err := sqlx.Open("postgres", psqlCredentials)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = psqldb.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	serverAddres, err := configs.AddresConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := server.Run(serverAddres, psqldb); err != http.ErrServerClosed {
		log.Fatal(err.Error())
	}
}
