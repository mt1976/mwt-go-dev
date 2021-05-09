package application

import (
	"net/http"

	globals "github.com/mt1976/mwt-go-dev/globals"
)

// SQLInjectionHander
func SQLInjectionHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	GenerateSourceViews(globals.SienaDB)
	HomePageHandler(w, r)
}
