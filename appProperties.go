package main

import (
	"log"
	"os"

	"github.com/jimlawless/cfg"
)

func getProperties(inPropertiesFile string) map[string]string {
	wctProperties := make(map[string]string)
	machineName, _ := os.Hostname()
	propertiesFileName := "config/" + inPropertiesFile
	localisedFileName := "config/" + machineName + "/" + inPropertiesFile
	//	log.Println("Testing", localisedFileName, fileExists(localisedFileName))
	//	log.Println("Testing", propertiesFileName, fileExists(propertiesFileName))
	if fileExists(localisedFileName) {
		propertiesFileName = localisedFileName
	}
	log.Println("Using Properties :", propertiesFileName)

	err := cfg.Load(propertiesFileName, wctProperties)

	if err != nil {
		log.Fatal(err)
	}
	return wctProperties
}
