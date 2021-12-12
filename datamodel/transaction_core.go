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
// Date & Time		    : 12/12/2021 at 16:13:14
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Transaction struct {

SienaReference        string
Status        string
ValueDate        string
MaturityDate        string
ContractNumber        string
ExternalReference        string
Book        string
MandatedUser        string
Portfolio        string
AgreementId        string
BackOfficeRefNo        string
ISIN        string
UTI        string
BookName        string
Centre        string
Firm        string
DealTypeShortName        string
FullDealType        string
TradeDate        string
DealtCcy        string
DealtAmount        string
AgainstAmount        string
AgainstCcy        string
AllInRate        string
MktRate        string
SettleCcy        string
Direction        string
NpvRate        string
OriginUser        string
PayInstruction        string
ReceiptInstruction        string
NIName        string
CCYPair        string
Instrument        string
PortfolioName        string
RVDate        string
RVMTM        string
CounterBook        string
CounterBookName        string
Party        string
PartyName        string
NameCentre        string
NameFirm        string
CustomerExternalView        string
CustomerSienaView        string
CompID        string
SienaDealer        string
DealOwner        string
DealOwnerMnemonic        string
EditedByUser        string
UTCOriginTime        string
UTCUpdateTime        string
MarginTrading        string
SwapPoints        string
SpotDate        string
SpotRate        string
MktSpotRate        string
SpotSalesMargin        string
SpotChlMargin        string
RerouteCcy        string
CustomerPayInstruction        string
CustomerReceiptInstruction        string
BackOfficeNotes        string
CustomerStatementNotes        string
NotesMargin        string
RequestedBy        string
EditReason        string
EditOtherReason        string
NICleanPrice        string
NIDirtyPrice        string
NIYield        string
NIClearingSystem        string
Acceptor        string
NIDiscount        string
FastPay        string
PaymentFee        string
PaymentFeePolicy        string
PaymentReason        string
PaymentDate        string
SettlementDate        string
FixingDate        string
VenueUTI        string
EditVersion        string
BrokeragePercentage        string
BrokerageAmount        string
BrokerageCurrency        string
BrokerageDate        string
AccountName        string
AccountNumber        string
CashBalance        string
DebitFrequency        string
CreditFrequency        string
ManuallyQuoted        string
LedgerBalance        string
SettAmtOutstanding        string
FeePercentage        string
FeeAmount        string
Venue        string
EURAmount        string
EUROtherAmount        string
LEI        string
Equity        string
Shares        string
QuoteExpiryDate        string
Commodity        string
PaymentSystemSienaView        string
PaymentSystemExternalView        string
SalesProfit        string
RejectReason        string
PaymentError        string
RepoPrincipal        string
FixingFrequency        string

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
	Transaction_TemplateList = "Transaction_List"
	Transaction_TemplateView = "Transaction_View"
	Transaction_TemplateEdit = "Transaction_Edit"
	Transaction_TemplateNew  = "Transaction_New"
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
	/// SQL Field Definitions
	///
	Transaction_SienaReference   = "SienaReference" // SienaReference is a String
	Transaction_Status   = "Status" // Status is a String
	Transaction_ValueDate   = "ValueDate" // ValueDate is a Time
	Transaction_MaturityDate   = "MaturityDate" // MaturityDate is a Time
	Transaction_ContractNumber   = "ContractNumber" // ContractNumber is a String
	Transaction_ExternalReference   = "ExternalReference" // ExternalReference is a String
	Transaction_Book   = "Book" // Book is a String
	Transaction_MandatedUser   = "MandatedUser" // MandatedUser is a String
	Transaction_Portfolio   = "Portfolio" // Portfolio is a String
	Transaction_AgreementId   = "AgreementId" // AgreementId is a Int
	Transaction_BackOfficeRefNo   = "BackOfficeRefNo" // BackOfficeRefNo is a String
	Transaction_ISIN   = "ISIN" // ISIN is a String
	Transaction_UTI   = "UTI" // UTI is a String
	Transaction_BookName   = "BookName" // BookName is a String
	Transaction_Centre   = "Centre" // Centre is a String
	Transaction_Firm   = "Firm" // Firm is a String
	Transaction_DealTypeShortName   = "DealTypeShortName" // DealTypeShortName is a String
	Transaction_FullDealType   = "FullDealType" // FullDealType is a String
	Transaction_TradeDate   = "TradeDate" // TradeDate is a Time
	Transaction_DealtCcy   = "DealtCcy" // DealtCcy is a String
	Transaction_DealtAmount   = "DealtAmount" // DealtAmount is a Float
	Transaction_AgainstAmount   = "AgainstAmount" // AgainstAmount is a Float
	Transaction_AgainstCcy   = "AgainstCcy" // AgainstCcy is a String
	Transaction_AllInRate   = "AllInRate" // AllInRate is a Float
	Transaction_MktRate   = "MktRate" // MktRate is a Float
	Transaction_SettleCcy   = "SettleCcy" // SettleCcy is a String
	Transaction_Direction   = "Direction" // Direction is a String
	Transaction_NpvRate   = "NpvRate" // NpvRate is a Float
	Transaction_OriginUser   = "OriginUser" // OriginUser is a String
	Transaction_PayInstruction   = "PayInstruction" // PayInstruction is a String
	Transaction_ReceiptInstruction   = "ReceiptInstruction" // ReceiptInstruction is a String
	Transaction_NIName   = "NIName" // NIName is a String
	Transaction_CCYPair   = "CCYPair" // CCYPair is a String
	Transaction_Instrument   = "Instrument" // Instrument is a String
	Transaction_PortfolioName   = "PortfolioName" // PortfolioName is a String
	Transaction_RVDate   = "RVDate" // RVDate is a Time
	Transaction_RVMTM   = "RVMTM" // RVMTM is a Float
	Transaction_CounterBook   = "CounterBook" // CounterBook is a String
	Transaction_CounterBookName   = "CounterBookName" // CounterBookName is a String
	Transaction_Party   = "Party" // Party is a String
	Transaction_PartyName   = "PartyName" // PartyName is a String
	Transaction_NameCentre   = "NameCentre" // NameCentre is a String
	Transaction_NameFirm   = "NameFirm" // NameFirm is a String
	Transaction_CustomerExternalView   = "CustomerExternalView" // CustomerExternalView is a String
	Transaction_CustomerSienaView   = "CustomerSienaView" // CustomerSienaView is a String
	Transaction_CompID   = "CompID" // CompID is a String
	Transaction_SienaDealer   = "SienaDealer" // SienaDealer is a String
	Transaction_DealOwner   = "DealOwner" // DealOwner is a String
	Transaction_DealOwnerMnemonic   = "DealOwnerMnemonic" // DealOwnerMnemonic is a String
	Transaction_EditedByUser   = "EditedByUser" // EditedByUser is a String
	Transaction_UTCOriginTime   = "UTCOriginTime" // UTCOriginTime is a String
	Transaction_UTCUpdateTime   = "UTCUpdateTime" // UTCUpdateTime is a String
	Transaction_MarginTrading   = "MarginTrading" // MarginTrading is a Bool
	Transaction_SwapPoints   = "SwapPoints" // SwapPoints is a Float
	Transaction_SpotDate   = "SpotDate" // SpotDate is a Time
	Transaction_SpotRate   = "SpotRate" // SpotRate is a Float
	Transaction_MktSpotRate   = "MktSpotRate" // MktSpotRate is a Float
	Transaction_SpotSalesMargin   = "SpotSalesMargin" // SpotSalesMargin is a Float
	Transaction_SpotChlMargin   = "SpotChlMargin" // SpotChlMargin is a Float
	Transaction_RerouteCcy   = "RerouteCcy" // RerouteCcy is a String
	Transaction_CustomerPayInstruction   = "CustomerPayInstruction" // CustomerPayInstruction is a String
	Transaction_CustomerReceiptInstruction   = "CustomerReceiptInstruction" // CustomerReceiptInstruction is a String
	Transaction_BackOfficeNotes   = "BackOfficeNotes" // BackOfficeNotes is a String
	Transaction_CustomerStatementNotes   = "customerStatementNotes" // CustomerStatementNotes is a String
	Transaction_NotesMargin   = "NotesMargin" // NotesMargin is a Float
	Transaction_RequestedBy   = "RequestedBy" // RequestedBy is a String
	Transaction_EditReason   = "EditReason" // EditReason is a String
	Transaction_EditOtherReason   = "EditOtherReason" // EditOtherReason is a String
	Transaction_NICleanPrice   = "NICleanPrice" // NICleanPrice is a Float
	Transaction_NIDirtyPrice   = "NIDirtyPrice" // NIDirtyPrice is a Float
	Transaction_NIYield   = "NIYield" // NIYield is a Float
	Transaction_NIClearingSystem   = "NIClearingSystem" // NIClearingSystem is a String
	Transaction_Acceptor   = "Acceptor" // Acceptor is a String
	Transaction_NIDiscount   = "NIDiscount" // NIDiscount is a Float
	Transaction_FastPay   = "FastPay" // FastPay is a Bool
	Transaction_PaymentFee   = "PaymentFee" // PaymentFee is a Float
	Transaction_PaymentFeePolicy   = "PaymentFeePolicy" // PaymentFeePolicy is a String
	Transaction_PaymentReason   = "PaymentReason" // PaymentReason is a String
	Transaction_PaymentDate   = "PaymentDate" // PaymentDate is a Time
	Transaction_SettlementDate   = "SettlementDate" // SettlementDate is a Time
	Transaction_FixingDate   = "FixingDate" // FixingDate is a Time
	Transaction_VenueUTI   = "VenueUTI" // VenueUTI is a String
	Transaction_EditVersion   = "EditVersion" // EditVersion is a Int
	Transaction_BrokeragePercentage   = "BrokeragePercentage" // BrokeragePercentage is a Float
	Transaction_BrokerageAmount   = "BrokerageAmount" // BrokerageAmount is a Float
	Transaction_BrokerageCurrency   = "BrokerageCurrency" // BrokerageCurrency is a String
	Transaction_BrokerageDate   = "BrokerageDate" // BrokerageDate is a Time
	Transaction_AccountName   = "AccountName" // AccountName is a String
	Transaction_AccountNumber   = "AccountNumber" // AccountNumber is a String
	Transaction_CashBalance   = "CashBalance" // CashBalance is a Float
	Transaction_DebitFrequency   = "DebitFrequency" // DebitFrequency is a String
	Transaction_CreditFrequency   = "CreditFrequency" // CreditFrequency is a String
	Transaction_ManuallyQuoted   = "ManuallyQuoted" // ManuallyQuoted is a String
	Transaction_LedgerBalance   = "LedgerBalance" // LedgerBalance is a Float
	Transaction_SettAmtOutstanding   = "SettAmtOutstanding" // SettAmtOutstanding is a Float
	Transaction_FeePercentage   = "FeePercentage" // FeePercentage is a Float
	Transaction_FeeAmount   = "FeeAmount" // FeeAmount is a Float
	Transaction_Venue   = "Venue" // Venue is a String
	Transaction_EURAmount   = "EURAmount" // EURAmount is a Float
	Transaction_EUROtherAmount   = "EUROtherAmount" // EUROtherAmount is a Float
	Transaction_LEI   = "LEI" // LEI is a String
	Transaction_Equity   = "Equity" // Equity is a String
	Transaction_Shares   = "Shares" // Shares is a Int
	Transaction_QuoteExpiryDate   = "QuoteExpiryDate" // QuoteExpiryDate is a Time
	Transaction_Commodity   = "Commodity" // Commodity is a String
	Transaction_PaymentSystemSienaView   = "PaymentSystemSienaView" // PaymentSystemSienaView is a String
	Transaction_PaymentSystemExternalView   = "PaymentSystemExternalView" // PaymentSystemExternalView is a String
	Transaction_SalesProfit   = "SalesProfit" // SalesProfit is a Float
	Transaction_RejectReason   = "RejectReason" // RejectReason is a String
	Transaction_PaymentError   = "PaymentError" // PaymentError is a String
	Transaction_RepoPrincipal   = "RepoPrincipal" // RepoPrincipal is a Float
	Transaction_FixingFrequency   = "FixingFrequency" // FixingFrequency is a String

	/// Definitions End
)
