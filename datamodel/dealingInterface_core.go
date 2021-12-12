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
// Date & Time		    : 12/12/2021 at 16:13:14
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type DealingInterface struct {

Name        string
AcceptReducedAmount        string
QuoteAsIndicative        string
RateTimeOut        string
PropagationDelay        string
CheckLiquidity        string
ChangeQuoteDirection        string
GenerateRejectedDeals        string
SpotUpdatesForForwardQuotes        string
SettlementInstructionStyle        string
CanRetractQuotes        string
CancelESPifNotPriced        string
CancelRFQSifNotPriced        string
CancelonDealingSuspended        string
CreditCheckedatDI        string
DuplicateCheckonExternalRef        string
LimitCheckQuote        string
LimitCheckonRFQDealSubmission        string
ListenonLimits        string
MarginStyle        string
UseRerouteDefinitionOnly        string
BypassConfirmation        string
DIOnAcceptance        string
IgnoreESPAmountRules        string

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
	DealingInterface_TemplateList = "DealingInterface_List"
	DealingInterface_TemplateView = "DealingInterface_View"
	DealingInterface_TemplateEdit = "DealingInterface_Edit"
	DealingInterface_TemplateNew  = "DealingInterface_New"
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
	/// SQL Field Definitions
	///
	DealingInterface_Name   = "Name" // Name is a String
	DealingInterface_AcceptReducedAmount   = "AcceptReducedAmount" // AcceptReducedAmount is a Bool
	DealingInterface_QuoteAsIndicative   = "QuoteAsIndicative" // QuoteAsIndicative is a Bool
	DealingInterface_RateTimeOut   = "RateTimeOut" // RateTimeOut is a Int
	DealingInterface_PropagationDelay   = "PropagationDelay" // PropagationDelay is a Int
	DealingInterface_CheckLiquidity   = "CheckLiquidity" // CheckLiquidity is a Bool
	DealingInterface_ChangeQuoteDirection   = "ChangeQuoteDirection" // ChangeQuoteDirection is a Bool
	DealingInterface_GenerateRejectedDeals   = "GenerateRejectedDeals" // GenerateRejectedDeals is a Bool
	DealingInterface_SpotUpdatesForForwardQuotes   = "SpotUpdatesForForwardQuotes" // SpotUpdatesForForwardQuotes is a Bool
	DealingInterface_SettlementInstructionStyle   = "SettlementInstructionStyle" // SettlementInstructionStyle is a String
	DealingInterface_CanRetractQuotes   = "CanRetractQuotes" // CanRetractQuotes is a Bool
	DealingInterface_CancelESPifNotPriced   = "CancelESPifNotPriced" // CancelESPifNotPriced is a Bool
	DealingInterface_CancelRFQSifNotPriced   = "CancelRFQSifNotPriced" // CancelRFQSifNotPriced is a Bool
	DealingInterface_CancelonDealingSuspended   = "CancelonDealingSuspended" // CancelonDealingSuspended is a Bool
	DealingInterface_CreditCheckedatDI   = "CreditCheckedatDI" // CreditCheckedatDI is a Bool
	DealingInterface_DuplicateCheckonExternalRef   = "DuplicateCheckonExternalRef" // DuplicateCheckonExternalRef is a Bool
	DealingInterface_LimitCheckQuote   = "LimitCheckQuote" // LimitCheckQuote is a Bool
	DealingInterface_LimitCheckonRFQDealSubmission   = "LimitCheckonRFQDealSubmission" // LimitCheckonRFQDealSubmission is a Bool
	DealingInterface_ListenonLimits   = "ListenonLimits" // ListenonLimits is a Bool
	DealingInterface_MarginStyle   = "MarginStyle" // MarginStyle is a String
	DealingInterface_UseRerouteDefinitionOnly   = "UseRerouteDefinitionOnly" // UseRerouteDefinitionOnly is a Bool
	DealingInterface_BypassConfirmation   = "BypassConfirmation" // BypassConfirmation is a Bool
	DealingInterface_DIOnAcceptance   = "DIOnAcceptance" // DIOnAcceptance is a Bool
	DealingInterface_IgnoreESPAmountRules   = "IgnoreESPAmountRules" // IgnoreESPAmountRules is a Bool

	/// Definitions End
)
