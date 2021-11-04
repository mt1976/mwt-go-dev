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
	overideJ := DataDispatcher_Job()
	overideJ.Description = fmt.Sprintf(DataDispatcher_Job().Description, actionType)
	overideJ.Period = period
	overideJ.Name = DataDispatcher_Job().Name + "_" + actionType
	overideJ.ID = DataDispatcher_Job().ID + "_" + actionType

	application.RegisterSchedule(overideJ)
	c.AddFunc(DataDispatcher_Job().Period, func() { DataDispatcher_Run(overideJ.Name) })
}

// RunJobRollover is a Rollover function
func DataDispatcher_Run(actionType string) {
	var message string
	/// CONTENT STARTS
	overideJ := DataDispatcher_Job()
	overideJ.Name = DataDispatcher_Job().Name + "_" + actionType
	overideJ.ID = DataDispatcher_Job().ID + "_" + actionType
	logStart(overideJ.Name)

	DispatchByType("RV" + actionType)

	/// CONTENT ENDS
	application.UpdateSchedule(overideJ, message)
	logEnd(overideJ.Name + " - " + actionType)
}

func DispatchByType(dispatchType string) {
	log.Printf("Dispatch      : Issue %q -> %q \n", dispatchType, globals.SienaProperties["rates"])

	/// BRANCH HERE FOR DELIVERY METHODS

}
