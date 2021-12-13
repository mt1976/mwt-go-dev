package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/account.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Account (account)
// Endpoint 	        : Account (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 13/12/2021 at 17:02:22
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Account struct {

SienaReference        string
CustomerSienaView        string
SienaCommonRef        string
Status        string
StartDate        string
MaturityDate        string
ContractNumber        string
ExternalReference        string
CCY        string
Book        string
MandatedUser        string
BackOfficeNotes        string
CashBalance        string
AccountNumber        string
AccountName        string
LedgerBalance        string
Portfolio        string
AgreementId        string
BackOfficeRefNo        string
ISIN        string
UTI        string
CCYName        string
BookName        string
PortfolioName        string
Centre        string
DealTypeKey        string
DealTypeShortName        string
InternalId        string
InternalDeleted        string
UpdatedTransactionId        string
UpdatedUserId        string
UpdatedDateTime        string
DeletedTransactionId        string
DeletedUserId        string
ChangeType        string
CCYDp        string
CompID        string
Firm        string
DealType        string
FullDealType        string
DealingInterface        string
DealtAmount        string
ParentContractNumber        string
InterestFrequency        string
InterestAction        string
InterestStrategy        string
InterestBasis        string
SienaDealer        string
DealOwner        string
OriginUser        string
EditedByUser        string
DealOwnerMnemonic        string
UTCOriginTime        string
UTCUpdateTime        string
CustomerStatementNotes        string
NotesMargin        string
RequestedBy        string
EditReason        string
EditOtherReason        string
NoticeDays        string
DebitFrequency        string
CreditFrequency        string
EURAmount        string
EUROtherAmount        string
PaymentSystemSienaView        string
PaymentSystemExternalView        string
CCY_Lookup        string
Book_Lookup        string
Portfolio_Lookup        string
Centre_Lookup        string
Firm_Lookup        string
DealtCA_Extra        string
AgainstCA_Extra        string
LedgerCA_Extra        string
CashBalanceCA_Extra        string

}

