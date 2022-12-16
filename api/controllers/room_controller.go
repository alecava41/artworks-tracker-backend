package controllers

import (
	"net/http"
	"strconv"

	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/services"
	"github.com/gin-gonic/gin"
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
	artID := c.Param("artId")

	id, err := strconv.Atoi(artID)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	room, err := u.service.GetOneRoom(uint(id))

	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": room,
	})
}
