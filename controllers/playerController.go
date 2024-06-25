package controllers

import (
	"slot-machine-api/models"
	"slot-machine-api/services"

	"github.com/gofiber/fiber/v2"
)

func CreatePlayer(c *fiber.Ctx) error {
	var player models.Player
	if err := c.BodyParser(&player); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	result, err := services.CreatePlayer(player)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func GetPlayer(c *fiber.Ctx) error {
	id := c.Params("id")

	player, err := services.GetPlayerByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return c.JSON(player)
}

func SuspendPlayer(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := services.SuspendPlayer(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(result)
}
