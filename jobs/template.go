package jobs

import (
	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
	cron "github.com/robfig/cron/v3"
)

func Template_Job() globals.JobDefinition {
	var j globals.JobDefinition
	j.ID = "ROLLOVER"
	j.Name = "ROLLOVER"
	j.Period = "10 1 * * *"
	j.Description = "Siena System Rollover Refresh"
	j.Type = globals.Monitor
	return j
}

func Template_Register(c cron.Cron) {
	application.RegisterSchedule(Template_Job().ID, Template_Job().Name, Template_Job().Description, Template_Job().Period, Template_Job().Type)
	c.AddFunc(Template_Job().Period, func() { Template_Run() })
}

// RunJobRollover is a Rollover function
func Template_Run() {
	logStart(Template_Job().Name)
	var message string
	/// CONTENT STARTS

	/// CONTENT ENDS
	application.UpdateSchedule(Template_Job().Name, Template_Job().Type, message)
	logEnd(Template_Job().Name)
}
