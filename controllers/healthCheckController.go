package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func LivenessCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func ReadinessCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}
