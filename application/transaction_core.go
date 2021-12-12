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
// Date & Time		    : 12/12/2021 at 16:13:14
// Who & Where		    : matttownsend on silicon.local
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
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Transaction
}

//transaction_Page provides the information for the template for an individual Transaction
type Transaction_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		SienaReference string
		Status string
		ValueDate string
		MaturityDate string
		ContractNumber string
		ExternalReference string
		Book string
		MandatedUser string
		Portfolio string
		AgreementId string
		BackOfficeRefNo string
		ISIN string
		UTI string
		BookName string
		Centre string
		Firm string
		DealTypeShortName string
		FullDealType string
		TradeDate string
		DealtCcy string
		DealtAmount string
		AgainstAmount string
		AgainstCcy string
		AllInRate string
		MktRate string
		SettleCcy string
		Direction string
		NpvRate string
		OriginUser string
		PayInstruction string
		ReceiptInstruction string
		NIName string
		CCYPair string
		Instrument string
		PortfolioName string
		RVDate string
		RVMTM string
		CounterBook string
		CounterBookName string
		Party string
		PartyName string
		NameCentre string
		NameFirm string
		CustomerExternalView string
		CustomerSienaView string
		CompID string
		SienaDealer string
		DealOwner string
		DealOwnerMnemonic string
		EditedByUser string
		UTCOriginTime string
		UTCUpdateTime string
		MarginTrading string
		SwapPoints string
		SpotDate string
		SpotRate string
		MktSpotRate string
		SpotSalesMargin string
		SpotChlMargin string
		RerouteCcy string
		CustomerPayInstruction string
		CustomerReceiptInstruction string
		BackOfficeNotes string
		CustomerStatementNotes string
		NotesMargin string
		RequestedBy string
		EditReason string
		EditOtherReason string
		NICleanPrice string
		NIDirtyPrice string
		NIYield string
		NIClearingSystem string
		Acceptor string
		NIDiscount string
		FastPay string
		PaymentFee string
		PaymentFeePolicy string
		PaymentReason string
		PaymentDate string
		SettlementDate string
		FixingDate string
		VenueUTI string
		EditVersion string
		BrokeragePercentage string
		BrokerageAmount string
		BrokerageCurrency string
		BrokerageDate string
		AccountName string
		AccountNumber string
		CashBalance string
		DebitFrequency string
		CreditFrequency string
		ManuallyQuoted string
		LedgerBalance string
		SettAmtOutstanding string
		FeePercentage string
		FeeAmount string
		Venue string
		EURAmount string
		EUROtherAmount string
		LEI string
		Equity string
		Shares string
		QuoteExpiryDate string
		Commodity string
		PaymentSystemSienaView string
		PaymentSystemExternalView string
		SalesProfit string
		RejectReason string
		PaymentError string
		RepoPrincipal string
		FixingFrequency string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	Transaction_Redirect = dm.Transaction_PathList
)

//Transaction_Publish annouces the endpoints available for this object
func Transaction_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.Transaction_PathList, Transaction_HandlerList)
	mux.HandleFunc(dm.Transaction_PathView, Transaction_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Siena", dm.Transaction_Title)
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

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Transaction_TemplateView, w, r, pageDetail)

}

//Transaction_HandlerEdit is the handler used generate the Edit page
func Transaction_HandlerEdit(w http.ResponseWriter, r *http.Request) {
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
		Title:       CardTitle(dm.Transaction_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Transaction_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Transaction_TemplateEdit, w, r, pageDetail)


}

