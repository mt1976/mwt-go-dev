package application
// ----------------------------------------------------------------
// Automatically generated  "/application/contactstream.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ContactStream (contactstream)
// Endpoint 	        : ContactStream (ID)
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





//ContactStream_Publish annouces the endpoints available for this object
func ContactStream_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.ContactStream_PathList, ContactStream_HandlerList)
	mux.HandleFunc(dm.ContactStream_PathView, ContactStream_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.ContactStream_Title)
    //No API
}


//ContactStream_HandlerList is the handler for the list page
func ContactStream_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.ContactStream
	noItems, returnList, _ := dao.ContactStream_GetList()

	pageDetail := dm.ContactStream_PageList{
		Title:            CardTitle(dm.ContactStream_Title, core.Action_List),
		PageTitle:        PageTitle(dm.ContactStream_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.ContactStream_TemplateList, w, r, pageDetail)

}


//ContactStream_HandlerView is the handler used to View a page
func ContactStream_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.ContactStream_QueryString)
	_, rD, _ := dao.ContactStream_GetByID(searchID)

	pageDetail := dm.ContactStream_Page{
		Title:       CardTitle(dm.ContactStream_Title, core.Action_View),
		PageTitle:   PageTitle(dm.ContactStream_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = contactstream_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.ContactStream_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the ContactStream Page 
func contactstream_PopulatePage(rD dm.ContactStream, pageDetail dm.ContactStream_Page) dm.ContactStream_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.StreamId = rD.StreamId
	pageDetail.ContactId = rD.ContactId
	pageDetail.Usr = rD.Usr
	pageDetail.Description = rD.Description
	pageDetail.StreamTypeId = rD.StreamTypeId
	pageDetail.StreamStatusId = rD.StreamStatusId
	pageDetail.RecordState = rD.RecordState
	pageDetail.CreatedDateTime = rD.CreatedDateTime
	pageDetail.OpenedDateTime = rD.OpenedDateTime
	pageDetail.ClosedDateTime = rD.ClosedDateTime
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	pageDetail.StreamTypeId_lookup = dao.ContactStreamType_GetLookup()
	
	
	
	pageDetail.StreamStatusId_lookup = dao.ContactStreamStatus_GetLookup()
	
	
	
	
	pageDetail.RecordState_lookup = dao.StubLists_Get("cmrecordstatus")
	
	
	
	
	
	
	
	
	pageDetail.StreamId_props = rD.StreamId_props
	pageDetail.ContactId_props = rD.ContactId_props
	pageDetail.Usr_props = rD.Usr_props
	pageDetail.Description_props = rD.Description_props
	pageDetail.StreamTypeId_props = rD.StreamTypeId_props
	pageDetail.StreamStatusId_props = rD.StreamStatusId_props
	pageDetail.RecordState_props = rD.RecordState_props
	pageDetail.CreatedDateTime_props = rD.CreatedDateTime_props
	pageDetail.OpenedDateTime_props = rD.OpenedDateTime_props
	pageDetail.ClosedDateTime_props = rD.ClosedDateTime_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	