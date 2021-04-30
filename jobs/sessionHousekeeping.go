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
	application.UpdateSchedule(actionType, Monitor, "")
	//logit(actionType, "*** DONE ***")
}