//Transaction_HandlerSave is the handler used process the saving of an Transaction
func Transaction_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("SienaReference"))

	var item dm.Transaction
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.SienaReference = r.FormValue(dm.Transaction_SienaReference)
		item.Status = r.FormValue(dm.Transaction_Status)
		item.ValueDate = r.FormValue(dm.Transaction_ValueDate)
		item.MaturityDate = r.FormValue(dm.Transaction_MaturityDate)
		item.ContractNumber = r.FormValue(dm.Transaction_ContractNumber)
		item.ExternalReference = r.FormValue(dm.Transaction_ExternalReference)
		item.Book = r.FormValue(dm.Transaction_Book)
		item.MandatedUser = r.FormValue(dm.Transaction_MandatedUser)
		item.Portfolio = r.FormValue(dm.Transaction_Portfolio)
		item.AgreementId = r.FormValue(dm.Transaction_AgreementId)
		item.BackOfficeRefNo = r.FormValue(dm.Transaction_BackOfficeRefNo)
		item.ISIN = r.FormValue(dm.Transaction_ISIN)
		item.UTI = r.FormValue(dm.Transaction_UTI)
		item.BookName = r.FormValue(dm.Transaction_BookName)
		item.Centre = r.FormValue(dm.Transaction_Centre)
		item.Firm = r.FormValue(dm.Transaction_Firm)
		item.DealTypeShortName = r.FormValue(dm.Transaction_DealTypeShortName)
		item.FullDealType = r.FormValue(dm.Transaction_FullDealType)
		item.TradeDate = r.FormValue(dm.Transaction_TradeDate)
		item.DealtCcy = r.FormValue(dm.Transaction_DealtCcy)
		item.DealtAmount = r.FormValue(dm.Transaction_DealtAmount)
		item.AgainstAmount = r.FormValue(dm.Transaction_AgainstAmount)
		item.AgainstCcy = r.FormValue(dm.Transaction_AgainstCcy)
		item.AllInRate = r.FormValue(dm.Transaction_AllInRate)
		item.MktRate = r.FormValue(dm.Transaction_MktRate)
		item.SettleCcy = r.FormValue(dm.Transaction_SettleCcy)
		item.Direction = r.FormValue(dm.Transaction_Direction)
		item.NpvRate = r.FormValue(dm.Transaction_NpvRate)
		item.OriginUser = r.FormValue(dm.Transaction_OriginUser)
		item.PayInstruction = r.FormValue(dm.Transaction_PayInstruction)
		item.ReceiptInstruction = r.FormValue(dm.Transaction_ReceiptInstruction)
		item.NIName = r.FormValue(dm.Transaction_NIName)
		item.CCYPair = r.FormValue(dm.Transaction_CCYPair)
		item.Instrument = r.FormValue(dm.Transaction_Instrument)
		item.PortfolioName = r.FormValue(dm.Transaction_PortfolioName)
		item.RVDate = r.FormValue(dm.Transaction_RVDate)
		item.RVMTM = r.FormValue(dm.Transaction_RVMTM)
		item.CounterBook = r.FormValue(dm.Transaction_CounterBook)
		item.CounterBookName = r.FormValue(dm.Transaction_CounterBookName)
		item.Party = r.FormValue(dm.Transaction_Party)
		item.PartyName = r.FormValue(dm.Transaction_PartyName)
		item.NameCentre = r.FormValue(dm.Transaction_NameCentre)
		item.NameFirm = r.FormValue(dm.Transaction_NameFirm)
		item.CustomerExternalView = r.FormValue(dm.Transaction_CustomerExternalView)
		item.CustomerSienaView = r.FormValue(dm.Transaction_CustomerSienaView)
		item.CompID = r.FormValue(dm.Transaction_CompID)
		item.SienaDealer = r.FormValue(dm.Transaction_SienaDealer)
		item.DealOwner = r.FormValue(dm.Transaction_DealOwner)
		item.DealOwnerMnemonic = r.FormValue(dm.Transaction_DealOwnerMnemonic)
		item.EditedByUser = r.FormValue(dm.Transaction_EditedByUser)
		item.UTCOriginTime = r.FormValue(dm.Transaction_UTCOriginTime)
		item.UTCUpdateTime = r.FormValue(dm.Transaction_UTCUpdateTime)
		item.MarginTrading = r.FormValue(dm.Transaction_MarginTrading)
		item.SwapPoints = r.FormValue(dm.Transaction_SwapPoints)
		item.SpotDate = r.FormValue(dm.Transaction_SpotDate)
		item.SpotRate = r.FormValue(dm.Transaction_SpotRate)
		item.MktSpotRate = r.FormValue(dm.Transaction_MktSpotRate)
		item.SpotSalesMargin = r.FormValue(dm.Transaction_SpotSalesMargin)
		item.SpotChlMargin = r.FormValue(dm.Transaction_SpotChlMargin)
		item.RerouteCcy = r.FormValue(dm.Transaction_RerouteCcy)
		item.CustomerPayInstruction = r.FormValue(dm.Transaction_CustomerPayInstruction)
		item.CustomerReceiptInstruction = r.FormValue(dm.Transaction_CustomerReceiptInstruction)
		item.BackOfficeNotes = r.FormValue(dm.Transaction_BackOfficeNotes)
		item.CustomerStatementNotes = r.FormValue(dm.Transaction_CustomerStatementNotes)
		item.NotesMargin = r.FormValue(dm.Transaction_NotesMargin)
		item.RequestedBy = r.FormValue(dm.Transaction_RequestedBy)
		item.EditReason = r.FormValue(dm.Transaction_EditReason)
		item.EditOtherReason = r.FormValue(dm.Transaction_EditOtherReason)
		item.NICleanPrice = r.FormValue(dm.Transaction_NICleanPrice)
		item.NIDirtyPrice = r.FormValue(dm.Transaction_NIDirtyPrice)
		item.NIYield = r.FormValue(dm.Transaction_NIYield)
		item.NIClearingSystem = r.FormValue(dm.Transaction_NIClearingSystem)
		item.Acceptor = r.FormValue(dm.Transaction_Acceptor)
		item.NIDiscount = r.FormValue(dm.Transaction_NIDiscount)
		item.FastPay = r.FormValue(dm.Transaction_FastPay)
		item.PaymentFee = r.FormValue(dm.Transaction_PaymentFee)
		item.PaymentFeePolicy = r.FormValue(dm.Transaction_PaymentFeePolicy)
		item.PaymentReason = r.FormValue(dm.Transaction_PaymentReason)
		item.PaymentDate = r.FormValue(dm.Transaction_PaymentDate)
		item.SettlementDate = r.FormValue(dm.Transaction_SettlementDate)
		item.FixingDate = r.FormValue(dm.Transaction_FixingDate)
		item.VenueUTI = r.FormValue(dm.Transaction_VenueUTI)
		item.EditVersion = r.FormValue(dm.Transaction_EditVersion)
		item.BrokeragePercentage = r.FormValue(dm.Transaction_BrokeragePercentage)
		item.BrokerageAmount = r.FormValue(dm.Transaction_BrokerageAmount)
		item.BrokerageCurrency = r.FormValue(dm.Transaction_BrokerageCurrency)
		item.BrokerageDate = r.FormValue(dm.Transaction_BrokerageDate)
		item.AccountName = r.FormValue(dm.Transaction_AccountName)
		item.AccountNumber = r.FormValue(dm.Transaction_AccountNumber)
		item.CashBalance = r.FormValue(dm.Transaction_CashBalance)
		item.DebitFrequency = r.FormValue(dm.Transaction_DebitFrequency)
		item.CreditFrequency = r.FormValue(dm.Transaction_CreditFrequency)
		item.ManuallyQuoted = r.FormValue(dm.Transaction_ManuallyQuoted)
		item.LedgerBalance = r.FormValue(dm.Transaction_LedgerBalance)
		item.SettAmtOutstanding = r.FormValue(dm.Transaction_SettAmtOutstanding)
		item.FeePercentage = r.FormValue(dm.Transaction_FeePercentage)
		item.FeeAmount = r.FormValue(dm.Transaction_FeeAmount)
		item.Venue = r.FormValue(dm.Transaction_Venue)
		item.EURAmount = r.FormValue(dm.Transaction_EURAmount)
		item.EUROtherAmount = r.FormValue(dm.Transaction_EUROtherAmount)
		item.LEI = r.FormValue(dm.Transaction_LEI)
		item.Equity = r.FormValue(dm.Transaction_Equity)
		item.Shares = r.FormValue(dm.Transaction_Shares)
		item.QuoteExpiryDate = r.FormValue(dm.Transaction_QuoteExpiryDate)
		item.Commodity = r.FormValue(dm.Transaction_Commodity)
		item.PaymentSystemSienaView = r.FormValue(dm.Transaction_PaymentSystemSienaView)
		item.PaymentSystemExternalView = r.FormValue(dm.Transaction_PaymentSystemExternalView)
		item.SalesProfit = r.FormValue(dm.Transaction_SalesProfit)
		item.RejectReason = r.FormValue(dm.Transaction_RejectReason)
		item.PaymentError = r.FormValue(dm.Transaction_PaymentError)
		item.RepoPrincipal = r.FormValue(dm.Transaction_RepoPrincipal)
		item.FixingFrequency = r.FormValue(dm.Transaction_FixingFrequency)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.Transaction_Store(item,r)	

	http.Redirect(w, r, Transaction_Redirect, http.StatusFound)
}

