package application
// ----------------------------------------------------------------
// Automatically generated  "/application/payee.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Payee (payee)
// Endpoint 	        : Payee (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 22/11/2021 at 11:25:56
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

//payee_PageList provides the information for the template for a list of Payees
type Payee_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Payee
}

//payee_Page provides the information for the template for an individual Payee
type Payee_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
		SourceTable string
		KeyCounterpartyFirm string
		KeyCounterpartyCentre string
		KeyCurrency string
		KeyName string
		KeyNumber string
		KeyDirection string
		KeyType string
		FullName string
		Address string
		PhoneNo string
		Country string
		Bic string
		Iban string
		AccountNo string
		FedWireNo string
		SortCode string
		BankName string
		BankPinCode string
		BankAddress string
		Reason string
		BankSettlementAcct string
		UpdatedUserId string
		Country_Impl string
		Firm_Impl string
		Centre_Impl string
		Currency_Impl string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	Country_Impl_List	[]dm.Country
	Firm_Impl_List	[]dm.Firm
	Centre_Impl_List	[]dm.Centre
	Currency_Impl_List	[]dm.Currency
	
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
}

const (
	Payee_Redirect = dm.Payee_PathList
)

//Payee_Publish annouces the endpoints available for this object
func Payee_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Payee_PathList, Payee_HandlerList)
	mux.HandleFunc(dm.Payee_PathView, Payee_HandlerView)
	
	
	
	
	logs.Publish("Siena", dm.Payee_Title)
	
	// payee_PublishImpl should be specified in application/payee_Impl.go
	// to provide the implementation for the special case.
	// override function should be defined as
	// payee_PublishImpl(mux http.ServeMux) {...}
	// TODO - this is a temporary hack to get the special case working
	// Add to main.go >>> payee_PublishImpl(mux)
	
	
}

//Payee_HandlerList is the handler for the list page
func Payee_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Payee
	noItems, returnList, _ := dao.Payee_GetList()


	pageDetail := Payee_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.Payee_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Payee_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//Payee_HandlerView is the handler used to View a page
func Payee_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Payee_QueryString)
	_, rD, _ := dao.Payee_GetByID(searchID)

	pageDetail := Payee_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Payee_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
	}

		// 
		// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
pageDetail.SourceTable = rD.SourceTable
pageDetail.KeyCounterpartyFirm = rD.KeyCounterpartyFirm
pageDetail.KeyCounterpartyCentre = rD.KeyCounterpartyCentre
pageDetail.KeyCurrency = rD.KeyCurrency
pageDetail.KeyName = rD.KeyName
pageDetail.KeyNumber = rD.KeyNumber
pageDetail.KeyDirection = rD.KeyDirection
pageDetail.KeyType = rD.KeyType
pageDetail.FullName = rD.FullName
pageDetail.Address = rD.Address
pageDetail.PhoneNo = rD.PhoneNo
pageDetail.Country = rD.Country
pageDetail.Bic = rD.Bic
pageDetail.Iban = rD.Iban
pageDetail.AccountNo = rD.AccountNo
pageDetail.FedWireNo = rD.FedWireNo
pageDetail.SortCode = rD.SortCode
pageDetail.BankName = rD.BankName
pageDetail.BankPinCode = rD.BankPinCode
pageDetail.BankAddress = rD.BankAddress
pageDetail.Reason = rD.Reason
pageDetail.BankSettlementAcct = rD.BankSettlementAcct
pageDetail.UpdatedUserId = rD.UpdatedUserId
// Automatically generated 22/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,Country_Lookup,_:= dao.Country_GetByID(rD.Country)
pageDetail.Country_Impl = Country_Lookup.Name
_,KeyCounterpartyFirm_Lookup,_:= dao.Firm_GetByID(rD.KeyCounterpartyFirm)
pageDetail.Firm_Impl = KeyCounterpartyFirm_Lookup.FullName
_,KeyCounterpartyCentre_Lookup,_:= dao.Centre_GetByID(rD.KeyCounterpartyCentre)
pageDetail.Centre_Impl = KeyCounterpartyCentre_Lookup.Name
_,KeyCurrency_Lookup,_:= dao.Currency_GetByID(rD.KeyCurrency)
pageDetail.Currency_Impl = KeyCurrency_Lookup.Name
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//


	// payee_HandlerViewImpl should be specified in application/payee_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// func payee_HandlerViewImpl(pageDetail Payee_Page) Payee_Page {return pageDetail}
	pageDetail = payee_HandlerViewImpl(pageDetail)

	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END


	t, _ := template.ParseFiles(core.GetTemplateID(dm.Payee_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Payee_HandlerEdit is the handler used generate the Edit page
func Payee_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Payee_QueryString)
	_, rD, _ := dao.Payee_GetByID(searchID)
	
	pageDetail := Payee_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Payee_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
	}

		// 
		// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
