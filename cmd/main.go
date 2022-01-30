package main

import (
	"fmt"
	"github.com/MollenAR/internOzonFintech/cmd/server"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start")

	if err := server.Run("localhost:8080"); err != http.ErrServerClosed {
		log.Fatal(err.Error())
	}
}
