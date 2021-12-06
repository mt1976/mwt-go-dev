package application

// ----------------------------------------------------------------
// Automatically generated  "/application/template.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyOnboarding (template)
// Endpoint 	        : CounterpartyOnboarding (CounterpartyOnboardingID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// CounterpartyOnboarding Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 28/11/2021 at 22:55:01
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
	"github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

//template_PageList provides the information for the template for a list of CounterpartyOnboardings
type CounterpartyOnboarding_PageList struct {
	SessionInfo dm.SessionInfo
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	ItemsOnPage int
	ItemList    []dm.CounterpartyOnboarding
}

//template_Page provides the information for the template for an individual CounterpartyOnboarding
type CounterpartyOnboarding_Page struct {
	SessionInfo dm.SessionInfo
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
	Item dm.CounterpartyOnboarding

	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END
}

const (
	CounterpartyOnboarding_Redirect = dm.CounterpartyOnboarding_PathList
)

//CounterpartyOnboarding_Publish annouces the endpoints available for this object
func CounterpartyOnboarding_Publish(mux http.ServeMux) {
	// mux.HandleFunc(dm.CounterpartyOnboarding_PathList, CounterpartyOnboarding_HandlerList)
	// mux.HandleFunc(dm.CounterpartyOnboarding_PathView, CounterpartyOnboarding_HandlerView)
	// mux.HandleFunc(dm.CounterpartyOnboarding_PathEdit, CounterpartyOnboarding_HandlerEdit)
	mux.HandleFunc(dm.CounterpartyOnboarding_PathNew, CounterpartyOnboarding_HandlerNew)
	// mux.HandleFunc(dm.CounterpartyOnboarding_PathSave, CounterpartyOnboarding_HandlerSave)
	// mux.HandleFunc(dm.CounterpartyOnboarding_PathDelete, CounterpartyOnboarding_HandlerDelete)
	logs.Publish("Application", dm.CounterpartyOnboarding_Title)

}

//CounterpartyOnboarding_HandlerList is the handler for the list page
func CounterpartyOnboarding_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	// inUTL := r.URL.Path
	// w.Header().Set("Content-Type", "text/html")
	// core.ServiceMessage(inUTL)

	// var returnList []dm.CounterpartyOnboarding
	// noItems, returnList, _ := dao.CounterpartyOnboarding_GetList()

	// pageDetail := CounterpartyOnboarding_PageList{
	// 	Title:       core.ApplicationProperties["appname"],
	// 	PageTitle:   PageTitle(dm.CounterpartyOnboarding_Title, core.Action_List),
	// 	ItemsOnPage: noItems,
	// 	ItemList:    returnList,
	// 	UserMenu:    UserMenu_Get(r),
	// 	UserRole:    core.GetUserRole(r),
	// }

	// t, _ := template.ParseFiles(core.GetCounterpartyOnboardingID(dm.CounterpartyOnboarding_CounterpartyOnboardingList, core.GetUserRole(r)))
	// t.Execute(w, pageDetail)
}

//CounterpartyOnboarding_HandlerView is the handler used to View a page
func CounterpartyOnboarding_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	// w.Header().Set("Content-Type", "text/html")
	// logs.Servicing(r.URL.Path)

	// searchID := core.GetURLparam(r, dm.CounterpartyOnboarding_QueryString)
	// _, rD, _ := dao.CounterpartyOnboarding_GetByID(searchID)

	// pageDetail := CounterpartyOnboarding_Page{
	// 	Title:     core.ApplicationProperties["appname"],
	// 	PageTitle: PageTitle(dm.CounterpartyOnboarding_Title, core.Action_View),
	// 	UserMenu:  UserMenu_Get(r),
	// 	UserRole:  core.GetUserRole(r),
	// }

	// //
	// // Automatically generated 28/11/2021 by matttownsend on silicon.local - START
	// pageDetail.SYSId = rD.SYSId
	// pageDetail.FIELD1 = rD.FIELD1
	// pageDetail.FIELD2 = rD.FIELD2
	// pageDetail.SYSCreated = rD.SYSCreated
	// pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	// pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	// pageDetail.SYSUpdated = rD.SYSUpdated
	// pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	// pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	// pageDetail.ID = rD.ID
	// // Automatically generated 28/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
	// // Automatically generated 28/11/2021 by matttownsend on silicon.local - END
	// //

	// // Automatically generated 28/11/2021 by matttownsend on silicon.local - END

	// t, _ := template.ParseFiles(core.GetCounterpartyOnboardingID(dm.CounterpartyOnboarding_CounterpartyOnboardingView, core.GetUserRole(r)))
	// t.Execute(w, pageDetail)

}

