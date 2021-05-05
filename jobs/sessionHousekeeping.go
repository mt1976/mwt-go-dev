package jobs

import (
	"log"

	application "github.com/mt1976/mwt-go-dev/application"
)

func RunJobSessionHousekeeping(actionType string) {
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
	application.UpdateSchedule("sessionhousekeeping", Monitor, message)
	//logit(actionType, "*** DONE ***")
}
