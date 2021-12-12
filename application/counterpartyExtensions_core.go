package application
// ----------------------------------------------------------------
// Automatically generated  "/application/counterpartyextensions.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyExtensions (counterpartyextensions)
// Endpoint 	        : CounterpartyExtensions (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:09
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//counterpartyextensions_PageList provides the information for the template for a list of CounterpartyExtensionss
type CounterpartyExtensions_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.CounterpartyExtensions
}

//counterpartyextensions_Page provides the information for the template for an individual CounterpartyExtensions
type CounterpartyExtensions_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		NameFirm string
		NameCentre string
		BICCode string
		ContactIndicator string
		CoverTrade string
		CustomerCategory string
		FSCSInclusive string
		FeeFactor string
		InactiveStatus string
		Indemnity string
		KnowYourCustomerStatus string
		LERLimitCarveOut string
		LastAmended string
		LastLogin string
		LossGivenDefault string
		MIC string
		ProtectedDepositor string
		RPTCurrency string
		RateTimeout string
		RateValidation string
		Registered string
		RegulatoryCategory string
		SecuredSettlement string
		SettlementLimitCarveOut string
		SortCode string
		Training string
		TrainingCode string
		TrainingReceived string
		Unencumbered string
		LEIExpiryDate string
		MIFIDReviewDate string
		GDPRReviewDate string
		DelegatedReporting string
		BOReconcile string
		MIFIDReportableDealsAllowed string
		SignedInvestmentAgreement string
		AppropriatenessAssessment string
		FinancialCounterparty string
		Collateralisation string
		PortfolioCode string
		ReconciliationLetterFrequency string
		DirectDealing string
		CompID string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	CounterpartyExtensions_Redirect = dm.CounterpartyExtensions_PathList
)

//CounterpartyExtensions_Publish annouces the endpoints available for this object
func CounterpartyExtensions_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.CounterpartyExtensions_PathList, CounterpartyExtensions_HandlerList)
	mux.HandleFunc(dm.CounterpartyExtensions_PathView, CounterpartyExtensions_HandlerView)
	mux.HandleFunc(dm.CounterpartyExtensions_PathEdit, CounterpartyExtensions_HandlerEdit)
	mux.HandleFunc(dm.CounterpartyExtensions_PathNew, CounterpartyExtensions_HandlerNew)
	mux.HandleFunc(dm.CounterpartyExtensions_PathSave, CounterpartyExtensions_HandlerSave)
	mux.HandleFunc(dm.CounterpartyExtensions_PathDelete, CounterpartyExtensions_HandlerDelete)
	logs.Publish("Siena", dm.CounterpartyExtensions_Title)
    //No API
}

//CounterpartyExtensions_HandlerList is the handler for the list page
func CounterpartyExtensions_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.CounterpartyExtensions
	noItems, returnList, _ := dao.CounterpartyExtensions_GetList()

	pageDetail := CounterpartyExtensions_PageList{
		Title:            CardTitle(dm.CounterpartyExtensions_Title, core.Action_List),
		PageTitle:        PageTitle(dm.CounterpartyExtensions_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.CounterpartyExtensions_TemplateList, w, r, pageDetail)

}

