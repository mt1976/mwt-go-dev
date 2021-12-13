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
// Date & Time		    : 13/12/2021 at 17:02:22
// Who & Where		    : matttownsend on silicon.local
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
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Account
}

//account_Page provides the information for the template for an individual Account
type Account_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 13/12/2021 by matttownsend on silicon.local - START
		SienaReference string
		CustomerSienaView string
		SienaCommonRef string
		Status string
		StartDate string
		MaturityDate string
		ContractNumber string
		ExternalReference string
		CCY string
		Book string
		MandatedUser string
		BackOfficeNotes string
		CashBalance string
		AccountNumber string
		AccountName string
		LedgerBalance string
		Portfolio string
		AgreementId string
		BackOfficeRefNo string
		ISIN string
		UTI string
		CCYName string
		BookName string
		PortfolioName string
		Centre string
		DealTypeKey string
		DealTypeShortName string
		InternalId string
		InternalDeleted string
		UpdatedTransactionId string
		UpdatedUserId string
		UpdatedDateTime string
		DeletedTransactionId string
		DeletedUserId string
		ChangeType string
		CCYDp string
		CompID string
		Firm string
		DealType string
		FullDealType string
		DealingInterface string
		DealtAmount string
		ParentContractNumber string
		InterestFrequency string
		InterestAction string
		InterestStrategy string
		InterestBasis string
		SienaDealer string
		DealOwner string
		OriginUser string
		EditedByUser string
		DealOwnerMnemonic string
		UTCOriginTime string
		UTCUpdateTime string
		CustomerStatementNotes string
		NotesMargin string
		RequestedBy string
		EditReason string
		EditOtherReason string
		NoticeDays string
		DebitFrequency string
		CreditFrequency string
		EURAmount string
		EUROtherAmount string
		PaymentSystemSienaView string
		PaymentSystemExternalView string
		CCY_Lookup string
		Book_Lookup string
		Portfolio_Lookup string
		Centre_Lookup string
		Firm_Lookup string
		DealtCA_Extra string
		AgainstCA_Extra string
		LedgerCA_Extra string
		CashBalanceCA_Extra string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	CCY_Lookup_List	[]dm.Currency
	Book_Lookup_List	[]dm.Book
	Portfolio_Lookup_List	[]dm.Portfolio
	Centre_Lookup_List	[]dm.Centre
	Firm_Lookup_List	[]dm.Firm
	
	
	
	
	
	// Automatically generated 13/12/2021 by matttownsend on silicon.local - END
}

const (
	Account_Redirect = dm.Account_PathList
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
	logs.Publish("Siena", dm.Account_Title)
    core.Catalog_Add(dm.Account_Title, dm.Account_Path, "", dm.Account_QueryString, "APP")
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

		// 
		// Automatically generated 13/12/2021 by matttownsend on silicon.local - START
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

pageDetail.DealtCA_Extra = rD.DealtCA_Extra
pageDetail.AgainstCA_Extra = rD.AgainstCA_Extra
pageDetail.LedgerCA_Extra = rD.LedgerCA_Extra
pageDetail.CashBalanceCA_Extra = rD.CashBalanceCA_Extra

// Automatically generated 13/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,CCY_Lookup_Name,_:= dao.Currency_GetByID(rD.CCY)
pageDetail.CCY_Lookup = CCY_Lookup_Name.Name
_,Book_Lookup_FullName,_:= dao.Book_GetByID(rD.Book)
pageDetail.Book_Lookup = Book_Lookup_FullName.FullName
_,Portfolio_Lookup_Description1,_:= dao.Portfolio_GetByID(rD.Portfolio)
pageDetail.Portfolio_Lookup = Portfolio_Lookup_Description1.Description1
_,Centre_Lookup_Name,_:= dao.Centre_GetByID(rD.Centre)
pageDetail.Centre_Lookup = Centre_Lookup_Name.Name
_,Firm_Lookup_FullName,_:= dao.Firm_GetByID(rD.Firm)
pageDetail.Firm_Lookup = Firm_Lookup_FullName.FullName
// Automatically generated 13/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 13/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Account_TemplateView, w, r, pageDetail)

}

