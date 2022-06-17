package application
// ----------------------------------------------------------------
// Automatically generated  "/application/externalmessage.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ExternalMessage (externalmessage)
// Endpoint 	        : ExternalMessage (Message)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:11
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//externalmessage_PageList provides the information for the template for a list of ExternalMessages
type ExternalMessage_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.ExternalMessage
}
//ExternalMessage_Redirect provides a page to return to aftern an action
const (
	ExternalMessage_Redirect = dm.ExternalMessage_PathList
)

//externalmessage_Page provides the information for the template for an individual ExternalMessage
type ExternalMessage_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	MessageID         string
	MessageFormat         string
	MessageDeliveredTo         string
	MessageBody         string
	MessageFilename         string
	MessageLife         string
	MessageDate         string
	MessageTime         string
	MessageTimeoutAction         string
	MessageACKNAK         string
	MessageACKNAK_lookup    []dm.Lookup_Item
	ResponseID         string
	ResponseFilename         string
	ResponseBody         string
	ResponseDate         string
	ResponseTime         string
	ResponseAdditionalInfo         string
	SYSCreated         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdated         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	MessageTimeout         string
	MessageEmitted         string
	ResponseRecieved         string
	MessageClass         string
	AppID         string
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//ExternalMessage_Publish annouces the endpoints available for this object
func ExternalMessage_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.ExternalMessage_Path, ExternalMessage_Handler)
	mux.HandleFunc(dm.ExternalMessage_PathList, ExternalMessage_HandlerList)
	mux.HandleFunc(dm.ExternalMessage_PathView, ExternalMessage_HandlerView)
	mux.HandleFunc(dm.ExternalMessage_PathEdit, ExternalMessage_HandlerEdit)
	//Cannot Create via GUI
	mux.HandleFunc(dm.ExternalMessage_PathSave, ExternalMessage_HandlerSave)
	mux.HandleFunc(dm.ExternalMessage_PathDelete, ExternalMessage_HandlerDelete)
	logs.Publish("Application", dm.ExternalMessage_Title)
    core.Catalog_Add(dm.ExternalMessage_Title, dm.ExternalMessage_Path, "", dm.ExternalMessage_QueryString, "Application")
}


//ExternalMessage_HandlerList is the handler for the list page
func ExternalMessage_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.ExternalMessage
	noItems, returnList, _ := dao.ExternalMessage_GetList()

	pageDetail := ExternalMessage_PageList{
		Title:            CardTitle(dm.ExternalMessage_Title, core.Action_List),
		PageTitle:        PageTitle(dm.ExternalMessage_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.ExternalMessage_TemplateList, w, r, pageDetail)

}


//ExternalMessage_HandlerView is the handler used to View a page
func ExternalMessage_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.ExternalMessage_QueryString)
	_, rD, _ := dao.ExternalMessage_GetByID(searchID)

	pageDetail := ExternalMessage_Page{
		Title:       CardTitle(dm.ExternalMessage_Title, core.Action_View),
		PageTitle:   PageTitle(dm.ExternalMessage_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = externalmessage_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.ExternalMessage_TemplateView, w, r, pageDetail)

}


//ExternalMessage_HandlerEdit is the handler used generate the Edit page
func ExternalMessage_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.ExternalMessage_QueryString)
	_, rD, _ := dao.ExternalMessage_GetByID(searchID)
	
	pageDetail := ExternalMessage_Page{
		Title:       CardTitle(dm.ExternalMessage_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.ExternalMessage_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = externalmessage_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.ExternalMessage_TemplateEdit, w, r, pageDetail)
}