//CounterpartyExtensions_HandlerView is the handler used to View a page
func CounterpartyExtensions_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyExtensions_QueryString)
	_, rD, _ := dao.CounterpartyExtensions_GetByID(searchID)

	pageDetail := CounterpartyExtensions_Page{
		Title:       CardTitle(dm.CounterpartyExtensions_Title, core.Action_View),
		PageTitle:   PageTitle(dm.CounterpartyExtensions_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.NameFirm = rD.NameFirm
pageDetail.NameCentre = rD.NameCentre
pageDetail.BICCode = rD.BICCode
pageDetail.ContactIndicator = rD.ContactIndicator
pageDetail.CoverTrade = rD.CoverTrade
pageDetail.CustomerCategory = rD.CustomerCategory
pageDetail.FSCSInclusive = rD.FSCSInclusive
pageDetail.FeeFactor = rD.FeeFactor
pageDetail.InactiveStatus = rD.InactiveStatus
pageDetail.Indemnity = rD.Indemnity
pageDetail.KnowYourCustomerStatus = rD.KnowYourCustomerStatus
pageDetail.LERLimitCarveOut = rD.LERLimitCarveOut
pageDetail.LastAmended = rD.LastAmended
pageDetail.LastLogin = rD.LastLogin
pageDetail.LossGivenDefault = rD.LossGivenDefault
pageDetail.MIC = rD.MIC
pageDetail.ProtectedDepositor = rD.ProtectedDepositor
pageDetail.RPTCurrency = rD.RPTCurrency
pageDetail.RateTimeout = rD.RateTimeout
pageDetail.RateValidation = rD.RateValidation
pageDetail.Registered = rD.Registered
pageDetail.RegulatoryCategory = rD.RegulatoryCategory
pageDetail.SecuredSettlement = rD.SecuredSettlement
pageDetail.SettlementLimitCarveOut = rD.SettlementLimitCarveOut
pageDetail.SortCode = rD.SortCode
pageDetail.Training = rD.Training
pageDetail.TrainingCode = rD.TrainingCode
pageDetail.TrainingReceived = rD.TrainingReceived
pageDetail.Unencumbered = rD.Unencumbered
pageDetail.LEIExpiryDate = rD.LEIExpiryDate
pageDetail.MIFIDReviewDate = rD.MIFIDReviewDate
pageDetail.GDPRReviewDate = rD.GDPRReviewDate
pageDetail.DelegatedReporting = rD.DelegatedReporting
pageDetail.BOReconcile = rD.BOReconcile
pageDetail.MIFIDReportableDealsAllowed = rD.MIFIDReportableDealsAllowed
pageDetail.SignedInvestmentAgreement = rD.SignedInvestmentAgreement
pageDetail.AppropriatenessAssessment = rD.AppropriatenessAssessment
pageDetail.FinancialCounterparty = rD.FinancialCounterparty
pageDetail.Collateralisation = rD.Collateralisation
pageDetail.PortfolioCode = rD.PortfolioCode
pageDetail.ReconciliationLetterFrequency = rD.ReconciliationLetterFrequency
pageDetail.DirectDealing = rD.DirectDealing
pageDetail.CompID = rD.CompID


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.CounterpartyExtensions_TemplateView, w, r, pageDetail)

}

//CounterpartyExtensions_HandlerEdit is the handler used generate the Edit page
func CounterpartyExtensions_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyExtensions_QueryString)
	_, rD, _ := dao.CounterpartyExtensions_GetByID(searchID)
	
	pageDetail := CounterpartyExtensions_Page{
		Title:       CardTitle(dm.CounterpartyExtensions_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.CounterpartyExtensions_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.NameFirm = rD.NameFirm
pageDetail.NameCentre = rD.NameCentre
pageDetail.BICCode = rD.BICCode
pageDetail.ContactIndicator = rD.ContactIndicator
pageDetail.CoverTrade = rD.CoverTrade
pageDetail.CustomerCategory = rD.CustomerCategory
pageDetail.FSCSInclusive = rD.FSCSInclusive
pageDetail.FeeFactor = rD.FeeFactor
pageDetail.InactiveStatus = rD.InactiveStatus
pageDetail.Indemnity = rD.Indemnity
pageDetail.KnowYourCustomerStatus = rD.KnowYourCustomerStatus
pageDetail.LERLimitCarveOut = rD.LERLimitCarveOut
pageDetail.LastAmended = rD.LastAmended
pageDetail.LastLogin = rD.LastLogin
pageDetail.LossGivenDefault = rD.LossGivenDefault
pageDetail.MIC = rD.MIC
pageDetail.ProtectedDepositor = rD.ProtectedDepositor
pageDetail.RPTCurrency = rD.RPTCurrency
pageDetail.RateTimeout = rD.RateTimeout
pageDetail.RateValidation = rD.RateValidation
pageDetail.Registered = rD.Registered
pageDetail.RegulatoryCategory = rD.RegulatoryCategory
pageDetail.SecuredSettlement = rD.SecuredSettlement
pageDetail.SettlementLimitCarveOut = rD.SettlementLimitCarveOut
pageDetail.SortCode = rD.SortCode
pageDetail.Training = rD.Training
pageDetail.TrainingCode = rD.TrainingCode
pageDetail.TrainingReceived = rD.TrainingReceived
pageDetail.Unencumbered = rD.Unencumbered
pageDetail.LEIExpiryDate = rD.LEIExpiryDate
pageDetail.MIFIDReviewDate = rD.MIFIDReviewDate
pageDetail.GDPRReviewDate = rD.GDPRReviewDate
pageDetail.DelegatedReporting = rD.DelegatedReporting
pageDetail.BOReconcile = rD.BOReconcile
pageDetail.MIFIDReportableDealsAllowed = rD.MIFIDReportableDealsAllowed
pageDetail.SignedInvestmentAgreement = rD.SignedInvestmentAgreement
pageDetail.AppropriatenessAssessment = rD.AppropriatenessAssessment
pageDetail.FinancialCounterparty = rD.FinancialCounterparty
pageDetail.Collateralisation = rD.Collateralisation
pageDetail.PortfolioCode = rD.PortfolioCode
pageDetail.ReconciliationLetterFrequency = rD.ReconciliationLetterFrequency
pageDetail.DirectDealing = rD.DirectDealing
pageDetail.CompID = rD.CompID


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.CounterpartyExtensions_TemplateEdit, w, r, pageDetail)


}

