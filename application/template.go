package application
// ----------------------------------------------------------------
// Automatically generated  "/application/template.go"
// ----------------------------------------------------------------
// Package            : application
// Object 			  : Template
// Endpoint Root 	  : Template
// Search QueryString : TemplateID
// From   			  : 
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 15/11/2021 at 23:50:09
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"html/template"
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//template_PageList provides the information for the template for a list of Templates
type template_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Template
}

//template_Page provides the information for the template for an individual Template
type template_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
		SYSId string
		FIELD1 string
		FIELD2 string
		SYSCreated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdated string
		SYSUpdatedHost string
		SYSUpdatedBy string
		ID string
	
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - END
}

const (
	Template_Redirect = dm.Template_PathList
)

//Template_Publish annouces the endpoints available for this object
func Template_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Template_PathList, Template_HandlerList)
	mux.HandleFunc(dm.Template_PathView, Template_HandlerView)
	mux.HandleFunc(dm.Template_PathEdit, Template_HandlerEdit)
	mux.HandleFunc(dm.Template_PathNew, Template_HandlerNew)
	mux.HandleFunc(dm.Template_PathSave, Template_HandlerSave)
	mux.HandleFunc(dm.Template_PathDelete, Template_HandlerDelete)
	logs.Publish("Application", dm.Template_Title)
}

//Template_HandlerList is the handler for the list page
func Template_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Template
	noItems, returnList, _ := dao.Template_GetList()


	pageDetail := template_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.Template_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Template_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//Template_HandlerView is the handler used to View a page
func Template_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Template_QueryString)
	_, rD, _ := dao.Template_GetByID(searchID)

	pageDetail := template_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Template_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
		// 
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
		SYSId: rD.SYSId,
		FIELD1: rD.FIELD1,
		FIELD2: rD.FIELD2,
		SYSCreated: rD.SYSCreated,
		SYSCreatedBy: rD.SYSCreatedBy,
		SYSCreatedHost: rD.SYSCreatedHost,
		SYSUpdated: rD.SYSUpdated,
		SYSUpdatedHost: rD.SYSUpdatedHost,
		SYSUpdatedBy: rD.SYSUpdatedBy,
		ID: rD.ID,
		
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - END
		//
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Template_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Template_HandlerEdit is the handler used generate the Edit page
func Template_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Template_QueryString)
	_, rD, _ := dao.Template_GetByID(searchID)
	
	pageDetail := template_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Template_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
			SYSId: rD.SYSId,
			FIELD1: rD.FIELD1,
			FIELD2: rD.FIELD2,
			SYSCreated: rD.SYSCreated,
			SYSCreatedBy: rD.SYSCreatedBy,
			SYSCreatedHost: rD.SYSCreatedHost,
			SYSUpdated: rD.SYSUpdated,
			SYSUpdatedHost: rD.SYSUpdatedHost,
			SYSUpdatedBy: rD.SYSUpdatedBy,
			ID: rD.ID,
		
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - END
		//Post Import Actions - START

		// Post Import Actions - END
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Template_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Template_HandlerSave is the handler used process the saving of an Template
func Template_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.Template

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue("SYSId")
		item.FIELD1 = r.FormValue("FIELD1")
		item.FIELD2 = r.FormValue("FIELD2")
		item.SYSCreated = r.FormValue("SYSCreated")
		item.SYSCreatedBy = r.FormValue("SYSCreatedBy")
		item.SYSCreatedHost = r.FormValue("SYSCreatedHost")
		item.SYSUpdated = r.FormValue("SYSUpdated")
		item.SYSUpdatedHost = r.FormValue("SYSUpdatedHost")
		item.SYSUpdatedBy = r.FormValue("SYSUpdatedBy")
		item.ID = r.FormValue("ID")
	
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - END

	dao.Template_Store(item)	

	http.Redirect(w, r, Template_Redirect, http.StatusFound)
}

//Template_HandlerNew is the handler used process the creation of an Template
func Template_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := template_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Template_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
			SYSId: "0",
			FIELD1: "",
			FIELD2: "",
			SYSCreated: "",
			SYSCreatedBy: "",
			SYSCreatedHost: "",
			SYSUpdated: "",
			SYSUpdatedHost: "",
			SYSUpdatedBy: "",
			ID: "",
		
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - END
		//
		// Post Import Actions - START

		// Post Import Actions - END
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Template_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Template_HandlerDelete is the handler used process the deletion of an Template
func Template_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Template_QueryString)

	dao.Template_Delete(searchID)	

	http.Redirect(w, r, Template_Redirect, http.StatusFound)
}