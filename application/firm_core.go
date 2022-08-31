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
// Date & Time		    : 28/06/2022 at 16:10:53
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





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

	pageDetail := dm.Firm_PageList{
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

	pageDetail := dm.Firm_Page{
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
	
	pageDetail := dm.Firm_Page{
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
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.FirmName = r.FormValue(dm.Firm_FirmName_scrn)
		item.FullName = r.FormValue(dm.Firm_FullName_scrn)
		item.Country = r.FormValue(dm.Firm_Country_scrn)
		item.Sector = r.FormValue(dm.Firm_Sector_scrn)
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Firm_Store(item,r)	
	http.Redirect(w, r, dm.Firm_Redirect, http.StatusFound)
}


//Firm_HandlerNew is the handler used process the creation of an Firm
func Firm_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.Firm_New()

	pageDetail := dm.Firm_Page{
		Title:       CardTitle(dm.Firm_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Firm_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = firm_PopulatePage(rD , pageDetail) 

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

	http.Redirect(w, r, dm.Firm_Redirect, http.StatusFound)
}


// Builds/Popuplates the Firm Page 
func firm_PopulatePage(rD dm.Firm, pageDetail dm.Firm_Page) dm.Firm_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.FirmName = rD.FirmName
	pageDetail.FullName = rD.FullName
	pageDetail.Country = rD.Country
	pageDetail.Sector = rD.Sector
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	pageDetail.Country_lookup = dao.Country_GetLookup()
	
	
	
	pageDetail.Sector_lookup = dao.Sector_GetLookup()
	
	
	
	pageDetail.FirmName_props = rD.FirmName_props
	pageDetail.FullName_props = rD.FullName_props
	pageDetail.Country_props = rD.Country_props
	pageDetail.Sector_props = rD.Sector_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	