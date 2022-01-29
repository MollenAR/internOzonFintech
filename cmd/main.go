package main

import (
	"fmt"
	"github.com/MollenAR/internOzonFintech/cmd/server"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	fmt.Println("start")

	if err := server.Run(":8080"); err != http.ErrServerClosed {
		echo.Logger.Fatal(err.Error())
	}
}
