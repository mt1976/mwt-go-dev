package application
// ----------------------------------------------------------------
// Automatically generated  "/application/test1.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Test1 (test1)
// Endpoint 	        : Test1 (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:33
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//test1_PageList provides the information for the template for a list of Test1s
type Test1_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Test1
}
//Test1_Redirect provides a page to return to aftern an action
const (
	
	Test1_Redirect = dm.Test1_PathList
	
)

//test1_Page provides the information for the template for an individual Test1
type Test1_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	ID         string
	ID_props     dm.FieldProperties
	Endpoint         string
	Endpoint_props     dm.FieldProperties
	Descr         string
	Descr_props     dm.FieldProperties
	Query         string
	Query_props     dm.FieldProperties
	Source         string
	Source_props     dm.FieldProperties
	Firm         string
	Firm_lookup    []dm.Lookup_Item
	Firm_props     dm.FieldProperties
	YN         string
	YN_lookup    []dm.Lookup_Item
	YN_props     dm.FieldProperties
	User         string
	User_lookup    []dm.Lookup_Item
	User_props     dm.FieldProperties
	Cheese         string
	Cheese_props     dm.FieldProperties
	Onion         string
	Onion_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Test1_Publish annouces the endpoints available for this object
func Test1_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Test1_Path, Test1_Handler)
	mux.HandleFunc(dm.Test1_PathList, Test1_HandlerList)
	mux.HandleFunc(dm.Test1_PathView, Test1_HandlerView)
	mux.HandleFunc(dm.Test1_PathEdit, Test1_HandlerEdit)
	mux.HandleFunc(dm.Test1_PathNew, Test1_HandlerNew)
	mux.HandleFunc(dm.Test1_PathSave, Test1_HandlerSave)
	mux.HandleFunc(dm.Test1_PathDelete, Test1_HandlerDelete)
	logs.Publish("Application", dm.Test1_Title)
    core.Catalog_Add(dm.Test1_Title, dm.Test1_Path, "", dm.Test1_QueryString, "Application")
}


//Test1_HandlerList is the handler for the list page
func Test1_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Test1
	noItems, returnList, _ := dao.Test1_GetList()

	pageDetail := Test1_PageList{
		Title:            CardTitle(dm.Test1_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Test1_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Test1_TemplateList, w, r, pageDetail)

}


//Test1_HandlerView is the handler used to View a page
func Test1_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Test1_QueryString)
	_, rD, _ := dao.Test1_GetByID(searchID)

	pageDetail := Test1_Page{
		Title:       CardTitle(dm.Test1_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Test1_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = test1_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Test1_TemplateView, w, r, pageDetail)

}


//Test1_HandlerEdit is the handler used generate the Edit page
func Test1_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Test1_QueryString)
	_, rD, _ := dao.Test1_GetByID(searchID)
	
	pageDetail := Test1_Page{
		Title:       CardTitle(dm.Test1_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Test1_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = test1_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Test1_TemplateEdit, w, r, pageDetail)
}


//Test1_HandlerSave is the handler used process the saving of an Test1
func Test1_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.Test1
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.ID = r.FormValue(dm.Test1_ID_scrn)
		item.Endpoint = r.FormValue(dm.Test1_Endpoint_scrn)
		item.Descr = r.FormValue(dm.Test1_Descr_scrn)
		item.Query = r.FormValue(dm.Test1_Query_scrn)
		item.Source = r.FormValue(dm.Test1_Source_scrn)
		item.Firm = r.FormValue(dm.Test1_Firm_scrn)
		item.YN = r.FormValue(dm.Test1_YN_scrn)
		item.User = r.FormValue(dm.Test1_User_scrn)
		item.Cheese = r.FormValue(dm.Test1_Cheese_scrn)
		item.Onion = r.FormValue(dm.Test1_Onion_scrn)
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Test1_Store(item,r)	
	http.Redirect(w, r, Test1_Redirect, http.StatusFound)
}


//Test1_HandlerNew is the handler used process the creation of an Test1
func Test1_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.Test1_New()

	pageDetail := Test1_Page{
		Title:       CardTitle(dm.Test1_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Test1_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = test1_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Test1_TemplateNew, w, r, pageDetail)

}	


//Test1_HandlerDelete is the handler used process the deletion of an Test1
func Test1_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Test1_QueryString)

	dao.Test1_Delete(searchID)	

	http.Redirect(w, r, Test1_Redirect, http.StatusFound)
}


// Builds/Popuplates the Test1 Page 
func test1_PopulatePage(rD dm.Test1, pageDetail Test1_Page) Test1_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.ID = rD.ID
	pageDetail.Endpoint = rD.Endpoint
	pageDetail.Descr = rD.Descr
	pageDetail.Query = rD.Query
	pageDetail.Source = rD.Source
	pageDetail.Firm = rD.Firm
	pageDetail.YN = rD.YN
	pageDetail.User = rD.User
	
	pageDetail.Cheese = rD.Cheese
	pageDetail.Onion = rD.Onion
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Firm_lookup = dao.Firm_GetLookup()
	
	
	
	
	pageDetail.YN_lookup = dao.StubLists_Get("")
	
	
	
	pageDetail.User_lookup = dao.StubLists_Get("")
	
	
	
	
	
	
	pageDetail.ID_props = rD.ID_props
	pageDetail.Endpoint_props = rD.Endpoint_props
	pageDetail.Descr_props = rD.Descr_props
	pageDetail.Query_props = rD.Query_props
	pageDetail.Source_props = rD.Source_props
	pageDetail.Firm_props = rD.Firm_props
	pageDetail.YN_props = rD.YN_props
	pageDetail.User_props = rD.User_props
	pageDetail.Cheese_props = rD.Cheese_props
	pageDetail.Onion_props = rD.Onion_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	