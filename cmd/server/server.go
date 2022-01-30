package server

import (
	"github.com/MollenAR/internOzonFintech/internal/middleware/addId"
	"github.com/MollenAR/internOzonFintech/internal/middleware/errorHandler"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/handler"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/repository/postgreSQL"
	// "github.com/MollenAR/internOzonFintech/internal/shortUrl/repository/postgreSQL"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/tarantool/go-tarantool"
)

func Run(addres string, psqlDb *sqlx.DB, tarantoolConn *tarantool.Connection) error {
	e := echo.New()
	e.Use(middleware.Recover())

	psqlRepo := postgreSQL.NewPsqlRepo(psqlDb)
	// tarantoolRepo := tRepo.NewTarantoolRepo(tarantoolConn)

	shortUrlUsecase := usecase.NewShortUrlUsecase(psqlRepo)
	// shortUrlUsecase := usecase.NewShortUrlUsecase(tarantoolRepo)

	shortUrlHandler := handler.NewShortUrlHandler(shortUrlUsecase)

	e.Use(addId.AddId)
	e.HTTPErrorHandler = errorHandler.ErrorHandler

	e.POST("/save", shortUrlHandler.SaveOriginalUrl)
	e.GET("/get/:shortUrl", shortUrlHandler.GetOriginalUrl)

	return e.Start(addres)
}
