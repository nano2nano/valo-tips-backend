package middleware

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	TxKey = "Tx"
)

func TransactionHandler(db *gorm.DB) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			tx := db.Begin()

			c.Set(TxKey, tx)

			if err := next(c); err != nil {
				tx.Rollback()
				return err
			}
			tx.Commit()

			return nil
		})
	}
}
