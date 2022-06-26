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
// Date & Time		    : 26/06/2022 at 18:48:21
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//contactstreamstatus_PageList provides the information for the template for a list of ContactStreamStatuss
type ContactStreamStatus_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.ContactStreamStatus
}
//ContactStreamStatus_Redirect provides a page to return to aftern an action
const (
	
	ContactStreamStatus_Redirect = dm.ContactStreamStatus_PathList
	
)

//contactstreamstatus_Page provides the information for the template for an individual ContactStreamStatus
type ContactStreamStatus_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	StatusId         string
	StatusId_props     dm.FieldProperties
	Description         string
	Description_props     dm.FieldProperties
	RecordState         string
	RecordState_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



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

	pageDetail := ContactStreamStatus_PageList{
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

	pageDetail := ContactStreamStatus_Page{
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
func contactstreamstatus_PopulatePage(rD dm.ContactStreamStatus, pageDetail ContactStreamStatus_Page) ContactStreamStatus_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.StatusId = rD.StatusId
	pageDetail.Description = rD.Description
	pageDetail.RecordState = rD.RecordState
	
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	pageDetail.StatusId_props = rD.StatusId_props
	pageDetail.Description_props = rD.Description_props
	pageDetail.RecordState_props = rD.RecordState_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	