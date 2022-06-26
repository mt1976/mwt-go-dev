package application
// ----------------------------------------------------------------
// Automatically generated  "/application/contactdeals.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ContactDeals (contactdeals)
// Endpoint 	        : ContactDeals (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 24/06/2022 at 16:01:55
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//contactdeals_PageList provides the information for the template for a list of ContactDealss
type ContactDeals_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.ContactDeals
}
//ContactDeals_Redirect provides a page to return to aftern an action
const (
	
	ContactDeals_Redirect = dm.ContactDeals_PathList
	
)

//contactdeals_Page provides the information for the template for an individual ContactDeals
type ContactDeals_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 24/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	RefNo         string
	NoteId         string
	RecordState         string
	// 
	// Dynamically generated 24/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//ContactDeals_Publish annouces the endpoints available for this object
func ContactDeals_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.ContactDeals_PathList, ContactDeals_HandlerList)
	mux.HandleFunc(dm.ContactDeals_PathView, ContactDeals_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.ContactDeals_Title)
    //No API
}


//ContactDeals_HandlerList is the handler for the list page
func ContactDeals_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.ContactDeals
	noItems, returnList, _ := dao.ContactDeals_GetList()

	pageDetail := ContactDeals_PageList{
		Title:            CardTitle(dm.ContactDeals_Title, core.Action_List),
		PageTitle:        PageTitle(dm.ContactDeals_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.ContactDeals_TemplateList, w, r, pageDetail)

}


//ContactDeals_HandlerView is the handler used to View a page
func ContactDeals_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.ContactDeals_QueryString)
	_, rD, _ := dao.ContactDeals_GetByID(searchID)

	pageDetail := ContactDeals_Page{
		Title:       CardTitle(dm.ContactDeals_Title, core.Action_View),
		PageTitle:   PageTitle(dm.ContactDeals_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = contactdeals_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.ContactDeals_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the ContactDeals Page 
func contactdeals_PopulatePage(rD dm.ContactDeals, pageDetail ContactDeals_Page) ContactDeals_Page {
	// START
	// Dynamically generated 24/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.RefNo = rD.RefNo
	pageDetail.NoteId = rD.NoteId
	pageDetail.RecordState = rD.RecordState
	
	
	//
	// Automatically generated 24/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	// 
	// Dynamically generated 24/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	