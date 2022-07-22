package Models

type CronExecutionResult struct {
	CronJobId     int    `json:"cron_job_id"`
	URL           string    `json:"url"`
	Output        string    `json:"output"`
	Error         string    `json:"error"`
	Time          int64 `json:"time"`
	StartTime     int64 `json:"start_time"`
	ExecutionTime int64 `json:"execution_time"`
	Status        int       `json:"status"`
}

func (ce *CronExecutionResult) TableName() string {
	return "cronexecutionresult"
}
