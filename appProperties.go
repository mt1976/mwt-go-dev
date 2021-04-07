package main

import (
	"log"

	"github.com/jimlawless/cfg"
)

func getProperties(propFile string) map[string]string {
	wctProperties := make(map[string]string)
	err := cfg.Load(propFile, wctProperties)
	if err != nil {
		log.Fatal(err)
	}
	return wctProperties
}
