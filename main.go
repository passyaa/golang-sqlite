package main

import (
	"golangApp/config"
	"golangApp/handlers"
	"golangApp/middlewares"
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	config.InitDB()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Apply the Basic Auth Middleware
	e.Use(middlewares.BasicAuthMiddleware())

	// Set up template renderer
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	// Routes
	e.GET("/users/:id", handlers.GetUser)
	e.GET("/users", handlers.GetAllUsers)
	e.GET("/groups/:id", handlers.GetGroup)
	e.GET("/groups", handlers.GetAllGroups)
	e.POST("/users", handlers.CreateUser)
	e.POST("/groups", handlers.CreateGroup)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.PUT("/users/:id/enable", handlers.EnableUser)
	e.PUT("/users/:id/disable", handlers.DisableUser)
	e.DELETE("/users/:id", handlers.DeleteUser)
	e.POST("/users/:id/groups/:group_id", handlers.AssignGroup)
	e.DELETE("/groups/:group_id", handlers.RemoveGroup)
	e.PUT("/users/:id/reset_password", handlers.ResetPassword)
	e.GET("/login", handlers.ShowLoginPage)       // Menampilkan halaman login
	e.POST("/login", handlers.HandleLogin)        // Memproses login
	e.GET("/", handlers.HomePage)                 // Halaman home setelah login berhasil

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
