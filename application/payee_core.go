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
// Date & Time		    : 14/06/2022 at 21:32:07
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
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
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Payee
}
//Payee_Redirect provides a page to return to aftern an action
const (
	Payee_Redirect = dm.Payee_PathList
)

//payee_Page provides the information for the template for an individual Payee
type Payee_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SourceTable         string
	KeyCounterpartyFirm         string
	KeyCounterpartyCentre         string
	KeyCurrency         string
	KeyName         string
	KeyNumber         string
	KeyDirection         string
	KeyType         string
	FullName         string
	Address         string
	PhoneNo         string
	Country         string
	Country_lookup    []dm.Lookup_Item
	Bic         string
	Iban         string
	AccountNo         string
	FedWireNo         string
	SortCode         string
	BankName         string
	BankPinCode         string
	BankAddress         string
	Reason         string
	BankSettlementAcct         string
	UpdatedUserId         string
	Status         string
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



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

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = payee_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Payee_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the Payee Page 
func payee_PopulatePage(rD dm.Payee, pageDetail Payee_Page) Payee_Page {
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Country_lookup = dao.Country_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	