//CounterpartyExtensions_HandlerSave is the handler used process the saving of an CounterpartyExtensions
func CounterpartyExtensions_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("CompID"))

	var item dm.CounterpartyExtensions
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.NameFirm = r.FormValue(dm.CounterpartyExtensions_NameFirm)
		item.NameCentre = r.FormValue(dm.CounterpartyExtensions_NameCentre)
		item.BICCode = r.FormValue(dm.CounterpartyExtensions_BICCode)
		item.ContactIndicator = r.FormValue(dm.CounterpartyExtensions_ContactIndicator)
		item.CoverTrade = r.FormValue(dm.CounterpartyExtensions_CoverTrade)
		item.CustomerCategory = r.FormValue(dm.CounterpartyExtensions_CustomerCategory)
		item.FSCSInclusive = r.FormValue(dm.CounterpartyExtensions_FSCSInclusive)
		item.FeeFactor = r.FormValue(dm.CounterpartyExtensions_FeeFactor)
		item.InactiveStatus = r.FormValue(dm.CounterpartyExtensions_InactiveStatus)
		item.Indemnity = r.FormValue(dm.CounterpartyExtensions_Indemnity)
		item.KnowYourCustomerStatus = r.FormValue(dm.CounterpartyExtensions_KnowYourCustomerStatus)
		item.LERLimitCarveOut = r.FormValue(dm.CounterpartyExtensions_LERLimitCarveOut)
		item.LastAmended = r.FormValue(dm.CounterpartyExtensions_LastAmended)
		item.LastLogin = r.FormValue(dm.CounterpartyExtensions_LastLogin)
		item.LossGivenDefault = r.FormValue(dm.CounterpartyExtensions_LossGivenDefault)
		item.MIC = r.FormValue(dm.CounterpartyExtensions_MIC)
		item.ProtectedDepositor = r.FormValue(dm.CounterpartyExtensions_ProtectedDepositor)
		item.RPTCurrency = r.FormValue(dm.CounterpartyExtensions_RPTCurrency)
		item.RateTimeout = r.FormValue(dm.CounterpartyExtensions_RateTimeout)
		item.RateValidation = r.FormValue(dm.CounterpartyExtensions_RateValidation)
		item.Registered = r.FormValue(dm.CounterpartyExtensions_Registered)
		item.RegulatoryCategory = r.FormValue(dm.CounterpartyExtensions_RegulatoryCategory)
		item.SecuredSettlement = r.FormValue(dm.CounterpartyExtensions_SecuredSettlement)
		item.SettlementLimitCarveOut = r.FormValue(dm.CounterpartyExtensions_SettlementLimitCarveOut)
		item.SortCode = r.FormValue(dm.CounterpartyExtensions_SortCode)
		item.Training = r.FormValue(dm.CounterpartyExtensions_Training)
		item.TrainingCode = r.FormValue(dm.CounterpartyExtensions_TrainingCode)
		item.TrainingReceived = r.FormValue(dm.CounterpartyExtensions_TrainingReceived)
		item.Unencumbered = r.FormValue(dm.CounterpartyExtensions_Unencumbered)
		item.LEIExpiryDate = r.FormValue(dm.CounterpartyExtensions_LEIExpiryDate)
		item.MIFIDReviewDate = r.FormValue(dm.CounterpartyExtensions_MIFIDReviewDate)
		item.GDPRReviewDate = r.FormValue(dm.CounterpartyExtensions_GDPRReviewDate)
		item.DelegatedReporting = r.FormValue(dm.CounterpartyExtensions_DelegatedReporting)
		item.BOReconcile = r.FormValue(dm.CounterpartyExtensions_BOReconcile)
		item.MIFIDReportableDealsAllowed = r.FormValue(dm.CounterpartyExtensions_MIFIDReportableDealsAllowed)
		item.SignedInvestmentAgreement = r.FormValue(dm.CounterpartyExtensions_SignedInvestmentAgreement)
		item.AppropriatenessAssessment = r.FormValue(dm.CounterpartyExtensions_AppropriatenessAssessment)
		item.FinancialCounterparty = r.FormValue(dm.CounterpartyExtensions_FinancialCounterparty)
		item.Collateralisation = r.FormValue(dm.CounterpartyExtensions_Collateralisation)
		item.PortfolioCode = r.FormValue(dm.CounterpartyExtensions_PortfolioCode)
		item.ReconciliationLetterFrequency = r.FormValue(dm.CounterpartyExtensions_ReconciliationLetterFrequency)
		item.DirectDealing = r.FormValue(dm.CounterpartyExtensions_DirectDealing)
		item.CompID = r.FormValue(dm.CounterpartyExtensions_CompID)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.CounterpartyExtensions_Store(item,r)	

	http.Redirect(w, r, CounterpartyExtensions_Redirect, http.StatusFound)
}

