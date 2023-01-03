package routes

import (
	"github.com/dipeshdulal/clean-gin/api/controllers"
	"github.com/dipeshdulal/clean-gin/api/middlewares"
	"github.com/dipeshdulal/clean-gin/lib"
)

// ArtworkRoutes struct
type ArtworkRoutes struct {
	logger            lib.Logger
	handler           lib.RequestHandler
	artworkController controllers.ArtworkController
	authMiddleware    middlewares.JWTAuthMiddleware
}

// Setup user routes
func (s ArtworkRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/api")
	{
		//api.GET("/artworks/:beaconId/media/:img", s.artworkController.GetArtworkImage)
		api.GET("/artworks/:beaconId", s.artworkController.GetOneArtwork)
		api.GET("/artworks/:beaconId/media/:img", s.artworkController.GetArtworkImage)
	}
}

// NewArtworkRoutes creates new user controller
func NewArtworkRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	artworkController controllers.ArtworkController,
	authMiddleware middlewares.JWTAuthMiddleware,
) ArtworkRoutes {
	return ArtworkRoutes{
		handler:           handler,
		logger:            logger,
		artworkController: artworkController,
		authMiddleware:    authMiddleware,
	}
}
