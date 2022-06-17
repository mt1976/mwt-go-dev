package application
// ----------------------------------------------------------------
// Automatically generated  "/application/counterpartyname.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyName (counterpartyname)
// Endpoint 	        : CounterpartyName (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//counterpartyname_PageList provides the information for the template for a list of CounterpartyNames
type CounterpartyName_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.CounterpartyName
}
//CounterpartyName_Redirect provides a page to return to aftern an action
const (
	CounterpartyName_Redirect = dm.CounterpartyName_PathList
)

//counterpartyname_Page provides the information for the template for an individual CounterpartyName
type CounterpartyName_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	NameFirm         string
	NameCentre         string
	FullName         string
	CompID         string
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//CounterpartyName_Publish annouces the endpoints available for this object
func CounterpartyName_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.CounterpartyName_PathList, CounterpartyName_HandlerList)
	mux.HandleFunc(dm.CounterpartyName_PathView, CounterpartyName_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.CounterpartyName_Title)
    //No API
}


//CounterpartyName_HandlerList is the handler for the list page
func CounterpartyName_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.CounterpartyName
	noItems, returnList, _ := dao.CounterpartyName_GetList()

	pageDetail := CounterpartyName_PageList{
		Title:            CardTitle(dm.CounterpartyName_Title, core.Action_List),
		PageTitle:        PageTitle(dm.CounterpartyName_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.CounterpartyName_TemplateList, w, r, pageDetail)

}


//CounterpartyName_HandlerView is the handler used to View a page
func CounterpartyName_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyName_QueryString)
	_, rD, _ := dao.CounterpartyName_GetByID(searchID)

	pageDetail := CounterpartyName_Page{
		Title:       CardTitle(dm.CounterpartyName_Title, core.Action_View),
		PageTitle:   PageTitle(dm.CounterpartyName_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterpartyname_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CounterpartyName_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the CounterpartyName Page 
func counterpartyname_PopulatePage(rD dm.CounterpartyName, pageDetail CounterpartyName_Page) CounterpartyName_Page {
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.NameFirm = rD.NameFirm
	pageDetail.NameCentre = rD.NameCentre
	pageDetail.FullName = rD.FullName
	pageDetail.CompID = rD.CompID
	
	
	//
	// Automatically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	