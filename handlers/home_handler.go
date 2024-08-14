package handlers

import (
	"golangApp/config"
	"golangApp/models"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// HomePage renders the home page with a list of all users
func HomePage(c echo.Context) error {
	var users []models.User

	// Mengambil semua pengguna dari database
	if err := config.DB.Preload("Groups").Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve users",
		})
	}

	// Render the HTML page
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"users": users,
	})
}
