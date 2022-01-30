package handler

import (
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ShortUrlHandler struct {
	ShortUrlUsecase model.ShortUrlUsecase
}

func NewShortUrlHandler(shurlUsecase model.ShortUrlUsecase) *ShortUrlHandler {
	return &ShortUrlHandler{
		ShortUrlUsecase: shurlUsecase,
	}
}

func (shurlHandler *ShortUrlHandler) SaveOriginalUrl(c echo.Context) error {
	var originalUrl string
	err := c.Bind(originalUrl)
	if err != nil {
		// todo errors
	}

	// todo validation

	ctx := c.Request().Context()
	response, err := shurlHandler.ShortUrlUsecase.SaveOriginalUrl(ctx, originalUrl)
	if err != nil {
		// todo errors
	}

	return c.JSON(http.StatusOK, response)
}

func (shurlHandler *ShortUrlHandler) GetOriginalUrl(c echo.Context) error {
	var shortlUrl string
	err := c.Bind(shortlUrl)
	if err != nil {
		// todo errors
	}

	// todo validation

	ctx := c.Request().Context()
	response, err := shurlHandler.ShortUrlUsecase.GetOriginalUrl(ctx, shortlUrl)
	if err != nil {
		// todo errors
	}

	return c.JSON(http.StatusOK, response)
}
