package application
// ----------------------------------------------------------------
// Automatically generated  "/application/firm.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Firm (firm)
// Endpoint 	        : Firm (FirmName)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:11
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//firm_PageList provides the information for the template for a list of Firms
type Firm_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Firm
}
//Firm_Redirect provides a page to return to aftern an action
const (
	Firm_Redirect = dm.Firm_PathList
)

//firm_Page provides the information for the template for an individual Firm
type Firm_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	FirmName         string
	FullName         string
	Country         string
	Country_lookup    []dm.Lookup_Item
	Sector         string
	Sector_lookup    []dm.Lookup_Item
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Firm_Publish annouces the endpoints available for this object
func Firm_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Firm_Path, Firm_Handler)
	mux.HandleFunc(dm.Firm_PathList, Firm_HandlerList)
	mux.HandleFunc(dm.Firm_PathView, Firm_HandlerView)
	mux.HandleFunc(dm.Firm_PathEdit, Firm_HandlerEdit)
	mux.HandleFunc(dm.Firm_PathNew, Firm_HandlerNew)
	mux.HandleFunc(dm.Firm_PathSave, Firm_HandlerSave)
	mux.HandleFunc(dm.Firm_PathDelete, Firm_HandlerDelete)
	logs.Publish("Application", dm.Firm_Title)
    core.Catalog_Add(dm.Firm_Title, dm.Firm_Path, "", dm.Firm_QueryString, "Application")
}


//Firm_HandlerList is the handler for the list page
func Firm_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Firm
	noItems, returnList, _ := dao.Firm_GetList()

	pageDetail := Firm_PageList{
		Title:            CardTitle(dm.Firm_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Firm_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Firm_TemplateList, w, r, pageDetail)

}


//Firm_HandlerView is the handler used to View a page
func Firm_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Firm_QueryString)
	_, rD, _ := dao.Firm_GetByID(searchID)

	pageDetail := Firm_Page{
		Title:       CardTitle(dm.Firm_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Firm_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = firm_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Firm_TemplateView, w, r, pageDetail)

}


//Firm_HandlerEdit is the handler used generate the Edit page
func Firm_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Firm_QueryString)
	_, rD, _ := dao.Firm_GetByID(searchID)
	
	pageDetail := Firm_Page{
		Title:       CardTitle(dm.Firm_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Firm_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = firm_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Firm_TemplateEdit, w, r, pageDetail)
}


//Firm_HandlerSave is the handler used process the saving of an Firm
func Firm_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("FirmName"))

	var item dm.Firm
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.FirmName = r.FormValue(dm.Firm_FirmName_scrn)
		item.FullName = r.FormValue(dm.Firm_FullName_scrn)
		item.Country = r.FormValue(dm.Firm_Country_scrn)
		item.Sector = r.FormValue(dm.Firm_Sector_scrn)
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Firm_Store(item,r)	
	http.Redirect(w, r, Firm_Redirect, http.StatusFound)
}


//Firm_HandlerNew is the handler used process the creation of an Firm
func Firm_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Firm_Page{
		Title:       CardTitle(dm.Firm_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Firm_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = firm_PopulatePage(dm.Firm{} , pageDetail) 

	ExecuteTemplate(dm.Firm_TemplateNew, w, r, pageDetail)

}	


//Firm_HandlerDelete is the handler used process the deletion of an Firm
func Firm_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Firm_QueryString)

	dao.Firm_Delete(searchID)	

	http.Redirect(w, r, Firm_Redirect, http.StatusFound)
}


// Builds/Popuplates the Firm Page 
func firm_PopulatePage(rD dm.Firm, pageDetail Firm_Page) Firm_Page {
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.FirmName = rD.FirmName
	pageDetail.FullName = rD.FullName
	pageDetail.Country = rD.Country
	pageDetail.Sector = rD.Sector
	
	
	//
	// Automatically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	pageDetail.Country_lookup = dao.Country_GetLookup()
	
	
	
	pageDetail.Sector_lookup = dao.Sector_GetLookup()
	
	
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	