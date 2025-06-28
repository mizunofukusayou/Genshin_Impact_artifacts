package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4@v4.13.4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger(), middleware.Recover())

	db, err := coredb.New()
	if err != nil {
		e.Logger.Fatal("Database connection failed", err)
	}

	e.POST("/artifacts", func(c echo.Context) error {
		var artifact coredb.Artifact
		if err := c.Bind(&artifact); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}
		if err := db.AddArtifact(artifact); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add artifact"})
		}
		return c.NoContent(http.StatusCreated)
	})
	e.POST("/weapons", func(c echo.Context) error {
		var weapon coredb.Weapon
		if err := c.Bind(&weapon); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}
		if err := db.AddWeapon(weapon); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add weapon"})
		}
		return c.NoContent(http.StatusCreated)
	})
	e.POST("/characters", func(c echo.Context) error {
		var character coredb.Character
		if err := c.Bind(&character); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}
		if err := db.AddCharacter(character); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add character"})
		}
		return c.NoContent(http.StatusCreated)
	})
}
