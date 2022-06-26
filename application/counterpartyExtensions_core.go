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
// Date & Time		    : 26/06/2022 at 18:48:22
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
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
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.CounterpartyExtensions
}
//CounterpartyExtensions_Redirect provides a page to return to aftern an action
const (
	
	CounterpartyExtensions_Redirect = dm.CounterpartyExtensions_PathList
	
)

//counterpartyextensions_Page provides the information for the template for an individual CounterpartyExtensions
type CounterpartyExtensions_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	NameFirm         string
	NameFirm_props     dm.FieldProperties
	NameCentre         string
	NameCentre_props     dm.FieldProperties
	BICCode         string
	BICCode_props     dm.FieldProperties
	ContactIndicator         string
	ContactIndicator_props     dm.FieldProperties
	CoverTrade         string
	CoverTrade_props     dm.FieldProperties
	CustomerCategory         string
	CustomerCategory_props     dm.FieldProperties
	FSCSInclusive         string
	FSCSInclusive_props     dm.FieldProperties
	FeeFactor         string
	FeeFactor_props     dm.FieldProperties
	InactiveStatus         string
	InactiveStatus_props     dm.FieldProperties
	Indemnity         string
	Indemnity_props     dm.FieldProperties
	KnowYourCustomerStatus         string
	KnowYourCustomerStatus_props     dm.FieldProperties
	LERLimitCarveOut         string
	LERLimitCarveOut_props     dm.FieldProperties
	LastAmended         string
	LastAmended_props     dm.FieldProperties
	LastLogin         string
	LastLogin_props     dm.FieldProperties
	LossGivenDefault         string
	LossGivenDefault_props     dm.FieldProperties
	MIC         string
	MIC_props     dm.FieldProperties
	ProtectedDepositor         string
	ProtectedDepositor_props     dm.FieldProperties
	RPTCurrency         string
	RPTCurrency_props     dm.FieldProperties
	RateTimeout         string
	RateTimeout_props     dm.FieldProperties
	RateValidation         string
	RateValidation_props     dm.FieldProperties
	Registered         string
	Registered_props     dm.FieldProperties
	RegulatoryCategory         string
	RegulatoryCategory_props     dm.FieldProperties
	SecuredSettlement         string
	SecuredSettlement_props     dm.FieldProperties
	SettlementLimitCarveOut         string
	SettlementLimitCarveOut_props     dm.FieldProperties
	SortCode         string
	SortCode_props     dm.FieldProperties
	Training         string
	Training_props     dm.FieldProperties
	TrainingCode         string
	TrainingCode_props     dm.FieldProperties
	TrainingReceived         string
	TrainingReceived_props     dm.FieldProperties
	Unencumbered         string
	Unencumbered_props     dm.FieldProperties
	LEIExpiryDate         string
	LEIExpiryDate_props     dm.FieldProperties
	MIFIDReviewDate         string
	MIFIDReviewDate_props     dm.FieldProperties
	GDPRReviewDate         string
	GDPRReviewDate_props     dm.FieldProperties
	DelegatedReporting         string
	DelegatedReporting_props     dm.FieldProperties
	BOReconcile         string
	BOReconcile_props     dm.FieldProperties
	MIFIDReportableDealsAllowed         string
	MIFIDReportableDealsAllowed_props     dm.FieldProperties
	SignedInvestmentAgreement         string
	SignedInvestmentAgreement_props     dm.FieldProperties
	AppropriatenessAssessment         string
	AppropriatenessAssessment_props     dm.FieldProperties
	FinancialCounterparty         string
	FinancialCounterparty_props     dm.FieldProperties
	Collateralisation         string
	Collateralisation_props     dm.FieldProperties
	PortfolioCode         string
	PortfolioCode_props     dm.FieldProperties
	ReconciliationLetterFrequency         string
	ReconciliationLetterFrequency_props     dm.FieldProperties
	DirectDealing         string
	DirectDealing_props     dm.FieldProperties
	CompID         string
	CompID_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//CounterpartyExtensions_Publish annouces the endpoints available for this object
