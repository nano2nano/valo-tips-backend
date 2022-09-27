package router

import (
	"valo-tips/internal/api"
	"valo-tips/internal/db"
	myMw "valo-tips/internal/middleware"

	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(echoMw.Logger())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAccessControlAllowOrigin},
	}))

	// Set Custom MiddleWare
	e.Use(myMw.TransactionHandler(db.Init()))

	// Routes
	e.GET("/", api.Status())
	v1 := e.Group("/api/v1")

	// Tip
	tip := v1.Group("/tip")
	tip.POST("", api.PostTip())
	tip.GET("", api.GetTips())
	tip.GET("/:id", api.GetTip())

	// Side
	side := v1.Group("/side")
	side.GET("", api.GetSides())

	// Image
	img := v1.Group("/img")
	img.POST("/upload", api.PostImg())
	img.GET("/:name", api.GetImg())
	return e
}
