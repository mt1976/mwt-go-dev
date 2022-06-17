package application
// ----------------------------------------------------------------
// Automatically generated  "/application/account.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Account (account)
// Endpoint 	        : Account (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:05
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//account_PageList provides the information for the template for a list of Accounts
type Account_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Account
}
//Account_Redirect provides a page to return to aftern an action
const (
	Account_Redirect = dm.Account_PathList
)

//account_Page provides the information for the template for an individual Account
type Account_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SienaReference         string
	CustomerSienaView         string
	SienaCommonRef         string
	Status         string
	StartDate         string
	MaturityDate         string
	ContractNumber         string
	ExternalReference         string
	CCY         string
	CCY_lookup    []dm.Lookup_Item
	Book         string
	Book_lookup    []dm.Lookup_Item
	MandatedUser         string
	BackOfficeNotes         string
	CashBalance         string
	AccountNumber         string
	AccountName         string
	LedgerBalance         string
	Portfolio         string
	Portfolio_lookup    []dm.Lookup_Item
	AgreementId         string
	BackOfficeRefNo         string
	ISIN         string
	UTI         string
	CCYName         string
	BookName         string
	PortfolioName         string
	Centre         string
	Centre_lookup    []dm.Lookup_Item
	DealTypeKey         string
	DealTypeShortName         string
	InternalId         string
	InternalDeleted         string
	UpdatedTransactionId         string
	UpdatedUserId         string
	UpdatedDateTime         string
	DeletedTransactionId         string
	DeletedUserId         string
	ChangeType         string
	CCYDp         string
	CompID         string
	Firm         string
	Firm_lookup    []dm.Lookup_Item
	DealType         string
	FullDealType         string
	DealingInterface         string
	DealtAmount         string
	ParentContractNumber         string
	InterestFrequency         string
	InterestAction         string
	InterestStrategy         string
	InterestBasis         string
	SienaDealer         string
	DealOwner         string
	OriginUser         string
	EditedByUser         string
	DealOwnerMnemonic         string
	UTCOriginTime         string
	UTCUpdateTime         string
	CustomerStatementNotes         string
	NotesMargin         string
	RequestedBy         string
	EditReason         string
	EditOtherReason         string
	NoticeDays         string
	DebitFrequency         string
	CreditFrequency         string
	EURAmount         string
	EUROtherAmount         string
	PaymentSystemSienaView         string
	PaymentSystemExternalView         string
	DealtCA         string
	AgainstCA         string
	LedgerCA         string
	CashBalanceCA         string
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Account_Publish annouces the endpoints available for this object
func Account_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Account_Path, Account_Handler)
	mux.HandleFunc(dm.Account_PathList, Account_HandlerList)
	mux.HandleFunc(dm.Account_PathView, Account_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Account_Title)
    core.Catalog_Add(dm.Account_Title, dm.Account_Path, "", dm.Account_QueryString, "Application")
}


//Account_HandlerList is the handler for the list page
func Account_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Account
	noItems, returnList, _ := dao.Account_GetList()

	pageDetail := Account_PageList{
		Title:            CardTitle(dm.Account_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Account_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Account_TemplateList, w, r, pageDetail)

}


