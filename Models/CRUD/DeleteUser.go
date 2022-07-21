package CRUD

import (
	"cronService/Database"
	"cronService/Models"
)

func DeleteUserById(user *Models.User,id string)(err error){

	if err=Database.DB.Where("id = ?", id).Delete(user).Error; err!=nil {
		return err
	}
	return nil
}

