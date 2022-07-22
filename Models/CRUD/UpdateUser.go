package CRUD

import (
	"cronService/Database"
	"cronService/Models"
)

func UpdateUser(user Models.User)(err error){
	var user_new Models.User

	err=Database.DB.Find(&user_new,"id=?",user.Id).Error;if err!=nil{
		return err
	}
	Database.DB.Model(user_new).Update(
		"PermissionID",
		user.PermissionID,
	)
	return nil
}
