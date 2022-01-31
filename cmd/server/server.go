package server

import (
	"github.com/MollenAR/internOzonFintech/internal/middleware/addId"
	"github.com/MollenAR/internOzonFintech/internal/middleware/errorHandler"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/handler"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/model"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/repository/postgreSQL"
	tRepo "github.com/MollenAR/internOzonFintech/internal/shortUrl/repository/tarantool"
	"github.com/pkg/errors"

	// "github.com/MollenAR/internOzonFintech/internal/shortUrl/repository/postgreSQL"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/tarantool/go-tarantool"
)

func Run(addres string, dbConn interface{}) error {
	e := echo.New()
	e.Use(middleware.Recover())

	var chosenRepo model.ShortUrlRepository

	switch dbConn.(type) {
	case *tarantool.Connection:
		chosenRepo = tRepo.NewTarantoolRepo(dbConn.(*tarantool.Connection))

	case *sqlx.DB:
		chosenRepo = postgreSQL.NewPsqlRepo(dbConn.(*sqlx.DB))
	default:
		return errors.New("wrong bd type")
	}
	// chosenRepo := postgreSQL.NewPsqlRepo(psqlDb)
	// chosenRepo := tRepo.NewTarantoolRepo(tarantoolConn)

	shortUrlUsecase := usecase.NewShortUrlUsecase(chosenRepo)

	shortUrlHandler := handler.NewShortUrlHandler(shortUrlUsecase)

	e.Use(addId.AddId)
	e.HTTPErrorHandler = errorHandler.ErrorHandler

	e.POST("/save", shortUrlHandler.SaveOriginalUrl)
	e.GET("/get/:shortUrl", shortUrlHandler.GetOriginalUrl)

	return e.Start(addres)
}
