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
// Date & Time		    : 26/06/2022 at 18:48:20
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//contactstream_PageList provides the information for the template for a list of ContactStreams
type ContactStream_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.ContactStream
}
//ContactStream_Redirect provides a page to return to aftern an action
const (
	
	ContactStream_Redirect = dm.ContactStream_PathList
	
)

//contactstream_Page provides the information for the template for an individual ContactStream
type ContactStream_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	StreamId         string
	StreamId_props     dm.FieldProperties
	ContactId         string
	ContactId_props     dm.FieldProperties
	Usr         string
	Usr_props     dm.FieldProperties
	Description         string
	Description_props     dm.FieldProperties
	StreamTypeId         string
	StreamTypeId_lookup    []dm.Lookup_Item
	StreamTypeId_props     dm.FieldProperties
	StreamStatusId         string
	StreamStatusId_lookup    []dm.Lookup_Item
	StreamStatusId_props     dm.FieldProperties
	RecordState         string
	RecordState_lookup    []dm.Lookup_Item
	RecordState_props     dm.FieldProperties
	CreatedDateTime         string
	CreatedDateTime_props     dm.FieldProperties
	OpenedDateTime         string
	OpenedDateTime_props     dm.FieldProperties
	ClosedDateTime         string
	ClosedDateTime_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



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

	pageDetail := ContactStream_PageList{
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

	pageDetail := ContactStream_Page{
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
func contactstream_PopulatePage(rD dm.ContactStream, pageDetail ContactStream_Page) ContactStream_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
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
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	