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

func DataDispatcher_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "DISPATCH"
	j.Name = "DISPATCH"
	j.Period = "10 1 * * *"
	j.Description = "Dispatch %q rates & pricing information"
	j.Type = core.Dispatcher
	return j
}

func DataDispatcher_Register(c *cron.Cron, actionType string, period string) {
	overideJ := DataDispatcher_Job()
	overideJ.Description = fmt.Sprintf(DataDispatcher_Job().Description, actionType)
	overideJ.Period = period
	overideJ.Name = DataDispatcher_Job().Name + "_" + actionType
	overideJ.ID = DataDispatcher_Job().ID + "_" + actionType

	application.Schedule_Register(overideJ)
	c.AddFunc(DataDispatcher_Job().Period, func() { DataDispatcher_Run(overideJ.Name) })
}

// RunJobRollover is a Rollover function
func DataDispatcher_Run(actionType string) {
	var message string
	/// CONTENT STARTS
	overideJ := DataDispatcher_Job()
	overideJ.Name = DataDispatcher_Job().Name + "_" + actionType
	overideJ.ID = DataDispatcher_Job().ID + "_" + actionType
	logs.StartJob(overideJ.Name)

	DispatchByType("RV" + actionType)

	/// CONTENT ENDS
	application.Schedule_Update(overideJ, message)
	logs.EndJob(overideJ.Name + " - " + actionType)
}

func DispatchByType(dispatchType string) {
	log.Printf("Dispatch      : Issue %q -> %q \n", dispatchType, core.SienaProperties["rates"])

	/// BRANCH HERE FOR DELIVERY METHODS

}
