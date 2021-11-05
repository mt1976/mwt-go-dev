package application

import (
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
)

//SQLInjection_Publish Publishes the endpoint
func SQLInjection_Publish(mux http.ServeMux) {
	core.LOG_mux("Application", "SQLInjection")
	mux.HandleFunc("/injectSQLViews/", SQLInjection_HandlerRun)
}

// SQLInjectionHander
func SQLInjection_HandlerRun(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	core.CreateDatabaseObjects(core.SienaDB, core.SienaPropertiesDB, "/config/database/views", false)
	HomePageHandler(w, r)
}
