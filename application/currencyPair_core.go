package application
// ----------------------------------------------------------------
// Automatically generated  "/application/currencypair.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CurrencyPair (currencypair)
// Endpoint 	        : CurrencyPair (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 22/11/2021 at 11:25:55
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

//currencypair_PageList provides the information for the template for a list of CurrencyPairs
type CurrencyPair_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.CurrencyPair
}

//currencypair_Page provides the information for the template for an individual CurrencyPair
type CurrencyPair_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
		CodeMajorCurrencyIsoCode string
		CodeMinorCurrencyIsoCode string
		ReciprocalActive string
		Code string
		MajorName string
		MinorName string
		Major_Impl string
		Minor_Impl string
	
	
	
	
	
	
	
	Major_Impl_List	[]dm.Currency
	Minor_Impl_List	[]dm.Currency
	
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
}

const (
	CurrencyPair_Redirect = dm.CurrencyPair_PathList
)

//CurrencyPair_Publish annouces the endpoints available for this object
func CurrencyPair_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.CurrencyPair_PathList, CurrencyPair_HandlerList)
	mux.HandleFunc(dm.CurrencyPair_PathView, CurrencyPair_HandlerView)
	mux.HandleFunc(dm.CurrencyPair_PathEdit, CurrencyPair_HandlerEdit)
	mux.HandleFunc(dm.CurrencyPair_PathNew, CurrencyPair_HandlerNew)
	mux.HandleFunc(dm.CurrencyPair_PathSave, CurrencyPair_HandlerSave)
	mux.HandleFunc(dm.CurrencyPair_PathDelete, CurrencyPair_HandlerDelete)
	logs.Publish("Siena", dm.CurrencyPair_Title)
	
}

//CurrencyPair_HandlerList is the handler for the list page
func CurrencyPair_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.CurrencyPair
	noItems, returnList, _ := dao.CurrencyPair_GetList()


	pageDetail := CurrencyPair_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.CurrencyPair_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.CurrencyPair_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//CurrencyPair_HandlerView is the handler used to View a page
func CurrencyPair_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CurrencyPair_QueryString)
	_, rD, _ := dao.CurrencyPair_GetByID(searchID)

	pageDetail := CurrencyPair_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.CurrencyPair_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
	}

		// 
		// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
pageDetail.CodeMajorCurrencyIsoCode = rD.CodeMajorCurrencyIsoCode
pageDetail.CodeMinorCurrencyIsoCode = rD.CodeMinorCurrencyIsoCode
pageDetail.ReciprocalActive = rD.ReciprocalActive
pageDetail.Code = rD.Code
pageDetail.MajorName = rD.MajorName
pageDetail.MinorName = rD.MinorName
// Automatically generated 22/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,CodeMajorCurrencyIsoCode_Lookup,_:= dao.Currency_GetByID(rD.CodeMajorCurrencyIsoCode)
pageDetail.Major_Impl = CodeMajorCurrencyIsoCode_Lookup.Name
_,CodeMinorCurrencyIsoCode_Lookup,_:= dao.Currency_GetByID(rD.CodeMinorCurrencyIsoCode)
pageDetail.Minor_Impl = CodeMinorCurrencyIsoCode_Lookup.Name
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END


	t, _ := template.ParseFiles(core.GetTemplateID(dm.CurrencyPair_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//CurrencyPair_HandlerEdit is the handler used generate the Edit page
func CurrencyPair_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CurrencyPair_QueryString)
	_, rD, _ := dao.CurrencyPair_GetByID(searchID)
	
	pageDetail := CurrencyPair_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.CurrencyPair_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
	}

		// 
		// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
pageDetail.CodeMajorCurrencyIsoCode = rD.CodeMajorCurrencyIsoCode
pageDetail.CodeMinorCurrencyIsoCode = rD.CodeMinorCurrencyIsoCode
pageDetail.ReciprocalActive = rD.ReciprocalActive
pageDetail.Code = rD.Code
pageDetail.MajorName = rD.MajorName
pageDetail.MinorName = rD.MinorName
// Automatically generated 22/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,CodeMajorCurrencyIsoCode_Lookup,_:= dao.Currency_GetByID(rD.CodeMajorCurrencyIsoCode)
pageDetail.Major_Impl = CodeMajorCurrencyIsoCode_Lookup.Name
_,pageDetail.Major_Impl_List,_ = dao.Currency_GetList()
_,CodeMinorCurrencyIsoCode_Lookup,_:= dao.Currency_GetByID(rD.CodeMinorCurrencyIsoCode)
pageDetail.Minor_Impl = CodeMinorCurrencyIsoCode_Lookup.Name
_,pageDetail.Minor_Impl_List,_ = dao.Currency_GetList()
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//

	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END

	t, _ := template.ParseFiles(core.GetTemplateID(dm.CurrencyPair_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//CurrencyPair_HandlerSave is the handler used process the saving of an CurrencyPair
func CurrencyPair_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.CurrencyPair

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
		item.CodeMajorCurrencyIsoCode = r.FormValue(dm.CurrencyPair_CodeMajorCurrencyIsoCode)
		item.CodeMinorCurrencyIsoCode = r.FormValue(dm.CurrencyPair_CodeMinorCurrencyIsoCode)
		item.ReciprocalActive = r.FormValue(dm.CurrencyPair_ReciprocalActive)
		item.Code = r.FormValue(dm.CurrencyPair_Code)
		item.MajorName = r.FormValue(dm.CurrencyPair_MajorName)
		item.MinorName = r.FormValue(dm.CurrencyPair_MinorName)
		item.Major_Impl = r.FormValue(dm.CurrencyPair_Major_Impl)
		item.Minor_Impl = r.FormValue(dm.CurrencyPair_Minor_Impl)
	
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END

	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END

	dao.CurrencyPair_Store(item)	

	http.Redirect(w, r, CurrencyPair_Redirect, http.StatusFound)
}

//CurrencyPair_HandlerNew is the handler used process the creation of an CurrencyPair
func CurrencyPair_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := CurrencyPair_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.CurrencyPair_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",	
	}

		// 
		// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
pageDetail.CodeMajorCurrencyIsoCode = ""
pageDetail.CodeMinorCurrencyIsoCode = ""
pageDetail.ReciprocalActive = ""
pageDetail.Code = ""
pageDetail.MajorName = ""
pageDetail.MinorName = ""
// Automatically generated 22/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
pageDetail.Major_Impl = ""
_,pageDetail.Major_Impl_List,_ = dao.Currency_GetList()
pageDetail.Minor_Impl = ""
_,pageDetail.Minor_Impl_List,_ = dao.Currency_GetList()
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//

	t, _ := template.ParseFiles(core.GetTemplateID(dm.CurrencyPair_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//CurrencyPair_HandlerDelete is the handler used process the deletion of an CurrencyPair
func CurrencyPair_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.CurrencyPair_QueryString)

	dao.CurrencyPair_Delete(searchID)	

	http.Redirect(w, r, CurrencyPair_Redirect, http.StatusFound)
}