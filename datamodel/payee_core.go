package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/payee.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Payee (payee)
// Endpoint 	        : Payee (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:32:07
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Payee defines the datamolde for the Payee object
type Payee struct {

SourceTable       string
KeyCounterpartyFirm       string
KeyCounterpartyCentre       string
KeyCurrency       string
KeyName       string
KeyNumber       string
KeyDirection       string
KeyType       string
FullName       string
Address       string
PhoneNo       string
Country       string
Country_lookup []Lookup_Item
Bic       string
Iban       string
AccountNo       string
FedWireNo       string
SortCode       string
BankName       string
BankPinCode       string
BankAddress       string
Reason       string
BankSettlementAcct       string
UpdatedUserId       string
Status       string

}

const (
	Payee_Title       = "Payee"
	Payee_SQLTable    = "sienaCounterpartyPayee"
	Payee_SQLSearchID = "KeyCounterpartyFirm"
	Payee_QueryString = "ID"
	///
	/// Handler Defintions
	///
	Payee_Template     = "Payee"
	Payee_TemplateList = "Payee_List"
	Payee_TemplateView = "Payee_View"
	Payee_TemplateEdit = "Payee_Edit"
	Payee_TemplateNew  = "Payee_New"
	///
	/// Handler Monitor Paths
	///
	Payee_Path       = "/API/Payee/"
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
Payee_Status   = "Status" // Status is a String

	/// Definitions End
)
