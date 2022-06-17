package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/transaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Transaction (transaction)
// Endpoint 	        : Transaction (Ref)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:03
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Transaction defines the datamolde for the Transaction object
type Transaction struct {

SienaReference       string
Status       string
ValueDate       string
MaturityDate       string
ContractNumber       string
ExternalReference       string
Book       string
MandatedUser       string
Portfolio       string
AgreementId       string
BackOfficeRefNo       string
ISIN       string
UTI       string
BookName       string
Centre       string
Firm       string
DealTypeShortName       string
FullDealType       string
TradeDate       string
DealtCcy       string
DealtAmount       string
AgainstAmount       string
AgainstCcy       string
AllInRate       string
MktRate       string
SettleCcy       string
Direction       string
NpvRate       string
OriginUser       string
PayInstruction       string
ReceiptInstruction       string
NIName       string
CCYPair       string
Instrument       string
PortfolioName       string
RVDate       string
RVMTM       string
CounterBook       string
CounterBookName       string
Party       string
PartyName       string
NameCentre       string
NameFirm       string
CustomerExternalView       string
CustomerSienaView       string
CompID       string
SienaDealer       string
DealOwner       string
DealOwnerMnemonic       string
EditedByUser       string
UTCOriginTime       string
UTCUpdateTime       string
MarginTrading       string
SwapPoints       string
SpotDate       string
SpotRate       string
MktSpotRate       string
SpotSalesMargin       string
SpotChlMargin       string
RerouteCcy       string
CustomerPayInstruction       string
CustomerReceiptInstruction       string
BackOfficeNotes       string
CustomerStatementNotes       string
NotesMargin       string
RequestedBy       string
EditReason       string
EditOtherReason       string
NICleanPrice       string
NIDirtyPrice       string
NIYield       string
NIClearingSystem       string
Acceptor       string
NIDiscount       string
FastPay       string
PaymentFee       string
PaymentFeePolicy       string
PaymentReason       string
PaymentDate       string
SettlementDate       string
FixingDate       string
VenueUTI       string
EditVersion       string
BrokeragePercentage       string
BrokerageAmount       string
BrokerageCurrency       string
BrokerageDate       string
AccountName       string
AccountNumber       string
CashBalance       string
DebitFrequency       string
CreditFrequency       string
ManuallyQuoted       string
LedgerBalance       string
SettAmtOutstanding       string
FeePercentage       string
FeeAmount       string
Venue       string
EURAmount       string
EUROtherAmount       string
LEI       string
Equity       string
Shares       string
QuoteExpiryDate       string
Commodity       string
PaymentSystemSienaView       string
PaymentSystemExternalView       string
SalesProfit       string
RejectReason       string
PaymentError       string
RepoPrincipal       string
FixingFrequency       string
Dealt       string
Against       string

}

