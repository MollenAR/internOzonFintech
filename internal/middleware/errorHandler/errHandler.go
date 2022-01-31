package errorHandler

import (
	"net/http"

	"github.com/MollenAR/internOzonFintech/internal/tools/errorTypes"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type errResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func ErrorHandler(err error, c echo.Context) {
	var response errResponse

	switch {
	case errors.As(err, &errorTypes.ErrTryAgainLater{}):
		response.Msg = "Возникла ошибка при обработке вашего запроса, попробуйте еще раз через несколько минут"
		response.Status = http.StatusInternalServerError

	case errors.As(err, &errorTypes.ErrWrongUsage{}):
		response.Msg = "Неверное использование сервиса, проверьте тело запроса"
		response.Status = http.StatusUnprocessableEntity

	case errors.As(err, &errorTypes.ErrWrongOriginalUrl{}):
		response.Msg = "Оригинальный URL указан в неверном формате"
		response.Status = http.StatusUnprocessableEntity

	case errors.As(err, &errorTypes.ErrWrongShortUrl{}):
		response.Msg = "Такого короткого URL не существует"
		response.Status = http.StatusNotFound

	default:
		response.Msg = "Возникла ошибка при обработке вашего запроса, попробуйте еще раз через несколько минут"
		response.Status = http.StatusInternalServerError
	}

	// todo logger

	c.Logger().Error(err.Error())
	errJson := c.JSON(response.Status, response)
	if errJson != nil {
		return
	}
}
