package routes

import (
	"github.com/eren-ay/golang-restapi/app/controllers/post"
	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App) {
	// Create Post
	app.Post("/posts", post.Create)
	// List All Posts
	app.Get("/posts", post.Index)
}
