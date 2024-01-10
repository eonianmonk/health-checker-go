package http

import (
	healthcheckergo "healthchecker"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)


func setLogger(logger *logrus.Logger)  func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Locals("logger", logger)
		return c.Next()
	}
}

func setPinger(pinger *healthcheckergo.Pinger)  func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Locals("pinger", pinger)
		return c.Next()
	}
}

func GetLogger(c *fiber.Ctx) *logrus.Logger {
	return c.Locals("logger").(*logrus.Logger)
}
func GetPinger(c *fiber.Ctx) *healthcheckergo.Pinger {
	return c.Locals("pinger").(*healthcheckergo.Pinger)
}
