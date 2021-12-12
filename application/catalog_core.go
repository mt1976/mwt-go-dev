package application
// ----------------------------------------------------------------
// Automatically generated  "/application/catalog.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Catalog (catalog)
// Endpoint 	        : Catalog (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:16:04
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//catalog_PageList provides the information for the template for a list of Catalogs
type Catalog_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Catalog
}

//catalog_Page provides the information for the template for an individual Catalog
type Catalog_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		ID string
		Endpoint string
		Descr string
		Query string
		Source string
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	Catalog_Redirect = dm.Catalog_PathList
)

//Catalog_Publish annouces the endpoints available for this object
func Catalog_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Catalog_Path, Catalog_Handler)
	mux.HandleFunc(dm.Catalog_PathList, Catalog_HandlerList)
	mux.HandleFunc(dm.Catalog_PathView, Catalog_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Catalog_Title)
    core.Catalog_Add(dm.Catalog_Title, dm.Catalog_Path, "", dm.Catalog_QueryString, "STATE")
}

//Catalog_HandlerList is the handler for the list page
func Catalog_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Catalog
	noItems, returnList, _ := dao.Catalog_GetList()

	pageDetail := Catalog_PageList{
		Title:            CardTitle(dm.Catalog_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Catalog_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Catalog_TemplateList, w, r, pageDetail)

}

//Catalog_HandlerView is the handler used to View a page
func Catalog_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Catalog_QueryString)
	_, rD, _ := dao.Catalog_GetByID(searchID)

	pageDetail := Catalog_Page{
		Title:       CardTitle(dm.Catalog_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Catalog_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.ID = rD.ID
pageDetail.Endpoint = rD.Endpoint
pageDetail.Descr = rD.Descr
pageDetail.Query = rD.Query
pageDetail.Source = rD.Source


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Catalog_TemplateView, w, r, pageDetail)

}

//Catalog_HandlerEdit is the handler used generate the Edit page
func Catalog_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Catalog_QueryString)
	_, rD, _ := dao.Catalog_GetByID(searchID)
	
	pageDetail := Catalog_Page{
		Title:       CardTitle(dm.Catalog_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Catalog_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.ID = rD.ID
pageDetail.Endpoint = rD.Endpoint
pageDetail.Descr = rD.Descr
pageDetail.Query = rD.Query
pageDetail.Source = rD.Source


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Catalog_TemplateEdit, w, r, pageDetail)


}

//Catalog_HandlerSave is the handler used process the saving of an Catalog
func Catalog_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.Catalog
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.ID = r.FormValue(dm.Catalog_ID)
		item.Endpoint = r.FormValue(dm.Catalog_Endpoint)
		item.Descr = r.FormValue(dm.Catalog_Descr)
		item.Query = r.FormValue(dm.Catalog_Query)
		item.Source = r.FormValue(dm.Catalog_Source)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.Catalog_Store(item,r)	

	http.Redirect(w, r, Catalog_Redirect, http.StatusFound)
}

//Catalog_HandlerNew is the handler used process the creation of an Catalog
func Catalog_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Catalog_Page{
		Title:       CardTitle(dm.Catalog_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Catalog_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.ID = ""
pageDetail.Endpoint = ""
pageDetail.Descr = ""
pageDetail.Query = ""
pageDetail.Source = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Catalog_TemplateNew, w, r, pageDetail)

}

//Catalog_HandlerDelete is the handler used process the deletion of an Catalog
func Catalog_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Catalog_QueryString)

	dao.Catalog_Delete(searchID)	

	http.Redirect(w, r, Catalog_Redirect, http.StatusFound)
}