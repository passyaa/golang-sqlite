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

	var userGroup models.UserGroup
	userGroup.UserID = userID
	userGroup.GroupID = groupID

	// Menetapkan grup ke pengguna
	if err := config.DB.Create(&userGroup).Error; err != nil {
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
