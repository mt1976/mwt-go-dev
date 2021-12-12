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
// Date & Time		    : 12/12/2021 at 16:13:17
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Owner struct {

UserName        string
FullName        string
Type        string
TradingEntity        string
DefaultEnterBook        string
EmailAddress        string
Enabled        string
ExternalUserIds        string
Language        string
LocalCurrency        string
Role        string
TelephoneNumber        string
TokenId        string
Entity        string
UserCode        string

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
	Owner_TemplateList = "Owner_List"
	Owner_TemplateView = "Owner_View"
	Owner_TemplateEdit = "Owner_Edit"
	Owner_TemplateNew  = "Owner_New"
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
	/// SQL Field Definitions
	///
	Owner_UserName   = "UserName" // UserName is a String
	Owner_FullName   = "FullName" // FullName is a String
	Owner_Type   = "Type" // Type is a String
	Owner_TradingEntity   = "TradingEntity" // TradingEntity is a String
	Owner_DefaultEnterBook   = "DefaultEnterBook" // DefaultEnterBook is a String
	Owner_EmailAddress   = "EmailAddress" // EmailAddress is a String
	Owner_Enabled   = "Enabled" // Enabled is a String
	Owner_ExternalUserIds   = "ExternalUserIds" // ExternalUserIds is a String
	Owner_Language   = "Language" // Language is a String
	Owner_LocalCurrency   = "LocalCurrency" // LocalCurrency is a String
	Owner_Role   = "Role" // Role is a String
	Owner_TelephoneNumber   = "TelephoneNumber" // TelephoneNumber is a String
	Owner_TokenId   = "TokenId" // TokenId is a String
	Owner_Entity   = "Entity" // Entity is a String
	Owner_UserCode   = "UserCode" // UserCode is a String

	/// Definitions End
)
