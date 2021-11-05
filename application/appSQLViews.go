package application

import (
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
)

// SQLInjectionHander
func SQLInjectionHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	core.CreateDatabaseObjects(core.SienaDB, core.SienaPropertiesDB, "/config/database/views", false)
	core.HomePageHandler(w, r)
}
