package application
// ----------------------------------------------------------------
// Automatically generated  "/application/account.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Account (account)
// Endpoint 	        : Account (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 29/06/2022 at 22:43:27
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





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

	pageDetail := dm.Account_PageList{
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

	pageDetail := dm.Account_Page{
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
func account_PopulatePage(rD dm.Account, pageDetail dm.Account_Page) dm.Account_Page {
	// START
	// Dynamically generated 29/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Automatically generated 29/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.CCY_lookup = dao.Currency_GetLookup()
	
	
	
	pageDetail.Book_lookup = dao.Book_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Portfolio_lookup = dao.Portfolio_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Centre_lookup = dao.Centre_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Firm_lookup = dao.Firm_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SienaReference_props = rD.SienaReference_props
	pageDetail.CustomerSienaView_props = rD.CustomerSienaView_props
	pageDetail.SienaCommonRef_props = rD.SienaCommonRef_props
	pageDetail.Status_props = rD.Status_props
	pageDetail.StartDate_props = rD.StartDate_props
	pageDetail.MaturityDate_props = rD.MaturityDate_props
	pageDetail.ContractNumber_props = rD.ContractNumber_props
	pageDetail.ExternalReference_props = rD.ExternalReference_props
	pageDetail.CCY_props = rD.CCY_props
	pageDetail.Book_props = rD.Book_props
	pageDetail.MandatedUser_props = rD.MandatedUser_props
	pageDetail.BackOfficeNotes_props = rD.BackOfficeNotes_props
	pageDetail.CashBalance_props = rD.CashBalance_props
	pageDetail.AccountNumber_props = rD.AccountNumber_props
	pageDetail.AccountName_props = rD.AccountName_props
	pageDetail.LedgerBalance_props = rD.LedgerBalance_props
	pageDetail.Portfolio_props = rD.Portfolio_props
	pageDetail.AgreementId_props = rD.AgreementId_props
	pageDetail.BackOfficeRefNo_props = rD.BackOfficeRefNo_props
	pageDetail.ISIN_props = rD.ISIN_props
	pageDetail.UTI_props = rD.UTI_props
	pageDetail.CCYName_props = rD.CCYName_props
	pageDetail.BookName_props = rD.BookName_props
	pageDetail.PortfolioName_props = rD.PortfolioName_props
	pageDetail.Centre_props = rD.Centre_props
	pageDetail.DealTypeKey_props = rD.DealTypeKey_props
	pageDetail.DealTypeShortName_props = rD.DealTypeShortName_props
	pageDetail.InternalId_props = rD.InternalId_props
	pageDetail.InternalDeleted_props = rD.InternalDeleted_props
	pageDetail.UpdatedTransactionId_props = rD.UpdatedTransactionId_props
	pageDetail.UpdatedUserId_props = rD.UpdatedUserId_props
	pageDetail.UpdatedDateTime_props = rD.UpdatedDateTime_props
	pageDetail.DeletedTransactionId_props = rD.DeletedTransactionId_props
	pageDetail.DeletedUserId_props = rD.DeletedUserId_props
	pageDetail.ChangeType_props = rD.ChangeType_props
	pageDetail.CCYDp_props = rD.CCYDp_props
	pageDetail.CompID_props = rD.CompID_props
	pageDetail.Firm_props = rD.Firm_props
	pageDetail.DealType_props = rD.DealType_props
	pageDetail.FullDealType_props = rD.FullDealType_props
	pageDetail.DealingInterface_props = rD.DealingInterface_props
	pageDetail.DealtAmount_props = rD.DealtAmount_props
	pageDetail.ParentContractNumber_props = rD.ParentContractNumber_props
	pageDetail.InterestFrequency_props = rD.InterestFrequency_props
	pageDetail.InterestAction_props = rD.InterestAction_props
	pageDetail.InterestStrategy_props = rD.InterestStrategy_props
	pageDetail.InterestBasis_props = rD.InterestBasis_props
	pageDetail.SienaDealer_props = rD.SienaDealer_props
	pageDetail.DealOwner_props = rD.DealOwner_props
	pageDetail.OriginUser_props = rD.OriginUser_props
	pageDetail.EditedByUser_props = rD.EditedByUser_props
	pageDetail.DealOwnerMnemonic_props = rD.DealOwnerMnemonic_props
	pageDetail.UTCOriginTime_props = rD.UTCOriginTime_props
	pageDetail.UTCUpdateTime_props = rD.UTCUpdateTime_props
	pageDetail.CustomerStatementNotes_props = rD.CustomerStatementNotes_props
	pageDetail.NotesMargin_props = rD.NotesMargin_props
	pageDetail.RequestedBy_props = rD.RequestedBy_props
	pageDetail.EditReason_props = rD.EditReason_props
	pageDetail.EditOtherReason_props = rD.EditOtherReason_props
	pageDetail.NoticeDays_props = rD.NoticeDays_props
	pageDetail.DebitFrequency_props = rD.DebitFrequency_props
	pageDetail.CreditFrequency_props = rD.CreditFrequency_props
	pageDetail.EURAmount_props = rD.EURAmount_props
	pageDetail.EUROtherAmount_props = rD.EUROtherAmount_props
	pageDetail.PaymentSystemSienaView_props = rD.PaymentSystemSienaView_props
	pageDetail.PaymentSystemExternalView_props = rD.PaymentSystemExternalView_props
	pageDetail.DealtCA_props = rD.DealtCA_props
	pageDetail.AgainstCA_props = rD.AgainstCA_props
	pageDetail.LedgerCA_props = rD.LedgerCA_props
	pageDetail.CashBalanceCA_props = rD.CashBalanceCA_props
	
	// 
	// Dynamically generated 29/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	