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
// Date & Time		    : 12/12/2021 at 16:13:09
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type CounterpartyExtensions struct {

NameFirm        string
NameCentre        string
BICCode        string
ContactIndicator        string
CoverTrade        string
CustomerCategory        string
FSCSInclusive        string
FeeFactor        string
InactiveStatus        string
Indemnity        string
KnowYourCustomerStatus        string
LERLimitCarveOut        string
LastAmended        string
LastLogin        string
LossGivenDefault        string
MIC        string
ProtectedDepositor        string
RPTCurrency        string
RateTimeout        string
RateValidation        string
Registered        string
RegulatoryCategory        string
SecuredSettlement        string
SettlementLimitCarveOut        string
SortCode        string
Training        string
TrainingCode        string
TrainingReceived        string
Unencumbered        string
LEIExpiryDate        string
MIFIDReviewDate        string
GDPRReviewDate        string
DelegatedReporting        string
BOReconcile        string
MIFIDReportableDealsAllowed        string
SignedInvestmentAgreement        string
AppropriatenessAssessment        string
FinancialCounterparty        string
Collateralisation        string
PortfolioCode        string
ReconciliationLetterFrequency        string
DirectDealing        string
CompID        string

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
	CounterpartyExtensions_TemplateList = "CounterpartyExtensions_List"
	CounterpartyExtensions_TemplateView = "CounterpartyExtensions_View"
	CounterpartyExtensions_TemplateEdit = "CounterpartyExtensions_Edit"
	CounterpartyExtensions_TemplateNew  = "CounterpartyExtensions_New"
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
	/// SQL Field Definitions
	///
	CounterpartyExtensions_NameFirm   = "NameFirm" // NameFirm is a String
	CounterpartyExtensions_NameCentre   = "NameCentre" // NameCentre is a String
	CounterpartyExtensions_BICCode   = "BICCode" // BICCode is a String
	CounterpartyExtensions_ContactIndicator   = "ContactIndicator" // ContactIndicator is a Bool
	CounterpartyExtensions_CoverTrade   = "CoverTrade" // CoverTrade is a Bool
	CounterpartyExtensions_CustomerCategory   = "CustomerCategory" // CustomerCategory is a String
	CounterpartyExtensions_FSCSInclusive   = "FSCSInclusive" // FSCSInclusive is a Bool
	CounterpartyExtensions_FeeFactor   = "FeeFactor" // FeeFactor is a Float
	CounterpartyExtensions_InactiveStatus   = "InactiveStatus" // InactiveStatus is a Bool
	CounterpartyExtensions_Indemnity   = "Indemnity" // Indemnity is a Bool
	CounterpartyExtensions_KnowYourCustomerStatus   = "KnowYourCustomerStatus" // KnowYourCustomerStatus is a Bool
	CounterpartyExtensions_LERLimitCarveOut   = "LERLimitCarveOut" // LERLimitCarveOut is a Bool
	CounterpartyExtensions_LastAmended   = "LastAmended" // LastAmended is a Time
	CounterpartyExtensions_LastLogin   = "LastLogin" // LastLogin is a Time
	CounterpartyExtensions_LossGivenDefault   = "LossGivenDefault" // LossGivenDefault is a Float
	CounterpartyExtensions_MIC   = "MIC" // MIC is a String
	CounterpartyExtensions_ProtectedDepositor   = "ProtectedDepositor" // ProtectedDepositor is a Bool
	CounterpartyExtensions_RPTCurrency   = "RPTCurrency" // RPTCurrency is a String
	CounterpartyExtensions_RateTimeout   = "RateTimeout" // RateTimeout is a Int
	CounterpartyExtensions_RateValidation   = "RateValidation" // RateValidation is a String
	CounterpartyExtensions_Registered   = "Registered" // Registered is a Bool
	CounterpartyExtensions_RegulatoryCategory   = "RegulatoryCategory" // RegulatoryCategory is a String
	CounterpartyExtensions_SecuredSettlement   = "SecuredSettlement" // SecuredSettlement is a Bool
	CounterpartyExtensions_SettlementLimitCarveOut   = "SettlementLimitCarveOut" // SettlementLimitCarveOut is a Bool
	CounterpartyExtensions_SortCode   = "SortCode" // SortCode is a String
	CounterpartyExtensions_Training   = "Training" // Training is a Bool
	CounterpartyExtensions_TrainingCode   = "TrainingCode" // TrainingCode is a String
	CounterpartyExtensions_TrainingReceived   = "TrainingReceived" // TrainingReceived is a Bool
	CounterpartyExtensions_Unencumbered   = "Unencumbered" // Unencumbered is a Bool
	CounterpartyExtensions_LEIExpiryDate   = "LEIExpiryDate" // LEIExpiryDate is a Time
	CounterpartyExtensions_MIFIDReviewDate   = "MIFIDReviewDate" // MIFIDReviewDate is a Time
	CounterpartyExtensions_GDPRReviewDate   = "GDPRReviewDate" // GDPRReviewDate is a Time
	CounterpartyExtensions_DelegatedReporting   = "DelegatedReporting" // DelegatedReporting is a String
	CounterpartyExtensions_BOReconcile   = "BOReconcile" // BOReconcile is a Bool
	CounterpartyExtensions_MIFIDReportableDealsAllowed   = "MIFIDReportableDealsAllowed" // MIFIDReportableDealsAllowed is a Bool
	CounterpartyExtensions_SignedInvestmentAgreement   = "SignedInvestmentAgreement" // SignedInvestmentAgreement is a Bool
	CounterpartyExtensions_AppropriatenessAssessment   = "AppropriatenessAssessment" // AppropriatenessAssessment is a Bool
	CounterpartyExtensions_FinancialCounterparty   = "FinancialCounterparty" // FinancialCounterparty is a Bool
	CounterpartyExtensions_Collateralisation   = "Collateralisation" // Collateralisation is a String
	CounterpartyExtensions_PortfolioCode   = "PortfolioCode" // PortfolioCode is a String
	CounterpartyExtensions_ReconciliationLetterFrequency   = "ReconciliationLetterFrequency" // ReconciliationLetterFrequency is a String
	CounterpartyExtensions_DirectDealing   = "DirectDealing" // DirectDealing is a Bool
	CounterpartyExtensions_CompID   = "CompID" // CompID is a String

	/// Definitions End
)
