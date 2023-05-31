package utils

import (
	"fmt"
	"log"

	"github.com/eren-ay/golang-restapi/pkg/routes"
	"github.com/eren-ay/golang-restapi/platform/database"
	"github.com/eren-ay/golang-restapi/platform/migrations"
	"github.com/gofiber/fiber/v2"
)

func CreateServer(port int) {
	// Create Fiber App
	app := fiber.New()
	routes.PostRoutes(app)
	err := database.Init()
	if err != nil {
		panic(err)
	}
	// database gorm AutoMigrate
	migrations.Migrate()

	// Start server
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
