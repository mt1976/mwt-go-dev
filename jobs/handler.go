package jobs

import (
	"log"
	"runtime"
	"strings"

	application "github.com/mt1976/mwt-go-dev/application"
	cron "github.com/robfig/cron/v3"
)

//CONST_CONFIG_FILE is cheese
const (
	Scheduled = "Scheduled"
	Adhoc     = "AdHoc"
)

// Start Initialises the Required Jobs
func Start() {
	//funcName = "Start"
	logit("JobHandler", "CRON SCHEDULER")
	c := cron.New()
	// Insert Jobs Here
	var period string

	period = "@every 20s"
	application.RegisterSchedule("heartbeat", "HeartBeat", "System Heartbeat", period)
	c.AddFunc(period, func() { RunJobHeartBeat(Scheduled) })
	logit("RunJobHeartBeat", period)

	period = "@every 10m"
	application.RegisterSchedule("fxspot", "FX Spot", "FX Spot rate from barchart.com", period)
	c.AddFunc(period, func() { RunJobFXSPOT(Scheduled) })
	logit("RunJobHeartBeat", period)

	period = "30 16 * * 1-5"
	application.RegisterSchedule("ecbrate", "ECB Rates", "ECB  Benchmark FX Spot rate", period)
	c.AddFunc(period, func() { RunJobECB(Scheduled) })
	logit("RunJobHeartBeat", period)

	period = "30 10 * * 1-5"
	application.RegisterSchedule("sonia", "BOE Sonia", "Current SONIA", period)
	c.AddFunc(period, func() { RunJobBOESONIA(Scheduled) })
	logit("RunJobHeartBeat", period)

	period = "58 7-19 * * 1-5"
	application.RegisterSchedule("fred", "Fred Series", "Current St Louis Fed Data", period)
	c.AddFunc(period, func() { RunJobFRED(Scheduled) })
	logit("RunJobHeartBeat", period)

	//c.AddFunc("@every 1m", func() { RunJobFXSPOT(Scheduled) })
	//logit("RunJobFXSPOT", "@every 1m")
	//log.Println(runtime.Caller(0))
	//log.Println(c.Entries())
	c.Start()
}

func logit(actionType string, data string) {
	_, caller, _, _ := runtime.Caller(1)
	outcall := strings.Split(caller, "/")
	depth := len(outcall) - 1
	depth2 := depth - 1
	//log.Println(len(outcall), depth, depth2)
	callerName := outcall[depth2] + "/" + outcall[depth]
	log.Println(callerName, actionType, data)
}
