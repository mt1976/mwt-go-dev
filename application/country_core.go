package application
// ----------------------------------------------------------------
// Automatically generated  "/application/country.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Country (country)
// Endpoint 	        : Country (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:31:59
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//country_PageList provides the information for the template for a list of Countrys
type Country_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Country
}
//Country_Redirect provides a page to return to aftern an action
const (
	Country_Redirect = dm.Country_PathList
)

//country_Page provides the information for the template for an individual Country
type Country_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	Code         string
	Name         string
	ShortCode         string
	EU_EEA         string
	HolidaysWeekend         string
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Country_Publish annouces the endpoints available for this object
func Country_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Country_Path, Country_Handler)
	mux.HandleFunc(dm.Country_PathList, Country_HandlerList)
	mux.HandleFunc(dm.Country_PathView, Country_HandlerView)
	mux.HandleFunc(dm.Country_PathEdit, Country_HandlerEdit)
	mux.HandleFunc(dm.Country_PathNew, Country_HandlerNew)
	mux.HandleFunc(dm.Country_PathSave, Country_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Country_Title)
    core.Catalog_Add(dm.Country_Title, dm.Country_Path, "", dm.Country_QueryString, "Application")
}


//Country_HandlerList is the handler for the list page
func Country_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Country
	noItems, returnList, _ := dao.Country_GetList()

	pageDetail := Country_PageList{
		Title:            CardTitle(dm.Country_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Country_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Country_TemplateList, w, r, pageDetail)

}


//Country_HandlerView is the handler used to View a page
func Country_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Country_QueryString)
	_, rD, _ := dao.Country_GetByID(searchID)

	pageDetail := Country_Page{
		Title:       CardTitle(dm.Country_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Country_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = country_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Country_TemplateView, w, r, pageDetail)

}


//Country_HandlerEdit is the handler used generate the Edit page
func Country_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Country_QueryString)
	_, rD, _ := dao.Country_GetByID(searchID)
	
	pageDetail := Country_Page{
		Title:       CardTitle(dm.Country_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Country_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = country_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Country_TemplateEdit, w, r, pageDetail)
}


//Country_HandlerSave is the handler used process the saving of an Country
func Country_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.Country
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.Code = r.FormValue(dm.Country_Code)
		item.Name = r.FormValue(dm.Country_Name)
		item.ShortCode = r.FormValue(dm.Country_ShortCode)
		item.EU_EEA = r.FormValue(dm.Country_EU_EEA)
		item.HolidaysWeekend = r.FormValue(dm.Country_HolidaysWeekend)
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Country_Store(item,r)	
	http.Redirect(w, r, Country_Redirect, http.StatusFound)
}


//Country_HandlerNew is the handler used process the creation of an Country
func Country_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Country_Page{
		Title:       CardTitle(dm.Country_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Country_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = country_PopulatePage(dm.Country{} , pageDetail) 

	ExecuteTemplate(dm.Country_TemplateNew, w, r, pageDetail)

}	



// Builds/Popuplates the Country Page 
func country_PopulatePage(rD dm.Country, pageDetail Country_Page) Country_Page {
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.Code = rD.Code
	pageDetail.Name = rD.Name
	pageDetail.ShortCode = rD.ShortCode
	pageDetail.EU_EEA = rD.EU_EEA
	pageDetail.HolidaysWeekend = rD.HolidaysWeekend
	
	
	//
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	