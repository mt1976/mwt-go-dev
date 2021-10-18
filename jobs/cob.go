package jobs

import (
	"fmt"

	globals "github.com/mt1976/mwt-go-dev/globals"
	"github.com/mt1976/mwt-go-dev/siena"
)

// RunJobHeartBeat is a HeartBeat function
func RunJobCob(actionType string) {
	//	funcName := "RunJobHeartBeat"
	//logit(actionType, ".")

	logStart(actionType)
	//globals.ApplicationDB = globals.GlobalsDatabasePoke(globals.ApplicationDB, globals.ApplicationPropertiesDB)
	globals.SienaDB = globals.GlobalsDatabasePoke(globals.SienaDB, globals.SienaPropertiesDB)
	//application.UpdateSchedule("heartbeat", globals.Monitor, "")
	oldSysDate := globals.SienaSystemDate
	_, tempDate, _ := siena.GetBusinessDate(globals.SienaDB)
	globals.SienaSystemDate = tempDate

	fmt.Printf("Old System Date: %v\n", oldSysDate)
	fmt.Printf("New System Date: %v\n", globals.SienaSystemDate)
	//info("System", globals.SienaProperties["name"])
	//nfo("System Date", globals.SienaSystemDate.Internal.Format(globals.DATEFORMATUSER))
	logEnd(actionType)

}
