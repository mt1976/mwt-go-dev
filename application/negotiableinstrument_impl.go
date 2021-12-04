package application

import (
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

//NegotiableInstrument_Publish annouces the endpoints available for this object
func NegotiableInstrument_Publish_Impl(mux http.ServeMux) {
	//mux.HandleFunc("/listGiltsDataStore/", application.ListLSEGiltsDataStoreHandler)
	//mux.HandleFunc("/viewLSEGiltsDataStore/", application.ViewLSEGiltsDataStoreHandler)
	mux.HandleFunc("/selectLSEGiltsDataStore/", NegotiableInstrument_HandlerSelect)
	mux.HandleFunc("/deselectLSEGiltsDataStore/", NegotiableInstrument_HandlerDeSelect)

	logs.Publish("Application", dm.NegotiableInstrument_Title)

}

func NegotiableInstrument_HandlerSelect(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, "LSEGiltsDataStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	core.ServiceMessageAction(inUTL, "Select", searchID)
	var thisRec AppNISelectedStoreItem
	thisRec.Id = searchID
	putNISelectedStore(thisRec, r)

	http.Redirect(w, r, NegotiableInstrument_Redirect, http.StatusFound)
}

func NegotiableInstrument_HandlerDeSelect(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, "LSEGiltsDataStore")
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	core.ServiceMessageAction(inUTL, "Deselect", searchID)
	//var thisRec AppNISelectedStoreItem
	//thisRec.Id = searchID
	deleteNISelectedStore(searchID)
	http.Redirect(w, r, NegotiableInstrument_Redirect, http.StatusFound)
}
