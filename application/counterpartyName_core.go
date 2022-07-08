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

	pageDetail := dm.CounterpartyName_PageList{
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

	pageDetail := dm.CounterpartyName_Page{
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
func counterpartyname_PopulatePage(rD dm.CounterpartyName, pageDetail dm.CounterpartyName_Page) dm.CounterpartyName_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.NameFirm = rD.NameFirm
	pageDetail.NameCentre = rD.NameCentre
	pageDetail.FullName = rD.FullName
	pageDetail.CompID = rD.CompID
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	pageDetail.NameFirm_props = rD.NameFirm_props
	pageDetail.NameCentre_props = rD.NameCentre_props
	pageDetail.FullName_props = rD.FullName_props
	pageDetail.CompID_props = rD.CompID_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	