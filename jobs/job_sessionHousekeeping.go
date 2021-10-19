package jobs

import (
	"log"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
	cron "github.com/robfig/cron/v3"
)

func SessionHouseKeeping_Job() globals.JobDefinition {
	var j globals.JobDefinition
	j.ID = "SESSION"
	j.Name = "SESSION"
	j.Period = "*/15 * * * *"
	j.Description = "Session Management"
	j.Type = globals.Monitor
	return j
}

func SessionHouseKeeping_Register(c *cron.Cron) {
	application.RegisterSchedule(SessionHouseKeeping_Job().ID, SessionHouseKeeping_Job().Name, SessionHouseKeeping_Job().Description, SessionHouseKeeping_Job().Period, SessionHouseKeeping_Job().Type)
	c.AddFunc(SessionHouseKeeping_Job().Period, func() { SessionHouseKeeping_Run() })
}

// RunJobRollover is a Rollover function
func SessionHouseKeeping_Run() {
	logStart(SessionHouseKeeping_Job().Name)
	/// CONTENT STARTS

	_, err := application.HousekeepSessionStore()
	// handle the error if there is one
	if err != nil {
		log.Println(err)
		panic(err)
	}
	message := ""
	if err != nil {
		message = err.Error()
	}

	/// CONTENT ENDS
	application.UpdateSchedule(SessionHouseKeeping_Job().Name, SessionHouseKeeping_Job().Type, message)
	logEnd(SessionHouseKeeping_Job().Name)
}
