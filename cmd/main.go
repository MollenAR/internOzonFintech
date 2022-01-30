package main

import (
	"fmt"
	"github.com/MollenAR/internOzonFintech/cmd/server"
	"net/http"
)

func main() {
	fmt.Println("start")

	if err := server.Run("localhost:8080"); err != http.ErrServerClosed {
		panic(err.Error())
	}
}
