// handlers/group_handler.go
package handlers

import (
	"golangApp/config"
	"golangApp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetGroup fetches a single group by ID
func GetGroup(c echo.Context) error {
	id := c.Param("id")
	var group models.Group

	// Mencari group berdasarkan ID
	if err := config.DB.First(&group, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Group not found",
		})
	}

	return c.JSON(http.StatusOK, group)
}

// GetAllGroups fetches all groups
func GetAllGroups(c echo.Context) error {
	var groups []models.Group

	// Mengambil semua grup
	if err := config.DB.Find(&groups).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve groups",
		})
	}

	return c.JSON(http.StatusOK, groups)
}

func CreateGroup(c echo.Context) error {
	var group models.Group

	// Bind input JSON to the group struct
	if err := c.Bind(&group); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
		})
	}

	// Validate the required fields
	if group.Name == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Group name is required",
		})
	}

	// Create the new group in the database
	if err := config.DB.Create(&group).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create group",
		})
	}

	return c.JSON(http.StatusCreated, group)
}

// AssignGroup assigns a group to a user
func AssignGroup(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid user ID",
		})
	}

	groupID, err := strconv.Atoi(c.Param("group_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid group ID",
		})
	}

	var user models.User
	var group models.Group

	// Mencari user berdasarkan ID
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	// Mencari group berdasarkan ID
	if err := config.DB.First(&group, groupID).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Group not found",
		})
	}

	// Menetapkan group ke user
	if err := config.DB.Model(&user).Association("Groups").Append(&group); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to assign group",
		})
	}

	return c.NoContent(http.StatusOK)
}

// RemoveGroup removes a group
func RemoveGroup(c echo.Context) error {
	id := c.Param("group_id")
	var group models.Group

	// Mencari group berdasarkan ID
	if err := config.DB.First(&group, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Group not found",
		})
	}

	// Menghapus grup dari database
	if err := config.DB.Delete(&group).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete group",
		})
	}

	return c.NoContent(http.StatusNoContent)
}
