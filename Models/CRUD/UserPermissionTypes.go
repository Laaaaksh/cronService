package CRUD

import (
	"awesomeProject3/Config"
	"cronService/Models"
)

func UserPermissionTypes(permission *Models.PermissionType)(err error){
	err = Config.DB.Create(permission).Error
	if err!=nil{
		return err
	}
	return nil
}
