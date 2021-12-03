package application

import (
	"log"
	"os"

	"github.com/jimlawless/cfg"
	core "github.com/mt1976/mwt-go-dev/core"
)

//comment
func OSBOLETEGetProperties(inPropertiesFile string) map[string]string {
	wctProperties := make(map[string]string)
	machineName, _ := os.Hostname()
	propertiesFileName := "config/" + inPropertiesFile
	localisedFileName := "config/" + machineName + "/" + inPropertiesFile
	//	log.Println("Testing", localisedFileName, core.FileExists(localisedFileName))
	//	log.Println("Testing", propertiesFileName, core.FileExists(propertiesFileName))
	if core.FileExists(localisedFileName) {
		propertiesFileName = localisedFileName
	}
	//log.Println("Using Properties :", propertiesFileName)

	err := cfg.Load(propertiesFileName, wctProperties)

	if err != nil {
		log.Fatal(err)
	}
	return wctProperties
}
