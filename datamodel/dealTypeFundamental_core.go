package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/dealtypefundamental.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DealTypeFundamental (dealtypefundamental)
// Endpoint 	        : DealTypeFundamental (DealTypeKey)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:15
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type DealTypeFundamental struct {

DealTypeKey        string
Amendment        string
DefaultFrequency        string
Directions        string
SettledTermDealType        string
Optn        string
AllowPledge        string
Takeup        string
MismatchDealType        string
SettledHedgeTermDealType        string
SettlementCode        string
TermSubType        string
FundingDealType        string
TransferType        string
TermDealType        string
NegotiableInstrumentType        string
Mismatch        string
ComplexTransferSubType        string
LayOffDealType        string
NIAccount        string
SimpleMMsubtype        string
SwapDealType        string
Positions        string
OptionOutright        string
SettledHedgeSpotDealType        string
StraightThroughInterestMethod        string
SubType        string
Rollover        string
DefaultIssuer        string
DefaultStartDate        string
Fee        string
NDF        string
FXFX        string
ONIA        string
MarginSubType        string
TransferDealType        string
IsFX        string
Ordr        string
OptionStyle        string
SpotDealType        string
CanIssue        string
CanShort        string
FXMarginTradingType        string
Internal        string
TicketBasename        string
InterestRateFutureType        string
TradingLimitProductCode        string
Forward        string
MaturityNotificationPeriod        string
NotificationEvents        string
SwapSubType        string
ProductCode        string
Funding        string
AllocationPricing        string
CancelPeriod        string
MMMarginTradingType        string
OptionSpot        string
Transfer        string
NotificationPeriod        string
Paymentdateshift        string
CloseOut        string
FXOptionPricing        string
SettledHedgeOutrightDealType        string
RepoBond        string
RepoTerm        string
RepoType        string
DateRule        string
CorpTransferDealType        string
GenerateFXImage        string
Variant        string
HedgeTermDealType        string
PricingModel        string
HedgeOutrightDealType        string
Fixing        string
Payment        string
MT        string
SettlementInstructionStyle        string
QuoteHistoryRequired        string
Brokerage        string
ExposureDisabled        string
CreditLine        string
Encumbered        string
InternalId        string
InternalDeleted        string
UpdatedTransactionId        string
UpdatedUserId        string
UpdatedDateTime        string
DeletedTransactionId        string
DeletedUserId        string
ChangeType        string

}

