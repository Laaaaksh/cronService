package CRUD

import (
	"cronService/Database"
	"cronService/Models"
)

func CreateUser(user * Models.User) (err error) {
	if err := Database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

