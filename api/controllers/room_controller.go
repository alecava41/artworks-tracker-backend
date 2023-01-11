package controllers

import (
	"errors"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

// RoomController data type
type RoomController struct {
	service services.RoomService
	logger  lib.Logger
}

// NewRoomController creates new user controller
func NewRoomController(roomService services.RoomService, logger lib.Logger) RoomController {
	return RoomController{
		service: roomService,
		logger:  logger,
	}
}

// GetOneRoom gets one user
func (u RoomController) GetOneRoom(c *gin.Context) {
	artID := c.Query("beacon")

	id, err := uuid.Parse(artID)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	lanID := c.DefaultQuery("lan", "en")

	room, err := u.service.GetOneRoom(id, lanID)

	if err != nil {
		u.logger.Error(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		return
	}

	c.JSON(200, room)
}