func CounterpartyExtensions_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.CounterpartyExtensions_PathList, CounterpartyExtensions_HandlerList)
	mux.HandleFunc(dm.CounterpartyExtensions_PathView, CounterpartyExtensions_HandlerView)
	mux.HandleFunc(dm.CounterpartyExtensions_PathEdit, CounterpartyExtensions_HandlerEdit)
	mux.HandleFunc(dm.CounterpartyExtensions_PathNew, CounterpartyExtensions_HandlerNew)
	mux.HandleFunc(dm.CounterpartyExtensions_PathSave, CounterpartyExtensions_HandlerSave)
	mux.HandleFunc(dm.CounterpartyExtensions_PathDelete, CounterpartyExtensions_HandlerDelete)
	logs.Publish("Application", dm.CounterpartyExtensions_Title)
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

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterpartyextensions_PopulatePage(rD , pageDetail) 

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

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterpartyextensions_PopulatePage(rD , pageDetail) 

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
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.NameFirm = r.FormValue(dm.CounterpartyExtensions_NameFirm_scrn)
		item.NameCentre = r.FormValue(dm.CounterpartyExtensions_NameCentre_scrn)
		item.BICCode = r.FormValue(dm.CounterpartyExtensions_BICCode_scrn)
		item.ContactIndicator = r.FormValue(dm.CounterpartyExtensions_ContactIndicator_scrn)
		item.CoverTrade = r.FormValue(dm.CounterpartyExtensions_CoverTrade_scrn)
		item.CustomerCategory = r.FormValue(dm.CounterpartyExtensions_CustomerCategory_scrn)
		item.FSCSInclusive = r.FormValue(dm.CounterpartyExtensions_FSCSInclusive_scrn)
		item.FeeFactor = r.FormValue(dm.CounterpartyExtensions_FeeFactor_scrn)
		item.InactiveStatus = r.FormValue(dm.CounterpartyExtensions_InactiveStatus_scrn)
		item.Indemnity = r.FormValue(dm.CounterpartyExtensions_Indemnity_scrn)
		item.KnowYourCustomerStatus = r.FormValue(dm.CounterpartyExtensions_KnowYourCustomerStatus_scrn)
		item.LERLimitCarveOut = r.FormValue(dm.CounterpartyExtensions_LERLimitCarveOut_scrn)
		item.LastAmended = r.FormValue(dm.CounterpartyExtensions_LastAmended_scrn)
		item.LastLogin = r.FormValue(dm.CounterpartyExtensions_LastLogin_scrn)
		item.LossGivenDefault = r.FormValue(dm.CounterpartyExtensions_LossGivenDefault_scrn)
		item.MIC = r.FormValue(dm.CounterpartyExtensions_MIC_scrn)
		item.ProtectedDepositor = r.FormValue(dm.CounterpartyExtensions_ProtectedDepositor_scrn)
		item.RPTCurrency = r.FormValue(dm.CounterpartyExtensions_RPTCurrency_scrn)
		item.RateTimeout = r.FormValue(dm.CounterpartyExtensions_RateTimeout_scrn)
		item.RateValidation = r.FormValue(dm.CounterpartyExtensions_RateValidation_scrn)
		item.Registered = r.FormValue(dm.CounterpartyExtensions_Registered_scrn)
		item.RegulatoryCategory = r.FormValue(dm.CounterpartyExtensions_RegulatoryCategory_scrn)
		item.SecuredSettlement = r.FormValue(dm.CounterpartyExtensions_SecuredSettlement_scrn)
		item.SettlementLimitCarveOut = r.FormValue(dm.CounterpartyExtensions_SettlementLimitCarveOut_scrn)
		item.SortCode = r.FormValue(dm.CounterpartyExtensions_SortCode_scrn)
		item.Training = r.FormValue(dm.CounterpartyExtensions_Training_scrn)
		item.TrainingCode = r.FormValue(dm.CounterpartyExtensions_TrainingCode_scrn)
		item.TrainingReceived = r.FormValue(dm.CounterpartyExtensions_TrainingReceived_scrn)
		item.Unencumbered = r.FormValue(dm.CounterpartyExtensions_Unencumbered_scrn)
		item.LEIExpiryDate = r.FormValue(dm.CounterpartyExtensions_LEIExpiryDate_scrn)
		item.MIFIDReviewDate = r.FormValue(dm.CounterpartyExtensions_MIFIDReviewDate_scrn)
		item.GDPRReviewDate = r.FormValue(dm.CounterpartyExtensions_GDPRReviewDate_scrn)
		item.DelegatedReporting = r.FormValue(dm.CounterpartyExtensions_DelegatedReporting_scrn)
		item.BOReconcile = r.FormValue(dm.CounterpartyExtensions_BOReconcile_scrn)
		item.MIFIDReportableDealsAllowed = r.FormValue(dm.CounterpartyExtensions_MIFIDReportableDealsAllowed_scrn)
		item.SignedInvestmentAgreement = r.FormValue(dm.CounterpartyExtensions_SignedInvestmentAgreement_scrn)
		item.AppropriatenessAssessment = r.FormValue(dm.CounterpartyExtensions_AppropriatenessAssessment_scrn)
		item.FinancialCounterparty = r.FormValue(dm.CounterpartyExtensions_FinancialCounterparty_scrn)
		item.Collateralisation = r.FormValue(dm.CounterpartyExtensions_Collateralisation_scrn)
		item.PortfolioCode = r.FormValue(dm.CounterpartyExtensions_PortfolioCode_scrn)
		item.ReconciliationLetterFrequency = r.FormValue(dm.CounterpartyExtensions_ReconciliationLetterFrequency_scrn)
		item.DirectDealing = r.FormValue(dm.CounterpartyExtensions_DirectDealing_scrn)
		item.CompID = r.FormValue(dm.CounterpartyExtensions_CompID_scrn)
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
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

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.CounterpartyExtensions_New()

	pageDetail := CounterpartyExtensions_Page{
		Title:       CardTitle(dm.CounterpartyExtensions_Title, core.Action_New),
		PageTitle:   PageTitle(dm.CounterpartyExtensions_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterpartyextensions_PopulatePage(rD , pageDetail) 

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


// Builds/Popuplates the CounterpartyExtensions Page 
func counterpartyextensions_PopulatePage(rD dm.CounterpartyExtensions, pageDetail CounterpartyExtensions_Page) CounterpartyExtensions_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.NameFirm_props = rD.NameFirm_props
	pageDetail.NameCentre_props = rD.NameCentre_props
	pageDetail.BICCode_props = rD.BICCode_props
	pageDetail.ContactIndicator_props = rD.ContactIndicator_props
	pageDetail.CoverTrade_props = rD.CoverTrade_props
	pageDetail.CustomerCategory_props = rD.CustomerCategory_props
	pageDetail.FSCSInclusive_props = rD.FSCSInclusive_props
	pageDetail.FeeFactor_props = rD.FeeFactor_props
	pageDetail.InactiveStatus_props = rD.InactiveStatus_props
	pageDetail.Indemnity_props = rD.Indemnity_props
	pageDetail.KnowYourCustomerStatus_props = rD.KnowYourCustomerStatus_props
	pageDetail.LERLimitCarveOut_props = rD.LERLimitCarveOut_props
	pageDetail.LastAmended_props = rD.LastAmended_props
	pageDetail.LastLogin_props = rD.LastLogin_props
	pageDetail.LossGivenDefault_props = rD.LossGivenDefault_props
	pageDetail.MIC_props = rD.MIC_props
	pageDetail.ProtectedDepositor_props = rD.ProtectedDepositor_props
	pageDetail.RPTCurrency_props = rD.RPTCurrency_props
	pageDetail.RateTimeout_props = rD.RateTimeout_props
	pageDetail.RateValidation_props = rD.RateValidation_props
	pageDetail.Registered_props = rD.Registered_props
	pageDetail.RegulatoryCategory_props = rD.RegulatoryCategory_props
	pageDetail.SecuredSettlement_props = rD.SecuredSettlement_props
	pageDetail.SettlementLimitCarveOut_props = rD.SettlementLimitCarveOut_props
	pageDetail.SortCode_props = rD.SortCode_props
	pageDetail.Training_props = rD.Training_props
	pageDetail.TrainingCode_props = rD.TrainingCode_props
	pageDetail.TrainingReceived_props = rD.TrainingReceived_props
	pageDetail.Unencumbered_props = rD.Unencumbered_props
	pageDetail.LEIExpiryDate_props = rD.LEIExpiryDate_props
	pageDetail.MIFIDReviewDate_props = rD.MIFIDReviewDate_props
	pageDetail.GDPRReviewDate_props = rD.GDPRReviewDate_props
	pageDetail.DelegatedReporting_props = rD.DelegatedReporting_props
	pageDetail.BOReconcile_props = rD.BOReconcile_props
	pageDetail.MIFIDReportableDealsAllowed_props = rD.MIFIDReportableDealsAllowed_props
	pageDetail.SignedInvestmentAgreement_props = rD.SignedInvestmentAgreement_props
	pageDetail.AppropriatenessAssessment_props = rD.AppropriatenessAssessment_props
	pageDetail.FinancialCounterparty_props = rD.FinancialCounterparty_props
	pageDetail.Collateralisation_props = rD.Collateralisation_props
	pageDetail.PortfolioCode_props = rD.PortfolioCode_props
	pageDetail.ReconciliationLetterFrequency_props = rD.ReconciliationLetterFrequency_props
	pageDetail.DirectDealing_props = rD.DirectDealing_props
	pageDetail.CompID_props = rD.CompID_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	