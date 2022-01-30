package server

import (
	"github.com/MollenAR/internOzonFintech/internal/middleware/addId"
	"github.com/MollenAR/internOzonFintech/internal/middleware/errorHandler"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/handler"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/repository/postgreSQL"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(addres string) error {
	e := echo.New()
	e.Use(middleware.Recover())

	psqlRepo := postgreSQL.NewPsqlRepo()

	shortUrlUsecase := usecase.NewShortUrlUsecase(psqlRepo)

	shortUrlHandler := handler.NewShortUrlHandler(shortUrlUsecase)

	e.Use(addId.AddId)
	e.HTTPErrorHandler = errorHandler.ErrorHandler

	e.POST("/save", shortUrlHandler.SaveOriginalUrl)
	e.GET("/get/:shortUrl", shortUrlHandler.GetOriginalUrl)

	return e.Start(addres)
}
