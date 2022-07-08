package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/counterpartyextensions.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyExtensions (counterpartyextensions)
// Endpoint 	        : CounterpartyExtensions (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:47
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CounterpartyExtensions defines the datamolde for the CounterpartyExtensions object
type CounterpartyExtensions struct {


NameFirm       string
NameFirm_props FieldProperties
NameCentre       string
NameCentre_props FieldProperties
BICCode       string
BICCode_props FieldProperties
ContactIndicator       string
ContactIndicator_props FieldProperties
CoverTrade       string
CoverTrade_props FieldProperties
CustomerCategory       string
CustomerCategory_props FieldProperties
FSCSInclusive       string
FSCSInclusive_props FieldProperties
FeeFactor       string
FeeFactor_props FieldProperties
InactiveStatus       string
InactiveStatus_props FieldProperties
Indemnity       string
Indemnity_props FieldProperties
KnowYourCustomerStatus       string
KnowYourCustomerStatus_props FieldProperties
LERLimitCarveOut       string
LERLimitCarveOut_props FieldProperties
LastAmended       string
LastAmended_props FieldProperties
LastLogin       string
LastLogin_props FieldProperties
LossGivenDefault       string
LossGivenDefault_props FieldProperties
MIC       string
MIC_props FieldProperties
ProtectedDepositor       string
ProtectedDepositor_props FieldProperties
RPTCurrency       string
RPTCurrency_props FieldProperties
RateTimeout       string
RateTimeout_props FieldProperties
RateValidation       string
RateValidation_props FieldProperties
Registered       string
Registered_props FieldProperties
RegulatoryCategory       string
RegulatoryCategory_props FieldProperties
SecuredSettlement       string
SecuredSettlement_props FieldProperties
SettlementLimitCarveOut       string
SettlementLimitCarveOut_props FieldProperties
SortCode       string
SortCode_props FieldProperties
Training       string
Training_props FieldProperties
TrainingCode       string
TrainingCode_props FieldProperties
TrainingReceived       string
TrainingReceived_props FieldProperties
Unencumbered       string
Unencumbered_props FieldProperties
LEIExpiryDate       string
LEIExpiryDate_props FieldProperties
MIFIDReviewDate       string
MIFIDReviewDate_props FieldProperties
GDPRReviewDate       string
GDPRReviewDate_props FieldProperties
DelegatedReporting       string
DelegatedReporting_props FieldProperties
BOReconcile       string
BOReconcile_props FieldProperties
MIFIDReportableDealsAllowed       string
MIFIDReportableDealsAllowed_props FieldProperties
SignedInvestmentAgreement       string
SignedInvestmentAgreement_props FieldProperties
AppropriatenessAssessment       string
AppropriatenessAssessment_props FieldProperties
FinancialCounterparty       string
FinancialCounterparty_props FieldProperties
Collateralisation       string
Collateralisation_props FieldProperties
PortfolioCode       string
PortfolioCode_props FieldProperties
ReconciliationLetterFrequency       string
ReconciliationLetterFrequency_props FieldProperties
DirectDealing       string
DirectDealing_props FieldProperties
CompID       string
CompID_props FieldProperties
 // Any lookups will be added below











































}

