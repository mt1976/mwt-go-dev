package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon.ico")
}

func renderTemplate(w http.ResponseWriter, tmpl string, p HomePage) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func getURLparam(r *http.Request, paramID string) string {
	fmt.Println(paramID)
	fmt.Println(r.URL)
	key := r.FormValue(paramID)
	log.Println("Url Param '" + paramID + "' is: " + string(key))
	return key
}

func doSnooze(inPollingInterval string) {
	pollingInterval, _ := strconv.Atoi(inPollingInterval)
	fmt.Println("Snoooze... Zzzzzz.... ", pollingInterval)
	time.Sleep(time.Duration(pollingInterval) * time.Second)
}
