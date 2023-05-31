package post

import (
	"github.com/eren-ay/golang-restapi/app/models"
	"github.com/eren-ay/golang-restapi/platform/database"
	"github.com/gofiber/fiber/v2"
)

func Create(ctx *fiber.Ctx) error {
	post := models.Post{}

	if err := ctx.BodyParser(&post); err != nil {
		return ctx.Status(503).JSON(err)
	}

	if err := database.Conn.Create(&post).Error; err != nil {
		return ctx.Status(503).JSON(err)
	}

	return ctx.JSON(post)
}

func Index(ctx *fiber.Ctx) error {
	posts := []models.Post{}
	database.Conn.Find(&posts)

	return ctx.JSON(posts)
}
