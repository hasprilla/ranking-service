package main

import (
	"github.com/hasprilla/ranking-service/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hasprilla/ranking-service/config"
	"github.com/hasprilla/ranking-service/controllers"
)

func main() {
	// Initialize Database (Optional for this service but good for consistency)
	config.ConnectDB()

	app := fiber.New(fiber.Config{
		AppName: "Sonifoy Ranking Service (Go)",
	})

	app.Use(logger.New())
	app.Use(recover.New())

	api := app.Group("/api/v1/ranking")
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "UP", "service": "ranking-service"})
	})
	api.Use(middleware.Protected()))

	api.Get("/artists", controllers.GetArtistRanking)
	api.Get("/fans", controllers.GetFanRanking)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "UP"})
	})

	log.Fatal(app.Listen(":8080"))
}
