package Helpers

import (
	"cronService/Database"
	"cronService/Models"
)

func CheckPermission(userName string, permission string) bool {

	var user Models.User
	if err := GetUserFromUserAuth(userName, &user); err != nil{
		return false
	}
	if user.UserType == "Admin"{
		return true
	}
	var permissionType Models.PermissionType
	if err := Database.DB.Where("id = ?", user.PermissionID).Find(&permissionType).Error; err != nil{
		return false
	}

	if permission == "delete"{
		if permissionType.Delete {
			return true
		}else {
			return false
		}
	}else if permission == "add"{
		if permissionType.Add {
			return true
		}else {
			return false
		}
	}else if permission == "update"{
		if permissionType.Update {
			return true
		}else {
			return false
		}
	}else if permission == "disable"{
		if permissionType.Disable {
			return true
		}else {
			return false
		}
	}else if permission == "enable"{
		if permissionType.Enable {
			return true
		}else {
			return false
		}
	}else if permission == "logs"{
		if permissionType.Logs {
			return true
		}else {
			return false
		}
	}


}
