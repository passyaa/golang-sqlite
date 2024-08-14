package middlewares

import (
	"github.com/labstack/echo/v4"
)

// BasicAuthMiddleware is a middleware to authenticate using Basic Auth
func BasicAuthMiddleware(username, password string, c echo.Context) (bool, error) {
	// Replace with your user validation logic
	if username == "spadmin" && password == "admin" {
		return true, nil
	}
	return false, nil
}
