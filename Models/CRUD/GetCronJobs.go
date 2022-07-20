package CRUD

import (
	"cronService/Database"
	"cronService/Models"
)


// getting all the cronjobs there in DB ---- perhaps we should add org_id to get cronjobs of certain org //

func GetCronJobs(CronJobs *[]Models.CronJob) (err error) {
	err = Database.DB.Find(CronJobs).Error
	if err != nil {
		return err
	}
	return nil
}