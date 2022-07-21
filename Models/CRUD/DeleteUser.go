package CRUD

import (
	"cronService/Database"
	"cronService/Models"
)

func DeleteUserById(user *Models.User,id string)(err error){

	Database.DB.Where("id = ?", id).Delete(user)
	return nil
}

