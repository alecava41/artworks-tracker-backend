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
		api.GET("/artworks/:artId", s.artworkController.GetOneArtwork)
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
