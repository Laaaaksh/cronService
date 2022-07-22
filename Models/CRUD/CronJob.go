package CRUD

import (
	"cronService/Database"
	"cronService/Models"
)

func ChangeCronJobStatus(cronjob *Models.CronJob) (err error) {

	if cronjob.Status == 0 { // changing cronjob status 0 means enable and 1 means disable
		cronjob.Status = 1
	} else {
		cronjob.Status = 0
	}
	Database.DB.Save(cronjob) // save

	return nil
}

func UpdateCronJob(cronjob *Models.CronJob) (err error) {

	Database.DB.Save(cronjob)
	return nil

}

func DeleteCronJob(cronjob *Models.CronJob, id string) (err error) {

	Database.DB.Where("id = ?", id).Delete(cronjob)
	return nil
}
