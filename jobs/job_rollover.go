package jobs

import (
	"fmt"
	"log"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
	"github.com/mt1976/mwt-go-dev/siena"
	cron "github.com/robfig/cron/v3"
)

func Rollover_Job() globals.JobDefinition {
	var j globals.JobDefinition
	j.ID = "ROLLOVER"
	j.Name = "ROLLOVER"
	j.Period = "10 1 * * *"
	j.Description = "Siena System Rollover Refresh"
	j.Type = globals.Monitor
	return j
}

func Rollover_Register(c *cron.Cron) {
	application.RegisterSchedule(Rollover_Job().ID, Rollover_Job().Name, Rollover_Job().Description, Rollover_Job().Period, Rollover_Job().Type)
	c.AddFunc(Rollover_Job().Period, func() { Rollover_Run() })
}

// RunJobRollover is a Rollover function
func Rollover_Run() {
	logStart(Rollover_Job().Name)
	/// CONTENT STARTS
	globals.Log_uptime()

	globals.SienaDB = globals.GlobalsDatabasePoke(globals.SienaDB, globals.SienaPropertiesDB)
	oldSysDate := globals.SienaSystemDate
	_, tempDate, _ := siena.GetBusinessDate(globals.SienaDB)
	globals.SienaSystemDate = tempDate

	log.Printf("Old System Date: %v\n", oldSysDate)
	log.Printf("New System Date: %v\n", globals.SienaSystemDate)

	message := fmt.Sprintf("Rolled from %v to %v", oldSysDate.Internal, globals.SienaSystemDate.Internal)

	application.UpdateSchedule("SRO", globals.Monitor, message)

	/// CONTENT ENDS
	application.UpdateSchedule(Rollover_Job().Name, Rollover_Job().Type, "")
	logEnd(Rollover_Job().Name)
}
