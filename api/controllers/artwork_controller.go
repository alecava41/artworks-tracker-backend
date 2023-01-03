package controllers

import (
	"errors"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"os"
	"path/filepath"
)

// ArtworkController data type
type ArtworkController struct {
	service services.ArtworkService
	logger  lib.Logger
}

// NewArtworkController creates new artwork controller
func NewArtworkController(artworkService services.ArtworkService, logger lib.Logger) ArtworkController {
	return ArtworkController{
		service: artworkService,
		logger:  logger,
	}
}

// GetOneArtwork gets one artwork
func (a ArtworkController) GetOneArtwork(c *gin.Context) {
	artID := c.Param("beaconId")

	id, err := uuid.Parse(artID)
	if err != nil {
		a.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	lanID := c.DefaultQuery("lan", "en")

	artwork, err := a.service.GetOneArtwork(id, lanID)

	if err != nil {
		a.logger.Error(err)

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

// GetArtworkImage gets one artwork
func (a ArtworkController) GetArtworkImage(c *gin.Context) {
	beaconId := c.Param("beaconId")
	imgNumber := c.Param("img")

	currentPath, err := os.Getwd()
	if err != nil {
		a.logger.Error(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	fullName := filepath.Join(currentPath, "..", "images", "artworks", beaconId, imgNumber)

	files, err := filepath.Glob(fullName + "*")
	if err != nil {
		a.logger.Error(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.File(files[0])
}
