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

// Simulator_SienaDealImporter_Job defines the job properties, name, period etc..
func Simulator_SienaDealImporter_Job() dm.JobDefinition {
	// ----------------------------------------------------------------
	// Job Definition
	// ----------------------------------------------------------------
	var j dm.JobDefinition
	j.ID = "Simulator_SienaDealImporter"
	j.Name = "Simulator_SienaDealImporter"
	j.Period = ""
	j.Description = "Deal Importer processing"
	j.Type = core.Monitor
	return j
}

func Simulator_SienaDealImporter_Watch() {
	logs.Success("Simulator_SienaDealImporter_Start")

	loc := core.SienaProperties["transactional_out"]
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

	application.Schedule_Register(Simulator_SienaDealImporter_Job())

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
					err := adaptor.Simulator_SienaDealImporter_ProcessResponse(event.Name)
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
