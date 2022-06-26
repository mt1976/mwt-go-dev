package application
// ----------------------------------------------------------------
// Automatically generated  "/application/contactstreamtype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ContactStreamType (contactstreamtype)
// Endpoint 	        : ContactStreamType (ID)
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

//contactstreamtype_PageList provides the information for the template for a list of ContactStreamTypes
type ContactStreamType_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.ContactStreamType
}
//ContactStreamType_Redirect provides a page to return to aftern an action
const (
	
	ContactStreamType_Redirect = dm.ContactStreamType_PathList
	
)

//contactstreamtype_Page provides the information for the template for an individual ContactStreamType
type ContactStreamType_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	TypeId         string
	TypeId_props     dm.FieldProperties
	Description         string
	Description_props     dm.FieldProperties
	RecordState         string
	RecordState_lookup    []dm.Lookup_Item
	RecordState_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//ContactStreamType_Publish annouces the endpoints available for this object
func ContactStreamType_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.ContactStreamType_PathList, ContactStreamType_HandlerList)
	mux.HandleFunc(dm.ContactStreamType_PathView, ContactStreamType_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.ContactStreamType_Title)
    //No API
}


//ContactStreamType_HandlerList is the handler for the list page
func ContactStreamType_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.ContactStreamType
	noItems, returnList, _ := dao.ContactStreamType_GetList()

	pageDetail := ContactStreamType_PageList{
		Title:            CardTitle(dm.ContactStreamType_Title, core.Action_List),
		PageTitle:        PageTitle(dm.ContactStreamType_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.ContactStreamType_TemplateList, w, r, pageDetail)

}


//ContactStreamType_HandlerView is the handler used to View a page
func ContactStreamType_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.ContactStreamType_QueryString)
	_, rD, _ := dao.ContactStreamType_GetByID(searchID)

	pageDetail := ContactStreamType_Page{
		Title:       CardTitle(dm.ContactStreamType_Title, core.Action_View),
		PageTitle:   PageTitle(dm.ContactStreamType_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = contactstreamtype_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.ContactStreamType_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the ContactStreamType Page 
func contactstreamtype_PopulatePage(rD dm.ContactStreamType, pageDetail ContactStreamType_Page) ContactStreamType_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.TypeId = rD.TypeId
	pageDetail.Description = rD.Description
	pageDetail.RecordState = rD.RecordState
	
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	pageDetail.RecordState_lookup = dao.StubLists_Get("cmrecordstate")
	
	
	pageDetail.TypeId_props = rD.TypeId_props
	pageDetail.Description_props = rD.Description_props
	pageDetail.RecordState_props = rD.RecordState_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	