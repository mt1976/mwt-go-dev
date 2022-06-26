package application
// ----------------------------------------------------------------
// Automatically generated  "/application/dealauditeventconversations.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DealAuditEventConversations (dealauditeventconversations)
// Endpoint 	        : DealAuditEventConversations (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:27
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//dealauditeventconversations_PageList provides the information for the template for a list of DealAuditEventConversationss
type DealAuditEventConversations_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.DealAuditEventConversations
}
//DealAuditEventConversations_Redirect provides a page to return to aftern an action
const (
	
	DealAuditEventConversations_Redirect = dm.DealAuditEventConversations_PathList
	
)

//dealauditeventconversations_Page provides the information for the template for an individual DealAuditEventConversations
type DealAuditEventConversations_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	DealRefNo         string
	DealRefNo_props     dm.FieldProperties
	EventIndex         string
	EventIndex_props     dm.FieldProperties
	MessageIndex         string
	MessageIndex_props     dm.FieldProperties
	Message         string
	Message_props     dm.FieldProperties
	InternalId         string
	InternalId_props     dm.FieldProperties
	InternalDeleted         string
	InternalDeleted_props     dm.FieldProperties
	UpdatedTransactionId         string
	UpdatedTransactionId_props     dm.FieldProperties
	UpdatedUserId         string
	UpdatedUserId_props     dm.FieldProperties
	UpdatedDateTime         string
	UpdatedDateTime_props     dm.FieldProperties
	DeletedTransactionId         string
	DeletedTransactionId_props     dm.FieldProperties
	DeletedUserId         string
	DeletedUserId_props     dm.FieldProperties
	ChangeType         string
	ChangeType_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//DealAuditEventConversations_Publish annouces the endpoints available for this object
func DealAuditEventConversations_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.DealAuditEventConversations_PathList, DealAuditEventConversations_HandlerList)
	mux.HandleFunc(dm.DealAuditEventConversations_PathView, DealAuditEventConversations_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.DealAuditEventConversations_Title)
    //No API
}


//DealAuditEventConversations_HandlerList is the handler for the list page
func DealAuditEventConversations_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.DealAuditEventConversations
	noItems, returnList, _ := dao.DealAuditEventConversations_GetList()

	pageDetail := DealAuditEventConversations_PageList{
		Title:            CardTitle(dm.DealAuditEventConversations_Title, core.Action_List),
		PageTitle:        PageTitle(dm.DealAuditEventConversations_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.DealAuditEventConversations_TemplateList, w, r, pageDetail)

}


//DealAuditEventConversations_HandlerView is the handler used to View a page
func DealAuditEventConversations_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DealAuditEventConversations_QueryString)
	_, rD, _ := dao.DealAuditEventConversations_GetByID(searchID)

	pageDetail := DealAuditEventConversations_Page{
		Title:       CardTitle(dm.DealAuditEventConversations_Title, core.Action_View),
		PageTitle:   PageTitle(dm.DealAuditEventConversations_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dealauditeventconversations_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.DealAuditEventConversations_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the DealAuditEventConversations Page 
func dealauditeventconversations_PopulatePage(rD dm.DealAuditEventConversations, pageDetail DealAuditEventConversations_Page) DealAuditEventConversations_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.DealRefNo = rD.DealRefNo
	pageDetail.EventIndex = rD.EventIndex
	pageDetail.MessageIndex = rD.MessageIndex
	pageDetail.Message = rD.Message
	pageDetail.InternalId = rD.InternalId
	pageDetail.InternalDeleted = rD.InternalDeleted
	pageDetail.UpdatedTransactionId = rD.UpdatedTransactionId
	pageDetail.UpdatedUserId = rD.UpdatedUserId
	pageDetail.UpdatedDateTime = rD.UpdatedDateTime
	pageDetail.DeletedTransactionId = rD.DeletedTransactionId
	pageDetail.DeletedUserId = rD.DeletedUserId
	pageDetail.ChangeType = rD.ChangeType
	
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.DealRefNo_props = rD.DealRefNo_props
	pageDetail.EventIndex_props = rD.EventIndex_props
	pageDetail.MessageIndex_props = rD.MessageIndex_props
	pageDetail.Message_props = rD.Message_props
	pageDetail.InternalId_props = rD.InternalId_props
	pageDetail.InternalDeleted_props = rD.InternalDeleted_props
	pageDetail.UpdatedTransactionId_props = rD.UpdatedTransactionId_props
	pageDetail.UpdatedUserId_props = rD.UpdatedUserId_props
	pageDetail.UpdatedDateTime_props = rD.UpdatedDateTime_props
	pageDetail.DeletedTransactionId_props = rD.DeletedTransactionId_props
	pageDetail.DeletedUserId_props = rD.DeletedUserId_props
	pageDetail.ChangeType_props = rD.ChangeType_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	