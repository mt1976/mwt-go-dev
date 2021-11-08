package jobs

import (
	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	cron "github.com/robfig/cron/v3"
)

func SessionHouseKeeping_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "SESSION"
	j.Name = "SESSION"
	j.Period = "*/15 * * * *"
	j.Description = "Session Management"
	j.Type = core.Monitor
	return j
}

func SessionHouseKeeping_Register(c *cron.Cron) {
	application.RegisterSchedule(SessionHouseKeeping_Job())
	c.AddFunc(SessionHouseKeeping_Job().Period, func() { SessionHouseKeeping_Run() })
}

// RunJobRollover is a Rollover function
func SessionHouseKeeping_Run() {
	logStart(SessionHouseKeeping_Job().Name)
	/// CONTENT STARTS

	// _, err := application.HousekeepSessionStore()
	// // handle the error if there is one
	// if err != nil {
	// 	log.Println(err)
	// 	panic(err)
	// }
	// message := ""
	// if err != nil {
	// 	message = err.Error()
	// }

	/// CONTENT ENDS
	application.UpdateSchedule(SessionHouseKeeping_Job(), "")
	logEnd(SessionHouseKeeping_Job().Name)
}
