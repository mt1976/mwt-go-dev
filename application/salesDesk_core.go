package application
// ----------------------------------------------------------------
// Automatically generated  "/application/salesdesk.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : SalesDesk (salesdesk)
// Endpoint 	        : SalesDesk (Desk)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 28/11/2021 at 22:55:00
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"html/template"
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//salesdesk_PageList provides the information for the template for a list of SalesDesks
type SalesDesk_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.SalesDesk
}

//salesdesk_Page provides the information for the template for an individual SalesDesk
type SalesDesk_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
		Name string
		ReportDealsOver string
		ReportDealsOverCCY string
		AccountTransferCutOffTime string
		AccountTransferCutOffTimeTimeZone string
		AccountTransferCutOffTimeCutOffPeriod string
	
	
	
	
	
	
	
	
	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END
}

const (
	SalesDesk_Redirect = dm.SalesDesk_PathList
)

//SalesDesk_Publish annouces the endpoints available for this object
func SalesDesk_Publish(mux http.ServeMux) {
	
	
	
	
	
	
	logs.Publish("Siena", dm.SalesDesk_Title)
	
}

//SalesDesk_HandlerList is the handler for the list page
func SalesDesk_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.SalesDesk
	noItems, returnList, _ := dao.SalesDesk_GetList()


	pageDetail := SalesDesk_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.SalesDesk_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.SalesDesk_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//SalesDesk_HandlerView is the handler used to View a page
func SalesDesk_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.SalesDesk_QueryString)
	_, rD, _ := dao.SalesDesk_GetByID(searchID)

	pageDetail := SalesDesk_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.SalesDesk_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

		// 
		// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
pageDetail.Name = rD.Name
pageDetail.ReportDealsOver = rD.ReportDealsOver
pageDetail.ReportDealsOverCCY = rD.ReportDealsOverCCY
pageDetail.AccountTransferCutOffTime = rD.AccountTransferCutOffTime
pageDetail.AccountTransferCutOffTimeTimeZone = rD.AccountTransferCutOffTimeTimeZone
pageDetail.AccountTransferCutOffTimeCutOffPeriod = rD.AccountTransferCutOffTimeCutOffPeriod
// Automatically generated 28/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 28/11/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END


	t, _ := template.ParseFiles(core.GetTemplateID(dm.SalesDesk_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//SalesDesk_HandlerEdit is the handler used generate the Edit page
func SalesDesk_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.SalesDesk_QueryString)
	_, rD, _ := dao.SalesDesk_GetByID(searchID)
	
	pageDetail := SalesDesk_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.SalesDesk_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

		// 
		// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
pageDetail.Name = rD.Name
pageDetail.ReportDealsOver = rD.ReportDealsOver
pageDetail.ReportDealsOverCCY = rD.ReportDealsOverCCY
pageDetail.AccountTransferCutOffTime = rD.AccountTransferCutOffTime
pageDetail.AccountTransferCutOffTimeTimeZone = rD.AccountTransferCutOffTimeTimeZone
pageDetail.AccountTransferCutOffTimeCutOffPeriod = rD.AccountTransferCutOffTimeCutOffPeriod
// Automatically generated 28/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 28/11/2021 by matttownsend on silicon.local - END
		//

	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END

	t, _ := template.ParseFiles(core.GetTemplateID(dm.SalesDesk_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//SalesDesk_HandlerSave is the handler used process the saving of an SalesDesk
func SalesDesk_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Name"))

	var item dm.SalesDesk
	// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
		item.Name = r.FormValue(dm.SalesDesk_Name)
		item.ReportDealsOver = r.FormValue(dm.SalesDesk_ReportDealsOver)
		item.ReportDealsOverCCY = r.FormValue(dm.SalesDesk_ReportDealsOverCCY)
		item.AccountTransferCutOffTime = r.FormValue(dm.SalesDesk_AccountTransferCutOffTime)
		item.AccountTransferCutOffTimeTimeZone = r.FormValue(dm.SalesDesk_AccountTransferCutOffTimeTimeZone)
		item.AccountTransferCutOffTimeCutOffPeriod = r.FormValue(dm.SalesDesk_AccountTransferCutOffTimeCutOffPeriod)
	
	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END

	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END

	dao.SalesDesk_Store(item)	

	http.Redirect(w, r, SalesDesk_Redirect, http.StatusFound)
}

//SalesDesk_HandlerNew is the handler used process the creation of an SalesDesk
func SalesDesk_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := SalesDesk_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.SalesDesk_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

		// 
		// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
pageDetail.Name = ""
pageDetail.ReportDealsOver = ""
pageDetail.ReportDealsOverCCY = ""
pageDetail.AccountTransferCutOffTime = ""
pageDetail.AccountTransferCutOffTimeTimeZone = ""
pageDetail.AccountTransferCutOffTimeCutOffPeriod = ""
// Automatically generated 28/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 28/11/2021 by matttownsend on silicon.local - END
		//

	t, _ := template.ParseFiles(core.GetTemplateID(dm.SalesDesk_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//SalesDesk_HandlerDelete is the handler used process the deletion of an SalesDesk
func SalesDesk_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.SalesDesk_QueryString)

	dao.SalesDesk_Delete(searchID)	

	http.Redirect(w, r, SalesDesk_Redirect, http.StatusFound)
}