//Account_HandlerView is the handler used to View a page
func Account_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Account_QueryString)
	_, rD, _ := dao.Account_GetByID(searchID)

	pageDetail := Account_Page{
		Title:       CardTitle(dm.Account_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Account_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = account_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Account_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the Account Page 
func account_PopulatePage(rD dm.Account, pageDetail Account_Page) Account_Page {
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SienaReference = rD.SienaReference
	pageDetail.CustomerSienaView = rD.CustomerSienaView
	pageDetail.SienaCommonRef = rD.SienaCommonRef
	pageDetail.Status = rD.Status
	pageDetail.StartDate = rD.StartDate
	pageDetail.MaturityDate = rD.MaturityDate
	pageDetail.ContractNumber = rD.ContractNumber
	pageDetail.ExternalReference = rD.ExternalReference
	pageDetail.CCY = rD.CCY
	pageDetail.Book = rD.Book
	pageDetail.MandatedUser = rD.MandatedUser
	pageDetail.BackOfficeNotes = rD.BackOfficeNotes
	pageDetail.CashBalance = rD.CashBalance
	pageDetail.AccountNumber = rD.AccountNumber
	pageDetail.AccountName = rD.AccountName
	pageDetail.LedgerBalance = rD.LedgerBalance
	pageDetail.Portfolio = rD.Portfolio
	pageDetail.AgreementId = rD.AgreementId
	pageDetail.BackOfficeRefNo = rD.BackOfficeRefNo
	pageDetail.ISIN = rD.ISIN
	pageDetail.UTI = rD.UTI
	pageDetail.CCYName = rD.CCYName
	pageDetail.BookName = rD.BookName
	pageDetail.PortfolioName = rD.PortfolioName
	pageDetail.Centre = rD.Centre
	pageDetail.DealTypeKey = rD.DealTypeKey
	pageDetail.DealTypeShortName = rD.DealTypeShortName
	pageDetail.InternalId = rD.InternalId
	pageDetail.InternalDeleted = rD.InternalDeleted
	pageDetail.UpdatedTransactionId = rD.UpdatedTransactionId
	pageDetail.UpdatedUserId = rD.UpdatedUserId
	pageDetail.UpdatedDateTime = rD.UpdatedDateTime
	pageDetail.DeletedTransactionId = rD.DeletedTransactionId
	pageDetail.DeletedUserId = rD.DeletedUserId
	pageDetail.ChangeType = rD.ChangeType
	pageDetail.CCYDp = rD.CCYDp
	pageDetail.CompID = rD.CompID
	pageDetail.Firm = rD.Firm
	pageDetail.DealType = rD.DealType
	pageDetail.FullDealType = rD.FullDealType
	pageDetail.DealingInterface = rD.DealingInterface
	pageDetail.DealtAmount = rD.DealtAmount
	pageDetail.ParentContractNumber = rD.ParentContractNumber
	pageDetail.InterestFrequency = rD.InterestFrequency
	pageDetail.InterestAction = rD.InterestAction
	pageDetail.InterestStrategy = rD.InterestStrategy
	pageDetail.InterestBasis = rD.InterestBasis
	pageDetail.SienaDealer = rD.SienaDealer
	pageDetail.DealOwner = rD.DealOwner
	pageDetail.OriginUser = rD.OriginUser
	pageDetail.EditedByUser = rD.EditedByUser
	pageDetail.DealOwnerMnemonic = rD.DealOwnerMnemonic
	pageDetail.UTCOriginTime = rD.UTCOriginTime
	pageDetail.UTCUpdateTime = rD.UTCUpdateTime
	pageDetail.CustomerStatementNotes = rD.CustomerStatementNotes
	pageDetail.NotesMargin = rD.NotesMargin
	pageDetail.RequestedBy = rD.RequestedBy
	pageDetail.EditReason = rD.EditReason
	pageDetail.EditOtherReason = rD.EditOtherReason
	pageDetail.NoticeDays = rD.NoticeDays
	pageDetail.DebitFrequency = rD.DebitFrequency
	pageDetail.CreditFrequency = rD.CreditFrequency
	pageDetail.EURAmount = rD.EURAmount
	pageDetail.EUROtherAmount = rD.EUROtherAmount
	pageDetail.PaymentSystemSienaView = rD.PaymentSystemSienaView
	pageDetail.PaymentSystemExternalView = rD.PaymentSystemExternalView
	
	pageDetail.DealtCA = rD.DealtCA
	pageDetail.AgainstCA = rD.AgainstCA
	pageDetail.LedgerCA = rD.LedgerCA
	pageDetail.CashBalanceCA = rD.CashBalanceCA
	
	//
	// Automatically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.CCY_lookup = dao.Currency_GetLookup()
	
	
	
	pageDetail.Book_lookup = dao.Book_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Portfolio_lookup = dao.Portfolio_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Centre_lookup = dao.Centre_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Firm_lookup = dao.Firm_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	