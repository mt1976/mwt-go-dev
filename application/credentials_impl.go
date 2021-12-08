package application

import (
	"net/http"
	"strconv"
	"time"

	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

func credentials_PublishImpl(mux http.ServeMux) http.ServeMux {
	return mux
}

func Credentials_HandlerBan(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, dm.Credentials_QueryString)
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	core.ServiceMessageAction(inUTL, "Ban", searchID)
	banCredentialsStore(searchID, r)
	http.Redirect(w, r, Credentials_Redirect, http.StatusFound)
}

func Credentials_HandlerActivate(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := core.GetURLparam(r, dm.Credentials_QueryString)
	if len(searchID) == 0 {
		searchID = r.FormValue("Id")
	}
	core.ServiceMessageAction(inUTL, "Activate", searchID)
	activateCredentialsStore(searchID, r)
	http.Redirect(w, r, Credentials_Redirect, http.StatusFound)

}

func getExpiryDate() string {
	expiryDate := time.Now()
	life, _ := strconv.Atoi(core.ApplicationProperties["credentialslife"])
	expiryDate = expiryDate.AddDate(0, 0, life)
	return expiryDate.Format(core.DATETIMEFORMATUSER)
}

func banCredentialsStore(id string, req *http.Request) {

	_, r, _ := dao.Credentials_GetByID(id)
	r.Expiry = ""
	r.Password = ""
	dao.Credentials_Store(r, req)
}

func activateCredentialsStore(id string, req *http.Request) {

	_, r, _ := dao.Credentials_GetByID(id)
	r.Expiry = getExpiryDate()
	dao.Credentials_Store(r, req)
}

func credentials_HandlerViewImpl(pageDetail Credentials_Page) Credentials_Page {
	return pageDetail
}

func credentials_HandlerEditImpl(pageDetail Credentials_Page) Credentials_Page {
	return pageDetail
}

func credentials_HandlerSaveImpl(item dm.Credentials) dm.Credentials {
	if item.Issued == "" {
		item.Issued = time.Now().Format(core.DATETIMEFORMATUSER)
	}
	if item.Expiry == "" {
		item.Expiry = getExpiryDate()
	}
	return item
}
