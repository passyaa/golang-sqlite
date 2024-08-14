package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// BasicAuthMiddleware is a middleware to authenticate using Basic Auth
func BasicAuthMiddleware(username, password string, c echo.Context) (bool, error) {
	// Replace with your user validation logic
	if username == "spadmin" && password == "admin" {
		return true, nil
	} else {
		return false, echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}
	return false, nil
}
