package http

import (
	healthcheckergo "healthchecker"
	"healthchecker/config"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Run(cfg config.Config) {
	if err := startFiber(cfg); err != nil {
		panic(err)
	}
}

func startFiber(cfg config.Config) error {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetOutput(os.Stdout)
	
	app := fiber.New()
	pinger := healthcheckergo.NewPinger(cfg.StopOnFail, cfg.Timeout, logger)
	app.Use(setLogger(logger),setPinger(pinger))
	setupRoutes(app)

	return app.Listen(cfg.Port)
}