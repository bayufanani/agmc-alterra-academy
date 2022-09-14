package database

import (
	"agmc/day2/config"
	"agmc/day2/models"
)

func GetUser() (interface{}, error) {
	var users []models.User
	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetOneUser(id int) (interface{}, error) {
	var users []models.User
	if e := config.DB.Where("id=?", id).Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func CreateUser(user *models.User) (interface{}, error) {
	if e := config.DB.Create(&user).Error; e != nil {
		return nil, e
	}
	return &user, nil
}

func UpdateUser(id int, user *models.User) (interface{}, error) {
	if e := config.DB.Model(&models.User{}).Where("id=?", id).Updates(&user).Error; e != nil {
		return nil, e
	}
	return &user, nil
}

func DeleteUser(id int) error {
	if e := config.DB.Delete(models.User{}, id).Error; e != nil {
		return e
	}
	return nil
}
