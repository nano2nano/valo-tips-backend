package main

import (
	"os"

	"valo-tips/internal/router"
)

func main() {
	router := router.Init()
	router.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))
}
