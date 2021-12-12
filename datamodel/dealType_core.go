package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/dealtype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DealType (dealtype)
// Endpoint 	        : DealType (DealTypeKey)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:14
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type DealType struct {

DealTypeKey        string
DealTypeShortName        string
HostKey        string
IsActive        string
Interbook        string
BackOfficeLink        string
HasTicket        string
CurrencyOverride        string
CurrencyHolderCurrency        string
AllBooks        string
FundamentalDealTypeKey        string
RelatedDealType        string
BookName        string
ExportMethod        string
DefaultUserLayoffBooks        string
RFQ        string
OBS        string
KID        string
InternalId        string
InternalDeleted        string
UpdatedTransactionId        string
UpdatedUserId        string
UpdatedDateTime        string
DeletedTransactionId        string
DeletedUserId        string
ChangeType        string

}

const (
	DealType_Title       = "Deal Type"
	DealType_SQLTable    = "sienaDealType"
	DealType_SQLSearchID = "DealTypeKey"
	DealType_QueryString = "DealTypeKey"
	///
	/// Handler Defintions
	///
	DealType_Template     = "DealType"
	DealType_TemplateList = "DealType_List"
	DealType_TemplateView = "DealType_View"
	DealType_TemplateEdit = "DealType_Edit"
	DealType_TemplateNew  = "DealType_New"
	///
	/// Handler Monitor Paths
	///
	DealType_Path       = "/API/DealType/"
	DealType_PathList   = "/DealTypeList/"
	DealType_PathView   = "/DealTypeView/"
	DealType_PathEdit   = "/DealTypeEdit/"
	DealType_PathNew    = "/DealTypeNew/"
	DealType_PathSave   = "/DealTypeSave/"
	DealType_PathDelete = "/DealTypeDelete/"
	///
	/// SQL Field Definitions
	///
	DealType_DealTypeKey   = "DealTypeKey" // DealTypeKey is a String
	DealType_DealTypeShortName   = "DealTypeShortName" // DealTypeShortName is a String
	DealType_HostKey   = "HostKey" // HostKey is a String
	DealType_IsActive   = "IsActive" // IsActive is a Bool
	DealType_Interbook   = "Interbook" // Interbook is a Bool
	DealType_BackOfficeLink   = "BackOfficeLink" // BackOfficeLink is a Bool
	DealType_HasTicket   = "HasTicket" // HasTicket is a Bool
	DealType_CurrencyOverride   = "CurrencyOverride" // CurrencyOverride is a Bool
	DealType_CurrencyHolderCurrency   = "CurrencyHolderCurrency" // CurrencyHolderCurrency is a String
	DealType_AllBooks   = "AllBooks" // AllBooks is a Bool
	DealType_FundamentalDealTypeKey   = "FundamentalDealTypeKey" // FundamentalDealTypeKey is a String
	DealType_RelatedDealType   = "RelatedDealType" // RelatedDealType is a String
	DealType_BookName   = "BookName" // BookName is a String
	DealType_ExportMethod   = "ExportMethod" // ExportMethod is a String
	DealType_DefaultUserLayoffBooks   = "DefaultUserLayoffBooks" // DefaultUserLayoffBooks is a Bool
	DealType_RFQ   = "RFQ" // RFQ is a Bool
	DealType_OBS   = "OBS" // OBS is a Bool
	DealType_KID   = "KID" // KID is a Bool
	DealType_InternalId   = "InternalId" // InternalId is a Int
	DealType_InternalDeleted   = "InternalDeleted" // InternalDeleted is a Time
	DealType_UpdatedTransactionId   = "UpdatedTransactionId" // UpdatedTransactionId is a String
	DealType_UpdatedUserId   = "UpdatedUserId" // UpdatedUserId is a String
	DealType_UpdatedDateTime   = "UpdatedDateTime" // UpdatedDateTime is a Time
	DealType_DeletedTransactionId   = "DeletedTransactionId" // DeletedTransactionId is a String
	DealType_DeletedUserId   = "DeletedUserId" // DeletedUserId is a String
	DealType_ChangeType   = "ChangeType" // ChangeType is a String

	/// Definitions End
)
