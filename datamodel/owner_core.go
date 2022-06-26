package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/owner.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Owner (owner)
// Endpoint 	        : Owner (Owner)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:30
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Owner defines the datamolde for the Owner object
type Owner struct {


UserName       string
UserName_props FieldProperties
FullName       string
FullName_props FieldProperties
Type       string
Type_props FieldProperties
TradingEntity       string
TradingEntity_props FieldProperties
DefaultEnterBook       string
DefaultEnterBook_props FieldProperties
EmailAddress       string
EmailAddress_props FieldProperties
Enabled       string
Enabled_props FieldProperties
ExternalUserIds       string
ExternalUserIds_props FieldProperties
Language       string
Language_props FieldProperties
LocalCurrency       string
LocalCurrency_props FieldProperties
Role       string
Role_props FieldProperties
TelephoneNumber       string
TelephoneNumber_props FieldProperties
TokenId       string
TokenId_props FieldProperties
Entity       string
Entity_props FieldProperties
UserCode       string
UserCode_props FieldProperties
 // Any lookups will be added below















}

const (
	Owner_Title       = "Owner"
	Owner_SQLTable    = "sienaOwner"
	Owner_SQLSearchID = "UserName"
	Owner_QueryString = "Owner"
	///
	/// Handler Defintions
	///
	Owner_Template     = "Owner"
	Owner_TemplateList = "/Owner/Owner_List"
	Owner_TemplateView = "/Owner/Owner_View"
	Owner_TemplateEdit = "/Owner/Owner_Edit"
	Owner_TemplateNew  = "/Owner/Owner_New"
	///
	/// Handler Monitor Paths
	///
	Owner_Path       = "/API/Owner/"
	Owner_PathList   = "/OwnerList/"
	Owner_PathView   = "/OwnerView/"
	Owner_PathEdit   = "/OwnerEdit/"
	Owner_PathNew    = "/OwnerNew/"
	Owner_PathSave   = "/OwnerSave/"
	Owner_PathDelete = "/OwnerDelete/"
	///
	///
	/// SQL Field Definitions
	///
Owner_UserName_sql   = "UserName" // UserName is a String
Owner_FullName_sql   = "FullName" // FullName is a String
Owner_Type_sql   = "Type" // Type is a String
Owner_TradingEntity_sql   = "TradingEntity" // TradingEntity is a String
Owner_DefaultEnterBook_sql   = "DefaultEnterBook" // DefaultEnterBook is a String
Owner_EmailAddress_sql   = "EmailAddress" // EmailAddress is a String
Owner_Enabled_sql   = "Enabled" // Enabled is a String
Owner_ExternalUserIds_sql   = "ExternalUserIds" // ExternalUserIds is a String
Owner_Language_sql   = "Language" // Language is a String
Owner_LocalCurrency_sql   = "LocalCurrency" // LocalCurrency is a String
Owner_Role_sql   = "Role" // Role is a String
Owner_TelephoneNumber_sql   = "TelephoneNumber" // TelephoneNumber is a String
Owner_TokenId_sql   = "TokenId" // TokenId is a String
Owner_Entity_sql   = "Entity" // Entity is a String
Owner_UserCode_sql   = "UserCode" // UserCode is a String

	/// Definitions End

	/// Application Field Definitions
	///
Owner_UserName_scrn   = "UserName" // UserName is a String
Owner_FullName_scrn   = "FullName" // FullName is a String
Owner_Type_scrn   = "Type" // Type is a String
Owner_TradingEntity_scrn   = "TradingEntity" // TradingEntity is a String
Owner_DefaultEnterBook_scrn   = "DefaultEnterBook" // DefaultEnterBook is a String
Owner_EmailAddress_scrn   = "EmailAddress" // EmailAddress is a String
Owner_Enabled_scrn   = "Enabled" // Enabled is a String
Owner_ExternalUserIds_scrn   = "ExternalUserIds" // ExternalUserIds is a String
Owner_Language_scrn   = "Language" // Language is a String
Owner_LocalCurrency_scrn   = "LocalCurrency" // LocalCurrency is a String
Owner_Role_scrn   = "Role" // Role is a String
Owner_TelephoneNumber_scrn   = "TelephoneNumber" // TelephoneNumber is a String
Owner_TokenId_scrn   = "TokenId" // TokenId is a String
Owner_Entity_scrn   = "Entity" // Entity is a String
Owner_UserCode_scrn   = "UserCode" // UserCode is a String

	/// Definitions End
)
