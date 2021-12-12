package application
// ----------------------------------------------------------------
// Automatically generated  "/application/currencypair.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CurrencyPair (currencypair)
// Endpoint 	        : CurrencyPair (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:11
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//currencypair_PageList provides the information for the template for a list of CurrencyPairs
type CurrencyPair_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.CurrencyPair
}

//currencypair_Page provides the information for the template for an individual CurrencyPair
type CurrencyPair_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		CodeMajorCurrencyIsoCode string
		CodeMinorCurrencyIsoCode string
		ReciprocalActive string
		Code string
		MajorName string
		MinorName string
		Major_Lookup string
		Minor_Lookup string
	
	
	
	
	
	
	
	Major_Lookup_List	[]dm.Currency
	Minor_Lookup_List	[]dm.Currency
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	CurrencyPair_Redirect = dm.CurrencyPair_PathList
)

//CurrencyPair_Publish annouces the endpoints available for this object
func CurrencyPair_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.CurrencyPair_Path, CurrencyPair_Handler)
	mux.HandleFunc(dm.CurrencyPair_PathList, CurrencyPair_HandlerList)
	mux.HandleFunc(dm.CurrencyPair_PathView, CurrencyPair_HandlerView)
	mux.HandleFunc(dm.CurrencyPair_PathEdit, CurrencyPair_HandlerEdit)
	mux.HandleFunc(dm.CurrencyPair_PathNew, CurrencyPair_HandlerNew)
	mux.HandleFunc(dm.CurrencyPair_PathSave, CurrencyPair_HandlerSave)
	mux.HandleFunc(dm.CurrencyPair_PathDelete, CurrencyPair_HandlerDelete)
	logs.Publish("Siena", dm.CurrencyPair_Title)
    core.Catalog_Add(dm.CurrencyPair_Title, dm.CurrencyPair_Path, "", dm.CurrencyPair_QueryString, "APP")
}

//CurrencyPair_HandlerList is the handler for the list page
func CurrencyPair_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.CurrencyPair
	noItems, returnList, _ := dao.CurrencyPair_GetList()

	pageDetail := CurrencyPair_PageList{
		Title:            CardTitle(dm.CurrencyPair_Title, core.Action_List),
		PageTitle:        PageTitle(dm.CurrencyPair_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.CurrencyPair_TemplateList, w, r, pageDetail)

}

//CurrencyPair_HandlerView is the handler used to View a page
func CurrencyPair_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CurrencyPair_QueryString)
	_, rD, _ := dao.CurrencyPair_GetByID(searchID)

	pageDetail := CurrencyPair_Page{
		Title:       CardTitle(dm.CurrencyPair_Title, core.Action_View),
		PageTitle:   PageTitle(dm.CurrencyPair_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.CodeMajorCurrencyIsoCode = rD.CodeMajorCurrencyIsoCode
pageDetail.CodeMinorCurrencyIsoCode = rD.CodeMinorCurrencyIsoCode
pageDetail.ReciprocalActive = rD.ReciprocalActive
pageDetail.Code = rD.Code
pageDetail.MajorName = rD.MajorName
pageDetail.MinorName = rD.MinorName


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,CodeMajorCurrencyIsoCode_Lookup_Name,_:= dao.Currency_GetByID(rD.CodeMajorCurrencyIsoCode)
pageDetail.Major_Lookup = CodeMajorCurrencyIsoCode_Lookup_Name.Name
_,CodeMinorCurrencyIsoCode_Lookup_Name,_:= dao.Currency_GetByID(rD.CodeMinorCurrencyIsoCode)
pageDetail.Minor_Lookup = CodeMinorCurrencyIsoCode_Lookup_Name.Name
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.CurrencyPair_TemplateView, w, r, pageDetail)

}

//CurrencyPair_HandlerEdit is the handler used generate the Edit page
func CurrencyPair_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CurrencyPair_QueryString)
	_, rD, _ := dao.CurrencyPair_GetByID(searchID)
	
	pageDetail := CurrencyPair_Page{
		Title:       CardTitle(dm.CurrencyPair_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.CurrencyPair_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.CodeMajorCurrencyIsoCode = rD.CodeMajorCurrencyIsoCode
pageDetail.CodeMinorCurrencyIsoCode = rD.CodeMinorCurrencyIsoCode
pageDetail.ReciprocalActive = rD.ReciprocalActive
pageDetail.Code = rD.Code
pageDetail.MajorName = rD.MajorName
pageDetail.MinorName = rD.MinorName


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,CodeMajorCurrencyIsoCode_Lookup_Name,_:= dao.Currency_GetByID(rD.CodeMajorCurrencyIsoCode)
pageDetail.Major_Lookup = CodeMajorCurrencyIsoCode_Lookup_Name.Name
_,pageDetail.Major_Lookup_List,_ = dao.Currency_GetList()
_,CodeMinorCurrencyIsoCode_Lookup_Name,_:= dao.Currency_GetByID(rD.CodeMinorCurrencyIsoCode)
pageDetail.Minor_Lookup = CodeMinorCurrencyIsoCode_Lookup_Name.Name
_,pageDetail.Minor_Lookup_List,_ = dao.Currency_GetList()
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.CurrencyPair_TemplateEdit, w, r, pageDetail)


}

//CurrencyPair_HandlerSave is the handler used process the saving of an CurrencyPair
func CurrencyPair_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.CurrencyPair
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.CodeMajorCurrencyIsoCode = r.FormValue(dm.CurrencyPair_CodeMajorCurrencyIsoCode)
		item.CodeMinorCurrencyIsoCode = r.FormValue(dm.CurrencyPair_CodeMinorCurrencyIsoCode)
		item.ReciprocalActive = r.FormValue(dm.CurrencyPair_ReciprocalActive)
		item.Code = r.FormValue(dm.CurrencyPair_Code)
		item.MajorName = r.FormValue(dm.CurrencyPair_MajorName)
		item.MinorName = r.FormValue(dm.CurrencyPair_MinorName)
		item.Major_Lookup = r.FormValue(dm.CurrencyPair_Major_Lookup)
		item.Minor_Lookup = r.FormValue(dm.CurrencyPair_Minor_Lookup)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.CurrencyPair_Store(item,r)	

	http.Redirect(w, r, CurrencyPair_Redirect, http.StatusFound)
}

//CurrencyPair_HandlerNew is the handler used process the creation of an CurrencyPair
func CurrencyPair_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := CurrencyPair_Page{
		Title:       CardTitle(dm.CurrencyPair_Title, core.Action_New),
		PageTitle:   PageTitle(dm.CurrencyPair_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.CodeMajorCurrencyIsoCode = ""
pageDetail.CodeMinorCurrencyIsoCode = ""
pageDetail.ReciprocalActive = ""
pageDetail.Code = ""
pageDetail.MajorName = ""
pageDetail.MinorName = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
pageDetail.Major_Lookup = ""
_,pageDetail.Major_Lookup_List,_ = dao.Currency_GetList()
pageDetail.Minor_Lookup = ""
_,pageDetail.Minor_Lookup_List,_ = dao.Currency_GetList()
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.CurrencyPair_TemplateNew, w, r, pageDetail)

}

//CurrencyPair_HandlerDelete is the handler used process the deletion of an CurrencyPair
func CurrencyPair_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.CurrencyPair_QueryString)

	dao.CurrencyPair_Delete(searchID)	

	http.Redirect(w, r, CurrencyPair_Redirect, http.StatusFound)
}