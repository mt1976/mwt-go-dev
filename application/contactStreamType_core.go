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
// Date & Time		    : 28/06/2022 at 16:10:47
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





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

	pageDetail := dm.ContactStreamType_PageList{
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

	pageDetail := dm.ContactStreamType_Page{
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
func contactstreamtype_PopulatePage(rD dm.ContactStreamType, pageDetail dm.ContactStreamType_Page) dm.ContactStreamType_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.TypeId = rD.TypeId
	pageDetail.Description = rD.Description
	pageDetail.RecordState = rD.RecordState
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	pageDetail.RecordState_lookup = dao.StubLists_Get("cmrecordstate")
	
	
	pageDetail.TypeId_props = rD.TypeId_props
	pageDetail.Description_props = rD.Description_props
	pageDetail.RecordState_props = rD.RecordState_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	