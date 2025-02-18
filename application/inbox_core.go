package application
// ----------------------------------------------------------------
// Automatically generated  "/application/inbox.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Inbox (inbox)
// Endpoint 	        : Inbox (Message)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:53
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





//Inbox_Publish annouces the endpoints available for this object
func Inbox_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Inbox_Path, Inbox_Handler)
	mux.HandleFunc(dm.Inbox_PathList, Inbox_HandlerList)
	mux.HandleFunc(dm.Inbox_PathView, Inbox_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	mux.HandleFunc(dm.Inbox_PathSave, Inbox_HandlerSave)
	mux.HandleFunc(dm.Inbox_PathDelete, Inbox_HandlerDelete)
	logs.Publish("Application", dm.Inbox_Title)
    core.Catalog_Add(dm.Inbox_Title, dm.Inbox_Path, "", dm.Inbox_QueryString, "Application")
}


//Inbox_HandlerList is the handler for the list page
func Inbox_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Inbox
	noItems, returnList, _ := dao.Inbox_GetList()

	pageDetail := dm.Inbox_PageList{
		Title:            CardTitle(dm.Inbox_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Inbox_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Inbox_TemplateList, w, r, pageDetail)

}


//Inbox_HandlerView is the handler used to View a page
func Inbox_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Inbox_QueryString)
	_, rD, _ := dao.Inbox_GetByID(searchID)

	pageDetail := dm.Inbox_Page{
		Title:       CardTitle(dm.Inbox_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Inbox_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = inbox_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Inbox_TemplateView, w, r, pageDetail)

}



//Inbox_HandlerSave is the handler used process the saving of an Inbox
func Inbox_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("MailId"))

	var item dm.Inbox
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Inbox_SYSId_scrn)
		item.SYSCreated = r.FormValue(dm.Inbox_SYSCreated_scrn)
		item.SYSWho = r.FormValue(dm.Inbox_SYSWho_scrn)
		item.SYSHost = r.FormValue(dm.Inbox_SYSHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Inbox_SYSUpdated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Inbox_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Inbox_SYSCreatedHost_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Inbox_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Inbox_SYSUpdatedHost_scrn)
		item.MailId = r.FormValue(dm.Inbox_MailId_scrn)
		item.MailTo = r.FormValue(dm.Inbox_MailTo_scrn)
		item.MailFrom = r.FormValue(dm.Inbox_MailFrom_scrn)
		item.MailSource = r.FormValue(dm.Inbox_MailSource_scrn)
		item.MailSent = r.FormValue(dm.Inbox_MailSent_scrn)
		item.MailUnread = r.FormValue(dm.Inbox_MailUnread_scrn)
		item.MailSubject = r.FormValue(dm.Inbox_MailSubject_scrn)
		item.MailContent = r.FormValue(dm.Inbox_MailContent_scrn)
		item.MailAcknowledged = r.FormValue(dm.Inbox_MailAcknowledged_scrn)
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Inbox_Store(item,r)	
	http.Redirect(w, r, dm.Inbox_Redirect, http.StatusFound)
}



//Inbox_HandlerDelete is the handler used process the deletion of an Inbox
func Inbox_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Inbox_QueryString)

	dao.Inbox_Delete(searchID)	

	http.Redirect(w, r, dm.Inbox_Redirect, http.StatusFound)
}


// Builds/Popuplates the Inbox Page 
func inbox_PopulatePage(rD dm.Inbox, pageDetail dm.Inbox_Page) dm.Inbox_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSWho = rD.SYSWho
	pageDetail.SYSHost = rD.SYSHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.MailId = rD.MailId
	pageDetail.MailTo = rD.MailTo
	pageDetail.MailFrom = rD.MailFrom
	pageDetail.MailSource = rD.MailSource
	pageDetail.MailSent = rD.MailSent
	pageDetail.MailUnread = rD.MailUnread
	pageDetail.MailSubject = rD.MailSubject
	pageDetail.MailContent = rD.MailContent
	pageDetail.MailAcknowledged = rD.MailAcknowledged
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.MailUnread_lookup = dao.StubLists_Get("tf")
	
	
	
	
	
	
	
	pageDetail.MailAcknowledged_lookup = dao.StubLists_Get("tf")
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSWho_props = rD.SYSWho_props
	pageDetail.SYSHost_props = rD.SYSHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.MailId_props = rD.MailId_props
	pageDetail.MailTo_props = rD.MailTo_props
	pageDetail.MailFrom_props = rD.MailFrom_props
	pageDetail.MailSource_props = rD.MailSource_props
	pageDetail.MailSent_props = rD.MailSent_props
	pageDetail.MailUnread_props = rD.MailUnread_props
	pageDetail.MailSubject_props = rD.MailSubject_props
	pageDetail.MailContent_props = rD.MailContent_props
	pageDetail.MailAcknowledged_props = rD.MailAcknowledged_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	