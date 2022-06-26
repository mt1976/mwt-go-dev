package application
// ----------------------------------------------------------------
// Automatically generated  "/application/credentialsaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CredentialsAction (credentialsaction)
// Endpoint 	        : CredentialsAction (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:25
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//credentialsaction_PageList provides the information for the template for a list of CredentialsActions
type CredentialsAction_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.CredentialsAction
}
//CredentialsAction_Redirect provides a page to return to aftern an action
const (
	
	CredentialsAction_Redirect = dm.CredentialsAction_PathView
	
)

//credentialsaction_Page provides the information for the template for an individual CredentialsAction
type CredentialsAction_Page struct {
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
	User         string
	User_lookup    []dm.Lookup_Item
	User_props     dm.FieldProperties
	Action         string
	Action_lookup    []dm.Lookup_Item
	Action_props     dm.FieldProperties
	Notes         string
	Notes_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//CredentialsAction_Publish annouces the endpoints available for this object
func CredentialsAction_Publish(mux http.ServeMux) {
	//No API
	//Cannot List via GUI
	mux.HandleFunc(dm.CredentialsAction_PathView, CredentialsAction_HandlerView)
	mux.HandleFunc(dm.CredentialsAction_PathEdit, CredentialsAction_HandlerEdit)
	mux.HandleFunc(dm.CredentialsAction_PathNew, CredentialsAction_HandlerNew)
	mux.HandleFunc(dm.CredentialsAction_PathSave, CredentialsAction_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.CredentialsAction_Title)
    //No API
}



//CredentialsAction_HandlerView is the handler used to View a page
func CredentialsAction_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CredentialsAction_QueryString)
	_, rD, _ := dao.CredentialsAction_GetByID(searchID)

	pageDetail := CredentialsAction_Page{
		Title:       CardTitle(dm.CredentialsAction_Title, core.Action_View),
		PageTitle:   PageTitle(dm.CredentialsAction_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = credentialsaction_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CredentialsAction_TemplateView, w, r, pageDetail)

}


//CredentialsAction_HandlerEdit is the handler used generate the Edit page
func CredentialsAction_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CredentialsAction_QueryString)
	_, rD, _ := dao.CredentialsAction_GetByID(searchID)
	
	pageDetail := CredentialsAction_Page{
		Title:       CardTitle(dm.CredentialsAction_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.CredentialsAction_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = credentialsaction_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CredentialsAction_TemplateEdit, w, r, pageDetail)
}


//CredentialsAction_HandlerSave is the handler used process the saving of an CredentialsAction
func CredentialsAction_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.CredentialsAction
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.ID = r.FormValue(dm.CredentialsAction_ID_scrn)
		item.User = r.FormValue(dm.CredentialsAction_User_scrn)
		item.Action = r.FormValue(dm.CredentialsAction_Action_scrn)
		item.Notes = r.FormValue(dm.CredentialsAction_Notes_scrn)
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.CredentialsAction_Store(item,r)	
	http.Redirect(w, r, CredentialsAction_Redirect, http.StatusFound)
}


//CredentialsAction_HandlerNew is the handler used process the creation of an CredentialsAction
func CredentialsAction_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.CredentialsAction_New()

	pageDetail := CredentialsAction_Page{
		Title:       CardTitle(dm.CredentialsAction_Title, core.Action_New),
		PageTitle:   PageTitle(dm.CredentialsAction_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = credentialsaction_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CredentialsAction_TemplateNew, w, r, pageDetail)

}	



// Builds/Popuplates the CredentialsAction Page 
func credentialsaction_PopulatePage(rD dm.CredentialsAction, pageDetail CredentialsAction_Page) CredentialsAction_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.ID = rD.ID
	pageDetail.User = rD.User
	pageDetail.Action = rD.Action
	pageDetail.Notes = rD.Notes
	
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	pageDetail.User_lookup = dao.Credentials_GetLookup()
	
	
	
	
	pageDetail.Action_lookup = dao.StubLists_Get("credentialStates")
	
	
	
	
	pageDetail.ID_props = rD.ID_props
	pageDetail.User_props = rD.User_props
	pageDetail.Action_props = rD.Action_props
	pageDetail.Notes_props = rD.Notes_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	