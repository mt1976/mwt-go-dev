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
	// Verbose c := cron.New(cron.WithLogger(
	//	cron.VerbosePrintfLogger(log.New(os.Stdout, "", log.LstdFlags))))
	c := cron.New()
	// Insert Jobs Here
	var period string
	var runType string

	//logit("info", c.Location().String())

	logit("cron Locale", c.Location().String())

	period = "*/5 * * * *"
	application.RegisterSchedule("heartbeat", "HEARTBEAT", "System Heartbeat", period, globals.Monitor)

	c.AddFunc(period, func() { RunJobHeartBeat("HEARTBEAT") })

	//logit("RunJobHeartBeat", period)

	if !globals.IsChildInstance {
		period = "*/10 7-19 * * 1-5"
		application.RegisterSchedule("fxspot", "FX_BARCHART", "FX Spot rate from barchart.com", period, globals.Aquirer)
		c.AddFunc(period, func() { RunJobFXSPOT("FX_BARCHART") })
	}

	if !globals.IsChildInstance {
		period = "30 16 * * 1-5"
		application.RegisterSchedule("ecbrate", "FX_ECB", "ECB  Benchmark FX Spot rate", period, globals.Aquirer)
		c.AddFunc(period, func() { RunJobECB("FX_ECB") })
		//logit("RunJobHeartBeat", period)
	}
	if !globals.IsChildInstance {
		period = "30 10 * * 1-5"
		application.RegisterSchedule("sonia", "SONIA_BOE", "Current SONIA", period, globals.Aquirer)
		c.AddFunc(period, func() { RunJobBOESONIA("SONIA_BOE") })
		//logit("RunJobHeartBeat", period)
	}
	if !globals.IsChildInstance {
		period = "58 7-19 * * 1-5"
		application.RegisterSchedule("fred", "BONDS_FRED", "Instruments - Fred", period, globals.Aquirer)
		c.AddFunc(period, func() { RunJobFRED("BONDS_FRED") })
		//logit("RunJobHeartBeat", period)
	}

	//if !globals.IsChildInstance {
	//	period = "10 7-19 * * 1-5"
	//	application.RegisterSchedule("LSE", "Bonds", "Retail Bonds data from http://www.londonstockexchange.com/exchange/prices-and-markets/retail-bonds/retail-bonds-search.html", period, globals.Aquirer)
	//	c.AddFunc(period, func() { RunJobLSE("LSE") })
	//logit("RunJobHeartBeat", period)
	//}

	if !globals.IsChildInstance {
		period = "15 7-19 * * 1-5"
		application.RegisterSchedule("FII", "BONDS_FII", "Instruments - Fixed Income Investor - Enrichment", period, globals.Aquirer)
		c.AddFunc(period, func() { RunJobFII("BONDS_FII") })
		//logit("RunJobHeartBeat", period)
	}

	period = "*/15 * * * *"
	application.RegisterSchedule("sessionhousekeeping", "SESSION", "Session Management", period, globals.Monitor)
	c.AddFunc(period, func() { RunJobSessionHousekeeping("SESSION") })
	//logit("RunJobSessionHousekeeping", period)
	period = "10 7-19 * * 1-5"
	application.RegisterSchedule("refresh", "FX_RATES", "Refresh all pricing data", period, globals.Dispatcher)
	c.AddFunc(period, func() { RunJobDispatch("FX_RATES") })
	//logit("RunJobDispatch", period)

	period = "*/10 6-21 * * 1-5"
	runType = "MARKET"
	application.RegisterSchedule(runType, runType, "Refresh "+runType+" data", period, globals.Dispatcher)
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

	period = "10 1 * * *"
	application.RegisterSchedule("SRO", "ROLLOVER", "Siena System Rollover", period, globals.Monitor)
	c.AddFunc(period, func() { RunJobCob("ROLLOVER") })

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
	log.Printf("Scheduler     : %v '%v' {%v}", actionType, data, callerName)
}

func logStart(data string) {
	logit("Job Start", data)
}

func logEnd(data string) {
	logit("Job End", data)
}
