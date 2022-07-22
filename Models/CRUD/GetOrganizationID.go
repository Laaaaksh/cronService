package CRUD

import (
	"cronService/Database"
	"cronService/Models"
	"fmt"
)

func GetOrgID(username string) (int, error) {
	var userAuth Models.UserAuthentication
	var user Models.User
	if err := Database.DB.Where("user_name = ?", username).Find(&userAuth).Error; err != nil {
		return 0, fmt.Errorf("user not found")
	}
	if err := Database.DB.Where("id = ?", userAuth.UserID).Find(&user).Error; err != nil {
		return 0, fmt.Errorf("user not found")
	}
	return user.OrganisationID, nil
}
