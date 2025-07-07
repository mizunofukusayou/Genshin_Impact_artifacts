package main

import (
	"./db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4@v4.13.4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger(), middleware.Recover())

	e.GET("/best-artifacts", db.SearchBestArtifacts())
}
