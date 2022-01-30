package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MollenAR/internOzonFintech/cmd/server"
	"github.com/MollenAR/internOzonFintech/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/tarantool/go-tarantool"
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

	user, pass, addr, err := configs.TarantoolConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	opts := tarantool.Opts{User: user, Pass: pass}
	tarantoolConn, err := tarantool.Connect(addr, opts)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = tarantoolConn.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}


	if err := server.Run(serverAddres, psqldb, tarantoolConn); err != http.ErrServerClosed {
		log.Fatal(err.Error())
	}
}
