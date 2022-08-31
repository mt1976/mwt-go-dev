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

	pageDetail := dm.Country_PageList{
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

	pageDetail := dm.Country_Page{
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
	
	pageDetail := dm.Country_Page{
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
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.Code = r.FormValue(dm.Country_Code_scrn)
		item.Name = r.FormValue(dm.Country_Name_scrn)
		item.ShortCode = r.FormValue(dm.Country_ShortCode_scrn)
		item.EU_EEA = r.FormValue(dm.Country_EU_EEA_scrn)
		item.HolidaysWeekend = r.FormValue(dm.Country_HolidaysWeekend_scrn)
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Country_Store(item,r)	
	http.Redirect(w, r, dm.Country_Redirect, http.StatusFound)
}


//Country_HandlerNew is the handler used process the creation of an Country
func Country_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.Country_New()

	pageDetail := dm.Country_Page{
		Title:       CardTitle(dm.Country_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Country_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = country_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Country_TemplateNew, w, r, pageDetail)

}	



// Builds/Popuplates the Country Page 
func country_PopulatePage(rD dm.Country, pageDetail dm.Country_Page) dm.Country_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.Code = rD.Code
	pageDetail.Name = rD.Name
	pageDetail.ShortCode = rD.ShortCode
	pageDetail.EU_EEA = rD.EU_EEA
	pageDetail.HolidaysWeekend = rD.HolidaysWeekend
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Code_props = rD.Code_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.ShortCode_props = rD.ShortCode_props
	pageDetail.EU_EEA_props = rD.EU_EEA_props
	pageDetail.HolidaysWeekend_props = rD.HolidaysWeekend_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	