package handler

import (
	"net/http"

	"github.com/MollenAR/internOzonFintech/internal/shortUrl/model"
	"github.com/MollenAR/internOzonFintech/internal/tools/errorTypes"
	"github.com/MollenAR/internOzonFintech/internal/tools/validation"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
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
	originalUrl := new(model.OriginalUrl)
	err := c.Bind(originalUrl)
	if err != nil {
		return errorTypes.ErrWrongUsage{
			Reason: err.Error(),
		}
	}

	err = validation.ValidateOriginalUrl(originalUrl.Url)
	if err != nil {
		return errorTypes.ErrWrongOriginalUrl{
			Reason: err.Error(),
		}
	}

	ctx := c.Request().Context()

	response, err := shurlHandler.ShortUrlUsecase.SaveOriginalUrl(ctx, originalUrl.Url)
	if err != nil {
		return errors.Wrap(err, "")
	}

	return c.JSON(http.StatusOK, response)
}

func (shurlHandler *ShortUrlHandler) GetOriginalUrl(c echo.Context) error {
	shortlUrl := c.Param("shortUrl")

	err := validation.ValidateShortUrl(shortlUrl)
	if err != nil {
		return errorTypes.ErrWrongShortUrl{
			Reason: err.Error(),
		}
	}

	ctx := c.Request().Context()

	response, err := shurlHandler.ShortUrlUsecase.GetOriginalUrl(ctx, shortlUrl)
	if err != nil {
		return errors.Wrap(err, "")
	}

	return c.JSON(http.StatusOK, response)
}
