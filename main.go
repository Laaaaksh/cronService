package main

import (
	"cronService/Database"
	"cronService/Models"
	"cronService/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)
var err error

func main() {
	Database.DB, err = gorm.Open("mysql", Database.DbURL(Database.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Database.DB.Close()
	Database.DB.AutoMigrate(&Models.User{})
	Database.DB.AutoMigrate(&Models.UserAuthentication{})
	Database.DB.AutoMigrate(&Models.CronJob{})
	Database.DB.AutoMigrate(&Models.CronExecutionResult{})
	Database.DB.AutoMigrate(&Models.Organization{})
	Database.DB.AutoMigrate(&Models.PermissionType{})
	r := Routes.Setuprouter()
	//running
	r.Run(":8080")
}