//CounterpartyExtensions_HandlerNew is the handler used process the creation of an CounterpartyExtensions
func CounterpartyExtensions_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := CounterpartyExtensions_Page{
		Title:       CardTitle(dm.CounterpartyExtensions_Title, core.Action_New),
		PageTitle:   PageTitle(dm.CounterpartyExtensions_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.NameFirm = ""
pageDetail.NameCentre = ""
pageDetail.BICCode = ""
pageDetail.ContactIndicator = ""
pageDetail.CoverTrade = ""
pageDetail.CustomerCategory = ""
pageDetail.FSCSInclusive = ""
pageDetail.FeeFactor = ""
pageDetail.InactiveStatus = ""
pageDetail.Indemnity = ""
pageDetail.KnowYourCustomerStatus = ""
pageDetail.LERLimitCarveOut = ""
pageDetail.LastAmended = ""
pageDetail.LastLogin = ""
pageDetail.LossGivenDefault = ""
pageDetail.MIC = ""
pageDetail.ProtectedDepositor = ""
pageDetail.RPTCurrency = ""
pageDetail.RateTimeout = ""
pageDetail.RateValidation = ""
pageDetail.Registered = ""
pageDetail.RegulatoryCategory = ""
pageDetail.SecuredSettlement = ""
pageDetail.SettlementLimitCarveOut = ""
pageDetail.SortCode = ""
pageDetail.Training = ""
pageDetail.TrainingCode = ""
pageDetail.TrainingReceived = ""
pageDetail.Unencumbered = ""
pageDetail.LEIExpiryDate = ""
pageDetail.MIFIDReviewDate = ""
pageDetail.GDPRReviewDate = ""
pageDetail.DelegatedReporting = ""
pageDetail.BOReconcile = ""
pageDetail.MIFIDReportableDealsAllowed = ""
pageDetail.SignedInvestmentAgreement = ""
pageDetail.AppropriatenessAssessment = ""
pageDetail.FinancialCounterparty = ""
pageDetail.Collateralisation = ""
pageDetail.PortfolioCode = ""
pageDetail.ReconciliationLetterFrequency = ""
pageDetail.DirectDealing = ""
pageDetail.CompID = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.CounterpartyExtensions_TemplateNew, w, r, pageDetail)

}

//CounterpartyExtensions_HandlerDelete is the handler used process the deletion of an CounterpartyExtensions
func CounterpartyExtensions_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.CounterpartyExtensions_QueryString)

	dao.CounterpartyExtensions_Delete(searchID)	

	http.Redirect(w, r, CounterpartyExtensions_Redirect, http.StatusFound)
}