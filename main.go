package main

import (
	"slot-machine-api/controllers"
	"slot-machine-api/jobs"
	"slot-machine-api/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize MongoDB and Redis
	utils.InitMongo()
	utils.InitRedis()

	// Run Initialization Job
	jobs.InitIndexes()

	// Player routes
	app.Post("/players", controllers.CreatePlayer)
	app.Get("/players/:id", controllers.GetPlayer)
	app.Put("/players/:id/suspend", controllers.SuspendPlayer)
	app.Get("/players/:id/games", controllers.GetPlayerGames)

	// Slot machine route
	app.Post("/play", controllers.PlaySlotMachine)

	// Health check routes
	app.Get("/health", controllers.HealthCheck)
	app.Get("/liveness", controllers.LivenessCheck)
	app.Get("/readiness", controllers.ReadinessCheck)

	app.Listen(":3000")
}
