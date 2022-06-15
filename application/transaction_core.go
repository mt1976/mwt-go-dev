package application
// ----------------------------------------------------------------
// Automatically generated  "/application/transaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Transaction (transaction)
// Endpoint 	        : Transaction (Ref)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:31:46
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//transaction_PageList provides the information for the template for a list of Transactions
type Transaction_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Transaction
}
//Transaction_Redirect provides a page to return to aftern an action
const (
	Transaction_Redirect = dm.Transaction_PathList
)

//transaction_Page provides the information for the template for an individual Transaction
type Transaction_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SienaReference         string
	Status         string
	ValueDate         string
	MaturityDate         string
	ContractNumber         string
	ExternalReference         string
	Book         string
	MandatedUser         string
	Portfolio         string
	AgreementId         string
	BackOfficeRefNo         string
	ISIN         string
	UTI         string
	BookName         string
	Centre         string
	Firm         string
	DealTypeShortName         string
	FullDealType         string
	TradeDate         string
	DealtCcy         string
	DealtAmount         string
	AgainstAmount         string
	AgainstCcy         string
	AllInRate         string
	MktRate         string
	SettleCcy         string
	Direction         string
	NpvRate         string
	OriginUser         string
	PayInstruction         string
	ReceiptInstruction         string
	NIName         string
	CCYPair         string
	Instrument         string
	PortfolioName         string
	RVDate         string
	RVMTM         string
	CounterBook         string
	CounterBookName         string
	Party         string
	PartyName         string
	NameCentre         string
	NameFirm         string
	CustomerExternalView         string
	CustomerSienaView         string
	CompID         string
	SienaDealer         string
	DealOwner         string
	DealOwnerMnemonic         string
	EditedByUser         string
	UTCOriginTime         string
	UTCUpdateTime         string
	MarginTrading         string
	SwapPoints         string
	SpotDate         string
	SpotRate         string
	MktSpotRate         string
	SpotSalesMargin         string
	SpotChlMargin         string
	RerouteCcy         string
	CustomerPayInstruction         string
	CustomerReceiptInstruction         string
	BackOfficeNotes         string
	CustomerStatementNotes         string
	NotesMargin         string
	RequestedBy         string
	EditReason         string
	EditOtherReason         string
	NICleanPrice         string
	NIDirtyPrice         string
	NIYield         string
	NIClearingSystem         string
	Acceptor         string
	NIDiscount         string
	FastPay         string
	PaymentFee         string
	PaymentFeePolicy         string
	PaymentReason         string
	PaymentDate         string
	SettlementDate         string
	FixingDate         string
	VenueUTI         string
	EditVersion         string
	BrokeragePercentage         string
	BrokerageAmount         string
	BrokerageCurrency         string
	BrokerageDate         string
	AccountName         string
	AccountNumber         string
	CashBalance         string
	DebitFrequency         string
	CreditFrequency         string
	ManuallyQuoted         string
	LedgerBalance         string
	SettAmtOutstanding         string
	FeePercentage         string
	FeeAmount         string
	Venue         string
	EURAmount         string
	EUROtherAmount         string
	LEI         string
	Equity         string
	Shares         string
	QuoteExpiryDate         string
	Commodity         string
	PaymentSystemSienaView         string
	PaymentSystemExternalView         string
	SalesProfit         string
	RejectReason         string
	PaymentError         string
	RepoPrincipal         string
	FixingFrequency         string
	Dealt         string
	Against         string
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Transaction_Publish annouces the endpoints available for this object
func Transaction_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.Transaction_PathList, Transaction_HandlerList)
	mux.HandleFunc(dm.Transaction_PathView, Transaction_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Transaction_Title)
    //No API
}


//Transaction_HandlerList is the handler for the list page
func Transaction_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Transaction
	noItems, returnList, _ := dao.Transaction_GetList()

	pageDetail := Transaction_PageList{
		Title:            CardTitle(dm.Transaction_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Transaction_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Transaction_TemplateList, w, r, pageDetail)

}


