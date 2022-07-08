package application
// ----------------------------------------------------------------
// Automatically generated  "/application/counterpartynotes.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyNotes (counterpartynotes)
// Endpoint 	        : CounterpartyNotes (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:48
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





//CounterpartyNotes_Publish annouces the endpoints available for this object
func CounterpartyNotes_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.CounterpartyNotes_PathList, CounterpartyNotes_HandlerList)
	mux.HandleFunc(dm.CounterpartyNotes_PathView, CounterpartyNotes_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.CounterpartyNotes_Title)
    //No API
}


//CounterpartyNotes_HandlerList is the handler for the list page
func CounterpartyNotes_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.CounterpartyNotes
	noItems, returnList, _ := dao.CounterpartyNotes_GetList()

	pageDetail := dm.CounterpartyNotes_PageList{
		Title:            CardTitle(dm.CounterpartyNotes_Title, core.Action_List),
		PageTitle:        PageTitle(dm.CounterpartyNotes_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.CounterpartyNotes_TemplateList, w, r, pageDetail)

}


//CounterpartyNotes_HandlerView is the handler used to View a page
func CounterpartyNotes_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyNotes_QueryString)
	_, rD, _ := dao.CounterpartyNotes_GetByID(searchID)

	pageDetail := dm.CounterpartyNotes_Page{
		Title:       CardTitle(dm.CounterpartyNotes_Title, core.Action_View),
		PageTitle:   PageTitle(dm.CounterpartyNotes_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterpartynotes_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CounterpartyNotes_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the CounterpartyNotes Page 
func counterpartynotes_PopulatePage(rD dm.CounterpartyNotes, pageDetail dm.CounterpartyNotes_Page) dm.CounterpartyNotes_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.NoteId = rD.NoteId
	pageDetail.StreamId = rD.StreamId
	pageDetail.Summary = rD.Summary
	pageDetail.Details = rD.Details
	pageDetail.RecordState = rD.RecordState
	pageDetail.CreatedBy = rD.CreatedBy
	pageDetail.CreatedDateTime = rD.CreatedDateTime
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	pageDetail.StreamId_lookup = dao.ContactStream_GetLookup()
	
	
	
	
	
	
	
	
	pageDetail.RecordState_lookup = dao.StubLists_Get("cmrecordstate")
	
	
	
	
	
	
	pageDetail.NoteId_props = rD.NoteId_props
	pageDetail.StreamId_props = rD.StreamId_props
	pageDetail.Summary_props = rD.Summary_props
	pageDetail.Details_props = rD.Details_props
	pageDetail.RecordState_props = rD.RecordState_props
	pageDetail.CreatedBy_props = rD.CreatedBy_props
	pageDetail.CreatedDateTime_props = rD.CreatedDateTime_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	