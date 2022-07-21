package CRUD

import (
	"cronService/Database"
	"cronService/Models"
	"time"
)


func CreateOrganization(orguser Models.OrgUser)(err error){
	createdTime := time.Now().Unix()
	org := Models.Organization{OrganisationName: orguser.OrganisationName, CreatedAt: createdTime}
	if err := Database.DB.Create(&org).Error; err != nil{
		return err
	}

	user := Models.User{UserType: "Admin", OrganisationID: org.Id, FirstName: orguser.AdminUserName, CreatedAt: createdTime}

	if err := Database.DB.Create(&user).Error; err != nil{
		return err
	}
	userauth := Models.UserAuthentication{UserID: user.Id, UserName: orguser.AdminUserName, HashPassword: orguser.Password, CreatedAt: createdTime}
	if err := Database.DB.Create(&userauth).Error; err != nil{
		return err
	}
	return nil
}
