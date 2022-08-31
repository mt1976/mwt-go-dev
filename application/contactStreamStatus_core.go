package application
// ----------------------------------------------------------------
// Automatically generated  "/application/contactstreamstatus.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ContactStreamStatus (contactstreamstatus)
// Endpoint 	        : ContactStreamStatus (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:46
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





//ContactStreamStatus_Publish annouces the endpoints available for this object
func ContactStreamStatus_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.ContactStreamStatus_PathList, ContactStreamStatus_HandlerList)
	mux.HandleFunc(dm.ContactStreamStatus_PathView, ContactStreamStatus_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.ContactStreamStatus_Title)
    //No API
}


//ContactStreamStatus_HandlerList is the handler for the list page
func ContactStreamStatus_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.ContactStreamStatus
	noItems, returnList, _ := dao.ContactStreamStatus_GetList()

	pageDetail := dm.ContactStreamStatus_PageList{
		Title:            CardTitle(dm.ContactStreamStatus_Title, core.Action_List),
		PageTitle:        PageTitle(dm.ContactStreamStatus_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.ContactStreamStatus_TemplateList, w, r, pageDetail)

}


//ContactStreamStatus_HandlerView is the handler used to View a page
func ContactStreamStatus_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.ContactStreamStatus_QueryString)
	_, rD, _ := dao.ContactStreamStatus_GetByID(searchID)

	pageDetail := dm.ContactStreamStatus_Page{
		Title:       CardTitle(dm.ContactStreamStatus_Title, core.Action_View),
		PageTitle:   PageTitle(dm.ContactStreamStatus_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = contactstreamstatus_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.ContactStreamStatus_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the ContactStreamStatus Page 
func contactstreamstatus_PopulatePage(rD dm.ContactStreamStatus, pageDetail dm.ContactStreamStatus_Page) dm.ContactStreamStatus_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.StatusId = rD.StatusId
	pageDetail.Description = rD.Description
	pageDetail.RecordState = rD.RecordState
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	pageDetail.StatusId_props = rD.StatusId_props
	pageDetail.Description_props = rD.Description_props
	pageDetail.RecordState_props = rD.RecordState_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	