//ExternalMessage_HandlerSave is the handler used process the saving of an ExternalMessage
func ExternalMessage_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("MessageID"))

	var item dm.ExternalMessage
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.ExternalMessage_SYSId_scrn)
		item.MessageID = r.FormValue(dm.ExternalMessage_MessageID_scrn)
		item.MessageFormat = r.FormValue(dm.ExternalMessage_MessageFormat_scrn)
		item.MessageDeliveredTo = r.FormValue(dm.ExternalMessage_MessageDeliveredTo_scrn)
		item.MessageBody = r.FormValue(dm.ExternalMessage_MessageBody_scrn)
		item.MessageFilename = r.FormValue(dm.ExternalMessage_MessageFilename_scrn)
		item.MessageLife = r.FormValue(dm.ExternalMessage_MessageLife_scrn)
		item.MessageDate = r.FormValue(dm.ExternalMessage_MessageDate_scrn)
		item.MessageTime = r.FormValue(dm.ExternalMessage_MessageTime_scrn)
		item.MessageTimeoutAction = r.FormValue(dm.ExternalMessage_MessageTimeoutAction_scrn)
		item.MessageACKNAK = r.FormValue(dm.ExternalMessage_MessageACKNAK_scrn)
		item.ResponseID = r.FormValue(dm.ExternalMessage_ResponseID_scrn)
		item.ResponseFilename = r.FormValue(dm.ExternalMessage_ResponseFilename_scrn)
		item.ResponseBody = r.FormValue(dm.ExternalMessage_ResponseBody_scrn)
		item.ResponseDate = r.FormValue(dm.ExternalMessage_ResponseDate_scrn)
		item.ResponseTime = r.FormValue(dm.ExternalMessage_ResponseTime_scrn)
		item.ResponseAdditionalInfo = r.FormValue(dm.ExternalMessage_ResponseAdditionalInfo_scrn)
		item.SYSCreated = r.FormValue(dm.ExternalMessage_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.ExternalMessage_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.ExternalMessage_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.ExternalMessage_SYSUpdated_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.ExternalMessage_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.ExternalMessage_SYSUpdatedHost_scrn)
		item.MessageTimeout = r.FormValue(dm.ExternalMessage_MessageTimeout_scrn)
		item.MessageEmitted = r.FormValue(dm.ExternalMessage_MessageEmitted_scrn)
		item.ResponseRecieved = r.FormValue(dm.ExternalMessage_ResponseRecieved_scrn)
		item.MessageClass = r.FormValue(dm.ExternalMessage_MessageClass_scrn)
		item.AppID = r.FormValue(dm.ExternalMessage_AppID_scrn)
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.ExternalMessage_Store(item,r)	
	http.Redirect(w, r, ExternalMessage_Redirect, http.StatusFound)
}



//ExternalMessage_HandlerDelete is the handler used process the deletion of an ExternalMessage
func ExternalMessage_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.ExternalMessage_QueryString)

	dao.ExternalMessage_Delete(searchID)	

	http.Redirect(w, r, ExternalMessage_Redirect, http.StatusFound)
}


// Builds/Popuplates the ExternalMessage Page 
func externalmessage_PopulatePage(rD dm.ExternalMessage, pageDetail ExternalMessage_Page) ExternalMessage_Page {
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.MessageID = rD.MessageID
	pageDetail.MessageFormat = rD.MessageFormat
	pageDetail.MessageDeliveredTo = rD.MessageDeliveredTo
	pageDetail.MessageBody = rD.MessageBody
	pageDetail.MessageFilename = rD.MessageFilename
	pageDetail.MessageLife = rD.MessageLife
	pageDetail.MessageDate = rD.MessageDate
	pageDetail.MessageTime = rD.MessageTime
	pageDetail.MessageTimeoutAction = rD.MessageTimeoutAction
	pageDetail.MessageACKNAK = rD.MessageACKNAK
	pageDetail.ResponseID = rD.ResponseID
	pageDetail.ResponseFilename = rD.ResponseFilename
	pageDetail.ResponseBody = rD.ResponseBody
	pageDetail.ResponseDate = rD.ResponseDate
	pageDetail.ResponseTime = rD.ResponseTime
	pageDetail.ResponseAdditionalInfo = rD.ResponseAdditionalInfo
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.MessageTimeout = rD.MessageTimeout
	pageDetail.MessageEmitted = rD.MessageEmitted
	pageDetail.ResponseRecieved = rD.ResponseRecieved
	pageDetail.MessageClass = rD.MessageClass
	pageDetail.AppID = rD.AppID
	
	
	//
	// Automatically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.MessageACKNAK_lookup = dao.StubLists_Get("messageACKNAK")
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	