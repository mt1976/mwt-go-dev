package application
// ----------------------------------------------------------------
// Automatically generated  "/application/dealauditevent.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DealAuditEvent (dealauditevent)
// Endpoint 	        : DealAuditEvent (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:09
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//dealauditevent_PageList provides the information for the template for a list of DealAuditEvents
type DealAuditEvent_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.DealAuditEvent
}
//DealAuditEvent_Redirect provides a page to return to aftern an action
const (
	DealAuditEvent_Redirect = dm.DealAuditEvent_PathList
)

//dealauditevent_Page provides the information for the template for an individual DealAuditEvent
type DealAuditEvent_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	DealRefNo         string
	EventIndex         string
	CommonRefNo         string
	Timestamp         string
	UTCTimestamp         string
	EventType         string
	Status         string
	LimitOrderStatus         string
	Usr         string
	DealingInterface         string
	SourceIP         string
	MessageID         string
	Details         string
	InternalId         string
	InternalDeleted         string
	UpdatedTransactionId         string
	UpdatedUserId         string
	UpdatedDateTime         string
	DeletedTransactionId         string
	DeletedUserId         string
	ChangeType         string
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//DealAuditEvent_Publish annouces the endpoints available for this object
func DealAuditEvent_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.DealAuditEvent_PathList, DealAuditEvent_HandlerList)
	mux.HandleFunc(dm.DealAuditEvent_PathView, DealAuditEvent_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.DealAuditEvent_Title)
    //No API
}


//DealAuditEvent_HandlerList is the handler for the list page
func DealAuditEvent_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.DealAuditEvent
	noItems, returnList, _ := dao.DealAuditEvent_GetList()

	pageDetail := DealAuditEvent_PageList{
		Title:            CardTitle(dm.DealAuditEvent_Title, core.Action_List),
		PageTitle:        PageTitle(dm.DealAuditEvent_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.DealAuditEvent_TemplateList, w, r, pageDetail)

}


//DealAuditEvent_HandlerView is the handler used to View a page
func DealAuditEvent_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DealAuditEvent_QueryString)
	_, rD, _ := dao.DealAuditEvent_GetByID(searchID)

	pageDetail := DealAuditEvent_Page{
		Title:       CardTitle(dm.DealAuditEvent_Title, core.Action_View),
		PageTitle:   PageTitle(dm.DealAuditEvent_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dealauditevent_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.DealAuditEvent_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the DealAuditEvent Page 
func dealauditevent_PopulatePage(rD dm.DealAuditEvent, pageDetail DealAuditEvent_Page) DealAuditEvent_Page {
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.DealRefNo = rD.DealRefNo
	pageDetail.EventIndex = rD.EventIndex
	pageDetail.CommonRefNo = rD.CommonRefNo
	pageDetail.Timestamp = rD.Timestamp
	pageDetail.UTCTimestamp = rD.UTCTimestamp
	pageDetail.EventType = rD.EventType
	pageDetail.Status = rD.Status
	pageDetail.LimitOrderStatus = rD.LimitOrderStatus
	pageDetail.Usr = rD.Usr
	pageDetail.DealingInterface = rD.DealingInterface
	pageDetail.SourceIP = rD.SourceIP
	pageDetail.MessageID = rD.MessageID
	pageDetail.Details = rD.Details
	pageDetail.InternalId = rD.InternalId
	pageDetail.InternalDeleted = rD.InternalDeleted
	pageDetail.UpdatedTransactionId = rD.UpdatedTransactionId
	pageDetail.UpdatedUserId = rD.UpdatedUserId
	pageDetail.UpdatedDateTime = rD.UpdatedDateTime
	pageDetail.DeletedTransactionId = rD.DeletedTransactionId
	pageDetail.DeletedUserId = rD.DeletedUserId
	pageDetail.ChangeType = rD.ChangeType
	
	
	//
	// Automatically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	