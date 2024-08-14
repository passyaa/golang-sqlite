package handlers

import (
	"golangApp/config"
	"golangApp/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo-contrib/session"
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

	// Simpan waktu login terakhir
	user.LastLogin = time.Now()
	config.DB.Save(&user)

	// Simpan status login di session
	sess, _ := session.Get("session", c)
	sess.Values["username"] = user.Username
	sess.Values["userID"] = user.ID
	sess.Save(c.Request(), c.Response())

	// Redirect ke halaman profil pengguna
	return c.Redirect(http.StatusFound, "/profile/"+strconv.Itoa(user.ID))
}

// UserProfile renders the profile page for the logged-in user
func UserProfile(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid user ID",
		})
	}

	var user models.User

	// Mencari user berdasarkan ID
	if err := config.DB.Preload("Groups").First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	// Render the HTML page
	return c.Render(http.StatusOK, "profile.html", user)
}

// Logout logs the user out and redirects to the home page
func Logout(c echo.Context) error {
	// Hapus sesi atau informasi login
	sess, _ := session.Get("session", c)
	sess.Values["username"] = nil
	sess.Values["userID"] = nil
	sess.Save(c.Request(), c.Response())

	// Redirect ke halaman utama
	return c.Redirect(http.StatusFound, "/")
}
