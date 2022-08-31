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
// Date & Time		    : 28/06/2022 at 16:10:50
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





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

	pageDetail := dm.DealAuditEvent_PageList{
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

	pageDetail := dm.DealAuditEvent_Page{
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
func dealauditevent_PopulatePage(rD dm.DealAuditEvent, pageDetail dm.DealAuditEvent_Page) dm.DealAuditEvent_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.DealRefNo_props = rD.DealRefNo_props
	pageDetail.EventIndex_props = rD.EventIndex_props
	pageDetail.CommonRefNo_props = rD.CommonRefNo_props
	pageDetail.Timestamp_props = rD.Timestamp_props
	pageDetail.UTCTimestamp_props = rD.UTCTimestamp_props
	pageDetail.EventType_props = rD.EventType_props
	pageDetail.Status_props = rD.Status_props
	pageDetail.LimitOrderStatus_props = rD.LimitOrderStatus_props
	pageDetail.Usr_props = rD.Usr_props
	pageDetail.DealingInterface_props = rD.DealingInterface_props
	pageDetail.SourceIP_props = rD.SourceIP_props
	pageDetail.MessageID_props = rD.MessageID_props
	pageDetail.Details_props = rD.Details_props
	pageDetail.InternalId_props = rD.InternalId_props
	pageDetail.InternalDeleted_props = rD.InternalDeleted_props
	pageDetail.UpdatedTransactionId_props = rD.UpdatedTransactionId_props
	pageDetail.UpdatedUserId_props = rD.UpdatedUserId_props
	pageDetail.UpdatedDateTime_props = rD.UpdatedDateTime_props
	pageDetail.DeletedTransactionId_props = rD.DeletedTransactionId_props
	pageDetail.DeletedUserId_props = rD.DeletedUserId_props
	pageDetail.ChangeType_props = rD.ChangeType_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	