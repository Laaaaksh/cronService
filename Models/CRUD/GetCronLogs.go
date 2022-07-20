package CRUD

import (
	"cronService/Database"
	"cronService/Models"
)

func GetCronLogs(logs *Models.CronJob, id string)(err error){
	if err = Database.DB.Where( "id = ?", id).Find(&logs).Error; err!=nil{
		return err
	}
	return nil
}
