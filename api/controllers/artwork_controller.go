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
	artID := c.Param("beaconId")

	id, err := uuid.Parse(artID)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	lanID := c.DefaultQuery("lan", "en")

	artwork, err := u.service.GetOneArtwork(id, lanID)

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

	c.JSON(200, artwork)
}
