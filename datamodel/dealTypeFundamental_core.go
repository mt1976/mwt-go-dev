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
// Date & Time		    : 17/06/2022 at 18:38:10
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//DealTypeFundamental defines the datamolde for the DealTypeFundamental object
type DealTypeFundamental struct {

DealTypeKey       string
Amendment       string
DefaultFrequency       string
Directions       string
SettledTermDealType       string
Optn       string
AllowPledge       string
Takeup       string
MismatchDealType       string
SettledHedgeTermDealType       string
SettlementCode       string
TermSubType       string
FundingDealType       string
TransferType       string
TermDealType       string
NegotiableInstrumentType       string
Mismatch       string
ComplexTransferSubType       string
LayOffDealType       string
NIAccount       string
SimpleMMsubtype       string
SwapDealType       string
Positions       string
OptionOutright       string
SettledHedgeSpotDealType       string
StraightThroughInterestMethod       string
SubType       string
Rollover       string
DefaultIssuer       string
DefaultStartDate       string
Fee       string
NDF       string
FXFX       string
ONIA       string
MarginSubType       string
TransferDealType       string
IsFX       string
Ordr       string
OptionStyle       string
SpotDealType       string
CanIssue       string
CanShort       string
FXMarginTradingType       string
Internal       string
TicketBasename       string
InterestRateFutureType       string
TradingLimitProductCode       string
Forward       string
MaturityNotificationPeriod       string
NotificationEvents       string
SwapSubType       string
ProductCode       string
Funding       string
AllocationPricing       string
CancelPeriod       string
MMMarginTradingType       string
OptionSpot       string
Transfer       string
NotificationPeriod       string
Paymentdateshift       string
CloseOut       string
FXOptionPricing       string
SettledHedgeOutrightDealType       string
RepoBond       string
RepoTerm       string
RepoType       string
DateRule       string
CorpTransferDealType       string
GenerateFXImage       string
Variant       string
HedgeTermDealType       string
PricingModel       string
HedgeOutrightDealType       string
Fixing       string
Payment       string
MT       string
SettlementInstructionStyle       string
QuoteHistoryRequired       string
Brokerage       string
ExposureDisabled       string
CreditLine       string
Encumbered       string
InternalId       string
InternalDeleted       string
UpdatedTransactionId       string
UpdatedUserId       string
UpdatedDateTime       string
DeletedTransactionId       string
DeletedUserId       string
ChangeType       string

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
	DealTypeFundamental_TemplateList = "/DealTypeFundamental/DealTypeFundamental_List"
	DealTypeFundamental_TemplateView = "/DealTypeFundamental/DealTypeFundamental_View"
	DealTypeFundamental_TemplateEdit = "/DealTypeFundamental/DealTypeFundamental_Edit"
	DealTypeFundamental_TemplateNew  = "/DealTypeFundamental/DealTypeFundamental_New"
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
	///
	/// SQL Field Definitions
	///
DealTypeFundamental_DealTypeKey_sql   = "DealTypeKey" // DealTypeKey is a String
DealTypeFundamental_Amendment_sql   = "Amendment" // Amendment is a Bool
DealTypeFundamental_DefaultFrequency_sql   = "DefaultFrequency" // DefaultFrequency is a Int
DealTypeFundamental_Directions_sql   = "Directions" // Directions is a String
DealTypeFundamental_SettledTermDealType_sql   = "SettledTermDealType" // SettledTermDealType is a String
DealTypeFundamental_Optn_sql   = "Optn" // Optn is a Bool
DealTypeFundamental_AllowPledge_sql   = "AllowPledge" // AllowPledge is a Bool
DealTypeFundamental_Takeup_sql   = "Takeup" // Takeup is a Bool
DealTypeFundamental_MismatchDealType_sql   = "MismatchDealType" // MismatchDealType is a String
DealTypeFundamental_SettledHedgeTermDealType_sql   = "SettledHedgeTermDealType" // SettledHedgeTermDealType is a String
DealTypeFundamental_SettlementCode_sql   = "SettlementCode" // SettlementCode is a String
DealTypeFundamental_TermSubType_sql   = "TermSubType" // TermSubType is a String
DealTypeFundamental_FundingDealType_sql   = "FundingDealType" // FundingDealType is a String
DealTypeFundamental_TransferType_sql   = "TransferType" // TransferType is a String
DealTypeFundamental_TermDealType_sql   = "TermDealType" // TermDealType is a String
DealTypeFundamental_NegotiableInstrumentType_sql   = "NegotiableInstrumentType" // NegotiableInstrumentType is a String
DealTypeFundamental_Mismatch_sql   = "Mismatch" // Mismatch is a Bool
DealTypeFundamental_ComplexTransferSubType_sql   = "ComplexTransferSubType" // ComplexTransferSubType is a String
DealTypeFundamental_LayOffDealType_sql   = "LayOffDealType" // LayOffDealType is a String
DealTypeFundamental_NIAccount_sql   = "NIAccount" // NIAccount is a Int
DealTypeFundamental_SimpleMMsubtype_sql   = "SimpleMMsubtype" // SimpleMMsubtype is a Int
DealTypeFundamental_SwapDealType_sql   = "SwapDealType" // SwapDealType is a String
DealTypeFundamental_Positions_sql   = "Positions" // Positions is a String
DealTypeFundamental_OptionOutright_sql   = "OptionOutright" // OptionOutright is a String
DealTypeFundamental_SettledHedgeSpotDealType_sql   = "SettledHedgeSpotDealType" // SettledHedgeSpotDealType is a String
DealTypeFundamental_StraightThroughInterestMethod_sql   = "StraightThroughInterestMethod" // StraightThroughInterestMethod is a Bool
DealTypeFundamental_SubType_sql   = "SubType" // SubType is a String
DealTypeFundamental_Rollover_sql   = "Rollover" // Rollover is a Bool
DealTypeFundamental_DefaultIssuer_sql   = "DefaultIssuer" // DefaultIssuer is a String
DealTypeFundamental_DefaultStartDate_sql   = "DefaultStartDate" // DefaultStartDate is a Int
DealTypeFundamental_Fee_sql   = "Fee" // Fee is a String
DealTypeFundamental_NDF_sql   = "NDF" // NDF is a Bool
DealTypeFundamental_FXFX_sql   = "FXFX" // FXFX is a Bool
DealTypeFundamental_ONIA_sql   = "ONIA" // ONIA is a Bool
DealTypeFundamental_MarginSubType_sql   = "MarginSubType" // MarginSubType is a Int
DealTypeFundamental_TransferDealType_sql   = "TransferDealType" // TransferDealType is a String
DealTypeFundamental_IsFX_sql   = "IsFX" // IsFX is a Bool
DealTypeFundamental_Ordr_sql   = "Ordr" // Ordr is a String
DealTypeFundamental_OptionStyle_sql   = "OptionStyle" // OptionStyle is a String
DealTypeFundamental_SpotDealType_sql   = "SpotDealType" // SpotDealType is a String
DealTypeFundamental_CanIssue_sql   = "CanIssue" // CanIssue is a Bool
DealTypeFundamental_CanShort_sql   = "CanShort" // CanShort is a Bool
DealTypeFundamental_FXMarginTradingType_sql   = "FXMarginTradingType" // FXMarginTradingType is a Int
DealTypeFundamental_Internal_sql   = "Internal" // Internal is a Bool
DealTypeFundamental_TicketBasename_sql   = "TicketBasename" // TicketBasename is a String
DealTypeFundamental_InterestRateFutureType_sql   = "InterestRateFutureType" // InterestRateFutureType is a String
DealTypeFundamental_TradingLimitProductCode_sql   = "TradingLimitProductCode" // TradingLimitProductCode is a String
DealTypeFundamental_Forward_sql   = "Forward" // Forward is a Bool
DealTypeFundamental_MaturityNotificationPeriod_sql   = "MaturityNotificationPeriod" // MaturityNotificationPeriod is a String
DealTypeFundamental_NotificationEvents_sql   = "NotificationEvents" // NotificationEvents is a String
DealTypeFundamental_SwapSubType_sql   = "SwapSubType" // SwapSubType is a String
DealTypeFundamental_ProductCode_sql   = "ProductCode" // ProductCode is a String
DealTypeFundamental_Funding_sql   = "Funding" // Funding is a Bool
DealTypeFundamental_AllocationPricing_sql   = "AllocationPricing" // AllocationPricing is a String
DealTypeFundamental_CancelPeriod_sql   = "CancelPeriod" // CancelPeriod is a String
DealTypeFundamental_MMMarginTradingType_sql   = "MMMarginTradingType" // MMMarginTradingType is a Int
DealTypeFundamental_OptionSpot_sql   = "OptionSpot" // OptionSpot is a String
DealTypeFundamental_Transfer_sql   = "Transfer" // Transfer is a Bool
DealTypeFundamental_NotificationPeriod_sql   = "NotificationPeriod" // NotificationPeriod is a String
DealTypeFundamental_Paymentdateshift_sql   = "Paymentdateshift" // Paymentdateshift is a Int
DealTypeFundamental_CloseOut_sql   = "CloseOut" // CloseOut is a Bool
DealTypeFundamental_FXOptionPricing_sql   = "FXOptionPricing" // FXOptionPricing is a String
DealTypeFundamental_SettledHedgeOutrightDealType_sql   = "SettledHedgeOutrightDealType" // SettledHedgeOutrightDealType is a String
DealTypeFundamental_RepoBond_sql   = "RepoBond" // RepoBond is a String
DealTypeFundamental_RepoTerm_sql   = "RepoTerm" // RepoTerm is a String
DealTypeFundamental_RepoType_sql   = "RepoType" // RepoType is a Int
DealTypeFundamental_DateRule_sql   = "DateRule" // DateRule is a String
DealTypeFundamental_CorpTransferDealType_sql   = "CorpTransferDealType" // CorpTransferDealType is a String
DealTypeFundamental_GenerateFXImage_sql   = "GenerateFXImage" // GenerateFXImage is a Bool
DealTypeFundamental_Variant_sql   = "Variant" // Variant is a String
DealTypeFundamental_HedgeTermDealType_sql   = "HedgeTermDealType" // HedgeTermDealType is a String
DealTypeFundamental_PricingModel_sql   = "PricingModel" // PricingModel is a String
DealTypeFundamental_HedgeOutrightDealType_sql   = "HedgeOutrightDealType" // HedgeOutrightDealType is a String
DealTypeFundamental_Fixing_sql   = "Fixing" // Fixing is a Bool
DealTypeFundamental_Payment_sql   = "Payment" // Payment is a Bool
DealTypeFundamental_MT_sql   = "MT" // MT is a Bool
DealTypeFundamental_SettlementInstructionStyle_sql   = "SettlementInstructionStyle" // SettlementInstructionStyle is a String
DealTypeFundamental_QuoteHistoryRequired_sql   = "QuoteHistoryRequired" // QuoteHistoryRequired is a Bool
DealTypeFundamental_Brokerage_sql   = "Brokerage" // Brokerage is a Bool
DealTypeFundamental_ExposureDisabled_sql   = "ExposureDisabled" // ExposureDisabled is a Bool
DealTypeFundamental_CreditLine_sql   = "CreditLine" // CreditLine is a String
DealTypeFundamental_Encumbered_sql   = "Encumbered" // Encumbered is a Bool
DealTypeFundamental_InternalId_sql   = "InternalId" // InternalId is a Int
DealTypeFundamental_InternalDeleted_sql   = "InternalDeleted" // InternalDeleted is a Time
DealTypeFundamental_UpdatedTransactionId_sql   = "UpdatedTransactionId" // UpdatedTransactionId is a String
DealTypeFundamental_UpdatedUserId_sql   = "UpdatedUserId" // UpdatedUserId is a String
DealTypeFundamental_UpdatedDateTime_sql   = "UpdatedDateTime" // UpdatedDateTime is a Time
DealTypeFundamental_DeletedTransactionId_sql   = "DeletedTransactionId" // DeletedTransactionId is a String
DealTypeFundamental_DeletedUserId_sql   = "DeletedUserId" // DeletedUserId is a String
DealTypeFundamental_ChangeType_sql   = "ChangeType" // ChangeType is a String

	/// Definitions End

	/// Application Field Definitions
	///
DealTypeFundamental_DealTypeKey_scrn   = "DealTypeKey" // DealTypeKey is a String
DealTypeFundamental_Amendment_scrn   = "Amendment" // Amendment is a Bool
DealTypeFundamental_DefaultFrequency_scrn   = "DefaultFrequency" // DefaultFrequency is a Int
DealTypeFundamental_Directions_scrn   = "Directions" // Directions is a String
DealTypeFundamental_SettledTermDealType_scrn   = "SettledTermDealType" // SettledTermDealType is a String
DealTypeFundamental_Optn_scrn   = "Optn" // Optn is a Bool
DealTypeFundamental_AllowPledge_scrn   = "AllowPledge" // AllowPledge is a Bool
DealTypeFundamental_Takeup_scrn   = "Takeup" // Takeup is a Bool
DealTypeFundamental_MismatchDealType_scrn   = "MismatchDealType" // MismatchDealType is a String
DealTypeFundamental_SettledHedgeTermDealType_scrn   = "SettledHedgeTermDealType" // SettledHedgeTermDealType is a String
DealTypeFundamental_SettlementCode_scrn   = "SettlementCode" // SettlementCode is a String
DealTypeFundamental_TermSubType_scrn   = "TermSubType" // TermSubType is a String
DealTypeFundamental_FundingDealType_scrn   = "FundingDealType" // FundingDealType is a String
DealTypeFundamental_TransferType_scrn   = "TransferType" // TransferType is a String
DealTypeFundamental_TermDealType_scrn   = "TermDealType" // TermDealType is a String
DealTypeFundamental_NegotiableInstrumentType_scrn   = "NegotiableInstrumentType" // NegotiableInstrumentType is a String
DealTypeFundamental_Mismatch_scrn   = "Mismatch" // Mismatch is a Bool
DealTypeFundamental_ComplexTransferSubType_scrn   = "ComplexTransferSubType" // ComplexTransferSubType is a String
DealTypeFundamental_LayOffDealType_scrn   = "LayOffDealType" // LayOffDealType is a String
DealTypeFundamental_NIAccount_scrn   = "NIAccount" // NIAccount is a Int
DealTypeFundamental_SimpleMMsubtype_scrn   = "SimpleMMsubtype" // SimpleMMsubtype is a Int
DealTypeFundamental_SwapDealType_scrn   = "SwapDealType" // SwapDealType is a String
DealTypeFundamental_Positions_scrn   = "Positions" // Positions is a String
DealTypeFundamental_OptionOutright_scrn   = "OptionOutright" // OptionOutright is a String
DealTypeFundamental_SettledHedgeSpotDealType_scrn   = "SettledHedgeSpotDealType" // SettledHedgeSpotDealType is a String
DealTypeFundamental_StraightThroughInterestMethod_scrn   = "StraightThroughInterestMethod" // StraightThroughInterestMethod is a Bool
DealTypeFundamental_SubType_scrn   = "SubType" // SubType is a String
DealTypeFundamental_Rollover_scrn   = "Rollover" // Rollover is a Bool
DealTypeFundamental_DefaultIssuer_scrn   = "DefaultIssuer" // DefaultIssuer is a String
DealTypeFundamental_DefaultStartDate_scrn   = "DefaultStartDate" // DefaultStartDate is a Int
DealTypeFundamental_Fee_scrn   = "Fee" // Fee is a String
DealTypeFundamental_NDF_scrn   = "NDF" // NDF is a Bool
DealTypeFundamental_FXFX_scrn   = "FXFX" // FXFX is a Bool
DealTypeFundamental_ONIA_scrn   = "ONIA" // ONIA is a Bool
DealTypeFundamental_MarginSubType_scrn   = "MarginSubType" // MarginSubType is a Int
DealTypeFundamental_TransferDealType_scrn   = "TransferDealType" // TransferDealType is a String
DealTypeFundamental_IsFX_scrn   = "IsFX" // IsFX is a Bool
DealTypeFundamental_Ordr_scrn   = "Ordr" // Ordr is a String
DealTypeFundamental_OptionStyle_scrn   = "OptionStyle" // OptionStyle is a String
DealTypeFundamental_SpotDealType_scrn   = "SpotDealType" // SpotDealType is a String
DealTypeFundamental_CanIssue_scrn   = "CanIssue" // CanIssue is a Bool
DealTypeFundamental_CanShort_scrn   = "CanShort" // CanShort is a Bool
DealTypeFundamental_FXMarginTradingType_scrn   = "FXMarginTradingType" // FXMarginTradingType is a Int
DealTypeFundamental_Internal_scrn   = "Internal" // Internal is a Bool
DealTypeFundamental_TicketBasename_scrn   = "TicketBasename" // TicketBasename is a String
DealTypeFundamental_InterestRateFutureType_scrn   = "InterestRateFutureType" // InterestRateFutureType is a String
DealTypeFundamental_TradingLimitProductCode_scrn   = "TradingLimitProductCode" // TradingLimitProductCode is a String
DealTypeFundamental_Forward_scrn   = "Forward" // Forward is a Bool
DealTypeFundamental_MaturityNotificationPeriod_scrn   = "MaturityNotificationPeriod" // MaturityNotificationPeriod is a String
DealTypeFundamental_NotificationEvents_scrn   = "NotificationEvents" // NotificationEvents is a String
DealTypeFundamental_SwapSubType_scrn   = "SwapSubType" // SwapSubType is a String
DealTypeFundamental_ProductCode_scrn   = "ProductCode" // ProductCode is a String
DealTypeFundamental_Funding_scrn   = "Funding" // Funding is a Bool
DealTypeFundamental_AllocationPricing_scrn   = "AllocationPricing" // AllocationPricing is a String
DealTypeFundamental_CancelPeriod_scrn   = "CancelPeriod" // CancelPeriod is a String
DealTypeFundamental_MMMarginTradingType_scrn   = "MMMarginTradingType" // MMMarginTradingType is a Int
DealTypeFundamental_OptionSpot_scrn   = "OptionSpot" // OptionSpot is a String
DealTypeFundamental_Transfer_scrn   = "Transfer" // Transfer is a Bool
DealTypeFundamental_NotificationPeriod_scrn   = "NotificationPeriod" // NotificationPeriod is a String
DealTypeFundamental_Paymentdateshift_scrn   = "Paymentdateshift" // Paymentdateshift is a Int
DealTypeFundamental_CloseOut_scrn   = "CloseOut" // CloseOut is a Bool
DealTypeFundamental_FXOptionPricing_scrn   = "FXOptionPricing" // FXOptionPricing is a String
DealTypeFundamental_SettledHedgeOutrightDealType_scrn   = "SettledHedgeOutrightDealType" // SettledHedgeOutrightDealType is a String
DealTypeFundamental_RepoBond_scrn   = "RepoBond" // RepoBond is a String
DealTypeFundamental_RepoTerm_scrn   = "RepoTerm" // RepoTerm is a String
DealTypeFundamental_RepoType_scrn   = "RepoType" // RepoType is a Int
DealTypeFundamental_DateRule_scrn   = "DateRule" // DateRule is a String
DealTypeFundamental_CorpTransferDealType_scrn   = "CorpTransferDealType" // CorpTransferDealType is a String
DealTypeFundamental_GenerateFXImage_scrn   = "GenerateFXImage" // GenerateFXImage is a Bool
DealTypeFundamental_Variant_scrn   = "Variant" // Variant is a String
DealTypeFundamental_HedgeTermDealType_scrn   = "HedgeTermDealType" // HedgeTermDealType is a String
DealTypeFundamental_PricingModel_scrn   = "PricingModel" // PricingModel is a String
DealTypeFundamental_HedgeOutrightDealType_scrn   = "HedgeOutrightDealType" // HedgeOutrightDealType is a String
DealTypeFundamental_Fixing_scrn   = "Fixing" // Fixing is a Bool
DealTypeFundamental_Payment_scrn   = "Payment" // Payment is a Bool
DealTypeFundamental_MT_scrn   = "MT" // MT is a Bool
DealTypeFundamental_SettlementInstructionStyle_scrn   = "SettlementInstructionStyle" // SettlementInstructionStyle is a String
DealTypeFundamental_QuoteHistoryRequired_scrn   = "QuoteHistoryRequired" // QuoteHistoryRequired is a Bool
DealTypeFundamental_Brokerage_scrn   = "Brokerage" // Brokerage is a Bool
DealTypeFundamental_ExposureDisabled_scrn   = "ExposureDisabled" // ExposureDisabled is a Bool
DealTypeFundamental_CreditLine_scrn   = "CreditLine" // CreditLine is a String
DealTypeFundamental_Encumbered_scrn   = "Encumbered" // Encumbered is a Bool
DealTypeFundamental_InternalId_scrn   = "InternalId" // InternalId is a Int
DealTypeFundamental_InternalDeleted_scrn   = "InternalDeleted" // InternalDeleted is a Time
DealTypeFundamental_UpdatedTransactionId_scrn   = "UpdatedTransactionId" // UpdatedTransactionId is a String
DealTypeFundamental_UpdatedUserId_scrn   = "UpdatedUserId" // UpdatedUserId is a String
DealTypeFundamental_UpdatedDateTime_scrn   = "UpdatedDateTime" // UpdatedDateTime is a Time
DealTypeFundamental_DeletedTransactionId_scrn   = "DeletedTransactionId" // DeletedTransactionId is a String
DealTypeFundamental_DeletedUserId_scrn   = "DeletedUserId" // DeletedUserId is a String
DealTypeFundamental_ChangeType_scrn   = "ChangeType" // ChangeType is a String

	/// Definitions End
)
