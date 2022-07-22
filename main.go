package main

import (
	"bytes"
	//"bufio"
	"cronService/Database"
	"cronService/Models"
	"cronService/Routes"
	"encoding/json"
	"github.com/gin-gonic/gin/binding"

	//"encoding/json"
	"fmt"
	//"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"go/scanner"
	"net/http"
	"time"
)

var err error

func extractTime(expression string) time.Duration {

	var updatetime time.Duration

	// do something for getting time in second from expression


	updatetime = 5 /// its a random time
	return updatetime

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
	r := Routes.Setuprouter()
	//running
	infinityTime := time.Unix(1<<63-62135596801, 999999999)
	fmt.Println(infinityTime)

	for  {
		var cronjobs []Models.CronJob
		Database.DB.Where("next_time < ?", time.Now()).Find(&cronjobs)
		for _, cronjob := range cronjobs {

			if cronjob.Status == 1 {
				cronjob.NextTime = infinityTime
				Database.DB.Save(cronjob)
			} else {
						var updatetime time.Duration
						updatetime = extractTime(cronjob.Expression)

				        // execute the job

				        var newlog Models.CronExecutionResult
						newlog.CronJobId = cronjob.Id
						newlog.URL = cronjob.URL
						newlog.Time = cronjob.CreatedAt
						newlog.StartTime = time.Now()



				        if cronjob.HttpMethod == "GET" {
							resp, err1 := http.Get(cronjob.URL)
							if err1 != nil {
								newlog.Status = 1
								newlog.Error = "can not execute"
								newlog.ExecutionTime = time.Now()
								newlog.Output = "can not execute the cronjob"
								Database.DB.Save(newlog)

								if cronjob.RetryCount <3 {
									cronjob.NextTime = cronjob.NextTime.Add(time.Minute * 5)
									cronjob.RetryCount +=1
									Database.DB.Save(cronjob)
								}else{
									cronjob.RetryCount = 0
									cronjob.NextTime = cronjob.NextTime.Add(time.Second * updatetime)
									Database.DB.Save(cronjob)
								}



								//panic(err1)
							}else{
								newlog.Status = 0
								newlog.ExecutionTime = time.Now()
								newlog.Output = resp.Status
								Database.DB.Save(newlog)

								cronjob.NextTime = cronjob.NextTime.Add(time.Second * updatetime)
								Database.DB.Save(cronjob)

							}
							//defer resp.Body.Close()
							//fmt.Println("Response status:", resp.Status)
							//scanner := bufio.NewScanner(resp.Body)
							//for i := 0; scanner.Scan() && i < 5; i++ {
							//	fmt.Println(scanner.Text())
							//}
							//if err := scanner.Err(); err != nil {
							//	panic(err)
							//}

						}

				if cronjob.HttpMethod == "POST" {

					postBody, _ := json.Marshal(cronjob.PostData)
					responseBody := bytes.NewBuffer(postBody)

					resp, err1 := http.Post(cronjob.URL, "application/json", responseBody)


					if err1 != nil {
						newlog.Status = 1
						newlog.Error = "can not execute"
						newlog.ExecutionTime = time.Now()
						newlog.Output = "can not execute the cronjob"
						Database.DB.Save(newlog)

						if cronjob.RetryCount <3 {
							cronjob.NextTime = cronjob.NextTime.Add(time.Minute * 5)
							cronjob.RetryCount +=1
							Database.DB.Save(cronjob)
						}else{
							cronjob.RetryCount = 0
							cronjob.NextTime = cronjob.NextTime.Add(time.Second * updatetime)
							Database.DB.Save(cronjob)
						}



						//panic(err1)
					}else{
						newlog.Status = 0
						newlog.ExecutionTime = time.Now()
						newlog.Output = resp.Status
						Database.DB.Save(newlog)

						cronjob.NextTime = cronjob.NextTime.Add(time.Second * updatetime)
						Database.DB.Save(cronjob)

					}

				}






            //fmt.Println(cronjob)


			}

		}
	}
	r.Run(":8080")


}