//Account_HandlerEdit is the handler used generate the Edit page
func Account_HandlerEdit(w http.ResponseWriter, r *http.Request) {
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
		Title:       CardTitle(dm.Account_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Account_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 13/12/2021 by matttownsend on silicon.local - START
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

pageDetail.DealtCA_Extra = rD.DealtCA_Extra
pageDetail.AgainstCA_Extra = rD.AgainstCA_Extra
pageDetail.LedgerCA_Extra = rD.LedgerCA_Extra
pageDetail.CashBalanceCA_Extra = rD.CashBalanceCA_Extra

// Automatically generated 13/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,CCY_Lookup_Name,_:= dao.Currency_GetByID(rD.CCY)
pageDetail.CCY_Lookup = CCY_Lookup_Name.Name
_,pageDetail.CCY_Lookup_List,_ = dao.Currency_GetList()
_,Book_Lookup_FullName,_:= dao.Book_GetByID(rD.Book)
pageDetail.Book_Lookup = Book_Lookup_FullName.FullName
_,pageDetail.Book_Lookup_List,_ = dao.Book_GetList()
_,Portfolio_Lookup_Description1,_:= dao.Portfolio_GetByID(rD.Portfolio)
pageDetail.Portfolio_Lookup = Portfolio_Lookup_Description1.Description1
_,pageDetail.Portfolio_Lookup_List,_ = dao.Portfolio_GetList()
_,Centre_Lookup_Name,_:= dao.Centre_GetByID(rD.Centre)
pageDetail.Centre_Lookup = Centre_Lookup_Name.Name
_,pageDetail.Centre_Lookup_List,_ = dao.Centre_GetList()
_,Firm_Lookup_FullName,_:= dao.Firm_GetByID(rD.Firm)
pageDetail.Firm_Lookup = Firm_Lookup_FullName.FullName
_,pageDetail.Firm_Lookup_List,_ = dao.Firm_GetList()
// Automatically generated 13/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 13/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Account_TemplateEdit, w, r, pageDetail)


}

//Account_HandlerSave is the handler used process the saving of an Account
func Account_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("SienaReference"))

	var item dm.Account
	// Automatically generated 13/12/2021 by matttownsend on silicon.local - START
		item.SienaReference = r.FormValue(dm.Account_SienaReference)
		item.CustomerSienaView = r.FormValue(dm.Account_CustomerSienaView)
		item.SienaCommonRef = r.FormValue(dm.Account_SienaCommonRef)
		item.Status = r.FormValue(dm.Account_Status)
		item.StartDate = r.FormValue(dm.Account_StartDate)
		item.MaturityDate = r.FormValue(dm.Account_MaturityDate)
		item.ContractNumber = r.FormValue(dm.Account_ContractNumber)
		item.ExternalReference = r.FormValue(dm.Account_ExternalReference)
		item.CCY = r.FormValue(dm.Account_CCY)
		item.Book = r.FormValue(dm.Account_Book)
		item.MandatedUser = r.FormValue(dm.Account_MandatedUser)
		item.BackOfficeNotes = r.FormValue(dm.Account_BackOfficeNotes)
		item.CashBalance = r.FormValue(dm.Account_CashBalance)
		item.AccountNumber = r.FormValue(dm.Account_AccountNumber)
		item.AccountName = r.FormValue(dm.Account_AccountName)
		item.LedgerBalance = r.FormValue(dm.Account_LedgerBalance)
		item.Portfolio = r.FormValue(dm.Account_Portfolio)
		item.AgreementId = r.FormValue(dm.Account_AgreementId)
		item.BackOfficeRefNo = r.FormValue(dm.Account_BackOfficeRefNo)
		item.ISIN = r.FormValue(dm.Account_ISIN)
		item.UTI = r.FormValue(dm.Account_UTI)
		item.CCYName = r.FormValue(dm.Account_CCYName)
		item.BookName = r.FormValue(dm.Account_BookName)
		item.PortfolioName = r.FormValue(dm.Account_PortfolioName)
		item.Centre = r.FormValue(dm.Account_Centre)
		item.DealTypeKey = r.FormValue(dm.Account_DealTypeKey)
		item.DealTypeShortName = r.FormValue(dm.Account_DealTypeShortName)
		item.InternalId = r.FormValue(dm.Account_InternalId)
		item.InternalDeleted = r.FormValue(dm.Account_InternalDeleted)
		item.UpdatedTransactionId = r.FormValue(dm.Account_UpdatedTransactionId)
		item.UpdatedUserId = r.FormValue(dm.Account_UpdatedUserId)
		item.UpdatedDateTime = r.FormValue(dm.Account_UpdatedDateTime)
		item.DeletedTransactionId = r.FormValue(dm.Account_DeletedTransactionId)
		item.DeletedUserId = r.FormValue(dm.Account_DeletedUserId)
		item.ChangeType = r.FormValue(dm.Account_ChangeType)
		item.CCYDp = r.FormValue(dm.Account_CCYDp)
		item.CompID = r.FormValue(dm.Account_CompID)
		item.Firm = r.FormValue(dm.Account_Firm)
		item.DealType = r.FormValue(dm.Account_DealType)
		item.FullDealType = r.FormValue(dm.Account_FullDealType)
		item.DealingInterface = r.FormValue(dm.Account_DealingInterface)
		item.DealtAmount = r.FormValue(dm.Account_DealtAmount)
		item.ParentContractNumber = r.FormValue(dm.Account_ParentContractNumber)
		item.InterestFrequency = r.FormValue(dm.Account_InterestFrequency)
		item.InterestAction = r.FormValue(dm.Account_InterestAction)
		item.InterestStrategy = r.FormValue(dm.Account_InterestStrategy)
		item.InterestBasis = r.FormValue(dm.Account_InterestBasis)
		item.SienaDealer = r.FormValue(dm.Account_SienaDealer)
		item.DealOwner = r.FormValue(dm.Account_DealOwner)
		item.OriginUser = r.FormValue(dm.Account_OriginUser)
		item.EditedByUser = r.FormValue(dm.Account_EditedByUser)
		item.DealOwnerMnemonic = r.FormValue(dm.Account_DealOwnerMnemonic)
		item.UTCOriginTime = r.FormValue(dm.Account_UTCOriginTime)
		item.UTCUpdateTime = r.FormValue(dm.Account_UTCUpdateTime)
		item.CustomerStatementNotes = r.FormValue(dm.Account_CustomerStatementNotes)
		item.NotesMargin = r.FormValue(dm.Account_NotesMargin)
		item.RequestedBy = r.FormValue(dm.Account_RequestedBy)
		item.EditReason = r.FormValue(dm.Account_EditReason)
		item.EditOtherReason = r.FormValue(dm.Account_EditOtherReason)
		item.NoticeDays = r.FormValue(dm.Account_NoticeDays)
		item.DebitFrequency = r.FormValue(dm.Account_DebitFrequency)
		item.CreditFrequency = r.FormValue(dm.Account_CreditFrequency)
		item.EURAmount = r.FormValue(dm.Account_EURAmount)
		item.EUROtherAmount = r.FormValue(dm.Account_EUROtherAmount)
		item.PaymentSystemSienaView = r.FormValue(dm.Account_PaymentSystemSienaView)
		item.PaymentSystemExternalView = r.FormValue(dm.Account_PaymentSystemExternalView)
		item.CCY_Lookup = r.FormValue(dm.Account_CCY_Lookup)
		item.Book_Lookup = r.FormValue(dm.Account_Book_Lookup)
		item.Portfolio_Lookup = r.FormValue(dm.Account_Portfolio_Lookup)
		item.Centre_Lookup = r.FormValue(dm.Account_Centre_Lookup)
		item.Firm_Lookup = r.FormValue(dm.Account_Firm_Lookup)
		item.DealtCA_Extra = r.FormValue(dm.Account_DealtCA_Extra)
		item.AgainstCA_Extra = r.FormValue(dm.Account_AgainstCA_Extra)
		item.LedgerCA_Extra = r.FormValue(dm.Account_LedgerCA_Extra)
		item.CashBalanceCA_Extra = r.FormValue(dm.Account_CashBalanceCA_Extra)
	

	// Automatically generated 13/12/2021 by matttownsend on silicon.local - END

	dao.Account_Store(item,r)	

	http.Redirect(w, r, Account_Redirect, http.StatusFound)
}

