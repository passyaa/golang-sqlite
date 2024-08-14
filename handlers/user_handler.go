// handlers/user_handler.go
package handlers

import (
	"golangApp/config"
	"golangApp/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// GetUser fetches a single user by ID
func GetUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	// Mengambil user beserta grup yang diassign menggunakan Preload
	if err := config.DB.Preload("Groups").First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	return c.JSON(http.StatusOK, user)
}

// GetAllUsers fetches all users
func GetAllUsers(c echo.Context) error {
	var users []models.User

	// Mengambil semua user beserta grup yang diassign menggunakan Preload
	if err := config.DB.Preload("Groups").Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve users",
		})
	}

	return c.JSON(http.StatusOK, users)
}

// CreateUser creates a new user
func CreateUser(c echo.Context) error {
	var user models.User

	// Bind input ke struct user
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
		})
	}

	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to hash password",
		})
	}
	user.Password = string(hashedPassword)

	// Menyimpan user baru ke database dengan GORM
	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create user",
		})
	}

	return c.JSON(http.StatusCreated, user)
}

// UpdateUser updates user attributes
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	// Mencari user berdasarkan ID dengan GORM
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	// Bind input ke struct user (hanya field yang ingin diupdate)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
		})
	}

	// Mengupdate user di database dengan GORM
	if err := config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update user",
		})
	}

	return c.JSON(http.StatusOK, user)
}

// EnableUser enables a user
func EnableUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	// Mencari user berdasarkan ID
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	// Mengubah is_enabled menjadi true
	user.IsEnabled = true
	if err := config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to enable user",
		})
	}

	return c.JSON(http.StatusOK, user)
}

// DisableUser disables a user
func DisableUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	// Mencari user berdasarkan ID
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	// Mengubah is_enabled menjadi false
	user.IsEnabled = false
	if err := config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to disable user",
		})
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user resource
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	// Mencari user berdasarkan ID
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	// Menghapus user dari database
	if err := config.DB.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete user",
		})
	}

	return c.NoContent(http.StatusNoContent)
}

// ResetPassword resets user password
func ResetPassword(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	// Mencari user berdasarkan ID
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	// Mendapatkan password baru dari request body
	type ResetPasswordRequest struct {
		NewPassword string `json:"new_password"`
	}

	var req ResetPasswordRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
		})
	}

	// Hash password baru
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to hash password",
		})
	}

	// Mengupdate password user di database
	user.Password = string(hashedPassword)
	if err := config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to reset password",
		})
	}

	return c.NoContent(http.StatusOK)
}
