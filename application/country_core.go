package application
// ----------------------------------------------------------------
// Automatically generated  "/application/country.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Country (country)
// Endpoint 	        : Country (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 28/11/2021 at 21:22:07
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"html/template"
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//country_PageList provides the information for the template for a list of Countrys
type Country_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Country
}

//country_Page provides the information for the template for an individual Country
type Country_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
		Code string
		Name string
		ShortCode string
		EU_EEA string
		HolidaysWeekend string
	
	
	
	
	
	
	
	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END
}

const (
	Country_Redirect = dm.Country_PathList
)

//Country_Publish annouces the endpoints available for this object
func Country_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Country_PathList, Country_HandlerList)
	mux.HandleFunc(dm.Country_PathView, Country_HandlerView)
	mux.HandleFunc(dm.Country_PathEdit, Country_HandlerEdit)
	mux.HandleFunc(dm.Country_PathNew, Country_HandlerNew)
	mux.HandleFunc(dm.Country_PathSave, Country_HandlerSave)
	
	logs.Publish("Siena", dm.Country_Title)
	
}

//Country_HandlerList is the handler for the list page
func Country_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Country
	noItems, returnList, _ := dao.Country_GetList()


	pageDetail := Country_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.Country_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Country_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//Country_HandlerView is the handler used to View a page
func Country_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Country_QueryString)
	_, rD, _ := dao.Country_GetByID(searchID)

	pageDetail := Country_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Country_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

		// 
		// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.ShortCode = rD.ShortCode
pageDetail.EU_EEA = rD.EU_EEA
pageDetail.HolidaysWeekend = rD.HolidaysWeekend
// Automatically generated 28/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 28/11/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END


	t, _ := template.ParseFiles(core.GetTemplateID(dm.Country_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Country_HandlerEdit is the handler used generate the Edit page
func Country_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Country_QueryString)
	_, rD, _ := dao.Country_GetByID(searchID)
	
	pageDetail := Country_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Country_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

		// 
		// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.ShortCode = rD.ShortCode
pageDetail.EU_EEA = rD.EU_EEA
pageDetail.HolidaysWeekend = rD.HolidaysWeekend
// Automatically generated 28/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 28/11/2021 by matttownsend on silicon.local - END
		//

	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Country_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Country_HandlerSave is the handler used process the saving of an Country
func Country_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.Country
	// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
		item.Code = r.FormValue(dm.Country_Code)
		item.Name = r.FormValue(dm.Country_Name)
		item.ShortCode = r.FormValue(dm.Country_ShortCode)
		item.EU_EEA = r.FormValue(dm.Country_EU_EEA)
		item.HolidaysWeekend = r.FormValue(dm.Country_HolidaysWeekend)
	
	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END

	// Automatically generated 28/11/2021 by matttownsend on silicon.local - END

	dao.Country_Store(item)	

	http.Redirect(w, r, Country_Redirect, http.StatusFound)
}

//Country_HandlerNew is the handler used process the creation of an Country
func Country_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Country_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Country_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

		// 
		// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
pageDetail.Code = ""
pageDetail.Name = ""
pageDetail.ShortCode = ""
pageDetail.EU_EEA = ""
pageDetail.HolidaysWeekend = ""
// Automatically generated 28/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 28/11/2021 by matttownsend on silicon.local - END
		//

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Country_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Country_HandlerDelete is the handler used process the deletion of an Country
func Country_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Country_QueryString)

	dao.Country_Delete(searchID)	

	http.Redirect(w, r, Country_Redirect, http.StatusFound)
}