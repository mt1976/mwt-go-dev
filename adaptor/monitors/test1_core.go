package monitors

import (
	"log"

	"github.com/fsnotify/fsnotify"

	core "github.com/mt1976/mwt-go-dev/core"
	application "github.com/mt1976/mwt-go-dev/application"
	adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
)

// Simulator_Test1_Job defines the job properties, name, period etc..
func Simulator_Test1_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "Simulator_Test1"
	j.Name = "Simulator_Test1"
	j.Period = ""
	j.Description = "Simulator Test1 Monitor"
	j.Type = core.Monitor
	return j
}

func Simulator_Test1_Watch() {
	logs.Success("Simulator_Test1_Start")

	loc := core.SienaProperties["swift_recv"]
	// if first char of loc is . then remove it
	if loc[0] == '.' {
		loc = loc[1:]
	}
	loc = core.PWD + loc

	logs.Information("Watching ", loc)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	application.Schedule_Register(Simulator_Test1_Job())

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				//log.Println("event:", event)
				// if event.Op&fsnotify.Write == fsnotify.Write {
				// 	fart(event.Name)
				// 	log.Println("modified file:", event.Name)
				// }
				if event.Op&fsnotify.Create == fsnotify.Create {
					logs.Event(event.Name)
					// Create a Test1_Simulator_ProcessResponse_impl job in  adaptor/monitor/Simulator_Test1_Impl.go
					err := adaptor.Test1_Simulator_ProcessResponse_impl(event.Name)
					if err != nil {
						logs.Error("Test1", err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(loc)
	if err != nil {
		log.Fatal(err)
	}
	<-done

}
