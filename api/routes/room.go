package routes

import (
	"github.com/dipeshdulal/clean-gin/api/controllers"
	"github.com/dipeshdulal/clean-gin/api/middlewares"
	"github.com/dipeshdulal/clean-gin/lib"
)

// RoomRoutes struct
type RoomRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	roomController controllers.RoomController
	authMiddleware middlewares.JWTAuthMiddleware
}

// Setup user routes
func (s RoomRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/api")
	{
		api.GET("/rooms/", s.roomController.GetOneRoom)
	}
}

// NewRoomRoutes creates new user controller
func NewRoomRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	roomController controllers.RoomController,
	authMiddleware middlewares.JWTAuthMiddleware,
) RoomRoutes {
	return RoomRoutes{
		handler:        handler,
		logger:         logger,
		roomController: roomController,
		authMiddleware: authMiddleware,
	}
}
