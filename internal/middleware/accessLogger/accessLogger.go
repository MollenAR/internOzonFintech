package accessLogger

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func WriteLog(r *http.Request, c echo.Context, start time.Time) {
	Id := r.Header.Get(echo.HeaderXRequestID)
	c.Logger().Info("Id = ", Id, " method = ", r.Method, " address = ", r.RemoteAddr, " RequestUri = ", r.RequestURI, " Request Time = ", time.Since(start))
}

func AccessLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		WriteLog(c.Request(), c, start)
		return err
	}
}
