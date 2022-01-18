package jobs

import (
	"fmt"
	"log"

	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	"github.com/mt1976/mwt-go-dev/logs"
	cron "github.com/robfig/cron/v3"
)

func Rollover_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "ROLLOVER"
	j.Name = "ROLLOVER"
	j.Period = "10 1 * * *"
	j.Description = "Siena System Rollover Refresh"
	j.Type = core.HouseKeeping
	return j
}

func Rollover_Register(c *cron.Cron) {
	//application.RegisterSchedule(Rollover_Job().ID, Rollover_Job().Name, Rollover_Job().Description, Rollover_Job().Period, Rollover_Job().Type)
	application.Schedule_Register(Rollover_Job())
	c.AddFunc(Rollover_Job().Period, func() { Rollover_Run() })
}

// RunJobRollover is a Rollover function
func Rollover_Run() {
	logs.StartJob(Rollover_Job().Name)
	/// CONTENT STARTS
	core.Log_uptime()

	core.SienaDB = core.Database_Poke(core.SienaDB, core.SienaPropertiesDB)
	oldSysDate := core.SienaSystemDate
	_, tempDate, _ := application.GetBusinessDate(core.SienaDB)
	core.SienaSystemDate = tempDate

	log.Printf("Old System Date: %v\n", oldSysDate)
	log.Printf("New System Date: %v\n", core.SienaSystemDate)

	message := fmt.Sprintf("Rolled from %v to %v", oldSysDate.Internal, core.SienaSystemDate.Internal)

	//application.UpdateSchedule("SRO", core.HouseKeeping, message)

	/// CONTENT ENDS
	application.Schedule_Update(Rollover_Job(), message)
	core.Notification_Normal(Rollover_Job().Name, message)
	logs.EndJob(Rollover_Job().Name)
}
