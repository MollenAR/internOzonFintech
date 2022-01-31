package main

import (
	"log"
	"net/http"

	"github.com/MollenAR/internOzonFintech/cmd/server"
	"github.com/MollenAR/internOzonFintech/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/tarantool/go-tarantool"
)

func main() {
	serverAddres, bdType, err := configs.ServerConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	var bdConn interface{}

	switch {
	case bdType == configs.BdTarantool:
		user, pass, addr, err := configs.TarantoolConfig()
		if err != nil {
			log.Fatal(err.Error())
		}

		opts := tarantool.Opts{User: user, Pass: pass}
		bdConn, err = tarantool.Connect(addr, opts)
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = bdConn.(*tarantool.Connection).Ping()
		if err != nil {
			log.Fatal(err.Error())
		}

	default:
		psqlCredentials, err := configs.PostgresConfig()
		if err != nil {
			log.Fatal(err.Error())
		}

		bdConn, err = sqlx.Open("postgres", psqlCredentials)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = bdConn.(*sqlx.DB).Ping()
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if err := server.Run(serverAddres, bdConn); err != http.ErrServerClosed {
		log.Fatal(err.Error())
	}
}
