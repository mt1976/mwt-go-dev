package monitors

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/mt1976/mwt-go-dev/adaptor"
	"github.com/mt1976/mwt-go-dev/core"
	"github.com/mt1976/mwt-go-dev/logs"
)

func StaticDataImporter_Watch() {
	logs.Success("StaticDataImporter_Start")

	loc := core.SienaProperties["static_in"]
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

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				// if event.Op&fsnotify.Write == fsnotify.Write {
				// 	fart(event.Name)
				// 	log.Println("modified file:", event.Name)
				// }
				if event.Op&fsnotify.Create == fsnotify.Create {
					fart(event.Name)
					log.Println("created file:", event.Name)
					err := adaptor.StaticImportProcessResponse(event.Name)
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

func fart(in string) {
	logs.Success("Poop" + in)
}