pageDetail.SourceTable = rD.SourceTable
pageDetail.KeyCounterpartyFirm = rD.KeyCounterpartyFirm
pageDetail.KeyCounterpartyCentre = rD.KeyCounterpartyCentre
pageDetail.KeyCurrency = rD.KeyCurrency
pageDetail.KeyName = rD.KeyName
pageDetail.KeyNumber = rD.KeyNumber
pageDetail.KeyDirection = rD.KeyDirection
pageDetail.KeyType = rD.KeyType
pageDetail.FullName = rD.FullName
pageDetail.Address = rD.Address
pageDetail.PhoneNo = rD.PhoneNo
pageDetail.Country = rD.Country
pageDetail.Bic = rD.Bic
pageDetail.Iban = rD.Iban
pageDetail.AccountNo = rD.AccountNo
pageDetail.FedWireNo = rD.FedWireNo
pageDetail.SortCode = rD.SortCode
pageDetail.BankName = rD.BankName
pageDetail.BankPinCode = rD.BankPinCode
pageDetail.BankAddress = rD.BankAddress
pageDetail.Reason = rD.Reason
pageDetail.BankSettlementAcct = rD.BankSettlementAcct
pageDetail.UpdatedUserId = rD.UpdatedUserId
// Automatically generated 22/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,Country_Lookup,_:= dao.Country_GetByID(rD.Country)
pageDetail.Country_Impl = Country_Lookup.Name
_,pageDetail.Country_Impl_List,_ = dao.Country_GetList()
_,KeyCounterpartyFirm_Lookup,_:= dao.Firm_GetByID(rD.KeyCounterpartyFirm)
pageDetail.Firm_Impl = KeyCounterpartyFirm_Lookup.FullName
_,pageDetail.Firm_Impl_List,_ = dao.Firm_GetList()
_,KeyCounterpartyCentre_Lookup,_:= dao.Centre_GetByID(rD.KeyCounterpartyCentre)
pageDetail.Centre_Impl = KeyCounterpartyCentre_Lookup.Name
_,pageDetail.Centre_Impl_List,_ = dao.Centre_GetList()
_,KeyCurrency_Lookup,_:= dao.Currency_GetByID(rD.KeyCurrency)
pageDetail.Currency_Impl = KeyCurrency_Lookup.Name
_,pageDetail.Currency_Impl_List,_ = dao.Currency_GetList()
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//

	// payee_HandlerEditImpl should be specified in application/payee_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// func payee_HandlerEditImpl(pageDetail Payee_Page) Payee_Page {return pageDetail}
	pageDetail = payee_HandlerEditImpl(pageDetail)

	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Payee_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Payee_HandlerSave is the handler used process the saving of an Payee
