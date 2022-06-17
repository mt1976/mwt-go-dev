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
// Date & Time		    : 17/06/2022 at 18:38:05
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Account defines the datamolde for the Account object
type Account struct {

SienaReference       string
CustomerSienaView       string
SienaCommonRef       string
Status       string
StartDate       string
MaturityDate       string
ContractNumber       string
ExternalReference       string
CCY       string
CCY_lookup []Lookup_Item
Book       string
Book_lookup []Lookup_Item
MandatedUser       string
BackOfficeNotes       string
CashBalance       string
AccountNumber       string
AccountName       string
LedgerBalance       string
Portfolio       string
Portfolio_lookup []Lookup_Item
AgreementId       string
BackOfficeRefNo       string
ISIN       string
UTI       string
CCYName       string
BookName       string
PortfolioName       string
Centre       string
Centre_lookup []Lookup_Item
DealTypeKey       string
DealTypeShortName       string
InternalId       string
InternalDeleted       string
UpdatedTransactionId       string
UpdatedUserId       string
UpdatedDateTime       string
DeletedTransactionId       string
DeletedUserId       string
ChangeType       string
CCYDp       string
CompID       string
Firm       string
Firm_lookup []Lookup_Item
DealType       string
FullDealType       string
DealingInterface       string
DealtAmount       string
ParentContractNumber       string
InterestFrequency       string
InterestAction       string
InterestStrategy       string
InterestBasis       string
SienaDealer       string
DealOwner       string
OriginUser       string
EditedByUser       string
DealOwnerMnemonic       string
UTCOriginTime       string
UTCUpdateTime       string
CustomerStatementNotes       string
NotesMargin       string
RequestedBy       string
EditReason       string
EditOtherReason       string
NoticeDays       string
DebitFrequency       string
CreditFrequency       string
EURAmount       string
EUROtherAmount       string
PaymentSystemSienaView       string
PaymentSystemExternalView       string
DealtCA       string
AgainstCA       string
LedgerCA       string
CashBalanceCA       string

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
	Account_TemplateList = "/Account/Account_List"
	Account_TemplateView = "/Account/Account_View"
	Account_TemplateEdit = "/Account/Account_Edit"
	Account_TemplateNew  = "/Account/Account_New"
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
	///
	/// SQL Field Definitions
	///
Account_SienaReference_sql   = "SienaReference" // SienaReference is a String
Account_CustomerSienaView_sql   = "CustomerSienaView" // CustomerSienaView is a String
Account_SienaCommonRef_sql   = "SienaCommonRef" // SienaCommonRef is a String
Account_Status_sql   = "Status" // Status is a String
Account_StartDate_sql   = "StartDate" // StartDate is a Time
Account_MaturityDate_sql   = "MaturityDate" // MaturityDate is a Time
Account_ContractNumber_sql   = "ContractNumber" // ContractNumber is a String
Account_ExternalReference_sql   = "ExternalReference" // ExternalReference is a String
Account_CCY_sql   = "CCY" // CCY is a String
Account_Book_sql   = "Book" // Book is a String
Account_MandatedUser_sql   = "MandatedUser" // MandatedUser is a String
Account_BackOfficeNotes_sql   = "BackOfficeNotes" // BackOfficeNotes is a String
Account_CashBalance_sql   = "CashBalance" // CashBalance is a Float
Account_AccountNumber_sql   = "AccountNumber" // AccountNumber is a String
Account_AccountName_sql   = "AccountName" // AccountName is a String
Account_LedgerBalance_sql   = "LedgerBalance" // LedgerBalance is a Float
Account_Portfolio_sql   = "Portfolio" // Portfolio is a String
Account_AgreementId_sql   = "AgreementId" // AgreementId is a Int
Account_BackOfficeRefNo_sql   = "BackOfficeRefNo" // BackOfficeRefNo is a String
Account_ISIN_sql   = "ISIN" // ISIN is a String
Account_UTI_sql   = "UTI" // UTI is a String
Account_CCYName_sql   = "CCYName" // CCYName is a String
Account_BookName_sql   = "BookName" // BookName is a String
Account_PortfolioName_sql   = "PortfolioName" // PortfolioName is a String
Account_Centre_sql   = "Centre" // Centre is a String
Account_DealTypeKey_sql   = "DealTypeKey" // DealTypeKey is a String
Account_DealTypeShortName_sql   = "DealTypeShortName" // DealTypeShortName is a String
Account_InternalId_sql   = "InternalId" // InternalId is a Int
Account_InternalDeleted_sql   = "InternalDeleted" // InternalDeleted is a Time
Account_UpdatedTransactionId_sql   = "UpdatedTransactionId" // UpdatedTransactionId is a String
Account_UpdatedUserId_sql   = "UpdatedUserId" // UpdatedUserId is a String
Account_UpdatedDateTime_sql   = "UpdatedDateTime" // UpdatedDateTime is a Time
Account_DeletedTransactionId_sql   = "DeletedTransactionId" // DeletedTransactionId is a String
Account_DeletedUserId_sql   = "DeletedUserId" // DeletedUserId is a String
Account_ChangeType_sql   = "ChangeType" // ChangeType is a String
Account_CCYDp_sql   = "CCYDp" // CCYDp is a Int
Account_CompID_sql   = "CompID" // CompID is a String
Account_Firm_sql   = "Firm" // Firm is a String
Account_DealType_sql   = "DealType" // DealType is a String
Account_FullDealType_sql   = "FullDealType" // FullDealType is a String
Account_DealingInterface_sql   = "DealingInterface" // DealingInterface is a String
Account_DealtAmount_sql   = "DealtAmount" // DealtAmount is a Float
Account_ParentContractNumber_sql   = "ParentContractNumber" // ParentContractNumber is a String
Account_InterestFrequency_sql   = "InterestFrequency" // InterestFrequency is a String
Account_InterestAction_sql   = "InterestAction" // InterestAction is a String
Account_InterestStrategy_sql   = "InterestStrategy" // InterestStrategy is a String
Account_InterestBasis_sql   = "InterestBasis" // InterestBasis is a String
Account_SienaDealer_sql   = "SienaDealer" // SienaDealer is a String
Account_DealOwner_sql   = "DealOwner" // DealOwner is a String
Account_OriginUser_sql   = "OriginUser" // OriginUser is a String
Account_EditedByUser_sql   = "EditedByUser" // EditedByUser is a String
Account_DealOwnerMnemonic_sql   = "DealOwnerMnemonic" // DealOwnerMnemonic is a String
Account_UTCOriginTime_sql   = "UTCOriginTime" // UTCOriginTime is a String
Account_UTCUpdateTime_sql   = "UTCUpdateTime" // UTCUpdateTime is a String
Account_CustomerStatementNotes_sql   = "customerStatementNotes" // CustomerStatementNotes is a String
Account_NotesMargin_sql   = "NotesMargin" // NotesMargin is a Float
Account_RequestedBy_sql   = "RequestedBy" // RequestedBy is a String
Account_EditReason_sql   = "EditReason" // EditReason is a String
Account_EditOtherReason_sql   = "EditOtherReason" // EditOtherReason is a String
Account_NoticeDays_sql   = "NoticeDays" // NoticeDays is a Int
Account_DebitFrequency_sql   = "DebitFrequency" // DebitFrequency is a String
Account_CreditFrequency_sql   = "CreditFrequency" // CreditFrequency is a String
Account_EURAmount_sql   = "EURAmount" // EURAmount is a Float
Account_EUROtherAmount_sql   = "EUROtherAmount" // EUROtherAmount is a Float
Account_PaymentSystemSienaView_sql   = "PaymentSystemSienaView" // PaymentSystemSienaView is a String
Account_PaymentSystemExternalView_sql   = "PaymentSystemExternalView" // PaymentSystemExternalView is a String
Account_DealtCA_sql   = "DealtCA" // DealtCA is a String
Account_AgainstCA_sql   = "AgainstCA" // AgainstCA is a String
Account_LedgerCA_sql   = "LedgerCA" // LedgerCA is a String
Account_CashBalanceCA_sql   = "CashBalanceCA" // CashBalanceCA is a String

	/// Definitions End

	/// Application Field Definitions
	///
Account_SienaReference_scrn   = "SienaReference" // SienaReference is a String
Account_CustomerSienaView_scrn   = "CustomerSienaView" // CustomerSienaView is a String
Account_SienaCommonRef_scrn   = "SienaCommonRef" // SienaCommonRef is a String
Account_Status_scrn   = "Status" // Status is a String
Account_StartDate_scrn   = "StartDate" // StartDate is a Time
Account_MaturityDate_scrn   = "MaturityDate" // MaturityDate is a Time
Account_ContractNumber_scrn   = "ContractNumber" // ContractNumber is a String
Account_ExternalReference_scrn   = "ExternalReference" // ExternalReference is a String
Account_CCY_scrn   = "CCY" // CCY is a String
Account_Book_scrn   = "Book" // Book is a String
Account_MandatedUser_scrn   = "MandatedUser" // MandatedUser is a String
Account_BackOfficeNotes_scrn   = "BackOfficeNotes" // BackOfficeNotes is a String
Account_CashBalance_scrn   = "CashBalance" // CashBalance is a Float
Account_AccountNumber_scrn   = "AccountNumber" // AccountNumber is a String
Account_AccountName_scrn   = "AccountName" // AccountName is a String
Account_LedgerBalance_scrn   = "LedgerBalance" // LedgerBalance is a Float
Account_Portfolio_scrn   = "Portfolio" // Portfolio is a String
Account_AgreementId_scrn   = "AgreementId" // AgreementId is a Int
Account_BackOfficeRefNo_scrn   = "BackOfficeRefNo" // BackOfficeRefNo is a String
Account_ISIN_scrn   = "ISIN" // ISIN is a String
Account_UTI_scrn   = "UTI" // UTI is a String
Account_CCYName_scrn   = "CCYName" // CCYName is a String
Account_BookName_scrn   = "BookName" // BookName is a String
Account_PortfolioName_scrn   = "PortfolioName" // PortfolioName is a String
Account_Centre_scrn   = "Centre" // Centre is a String
Account_DealTypeKey_scrn   = "DealTypeKey" // DealTypeKey is a String
Account_DealTypeShortName_scrn   = "DealTypeShortName" // DealTypeShortName is a String
Account_InternalId_scrn   = "InternalId" // InternalId is a Int
Account_InternalDeleted_scrn   = "InternalDeleted" // InternalDeleted is a Time
Account_UpdatedTransactionId_scrn   = "UpdatedTransactionId" // UpdatedTransactionId is a String
Account_UpdatedUserId_scrn   = "UpdatedUserId" // UpdatedUserId is a String
Account_UpdatedDateTime_scrn   = "UpdatedDateTime" // UpdatedDateTime is a Time
Account_DeletedTransactionId_scrn   = "DeletedTransactionId" // DeletedTransactionId is a String
Account_DeletedUserId_scrn   = "DeletedUserId" // DeletedUserId is a String
Account_ChangeType_scrn   = "ChangeType" // ChangeType is a String
Account_CCYDp_scrn   = "CCYDp" // CCYDp is a Int
Account_CompID_scrn   = "CompID" // CompID is a String
Account_Firm_scrn   = "Firm" // Firm is a String
Account_DealType_scrn   = "DealType" // DealType is a String
Account_FullDealType_scrn   = "FullDealType" // FullDealType is a String
Account_DealingInterface_scrn   = "DealingInterface" // DealingInterface is a String
Account_DealtAmount_scrn   = "DealtAmount" // DealtAmount is a Float
Account_ParentContractNumber_scrn   = "ParentContractNumber" // ParentContractNumber is a String
Account_InterestFrequency_scrn   = "InterestFrequency" // InterestFrequency is a String
Account_InterestAction_scrn   = "InterestAction" // InterestAction is a String
Account_InterestStrategy_scrn   = "InterestStrategy" // InterestStrategy is a String
Account_InterestBasis_scrn   = "InterestBasis" // InterestBasis is a String
Account_SienaDealer_scrn   = "SienaDealer" // SienaDealer is a String
Account_DealOwner_scrn   = "DealOwner" // DealOwner is a String
Account_OriginUser_scrn   = "OriginUser" // OriginUser is a String
Account_EditedByUser_scrn   = "EditedByUser" // EditedByUser is a String
Account_DealOwnerMnemonic_scrn   = "DealOwnerMnemonic" // DealOwnerMnemonic is a String
Account_UTCOriginTime_scrn   = "UTCOriginTime" // UTCOriginTime is a String
Account_UTCUpdateTime_scrn   = "UTCUpdateTime" // UTCUpdateTime is a String
Account_CustomerStatementNotes_scrn   = "CustomerStatementNotes" // CustomerStatementNotes is a String
Account_NotesMargin_scrn   = "NotesMargin" // NotesMargin is a Float
Account_RequestedBy_scrn   = "RequestedBy" // RequestedBy is a String
Account_EditReason_scrn   = "EditReason" // EditReason is a String
Account_EditOtherReason_scrn   = "EditOtherReason" // EditOtherReason is a String
Account_NoticeDays_scrn   = "NoticeDays" // NoticeDays is a Int
Account_DebitFrequency_scrn   = "DebitFrequency" // DebitFrequency is a String
Account_CreditFrequency_scrn   = "CreditFrequency" // CreditFrequency is a String
Account_EURAmount_scrn   = "EURAmount" // EURAmount is a Float
Account_EUROtherAmount_scrn   = "EUROtherAmount" // EUROtherAmount is a Float
Account_PaymentSystemSienaView_scrn   = "PaymentSystemSienaView" // PaymentSystemSienaView is a String
Account_PaymentSystemExternalView_scrn   = "PaymentSystemExternalView" // PaymentSystemExternalView is a String
Account_DealtCA_scrn   = "DealtCA" // DealtCA is a String
Account_AgainstCA_scrn   = "AgainstCA" // AgainstCA is a String
Account_LedgerCA_scrn   = "LedgerCA" // LedgerCA is a String
Account_CashBalanceCA_scrn   = "CashBalanceCA" // CashBalanceCA is a String

	/// Definitions End
)