const (
	CounterpartyExtensions_Title       = "Counterparty Extension"
	CounterpartyExtensions_SQLTable    = "sienaCounterpartyExtensions"
	CounterpartyExtensions_SQLSearchID = "CompID"
	CounterpartyExtensions_QueryString = "ID"
	///
	/// Handler Defintions
	///
	CounterpartyExtensions_Template     = "CounterpartyExtensions"
	CounterpartyExtensions_TemplateList = "/CounterpartyExtensions/CounterpartyExtensions_List"
	CounterpartyExtensions_TemplateView = "/CounterpartyExtensions/CounterpartyExtensions_View"
	CounterpartyExtensions_TemplateEdit = "/CounterpartyExtensions/CounterpartyExtensions_Edit"
	CounterpartyExtensions_TemplateNew  = "/CounterpartyExtensions/CounterpartyExtensions_New"
	///
	/// Handler Monitor Paths
	///
	CounterpartyExtensions_Path       = "/API/CounterpartyExtensions/"
	CounterpartyExtensions_PathList   = "/CounterpartyExtensionsList/"
	CounterpartyExtensions_PathView   = "/CounterpartyExtensionsView/"
	CounterpartyExtensions_PathEdit   = "/CounterpartyExtensionsEdit/"
	CounterpartyExtensions_PathNew    = "/CounterpartyExtensionsNew/"
	CounterpartyExtensions_PathSave   = "/CounterpartyExtensionsSave/"
	CounterpartyExtensions_PathDelete = "/CounterpartyExtensionsDelete/"
	///
	//CounterpartyExtensions_Redirect provides a page to return to aftern an action
	CounterpartyExtensions_Redirect = CounterpartyExtensions_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	CounterpartyExtensions_NameFirm_sql   = "NameFirm" // NameFirm is a String
	CounterpartyExtensions_NameCentre_sql   = "NameCentre" // NameCentre is a String
	CounterpartyExtensions_BICCode_sql   = "BICCode" // BICCode is a String
	CounterpartyExtensions_ContactIndicator_sql   = "ContactIndicator" // ContactIndicator is a Bool
	CounterpartyExtensions_CoverTrade_sql   = "CoverTrade" // CoverTrade is a Bool
	CounterpartyExtensions_CustomerCategory_sql   = "CustomerCategory" // CustomerCategory is a String
	CounterpartyExtensions_FSCSInclusive_sql   = "FSCSInclusive" // FSCSInclusive is a Bool
	CounterpartyExtensions_FeeFactor_sql   = "FeeFactor" // FeeFactor is a Float
	CounterpartyExtensions_InactiveStatus_sql   = "InactiveStatus" // InactiveStatus is a Bool
	CounterpartyExtensions_Indemnity_sql   = "Indemnity" // Indemnity is a Bool
	CounterpartyExtensions_KnowYourCustomerStatus_sql   = "KnowYourCustomerStatus" // KnowYourCustomerStatus is a Bool
	CounterpartyExtensions_LERLimitCarveOut_sql   = "LERLimitCarveOut" // LERLimitCarveOut is a Bool
	CounterpartyExtensions_LastAmended_sql   = "LastAmended" // LastAmended is a Time
	CounterpartyExtensions_LastLogin_sql   = "LastLogin" // LastLogin is a Time
	CounterpartyExtensions_LossGivenDefault_sql   = "LossGivenDefault" // LossGivenDefault is a Float
	CounterpartyExtensions_MIC_sql   = "MIC" // MIC is a String
	CounterpartyExtensions_ProtectedDepositor_sql   = "ProtectedDepositor" // ProtectedDepositor is a Bool
	CounterpartyExtensions_RPTCurrency_sql   = "RPTCurrency" // RPTCurrency is a String
	CounterpartyExtensions_RateTimeout_sql   = "RateTimeout" // RateTimeout is a Int
	CounterpartyExtensions_RateValidation_sql   = "RateValidation" // RateValidation is a String
	CounterpartyExtensions_Registered_sql   = "Registered" // Registered is a Bool
	CounterpartyExtensions_RegulatoryCategory_sql   = "RegulatoryCategory" // RegulatoryCategory is a String
	CounterpartyExtensions_SecuredSettlement_sql   = "SecuredSettlement" // SecuredSettlement is a Bool
	CounterpartyExtensions_SettlementLimitCarveOut_sql   = "SettlementLimitCarveOut" // SettlementLimitCarveOut is a Bool
	CounterpartyExtensions_SortCode_sql   = "SortCode" // SortCode is a String
	CounterpartyExtensions_Training_sql   = "Training" // Training is a Bool
	CounterpartyExtensions_TrainingCode_sql   = "TrainingCode" // TrainingCode is a String
	CounterpartyExtensions_TrainingReceived_sql   = "TrainingReceived" // TrainingReceived is a Bool
	CounterpartyExtensions_Unencumbered_sql   = "Unencumbered" // Unencumbered is a Bool
	CounterpartyExtensions_LEIExpiryDate_sql   = "LEIExpiryDate" // LEIExpiryDate is a Time
	CounterpartyExtensions_MIFIDReviewDate_sql   = "MIFIDReviewDate" // MIFIDReviewDate is a Time
	CounterpartyExtensions_GDPRReviewDate_sql   = "GDPRReviewDate" // GDPRReviewDate is a Time
	CounterpartyExtensions_DelegatedReporting_sql   = "DelegatedReporting" // DelegatedReporting is a String
	CounterpartyExtensions_BOReconcile_sql   = "BOReconcile" // BOReconcile is a Bool
	CounterpartyExtensions_MIFIDReportableDealsAllowed_sql   = "MIFIDReportableDealsAllowed" // MIFIDReportableDealsAllowed is a Bool
	CounterpartyExtensions_SignedInvestmentAgreement_sql   = "SignedInvestmentAgreement" // SignedInvestmentAgreement is a Bool
	CounterpartyExtensions_AppropriatenessAssessment_sql   = "AppropriatenessAssessment" // AppropriatenessAssessment is a Bool
	CounterpartyExtensions_FinancialCounterparty_sql   = "FinancialCounterparty" // FinancialCounterparty is a Bool
	CounterpartyExtensions_Collateralisation_sql   = "Collateralisation" // Collateralisation is a String
	CounterpartyExtensions_PortfolioCode_sql   = "PortfolioCode" // PortfolioCode is a String
	CounterpartyExtensions_ReconciliationLetterFrequency_sql   = "ReconciliationLetterFrequency" // ReconciliationLetterFrequency is a String
	CounterpartyExtensions_DirectDealing_sql   = "DirectDealing" // DirectDealing is a Bool
	CounterpartyExtensions_CompID_sql   = "CompID" // CompID is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	CounterpartyExtensions_NameFirm_scrn   = "NameFirm" // NameFirm is a String
	CounterpartyExtensions_NameCentre_scrn   = "NameCentre" // NameCentre is a String
	CounterpartyExtensions_BICCode_scrn   = "BICCode" // BICCode is a String
	CounterpartyExtensions_ContactIndicator_scrn   = "ContactIndicator" // ContactIndicator is a Bool
	CounterpartyExtensions_CoverTrade_scrn   = "CoverTrade" // CoverTrade is a Bool
	CounterpartyExtensions_CustomerCategory_scrn   = "CustomerCategory" // CustomerCategory is a String
	CounterpartyExtensions_FSCSInclusive_scrn   = "FSCSInclusive" // FSCSInclusive is a Bool
	CounterpartyExtensions_FeeFactor_scrn   = "FeeFactor" // FeeFactor is a Float
	CounterpartyExtensions_InactiveStatus_scrn   = "InactiveStatus" // InactiveStatus is a Bool
	CounterpartyExtensions_Indemnity_scrn   = "Indemnity" // Indemnity is a Bool
	CounterpartyExtensions_KnowYourCustomerStatus_scrn   = "KnowYourCustomerStatus" // KnowYourCustomerStatus is a Bool
	CounterpartyExtensions_LERLimitCarveOut_scrn   = "LERLimitCarveOut" // LERLimitCarveOut is a Bool
	CounterpartyExtensions_LastAmended_scrn   = "LastAmended" // LastAmended is a Time
	CounterpartyExtensions_LastLogin_scrn   = "LastLogin" // LastLogin is a Time
	CounterpartyExtensions_LossGivenDefault_scrn   = "LossGivenDefault" // LossGivenDefault is a Float
	CounterpartyExtensions_MIC_scrn   = "MIC" // MIC is a String
	CounterpartyExtensions_ProtectedDepositor_scrn   = "ProtectedDepositor" // ProtectedDepositor is a Bool
	CounterpartyExtensions_RPTCurrency_scrn   = "RPTCurrency" // RPTCurrency is a String
	CounterpartyExtensions_RateTimeout_scrn   = "RateTimeout" // RateTimeout is a Int
	CounterpartyExtensions_RateValidation_scrn   = "RateValidation" // RateValidation is a String
	CounterpartyExtensions_Registered_scrn   = "Registered" // Registered is a Bool
	CounterpartyExtensions_RegulatoryCategory_scrn   = "RegulatoryCategory" // RegulatoryCategory is a String
	CounterpartyExtensions_SecuredSettlement_scrn   = "SecuredSettlement" // SecuredSettlement is a Bool
	CounterpartyExtensions_SettlementLimitCarveOut_scrn   = "SettlementLimitCarveOut" // SettlementLimitCarveOut is a Bool
	CounterpartyExtensions_SortCode_scrn   = "SortCode" // SortCode is a String
	CounterpartyExtensions_Training_scrn   = "Training" // Training is a Bool
	CounterpartyExtensions_TrainingCode_scrn   = "TrainingCode" // TrainingCode is a String
	CounterpartyExtensions_TrainingReceived_scrn   = "TrainingReceived" // TrainingReceived is a Bool
	CounterpartyExtensions_Unencumbered_scrn   = "Unencumbered" // Unencumbered is a Bool
	CounterpartyExtensions_LEIExpiryDate_scrn   = "LEIExpiryDate" // LEIExpiryDate is a Time
	CounterpartyExtensions_MIFIDReviewDate_scrn   = "MIFIDReviewDate" // MIFIDReviewDate is a Time
	CounterpartyExtensions_GDPRReviewDate_scrn   = "GDPRReviewDate" // GDPRReviewDate is a Time
	CounterpartyExtensions_DelegatedReporting_scrn   = "DelegatedReporting" // DelegatedReporting is a String
	CounterpartyExtensions_BOReconcile_scrn   = "BOReconcile" // BOReconcile is a Bool
	CounterpartyExtensions_MIFIDReportableDealsAllowed_scrn   = "MIFIDReportableDealsAllowed" // MIFIDReportableDealsAllowed is a Bool
	CounterpartyExtensions_SignedInvestmentAgreement_scrn   = "SignedInvestmentAgreement" // SignedInvestmentAgreement is a Bool
	CounterpartyExtensions_AppropriatenessAssessment_scrn   = "AppropriatenessAssessment" // AppropriatenessAssessment is a Bool
	CounterpartyExtensions_FinancialCounterparty_scrn   = "FinancialCounterparty" // FinancialCounterparty is a Bool
	CounterpartyExtensions_Collateralisation_scrn   = "Collateralisation" // Collateralisation is a String
	CounterpartyExtensions_PortfolioCode_scrn   = "PortfolioCode" // PortfolioCode is a String
	CounterpartyExtensions_ReconciliationLetterFrequency_scrn   = "ReconciliationLetterFrequency" // ReconciliationLetterFrequency is a String
	CounterpartyExtensions_DirectDealing_scrn   = "DirectDealing" // DirectDealing is a Bool
	CounterpartyExtensions_CompID_scrn   = "CompID" // CompID is a String

	/// Definitions End
	///
)

