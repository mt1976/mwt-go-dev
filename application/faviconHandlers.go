package application

import "net/http"

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon.ico")
}

func Favicon16Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon-16x16.png")
}

func Favicon32Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon-32x32.png")
}

func FaviconManifestHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "site.webmanifest")
}

func FaviconBrowserConfigHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "browserconfig.xml")
}
