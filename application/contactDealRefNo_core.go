package application
// ----------------------------------------------------------------
// Automatically generated  "/application/contactdealrefno.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ContactDealRefNo (contactdealrefno)
// Endpoint 	        : ContactDealRefNo (ID)
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

//contactdealrefno_PageList provides the information for the template for a list of ContactDealRefNos
type ContactDealRefNo_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.ContactDealRefNo
}
//ContactDealRefNo_Redirect provides a page to return to aftern an action
const (
	
	ContactDealRefNo_Redirect = dm.ContactDealRefNo_PathList
	
)

//contactdealrefno_Page provides the information for the template for an individual ContactDealRefNo
type ContactDealRefNo_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	RefNo         string
	RefNo_lookup    []dm.Lookup_Item
	RefNo_props     dm.FieldProperties
	NoteId         string
	NoteId_lookup    []dm.Lookup_Item
	NoteId_props     dm.FieldProperties
	RecordState         string
	RecordState_lookup    []dm.Lookup_Item
	RecordState_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//ContactDealRefNo_Publish annouces the endpoints available for this object
func ContactDealRefNo_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.ContactDealRefNo_PathList, ContactDealRefNo_HandlerList)
	mux.HandleFunc(dm.ContactDealRefNo_PathView, ContactDealRefNo_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.ContactDealRefNo_Title)
    //No API
}


//ContactDealRefNo_HandlerList is the handler for the list page
func ContactDealRefNo_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.ContactDealRefNo
	noItems, returnList, _ := dao.ContactDealRefNo_GetList()

	pageDetail := ContactDealRefNo_PageList{
		Title:            CardTitle(dm.ContactDealRefNo_Title, core.Action_List),
		PageTitle:        PageTitle(dm.ContactDealRefNo_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.ContactDealRefNo_TemplateList, w, r, pageDetail)

}


//ContactDealRefNo_HandlerView is the handler used to View a page
func ContactDealRefNo_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.ContactDealRefNo_QueryString)
	_, rD, _ := dao.ContactDealRefNo_GetByID(searchID)

	pageDetail := ContactDealRefNo_Page{
		Title:       CardTitle(dm.ContactDealRefNo_Title, core.Action_View),
		PageTitle:   PageTitle(dm.ContactDealRefNo_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = contactdealrefno_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.ContactDealRefNo_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the ContactDealRefNo Page 
func contactdealrefno_PopulatePage(rD dm.ContactDealRefNo, pageDetail ContactDealRefNo_Page) ContactDealRefNo_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.RefNo = rD.RefNo
	pageDetail.NoteId = rD.NoteId
	pageDetail.RecordState = rD.RecordState
	
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	pageDetail.RefNo_lookup = dao.Transaction_GetLookup()
	
	
	
	pageDetail.NoteId_lookup = dao.CounterpartyNotes_GetLookup()
	
	
	
	
	pageDetail.RecordState_lookup = dao.StubLists_Get("cmrecordstate")
	
	
	pageDetail.RefNo_props = rD.RefNo_props
	pageDetail.NoteId_props = rD.NoteId_props
	pageDetail.RecordState_props = rD.RecordState_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	