const (
	DealTypeFundamental_Title       = "Deal Type Fundamental"
	DealTypeFundamental_SQLTable    = "sienaDealTypeFundamentals"
	DealTypeFundamental_SQLSearchID = "DealTypeKey"
	DealTypeFundamental_QueryString = "DealTypeKey"
	///
	/// Handler Defintions
	///
	DealTypeFundamental_Template     = "DealTypeFundamental"
	DealTypeFundamental_TemplateList = "DealTypeFundamental_List"
	DealTypeFundamental_TemplateView = "DealTypeFundamental_View"
	DealTypeFundamental_TemplateEdit = "DealTypeFundamental_Edit"
	DealTypeFundamental_TemplateNew  = "DealTypeFundamental_New"
	///
	/// Handler Monitor Paths
	///
	DealTypeFundamental_Path       = "/API/DealTypeFundamental/"
	DealTypeFundamental_PathList   = "/DealTypeFundamentalList/"
	DealTypeFundamental_PathView   = "/DealTypeFundamentalView/"
	DealTypeFundamental_PathEdit   = "/DealTypeFundamentalEdit/"
	DealTypeFundamental_PathNew    = "/DealTypeFundamentalNew/"
	DealTypeFundamental_PathSave   = "/DealTypeFundamentalSave/"
	DealTypeFundamental_PathDelete = "/DealTypeFundamentalDelete/"
	///
	/// SQL Field Definitions
	///
	DealTypeFundamental_DealTypeKey   = "DealTypeKey" // DealTypeKey is a String
	DealTypeFundamental_Amendment   = "Amendment" // Amendment is a Bool
	DealTypeFundamental_DefaultFrequency   = "DefaultFrequency" // DefaultFrequency is a Int
	DealTypeFundamental_Directions   = "Directions" // Directions is a String
	DealTypeFundamental_SettledTermDealType   = "SettledTermDealType" // SettledTermDealType is a String
	DealTypeFundamental_Optn   = "Optn" // Optn is a Bool
	DealTypeFundamental_AllowPledge   = "AllowPledge" // AllowPledge is a Bool
	DealTypeFundamental_Takeup   = "Takeup" // Takeup is a Bool
	DealTypeFundamental_MismatchDealType   = "MismatchDealType" // MismatchDealType is a String
	DealTypeFundamental_SettledHedgeTermDealType   = "SettledHedgeTermDealType" // SettledHedgeTermDealType is a String
	DealTypeFundamental_SettlementCode   = "SettlementCode" // SettlementCode is a String
	DealTypeFundamental_TermSubType   = "TermSubType" // TermSubType is a String
	DealTypeFundamental_FundingDealType   = "FundingDealType" // FundingDealType is a String
	DealTypeFundamental_TransferType   = "TransferType" // TransferType is a String
	DealTypeFundamental_TermDealType   = "TermDealType" // TermDealType is a String
	DealTypeFundamental_NegotiableInstrumentType   = "NegotiableInstrumentType" // NegotiableInstrumentType is a String
	DealTypeFundamental_Mismatch   = "Mismatch" // Mismatch is a Bool
	DealTypeFundamental_ComplexTransferSubType   = "ComplexTransferSubType" // ComplexTransferSubType is a String
	DealTypeFundamental_LayOffDealType   = "LayOffDealType" // LayOffDealType is a String
	DealTypeFundamental_NIAccount   = "NIAccount" // NIAccount is a Int
	DealTypeFundamental_SimpleMMsubtype   = "SimpleMMsubtype" // SimpleMMsubtype is a Int
	DealTypeFundamental_SwapDealType   = "SwapDealType" // SwapDealType is a String
	DealTypeFundamental_Positions   = "Positions" // Positions is a String
	DealTypeFundamental_OptionOutright   = "OptionOutright" // OptionOutright is a String
	DealTypeFundamental_SettledHedgeSpotDealType   = "SettledHedgeSpotDealType" // SettledHedgeSpotDealType is a String
	DealTypeFundamental_StraightThroughInterestMethod   = "StraightThroughInterestMethod" // StraightThroughInterestMethod is a Bool
	DealTypeFundamental_SubType   = "SubType" // SubType is a String
	DealTypeFundamental_Rollover   = "Rollover" // Rollover is a Bool
	DealTypeFundamental_DefaultIssuer   = "DefaultIssuer" // DefaultIssuer is a String
	DealTypeFundamental_DefaultStartDate   = "DefaultStartDate" // DefaultStartDate is a Int
	DealTypeFundamental_Fee   = "Fee" // Fee is a String
	DealTypeFundamental_NDF   = "NDF" // NDF is a Bool
	DealTypeFundamental_FXFX   = "FXFX" // FXFX is a Bool
	DealTypeFundamental_ONIA   = "ONIA" // ONIA is a Bool
	DealTypeFundamental_MarginSubType   = "MarginSubType" // MarginSubType is a Int
	DealTypeFundamental_TransferDealType   = "TransferDealType" // TransferDealType is a String
	DealTypeFundamental_IsFX   = "IsFX" // IsFX is a Bool
	DealTypeFundamental_Ordr   = "Ordr" // Ordr is a String
	DealTypeFundamental_OptionStyle   = "OptionStyle" // OptionStyle is a String
	DealTypeFundamental_SpotDealType   = "SpotDealType" // SpotDealType is a String
	DealTypeFundamental_CanIssue   = "CanIssue" // CanIssue is a Bool
	DealTypeFundamental_CanShort   = "CanShort" // CanShort is a Bool
	DealTypeFundamental_FXMarginTradingType   = "FXMarginTradingType" // FXMarginTradingType is a Int
	DealTypeFundamental_Internal   = "Internal" // Internal is a Bool
	DealTypeFundamental_TicketBasename   = "TicketBasename" // TicketBasename is a String
	DealTypeFundamental_InterestRateFutureType   = "InterestRateFutureType" // InterestRateFutureType is a String
	DealTypeFundamental_TradingLimitProductCode   = "TradingLimitProductCode" // TradingLimitProductCode is a String
	DealTypeFundamental_Forward   = "Forward" // Forward is a Bool
	DealTypeFundamental_MaturityNotificationPeriod   = "MaturityNotificationPeriod" // MaturityNotificationPeriod is a String
	DealTypeFundamental_NotificationEvents   = "NotificationEvents" // NotificationEvents is a String
	DealTypeFundamental_SwapSubType   = "SwapSubType" // SwapSubType is a String
	DealTypeFundamental_ProductCode   = "ProductCode" // ProductCode is a String
	DealTypeFundamental_Funding   = "Funding" // Funding is a Bool
	DealTypeFundamental_AllocationPricing   = "AllocationPricing" // AllocationPricing is a String
	DealTypeFundamental_CancelPeriod   = "CancelPeriod" // CancelPeriod is a String
	DealTypeFundamental_MMMarginTradingType   = "MMMarginTradingType" // MMMarginTradingType is a Int
	DealTypeFundamental_OptionSpot   = "OptionSpot" // OptionSpot is a String
	DealTypeFundamental_Transfer   = "Transfer" // Transfer is a Bool
	DealTypeFundamental_NotificationPeriod   = "NotificationPeriod" // NotificationPeriod is a String
	DealTypeFundamental_Paymentdateshift   = "Paymentdateshift" // Paymentdateshift is a Int
	DealTypeFundamental_CloseOut   = "CloseOut" // CloseOut is a Bool
	DealTypeFundamental_FXOptionPricing   = "FXOptionPricing" // FXOptionPricing is a String
	DealTypeFundamental_SettledHedgeOutrightDealType   = "SettledHedgeOutrightDealType" // SettledHedgeOutrightDealType is a String
	DealTypeFundamental_RepoBond   = "RepoBond" // RepoBond is a String
	DealTypeFundamental_RepoTerm   = "RepoTerm" // RepoTerm is a String
	DealTypeFundamental_RepoType   = "RepoType" // RepoType is a Int
	DealTypeFundamental_DateRule   = "DateRule" // DateRule is a String
	DealTypeFundamental_CorpTransferDealType   = "CorpTransferDealType" // CorpTransferDealType is a String
	DealTypeFundamental_GenerateFXImage   = "GenerateFXImage" // GenerateFXImage is a Bool
	DealTypeFundamental_Variant   = "Variant" // Variant is a String
	DealTypeFundamental_HedgeTermDealType   = "HedgeTermDealType" // HedgeTermDealType is a String
	DealTypeFundamental_PricingModel   = "PricingModel" // PricingModel is a String
	DealTypeFundamental_HedgeOutrightDealType   = "HedgeOutrightDealType" // HedgeOutrightDealType is a String
	DealTypeFundamental_Fixing   = "Fixing" // Fixing is a Bool
	DealTypeFundamental_Payment   = "Payment" // Payment is a Bool
	DealTypeFundamental_MT   = "MT" // MT is a Bool
	DealTypeFundamental_SettlementInstructionStyle   = "SettlementInstructionStyle" // SettlementInstructionStyle is a String
	DealTypeFundamental_QuoteHistoryRequired   = "QuoteHistoryRequired" // QuoteHistoryRequired is a Bool
	DealTypeFundamental_Brokerage   = "Brokerage" // Brokerage is a Bool
	DealTypeFundamental_ExposureDisabled   = "ExposureDisabled" // ExposureDisabled is a Bool
	DealTypeFundamental_CreditLine   = "CreditLine" // CreditLine is a String
	DealTypeFundamental_Encumbered   = "Encumbered" // Encumbered is a Bool
	DealTypeFundamental_InternalId   = "InternalId" // InternalId is a Int
	DealTypeFundamental_InternalDeleted   = "InternalDeleted" // InternalDeleted is a Time
	DealTypeFundamental_UpdatedTransactionId   = "UpdatedTransactionId" // UpdatedTransactionId is a String
	DealTypeFundamental_UpdatedUserId   = "UpdatedUserId" // UpdatedUserId is a String
	DealTypeFundamental_UpdatedDateTime   = "UpdatedDateTime" // UpdatedDateTime is a Time
	DealTypeFundamental_DeletedTransactionId   = "DeletedTransactionId" // DeletedTransactionId is a String
	DealTypeFundamental_DeletedUserId   = "DeletedUserId" // DeletedUserId is a String
	DealTypeFundamental_ChangeType   = "ChangeType" // ChangeType is a String

	/// Definitions End
)
