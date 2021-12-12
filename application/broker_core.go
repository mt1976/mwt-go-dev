package application
// ----------------------------------------------------------------
// Automatically generated  "/application/broker.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Broker (broker)
// Endpoint 	        : Broker (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:07
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//broker_PageList provides the information for the template for a list of Brokers
type Broker_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Broker
}

//broker_Page provides the information for the template for an individual Broker
type Broker_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		Code string
		Name string
		FullName string
		Contact string
		Address string
		LEI string
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	Broker_Redirect = dm.Broker_PathList
)

//Broker_Publish annouces the endpoints available for this object
func Broker_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Broker_Path, Broker_Handler)
	mux.HandleFunc(dm.Broker_PathList, Broker_HandlerList)
	mux.HandleFunc(dm.Broker_PathView, Broker_HandlerView)
	mux.HandleFunc(dm.Broker_PathEdit, Broker_HandlerEdit)
	mux.HandleFunc(dm.Broker_PathNew, Broker_HandlerNew)
	mux.HandleFunc(dm.Broker_PathSave, Broker_HandlerSave)
	mux.HandleFunc(dm.Broker_PathDelete, Broker_HandlerDelete)
	logs.Publish("Siena", dm.Broker_Title)
    core.Catalog_Add(dm.Broker_Title, dm.Broker_Path, "", dm.Broker_QueryString, "APP")
}

//Broker_HandlerList is the handler for the list page
func Broker_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Broker
	noItems, returnList, _ := dao.Broker_GetList()

	pageDetail := Broker_PageList{
		Title:            CardTitle(dm.Broker_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Broker_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Broker_TemplateList, w, r, pageDetail)

}

//Broker_HandlerView is the handler used to View a page
func Broker_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Broker_QueryString)
	_, rD, _ := dao.Broker_GetByID(searchID)

	pageDetail := Broker_Page{
		Title:       CardTitle(dm.Broker_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Broker_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.FullName = rD.FullName
pageDetail.Contact = rD.Contact
pageDetail.Address = rD.Address
pageDetail.LEI = rD.LEI


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Broker_TemplateView, w, r, pageDetail)

}

//Broker_HandlerEdit is the handler used generate the Edit page
func Broker_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Broker_QueryString)
	_, rD, _ := dao.Broker_GetByID(searchID)
	
	pageDetail := Broker_Page{
		Title:       CardTitle(dm.Broker_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Broker_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.FullName = rD.FullName
pageDetail.Contact = rD.Contact
pageDetail.Address = rD.Address
pageDetail.LEI = rD.LEI


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Broker_TemplateEdit, w, r, pageDetail)


}

//Broker_HandlerSave is the handler used process the saving of an Broker
func Broker_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.Broker
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.Code = r.FormValue(dm.Broker_Code)
		item.Name = r.FormValue(dm.Broker_Name)
		item.FullName = r.FormValue(dm.Broker_FullName)
		item.Contact = r.FormValue(dm.Broker_Contact)
		item.Address = r.FormValue(dm.Broker_Address)
		item.LEI = r.FormValue(dm.Broker_LEI)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.Broker_Store(item,r)	

	http.Redirect(w, r, Broker_Redirect, http.StatusFound)
}

//Broker_HandlerNew is the handler used process the creation of an Broker
func Broker_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Broker_Page{
		Title:       CardTitle(dm.Broker_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Broker_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.Code = ""
pageDetail.Name = ""
pageDetail.FullName = ""
pageDetail.Contact = ""
pageDetail.Address = ""
pageDetail.LEI = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Broker_TemplateNew, w, r, pageDetail)

}

//Broker_HandlerDelete is the handler used process the deletion of an Broker
func Broker_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Broker_QueryString)

	dao.Broker_Delete(searchID)	

	http.Redirect(w, r, Broker_Redirect, http.StatusFound)
}