//counterpartyextensions_PageList provides the information for the template for a list of CounterpartyExtensionss
type CounterpartyExtensions_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []CounterpartyExtensions
	Context	 appContext
}

//counterpartyextensions_Page provides the information for the template for an individual CounterpartyExtensions
type CounterpartyExtensions_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	NameFirm         string
	NameFirm_props     FieldProperties
	NameCentre         string
	NameCentre_props     FieldProperties
	BICCode         string
	BICCode_props     FieldProperties
	ContactIndicator         string
	ContactIndicator_props     FieldProperties
	CoverTrade         string
	CoverTrade_props     FieldProperties
	CustomerCategory         string
	CustomerCategory_props     FieldProperties
	FSCSInclusive         string
	FSCSInclusive_props     FieldProperties
	FeeFactor         string
	FeeFactor_props     FieldProperties
	InactiveStatus         string
	InactiveStatus_props     FieldProperties
	Indemnity         string
	Indemnity_props     FieldProperties
	KnowYourCustomerStatus         string
	KnowYourCustomerStatus_props     FieldProperties
	LERLimitCarveOut         string
	LERLimitCarveOut_props     FieldProperties
	LastAmended         string
	LastAmended_props     FieldProperties
	LastLogin         string
	LastLogin_props     FieldProperties
	LossGivenDefault         string
	LossGivenDefault_props     FieldProperties
	MIC         string
	MIC_props     FieldProperties
	ProtectedDepositor         string
	ProtectedDepositor_props     FieldProperties
	RPTCurrency         string
	RPTCurrency_props     FieldProperties
	RateTimeout         string
	RateTimeout_props     FieldProperties
	RateValidation         string
	RateValidation_props     FieldProperties
	Registered         string
	Registered_props     FieldProperties
	RegulatoryCategory         string
	RegulatoryCategory_props     FieldProperties
	SecuredSettlement         string
	SecuredSettlement_props     FieldProperties
	SettlementLimitCarveOut         string
	SettlementLimitCarveOut_props     FieldProperties
	SortCode         string
	SortCode_props     FieldProperties
	Training         string
	Training_props     FieldProperties
	TrainingCode         string
	TrainingCode_props     FieldProperties
	TrainingReceived         string
	TrainingReceived_props     FieldProperties
	Unencumbered         string
	Unencumbered_props     FieldProperties
	LEIExpiryDate         string
	LEIExpiryDate_props     FieldProperties
	MIFIDReviewDate         string
	MIFIDReviewDate_props     FieldProperties
	GDPRReviewDate         string
	GDPRReviewDate_props     FieldProperties
	DelegatedReporting         string
	DelegatedReporting_props     FieldProperties
	BOReconcile         string
	BOReconcile_props     FieldProperties
	MIFIDReportableDealsAllowed         string
	MIFIDReportableDealsAllowed_props     FieldProperties
	SignedInvestmentAgreement         string
	SignedInvestmentAgreement_props     FieldProperties
	AppropriatenessAssessment         string
	AppropriatenessAssessment_props     FieldProperties
	FinancialCounterparty         string
	FinancialCounterparty_props     FieldProperties
	Collateralisation         string
	Collateralisation_props     FieldProperties
	PortfolioCode         string
	PortfolioCode_props     FieldProperties
	ReconciliationLetterFrequency         string
	ReconciliationLetterFrequency_props     FieldProperties
	DirectDealing         string
	DirectDealing_props     FieldProperties
	CompID         string
	CompID_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}