//Account_HandlerNew is the handler used process the creation of an Account
func Account_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Account_Page{
		Title:       CardTitle(dm.Account_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Account_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 13/12/2021 by matttownsend on silicon.local - START
pageDetail.SienaReference = ""
pageDetail.CustomerSienaView = ""
pageDetail.SienaCommonRef = ""
pageDetail.Status = ""
pageDetail.StartDate = ""
pageDetail.MaturityDate = ""
pageDetail.ContractNumber = ""
pageDetail.ExternalReference = ""
pageDetail.CCY = ""
pageDetail.Book = ""
pageDetail.MandatedUser = ""
pageDetail.BackOfficeNotes = ""
pageDetail.CashBalance = ""
pageDetail.AccountNumber = ""
pageDetail.AccountName = ""
pageDetail.LedgerBalance = ""
pageDetail.Portfolio = ""
pageDetail.AgreementId = ""
pageDetail.BackOfficeRefNo = ""
pageDetail.ISIN = ""
pageDetail.UTI = ""
pageDetail.CCYName = ""
pageDetail.BookName = ""
pageDetail.PortfolioName = ""
pageDetail.Centre = ""
pageDetail.DealTypeKey = ""
pageDetail.DealTypeShortName = ""
pageDetail.InternalId = ""
pageDetail.InternalDeleted = ""
pageDetail.UpdatedTransactionId = ""
pageDetail.UpdatedUserId = ""
pageDetail.UpdatedDateTime = ""
pageDetail.DeletedTransactionId = ""
pageDetail.DeletedUserId = ""
pageDetail.ChangeType = ""
pageDetail.CCYDp = ""
pageDetail.CompID = ""
pageDetail.Firm = ""
pageDetail.DealType = ""
pageDetail.FullDealType = ""
pageDetail.DealingInterface = ""
pageDetail.DealtAmount = ""
pageDetail.ParentContractNumber = ""
pageDetail.InterestFrequency = ""
pageDetail.InterestAction = ""
pageDetail.InterestStrategy = ""
pageDetail.InterestBasis = ""
pageDetail.SienaDealer = ""
pageDetail.DealOwner = ""
pageDetail.OriginUser = ""
pageDetail.EditedByUser = ""
pageDetail.DealOwnerMnemonic = ""
pageDetail.UTCOriginTime = ""
pageDetail.UTCUpdateTime = ""
pageDetail.CustomerStatementNotes = ""
pageDetail.NotesMargin = ""
pageDetail.RequestedBy = ""
pageDetail.EditReason = ""
pageDetail.EditOtherReason = ""
pageDetail.NoticeDays = ""
pageDetail.DebitFrequency = ""
pageDetail.CreditFrequency = ""
pageDetail.EURAmount = ""
pageDetail.EUROtherAmount = ""
pageDetail.PaymentSystemSienaView = ""
pageDetail.PaymentSystemExternalView = ""

pageDetail.DealtCA_Extra = ""
pageDetail.AgainstCA_Extra = ""
pageDetail.LedgerCA_Extra = ""
pageDetail.CashBalanceCA_Extra = ""

// Automatically generated 13/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
pageDetail.CCY_Lookup = ""
_,pageDetail.CCY_Lookup_List,_ = dao.Currency_GetList()
pageDetail.Book_Lookup = ""
_,pageDetail.Book_Lookup_List,_ = dao.Book_GetList()
pageDetail.Portfolio_Lookup = ""
_,pageDetail.Portfolio_Lookup_List,_ = dao.Portfolio_GetList()
pageDetail.Centre_Lookup = ""
_,pageDetail.Centre_Lookup_List,_ = dao.Centre_GetList()
pageDetail.Firm_Lookup = ""
_,pageDetail.Firm_Lookup_List,_ = dao.Firm_GetList()
// Automatically generated 13/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Account_TemplateNew, w, r, pageDetail)

}

//Account_HandlerDelete is the handler used process the deletion of an Account
func Account_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Account_QueryString)

	dao.Account_Delete(searchID)	

	http.Redirect(w, r, Account_Redirect, http.StatusFound)
}