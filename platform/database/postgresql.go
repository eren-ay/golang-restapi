package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "todos"
)

var Conn *gorm.DB

func Init() error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Conn = conn
	if err != nil {
		return err
	}

	return nil
}
