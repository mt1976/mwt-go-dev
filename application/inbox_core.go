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
// Date & Time		    : 26/06/2022 at 18:48:29
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//inbox_PageList provides the information for the template for a list of Inboxs
type Inbox_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Inbox
}
//Inbox_Redirect provides a page to return to aftern an action
const (
	
	Inbox_Redirect = dm.Inbox_PathList
	
)

//inbox_Page provides the information for the template for an individual Inbox
type Inbox_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     dm.FieldProperties
	SYSCreated         string
	SYSCreated_props     dm.FieldProperties
	SYSWho         string
	SYSWho_props     dm.FieldProperties
	SYSHost         string
	SYSHost_props     dm.FieldProperties
	SYSUpdated         string
	SYSUpdated_props     dm.FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props     dm.FieldProperties
	SYSCreatedHost         string
	SYSCreatedHost_props     dm.FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props     dm.FieldProperties
	SYSUpdatedHost         string
	SYSUpdatedHost_props     dm.FieldProperties
	MailId         string
	MailId_props     dm.FieldProperties
	MailTo         string
	MailTo_props     dm.FieldProperties
	MailFrom         string
	MailFrom_props     dm.FieldProperties
	MailSource         string
	MailSource_props     dm.FieldProperties
	MailSent         string
	MailSent_props     dm.FieldProperties
	MailUnread         string
	MailUnread_lookup    []dm.Lookup_Item
	MailUnread_props     dm.FieldProperties
	MailSubject         string
	MailSubject_props     dm.FieldProperties
	MailContent         string
	MailContent_props     dm.FieldProperties
	MailAcknowledged         string
	MailAcknowledged_lookup    []dm.Lookup_Item
	MailAcknowledged_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Inbox_Publish annouces the endpoints available for this object
func Inbox_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Inbox_Path, Inbox_Handler)
	mux.HandleFunc(dm.Inbox_PathList, Inbox_HandlerList)
	mux.HandleFunc(dm.Inbox_PathView, Inbox_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
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

	pageDetail := Inbox_PageList{
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

	pageDetail := Inbox_Page{
		Title:       CardTitle(dm.Inbox_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Inbox_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = inbox_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Inbox_TemplateView, w, r, pageDetail)

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

	http.Redirect(w, r, Inbox_Redirect, http.StatusFound)
}


// Builds/Popuplates the Inbox Page 
func inbox_PopulatePage(rD dm.Inbox, pageDetail Inbox_Page) Inbox_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
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
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	