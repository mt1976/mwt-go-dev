package jobs

import (
	application "github.com/mt1976/mwt-go-dev/application"
)

// RunJobHeartBeat is a HeartBeat function
func RunJobHeartBeat(actionType string) {
	//	funcName := "RunJobHeartBeat"
	//logit(actionType, ".")
	application.UpdateSchedule("heartbeat", Monitor, "")
}
