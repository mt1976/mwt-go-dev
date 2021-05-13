package jobs

import (
	"log"
	"runtime"
	"strings"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"

	cron "github.com/robfig/cron/v3"
)

//CONST_CONFIG_FILE is cheese

// Start Initialises the Required Jobs
func Start() {
	//funcName = "Start"
	//logit("JobHandler", "CRON SCHEDULER")
	c := cron.New()
	// Insert Jobs Here
	var period string
	var runType string

	period = "*/5 * * * *"
	application.RegisterSchedule("heartbeat", "HeartBeat", "System Heartbeat", period, globals.Monitor)
	c.AddFunc(period, func() { RunJobHeartBeat("heartbeat") })
	//logit("RunJobHeartBeat", period)

	if !globals.IsChildInstance {
		period = "*/10 7-19 * * 1-5"
		application.RegisterSchedule("fxspot", "FX Spot", "FX Spot rate from barchart.com", period, globals.Aquirer)
		c.AddFunc(period, func() { RunJobFXSPOT("fxspot") })
	}

	if !globals.IsChildInstance {
		period = "30 16 * * 1-5"
		application.RegisterSchedule("ecbrate", "ECB Rates", "ECB  Benchmark FX Spot rate", period, globals.Aquirer)
		c.AddFunc(period, func() { RunJobECB("ecbrate") })
		//logit("RunJobHeartBeat", period)
	}
	if !globals.IsChildInstance {
		period = "30 10 * * 1-5"
		application.RegisterSchedule("sonia", "BOE Sonia", "Current SONIA", period, globals.Aquirer)
		c.AddFunc(period, func() { RunJobBOESONIA("sonia") })
		//logit("RunJobHeartBeat", period)
	}
	if !globals.IsChildInstance {
		period = "58 7-19 * * 1-5"
		application.RegisterSchedule("fred", "Fred Series", "Current St Louis Fed Data", period, globals.Aquirer)
		c.AddFunc(period, func() { RunJobFRED("fred") })
		//logit("RunJobHeartBeat", period)
	}

	period = "*/15 * * * *"
	application.RegisterSchedule("sessionhousekeeping", "Session Management", "Manage Application Sessions", period, globals.Monitor)
	c.AddFunc(period, func() { RunJobSessionHousekeeping("sessionhousekeeping") })
	//logit("RunJobSessionHousekeeping", period)
	period = "10 7-19 * * 1-5"
	application.RegisterSchedule("refresh", "Refresh", "Refresh all pricing data", period, globals.Dispatcher)
	c.AddFunc(period, func() { RunJobDispatch("refresh") })
	//logit("RunJobDispatch", period)

	period = "*/10 6-21 * * 1-5"
	runType = "MARKET"
	application.RegisterSchedule(runType, runType, "Refresh all "+runType+" data", period, globals.Dispatcher)
	c.AddFunc(period, func() { RunJobDispatchType(runType) })
	//logit("RunJobDispatchType-"+runType, period)

	period = "35 17 * * 1-5"
	runType = "EONIA"
	application.RegisterSchedule(runType, runType, "Refresh all "+runType+" data", period, globals.Dispatcher)
	c.AddFunc(period, func() { RunJobDispatchType(runType) })
	//logit("RunJobDispatchType-"+runType, period)

	period = "35 11 * * 1-5"
	runType = "SONIA"
	application.RegisterSchedule(runType, runType, "Refresh all "+runType+" data", period, globals.Dispatcher)
	c.AddFunc(period, func() { RunJobDispatchType(runType) })
	//logit("RunJobDispatchType-"+runType, period)

	period = "35 17 * * 1-5"
	runType = "SOFR"
	application.RegisterSchedule(runType, runType, "Refresh all "+runType+" data", period, globals.Dispatcher)
	c.AddFunc(period, func() { RunJobDispatchType(runType) })
	//logit("RunJobDispatchType-"+runType, period)

	period = "35 17 * * 1-5"
	runType = "ESTR"
	application.RegisterSchedule(runType, runType, "Refresh all "+runType+" data", period, globals.Dispatcher)
	c.AddFunc(period, func() { RunJobDispatchType(runType) })
	//logit("RunJobDispatchType-"+runType, period)

	period = "35 17 * * 1-5"
	runType = "TONAR"
	application.RegisterSchedule(runType, runType, "Refresh all "+runType+" data", period, globals.Dispatcher)
	c.AddFunc(period, func() { RunJobDispatchType(runType) })
	//logit("RunJobDispatchType-"+runType, period)

	period = "35 17 * * 1-5"
	runType = "EURIBOR"
	application.RegisterSchedule(runType, runType, "Refresh all "+runType+" data", period, globals.Dispatcher)
	c.AddFunc(period, func() { RunJobDispatchType(runType) })
	//logit("RunJobDispatchType-"+runType, period)

	period = "35 16 * * 1-5"
	runType = "ECB"
	application.RegisterSchedule(runType, runType, "Refresh all "+runType+" data", period, globals.Dispatcher)
	c.AddFunc(period, func() { RunJobDispatchType(runType) })
	//logit("RunJobDispatchType-"+runType, period)

	period = "*/30 6-21 * * 1-5"
	runType = "NI"
	application.RegisterSchedule(runType, runType, "Refresh all "+runType+" data", period, globals.Dispatcher)
	c.AddFunc(period, func() { RunJobDispatchType(runType) })
	//logit("RunJobDispatchType-"+runType, period)

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
