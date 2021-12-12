package application
// ----------------------------------------------------------------
// Automatically generated  "/application/sector.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Sector (sector)
// Endpoint 	        : Sector (Sector)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:19
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//sector_PageList provides the information for the template for a list of Sectors
type Sector_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Sector
}

//sector_Page provides the information for the template for an individual Sector
type Sector_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		Code string
		Name string
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	Sector_Redirect = dm.Sector_PathList
)

//Sector_Publish annouces the endpoints available for this object
func Sector_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Sector_Path, Sector_Handler)
	mux.HandleFunc(dm.Sector_PathList, Sector_HandlerList)
	mux.HandleFunc(dm.Sector_PathView, Sector_HandlerView)
	mux.HandleFunc(dm.Sector_PathEdit, Sector_HandlerEdit)
	mux.HandleFunc(dm.Sector_PathNew, Sector_HandlerNew)
	mux.HandleFunc(dm.Sector_PathSave, Sector_HandlerSave)
	mux.HandleFunc(dm.Sector_PathDelete, Sector_HandlerDelete)
	logs.Publish("Siena", dm.Sector_Title)
    core.Catalog_Add(dm.Sector_Title, dm.Sector_Path, "", dm.Sector_QueryString, "APP")
}

//Sector_HandlerList is the handler for the list page
func Sector_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Sector
	noItems, returnList, _ := dao.Sector_GetList()

	pageDetail := Sector_PageList{
		Title:            CardTitle(dm.Sector_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Sector_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Sector_TemplateList, w, r, pageDetail)

}

//Sector_HandlerView is the handler used to View a page
func Sector_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Sector_QueryString)
	_, rD, _ := dao.Sector_GetByID(searchID)

	pageDetail := Sector_Page{
		Title:       CardTitle(dm.Sector_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Sector_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Sector_TemplateView, w, r, pageDetail)

}

//Sector_HandlerEdit is the handler used generate the Edit page
func Sector_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Sector_QueryString)
	_, rD, _ := dao.Sector_GetByID(searchID)
	
	pageDetail := Sector_Page{
		Title:       CardTitle(dm.Sector_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Sector_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Sector_TemplateEdit, w, r, pageDetail)


}

//Sector_HandlerSave is the handler used process the saving of an Sector
func Sector_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.Sector
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.Code = r.FormValue(dm.Sector_Code)
		item.Name = r.FormValue(dm.Sector_Name)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.Sector_Store(item,r)	

	http.Redirect(w, r, Sector_Redirect, http.StatusFound)
}

//Sector_HandlerNew is the handler used process the creation of an Sector
func Sector_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Sector_Page{
		Title:       CardTitle(dm.Sector_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Sector_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.Code = ""
pageDetail.Name = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Sector_TemplateNew, w, r, pageDetail)

}

//Sector_HandlerDelete is the handler used process the deletion of an Sector
func Sector_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Sector_QueryString)

	dao.Sector_Delete(searchID)	

	http.Redirect(w, r, Sector_Redirect, http.StatusFound)
}