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
// Date & Time		    : 26/06/2022 at 18:48:31
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Payee defines the datamolde for the Payee object
type Payee struct {


SourceTable       string
SourceTable_props FieldProperties
KeyCounterpartyFirm       string
KeyCounterpartyFirm_props FieldProperties
KeyCounterpartyCentre       string
KeyCounterpartyCentre_props FieldProperties
KeyCurrency       string
KeyCurrency_props FieldProperties
KeyName       string
KeyName_props FieldProperties
KeyNumber       string
KeyNumber_props FieldProperties
KeyDirection       string
KeyDirection_props FieldProperties
KeyType       string
KeyType_props FieldProperties
FullName       string
FullName_props FieldProperties
Address       string
Address_props FieldProperties
PhoneNo       string
PhoneNo_props FieldProperties
Country       string
Country_props FieldProperties
Bic       string
Bic_props FieldProperties
Iban       string
Iban_props FieldProperties
AccountNo       string
AccountNo_props FieldProperties
FedWireNo       string
FedWireNo_props FieldProperties
SortCode       string
SortCode_props FieldProperties
BankName       string
BankName_props FieldProperties
BankPinCode       string
BankPinCode_props FieldProperties
BankAddress       string
BankAddress_props FieldProperties
Reason       string
Reason_props FieldProperties
BankSettlementAcct       string
BankSettlementAcct_props FieldProperties
UpdatedUserId       string
UpdatedUserId_props FieldProperties
Status       string
Status_props FieldProperties
 // Any lookups will be added below










Country_lookup []Lookup_Item













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
	Payee_TemplateList = "/Payee/Payee_List"
	Payee_TemplateView = "/Payee/Payee_View"
	Payee_TemplateEdit = "/Payee/Payee_Edit"
	Payee_TemplateNew  = "/Payee/Payee_New"
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
	///
	/// SQL Field Definitions
	///
Payee_SourceTable_sql   = "SourceTable" // SourceTable is a String
Payee_KeyCounterpartyFirm_sql   = "KeyCounterpartyFirm" // KeyCounterpartyFirm is a String
Payee_KeyCounterpartyCentre_sql   = "KeyCounterpartyCentre" // KeyCounterpartyCentre is a String
Payee_KeyCurrency_sql   = "KeyCurrency" // KeyCurrency is a String
Payee_KeyName_sql   = "KeyName" // KeyName is a String
Payee_KeyNumber_sql   = "KeyNumber" // KeyNumber is a String
Payee_KeyDirection_sql   = "KeyDirection" // KeyDirection is a String
Payee_KeyType_sql   = "KeyType" // KeyType is a String
Payee_FullName_sql   = "FullName" // FullName is a String
Payee_Address_sql   = "Address" // Address is a String
Payee_PhoneNo_sql   = "PhoneNo" // PhoneNo is a String
Payee_Country_sql   = "Country" // Country is a String
Payee_Bic_sql   = "Bic" // Bic is a String
Payee_Iban_sql   = "Iban" // Iban is a String
Payee_AccountNo_sql   = "AccountNo" // AccountNo is a String
Payee_FedWireNo_sql   = "FedWireNo" // FedWireNo is a String
Payee_SortCode_sql   = "SortCode" // SortCode is a String
Payee_BankName_sql   = "BankName" // BankName is a String
Payee_BankPinCode_sql   = "BankPinCode" // BankPinCode is a String
Payee_BankAddress_sql   = "BankAddress" // BankAddress is a String
Payee_Reason_sql   = "Reason" // Reason is a String
Payee_BankSettlementAcct_sql   = "BankSettlementAcct" // BankSettlementAcct is a Bool
Payee_UpdatedUserId_sql   = "UpdatedUserId" // UpdatedUserId is a String
Payee_Status_sql   = "Status" // Status is a String

	/// Definitions End

	/// Application Field Definitions
	///
Payee_SourceTable_scrn   = "SourceTable" // SourceTable is a String
Payee_KeyCounterpartyFirm_scrn   = "KeyCounterpartyFirm" // KeyCounterpartyFirm is a String
Payee_KeyCounterpartyCentre_scrn   = "KeyCounterpartyCentre" // KeyCounterpartyCentre is a String
Payee_KeyCurrency_scrn   = "KeyCurrency" // KeyCurrency is a String
Payee_KeyName_scrn   = "KeyName" // KeyName is a String
Payee_KeyNumber_scrn   = "KeyNumber" // KeyNumber is a String
Payee_KeyDirection_scrn   = "KeyDirection" // KeyDirection is a String
Payee_KeyType_scrn   = "KeyType" // KeyType is a String
Payee_FullName_scrn   = "FullName" // FullName is a String
Payee_Address_scrn   = "Address" // Address is a String
Payee_PhoneNo_scrn   = "PhoneNo" // PhoneNo is a String
Payee_Country_scrn   = "Country" // Country is a String
Payee_Bic_scrn   = "Bic" // Bic is a String
Payee_Iban_scrn   = "Iban" // Iban is a String
Payee_AccountNo_scrn   = "AccountNo" // AccountNo is a String
Payee_FedWireNo_scrn   = "FedWireNo" // FedWireNo is a String
Payee_SortCode_scrn   = "SortCode" // SortCode is a String
Payee_BankName_scrn   = "BankName" // BankName is a String
Payee_BankPinCode_scrn   = "BankPinCode" // BankPinCode is a String
Payee_BankAddress_scrn   = "BankAddress" // BankAddress is a String
Payee_Reason_scrn   = "Reason" // Reason is a String
Payee_BankSettlementAcct_scrn   = "BankSettlementAcct" // BankSettlementAcct is a Bool
Payee_UpdatedUserId_scrn   = "UpdatedUserId" // UpdatedUserId is a String
Payee_Status_scrn   = "Status" // Status is a String

	/// Definitions End
)