const (
	Account_Title       = "Account"
	Account_SQLTable    = "sienaAccount"
	Account_SQLSearchID = "SienaReference"
	Account_QueryString = "AccountNo"
	///
	/// Handler Defintions
	///
	Account_Template     = "Account"
	Account_TemplateList = "Account_List"
	Account_TemplateView = "Account_View"
	Account_TemplateEdit = "Account_Edit"
	Account_TemplateNew  = "Account_New"
	///
	/// Handler Monitor Paths
	///
	Account_Path       = "/API/Account/"
	Account_PathList   = "/AccountList/"
	Account_PathView   = "/AccountView/"
	Account_PathEdit   = "/AccountEdit/"
	Account_PathNew    = "/AccountNew/"
	Account_PathSave   = "/AccountSave/"
	Account_PathDelete = "/AccountDelete/"
	///
	/// SQL Field Definitions
	///
	Account_SienaReference   = "SienaReference" // SienaReference is a String
	Account_CustomerSienaView   = "CustomerSienaView" // CustomerSienaView is a String
	Account_SienaCommonRef   = "SienaCommonRef" // SienaCommonRef is a String
	Account_Status   = "Status" // Status is a String
	Account_StartDate   = "StartDate" // StartDate is a Time
	Account_MaturityDate   = "MaturityDate" // MaturityDate is a Time
	Account_ContractNumber   = "ContractNumber" // ContractNumber is a String
	Account_ExternalReference   = "ExternalReference" // ExternalReference is a String
	Account_CCY   = "CCY" // CCY is a String
	Account_Book   = "Book" // Book is a String
	Account_MandatedUser   = "MandatedUser" // MandatedUser is a String
	Account_BackOfficeNotes   = "BackOfficeNotes" // BackOfficeNotes is a String
	Account_CashBalance   = "CashBalance" // CashBalance is a Float
	Account_AccountNumber   = "AccountNumber" // AccountNumber is a String
	Account_AccountName   = "AccountName" // AccountName is a String
	Account_LedgerBalance   = "LedgerBalance" // LedgerBalance is a Float
	Account_Portfolio   = "Portfolio" // Portfolio is a String
	Account_AgreementId   = "AgreementId" // AgreementId is a Int
	Account_BackOfficeRefNo   = "BackOfficeRefNo" // BackOfficeRefNo is a String
	Account_ISIN   = "ISIN" // ISIN is a String
	Account_UTI   = "UTI" // UTI is a String
	Account_CCYName   = "CCYName" // CCYName is a String
	Account_BookName   = "BookName" // BookName is a String
	Account_PortfolioName   = "PortfolioName" // PortfolioName is a String
	Account_Centre   = "Centre" // Centre is a String
	Account_DealTypeKey   = "DealTypeKey" // DealTypeKey is a String
	Account_DealTypeShortName   = "DealTypeShortName" // DealTypeShortName is a String
	Account_InternalId   = "InternalId" // InternalId is a Int
	Account_InternalDeleted   = "InternalDeleted" // InternalDeleted is a Time
	Account_UpdatedTransactionId   = "UpdatedTransactionId" // UpdatedTransactionId is a String
	Account_UpdatedUserId   = "UpdatedUserId" // UpdatedUserId is a String
	Account_UpdatedDateTime   = "UpdatedDateTime" // UpdatedDateTime is a Time
	Account_DeletedTransactionId   = "DeletedTransactionId" // DeletedTransactionId is a String
	Account_DeletedUserId   = "DeletedUserId" // DeletedUserId is a String
	Account_ChangeType   = "ChangeType" // ChangeType is a String
	Account_CCYDp   = "CCYDp" // CCYDp is a Int
	Account_CompID   = "CompID" // CompID is a String
	Account_Firm   = "Firm" // Firm is a String
	Account_DealType   = "DealType" // DealType is a String
	Account_FullDealType   = "FullDealType" // FullDealType is a String
	Account_DealingInterface   = "DealingInterface" // DealingInterface is a String
	Account_DealtAmount   = "DealtAmount" // DealtAmount is a Float
	Account_ParentContractNumber   = "ParentContractNumber" // ParentContractNumber is a String
	Account_InterestFrequency   = "InterestFrequency" // InterestFrequency is a String
	Account_InterestAction   = "InterestAction" // InterestAction is a String
	Account_InterestStrategy   = "InterestStrategy" // InterestStrategy is a String
	Account_InterestBasis   = "InterestBasis" // InterestBasis is a String
	Account_SienaDealer   = "SienaDealer" // SienaDealer is a String
	Account_DealOwner   = "DealOwner" // DealOwner is a String
	Account_OriginUser   = "OriginUser" // OriginUser is a String
	Account_EditedByUser   = "EditedByUser" // EditedByUser is a String
	Account_DealOwnerMnemonic   = "DealOwnerMnemonic" // DealOwnerMnemonic is a String
	Account_UTCOriginTime   = "UTCOriginTime" // UTCOriginTime is a String
	Account_UTCUpdateTime   = "UTCUpdateTime" // UTCUpdateTime is a String
	Account_CustomerStatementNotes   = "customerStatementNotes" // CustomerStatementNotes is a String
	Account_NotesMargin   = "NotesMargin" // NotesMargin is a Float
	Account_RequestedBy   = "RequestedBy" // RequestedBy is a String
	Account_EditReason   = "EditReason" // EditReason is a String
	Account_EditOtherReason   = "EditOtherReason" // EditOtherReason is a String
	Account_NoticeDays   = "NoticeDays" // NoticeDays is a Int
	Account_DebitFrequency   = "DebitFrequency" // DebitFrequency is a String
	Account_CreditFrequency   = "CreditFrequency" // CreditFrequency is a String
	Account_EURAmount   = "EURAmount" // EURAmount is a Float
	Account_EUROtherAmount   = "EUROtherAmount" // EUROtherAmount is a Float
	Account_PaymentSystemSienaView   = "PaymentSystemSienaView" // PaymentSystemSienaView is a String
	Account_PaymentSystemExternalView   = "PaymentSystemExternalView" // PaymentSystemExternalView is a String
	Account_CCY_Lookup   = "CCY_Lookup" // CCY_Lookup is a String
	Account_Book_Lookup   = "Book_Lookup" // Book_Lookup is a String
	Account_Portfolio_Lookup   = "Portfolio_Lookup" // Portfolio_Lookup is a String
	Account_Centre_Lookup   = "Centre_Lookup" // Centre_Lookup is a String
	Account_Firm_Lookup   = "Firm_Lookup" // Firm_Lookup is a String
	Account_DealtCA_Extra   = "DealtCA_Extra" // DealtCA_Extra is a String
	Account_AgainstCA_Extra   = "AgainstCA_Extra" // AgainstCA_Extra is a String
	Account_LedgerCA_Extra   = "LedgerCA_Extra" // LedgerCA_Extra is a String
	Account_CashBalanceCA_Extra   = "CashBalanceCA_Extra" // CashBalanceCA_Extra is a String

	/// Definitions End
)
