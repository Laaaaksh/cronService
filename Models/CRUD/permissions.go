package CRUD

import (
	"cronService/Controller"
	"cronService/Database"
	"cronService/Models"
)

func CreateUserType(userType *Models.PermissionUserType) (err error){
	if err := Database.DB.Create(userType).Error; err != nil{
		return err
	}
	return nil
}

func VerifyCredentials(cred Controller.Credentials) (flag bool){
	var user Models.UserAuthentication
	if err:= Database.DB.Where("user_name = ?", cred.UserName).Find(&user).Error; err != nil{
		return false
	}
	if user.HashPassword != cred.Password{  // need to convert the password to hash to check with original
		return false
	}
	return true
}

func CreateToken(usertoken *Models.UserToken) (err error) {
	if err:= Database.DB.Create(usertoken).Error; err != nil{
		return err
	}
	return nil
}

func AuthorizeAdmin(user_name string) (flag bool){
	var user Models.User

	if err := Database.DB.Where("user_name = ?", user_name).Find(&user).Error; err != nil{
		return false
	}
	if user.UserType != "Admin"{
		return false
	}
	return true
}