const (
	Transaction_Title       = "Transaction"
	Transaction_SQLTable    = "sienaDealList"
	Transaction_SQLSearchID = "SienaReference"
	Transaction_QueryString = "Ref"
	///
	/// Handler Defintions
	///
	Transaction_Template     = "Transaction"
	Transaction_TemplateList = "/Transaction/Transaction_List"
	Transaction_TemplateView = "/Transaction/Transaction_View"
	Transaction_TemplateEdit = "/Transaction/Transaction_Edit"
	Transaction_TemplateNew  = "/Transaction/Transaction_New"
	///
	/// Handler Monitor Paths
	///
	Transaction_Path       = "/API/Transaction/"
	Transaction_PathList   = "/TransactionList/"
	Transaction_PathView   = "/TransactionView/"
	Transaction_PathEdit   = "/TransactionEdit/"
	Transaction_PathNew    = "/TransactionNew/"
	Transaction_PathSave   = "/TransactionSave/"
	Transaction_PathDelete = "/TransactionDelete/"
	///
	///
	/// SQL Field Definitions
	///
Transaction_SienaReference_sql   = "SienaReference" // SienaReference is a String
Transaction_Status_sql   = "Status" // Status is a String
Transaction_ValueDate_sql   = "ValueDate" // ValueDate is a Time
Transaction_MaturityDate_sql   = "MaturityDate" // MaturityDate is a Time
Transaction_ContractNumber_sql   = "ContractNumber" // ContractNumber is a String
Transaction_ExternalReference_sql   = "ExternalReference" // ExternalReference is a String
Transaction_Book_sql   = "Book" // Book is a String
Transaction_MandatedUser_sql   = "MandatedUser" // MandatedUser is a String
Transaction_Portfolio_sql   = "Portfolio" // Portfolio is a String
Transaction_AgreementId_sql   = "AgreementId" // AgreementId is a Int
Transaction_BackOfficeRefNo_sql   = "BackOfficeRefNo" // BackOfficeRefNo is a String
Transaction_ISIN_sql   = "ISIN" // ISIN is a String
Transaction_UTI_sql   = "UTI" // UTI is a String
Transaction_BookName_sql   = "BookName" // BookName is a String
Transaction_Centre_sql   = "Centre" // Centre is a String
Transaction_Firm_sql   = "Firm" // Firm is a String
Transaction_DealTypeShortName_sql   = "DealTypeShortName" // DealTypeShortName is a String
Transaction_FullDealType_sql   = "FullDealType" // FullDealType is a String
Transaction_TradeDate_sql   = "TradeDate" // TradeDate is a Time
Transaction_DealtCcy_sql   = "DealtCcy" // DealtCcy is a String
Transaction_DealtAmount_sql   = "DealtAmount" // DealtAmount is a Float
Transaction_AgainstAmount_sql   = "AgainstAmount" // AgainstAmount is a Float
Transaction_AgainstCcy_sql   = "AgainstCcy" // AgainstCcy is a String
Transaction_AllInRate_sql   = "AllInRate" // AllInRate is a Float
Transaction_MktRate_sql   = "MktRate" // MktRate is a Float
Transaction_SettleCcy_sql   = "SettleCcy" // SettleCcy is a String
Transaction_Direction_sql   = "Direction" // Direction is a String
Transaction_NpvRate_sql   = "NpvRate" // NpvRate is a Float
Transaction_OriginUser_sql   = "OriginUser" // OriginUser is a String
Transaction_PayInstruction_sql   = "PayInstruction" // PayInstruction is a String
Transaction_ReceiptInstruction_sql   = "ReceiptInstruction" // ReceiptInstruction is a String
Transaction_NIName_sql   = "NIName" // NIName is a String
Transaction_CCYPair_sql   = "CCYPair" // CCYPair is a String
Transaction_Instrument_sql   = "Instrument" // Instrument is a String
Transaction_PortfolioName_sql   = "PortfolioName" // PortfolioName is a String
Transaction_RVDate_sql   = "RVDate" // RVDate is a Time
Transaction_RVMTM_sql   = "RVMTM" // RVMTM is a Float
Transaction_CounterBook_sql   = "CounterBook" // CounterBook is a String
Transaction_CounterBookName_sql   = "CounterBookName" // CounterBookName is a String
Transaction_Party_sql   = "Party" // Party is a String
Transaction_PartyName_sql   = "PartyName" // PartyName is a String
Transaction_NameCentre_sql   = "NameCentre" // NameCentre is a String
Transaction_NameFirm_sql   = "NameFirm" // NameFirm is a String
Transaction_CustomerExternalView_sql   = "CustomerExternalView" // CustomerExternalView is a String
Transaction_CustomerSienaView_sql   = "CustomerSienaView" // CustomerSienaView is a String
Transaction_CompID_sql   = "CompID" // CompID is a String
Transaction_SienaDealer_sql   = "SienaDealer" // SienaDealer is a String
Transaction_DealOwner_sql   = "DealOwner" // DealOwner is a String
Transaction_DealOwnerMnemonic_sql   = "DealOwnerMnemonic" // DealOwnerMnemonic is a String
Transaction_EditedByUser_sql   = "EditedByUser" // EditedByUser is a String
Transaction_UTCOriginTime_sql   = "UTCOriginTime" // UTCOriginTime is a String
Transaction_UTCUpdateTime_sql   = "UTCUpdateTime" // UTCUpdateTime is a String
Transaction_MarginTrading_sql   = "MarginTrading" // MarginTrading is a Bool
Transaction_SwapPoints_sql   = "SwapPoints" // SwapPoints is a Float
Transaction_SpotDate_sql   = "SpotDate" // SpotDate is a Time
Transaction_SpotRate_sql   = "SpotRate" // SpotRate is a Float
Transaction_MktSpotRate_sql   = "MktSpotRate" // MktSpotRate is a Float
Transaction_SpotSalesMargin_sql   = "SpotSalesMargin" // SpotSalesMargin is a Float
Transaction_SpotChlMargin_sql   = "SpotChlMargin" // SpotChlMargin is a Float
Transaction_RerouteCcy_sql   = "RerouteCcy" // RerouteCcy is a String
Transaction_CustomerPayInstruction_sql   = "CustomerPayInstruction" // CustomerPayInstruction is a String
Transaction_CustomerReceiptInstruction_sql   = "CustomerReceiptInstruction" // CustomerReceiptInstruction is a String
Transaction_BackOfficeNotes_sql   = "BackOfficeNotes" // BackOfficeNotes is a String
Transaction_CustomerStatementNotes_sql   = "customerStatementNotes" // CustomerStatementNotes is a String
Transaction_NotesMargin_sql   = "NotesMargin" // NotesMargin is a Float
Transaction_RequestedBy_sql   = "RequestedBy" // RequestedBy is a String
Transaction_EditReason_sql   = "EditReason" // EditReason is a String
Transaction_EditOtherReason_sql   = "EditOtherReason" // EditOtherReason is a String
Transaction_NICleanPrice_sql   = "NICleanPrice" // NICleanPrice is a Float
Transaction_NIDirtyPrice_sql   = "NIDirtyPrice" // NIDirtyPrice is a Float
Transaction_NIYield_sql   = "NIYield" // NIYield is a Float
Transaction_NIClearingSystem_sql   = "NIClearingSystem" // NIClearingSystem is a String
Transaction_Acceptor_sql   = "Acceptor" // Acceptor is a String
Transaction_NIDiscount_sql   = "NIDiscount" // NIDiscount is a Float
Transaction_FastPay_sql   = "FastPay" // FastPay is a Bool
Transaction_PaymentFee_sql   = "PaymentFee" // PaymentFee is a Float
Transaction_PaymentFeePolicy_sql   = "PaymentFeePolicy" // PaymentFeePolicy is a String
Transaction_PaymentReason_sql   = "PaymentReason" // PaymentReason is a String
Transaction_PaymentDate_sql   = "PaymentDate" // PaymentDate is a Time
Transaction_SettlementDate_sql   = "SettlementDate" // SettlementDate is a Time
Transaction_FixingDate_sql   = "FixingDate" // FixingDate is a Time
Transaction_VenueUTI_sql   = "VenueUTI" // VenueUTI is a String
Transaction_EditVersion_sql   = "EditVersion" // EditVersion is a Int
Transaction_BrokeragePercentage_sql   = "BrokeragePercentage" // BrokeragePercentage is a Float
Transaction_BrokerageAmount_sql   = "BrokerageAmount" // BrokerageAmount is a Float
Transaction_BrokerageCurrency_sql   = "BrokerageCurrency" // BrokerageCurrency is a String
Transaction_BrokerageDate_sql   = "BrokerageDate" // BrokerageDate is a Time
Transaction_AccountName_sql   = "AccountName" // AccountName is a String
Transaction_AccountNumber_sql   = "AccountNumber" // AccountNumber is a String
Transaction_CashBalance_sql   = "CashBalance" // CashBalance is a Float
Transaction_DebitFrequency_sql   = "DebitFrequency" // DebitFrequency is a String
Transaction_CreditFrequency_sql   = "CreditFrequency" // CreditFrequency is a String
Transaction_ManuallyQuoted_sql   = "ManuallyQuoted" // ManuallyQuoted is a String
Transaction_LedgerBalance_sql   = "LedgerBalance" // LedgerBalance is a Float
Transaction_SettAmtOutstanding_sql   = "SettAmtOutstanding" // SettAmtOutstanding is a Float
Transaction_FeePercentage_sql   = "FeePercentage" // FeePercentage is a Float
Transaction_FeeAmount_sql   = "FeeAmount" // FeeAmount is a Float
Transaction_Venue_sql   = "Venue" // Venue is a String
Transaction_EURAmount_sql   = "EURAmount" // EURAmount is a Float
Transaction_EUROtherAmount_sql   = "EUROtherAmount" // EUROtherAmount is a Float
Transaction_LEI_sql   = "LEI" // LEI is a String
Transaction_Equity_sql   = "Equity" // Equity is a String
Transaction_Shares_sql   = "Shares" // Shares is a Int
Transaction_QuoteExpiryDate_sql   = "QuoteExpiryDate" // QuoteExpiryDate is a Time
Transaction_Commodity_sql   = "Commodity" // Commodity is a String
Transaction_PaymentSystemSienaView_sql   = "PaymentSystemSienaView" // PaymentSystemSienaView is a String
Transaction_PaymentSystemExternalView_sql   = "PaymentSystemExternalView" // PaymentSystemExternalView is a String
Transaction_SalesProfit_sql   = "SalesProfit" // SalesProfit is a Float
Transaction_RejectReason_sql   = "RejectReason" // RejectReason is a String
Transaction_PaymentError_sql   = "PaymentError" // PaymentError is a String
Transaction_RepoPrincipal_sql   = "RepoPrincipal" // RepoPrincipal is a Float
Transaction_FixingFrequency_sql   = "FixingFrequency" // FixingFrequency is a String
Transaction_Dealt_sql   = "Dealt" // Dealt is a String
Transaction_Against_sql   = "Against" // Against is a String

	/// Definitions End

	/// Application Field Definitions
	///
Transaction_SienaReference_scrn   = "SienaReference" // SienaReference is a String
Transaction_Status_scrn   = "Status" // Status is a String
Transaction_ValueDate_scrn   = "ValueDate" // ValueDate is a Time
Transaction_MaturityDate_scrn   = "MaturityDate" // MaturityDate is a Time
Transaction_ContractNumber_scrn   = "ContractNumber" // ContractNumber is a String
Transaction_ExternalReference_scrn   = "ExternalReference" // ExternalReference is a String
Transaction_Book_scrn   = "Book" // Book is a String
Transaction_MandatedUser_scrn   = "MandatedUser" // MandatedUser is a String
Transaction_Portfolio_scrn   = "Portfolio" // Portfolio is a String
Transaction_AgreementId_scrn   = "AgreementId" // AgreementId is a Int
Transaction_BackOfficeRefNo_scrn   = "BackOfficeRefNo" // BackOfficeRefNo is a String
Transaction_ISIN_scrn   = "ISIN" // ISIN is a String
Transaction_UTI_scrn   = "UTI" // UTI is a String
Transaction_BookName_scrn   = "BookName" // BookName is a String
Transaction_Centre_scrn   = "Centre" // Centre is a String
Transaction_Firm_scrn   = "Firm" // Firm is a String
Transaction_DealTypeShortName_scrn   = "DealTypeShortName" // DealTypeShortName is a String
Transaction_FullDealType_scrn   = "FullDealType" // FullDealType is a String
Transaction_TradeDate_scrn   = "TradeDate" // TradeDate is a Time
Transaction_DealtCcy_scrn   = "DealtCcy" // DealtCcy is a String
Transaction_DealtAmount_scrn   = "DealtAmount" // DealtAmount is a Float
Transaction_AgainstAmount_scrn   = "AgainstAmount" // AgainstAmount is a Float
Transaction_AgainstCcy_scrn   = "AgainstCcy" // AgainstCcy is a String
Transaction_AllInRate_scrn   = "AllInRate" // AllInRate is a Float
Transaction_MktRate_scrn   = "MktRate" // MktRate is a Float
Transaction_SettleCcy_scrn   = "SettleCcy" // SettleCcy is a String
Transaction_Direction_scrn   = "Direction" // Direction is a String
Transaction_NpvRate_scrn   = "NpvRate" // NpvRate is a Float
Transaction_OriginUser_scrn   = "OriginUser" // OriginUser is a String
Transaction_PayInstruction_scrn   = "PayInstruction" // PayInstruction is a String
Transaction_ReceiptInstruction_scrn   = "ReceiptInstruction" // ReceiptInstruction is a String
Transaction_NIName_scrn   = "NIName" // NIName is a String
Transaction_CCYPair_scrn   = "CCYPair" // CCYPair is a String
Transaction_Instrument_scrn   = "Instrument" // Instrument is a String
Transaction_PortfolioName_scrn   = "PortfolioName" // PortfolioName is a String
Transaction_RVDate_scrn   = "RVDate" // RVDate is a Time
Transaction_RVMTM_scrn   = "RVMTM" // RVMTM is a Float
Transaction_CounterBook_scrn   = "CounterBook" // CounterBook is a String
Transaction_CounterBookName_scrn   = "CounterBookName" // CounterBookName is a String
Transaction_Party_scrn   = "Party" // Party is a String
Transaction_PartyName_scrn   = "PartyName" // PartyName is a String
Transaction_NameCentre_scrn   = "NameCentre" // NameCentre is a String
Transaction_NameFirm_scrn   = "NameFirm" // NameFirm is a String
Transaction_CustomerExternalView_scrn   = "CustomerExternalView" // CustomerExternalView is a String
Transaction_CustomerSienaView_scrn   = "CustomerSienaView" // CustomerSienaView is a String
Transaction_CompID_scrn   = "CompID" // CompID is a String
Transaction_SienaDealer_scrn   = "SienaDealer" // SienaDealer is a String
Transaction_DealOwner_scrn   = "DealOwner" // DealOwner is a String
Transaction_DealOwnerMnemonic_scrn   = "DealOwnerMnemonic" // DealOwnerMnemonic is a String
Transaction_EditedByUser_scrn   = "EditedByUser" // EditedByUser is a String
Transaction_UTCOriginTime_scrn   = "UTCOriginTime" // UTCOriginTime is a String
Transaction_UTCUpdateTime_scrn   = "UTCUpdateTime" // UTCUpdateTime is a String
Transaction_MarginTrading_scrn   = "MarginTrading" // MarginTrading is a Bool
Transaction_SwapPoints_scrn   = "SwapPoints" // SwapPoints is a Float
Transaction_SpotDate_scrn   = "SpotDate" // SpotDate is a Time
Transaction_SpotRate_scrn   = "SpotRate" // SpotRate is a Float
Transaction_MktSpotRate_scrn   = "MktSpotRate" // MktSpotRate is a Float
Transaction_SpotSalesMargin_scrn   = "SpotSalesMargin" // SpotSalesMargin is a Float
Transaction_SpotChlMargin_scrn   = "SpotChlMargin" // SpotChlMargin is a Float
Transaction_RerouteCcy_scrn   = "RerouteCcy" // RerouteCcy is a String
Transaction_CustomerPayInstruction_scrn   = "CustomerPayInstruction" // CustomerPayInstruction is a String
Transaction_CustomerReceiptInstruction_scrn   = "CustomerReceiptInstruction" // CustomerReceiptInstruction is a String
Transaction_BackOfficeNotes_scrn   = "BackOfficeNotes" // BackOfficeNotes is a String
Transaction_CustomerStatementNotes_scrn   = "CustomerStatementNotes" // CustomerStatementNotes is a String
Transaction_NotesMargin_scrn   = "NotesMargin" // NotesMargin is a Float
Transaction_RequestedBy_scrn   = "RequestedBy" // RequestedBy is a String
Transaction_EditReason_scrn   = "EditReason" // EditReason is a String
Transaction_EditOtherReason_scrn   = "EditOtherReason" // EditOtherReason is a String
Transaction_NICleanPrice_scrn   = "NICleanPrice" // NICleanPrice is a Float
Transaction_NIDirtyPrice_scrn   = "NIDirtyPrice" // NIDirtyPrice is a Float
Transaction_NIYield_scrn   = "NIYield" // NIYield is a Float
Transaction_NIClearingSystem_scrn   = "NIClearingSystem" // NIClearingSystem is a String
Transaction_Acceptor_scrn   = "Acceptor" // Acceptor is a String
Transaction_NIDiscount_scrn   = "NIDiscount" // NIDiscount is a Float
Transaction_FastPay_scrn   = "FastPay" // FastPay is a Bool
Transaction_PaymentFee_scrn   = "PaymentFee" // PaymentFee is a Float
Transaction_PaymentFeePolicy_scrn   = "PaymentFeePolicy" // PaymentFeePolicy is a String
Transaction_PaymentReason_scrn   = "PaymentReason" // PaymentReason is a String
Transaction_PaymentDate_scrn   = "PaymentDate" // PaymentDate is a Time
Transaction_SettlementDate_scrn   = "SettlementDate" // SettlementDate is a Time
Transaction_FixingDate_scrn   = "FixingDate" // FixingDate is a Time
Transaction_VenueUTI_scrn   = "VenueUTI" // VenueUTI is a String
Transaction_EditVersion_scrn   = "EditVersion" // EditVersion is a Int
Transaction_BrokeragePercentage_scrn   = "BrokeragePercentage" // BrokeragePercentage is a Float
Transaction_BrokerageAmount_scrn   = "BrokerageAmount" // BrokerageAmount is a Float
Transaction_BrokerageCurrency_scrn   = "BrokerageCurrency" // BrokerageCurrency is a String
Transaction_BrokerageDate_scrn   = "BrokerageDate" // BrokerageDate is a Time
Transaction_AccountName_scrn   = "AccountName" // AccountName is a String
Transaction_AccountNumber_scrn   = "AccountNumber" // AccountNumber is a String
Transaction_CashBalance_scrn   = "CashBalance" // CashBalance is a Float
Transaction_DebitFrequency_scrn   = "DebitFrequency" // DebitFrequency is a String
Transaction_CreditFrequency_scrn   = "CreditFrequency" // CreditFrequency is a String
Transaction_ManuallyQuoted_scrn   = "ManuallyQuoted" // ManuallyQuoted is a String
Transaction_LedgerBalance_scrn   = "LedgerBalance" // LedgerBalance is a Float
Transaction_SettAmtOutstanding_scrn   = "SettAmtOutstanding" // SettAmtOutstanding is a Float
Transaction_FeePercentage_scrn   = "FeePercentage" // FeePercentage is a Float
Transaction_FeeAmount_scrn   = "FeeAmount" // FeeAmount is a Float
Transaction_Venue_scrn   = "Venue" // Venue is a String
Transaction_EURAmount_scrn   = "EURAmount" // EURAmount is a Float
Transaction_EUROtherAmount_scrn   = "EUROtherAmount" // EUROtherAmount is a Float
Transaction_LEI_scrn   = "LEI" // LEI is a String
Transaction_Equity_scrn   = "Equity" // Equity is a String
Transaction_Shares_scrn   = "Shares" // Shares is a Int
Transaction_QuoteExpiryDate_scrn   = "QuoteExpiryDate" // QuoteExpiryDate is a Time
Transaction_Commodity_scrn   = "Commodity" // Commodity is a String
Transaction_PaymentSystemSienaView_scrn   = "PaymentSystemSienaView" // PaymentSystemSienaView is a String
Transaction_PaymentSystemExternalView_scrn   = "PaymentSystemExternalView" // PaymentSystemExternalView is a String
Transaction_SalesProfit_scrn   = "SalesProfit" // SalesProfit is a Float
Transaction_RejectReason_scrn   = "RejectReason" // RejectReason is a String
Transaction_PaymentError_scrn   = "PaymentError" // PaymentError is a String
Transaction_RepoPrincipal_scrn   = "RepoPrincipal" // RepoPrincipal is a Float
Transaction_FixingFrequency_scrn   = "FixingFrequency" // FixingFrequency is a String
Transaction_Dealt_scrn   = "Dealt" // Dealt is a String
Transaction_Against_scrn   = "Against" // Against is a String

	/// Definitions End
)
