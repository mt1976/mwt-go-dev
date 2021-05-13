package jobs

import (
	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

// RunJobHeartBeat is a HeartBeat function
func RunJobHeartBeat(actionType string) {
	//	funcName := "RunJobHeartBeat"
	//logit(actionType, ".")

	globals.ApplicationDB = globals.GlobalsDatabasePoke(globals.ApplicationDB, globals.ApplicationPropertiesDB)
	globals.SienaDB = globals.GlobalsDatabasePoke(globals.SienaDB, globals.SienaPropertiesDB)
	application.UpdateSchedule("heartbeat", globals.Monitor, "")
}
