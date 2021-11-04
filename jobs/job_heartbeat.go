package jobs

import (
	"fmt"

	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	"github.com/robfig/cron/v3"
)

func HeartBeat_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "HEARTBEAT"
	j.Name = "HEARTBEAT"
	j.Period = "*/5 * * * *"
	j.Description = "System Heartbeat"
	j.Type = core.Monitor
	return j
}

func HeartBeat_Register(c *cron.Cron) {
	application.RegisterSchedule(HeartBeat_Job())
	c.AddFunc(HeartBeat_Job().Period, func() { HeartBeat_Run() })
}

// RunJobHeartBeat is a HeartBeat function
func HeartBeat_Run() {

	logStart(HeartBeat_Job().Name)
	core.Log_uptime()
	core.ApplicationDB = core.GlobalsDatabasePoke(core.ApplicationDB, core.ApplicationPropertiesDB)
	core.SienaDB = core.GlobalsDatabasePoke(core.SienaDB, core.SienaPropertiesDB)

	message := fmt.Sprintf("Uptime = %v", core.Uptime())

	application.UpdateSchedule(HeartBeat_Job(), message)
	logEnd(HeartBeat_Job().Name)
}