func Payee_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("KeyCounterpartyFirm"))

	var item dm.Payee

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
		item.SourceTable = r.FormValue(dm.Payee_SourceTable)
		item.KeyCounterpartyFirm = r.FormValue(dm.Payee_KeyCounterpartyFirm)
		item.KeyCounterpartyCentre = r.FormValue(dm.Payee_KeyCounterpartyCentre)
		item.KeyCurrency = r.FormValue(dm.Payee_KeyCurrency)
		item.KeyName = r.FormValue(dm.Payee_KeyName)
		item.KeyNumber = r.FormValue(dm.Payee_KeyNumber)
		item.KeyDirection = r.FormValue(dm.Payee_KeyDirection)
		item.KeyType = r.FormValue(dm.Payee_KeyType)
		item.FullName = r.FormValue(dm.Payee_FullName)
		item.Address = r.FormValue(dm.Payee_Address)
		item.PhoneNo = r.FormValue(dm.Payee_PhoneNo)
		item.Country = r.FormValue(dm.Payee_Country)
		item.Bic = r.FormValue(dm.Payee_Bic)
		item.Iban = r.FormValue(dm.Payee_Iban)
		item.AccountNo = r.FormValue(dm.Payee_AccountNo)
		item.FedWireNo = r.FormValue(dm.Payee_FedWireNo)
		item.SortCode = r.FormValue(dm.Payee_SortCode)
		item.BankName = r.FormValue(dm.Payee_BankName)
		item.BankPinCode = r.FormValue(dm.Payee_BankPinCode)
		item.BankAddress = r.FormValue(dm.Payee_BankAddress)
		item.Reason = r.FormValue(dm.Payee_Reason)
		item.BankSettlementAcct = r.FormValue(dm.Payee_BankSettlementAcct)
		item.UpdatedUserId = r.FormValue(dm.Payee_UpdatedUserId)
		item.Country_Impl = r.FormValue(dm.Payee_Country_Impl)
		item.Firm_Impl = r.FormValue(dm.Payee_Firm_Impl)
		item.Centre_Impl = r.FormValue(dm.Payee_Centre_Impl)
		item.Currency_Impl = r.FormValue(dm.Payee_Currency_Impl)
	
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END

	// payee_HandlerSaveImpl should be specified in application/payee_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// func payee_HandlerSaveImpl(item dm.Payee) dm.Payee {return item}
	item = payee_HandlerSaveImpl(item)

	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END

	dao.Payee_Store(item)	

	http.Redirect(w, r, Payee_Redirect, http.StatusFound)
}

//Payee_HandlerNew is the handler used process the creation of an Payee
func Payee_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Payee_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Payee_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",	
	}

		// 
		// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
pageDetail.SourceTable = ""
pageDetail.KeyCounterpartyFirm = ""
pageDetail.KeyCounterpartyCentre = ""
pageDetail.KeyCurrency = ""
pageDetail.KeyName = ""
pageDetail.KeyNumber = ""
pageDetail.KeyDirection = ""
pageDetail.KeyType = ""
pageDetail.FullName = ""
pageDetail.Address = ""
pageDetail.PhoneNo = ""
pageDetail.Country = ""
pageDetail.Bic = ""
pageDetail.Iban = ""
pageDetail.AccountNo = ""
pageDetail.FedWireNo = ""
pageDetail.SortCode = ""
pageDetail.BankName = ""
pageDetail.BankPinCode = ""
pageDetail.BankAddress = ""
pageDetail.Reason = ""
pageDetail.BankSettlementAcct = ""
pageDetail.UpdatedUserId = ""
// Automatically generated 22/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
pageDetail.Country_Impl = ""
_,pageDetail.Country_Impl_List,_ = dao.Country_GetList()
pageDetail.Firm_Impl = ""
_,pageDetail.Firm_Impl_List,_ = dao.Firm_GetList()
pageDetail.Centre_Impl = ""
_,pageDetail.Centre_Impl_List,_ = dao.Centre_GetList()
pageDetail.Currency_Impl = ""
_,pageDetail.Currency_Impl_List,_ = dao.Currency_GetList()
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Payee_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Payee_HandlerDelete is the handler used process the deletion of an Payee
func Payee_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Payee_QueryString)

	dao.Payee_Delete(searchID)	

	http.Redirect(w, r, Payee_Redirect, http.StatusFound)
}