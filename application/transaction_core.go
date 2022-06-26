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
// Date & Time		    : 26/06/2022 at 18:48:11
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
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SienaReference         string
	SienaReference_props     dm.FieldProperties
	Status         string
	Status_props     dm.FieldProperties
	ValueDate         string
	ValueDate_props     dm.FieldProperties
	MaturityDate         string
	MaturityDate_props     dm.FieldProperties
	ContractNumber         string
	ContractNumber_props     dm.FieldProperties
	ExternalReference         string
	ExternalReference_props     dm.FieldProperties
	Book         string
	Book_props     dm.FieldProperties
	MandatedUser         string
	MandatedUser_props     dm.FieldProperties
	Portfolio         string
	Portfolio_props     dm.FieldProperties
	AgreementId         string
	AgreementId_props     dm.FieldProperties
	BackOfficeRefNo         string
	BackOfficeRefNo_props     dm.FieldProperties
	ISIN         string
	ISIN_props     dm.FieldProperties
	UTI         string
	UTI_props     dm.FieldProperties
	BookName         string
	BookName_props     dm.FieldProperties
	Centre         string
	Centre_props     dm.FieldProperties
	Firm         string
	Firm_props     dm.FieldProperties
	DealTypeShortName         string
	DealTypeShortName_props     dm.FieldProperties
	FullDealType         string
	FullDealType_props     dm.FieldProperties
	TradeDate         string
	TradeDate_props     dm.FieldProperties
	DealtCcy         string
	DealtCcy_props     dm.FieldProperties
	DealtAmount         string
	DealtAmount_props     dm.FieldProperties
	AgainstAmount         string
	AgainstAmount_props     dm.FieldProperties
	AgainstCcy         string
	AgainstCcy_props     dm.FieldProperties
	AllInRate         string
	AllInRate_props     dm.FieldProperties
	MktRate         string
	MktRate_props     dm.FieldProperties
	SettleCcy         string
	SettleCcy_props     dm.FieldProperties
	Direction         string
	Direction_props     dm.FieldProperties
	NpvRate         string
	NpvRate_props     dm.FieldProperties
	OriginUser         string
	OriginUser_props     dm.FieldProperties
	PayInstruction         string
	PayInstruction_props     dm.FieldProperties
	ReceiptInstruction         string
	ReceiptInstruction_props     dm.FieldProperties
	NIName         string
	NIName_props     dm.FieldProperties
	CCYPair         string
	CCYPair_props     dm.FieldProperties
	Instrument         string
	Instrument_props     dm.FieldProperties
	PortfolioName         string
	PortfolioName_props     dm.FieldProperties
	RVDate         string
	RVDate_props     dm.FieldProperties
	RVMTM         string
	RVMTM_props     dm.FieldProperties
	CounterBook         string
	CounterBook_props     dm.FieldProperties
	CounterBookName         string
	CounterBookName_props     dm.FieldProperties
	Party         string
	Party_props     dm.FieldProperties
	PartyName         string
	PartyName_props     dm.FieldProperties
	NameCentre         string
	NameCentre_props     dm.FieldProperties
	NameFirm         string
	NameFirm_props     dm.FieldProperties
	CustomerExternalView         string
	CustomerExternalView_props     dm.FieldProperties
	CustomerSienaView         string
	CustomerSienaView_props     dm.FieldProperties
	CompID         string
	CompID_props     dm.FieldProperties
	SienaDealer         string
	SienaDealer_props     dm.FieldProperties
	DealOwner         string
	DealOwner_props     dm.FieldProperties
	DealOwnerMnemonic         string
	DealOwnerMnemonic_props     dm.FieldProperties
	EditedByUser         string
	EditedByUser_props     dm.FieldProperties
	UTCOriginTime         string
	UTCOriginTime_props     dm.FieldProperties
	UTCUpdateTime         string
	UTCUpdateTime_props     dm.FieldProperties
	MarginTrading         string
	MarginTrading_props     dm.FieldProperties
	SwapPoints         string
	SwapPoints_props     dm.FieldProperties
	SpotDate         string
	SpotDate_props     dm.FieldProperties
	SpotRate         string
	SpotRate_props     dm.FieldProperties
	MktSpotRate         string
	MktSpotRate_props     dm.FieldProperties
	SpotSalesMargin         string
	SpotSalesMargin_props     dm.FieldProperties
	SpotChlMargin         string
	SpotChlMargin_props     dm.FieldProperties
	RerouteCcy         string
	RerouteCcy_props     dm.FieldProperties
	CustomerPayInstruction         string
	CustomerPayInstruction_props     dm.FieldProperties
	CustomerReceiptInstruction         string
	CustomerReceiptInstruction_props     dm.FieldProperties
	BackOfficeNotes         string
	BackOfficeNotes_props     dm.FieldProperties
	CustomerStatementNotes         string
	CustomerStatementNotes_props     dm.FieldProperties
	NotesMargin         string
	NotesMargin_props     dm.FieldProperties
	RequestedBy         string
	RequestedBy_props     dm.FieldProperties
	EditReason         string
	EditReason_props     dm.FieldProperties
	EditOtherReason         string
	EditOtherReason_props     dm.FieldProperties
	NICleanPrice         string
	NICleanPrice_props     dm.FieldProperties
	NIDirtyPrice         string
	NIDirtyPrice_props     dm.FieldProperties
	NIYield         string
	NIYield_props     dm.FieldProperties
	NIClearingSystem         string
	NIClearingSystem_props     dm.FieldProperties
	Acceptor         string
	Acceptor_props     dm.FieldProperties
	NIDiscount         string
	NIDiscount_props     dm.FieldProperties
	FastPay         string
	FastPay_props     dm.FieldProperties
	PaymentFee         string
	PaymentFee_props     dm.FieldProperties
	PaymentFeePolicy         string
	PaymentFeePolicy_props     dm.FieldProperties
	PaymentReason         string
	PaymentReason_props     dm.FieldProperties
	PaymentDate         string
	PaymentDate_props     dm.FieldProperties
	SettlementDate         string
	SettlementDate_props     dm.FieldProperties
	FixingDate         string
	FixingDate_props     dm.FieldProperties
	VenueUTI         string
	VenueUTI_props     dm.FieldProperties
	EditVersion         string
	EditVersion_props     dm.FieldProperties
	BrokeragePercentage         string
	BrokeragePercentage_props     dm.FieldProperties
	BrokerageAmount         string
	BrokerageAmount_props     dm.FieldProperties
	BrokerageCurrency         string
	BrokerageCurrency_props     dm.FieldProperties
	BrokerageDate         string
	BrokerageDate_props     dm.FieldProperties
	AccountName         string
	AccountName_props     dm.FieldProperties
	AccountNumber         string
	AccountNumber_props     dm.FieldProperties
	CashBalance         string
	CashBalance_props     dm.FieldProperties
	DebitFrequency         string
	DebitFrequency_props     dm.FieldProperties
	CreditFrequency         string
	CreditFrequency_props     dm.FieldProperties
	ManuallyQuoted         string
	ManuallyQuoted_props     dm.FieldProperties
	LedgerBalance         string
	LedgerBalance_props     dm.FieldProperties
	SettAmtOutstanding         string
	SettAmtOutstanding_props     dm.FieldProperties
	FeePercentage         string
	FeePercentage_props     dm.FieldProperties
	FeeAmount         string
	FeeAmount_props     dm.FieldProperties
	Venue         string
	Venue_props     dm.FieldProperties
	EURAmount         string
	EURAmount_props     dm.FieldProperties
	EUROtherAmount         string
	EUROtherAmount_props     dm.FieldProperties
	LEI         string
	LEI_props     dm.FieldProperties
	Equity         string
	Equity_props     dm.FieldProperties
	Shares         string
	Shares_props     dm.FieldProperties
	QuoteExpiryDate         string
	QuoteExpiryDate_props     dm.FieldProperties
	Commodity         string
	Commodity_props     dm.FieldProperties
	PaymentSystemSienaView         string
	PaymentSystemSienaView_props     dm.FieldProperties
	PaymentSystemExternalView         string
	PaymentSystemExternalView_props     dm.FieldProperties
	SalesProfit         string
	SalesProfit_props     dm.FieldProperties
	RejectReason         string
	RejectReason_props     dm.FieldProperties
	PaymentError         string
	PaymentError_props     dm.FieldProperties
	RepoPrincipal         string
	RepoPrincipal_props     dm.FieldProperties
	FixingFrequency         string
	FixingFrequency_props     dm.FieldProperties
	Dealt         string
	Dealt_props     dm.FieldProperties
	Against         string
	Against_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SienaReference_props = rD.SienaReference_props
	pageDetail.Status_props = rD.Status_props
	pageDetail.ValueDate_props = rD.ValueDate_props
	pageDetail.MaturityDate_props = rD.MaturityDate_props
	pageDetail.ContractNumber_props = rD.ContractNumber_props
	pageDetail.ExternalReference_props = rD.ExternalReference_props
	pageDetail.Book_props = rD.Book_props
	pageDetail.MandatedUser_props = rD.MandatedUser_props
	pageDetail.Portfolio_props = rD.Portfolio_props
	pageDetail.AgreementId_props = rD.AgreementId_props
	pageDetail.BackOfficeRefNo_props = rD.BackOfficeRefNo_props
	pageDetail.ISIN_props = rD.ISIN_props
	pageDetail.UTI_props = rD.UTI_props
	pageDetail.BookName_props = rD.BookName_props
	pageDetail.Centre_props = rD.Centre_props
	pageDetail.Firm_props = rD.Firm_props
	pageDetail.DealTypeShortName_props = rD.DealTypeShortName_props
	pageDetail.FullDealType_props = rD.FullDealType_props
	pageDetail.TradeDate_props = rD.TradeDate_props
	pageDetail.DealtCcy_props = rD.DealtCcy_props
	pageDetail.DealtAmount_props = rD.DealtAmount_props
	pageDetail.AgainstAmount_props = rD.AgainstAmount_props
	pageDetail.AgainstCcy_props = rD.AgainstCcy_props
	pageDetail.AllInRate_props = rD.AllInRate_props
	pageDetail.MktRate_props = rD.MktRate_props
	pageDetail.SettleCcy_props = rD.SettleCcy_props
	pageDetail.Direction_props = rD.Direction_props
	pageDetail.NpvRate_props = rD.NpvRate_props
	pageDetail.OriginUser_props = rD.OriginUser_props
	pageDetail.PayInstruction_props = rD.PayInstruction_props
	pageDetail.ReceiptInstruction_props = rD.ReceiptInstruction_props
	pageDetail.NIName_props = rD.NIName_props
	pageDetail.CCYPair_props = rD.CCYPair_props
	pageDetail.Instrument_props = rD.Instrument_props
	pageDetail.PortfolioName_props = rD.PortfolioName_props
	pageDetail.RVDate_props = rD.RVDate_props
	pageDetail.RVMTM_props = rD.RVMTM_props
	pageDetail.CounterBook_props = rD.CounterBook_props
	pageDetail.CounterBookName_props = rD.CounterBookName_props
	pageDetail.Party_props = rD.Party_props
	pageDetail.PartyName_props = rD.PartyName_props
	pageDetail.NameCentre_props = rD.NameCentre_props
	pageDetail.NameFirm_props = rD.NameFirm_props
	pageDetail.CustomerExternalView_props = rD.CustomerExternalView_props
	pageDetail.CustomerSienaView_props = rD.CustomerSienaView_props
	pageDetail.CompID_props = rD.CompID_props
	pageDetail.SienaDealer_props = rD.SienaDealer_props
	pageDetail.DealOwner_props = rD.DealOwner_props
	pageDetail.DealOwnerMnemonic_props = rD.DealOwnerMnemonic_props
	pageDetail.EditedByUser_props = rD.EditedByUser_props
	pageDetail.UTCOriginTime_props = rD.UTCOriginTime_props
	pageDetail.UTCUpdateTime_props = rD.UTCUpdateTime_props
	pageDetail.MarginTrading_props = rD.MarginTrading_props
	pageDetail.SwapPoints_props = rD.SwapPoints_props
	pageDetail.SpotDate_props = rD.SpotDate_props
	pageDetail.SpotRate_props = rD.SpotRate_props
	pageDetail.MktSpotRate_props = rD.MktSpotRate_props
	pageDetail.SpotSalesMargin_props = rD.SpotSalesMargin_props
	pageDetail.SpotChlMargin_props = rD.SpotChlMargin_props
	pageDetail.RerouteCcy_props = rD.RerouteCcy_props
	pageDetail.CustomerPayInstruction_props = rD.CustomerPayInstruction_props
	pageDetail.CustomerReceiptInstruction_props = rD.CustomerReceiptInstruction_props
	pageDetail.BackOfficeNotes_props = rD.BackOfficeNotes_props
	pageDetail.CustomerStatementNotes_props = rD.CustomerStatementNotes_props
	pageDetail.NotesMargin_props = rD.NotesMargin_props
	pageDetail.RequestedBy_props = rD.RequestedBy_props
	pageDetail.EditReason_props = rD.EditReason_props
	pageDetail.EditOtherReason_props = rD.EditOtherReason_props
	pageDetail.NICleanPrice_props = rD.NICleanPrice_props
	pageDetail.NIDirtyPrice_props = rD.NIDirtyPrice_props
	pageDetail.NIYield_props = rD.NIYield_props
	pageDetail.NIClearingSystem_props = rD.NIClearingSystem_props
	pageDetail.Acceptor_props = rD.Acceptor_props
	pageDetail.NIDiscount_props = rD.NIDiscount_props
	pageDetail.FastPay_props = rD.FastPay_props
	pageDetail.PaymentFee_props = rD.PaymentFee_props
	pageDetail.PaymentFeePolicy_props = rD.PaymentFeePolicy_props
	pageDetail.PaymentReason_props = rD.PaymentReason_props
	pageDetail.PaymentDate_props = rD.PaymentDate_props
	pageDetail.SettlementDate_props = rD.SettlementDate_props
	pageDetail.FixingDate_props = rD.FixingDate_props
	pageDetail.VenueUTI_props = rD.VenueUTI_props
	pageDetail.EditVersion_props = rD.EditVersion_props
	pageDetail.BrokeragePercentage_props = rD.BrokeragePercentage_props
	pageDetail.BrokerageAmount_props = rD.BrokerageAmount_props
	pageDetail.BrokerageCurrency_props = rD.BrokerageCurrency_props
	pageDetail.BrokerageDate_props = rD.BrokerageDate_props
	pageDetail.AccountName_props = rD.AccountName_props
	pageDetail.AccountNumber_props = rD.AccountNumber_props
	pageDetail.CashBalance_props = rD.CashBalance_props
	pageDetail.DebitFrequency_props = rD.DebitFrequency_props
	pageDetail.CreditFrequency_props = rD.CreditFrequency_props
	pageDetail.ManuallyQuoted_props = rD.ManuallyQuoted_props
	pageDetail.LedgerBalance_props = rD.LedgerBalance_props
	pageDetail.SettAmtOutstanding_props = rD.SettAmtOutstanding_props
	pageDetail.FeePercentage_props = rD.FeePercentage_props
	pageDetail.FeeAmount_props = rD.FeeAmount_props
	pageDetail.Venue_props = rD.Venue_props
	pageDetail.EURAmount_props = rD.EURAmount_props
	pageDetail.EUROtherAmount_props = rD.EUROtherAmount_props
	pageDetail.LEI_props = rD.LEI_props
	pageDetail.Equity_props = rD.Equity_props
	pageDetail.Shares_props = rD.Shares_props
	pageDetail.QuoteExpiryDate_props = rD.QuoteExpiryDate_props
	pageDetail.Commodity_props = rD.Commodity_props
	pageDetail.PaymentSystemSienaView_props = rD.PaymentSystemSienaView_props
	pageDetail.PaymentSystemExternalView_props = rD.PaymentSystemExternalView_props
	pageDetail.SalesProfit_props = rD.SalesProfit_props
	pageDetail.RejectReason_props = rD.RejectReason_props
	pageDetail.PaymentError_props = rD.PaymentError_props
	pageDetail.RepoPrincipal_props = rD.RepoPrincipal_props
	pageDetail.FixingFrequency_props = rD.FixingFrequency_props
	pageDetail.Dealt_props = rD.Dealt_props
	pageDetail.Against_props = rD.Against_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	