//CounterpartyOnboarding_HandlerEdit is the handler used generate the Edit page
func CounterpartyOnboarding_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	// searchID := core.GetURLparam(r, dm.CounterpartyOnboarding_QueryString)
	// _, rD, _ := dao.CounterpartyOnboarding_GetByID(searchID)

	// pageDetail := CounterpartyOnboarding_Page{
	// 	Title:     core.ApplicationProperties["appname"],
	// 	PageTitle: PageTitle(dm.CounterpartyOnboarding_Title, core.Action_Edit),
	// 	UserMenu:  UserMenu_Get(r),
	// 	UserRole:  core.GetUserRole(r),
	// }

	// //
	// // Automatically generated 28/11/2021 by matttownsend on silicon.local - START
	// pageDetail.SYSId = rD.SYSId
	// pageDetail.FIELD1 = rD.FIELD1
	// pageDetail.FIELD2 = rD.FIELD2
	// pageDetail.SYSCreated = rD.SYSCreated
	// pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	// pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	// pageDetail.SYSUpdated = rD.SYSUpdated
	// pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	// pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	// pageDetail.ID = rD.ID
	// // Automatically generated 28/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
	// // Automatically generated 28/11/2021 by matttownsend on silicon.local - END
	// //

	// // Automatically generated 28/11/2021 by matttownsend on silicon.local - END

	// t, _ := template.ParseFiles(core.GetCounterpartyOnboardingID(dm.CounterpartyOnboarding_CounterpartyOnboardingEdit, core.GetUserRole(r)))
	// t.Execute(w, pageDetail)

}

//CounterpartyOnboarding_HandlerSave is the handler used process the saving of an CounterpartyOnboarding
func CounterpartyOnboarding_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	// w.Header().Set("Content-Type", "text/html")
	// logs.Servicing(r.URL.Path + r.FormValue("ID"))

	// var item dm.CounterpartyOnboarding
	// // Automatically generated 28/11/2021 by matttownsend on silicon.local - START
	// item.SYSId = r.FormValue(dm.CounterpartyOnboarding_SYSId)
	// item.FIELD1 = r.FormValue(dm.CounterpartyOnboarding_FIELD1)
	// item.FIELD2 = r.FormValue(dm.CounterpartyOnboarding_FIELD2)
	// item.SYSCreated = r.FormValue(dm.CounterpartyOnboarding_SYSCreated)
	// item.SYSCreatedBy = r.FormValue(dm.CounterpartyOnboarding_SYSCreatedBy)
	// item.SYSCreatedHost = r.FormValue(dm.CounterpartyOnboarding_SYSCreatedHost)
	// item.SYSUpdated = r.FormValue(dm.CounterpartyOnboarding_SYSUpdated)
	// item.SYSUpdatedHost = r.FormValue(dm.CounterpartyOnboarding_SYSUpdatedHost)
	// item.SYSUpdatedBy = r.FormValue(dm.CounterpartyOnboarding_SYSUpdatedBy)
	// item.ID = r.FormValue(dm.CounterpartyOnboarding_ID)

	// // Automatically generated 28/11/2021 by matttownsend on silicon.local - END

	// // Automatically generated 28/11/2021 by matttownsend on silicon.local - END

	// dao.CounterpartyOnboarding_Store(item)

	// http.Redirect(w, r, CounterpartyOnboarding_Redirect, http.StatusFound)
}

//CounterpartyOnboarding_HandlerNew is the handler used process the creation of an CounterpartyOnboarding
func CounterpartyOnboarding_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := CounterpartyOnboarding_Page{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: PageTitle(dm.CounterpartyOnboarding_Title, core.Action_New),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
	}

	pageDetail.Item = dao.CounterpartyOnboarding_Build()
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	ExecuteTemplate(dm.CounterpartyOnboarding_CounterpartyOnboardingNew, w, r, pageDetail)

}

//CounterpartyOnboarding_HandlerDelete is the handler used process the deletion of an CounterpartyOnboarding
func CounterpartyOnboarding_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	//searchID := core.GetURLparam(r, dm.CounterpartyOnboarding_QueryString)

	//	dao.CounterpartyOnboarding_Delete(searchID)

	http.Redirect(w, r, CounterpartyOnboarding_Redirect, http.StatusFound)
}
