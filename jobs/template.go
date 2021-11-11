package jobs

import (
	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	"github.com/mt1976/mwt-go-dev/logs"
	cron "github.com/robfig/cron/v3"
)

func Template_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "ROLLOVER"
	j.Name = "ROLLOVER"
	j.Period = "10 1 * * *"
	j.Description = "Siena System Rollover Refresh"
	j.Type = core.Monitor
	return j
}

func Template_Register(c cron.Cron) {
	application.RegisterSchedule(Template_Job())
	c.AddFunc(Template_Job().Period, func() { Template_Run() })
}

// RunJobRollover is a Rollover function
func Template_Run() {
	logs.StartJob(Template_Job().Name)
	var message string
	/// CONTENT STARTS

	/// CONTENT ENDS
	application.UpdateSchedule(Template_Job(), message)
	logs.EndJob(Template_Job().Name)
}
