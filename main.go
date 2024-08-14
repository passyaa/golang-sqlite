package main

import (
	"golangApp/config"
	"golangApp/handlers"
	"golangApp/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.InitDB()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Apply the Basic Auth Middleware
	e.Use(middlewares.BasicAuthMiddleware())

	// Routes
	e.GET("/users/:id", handlers.GetUser)
	e.GET("/users", handlers.GetAllUsers)
	e.GET("/groups/:id", handlers.GetGroup)
	e.GET("/groups", handlers.GetAllGroups)
	e.POST("/users", handlers.CreateUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.PUT("/users/:id/enable", handlers.EnableUser)
	e.PUT("/users/:id/disable", handlers.DisableUser)
	e.DELETE("/users/:id", handlers.DeleteUser)
	e.POST("/users/:id/groups/:group_id", handlers.AssignGroup)
	e.DELETE("/groups/:group_id", handlers.RemoveGroup)
	e.PUT("/users/:id/reset_password", handlers.ResetPassword)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
