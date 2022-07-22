package CRUD

import (
	"cronService/Database"
	"cronService/Models"
	"fmt"
	"time"
)

func CreateUser(userauth * Models.UserUserauth, id int) (err error) {
	createdTime := time.Now().Unix()
	var existing_user Models.UserAuthentication
	fmt.Println("username",userauth.UserName)
	err = Database.DB.First(&existing_user,"user_name = ?", userauth.UserName).Error
	fmt.Println("createuser_models",err)
	if err == nil{
		return fmt.Errorf("username is not available")
	}
	//
	user := Models.User{UserType: "User", OrganisationID: id, FirstName: userauth.FirstName, CreatedAt: createdTime}
	if err := Database.DB.Create(&user).Error; err != nil{
		return err
	}
	userAuth := Models.UserAuthentication{UserID: user.Id, UserName: userauth.UserName, HashPassword: userauth.HashPassword, CreatedAt: createdTime}
	if err := Database.DB.Create(&userAuth).Error; err != nil{
		return err
	}
	return nil
}

