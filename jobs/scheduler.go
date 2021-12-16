package jobs

import (
	"fmt"
	"runtime"
	"strings"

	core "github.com/mt1976/mwt-go-dev/core"
	"github.com/mt1976/mwt-go-dev/logs"

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
	//var period string
	//var runType string

	//logit("info", c.Location().String())

	//logit("cron Locale", c.Location().String())

	HeartBeat_Register(c)

	if !core.IsChildInstance {
		RatesFXSpot_Register(c)
	}

	if !core.IsChildInstance {
		RatesFXECB_Register(c)
	}
	if !core.IsChildInstance {
		IndexSONIABOE_Register(c)
	}
	if !core.IsChildInstance {
		InstFRED_Register(c)
	}

	if !core.IsChildInstance {
		InstFII_Register(c)
	}

	ExternalMessage_Register(c)

	SessionHouseKeeping_Register(c)
	DataDispatcher_Register(c, "MARKET", "*/10 6-21 * * 1-5")
	DataDispatcher_Register(c, "EONIA", "35 17 * * 1-5")
	DataDispatcher_Register(c, "SONIA", "35 11 * * 1-5")
	DataDispatcher_Register(c, "SOFR", "35 17 * * 1-5")
	DataDispatcher_Register(c, "ESTR", "35 17 * * 1-5")
	DataDispatcher_Register(c, "TONAR", "35 17 * * 1-5")
	DataDispatcher_Register(c, "ECB", "35 16 * * 1-5")
	DataDispatcher_Register(c, "EURIBOR", "35 17 * * 1-5")
	DataDispatcher_Register(c, "NI", "*/30 6-21 * * 1-5")
	Rollover_Register(c)

	c.Start()
	//fmt.Println(len(c.Entries()))
	//for i, e := range c.Entries() {
	//	fmt.Print(i)
	//		fmt.Printf("e: %v\n", e)
	//		fmt.Println(e.WrappedJob)
	//	}
}

func logit(actionType string, data string) {
	_, caller, _, _ := runtime.Caller(1)
	outcall := strings.Split(caller, "/")
	depth := len(outcall) - 1
	depth2 := depth - 1
	//log.Println(len(outcall), depth, depth2)
	callerName := outcall[depth2] + "/" + outcall[depth]
	op := fmt.Sprintf("%v '%v' {%v}", actionType, data, callerName)
	logs.Message("Scheduler", op)
}