//Transaction_HandlerNew is the handler used process the creation of an Transaction
func Transaction_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Transaction_Page{
		Title:       CardTitle(dm.Transaction_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Transaction_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SienaReference = ""
pageDetail.Status = ""
pageDetail.ValueDate = ""
pageDetail.MaturityDate = ""
pageDetail.ContractNumber = ""
pageDetail.ExternalReference = ""
pageDetail.Book = ""
pageDetail.MandatedUser = ""
pageDetail.Portfolio = ""
pageDetail.AgreementId = ""
pageDetail.BackOfficeRefNo = ""
pageDetail.ISIN = ""
pageDetail.UTI = ""
pageDetail.BookName = ""
pageDetail.Centre = ""
pageDetail.Firm = ""
pageDetail.DealTypeShortName = ""
pageDetail.FullDealType = ""
pageDetail.TradeDate = ""
pageDetail.DealtCcy = ""
pageDetail.DealtAmount = ""
pageDetail.AgainstAmount = ""
pageDetail.AgainstCcy = ""
pageDetail.AllInRate = ""
pageDetail.MktRate = ""
pageDetail.SettleCcy = ""
pageDetail.Direction = ""
pageDetail.NpvRate = ""
pageDetail.OriginUser = ""
pageDetail.PayInstruction = ""
pageDetail.ReceiptInstruction = ""
pageDetail.NIName = ""
pageDetail.CCYPair = ""
pageDetail.Instrument = ""
pageDetail.PortfolioName = ""
pageDetail.RVDate = ""
pageDetail.RVMTM = ""
pageDetail.CounterBook = ""
pageDetail.CounterBookName = ""
pageDetail.Party = ""
pageDetail.PartyName = ""
pageDetail.NameCentre = ""
pageDetail.NameFirm = ""
pageDetail.CustomerExternalView = ""
pageDetail.CustomerSienaView = ""
pageDetail.CompID = ""
pageDetail.SienaDealer = ""
pageDetail.DealOwner = ""
pageDetail.DealOwnerMnemonic = ""
pageDetail.EditedByUser = ""
pageDetail.UTCOriginTime = ""
pageDetail.UTCUpdateTime = ""
pageDetail.MarginTrading = ""
pageDetail.SwapPoints = ""
pageDetail.SpotDate = ""
pageDetail.SpotRate = ""
pageDetail.MktSpotRate = ""
pageDetail.SpotSalesMargin = ""
pageDetail.SpotChlMargin = ""
pageDetail.RerouteCcy = ""
pageDetail.CustomerPayInstruction = ""
pageDetail.CustomerReceiptInstruction = ""
pageDetail.BackOfficeNotes = ""
pageDetail.CustomerStatementNotes = ""
pageDetail.NotesMargin = ""
pageDetail.RequestedBy = ""
pageDetail.EditReason = ""
pageDetail.EditOtherReason = ""
pageDetail.NICleanPrice = ""
pageDetail.NIDirtyPrice = ""
pageDetail.NIYield = ""
pageDetail.NIClearingSystem = ""
pageDetail.Acceptor = ""
pageDetail.NIDiscount = ""
pageDetail.FastPay = ""
pageDetail.PaymentFee = ""
pageDetail.PaymentFeePolicy = ""
pageDetail.PaymentReason = ""
pageDetail.PaymentDate = ""
pageDetail.SettlementDate = ""
pageDetail.FixingDate = ""
pageDetail.VenueUTI = ""
pageDetail.EditVersion = ""
pageDetail.BrokeragePercentage = ""
pageDetail.BrokerageAmount = ""
pageDetail.BrokerageCurrency = ""
pageDetail.BrokerageDate = ""
pageDetail.AccountName = ""
pageDetail.AccountNumber = ""
pageDetail.CashBalance = ""
pageDetail.DebitFrequency = ""
pageDetail.CreditFrequency = ""
pageDetail.ManuallyQuoted = ""
pageDetail.LedgerBalance = ""
pageDetail.SettAmtOutstanding = ""
pageDetail.FeePercentage = ""
pageDetail.FeeAmount = ""
pageDetail.Venue = ""
pageDetail.EURAmount = ""
pageDetail.EUROtherAmount = ""
pageDetail.LEI = ""
pageDetail.Equity = ""
pageDetail.Shares = ""
pageDetail.QuoteExpiryDate = ""
pageDetail.Commodity = ""
pageDetail.PaymentSystemSienaView = ""
pageDetail.PaymentSystemExternalView = ""
pageDetail.SalesProfit = ""
pageDetail.RejectReason = ""
pageDetail.PaymentError = ""
pageDetail.RepoPrincipal = ""
pageDetail.FixingFrequency = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Transaction_TemplateNew, w, r, pageDetail)

}

//Transaction_HandlerDelete is the handler used process the deletion of an Transaction
func Transaction_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Transaction_QueryString)

	dao.Transaction_Delete(searchID)	

	http.Redirect(w, r, Transaction_Redirect, http.StatusFound)
}