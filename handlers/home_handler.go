package handlers

import (
	"golangApp/config"
	"golangApp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HomePage renders the home page with a list of all users and groups
func HomePage(c echo.Context) error {
	var users []models.User
	var groups []models.Group

	// Mengambil semua pengguna dari database
	if err := config.DB.Preload("Groups").Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve users",
		})
	}

	// Mengambil semua grup dari database
	if err := config.DB.Find(&groups).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve groups",
		})
	}

	// Render the HTML page
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"users":  users,
		"groups": groups,
	})
}
