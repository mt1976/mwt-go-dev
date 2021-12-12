package application
// ----------------------------------------------------------------
// Automatically generated  "/application/payee.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Payee (payee)
// Endpoint 	        : Payee (ID)
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

//payee_PageList provides the information for the template for a list of Payees
type Payee_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Payee
}

//payee_Page provides the information for the template for an individual Payee
type Payee_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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
		Country_Lookup string
		Firm_Lookup string
		Centre_Lookup string
		Currency_Lookup string
		Status_Extra string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	Country_Lookup_List	[]dm.Country
	Firm_Lookup_List	[]dm.Firm
	Centre_Lookup_List	[]dm.Centre
	Currency_Lookup_List	[]dm.Currency
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	Payee_Redirect = dm.Payee_PathList
)

//Payee_Publish annouces the endpoints available for this object
func Payee_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.Payee_PathList, Payee_HandlerList)
	mux.HandleFunc(dm.Payee_PathView, Payee_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Siena", dm.Payee_Title)
    //No API
}

//Payee_HandlerList is the handler for the list page
func Payee_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Payee
	noItems, returnList, _ := dao.Payee_GetList()

	pageDetail := Payee_PageList{
		Title:            CardTitle(dm.Payee_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Payee_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Payee_TemplateList, w, r, pageDetail)

}

//Payee_HandlerView is the handler used to View a page
func Payee_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Payee_QueryString)
	_, rD, _ := dao.Payee_GetByID(searchID)

	pageDetail := Payee_Page{
		Title:       CardTitle(dm.Payee_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Payee_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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

pageDetail.Status_Extra = rD.Status_Extra

// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,Country_Lookup_Name,_:= dao.Country_GetByID(rD.Country)
pageDetail.Country_Lookup = Country_Lookup_Name.Name
_,KeyCounterpartyFirm_Lookup_FullName,_:= dao.Firm_GetByID(rD.KeyCounterpartyFirm)
pageDetail.Firm_Lookup = KeyCounterpartyFirm_Lookup_FullName.FullName
_,KeyCounterpartyCentre_Lookup_Name,_:= dao.Centre_GetByID(rD.KeyCounterpartyCentre)
pageDetail.Centre_Lookup = KeyCounterpartyCentre_Lookup_Name.Name
_,KeyCurrency_Lookup_Name,_:= dao.Currency_GetByID(rD.KeyCurrency)
pageDetail.Currency_Lookup = KeyCurrency_Lookup_Name.Name
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Payee_TemplateView, w, r, pageDetail)

}

//Payee_HandlerEdit is the handler used generate the Edit page
func Payee_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Payee_QueryString)
	_, rD, _ := dao.Payee_GetByID(searchID)
	
	pageDetail := Payee_Page{
		Title:       CardTitle(dm.Payee_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Payee_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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

pageDetail.Status_Extra = rD.Status_Extra

// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,Country_Lookup_Name,_:= dao.Country_GetByID(rD.Country)
pageDetail.Country_Lookup = Country_Lookup_Name.Name
_,pageDetail.Country_Lookup_List,_ = dao.Country_GetList()
_,KeyCounterpartyFirm_Lookup_FullName,_:= dao.Firm_GetByID(rD.KeyCounterpartyFirm)
pageDetail.Firm_Lookup = KeyCounterpartyFirm_Lookup_FullName.FullName
_,pageDetail.Firm_Lookup_List,_ = dao.Firm_GetList()
_,KeyCounterpartyCentre_Lookup_Name,_:= dao.Centre_GetByID(rD.KeyCounterpartyCentre)
pageDetail.Centre_Lookup = KeyCounterpartyCentre_Lookup_Name.Name
_,pageDetail.Centre_Lookup_List,_ = dao.Centre_GetList()
_,KeyCurrency_Lookup_Name,_:= dao.Currency_GetByID(rD.KeyCurrency)
pageDetail.Currency_Lookup = KeyCurrency_Lookup_Name.Name
_,pageDetail.Currency_Lookup_List,_ = dao.Currency_GetList()
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Payee_TemplateEdit, w, r, pageDetail)


}

//Payee_HandlerSave is the handler used process the saving of an Payee
func Payee_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("KeyCounterpartyFirm"))

	var item dm.Payee
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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
		item.Country_Lookup = r.FormValue(dm.Payee_Country_Lookup)
		item.Firm_Lookup = r.FormValue(dm.Payee_Firm_Lookup)
		item.Centre_Lookup = r.FormValue(dm.Payee_Centre_Lookup)
		item.Currency_Lookup = r.FormValue(dm.Payee_Currency_Lookup)
		item.Status_Extra = r.FormValue(dm.Payee_Status_Extra)
		
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.Payee_Store(item,r)	

	http.Redirect(w, r, Payee_Redirect, http.StatusFound)
}

//Payee_HandlerNew is the handler used process the creation of an Payee
func Payee_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Payee_Page{
		Title:       CardTitle(dm.Payee_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Payee_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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

pageDetail.Status_Extra = ""

// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
pageDetail.Country_Lookup = ""
_,pageDetail.Country_Lookup_List,_ = dao.Country_GetList()
pageDetail.Firm_Lookup = ""
_,pageDetail.Firm_Lookup_List,_ = dao.Firm_GetList()
pageDetail.Centre_Lookup = ""
_,pageDetail.Centre_Lookup_List,_ = dao.Centre_GetList()
pageDetail.Currency_Lookup = ""
_,pageDetail.Currency_Lookup_List,_ = dao.Currency_GetList()
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Payee_TemplateNew, w, r, pageDetail)

}

//Payee_HandlerDelete is the handler used process the deletion of an Payee
func Payee_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Payee_QueryString)

	dao.Payee_Delete(searchID)	

	http.Redirect(w, r, Payee_Redirect, http.StatusFound)
}