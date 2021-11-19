package application
// ----------------------------------------------------------------
// Automatically generated  "/application/counterpartygroup.go"
// ----------------------------------------------------------------
// Package            : application
// Object 			  : CounterpartyGroup
// Endpoint Root 	  : CounterpartyGroup
// Search QueryString : Group
// From   			  : 
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 19/11/2021 at 17:16:03
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

//counterpartygroup_PageList provides the information for the template for a list of CounterpartyGroups
type counterpartygroup_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.CounterpartyGroup
}

//counterpartygroup_Page provides the information for the template for an individual CounterpartyGroup
type counterpartygroup_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 19/11/2021 by matttownsend on silicon.local - START
		Name string
		CountryCode string
		SuperGroup string
		Country_Enri string
		Parent_Enri string
	


	
	
	
	Country_Enri_List	[]dm.Country
	Parent_Enri_List	[]dm.CounterpartyGroup
	

	// Automatically generated 19/11/2021 by matttownsend on silicon.local - END
}

const (
	CounterpartyGroup_Redirect = dm.CounterpartyGroup_PathList
)

//CounterpartyGroup_Publish annouces the endpoints available for this object
func CounterpartyGroup_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.CounterpartyGroup_PathList, CounterpartyGroup_HandlerList)
	mux.HandleFunc(dm.CounterpartyGroup_PathView, CounterpartyGroup_HandlerView)
	mux.HandleFunc(dm.CounterpartyGroup_PathEdit, CounterpartyGroup_HandlerEdit)
	mux.HandleFunc(dm.CounterpartyGroup_PathNew, CounterpartyGroup_HandlerNew)
	mux.HandleFunc(dm.CounterpartyGroup_PathSave, CounterpartyGroup_HandlerSave)
	mux.HandleFunc(dm.CounterpartyGroup_PathDelete, CounterpartyGroup_HandlerDelete)
	logs.Publish("Siena", dm.CounterpartyGroup_Title)
}

//CounterpartyGroup_HandlerList is the handler for the list page
func CounterpartyGroup_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.CounterpartyGroup
	noItems, returnList, _ := dao.CounterpartyGroup_GetList()


	pageDetail := counterpartygroup_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.CounterpartyGroup_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.CounterpartyGroup_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//CounterpartyGroup_HandlerView is the handler used to View a page
func CounterpartyGroup_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyGroup_QueryString)
	_, rD, _ := dao.CounterpartyGroup_GetByID(searchID)

	pageDetail := counterpartygroup_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.CounterpartyGroup_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
	}

		// 
		// Automatically generated 19/11/2021 by matttownsend on silicon.local - START
pageDetail.Name = rD.Name
pageDetail.CountryCode = rD.CountryCode
pageDetail.SuperGroup = rD.SuperGroup
// Automatically generated 19/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,CountryCode_Lookup,_:= dao.Country_GetByID(rD.CountryCode)
pageDetail.Country_Enri = CountryCode_Lookup.Name
_,SuperGroup_Lookup,_:= dao.CounterpartyGroup_GetByID(rD.SuperGroup)
pageDetail.Parent_Enri = SuperGroup_Lookup.Name
// Automatically generated 19/11/2021 by matttownsend on silicon.local - END
		//

	t, _ := template.ParseFiles(core.GetTemplateID(dm.CounterpartyGroup_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//CounterpartyGroup_HandlerEdit is the handler used generate the Edit page
func CounterpartyGroup_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyGroup_QueryString)
	_, rD, _ := dao.CounterpartyGroup_GetByID(searchID)
	
	pageDetail := counterpartygroup_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.CounterpartyGroup_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
	}

		// 
		// Automatically generated 19/11/2021 by matttownsend on silicon.local - START
pageDetail.Name = rD.Name
pageDetail.CountryCode = rD.CountryCode
pageDetail.SuperGroup = rD.SuperGroup
// Automatically generated 19/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,CountryCode_Lookup,_:= dao.Country_GetByID(rD.CountryCode)
pageDetail.Country_Enri = CountryCode_Lookup.Name
_,pageDetail.Country_Enri_List,_ = dao.Country_GetList()
_,SuperGroup_Lookup,_:= dao.CounterpartyGroup_GetByID(rD.SuperGroup)
pageDetail.Parent_Enri = SuperGroup_Lookup.Name
_,pageDetail.Parent_Enri_List,_ = dao.CounterpartyGroup_GetList()
// Automatically generated 19/11/2021 by matttownsend on silicon.local - END
		//


	t, _ := template.ParseFiles(core.GetTemplateID(dm.CounterpartyGroup_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//CounterpartyGroup_HandlerSave is the handler used process the saving of an CounterpartyGroup
func CounterpartyGroup_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.CounterpartyGroup

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 19/11/2021 by matttownsend on silicon.local - START
		item.Name = r.FormValue(dm.CounterpartyGroup_Name)
		item.CountryCode = r.FormValue(dm.CounterpartyGroup_CountryCode)
		item.SuperGroup = r.FormValue(dm.CounterpartyGroup_SuperGroup)
		item.Country_Enri = r.FormValue(dm.CounterpartyGroup_Country_Enri)
		item.Parent_Enri = r.FormValue(dm.CounterpartyGroup_Parent_Enri)
	
	// Automatically generated 19/11/2021 by matttownsend on silicon.local - END

	dao.CounterpartyGroup_Store(item)	

	http.Redirect(w, r, CounterpartyGroup_Redirect, http.StatusFound)
}

//CounterpartyGroup_HandlerNew is the handler used process the creation of an CounterpartyGroup
func CounterpartyGroup_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := counterpartygroup_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.CounterpartyGroup_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",	
	}

		// 
		// Automatically generated 19/11/2021 by matttownsend on silicon.local - START
pageDetail.Name = ""
pageDetail.CountryCode = ""
pageDetail.SuperGroup = ""
// Automatically generated 19/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
pageDetail.Country_Enri = ""
_,pageDetail.Country_Enri_List,_ = dao.Country_GetList()
pageDetail.Parent_Enri = ""
_,pageDetail.Parent_Enri_List,_ = dao.CounterpartyGroup_GetList()
// Automatically generated 19/11/2021 by matttownsend on silicon.local - END
		//

	t, _ := template.ParseFiles(core.GetTemplateID(dm.CounterpartyGroup_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//CounterpartyGroup_HandlerDelete is the handler used process the deletion of an CounterpartyGroup
func CounterpartyGroup_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.CounterpartyGroup_QueryString)

	dao.CounterpartyGroup_Delete(searchID)	

	http.Redirect(w, r, CounterpartyGroup_Redirect, http.StatusFound)
}