//Transaction_HandlerView is the handler used to View a page
func Transaction_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Transaction_QueryString)
	_, rD, _ := dao.Transaction_GetByID(searchID)

	pageDetail := Transaction_Page{
		Title:       CardTitle(dm.Transaction_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Transaction_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = transaction_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Transaction_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the Transaction Page 
func transaction_PopulatePage(rD dm.Transaction, pageDetail Transaction_Page) Transaction_Page {
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SienaReference = rD.SienaReference
	pageDetail.Status = rD.Status
	pageDetail.ValueDate = rD.ValueDate
	pageDetail.MaturityDate = rD.MaturityDate
	pageDetail.ContractNumber = rD.ContractNumber
	pageDetail.ExternalReference = rD.ExternalReference
	pageDetail.Book = rD.Book
	pageDetail.MandatedUser = rD.MandatedUser
	pageDetail.Portfolio = rD.Portfolio
	pageDetail.AgreementId = rD.AgreementId
	pageDetail.BackOfficeRefNo = rD.BackOfficeRefNo
	pageDetail.ISIN = rD.ISIN
	pageDetail.UTI = rD.UTI
	pageDetail.BookName = rD.BookName
	pageDetail.Centre = rD.Centre
	pageDetail.Firm = rD.Firm
	pageDetail.DealTypeShortName = rD.DealTypeShortName
	pageDetail.FullDealType = rD.FullDealType
	pageDetail.TradeDate = rD.TradeDate
	pageDetail.DealtCcy = rD.DealtCcy
	pageDetail.DealtAmount = rD.DealtAmount
	pageDetail.AgainstAmount = rD.AgainstAmount
	pageDetail.AgainstCcy = rD.AgainstCcy
	pageDetail.AllInRate = rD.AllInRate
	pageDetail.MktRate = rD.MktRate
	pageDetail.SettleCcy = rD.SettleCcy
	pageDetail.Direction = rD.Direction
	pageDetail.NpvRate = rD.NpvRate
	pageDetail.OriginUser = rD.OriginUser
	pageDetail.PayInstruction = rD.PayInstruction
	pageDetail.ReceiptInstruction = rD.ReceiptInstruction
	pageDetail.NIName = rD.NIName
	pageDetail.CCYPair = rD.CCYPair
	pageDetail.Instrument = rD.Instrument
	pageDetail.PortfolioName = rD.PortfolioName
	pageDetail.RVDate = rD.RVDate
	pageDetail.RVMTM = rD.RVMTM
	pageDetail.CounterBook = rD.CounterBook
	pageDetail.CounterBookName = rD.CounterBookName
	pageDetail.Party = rD.Party
	pageDetail.PartyName = rD.PartyName
	pageDetail.NameCentre = rD.NameCentre
	pageDetail.NameFirm = rD.NameFirm
	pageDetail.CustomerExternalView = rD.CustomerExternalView
	pageDetail.CustomerSienaView = rD.CustomerSienaView
	pageDetail.CompID = rD.CompID
	pageDetail.SienaDealer = rD.SienaDealer
	pageDetail.DealOwner = rD.DealOwner
	pageDetail.DealOwnerMnemonic = rD.DealOwnerMnemonic
	pageDetail.EditedByUser = rD.EditedByUser
	pageDetail.UTCOriginTime = rD.UTCOriginTime
	pageDetail.UTCUpdateTime = rD.UTCUpdateTime
	pageDetail.MarginTrading = rD.MarginTrading
	pageDetail.SwapPoints = rD.SwapPoints
	pageDetail.SpotDate = rD.SpotDate
	pageDetail.SpotRate = rD.SpotRate
	pageDetail.MktSpotRate = rD.MktSpotRate
	pageDetail.SpotSalesMargin = rD.SpotSalesMargin
	pageDetail.SpotChlMargin = rD.SpotChlMargin
	pageDetail.RerouteCcy = rD.RerouteCcy
	pageDetail.CustomerPayInstruction = rD.CustomerPayInstruction
	pageDetail.CustomerReceiptInstruction = rD.CustomerReceiptInstruction
	pageDetail.BackOfficeNotes = rD.BackOfficeNotes
	pageDetail.CustomerStatementNotes = rD.CustomerStatementNotes
	pageDetail.NotesMargin = rD.NotesMargin
	pageDetail.RequestedBy = rD.RequestedBy
	pageDetail.EditReason = rD.EditReason
	pageDetail.EditOtherReason = rD.EditOtherReason
	pageDetail.NICleanPrice = rD.NICleanPrice
	pageDetail.NIDirtyPrice = rD.NIDirtyPrice
	pageDetail.NIYield = rD.NIYield
	pageDetail.NIClearingSystem = rD.NIClearingSystem
	pageDetail.Acceptor = rD.Acceptor
	pageDetail.NIDiscount = rD.NIDiscount
	pageDetail.FastPay = rD.FastPay
	pageDetail.PaymentFee = rD.PaymentFee
	pageDetail.PaymentFeePolicy = rD.PaymentFeePolicy
	pageDetail.PaymentReason = rD.PaymentReason
	pageDetail.PaymentDate = rD.PaymentDate
	pageDetail.SettlementDate = rD.SettlementDate
	pageDetail.FixingDate = rD.FixingDate
	pageDetail.VenueUTI = rD.VenueUTI
	pageDetail.EditVersion = rD.EditVersion
	pageDetail.BrokeragePercentage = rD.BrokeragePercentage
	pageDetail.BrokerageAmount = rD.BrokerageAmount
	pageDetail.BrokerageCurrency = rD.BrokerageCurrency
	pageDetail.BrokerageDate = rD.BrokerageDate
	pageDetail.AccountName = rD.AccountName
	pageDetail.AccountNumber = rD.AccountNumber
	pageDetail.CashBalance = rD.CashBalance
	pageDetail.DebitFrequency = rD.DebitFrequency
	pageDetail.CreditFrequency = rD.CreditFrequency
	pageDetail.ManuallyQuoted = rD.ManuallyQuoted
	pageDetail.LedgerBalance = rD.LedgerBalance
	pageDetail.SettAmtOutstanding = rD.SettAmtOutstanding
	pageDetail.FeePercentage = rD.FeePercentage
	pageDetail.FeeAmount = rD.FeeAmount
	pageDetail.Venue = rD.Venue
	pageDetail.EURAmount = rD.EURAmount
	pageDetail.EUROtherAmount = rD.EUROtherAmount
	pageDetail.LEI = rD.LEI
	pageDetail.Equity = rD.Equity
	pageDetail.Shares = rD.Shares
	pageDetail.QuoteExpiryDate = rD.QuoteExpiryDate
	pageDetail.Commodity = rD.Commodity
	pageDetail.PaymentSystemSienaView = rD.PaymentSystemSienaView
	pageDetail.PaymentSystemExternalView = rD.PaymentSystemExternalView
	pageDetail.SalesProfit = rD.SalesProfit
	pageDetail.RejectReason = rD.RejectReason
	pageDetail.PaymentError = rD.PaymentError
	pageDetail.RepoPrincipal = rD.RepoPrincipal
	pageDetail.FixingFrequency = rD.FixingFrequency
	
	pageDetail.Dealt = rD.Dealt
	pageDetail.Against = rD.Against
	
	//
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	