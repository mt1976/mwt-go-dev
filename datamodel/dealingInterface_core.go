package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/dealinginterface.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DealingInterface (dealinginterface)
// Endpoint 	        : DealingInterface (Name)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:10
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//DealingInterface defines the datamolde for the DealingInterface object
type DealingInterface struct {

Name       string
AcceptReducedAmount       string
QuoteAsIndicative       string
RateTimeOut       string
PropagationDelay       string
CheckLiquidity       string
ChangeQuoteDirection       string
GenerateRejectedDeals       string
SpotUpdatesForForwardQuotes       string
SettlementInstructionStyle       string
CanRetractQuotes       string
CancelESPifNotPriced       string
CancelRFQSifNotPriced       string
CancelonDealingSuspended       string
CreditCheckedatDI       string
DuplicateCheckonExternalRef       string
LimitCheckQuote       string
LimitCheckonRFQDealSubmission       string
ListenonLimits       string
MarginStyle       string
UseRerouteDefinitionOnly       string
BypassConfirmation       string
DIOnAcceptance       string
IgnoreESPAmountRules       string

}

const (
	DealingInterface_Title       = "Dealing Interface"
	DealingInterface_SQLTable    = "sienaDealingInterface"
	DealingInterface_SQLSearchID = "Name"
	DealingInterface_QueryString = "Name"
	///
	/// Handler Defintions
	///
	DealingInterface_Template     = "DealingInterface"
	DealingInterface_TemplateList = "/DealingInterface/DealingInterface_List"
	DealingInterface_TemplateView = "/DealingInterface/DealingInterface_View"
	DealingInterface_TemplateEdit = "/DealingInterface/DealingInterface_Edit"
	DealingInterface_TemplateNew  = "/DealingInterface/DealingInterface_New"
	///
	/// Handler Monitor Paths
	///
	DealingInterface_Path       = "/API/DealingInterface/"
	DealingInterface_PathList   = "/DealingInterfaceList/"
	DealingInterface_PathView   = "/DealingInterfaceView/"
	DealingInterface_PathEdit   = "/DealingInterfaceEdit/"
	DealingInterface_PathNew    = "/DealingInterfaceNew/"
	DealingInterface_PathSave   = "/DealingInterfaceSave/"
	DealingInterface_PathDelete = "/DealingInterfaceDelete/"
	///
	///
	/// SQL Field Definitions
	///
DealingInterface_Name_sql   = "Name" // Name is a String
DealingInterface_AcceptReducedAmount_sql   = "AcceptReducedAmount" // AcceptReducedAmount is a Bool
DealingInterface_QuoteAsIndicative_sql   = "QuoteAsIndicative" // QuoteAsIndicative is a Bool
DealingInterface_RateTimeOut_sql   = "RateTimeOut" // RateTimeOut is a Int
DealingInterface_PropagationDelay_sql   = "PropagationDelay" // PropagationDelay is a Int
DealingInterface_CheckLiquidity_sql   = "CheckLiquidity" // CheckLiquidity is a Bool
DealingInterface_ChangeQuoteDirection_sql   = "ChangeQuoteDirection" // ChangeQuoteDirection is a Bool
DealingInterface_GenerateRejectedDeals_sql   = "GenerateRejectedDeals" // GenerateRejectedDeals is a Bool
DealingInterface_SpotUpdatesForForwardQuotes_sql   = "SpotUpdatesForForwardQuotes" // SpotUpdatesForForwardQuotes is a Bool
DealingInterface_SettlementInstructionStyle_sql   = "SettlementInstructionStyle" // SettlementInstructionStyle is a String
DealingInterface_CanRetractQuotes_sql   = "CanRetractQuotes" // CanRetractQuotes is a Bool
DealingInterface_CancelESPifNotPriced_sql   = "CancelESPifNotPriced" // CancelESPifNotPriced is a Bool
DealingInterface_CancelRFQSifNotPriced_sql   = "CancelRFQSifNotPriced" // CancelRFQSifNotPriced is a Bool
DealingInterface_CancelonDealingSuspended_sql   = "CancelonDealingSuspended" // CancelonDealingSuspended is a Bool
DealingInterface_CreditCheckedatDI_sql   = "CreditCheckedatDI" // CreditCheckedatDI is a Bool
DealingInterface_DuplicateCheckonExternalRef_sql   = "DuplicateCheckonExternalRef" // DuplicateCheckonExternalRef is a Bool
DealingInterface_LimitCheckQuote_sql   = "LimitCheckQuote" // LimitCheckQuote is a Bool
DealingInterface_LimitCheckonRFQDealSubmission_sql   = "LimitCheckonRFQDealSubmission" // LimitCheckonRFQDealSubmission is a Bool
DealingInterface_ListenonLimits_sql   = "ListenonLimits" // ListenonLimits is a Bool
DealingInterface_MarginStyle_sql   = "MarginStyle" // MarginStyle is a String
DealingInterface_UseRerouteDefinitionOnly_sql   = "UseRerouteDefinitionOnly" // UseRerouteDefinitionOnly is a Bool
DealingInterface_BypassConfirmation_sql   = "BypassConfirmation" // BypassConfirmation is a Bool
DealingInterface_DIOnAcceptance_sql   = "DIOnAcceptance" // DIOnAcceptance is a Bool
DealingInterface_IgnoreESPAmountRules_sql   = "IgnoreESPAmountRules" // IgnoreESPAmountRules is a Bool

	/// Definitions End

	/// Application Field Definitions
	///
DealingInterface_Name_scrn   = "Name" // Name is a String
DealingInterface_AcceptReducedAmount_scrn   = "AcceptReducedAmount" // AcceptReducedAmount is a Bool
DealingInterface_QuoteAsIndicative_scrn   = "QuoteAsIndicative" // QuoteAsIndicative is a Bool
DealingInterface_RateTimeOut_scrn   = "RateTimeOut" // RateTimeOut is a Int
DealingInterface_PropagationDelay_scrn   = "PropagationDelay" // PropagationDelay is a Int
DealingInterface_CheckLiquidity_scrn   = "CheckLiquidity" // CheckLiquidity is a Bool
DealingInterface_ChangeQuoteDirection_scrn   = "ChangeQuoteDirection" // ChangeQuoteDirection is a Bool
DealingInterface_GenerateRejectedDeals_scrn   = "GenerateRejectedDeals" // GenerateRejectedDeals is a Bool
DealingInterface_SpotUpdatesForForwardQuotes_scrn   = "SpotUpdatesForForwardQuotes" // SpotUpdatesForForwardQuotes is a Bool
DealingInterface_SettlementInstructionStyle_scrn   = "SettlementInstructionStyle" // SettlementInstructionStyle is a String
DealingInterface_CanRetractQuotes_scrn   = "CanRetractQuotes" // CanRetractQuotes is a Bool
DealingInterface_CancelESPifNotPriced_scrn   = "CancelESPifNotPriced" // CancelESPifNotPriced is a Bool
DealingInterface_CancelRFQSifNotPriced_scrn   = "CancelRFQSifNotPriced" // CancelRFQSifNotPriced is a Bool
DealingInterface_CancelonDealingSuspended_scrn   = "CancelonDealingSuspended" // CancelonDealingSuspended is a Bool
DealingInterface_CreditCheckedatDI_scrn   = "CreditCheckedatDI" // CreditCheckedatDI is a Bool
DealingInterface_DuplicateCheckonExternalRef_scrn   = "DuplicateCheckonExternalRef" // DuplicateCheckonExternalRef is a Bool
DealingInterface_LimitCheckQuote_scrn   = "LimitCheckQuote" // LimitCheckQuote is a Bool
DealingInterface_LimitCheckonRFQDealSubmission_scrn   = "LimitCheckonRFQDealSubmission" // LimitCheckonRFQDealSubmission is a Bool
DealingInterface_ListenonLimits_scrn   = "ListenonLimits" // ListenonLimits is a Bool
DealingInterface_MarginStyle_scrn   = "MarginStyle" // MarginStyle is a String
DealingInterface_UseRerouteDefinitionOnly_scrn   = "UseRerouteDefinitionOnly" // UseRerouteDefinitionOnly is a Bool
DealingInterface_BypassConfirmation_scrn   = "BypassConfirmation" // BypassConfirmation is a Bool
DealingInterface_DIOnAcceptance_scrn   = "DIOnAcceptance" // DIOnAcceptance is a Bool
DealingInterface_IgnoreESPAmountRules_scrn   = "IgnoreESPAmountRules" // IgnoreESPAmountRules is a Bool

	/// Definitions End
)
