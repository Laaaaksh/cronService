package Models

import (
	"time"
)

type CronExecutionResult struct {
	CronJobId     int    `json:"cron_job_id"`
	URL           string    `json:"url"`
	Output        string    `json:"output"`
	Error         string    `json:"error"`
	Time          time.Time `json:"time"`
	StartTime     time.Time `json:"start_time"`
	ExecutionTime time.Time `json:"execution_time"`
	Status        int       `json:"status"`
}

func (ce *CronExecutionResult) TableName() string {
	return "cronexecutionresult"
}
