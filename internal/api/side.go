package api

import (
	"net/http"

	"valo-tips/internal/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetSides() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorm.DB)

		sides := new(model.Sides)
		if err := sides.GetAll(tx); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Does not exists.")
		}
		return c.JSON(http.StatusOK, sides)
	}
}
