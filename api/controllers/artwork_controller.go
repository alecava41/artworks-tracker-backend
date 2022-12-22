package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/services"
	"github.com/gin-gonic/gin"
)

// ArtworkController data type
type ArtworkController struct {
	service services.ArtworkService
	logger  lib.Logger
}

// NewArtworkController creates new user controller
func NewArtworkController(roomService services.ArtworkService, logger lib.Logger) ArtworkController {
	return ArtworkController{
		service: roomService,
		logger:  logger,
	}
}

// GetOneArtwork gets one user
func (u ArtworkController) GetOneArtwork(c *gin.Context) {
	artID := c.Param("artId")

	id, err := strconv.Atoi(artID)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	lanID := c.Query("lan")
	if lanID == "" {
		u.logger.Error(errors.New("language query parameter missing"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	room, err := u.service.GetOneArtwork(uint(id), lanID)

	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, room)
}
