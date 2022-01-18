package monitors

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/mt1976/mwt-go-dev/adaptor"
	"github.com/mt1976/mwt-go-dev/application"
	"github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	"github.com/mt1976/mwt-go-dev/logs"
)

// Simulator_SienaStaticDataImporter_Job defines the job properties, name, period etc..
func Simulator_SienaStaticDataImporter_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "Simulator_SienaStaticDataImporter"
	j.Name = "Simulator_SienaStaticDataImporter"
	j.Period = ""
	j.Description = "Static Data Response processing"
	j.Type = core.Monitor
	return j
}

func Simulator_SienaStaticDataImporter_Watch() {
	logs.Success("Simulator_SienaStaticDataImporter_Start")

	loc := core.SienaProperties["static_out"]
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

	application.Schedule_Register(Simulator_SienaStaticDataImporter_Job())

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
					err := adaptor.Simulator_SienaStaticDataImporter_ProcessResponse(event.Name)
					if err != nil {
						logs.Error("SDI", err)
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
