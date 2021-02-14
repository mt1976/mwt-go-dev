package main

import (
	"html/template"
	"net/http"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon.ico")
}

func renderTemplate(w http.ResponseWriter, tmpl string, p HomePage) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}
