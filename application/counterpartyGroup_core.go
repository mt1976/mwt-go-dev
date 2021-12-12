package application
// ----------------------------------------------------------------
// Automatically generated  "/application/counterpartygroup.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyGroup (counterpartygroup)
// Endpoint 	        : CounterpartyGroup (Group)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:09
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//counterpartygroup_PageList provides the information for the template for a list of CounterpartyGroups
type CounterpartyGroup_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.CounterpartyGroup
}

//counterpartygroup_Page provides the information for the template for an individual CounterpartyGroup
type CounterpartyGroup_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		Name string
		CountryCode string
		SuperGroup string
		Country_Lookup string
		Parent_Lookup string
	
	
	
	
	Country_Lookup_List	[]dm.Country
	Parent_Lookup_List	[]dm.CounterpartyGroup
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	CounterpartyGroup_Redirect = dm.CounterpartyGroup_PathList
)

//CounterpartyGroup_Publish annouces the endpoints available for this object
func CounterpartyGroup_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.CounterpartyGroup_PathList, CounterpartyGroup_HandlerList)
	mux.HandleFunc(dm.CounterpartyGroup_PathView, CounterpartyGroup_HandlerView)
	mux.HandleFunc(dm.CounterpartyGroup_PathEdit, CounterpartyGroup_HandlerEdit)
	mux.HandleFunc(dm.CounterpartyGroup_PathNew, CounterpartyGroup_HandlerNew)
	mux.HandleFunc(dm.CounterpartyGroup_PathSave, CounterpartyGroup_HandlerSave)
	mux.HandleFunc(dm.CounterpartyGroup_PathDelete, CounterpartyGroup_HandlerDelete)
	logs.Publish("Siena", dm.CounterpartyGroup_Title)
    //No API
}

//CounterpartyGroup_HandlerList is the handler for the list page
func CounterpartyGroup_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.CounterpartyGroup
	noItems, returnList, _ := dao.CounterpartyGroup_GetList()

	pageDetail := CounterpartyGroup_PageList{
		Title:            CardTitle(dm.CounterpartyGroup_Title, core.Action_List),
		PageTitle:        PageTitle(dm.CounterpartyGroup_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.CounterpartyGroup_TemplateList, w, r, pageDetail)

}

//CounterpartyGroup_HandlerView is the handler used to View a page
func CounterpartyGroup_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyGroup_QueryString)
	_, rD, _ := dao.CounterpartyGroup_GetByID(searchID)

	pageDetail := CounterpartyGroup_Page{
		Title:       CardTitle(dm.CounterpartyGroup_Title, core.Action_View),
		PageTitle:   PageTitle(dm.CounterpartyGroup_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.Name = rD.Name
pageDetail.CountryCode = rD.CountryCode
pageDetail.SuperGroup = rD.SuperGroup


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,CountryCode_Lookup_Name,_:= dao.Country_GetByID(rD.CountryCode)
pageDetail.Country_Lookup = CountryCode_Lookup_Name.Name
_,SuperGroup_Lookup_Name,_:= dao.CounterpartyGroup_GetByID(rD.SuperGroup)
pageDetail.Parent_Lookup = SuperGroup_Lookup_Name.Name
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.CounterpartyGroup_TemplateView, w, r, pageDetail)

}

//CounterpartyGroup_HandlerEdit is the handler used generate the Edit page
func CounterpartyGroup_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyGroup_QueryString)
	_, rD, _ := dao.CounterpartyGroup_GetByID(searchID)
	
	pageDetail := CounterpartyGroup_Page{
		Title:       CardTitle(dm.CounterpartyGroup_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.CounterpartyGroup_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.Name = rD.Name
pageDetail.CountryCode = rD.CountryCode
pageDetail.SuperGroup = rD.SuperGroup


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,CountryCode_Lookup_Name,_:= dao.Country_GetByID(rD.CountryCode)
pageDetail.Country_Lookup = CountryCode_Lookup_Name.Name
_,pageDetail.Country_Lookup_List,_ = dao.Country_GetList()
_,SuperGroup_Lookup_Name,_:= dao.CounterpartyGroup_GetByID(rD.SuperGroup)
pageDetail.Parent_Lookup = SuperGroup_Lookup_Name.Name
_,pageDetail.Parent_Lookup_List,_ = dao.CounterpartyGroup_GetList()
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.CounterpartyGroup_TemplateEdit, w, r, pageDetail)


}

//CounterpartyGroup_HandlerSave is the handler used process the saving of an CounterpartyGroup
func CounterpartyGroup_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Name"))

	var item dm.CounterpartyGroup
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.Name = r.FormValue(dm.CounterpartyGroup_Name)
		item.CountryCode = r.FormValue(dm.CounterpartyGroup_CountryCode)
		item.SuperGroup = r.FormValue(dm.CounterpartyGroup_SuperGroup)
		item.Country_Lookup = r.FormValue(dm.CounterpartyGroup_Country_Lookup)
		item.Parent_Lookup = r.FormValue(dm.CounterpartyGroup_Parent_Lookup)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.CounterpartyGroup_Store(item,r)	

	http.Redirect(w, r, CounterpartyGroup_Redirect, http.StatusFound)
}

//CounterpartyGroup_HandlerNew is the handler used process the creation of an CounterpartyGroup
func CounterpartyGroup_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := CounterpartyGroup_Page{
		Title:       CardTitle(dm.CounterpartyGroup_Title, core.Action_New),
		PageTitle:   PageTitle(dm.CounterpartyGroup_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.Name = ""
pageDetail.CountryCode = ""
pageDetail.SuperGroup = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
pageDetail.Country_Lookup = ""
_,pageDetail.Country_Lookup_List,_ = dao.Country_GetList()
pageDetail.Parent_Lookup = ""
_,pageDetail.Parent_Lookup_List,_ = dao.CounterpartyGroup_GetList()
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.CounterpartyGroup_TemplateNew, w, r, pageDetail)

}

//CounterpartyGroup_HandlerDelete is the handler used process the deletion of an CounterpartyGroup
func CounterpartyGroup_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.CounterpartyGroup_QueryString)

	dao.CounterpartyGroup_Delete(searchID)	

	http.Redirect(w, r, CounterpartyGroup_Redirect, http.StatusFound)
}