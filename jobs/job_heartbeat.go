package jobs

import (
	"fmt"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
	"github.com/robfig/cron/v3"
)

func HeartBeat_Job() globals.JobDefinition {
	var j globals.JobDefinition
	j.ID = "HEARTBEAT"
	j.Name = "HEARTBEAT"
	j.Period = "*/5 * * * *"
	j.Description = "System Heartbeat"
	j.Type = globals.Monitor
	return j
}

func HeartBeat_Register(c *cron.Cron) {
	application.RegisterSchedule(HeartBeat_Job())
	c.AddFunc(HeartBeat_Job().Period, func() { HeartBeat_Run() })
}

// RunJobHeartBeat is a HeartBeat function
func HeartBeat_Run() {

	logStart(HeartBeat_Job().Name)
	globals.Log_uptime()
	globals.ApplicationDB = globals.GlobalsDatabasePoke(globals.ApplicationDB, globals.ApplicationPropertiesDB)
	globals.SienaDB = globals.GlobalsDatabasePoke(globals.SienaDB, globals.SienaPropertiesDB)

	message := fmt.Sprintf("Uptime = %v", globals.Uptime())

	application.UpdateSchedule(HeartBeat_Job(), message)
	logEnd(HeartBeat_Job().Name)
}
