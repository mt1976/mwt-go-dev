package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/payee.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Payee (payee)
// Endpoint 	        : Payee (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 03/12/2021 at 13:16:59
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Payee struct {

SourceTable        string
KeyCounterpartyFirm        string
KeyCounterpartyCentre        string
KeyCurrency        string
KeyName        string
KeyNumber        string
KeyDirection        string
KeyType        string
FullName        string
Address        string
PhoneNo        string
Country        string
Bic        string
Iban        string
AccountNo        string
FedWireNo        string
SortCode        string
BankName        string
BankPinCode        string
BankAddress        string
Reason        string
BankSettlementAcct        string
UpdatedUserId        string
Country_Impl        string
Firm_Impl        string
Centre_Impl        string
Currency_Impl        string

}

const (
	Payee_Title       = "Payee"
	Payee_SQLTable    = "sienaCounterpartyPayee"
	Payee_SQLSearchID = "KeyCounterpartyFirm"
	Payee_QueryString = "ID"
	///
	/// Handler Defintions
	///
	Payee_TemplateList = "Payee_List"
	Payee_TemplateView = "Payee_View"
	Payee_TemplateEdit = "Payee_Edit"
	Payee_TemplateNew  = "Payee_New"
	///
	/// Handler Monitor Paths
	///
	Payee_PathList   = "/PayeeList/"
	Payee_PathView   = "/PayeeView/"
	Payee_PathEdit   = "/PayeeEdit/"
	Payee_PathNew    = "/PayeeNew/"
	Payee_PathSave   = "/PayeeSave/"
	Payee_PathDelete = "/PayeeDelete/"
	///
	/// SQL Field Definitions
	///
	Payee_SourceTable   = "SourceTable" // SourceTable is a String
	Payee_KeyCounterpartyFirm   = "KeyCounterpartyFirm" // KeyCounterpartyFirm is a String
	Payee_KeyCounterpartyCentre   = "KeyCounterpartyCentre" // KeyCounterpartyCentre is a String
	Payee_KeyCurrency   = "KeyCurrency" // KeyCurrency is a String
	Payee_KeyName   = "KeyName" // KeyName is a String
	Payee_KeyNumber   = "KeyNumber" // KeyNumber is a String
	Payee_KeyDirection   = "KeyDirection" // KeyDirection is a String
	Payee_KeyType   = "KeyType" // KeyType is a String
	Payee_FullName   = "FullName" // FullName is a String
	Payee_Address   = "Address" // Address is a String
	Payee_PhoneNo   = "PhoneNo" // PhoneNo is a String
	Payee_Country   = "Country" // Country is a String
	Payee_Bic   = "Bic" // Bic is a String
	Payee_Iban   = "Iban" // Iban is a String
	Payee_AccountNo   = "AccountNo" // AccountNo is a String
	Payee_FedWireNo   = "FedWireNo" // FedWireNo is a String
	Payee_SortCode   = "SortCode" // SortCode is a String
	Payee_BankName   = "BankName" // BankName is a String
	Payee_BankPinCode   = "BankPinCode" // BankPinCode is a String
	Payee_BankAddress   = "BankAddress" // BankAddress is a String
	Payee_Reason   = "Reason" // Reason is a String
	Payee_BankSettlementAcct   = "BankSettlementAcct" // BankSettlementAcct is a Bool
	Payee_UpdatedUserId   = "UpdatedUserId" // UpdatedUserId is a String
	Payee_Country_Impl   = "Country_Impl" // Country_Impl is a String
	Payee_Firm_Impl   = "Firm_Impl" // Firm_Impl is a String
	Payee_Centre_Impl   = "Centre_Impl" // Centre_Impl is a String
	Payee_Currency_Impl   = "Currency_Impl" // Currency_Impl is a String

	/// Definitions End
)
