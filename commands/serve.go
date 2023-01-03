package commands

import (
	"context"
	"github.com/brutella/dnssd"
	"github.com/dipeshdulal/clean-gin/api/middlewares"
	"github.com/dipeshdulal/clean-gin/api/routes"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/spf13/cobra"
	"log"
)

// ServeCommand test command
type ServeCommand struct{}

func (s *ServeCommand) Short() string {
	return "serve application"
}

func (s *ServeCommand) Setup(*cobra.Command) {}

func (s *ServeCommand) Run() lib.CommandRunner {
	return func(
		middleware middlewares.Middlewares,
		env lib.Env,
		router lib.RequestHandler,
		route routes.Routes,
		logger lib.Logger,
		database lib.Database,
	) {
		middleware.Setup()
		route.Setup()

		cfg := dnssd.Config{
			Name: "MuseumBackend",
			Type: "_http._tcp",
			Port: 5000,
		}

		sv, err := dnssd.NewService(cfg)
		if err != nil {
			log.Println(err.Error())
		}

		rp, err := dnssd.NewResponder()
		if err != nil {
			log.Println(err.Error())
		}

		_, err = rp.Add(sv)
		if err != nil {
			log.Println(err.Error())
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go func() {
			err = rp.Respond(ctx)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}()

		logger.Info("Running server")
		if env.ServerPort == "" {
			_ = router.Gin.Run()
		} else {
			_ = router.Gin.Run(":" + env.ServerPort)
		}
	}
}

func NewServeCommand() *ServeCommand {
	return &ServeCommand{}
}
