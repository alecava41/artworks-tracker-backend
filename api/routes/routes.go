package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewRoomRoutes),
	fx.Provide(NewArtworkRoutes),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	//userRoutes UserRoutes,
	roomRoutes RoomRoutes,
	//authRoutes AuthRoutes,
	artworkRouter ArtworkRoutes,
) Routes {
	return Routes{
		//userRoutes,
		roomRoutes,
		//authRoutes,
		artworkRouter,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
