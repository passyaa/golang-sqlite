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

	// Set up template renderer
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	// Routes that don't require authentication
	e.GET("/", handlers.HomePage) // Halaman home dapat diakses tanpa autentikasi
	e.GET("/login", handlers.ShowLoginPage)
	e.POST("/login", handlers.HandleLogin)

	// Route to show the user's profile
	e.GET("/profile/:id", handlers.UserProfile)

	// Group of routes that require authentication
	auth := e.Group("/api/v1")

	// Apply the Basic Auth Middleware only to specific routes
	auth.Use(middleware.BasicAuth(middlewares.BasicAuthMiddleware))

	auth.GET("/users/:id", handlers.GetUser)
	auth.GET("/users", handlers.GetAllUsers)
	auth.GET("/groups/:id", handlers.GetGroup)
	auth.GET("/groups", handlers.GetAllGroups)
	auth.POST("/users", handlers.CreateUser)
	auth.POST("/groups", handlers.CreateGroup)
	auth.PUT("/users/:id", handlers.UpdateUser)
	auth.PUT("/users/:id/enable", handlers.EnableUser)
	auth.PUT("/users/:id/disable", handlers.DisableUser)
	auth.DELETE("/users/:id", handlers.DeleteUser)
	auth.POST("/users/:id/groups/:group_id", handlers.AssignGroup)
	auth.DELETE("/users/:id/groups/:group_id", handlers.RemoveAssignGroup)
	auth.DELETE("/groups/:group_id", handlers.RemoveGroup)
	auth.PUT("/users/:id/reset_password", handlers.ResetPassword)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
