package handlers

import (
	"golangApp/config"
	"golangApp/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// ShowLoginPage renders the login page
func ShowLoginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

// HandleLogin handles the login form submission
func HandleLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Mencari user berdasarkan username
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return c.Render(http.StatusOK, "login.html", map[string]interface{}{
			"error": "Invalid username or password",
		})
	}

	// Verifikasi password dengan bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return c.Render(http.StatusOK, "login.html", map[string]interface{}{
			"error": "Invalid username or password",
		})
	}

	// Jika username dan password valid, redirect ke halaman home
	return c.Redirect(http.StatusFound, "/")
}
