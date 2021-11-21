package application
// ----------------------------------------------------------------
// Automatically generated  "/application/accountladder.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : AccountLadder (accountladder)
// Endpoint 	        : AccountLadder (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 19:48:01
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

//accountladder_PageList provides the information for the template for a list of AccountLadders
type AccountLadder_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.AccountLadder
}

//accountladder_Page provides the information for the template for an individual AccountLadder
type AccountLadder_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
		SienaReference string
		BusinessDate string
		ContractNumber string
		Balance string
		DealtCcy string
		AmountDp string
	
	
	
	
	
	
	
	
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
}

const (
	AccountLadder_Redirect = dm.AccountLadder_PathList
)

//AccountLadder_Publish annouces the endpoints available for this object
func AccountLadder_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.AccountLadder_PathList, AccountLadder_HandlerList)
	mux.HandleFunc(dm.AccountLadder_PathView, AccountLadder_HandlerView)
	
	
	
	
	logs.Publish("Siena", dm.AccountLadder_Title)
	
}

//AccountLadder_HandlerList is the handler for the list page
func AccountLadder_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.AccountLadder
	noItems, returnList, _ := dao.AccountLadder_GetList()


	pageDetail := AccountLadder_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.AccountLadder_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.AccountLadder_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//AccountLadder_HandlerView is the handler used to View a page
func AccountLadder_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.AccountLadder_QueryString)
	_, rD, _ := dao.AccountLadder_GetByID(searchID)

	pageDetail := AccountLadder_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.AccountLadder_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
	}

		// 
		// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
pageDetail.SienaReference = rD.SienaReference
pageDetail.BusinessDate = rD.BusinessDate
pageDetail.ContractNumber = rD.ContractNumber
pageDetail.Balance = rD.Balance
pageDetail.DealtCcy = rD.DealtCcy
pageDetail.AmountDp = rD.AmountDp
// Automatically generated 21/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 21/11/2021 by matttownsend on silicon.local - END


	t, _ := template.ParseFiles(core.GetTemplateID(dm.AccountLadder_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//AccountLadder_HandlerEdit is the handler used generate the Edit page
func AccountLadder_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.AccountLadder_QueryString)
	_, rD, _ := dao.AccountLadder_GetByID(searchID)
	
	pageDetail := AccountLadder_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.AccountLadder_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
	}

		// 
		// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
pageDetail.SienaReference = rD.SienaReference
pageDetail.BusinessDate = rD.BusinessDate
pageDetail.ContractNumber = rD.ContractNumber
pageDetail.Balance = rD.Balance
pageDetail.DealtCcy = rD.DealtCcy
pageDetail.AmountDp = rD.AmountDp
// Automatically generated 21/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
		//

	// Automatically generated 21/11/2021 by matttownsend on silicon.local - END

	t, _ := template.ParseFiles(core.GetTemplateID(dm.AccountLadder_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//AccountLadder_HandlerSave is the handler used process the saving of an AccountLadder
func AccountLadder_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("SienaReference"))

	var item dm.AccountLadder

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
		item.SienaReference = r.FormValue(dm.AccountLadder_SienaReference)
		item.BusinessDate = r.FormValue(dm.AccountLadder_BusinessDate)
		item.ContractNumber = r.FormValue(dm.AccountLadder_ContractNumber)
		item.Balance = r.FormValue(dm.AccountLadder_Balance)
		item.DealtCcy = r.FormValue(dm.AccountLadder_DealtCcy)
		item.AmountDp = r.FormValue(dm.AccountLadder_AmountDp)
	
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - END

	// Automatically generated 21/11/2021 by matttownsend on silicon.local - END

	dao.AccountLadder_Store(item)	

	http.Redirect(w, r, AccountLadder_Redirect, http.StatusFound)
}

//AccountLadder_HandlerNew is the handler used process the creation of an AccountLadder
func AccountLadder_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := AccountLadder_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.AccountLadder_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",	
	}

		// 
		// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
pageDetail.SienaReference = ""
pageDetail.BusinessDate = ""
pageDetail.ContractNumber = ""
pageDetail.Balance = ""
pageDetail.DealtCcy = ""
pageDetail.AmountDp = ""
// Automatically generated 21/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
		//

	t, _ := template.ParseFiles(core.GetTemplateID(dm.AccountLadder_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//AccountLadder_HandlerDelete is the handler used process the deletion of an AccountLadder
func AccountLadder_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.AccountLadder_QueryString)

	dao.AccountLadder_Delete(searchID)	

	http.Redirect(w, r, AccountLadder_Redirect, http.StatusFound)
}