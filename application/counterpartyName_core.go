package application
// ----------------------------------------------------------------
// Automatically generated  "/application/counterpartyname.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyName (counterpartyname)
// Endpoint 	        : CounterpartyName (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:10
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//counterpartyname_PageList provides the information for the template for a list of CounterpartyNames
type CounterpartyName_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.CounterpartyName
}

//counterpartyname_Page provides the information for the template for an individual CounterpartyName
type CounterpartyName_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		NameFirm string
		NameCentre string
		FullName string
		CompID string
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	CounterpartyName_Redirect = dm.CounterpartyName_PathList
)

//CounterpartyName_Publish annouces the endpoints available for this object
func CounterpartyName_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.CounterpartyName_PathList, CounterpartyName_HandlerList)
	mux.HandleFunc(dm.CounterpartyName_PathView, CounterpartyName_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Siena", dm.CounterpartyName_Title)
    //No API
}

//CounterpartyName_HandlerList is the handler for the list page
func CounterpartyName_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.CounterpartyName
	noItems, returnList, _ := dao.CounterpartyName_GetList()

	pageDetail := CounterpartyName_PageList{
		Title:            CardTitle(dm.CounterpartyName_Title, core.Action_List),
		PageTitle:        PageTitle(dm.CounterpartyName_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.CounterpartyName_TemplateList, w, r, pageDetail)

}

//CounterpartyName_HandlerView is the handler used to View a page
func CounterpartyName_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyName_QueryString)
	_, rD, _ := dao.CounterpartyName_GetByID(searchID)

	pageDetail := CounterpartyName_Page{
		Title:       CardTitle(dm.CounterpartyName_Title, core.Action_View),
		PageTitle:   PageTitle(dm.CounterpartyName_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.NameFirm = rD.NameFirm
pageDetail.NameCentre = rD.NameCentre
pageDetail.FullName = rD.FullName
pageDetail.CompID = rD.CompID


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.CounterpartyName_TemplateView, w, r, pageDetail)

}

//CounterpartyName_HandlerEdit is the handler used generate the Edit page
func CounterpartyName_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyName_QueryString)
	_, rD, _ := dao.CounterpartyName_GetByID(searchID)
	
	pageDetail := CounterpartyName_Page{
		Title:       CardTitle(dm.CounterpartyName_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.CounterpartyName_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.NameFirm = rD.NameFirm
pageDetail.NameCentre = rD.NameCentre
pageDetail.FullName = rD.FullName
pageDetail.CompID = rD.CompID


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.CounterpartyName_TemplateEdit, w, r, pageDetail)


}

//CounterpartyName_HandlerSave is the handler used process the saving of an CounterpartyName
func CounterpartyName_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("CompID"))

	var item dm.CounterpartyName
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.NameFirm = r.FormValue(dm.CounterpartyName_NameFirm)
		item.NameCentre = r.FormValue(dm.CounterpartyName_NameCentre)
		item.FullName = r.FormValue(dm.CounterpartyName_FullName)
		item.CompID = r.FormValue(dm.CounterpartyName_CompID)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.CounterpartyName_Store(item,r)	

	http.Redirect(w, r, CounterpartyName_Redirect, http.StatusFound)
}

//CounterpartyName_HandlerNew is the handler used process the creation of an CounterpartyName
func CounterpartyName_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := CounterpartyName_Page{
		Title:       CardTitle(dm.CounterpartyName_Title, core.Action_New),
		PageTitle:   PageTitle(dm.CounterpartyName_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.NameFirm = ""
pageDetail.NameCentre = ""
pageDetail.FullName = ""
pageDetail.CompID = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.CounterpartyName_TemplateNew, w, r, pageDetail)

}

//CounterpartyName_HandlerDelete is the handler used process the deletion of an CounterpartyName
func CounterpartyName_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.CounterpartyName_QueryString)

	dao.CounterpartyName_Delete(searchID)	

	http.Redirect(w, r, CounterpartyName_Redirect, http.StatusFound)
}