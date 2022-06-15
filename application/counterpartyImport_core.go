package application
// ----------------------------------------------------------------
// Automatically generated  "/application/counterpartyimport.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyImport (counterpartyimport)
// Endpoint 	        : CounterpartyImport (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:31:57
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//counterpartyimport_PageList provides the information for the template for a list of CounterpartyImports
type CounterpartyImport_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.CounterpartyImport
}
//CounterpartyImport_Redirect provides a page to return to aftern an action
const (
	CounterpartyImport_Redirect = dm.CounterpartyImport_PathList
)

//counterpartyimport_Page provides the information for the template for an individual CounterpartyImport
type CounterpartyImport_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	KeyImportID         string
	Firm         string
	Centre         string
	FirmName         string
	CentreName         string
	KeyOriginID         string
	FullName         string
	CompID         string
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//CounterpartyImport_Publish annouces the endpoints available for this object
func CounterpartyImport_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.CounterpartyImport_PathList, CounterpartyImport_HandlerList)
	mux.HandleFunc(dm.CounterpartyImport_PathView, CounterpartyImport_HandlerView)
	mux.HandleFunc(dm.CounterpartyImport_PathEdit, CounterpartyImport_HandlerEdit)
	mux.HandleFunc(dm.CounterpartyImport_PathNew, CounterpartyImport_HandlerNew)
	mux.HandleFunc(dm.CounterpartyImport_PathSave, CounterpartyImport_HandlerSave)
	mux.HandleFunc(dm.CounterpartyImport_PathDelete, CounterpartyImport_HandlerDelete)
	logs.Publish("Application", dm.CounterpartyImport_Title)
    //No API
}


//CounterpartyImport_HandlerList is the handler for the list page
func CounterpartyImport_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.CounterpartyImport
	noItems, returnList, _ := dao.CounterpartyImport_GetList()

	pageDetail := CounterpartyImport_PageList{
		Title:            CardTitle(dm.CounterpartyImport_Title, core.Action_List),
		PageTitle:        PageTitle(dm.CounterpartyImport_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.CounterpartyImport_TemplateList, w, r, pageDetail)

}


//CounterpartyImport_HandlerView is the handler used to View a page
func CounterpartyImport_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyImport_QueryString)
	_, rD, _ := dao.CounterpartyImport_GetByID(searchID)

	pageDetail := CounterpartyImport_Page{
		Title:       CardTitle(dm.CounterpartyImport_Title, core.Action_View),
		PageTitle:   PageTitle(dm.CounterpartyImport_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterpartyimport_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CounterpartyImport_TemplateView, w, r, pageDetail)

}


//CounterpartyImport_HandlerEdit is the handler used generate the Edit page
func CounterpartyImport_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyImport_QueryString)
	_, rD, _ := dao.CounterpartyImport_GetByID(searchID)
	
	pageDetail := CounterpartyImport_Page{
		Title:       CardTitle(dm.CounterpartyImport_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.CounterpartyImport_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterpartyimport_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CounterpartyImport_TemplateEdit, w, r, pageDetail)
}


//CounterpartyImport_HandlerSave is the handler used process the saving of an CounterpartyImport
func CounterpartyImport_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("CompID"))

	var item dm.CounterpartyImport
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.KeyImportID = r.FormValue(dm.CounterpartyImport_KeyImportID)
		item.Firm = r.FormValue(dm.CounterpartyImport_Firm)
		item.Centre = r.FormValue(dm.CounterpartyImport_Centre)
		item.FirmName = r.FormValue(dm.CounterpartyImport_FirmName)
		item.CentreName = r.FormValue(dm.CounterpartyImport_CentreName)
		item.KeyOriginID = r.FormValue(dm.CounterpartyImport_KeyOriginID)
		item.FullName = r.FormValue(dm.CounterpartyImport_FullName)
		item.CompID = r.FormValue(dm.CounterpartyImport_CompID)
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.CounterpartyImport_Store(item,r)	
	http.Redirect(w, r, CounterpartyImport_Redirect, http.StatusFound)
}


//CounterpartyImport_HandlerNew is the handler used process the creation of an CounterpartyImport
func CounterpartyImport_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := CounterpartyImport_Page{
		Title:       CardTitle(dm.CounterpartyImport_Title, core.Action_New),
		PageTitle:   PageTitle(dm.CounterpartyImport_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterpartyimport_PopulatePage(dm.CounterpartyImport{} , pageDetail) 

	ExecuteTemplate(dm.CounterpartyImport_TemplateNew, w, r, pageDetail)

}	


//CounterpartyImport_HandlerDelete is the handler used process the deletion of an CounterpartyImport
func CounterpartyImport_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.CounterpartyImport_QueryString)

	dao.CounterpartyImport_Delete(searchID)	

	http.Redirect(w, r, CounterpartyImport_Redirect, http.StatusFound)
}


// Builds/Popuplates the CounterpartyImport Page 
func counterpartyimport_PopulatePage(rD dm.CounterpartyImport, pageDetail CounterpartyImport_Page) CounterpartyImport_Page {
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.KeyImportID = rD.KeyImportID
	pageDetail.Firm = rD.Firm
	pageDetail.Centre = rD.Centre
	pageDetail.FirmName = rD.FirmName
	pageDetail.CentreName = rD.CentreName
	pageDetail.KeyOriginID = rD.KeyOriginID
	pageDetail.FullName = rD.FullName
	pageDetail.CompID = rD.CompID
	
	
	//
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	