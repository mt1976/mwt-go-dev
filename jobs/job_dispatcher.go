package jobs

import (
	"fmt"
	"log"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
	cron "github.com/robfig/cron/v3"
)

func DataDispatcher_Job() globals.JobDefinition {
	var j globals.JobDefinition
	j.ID = "DISPATCH"
	j.Name = "DISPATCH"
	j.Period = "10 1 * * *"
	j.Description = "Dispatch %q rates & pricing information"
	j.Type = globals.Dispatcher
	return j
}

func DataDispatcher_Register(c *cron.Cron, actionType string, period string) {
	ovr_description := fmt.Sprintf(DataDispatcher_Job().Description, actionType)
	ovr_period := period
	ovr_name := DataDispatcher_Job().Name + "_" + actionType
	ovr_id := DataDispatcher_Job().ID + "_" + actionType

	application.RegisterSchedule(ovr_id, ovr_name, ovr_description, ovr_period, DataDispatcher_Job().Type)
	c.AddFunc(DataDispatcher_Job().Period, func() { DataDispatcher_Run(ovr_name) })
}

// RunJobRollover is a Rollover function
func DataDispatcher_Run(actionType string) {
	logStart(DataDispatcher_Job().Name + " - " + actionType)
	var message string
	/// CONTENT STARTS

	DispatchByType("RV" + actionType)

	/// CONTENT ENDS
	application.UpdateSchedule(DataDispatcher_Job().Name, DataDispatcher_Job().Type, message)
	logEnd(DataDispatcher_Job().Name + " - " + actionType)
}

func DispatchByType(dispatchType string) {
	log.Printf("Dispatch      : Issue %q -> %q \n", dispatchType, globals.SienaProperties["rates"])

	/// BRANCH HERE FOR DELIVERY METHODS

}
