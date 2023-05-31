package models

import (
	"gorm.io/gorm"
)

type Post struct {
	ID      uint    `gorm:"primary key;autoIncrement" json:"id"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) error {

	return nil
}
