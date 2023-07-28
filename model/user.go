package model

import (
	"flashcards-api/app/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID string `gorm:"" json:"id"`
	Name string `gorm:"" json:"name"`
	Balance  string `gorm:"" json:"balance"`
	Password string `gorm:"" json:"password"`
}

func init() {
	database.DB.AutoMigrate(&User{})
}
