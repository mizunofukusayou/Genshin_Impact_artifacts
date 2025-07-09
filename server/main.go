package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mizunofukusayou/Genshin_Impact_artifacts/db"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger(), middleware.Recover())

	e.GET("/best-artifacts", db.SearchBestArtifacts)
}
