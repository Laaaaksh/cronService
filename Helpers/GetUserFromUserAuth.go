package Helpers

import (
	"cronService/Database"
	"cronService/Models"
)

func GetUserFromUserAuth(username string, user *Models.User) error{
	var userAuth Models.UserAuthentication
	if err:= Database.DB.Where("user_name = ?", username).Find(&userAuth).Error; err != nil{
		return err
	}
	if err:= Database.DB.Where("id = ?", userAuth.UserID).Find(&user).Error; err != nil{
		return err
	}
	return nil
}
