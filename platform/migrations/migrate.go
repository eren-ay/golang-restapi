package migrations

import (
	"github.com/eren-ay/golang-restapi/app/models"
	"github.com/eren-ay/golang-restapi/platform/database"
)

func Migrate() {
	database.Conn.AutoMigrate(&models.Post{})
}
