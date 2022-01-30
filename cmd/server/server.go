package server

import (
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/handler"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/repository/postgreSQL"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/usecase"
	"github.com/labstack/echo/v4"
)

func Run(addres string) error {
	e := echo.New()

	psqlRepo := postgreSQL.NewPsqlRepo()

	shortUrlUsecase := usecase.NewShortUrlUsecase(psqlRepo)

	shortUrlHandler := handler.NewShortUrlHandler(shortUrlUsecase)

	e.POST("/save", shortUrlHandler.SaveOriginalUrl)
	e.GET("/get/:shortUrl", shortUrlHandler.GetOriginalUrl)

	return e.Start(addres)
}
