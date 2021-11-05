package application

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/lnquy/cron"
	hcron "github.com/lnquy/cron"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

func RegisterSchedule(thisJob dm.JobDefinition) {
	var s appScheduleStoreItem
	s.Id = thisJob.ID + core.IDSep + thisJob.Type
	s.Name = thisJob.Name
	s.Description = thisJob.Description
	s.Schedule = thisJob.Period
	s.Started = time.Now().Format(core.DATETIMEFORMATUSER)
	s.Lastrun = ""
	s.Message = ""
	s.SYSCreated = time.Now().Format(core.DATETIMEFORMATUSER)
	currentUserID, _ := user.Current()
	host, _ := os.Hostname()
	s.SYSWho = currentUserID.Name
	s.SYSHost = host
	s.SYSUpdated = time.Now().Format(core.DATETIMEFORMATUSER)
	s.Type = thisJob.Type
	//log.Println("STORE", s)

	registerIt := true

	if core.IsChildInstance {
		if s.Type == core.Aquirer {
			registerIt = false
		}
	}
	if registerIt {
		putScheduleStore(s)
		//desc := GetCronScheduleHuman(s.Schedule)
		op := fmt.Sprintf("Scheduled Job : %-11s %-20s %-20s %q", s.Type, s.Name, s.Schedule, GetCronScheduleHuman(s.Schedule))
		core.LOG_success(op)
	}
}

func UpdateSchedule(thisJob dm.JobDefinition, message string) {
	scheduleID := thisJob.ID + core.IDSep + thisJob.Type
	if len(scheduleID) > 1 {
		_, s, _ := GetScheduleStoreByID(scheduleID)
		if len(s.Name) > 0 {
			s.Lastrun = time.Now().Format(core.DATETIMEFORMATUSER)
			s.Message = message
			thisMess := fmt.Sprintf("Ran Job - %-11s %-20s %q", thisJob.Type, s.Name, message)
			core.Logit("Scheduler", thisMess)
			putScheduleStore(s)
		} else {
			thisMess := fmt.Sprintf("Update Schedule Schedule with '%s','%s','%s' ScheduleID = '%s'", thisJob.ID, thisJob.ID, message, scheduleID)
			core.Logit("Scheduler", thisMess)
			//spew.Dump(s)
		}
	} else {
		thisMess := fmt.Sprintf("Update Schedule Called with '%s','%s','%s'", thisJob.ID, thisJob.Type, message)
		core.Logit("Scheduler", thisMess)
	}
}

func GetCronScheduleHuman(in string) string {
	desc := ""
	if len(in) != 0 {
		exprDesc, err := hcron.NewDescriptor(
			cron.Use24HourTimeFormat(true),
			cron.DayOfWeekStartsAtOne(false),
			cron.Verbose(true),
			cron.SetLogger(log.New(os.Stdout, "cron: ", 0)),
			cron.SetLocales(hcron.Locale_en),
		)
		if err != nil {
			log.Panicf("failed to create CRON expression descriptor: %s", err)
		}
		desc, err = exprDesc.ToDescription(in, hcron.Locale_en)
		if err != nil {
			log.Panicf("failed to convert CRON expression to human readable description: %s", err)
		}
	}

	return desc
}
