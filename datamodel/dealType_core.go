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
// Date & Time		    : 28/06/2022 at 16:10:51
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//DealType defines the datamolde for the DealType object
type DealType struct {


DealTypeKey       string
DealTypeKey_props FieldProperties
DealTypeShortName       string
DealTypeShortName_props FieldProperties
HostKey       string
HostKey_props FieldProperties
IsActive       string
IsActive_props FieldProperties
Interbook       string
Interbook_props FieldProperties
BackOfficeLink       string
BackOfficeLink_props FieldProperties
HasTicket       string
HasTicket_props FieldProperties
CurrencyOverride       string
CurrencyOverride_props FieldProperties
CurrencyHolderCurrency       string
CurrencyHolderCurrency_props FieldProperties
AllBooks       string
AllBooks_props FieldProperties
FundamentalDealTypeKey       string
FundamentalDealTypeKey_props FieldProperties
RelatedDealType       string
RelatedDealType_props FieldProperties
BookName       string
BookName_props FieldProperties
ExportMethod       string
ExportMethod_props FieldProperties
DefaultUserLayoffBooks       string
DefaultUserLayoffBooks_props FieldProperties
RFQ       string
RFQ_props FieldProperties
OBS       string
OBS_props FieldProperties
KID       string
KID_props FieldProperties
InternalId       string
InternalId_props FieldProperties
InternalDeleted       string
InternalDeleted_props FieldProperties
UpdatedTransactionId       string
UpdatedTransactionId_props FieldProperties
UpdatedUserId       string
UpdatedUserId_props FieldProperties
UpdatedDateTime       string
UpdatedDateTime_props FieldProperties
DeletedTransactionId       string
DeletedTransactionId_props FieldProperties
DeletedUserId       string
DeletedUserId_props FieldProperties
ChangeType       string
ChangeType_props FieldProperties
 // Any lookups will be added below


























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
	DealType_TemplateList = "/DealType/DealType_List"
	DealType_TemplateView = "/DealType/DealType_View"
	DealType_TemplateEdit = "/DealType/DealType_Edit"
	DealType_TemplateNew  = "/DealType/DealType_New"
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
	//DealType_Redirect provides a page to return to aftern an action
	DealType_Redirect = DealType_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	DealType_DealTypeKey_sql   = "DealTypeKey" // DealTypeKey is a String
	DealType_DealTypeShortName_sql   = "DealTypeShortName" // DealTypeShortName is a String
	DealType_HostKey_sql   = "HostKey" // HostKey is a String
	DealType_IsActive_sql   = "IsActive" // IsActive is a Bool
	DealType_Interbook_sql   = "Interbook" // Interbook is a Bool
	DealType_BackOfficeLink_sql   = "BackOfficeLink" // BackOfficeLink is a Bool
	DealType_HasTicket_sql   = "HasTicket" // HasTicket is a Bool
	DealType_CurrencyOverride_sql   = "CurrencyOverride" // CurrencyOverride is a Bool
	DealType_CurrencyHolderCurrency_sql   = "CurrencyHolderCurrency" // CurrencyHolderCurrency is a String
	DealType_AllBooks_sql   = "AllBooks" // AllBooks is a Bool
	DealType_FundamentalDealTypeKey_sql   = "FundamentalDealTypeKey" // FundamentalDealTypeKey is a String
	DealType_RelatedDealType_sql   = "RelatedDealType" // RelatedDealType is a String
	DealType_BookName_sql   = "BookName" // BookName is a String
	DealType_ExportMethod_sql   = "ExportMethod" // ExportMethod is a String
	DealType_DefaultUserLayoffBooks_sql   = "DefaultUserLayoffBooks" // DefaultUserLayoffBooks is a Bool
	DealType_RFQ_sql   = "RFQ" // RFQ is a Bool
	DealType_OBS_sql   = "OBS" // OBS is a Bool
	DealType_KID_sql   = "KID" // KID is a Bool
	DealType_InternalId_sql   = "InternalId" // InternalId is a Int
	DealType_InternalDeleted_sql   = "InternalDeleted" // InternalDeleted is a Time
	DealType_UpdatedTransactionId_sql   = "UpdatedTransactionId" // UpdatedTransactionId is a String
	DealType_UpdatedUserId_sql   = "UpdatedUserId" // UpdatedUserId is a String
	DealType_UpdatedDateTime_sql   = "UpdatedDateTime" // UpdatedDateTime is a Time
	DealType_DeletedTransactionId_sql   = "DeletedTransactionId" // DeletedTransactionId is a String
	DealType_DeletedUserId_sql   = "DeletedUserId" // DeletedUserId is a String
	DealType_ChangeType_sql   = "ChangeType" // ChangeType is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	DealType_DealTypeKey_scrn   = "DealTypeKey" // DealTypeKey is a String
	DealType_DealTypeShortName_scrn   = "DealTypeShortName" // DealTypeShortName is a String
	DealType_HostKey_scrn   = "HostKey" // HostKey is a String
	DealType_IsActive_scrn   = "IsActive" // IsActive is a Bool
	DealType_Interbook_scrn   = "Interbook" // Interbook is a Bool
	DealType_BackOfficeLink_scrn   = "BackOfficeLink" // BackOfficeLink is a Bool
	DealType_HasTicket_scrn   = "HasTicket" // HasTicket is a Bool
	DealType_CurrencyOverride_scrn   = "CurrencyOverride" // CurrencyOverride is a Bool
	DealType_CurrencyHolderCurrency_scrn   = "CurrencyHolderCurrency" // CurrencyHolderCurrency is a String
	DealType_AllBooks_scrn   = "AllBooks" // AllBooks is a Bool
	DealType_FundamentalDealTypeKey_scrn   = "FundamentalDealTypeKey" // FundamentalDealTypeKey is a String
	DealType_RelatedDealType_scrn   = "RelatedDealType" // RelatedDealType is a String
	DealType_BookName_scrn   = "BookName" // BookName is a String
	DealType_ExportMethod_scrn   = "ExportMethod" // ExportMethod is a String
	DealType_DefaultUserLayoffBooks_scrn   = "DefaultUserLayoffBooks" // DefaultUserLayoffBooks is a Bool
	DealType_RFQ_scrn   = "RFQ" // RFQ is a Bool
	DealType_OBS_scrn   = "OBS" // OBS is a Bool
	DealType_KID_scrn   = "KID" // KID is a Bool
	DealType_InternalId_scrn   = "InternalId" // InternalId is a Int
	DealType_InternalDeleted_scrn   = "InternalDeleted" // InternalDeleted is a Time
	DealType_UpdatedTransactionId_scrn   = "UpdatedTransactionId" // UpdatedTransactionId is a String
	DealType_UpdatedUserId_scrn   = "UpdatedUserId" // UpdatedUserId is a String
	DealType_UpdatedDateTime_scrn   = "UpdatedDateTime" // UpdatedDateTime is a Time
	DealType_DeletedTransactionId_scrn   = "DeletedTransactionId" // DeletedTransactionId is a String
	DealType_DeletedUserId_scrn   = "DeletedUserId" // DeletedUserId is a String
	DealType_ChangeType_scrn   = "ChangeType" // ChangeType is a String

	/// Definitions End
	///
)

