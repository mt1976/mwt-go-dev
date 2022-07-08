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
// Date & Time		    : 28/06/2022 at 16:10:55
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
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
	logs.Publish("Application", dm.Payee_Title)
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

	pageDetail := dm.Payee_PageList{
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

	pageDetail := dm.Payee_Page{
		Title:       CardTitle(dm.Payee_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Payee_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = payee_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Payee_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the Payee Page 
func payee_PopulatePage(rD dm.Payee, pageDetail dm.Payee_Page) dm.Payee_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	pageDetail.Status = rD.Status
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Country_lookup = dao.Country_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SourceTable_props = rD.SourceTable_props
	pageDetail.KeyCounterpartyFirm_props = rD.KeyCounterpartyFirm_props
	pageDetail.KeyCounterpartyCentre_props = rD.KeyCounterpartyCentre_props
	pageDetail.KeyCurrency_props = rD.KeyCurrency_props
	pageDetail.KeyName_props = rD.KeyName_props
	pageDetail.KeyNumber_props = rD.KeyNumber_props
	pageDetail.KeyDirection_props = rD.KeyDirection_props
	pageDetail.KeyType_props = rD.KeyType_props
	pageDetail.FullName_props = rD.FullName_props
	pageDetail.Address_props = rD.Address_props
	pageDetail.PhoneNo_props = rD.PhoneNo_props
	pageDetail.Country_props = rD.Country_props
	pageDetail.Bic_props = rD.Bic_props
	pageDetail.Iban_props = rD.Iban_props
	pageDetail.AccountNo_props = rD.AccountNo_props
	pageDetail.FedWireNo_props = rD.FedWireNo_props
	pageDetail.SortCode_props = rD.SortCode_props
	pageDetail.BankName_props = rD.BankName_props
	pageDetail.BankPinCode_props = rD.BankPinCode_props
	pageDetail.BankAddress_props = rD.BankAddress_props
	pageDetail.Reason_props = rD.Reason_props
	pageDetail.BankSettlementAcct_props = rD.BankSettlementAcct_props
	pageDetail.UpdatedUserId_props = rD.UpdatedUserId_props
	pageDetail.Status_props = rD.Status_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	