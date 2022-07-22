package main

import (
	"bytes"
	"cronService/Routes"
	"github.com/goccy/go-json"
	"math"
	"strconv"

	//"bufio"
	"cronService/Database"
	"cronService/Models"
	"net/http"

	//"github.com/gin-gonic/gin/binding"

	//"encoding/json"
	"fmt"
	//"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"time"
)

var err error

func extractTime(expression string)int64 {

	var nextTime int
	if len(expression)<10{
		//print("hey")
		return 10000000
	}
	for i:=0; i<5;i++{
		if i<3 {
			temp, _ := strconv.Atoi(expression[2*i : 2*(i+1)])
			nextTime += temp * (int(math.Pow(60, float64(i))))
			continue
		}
		if i==3 {
			temp, _ := strconv.Atoi(expression[2*i : 2*(i+1)])
			nextTime += temp * 3600*24
		}
		if i==4{
			temp, _ := strconv.Atoi(expression[2*i : 2*(i+1)])
			nextTime += temp * 3600*24*30
		}
	}
	return int64(nextTime)
}

func main() {
	Database.DB, err = gorm.Open("mysql", Database.DbURL(Database.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	fmt.Println("infiniteloop1")
	defer Database.DB.Close()
	Database.DB.AutoMigrate(&Models.User{})
	Database.DB.AutoMigrate(&Models.UserAuthentication{})
	Database.DB.AutoMigrate(&Models.CronJob{})
	Database.DB.AutoMigrate(&Models.CronExecutionResult{})
	Database.DB.AutoMigrate(&Models.Organization{})
	Database.DB.AutoMigrate(&Models.PermissionType{})
	go func() {
		infinityTime := time.Unix(1<<63-62135596801, 999999999).Unix()
		for {
			var cronjobs []Models.CronJob
			Database.DB.Where("next_time < ?", time.Now().Unix()).Find(&cronjobs)
			for _, cronjob := range cronjobs {

				if cronjob.Status == 1 {
					cronjob.NextTime = infinityTime
					Database.DB.Save(&cronjob)
				} else {
					updatetime := extractTime(cronjob.Expression)

					// execute the job

					var newlog Models.CronExecutionResult
					newlog.CronJobId = cronjob.Id
					newlog.URL = cronjob.URL
					newlog.Time = cronjob.CreatedAt
					newlog.StartTime = time.Now().Unix()

					if cronjob.HttpMethod == "GET" {
						resp, err1 := http.Get(cronjob.URL)
						if err1 != nil {
							newlog.Status = 1
							newlog.Error = "can not execute"
							newlog.ExecutionTime = time.Now().Unix()
							newlog.Output = "can not execute the cronjob"
							Database.DB.Create(&newlog)

							if cronjob.RetryCount < 3 {
								cronjob.NextTime = cronjob.NextTime + 60
								cronjob.RetryCount += 1
								Database.DB.Save(&cronjob)
							} else {
								cronjob.NextTime = cronjob.NextTime - int64(cronjob.RetryCount*60) + updatetime
								cronjob.RetryCount = 0

								Database.DB.Save(&cronjob)
							}

							//panic(err1)
						} else {
							newlog.Status = 0
							newlog.ExecutionTime = time.Now().Unix()
							newlog.Output = resp.Status
							Database.DB.Create(&newlog)

							cronjob.NextTime = cronjob.NextTime + updatetime
							Database.DB.Save(&cronjob)

						}
					}

					if cronjob.HttpMethod == "POST" {

						postBody, _ := json.Marshal(cronjob.PostData)
						responseBody := bytes.NewBuffer(postBody)

						resp, err1 := http.Post(cronjob.URL, "application/json", responseBody)

						if err1 != nil {
							newlog.Status = 1
							newlog.Error = "can not execute"
							newlog.ExecutionTime = time.Now().Unix()
							newlog.Output = "can not execute the cronjob"
							Database.DB.Create(&newlog)

							if cronjob.RetryCount < 3 {
								cronjob.NextTime = cronjob.NextTime + 60
								cronjob.RetryCount += 1
								Database.DB.Save(&cronjob)
							} else {
								cronjob.NextTime = cronjob.NextTime - int64(cronjob.RetryCount*60) + updatetime
								cronjob.RetryCount = 0

								Database.DB.Save(&cronjob)
							}
							//panic(err1)
						} else {
							newlog.Status = 0
							newlog.ExecutionTime = time.Now().Unix()
							newlog.Output = resp.Status
							Database.DB.Create(&newlog)
							cronjob.NextTime = cronjob.NextTime + updatetime
							Database.DB.Save(&cronjob)

						}

					}

				}

			}
		}
	}()
	r := Routes.Setuprouter()
	//running
	r.Run(":8080")


}
