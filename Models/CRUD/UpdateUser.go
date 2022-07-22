package CRUD

import (
	"cronService/Database"
	"cronService/Models"
	"time"
)

func UpdateUser(user Models.User)(err error){
	var user_new Models.User

	err=Database.DB.Find(&user_new,"id=?",user.Id).Error;if err!=nil{
		return err
	}
	user_new.PermissionID = user.PermissionID
	user_new.UpdatedAT = time.Now().Unix()
	user.UpdatedAT = time.Now().Unix()
	Database.DB.Save(user_new)
	//Database.DB.Model(user_new).Update(Models.User{PermissionID: user.PermissionID, UpdatedAT: time.Now().Unix()})
	return nil
}
