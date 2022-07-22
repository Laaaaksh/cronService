package CRUD

import (
	"cronService/Database"
	"cronService/Models"
)

func CreatePermissions(permission *Models.PermissionType) (err error) {
	err = Database.DB.Create(permission).Error
	if err != nil {
		return err
	}
	return nil
}