//dealtype_PageList provides the information for the template for a list of DealTypes
type DealType_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []DealType
	Context	 appContext
}

//dealtype_Page provides the information for the template for an individual DealType
type DealType_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	DealTypeKey         string
	DealTypeKey_props     FieldProperties
	DealTypeShortName         string
	DealTypeShortName_props     FieldProperties
	HostKey         string
	HostKey_props     FieldProperties
	IsActive         string
	IsActive_props     FieldProperties
	Interbook         string
	Interbook_props     FieldProperties
	BackOfficeLink         string
	BackOfficeLink_props     FieldProperties
	HasTicket         string
	HasTicket_props     FieldProperties
	CurrencyOverride         string
	CurrencyOverride_props     FieldProperties
	CurrencyHolderCurrency         string
	CurrencyHolderCurrency_props     FieldProperties
	AllBooks         string
	AllBooks_props     FieldProperties
	FundamentalDealTypeKey         string
	FundamentalDealTypeKey_props     FieldProperties
	RelatedDealType         string
	RelatedDealType_props     FieldProperties
	BookName         string
	BookName_props     FieldProperties
	ExportMethod         string
	ExportMethod_props     FieldProperties
	DefaultUserLayoffBooks         string
	DefaultUserLayoffBooks_props     FieldProperties
	RFQ         string
	RFQ_props     FieldProperties
	OBS         string
	OBS_props     FieldProperties
	KID         string
	KID_props     FieldProperties
	InternalId         string
	InternalId_props     FieldProperties
	InternalDeleted         string
	InternalDeleted_props     FieldProperties
	UpdatedTransactionId         string
	UpdatedTransactionId_props     FieldProperties
	UpdatedUserId         string
	UpdatedUserId_props     FieldProperties
	UpdatedDateTime         string
	UpdatedDateTime_props     FieldProperties
	DeletedTransactionId         string
	DeletedTransactionId_props     FieldProperties
	DeletedUserId         string
	DeletedUserId_props     FieldProperties
	ChangeType         string
	ChangeType_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}