package middlewares

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
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
}
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if sess.Values["username"] == nil {
			return c.Redirect(http.StatusFound, "/login")
		}
		return next(c)
	}
}
