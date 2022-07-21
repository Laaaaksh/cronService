package CRUD

import (
	"cronService/Database"
	"cronService/Models"
	"time"
)

func CreateUser(userauth * Models.UserUserauth, id int) (err error) {
	createdTime := time.Now().Unix()
	user := Models.User{UserType: "User", OrganisationID: id, FirstName: userauth.FirstName, CreatedAt: createdTime}

	if err := Database.DB.Create(&user).Error; err != nil{
		return err
	}
	userAuth := Models.UserAuthentication{UserID: user.Id, UserName: userauth.UserName, HashPassword: userauth.HashPassword, CreatedAt: createdTime}
	if err := Database.DB.Create(&userAuth).Error; err != nil{
		return err
	}
	if err := Database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

