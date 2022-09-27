package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Status() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "api is working")
	}
}
