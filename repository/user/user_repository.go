package user

import (
	"TRANSFERSYSTEM/app/database"
	"TRANSFERSYSTEM/model"
	"fmt"
)

type User struct {
	model.User
}

func GetAll() []model.User {
	var users []model.User
	database.DB.Find(&users)

	return users
}

func FindById(id string) model.User {
	var user model.User
	database.DB.Where("ID = ?", id).Find(&user)

	return user
}

func FindBy(whereConditions map[string]string) model.User {
	var user model.User

	for key, val := range whereConditions {
		database.DB = database.DB.Where(fmt.Sprintf("%s = ?", key), val)
	}

	database.DB.Find(&user)

	return user
}

func UpdateBalance(id string, balance float64) model.User {
	user := FindById(id)
	database.DB.Model(&user).Where("ID = ?", id).Update("Balance", balance)

	return user
}

func (s *User) Create() {
	database.DB.Create(&s)
}