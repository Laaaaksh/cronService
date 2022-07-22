package CRUD

import (
	"cronService/Database"
	"cronService/Models"
)

func GetPermissions(permissions *[]Models.PermissionType)(err error){
	err = Database.DB.Find(permissions).Error
	if err!=nil{
		return err
	}
	return nil
}

