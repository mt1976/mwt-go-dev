package application
// ----------------------------------------------------------------
// Automatically generated  "/application/firm.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Firm (firm)
// Endpoint 	        : Firm (FirmName)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 01/12/2021 at 20:36:40
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//firm_PageList provides the information for the template for a list of Firms
type Firm_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Firm
}

//firm_Page provides the information for the template for an individual Firm
type Firm_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// Automatically generated 01/12/2021 by matttownsend on silicon.local - START
		FirmName string
		FullName string
		Country string
		Sector string
		Sector_Impl string
		Country_Impl string
	
	
	
	
	
	Sector_Impl_List	[]dm.Sector
	Country_Impl_List	[]dm.Country
	
	// Automatically generated 01/12/2021 by matttownsend on silicon.local - END
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

	pageDetail := Firm_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.Firm_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

		ExecuteTemplate(dm.Firm_TemplateList, w, r, pageDetail)

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

	pageDetail := Firm_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Firm_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

		// 
		// Automatically generated 01/12/2021 by matttownsend on silicon.local - START
pageDetail.FirmName = rD.FirmName
pageDetail.FullName = rD.FullName
pageDetail.Country = rD.Country
pageDetail.Sector = rD.Sector
// Automatically generated 01/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,Sector_Lookup,_:= dao.Sector_GetByID(rD.Sector)
pageDetail.Sector_Impl = Sector_Lookup.Name
_,Country_Lookup,_:= dao.Country_GetByID(rD.Country)
pageDetail.Country_Impl = Country_Lookup.Name
// Automatically generated 01/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 01/12/2021 by matttownsend on silicon.local - END


		ExecuteTemplate(dm.Firm_TemplateView, w, r, pageDetail)


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
	
	pageDetail := Firm_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Firm_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

		// 
		// Automatically generated 01/12/2021 by matttownsend on silicon.local - START
pageDetail.FirmName = rD.FirmName
pageDetail.FullName = rD.FullName
pageDetail.Country = rD.Country
pageDetail.Sector = rD.Sector
// Automatically generated 01/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,Sector_Lookup,_:= dao.Sector_GetByID(rD.Sector)
pageDetail.Sector_Impl = Sector_Lookup.Name
_,pageDetail.Sector_Impl_List,_ = dao.Sector_GetList()
_,Country_Lookup,_:= dao.Country_GetByID(rD.Country)
pageDetail.Country_Impl = Country_Lookup.Name
_,pageDetail.Country_Impl_List,_ = dao.Country_GetList()
// Automatically generated 01/12/2021 by matttownsend on silicon.local - END
		//

	// Automatically generated 01/12/2021 by matttownsend on silicon.local - END

		ExecuteTemplate(dm.Firm_TemplateEdit, w, r, pageDetail)


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
	logs.Servicing(r.URL.Path+r.FormValue("FirmName"))

	var item dm.Firm
	// Automatically generated 01/12/2021 by matttownsend on silicon.local - START
		item.FirmName = r.FormValue(dm.Firm_FirmName)
		item.FullName = r.FormValue(dm.Firm_FullName)
		item.Country = r.FormValue(dm.Firm_Country)
		item.Sector = r.FormValue(dm.Firm_Sector)
		item.Sector_Impl = r.FormValue(dm.Firm_Sector_Impl)
		item.Country_Impl = r.FormValue(dm.Firm_Country_Impl)
	
	// Automatically generated 01/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 01/12/2021 by matttownsend on silicon.local - END

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

	pageDetail := Firm_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Firm_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

		// 
		// Automatically generated 01/12/2021 by matttownsend on silicon.local - START
pageDetail.FirmName = ""
pageDetail.FullName = ""
pageDetail.Country = ""
pageDetail.Sector = ""
// Automatically generated 01/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
pageDetail.Sector_Impl = ""
_,pageDetail.Sector_Impl_List,_ = dao.Sector_GetList()
pageDetail.Country_Impl = ""
_,pageDetail.Country_Impl_List,_ = dao.Country_GetList()
// Automatically generated 01/12/2021 by matttownsend on silicon.local - END
		//

		ExecuteTemplate(dm.Firm_TemplateNew, w, r, pageDetail)

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