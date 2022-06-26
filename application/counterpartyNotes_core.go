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
// Date & Time		    : 26/06/2022 at 18:48:24
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//counterpartynotes_PageList provides the information for the template for a list of CounterpartyNotess
type CounterpartyNotes_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.CounterpartyNotes
}
//CounterpartyNotes_Redirect provides a page to return to aftern an action
const (
	
	CounterpartyNotes_Redirect = dm.CounterpartyNotes_PathList
	
)

//counterpartynotes_Page provides the information for the template for an individual CounterpartyNotes
type CounterpartyNotes_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	NoteId         string
	NoteId_props     dm.FieldProperties
	StreamId         string
	StreamId_lookup    []dm.Lookup_Item
	StreamId_props     dm.FieldProperties
	Summary         string
	Summary_props     dm.FieldProperties
	Details         string
	Details_props     dm.FieldProperties
	RecordState         string
	RecordState_lookup    []dm.Lookup_Item
	RecordState_props     dm.FieldProperties
	CreatedBy         string
	CreatedBy_props     dm.FieldProperties
	CreatedDateTime         string
	CreatedDateTime_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



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

	pageDetail := CounterpartyNotes_PageList{
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

	pageDetail := CounterpartyNotes_Page{
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
func counterpartynotes_PopulatePage(rD dm.CounterpartyNotes, pageDetail CounterpartyNotes_Page) CounterpartyNotes_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.NoteId = rD.NoteId
	pageDetail.StreamId = rD.StreamId
	pageDetail.Summary = rD.Summary
	pageDetail.Details = rD.Details
	pageDetail.RecordState = rD.RecordState
	pageDetail.CreatedBy = rD.CreatedBy
	pageDetail.CreatedDateTime = rD.CreatedDateTime
	
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
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
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	