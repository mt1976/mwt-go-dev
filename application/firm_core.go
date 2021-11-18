package application
// ----------------------------------------------------------------
// Automatically generated  "/application/firm.go"
// ----------------------------------------------------------------
// Package            : application
// Object 			  : Firm
// Endpoint Root 	  : Firm
// Search QueryString : FirmName
// From   			  : 
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 18/11/2021 at 21:34:20
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

//firm_PageList provides the information for the template for a list of Firms
type firm_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Firm
}

//firm_Page provides the information for the template for an individual Firm
type firm_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
		FirmName string
		FullName string
		Country string
		Sector string
		SectorName string
		CountryName string
	
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - END
}

const (
	Firm_Redirect = dm.Firm_PathList
)

//Firm_Publish annouces the endpoints available for this object
func Firm_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Firm_PathList, Firm_HandlerList)
	mux.HandleFunc(dm.Firm_PathView, Firm_HandlerView)
	mux.HandleFunc(dm.Firm_PathEdit, Firm_HandlerEdit)
	mux.HandleFunc(dm.Firm_PathNew, Firm_HandlerNew)
	mux.HandleFunc(dm.Firm_PathSave, Firm_HandlerSave)
	mux.HandleFunc(dm.Firm_PathDelete, Firm_HandlerDelete)
	logs.Publish("Siena", dm.Firm_Title)
}

//Firm_HandlerList is the handler for the list page
func Firm_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Firm
	noItems, returnList, _ := dao.Firm_GetList()


	pageDetail := firm_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.Firm_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Firm_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//Firm_HandlerView is the handler used to View a page
func Firm_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Firm_QueryString)
	_, rD, _ := dao.Firm_GetByID(searchID)

	pageDetail := firm_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Firm_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
		// 
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
		FirmName: rD.FirmName,
		FullName: rD.FullName,
		Country: rD.Country,
		Sector: rD.Sector,
		SectorName: rD.SectorName,
		CountryName: rD.CountryName,
		
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - END
		//
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Firm_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Firm_HandlerEdit is the handler used generate the Edit page
func Firm_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Firm_QueryString)
	_, rD, _ := dao.Firm_GetByID(searchID)
	
	pageDetail := firm_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Firm_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
			FirmName: rD.FirmName,
			FullName: rD.FullName,
			Country: rD.Country,
			Sector: rD.Sector,
			SectorName: rD.SectorName,
			CountryName: rD.CountryName,
		
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - END
		//Post Import Actions - START

		// Post Import Actions - END
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Firm_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Firm_HandlerSave is the handler used process the saving of an Firm
func Firm_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.Firm

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
		item.FirmName = r.FormValue(dm.Firm_FirmName)
		item.FullName = r.FormValue(dm.Firm_FullName)
		item.Country = r.FormValue(dm.Firm_Country)
		item.Sector = r.FormValue(dm.Firm_Sector)
		item.SectorName = r.FormValue(dm.Firm_SectorName)
		item.CountryName = r.FormValue(dm.Firm_CountryName)
	
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - END

	dao.Firm_Store(item)	

	http.Redirect(w, r, Firm_Redirect, http.StatusFound)
}

//Firm_HandlerNew is the handler used process the creation of an Firm
func Firm_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := firm_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Firm_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
			FirmName: "",
			FullName: "",
			Country: "",
			Sector: "",
			SectorName: "",
			CountryName: "",
		
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - END
		//
		// Post Import Actions - START

		// Post Import Actions - END
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Firm_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Firm_HandlerDelete is the handler used process the deletion of an Firm
func Firm_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Firm_QueryString)

	dao.Firm_Delete(searchID)	

	http.Redirect(w, r, Firm_Redirect, http.StatusFound)
}