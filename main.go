package main

import "C"
import (
	"cronService/Controller"
	"cronService/Database"
	"cronService/Models"
	"cronService/Models/CRUD"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/robfig/cron/v3"
	"time"
)
var err error
func Test() {
	fmt.Println("i was call")
}
func RunCronJobs(){
	c := cron.New()
	c.AddFunc("@every 1h", Test)
	c.Start()
}
func main(){
	// Opening the database
	Database.DB, err = gorm.Open("mysql", Database.DbURL(Database.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	for {
		time.Sleep(time.Minute)

		var apiData []Models.CronJob
		CRUD.GetCronJobs(& apiData)
		
		for {
			apiData.
		}


	}

}
