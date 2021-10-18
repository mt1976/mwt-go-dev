package jobs

import (
	"log"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

func RunJobSessionHousekeeping(actionType string) {
	logStart(actionType)
	//logit(actionType, "*** START ***")
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

	application.UpdateSchedule("sessionhousekeeping", globals.Monitor, message)
	//logit(actionType, "*** DONE ***")
	logEnd(actionType)
}
