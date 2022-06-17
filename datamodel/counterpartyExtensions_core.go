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
// Date & Time		    : 17/06/2022 at 18:38:07
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CounterpartyExtensions defines the datamolde for the CounterpartyExtensions object
type CounterpartyExtensions struct {

NameFirm       string
NameCentre       string
BICCode       string
ContactIndicator       string
CoverTrade       string
CustomerCategory       string
FSCSInclusive       string
FeeFactor       string
InactiveStatus       string
Indemnity       string
KnowYourCustomerStatus       string
LERLimitCarveOut       string
LastAmended       string
LastLogin       string
LossGivenDefault       string
MIC       string
ProtectedDepositor       string
RPTCurrency       string
RateTimeout       string
RateValidation       string
Registered       string
RegulatoryCategory       string
SecuredSettlement       string
SettlementLimitCarveOut       string
SortCode       string
Training       string
TrainingCode       string
TrainingReceived       string
Unencumbered       string
LEIExpiryDate       string
MIFIDReviewDate       string
GDPRReviewDate       string
DelegatedReporting       string
BOReconcile       string
MIFIDReportableDealsAllowed       string
SignedInvestmentAgreement       string
AppropriatenessAssessment       string
FinancialCounterparty       string
Collateralisation       string
PortfolioCode       string
ReconciliationLetterFrequency       string
DirectDealing       string
CompID       string

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
)
