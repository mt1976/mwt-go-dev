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
// Date & Time		    : 12/12/2021 at 16:13:15
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
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Firm
}

//firm_Page provides the information for the template for an individual Firm
type Firm_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		FirmName string
		FullName string
		Country string
		Sector string
		Sector_Lookup string
		Country_Lookup string
	
	
	
	
	
	Sector_Lookup_List	[]dm.Sector
	Country_Lookup_List	[]dm.Country
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	Firm_Redirect = dm.Firm_PathList
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
	logs.Publish("Siena", dm.Firm_Title)
    core.Catalog_Add(dm.Firm_Title, dm.Firm_Path, "", dm.Firm_QueryString, "APP")
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

	pageDetail := Firm_PageList{
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

	pageDetail := Firm_Page{
		Title:       CardTitle(dm.Firm_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Firm_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.FirmName = rD.FirmName
pageDetail.FullName = rD.FullName
pageDetail.Country = rD.Country
pageDetail.Sector = rD.Sector


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,Sector_Lookup_Name,_:= dao.Sector_GetByID(rD.Sector)
pageDetail.Sector_Lookup = Sector_Lookup_Name.Name
_,Country_Lookup_Name,_:= dao.Country_GetByID(rD.Country)
pageDetail.Country_Lookup = Country_Lookup_Name.Name
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

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
	
	pageDetail := Firm_Page{
		Title:       CardTitle(dm.Firm_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Firm_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.FirmName = rD.FirmName
pageDetail.FullName = rD.FullName
pageDetail.Country = rD.Country
pageDetail.Sector = rD.Sector


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,Sector_Lookup_Name,_:= dao.Sector_GetByID(rD.Sector)
pageDetail.Sector_Lookup = Sector_Lookup_Name.Name
_,pageDetail.Sector_Lookup_List,_ = dao.Sector_GetList()
_,Country_Lookup_Name,_:= dao.Country_GetByID(rD.Country)
pageDetail.Country_Lookup = Country_Lookup_Name.Name
_,pageDetail.Country_Lookup_List,_ = dao.Country_GetList()
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

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
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.FirmName = r.FormValue(dm.Firm_FirmName)
		item.FullName = r.FormValue(dm.Firm_FullName)
		item.Country = r.FormValue(dm.Firm_Country)
		item.Sector = r.FormValue(dm.Firm_Sector)
		item.Sector_Lookup = r.FormValue(dm.Firm_Sector_Lookup)
		item.Country_Lookup = r.FormValue(dm.Firm_Country_Lookup)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.Firm_Store(item,r)	

	http.Redirect(w, r, Firm_Redirect, http.StatusFound)
}

//Firm_HandlerNew is the handler used process the creation of an Firm
func Firm_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Firm_Page{
		Title:       CardTitle(dm.Firm_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Firm_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.FirmName = ""
pageDetail.FullName = ""
pageDetail.Country = ""
pageDetail.Sector = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
pageDetail.Sector_Lookup = ""
_,pageDetail.Sector_Lookup_List,_ = dao.Sector_GetList()
pageDetail.Country_Lookup = ""
_,pageDetail.Country_Lookup_List,_ = dao.Country_GetList()
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

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

	http.Redirect(w, r, Firm_Redirect, http.StatusFound)
}