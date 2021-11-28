package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/account.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Account (account)
// Endpoint 	        : Account (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 28/11/2021 at 22:54:51
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
Firm        string
CCYDp        string
CCY_Impl        string
Book_Impl        string
Portfolio_Impl        string
Centre_Impl        string
Firm_Impl        string

}

const (
	Account_Title       = "Accounts"
	Account_SQLTable    = "sienaAccount"
	Account_SQLSearchID = "SienaReference"
	Account_QueryString = "AccountNo"
	///
	/// Handler Defintions
	///
	Account_TemplateList = "Account_List"
	Account_TemplateView = "Account_View"
	Account_TemplateEdit = "Account_Edit"
	Account_TemplateNew  = "Account_New"
	///
	/// Handler Monitor Paths
	///
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
	Account_Firm   = "Firm" // Firm is a String
	Account_CCYDp   = "CCYDp" // CCYDp is a Int
	Account_CCY_Impl   = "CCY_Impl" // CCY_Impl is a String
	Account_Book_Impl   = "Book_Impl" // Book_Impl is a String
	Account_Portfolio_Impl   = "Portfolio_Impl" // Portfolio_Impl is a String
	Account_Centre_Impl   = "Centre_Impl" // Centre_Impl is a String
	Account_Firm_Impl   = "Firm_Impl" // Firm_Impl is a String

	/// Definitions End
)
