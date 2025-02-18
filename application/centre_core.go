package application
// ----------------------------------------------------------------
// Automatically generated  "/application/centre.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Centre (centre)
// Endpoint 	        : Centre (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:44
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





//Centre_Publish annouces the endpoints available for this object
func Centre_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Centre_Path, Centre_Handler)
	mux.HandleFunc(dm.Centre_PathList, Centre_HandlerList)
	mux.HandleFunc(dm.Centre_PathView, Centre_HandlerView)
	mux.HandleFunc(dm.Centre_PathEdit, Centre_HandlerEdit)
	mux.HandleFunc(dm.Centre_PathNew, Centre_HandlerNew)
	mux.HandleFunc(dm.Centre_PathSave, Centre_HandlerSave)
	mux.HandleFunc(dm.Centre_PathDelete, Centre_HandlerDelete)
	logs.Publish("Application", dm.Centre_Title)
    core.Catalog_Add(dm.Centre_Title, dm.Centre_Path, "", dm.Centre_QueryString, "Application")
}


//Centre_HandlerList is the handler for the list page
func Centre_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Centre
	noItems, returnList, _ := dao.Centre_GetList()

	pageDetail := dm.Centre_PageList{
		Title:            CardTitle(dm.Centre_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Centre_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Centre_TemplateList, w, r, pageDetail)

}


//Centre_HandlerView is the handler used to View a page
func Centre_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Centre_QueryString)
	_, rD, _ := dao.Centre_GetByID(searchID)

	pageDetail := dm.Centre_Page{
		Title:       CardTitle(dm.Centre_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Centre_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = centre_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Centre_TemplateView, w, r, pageDetail)

}


//Centre_HandlerEdit is the handler used generate the Edit page
func Centre_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Centre_QueryString)
	_, rD, _ := dao.Centre_GetByID(searchID)
	
	pageDetail := dm.Centre_Page{
		Title:       CardTitle(dm.Centre_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Centre_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = centre_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Centre_TemplateEdit, w, r, pageDetail)
}


//Centre_HandlerSave is the handler used process the saving of an Centre
func Centre_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.Centre
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.Code = r.FormValue(dm.Centre_Code_scrn)
		item.Name = r.FormValue(dm.Centre_Name_scrn)
		item.Country = r.FormValue(dm.Centre_Country_scrn)
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Centre_Store(item,r)	
	http.Redirect(w, r, dm.Centre_Redirect, http.StatusFound)
}


//Centre_HandlerNew is the handler used process the creation of an Centre
func Centre_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.Centre_New()

	pageDetail := dm.Centre_Page{
		Title:       CardTitle(dm.Centre_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Centre_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = centre_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Centre_TemplateNew, w, r, pageDetail)

}	


//Centre_HandlerDelete is the handler used process the deletion of an Centre
func Centre_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Centre_QueryString)

	dao.Centre_Delete(searchID)	

	http.Redirect(w, r, dm.Centre_Redirect, http.StatusFound)
}


// Builds/Popuplates the Centre Page 
func centre_PopulatePage(rD dm.Centre, pageDetail dm.Centre_Page) dm.Centre_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.Code = rD.Code
	pageDetail.Name = rD.Name
	pageDetail.Country = rD.Country
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	pageDetail.Country_lookup = dao.Country_GetLookup()
	
	
	
	pageDetail.Code_props = rD.Code_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.Country_props = rD.Country_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	