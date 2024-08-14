package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// BasicAuthMiddleware provides a middleware for Basic Authentication
func BasicAuthMiddleware() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Replace with your user validation logic
		if username == "admin" && password == "admin" {
			return true, nil
		}
		return false, nil
	})
}
