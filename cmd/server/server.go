package server

import (
	"github.com/labstack/echo/v4"
)

func Run(addres string) error {
	e := echo.New()

	return e.Start(addres)
}
