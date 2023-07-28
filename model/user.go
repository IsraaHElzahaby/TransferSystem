package model

import (
	"TRANSFERSYSTEM/app/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID string `gorm:"" json:"id"`
	Name string `gorm:"" json:"name"`
	Balance  float64 `gorm:"" json:"balance"`
	Password string `gorm:"" json:"password"`
}

func init() {
	database.DB.AutoMigrate(&User{})
}
