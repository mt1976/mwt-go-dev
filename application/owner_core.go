package application
// ----------------------------------------------------------------
// Automatically generated  "/application/owner.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Owner (owner)
// Endpoint 	        : Owner (Owner)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:17
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//owner_PageList provides the information for the template for a list of Owners
type Owner_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Owner
}

//owner_Page provides the information for the template for an individual Owner
type Owner_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		UserName string
		FullName string
		Type string
		TradingEntity string
		DefaultEnterBook string
		EmailAddress string
		Enabled string
		ExternalUserIds string
		Language string
		LocalCurrency string
		Role string
		TelephoneNumber string
		TokenId string
		Entity string
		UserCode string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	Owner_Redirect = dm.Owner_PathList
)

//Owner_Publish annouces the endpoints available for this object
func Owner_Publish(mux http.ServeMux) {
	//No API
	//Cannot List via GUI
	//Cannot View via GUI
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Siena", dm.Owner_Title)
    //No API
}

//Owner_HandlerList is the handler for the list page
func Owner_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Owner
	noItems, returnList, _ := dao.Owner_GetList()

	pageDetail := Owner_PageList{
		Title:            CardTitle(dm.Owner_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Owner_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Owner_TemplateList, w, r, pageDetail)

}

//Owner_HandlerView is the handler used to View a page
func Owner_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Owner_QueryString)
	_, rD, _ := dao.Owner_GetByID(searchID)

	pageDetail := Owner_Page{
		Title:       CardTitle(dm.Owner_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Owner_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.UserName = rD.UserName
pageDetail.FullName = rD.FullName
pageDetail.Type = rD.Type
pageDetail.TradingEntity = rD.TradingEntity
pageDetail.DefaultEnterBook = rD.DefaultEnterBook
pageDetail.EmailAddress = rD.EmailAddress
pageDetail.Enabled = rD.Enabled
pageDetail.ExternalUserIds = rD.ExternalUserIds
pageDetail.Language = rD.Language
pageDetail.LocalCurrency = rD.LocalCurrency
pageDetail.Role = rD.Role
pageDetail.TelephoneNumber = rD.TelephoneNumber
pageDetail.TokenId = rD.TokenId
pageDetail.Entity = rD.Entity
pageDetail.UserCode = rD.UserCode


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Owner_TemplateView, w, r, pageDetail)

}

//Owner_HandlerEdit is the handler used generate the Edit page
func Owner_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Owner_QueryString)
	_, rD, _ := dao.Owner_GetByID(searchID)
	
	pageDetail := Owner_Page{
		Title:       CardTitle(dm.Owner_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Owner_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.UserName = rD.UserName
pageDetail.FullName = rD.FullName
pageDetail.Type = rD.Type
pageDetail.TradingEntity = rD.TradingEntity
pageDetail.DefaultEnterBook = rD.DefaultEnterBook
pageDetail.EmailAddress = rD.EmailAddress
pageDetail.Enabled = rD.Enabled
pageDetail.ExternalUserIds = rD.ExternalUserIds
pageDetail.Language = rD.Language
pageDetail.LocalCurrency = rD.LocalCurrency
pageDetail.Role = rD.Role
pageDetail.TelephoneNumber = rD.TelephoneNumber
pageDetail.TokenId = rD.TokenId
pageDetail.Entity = rD.Entity
pageDetail.UserCode = rD.UserCode


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Owner_TemplateEdit, w, r, pageDetail)


}

//Owner_HandlerSave is the handler used process the saving of an Owner
func Owner_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("UserName"))

	var item dm.Owner
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.UserName = r.FormValue(dm.Owner_UserName)
		item.FullName = r.FormValue(dm.Owner_FullName)
		item.Type = r.FormValue(dm.Owner_Type)
		item.TradingEntity = r.FormValue(dm.Owner_TradingEntity)
		item.DefaultEnterBook = r.FormValue(dm.Owner_DefaultEnterBook)
		item.EmailAddress = r.FormValue(dm.Owner_EmailAddress)
		item.Enabled = r.FormValue(dm.Owner_Enabled)
		item.ExternalUserIds = r.FormValue(dm.Owner_ExternalUserIds)
		item.Language = r.FormValue(dm.Owner_Language)
		item.LocalCurrency = r.FormValue(dm.Owner_LocalCurrency)
		item.Role = r.FormValue(dm.Owner_Role)
		item.TelephoneNumber = r.FormValue(dm.Owner_TelephoneNumber)
		item.TokenId = r.FormValue(dm.Owner_TokenId)
		item.Entity = r.FormValue(dm.Owner_Entity)
		item.UserCode = r.FormValue(dm.Owner_UserCode)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.Owner_Store(item,r)	

	http.Redirect(w, r, Owner_Redirect, http.StatusFound)
}

//Owner_HandlerNew is the handler used process the creation of an Owner
func Owner_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Owner_Page{
		Title:       CardTitle(dm.Owner_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Owner_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.UserName = ""
pageDetail.FullName = ""
pageDetail.Type = ""
pageDetail.TradingEntity = ""
pageDetail.DefaultEnterBook = ""
pageDetail.EmailAddress = ""
pageDetail.Enabled = ""
pageDetail.ExternalUserIds = ""
pageDetail.Language = ""
pageDetail.LocalCurrency = ""
pageDetail.Role = ""
pageDetail.TelephoneNumber = ""
pageDetail.TokenId = ""
pageDetail.Entity = ""
pageDetail.UserCode = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Owner_TemplateNew, w, r, pageDetail)

}

//Owner_HandlerDelete is the handler used process the deletion of an Owner
func Owner_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Owner_QueryString)

	dao.Owner_Delete(searchID)	

	http.Redirect(w, r, Owner_Redirect, http.StatusFound)
}