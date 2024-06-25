package controllers

import (
	"slot-machine-api/services"

	"github.com/gofiber/fiber/v2"
)

func PlaySlotMachine(c *fiber.Ctx) error {
	var request struct {
		PlayerID string `json:"player_id"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	result, err := services.Play(request.PlayerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(result)
}

func GetPlayerGames(c *fiber.Ctx) error {
	playerID := c.Params("id")

	games, err := services.GetGamesByPlayerID(playerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(games)
}
