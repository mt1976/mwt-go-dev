package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/portfolio.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Portfolio (portfolio)
// Endpoint 	        : Portfolio (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:55
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Portfolio defines the datamolde for the Portfolio object
type Portfolio struct {


Code       string
Code_props FieldProperties
Description1       string
Description1_props FieldProperties
Description2       string
Description2_props FieldProperties
IsDefault       string
IsDefault_props FieldProperties
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
	Portfolio_Title       = "Bank Porfolio"
	Portfolio_SQLTable    = "sienaPortfolio"
	Portfolio_SQLSearchID = "Code"
	Portfolio_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Portfolio_Template     = "Portfolio"
	Portfolio_TemplateList = "/Portfolio/Portfolio_List"
	Portfolio_TemplateView = "/Portfolio/Portfolio_View"
	Portfolio_TemplateEdit = "/Portfolio/Portfolio_Edit"
	Portfolio_TemplateNew  = "/Portfolio/Portfolio_New"
	///
	/// Handler Monitor Paths
	///
	Portfolio_Path       = "/API/Portfolio/"
	Portfolio_PathList   = "/PortfolioList/"
	Portfolio_PathView   = "/PortfolioView/"
	Portfolio_PathEdit   = "/PortfolioEdit/"
	Portfolio_PathNew    = "/PortfolioNew/"
	Portfolio_PathSave   = "/PortfolioSave/"
	Portfolio_PathDelete = "/PortfolioDelete/"
	///
	//Portfolio_Redirect provides a page to return to aftern an action
	Portfolio_Redirect = Portfolio_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	Portfolio_Code_sql   = "Code" // Code is a String
	Portfolio_Description1_sql   = "Description1" // Description1 is a String
	Portfolio_Description2_sql   = "Description2" // Description2 is a String
	Portfolio_IsDefault_sql   = "isDefault" // IsDefault is a Bool
	Portfolio_InternalId_sql   = "InternalId" // InternalId is a Int
	Portfolio_InternalDeleted_sql   = "InternalDeleted" // InternalDeleted is a Time
	Portfolio_UpdatedTransactionId_sql   = "UpdatedTransactionId" // UpdatedTransactionId is a String
	Portfolio_UpdatedUserId_sql   = "UpdatedUserId" // UpdatedUserId is a String
	Portfolio_UpdatedDateTime_sql   = "UpdatedDateTime" // UpdatedDateTime is a Time
	Portfolio_DeletedTransactionId_sql   = "DeletedTransactionId" // DeletedTransactionId is a String
	Portfolio_DeletedUserId_sql   = "DeletedUserId" // DeletedUserId is a String
	Portfolio_ChangeType_sql   = "ChangeType" // ChangeType is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Portfolio_Code_scrn   = "Code" // Code is a String
	Portfolio_Description1_scrn   = "Description1" // Description1 is a String
	Portfolio_Description2_scrn   = "Description2" // Description2 is a String
	Portfolio_IsDefault_scrn   = "IsDefault" // IsDefault is a Bool
	Portfolio_InternalId_scrn   = "InternalId" // InternalId is a Int
	Portfolio_InternalDeleted_scrn   = "InternalDeleted" // InternalDeleted is a Time
	Portfolio_UpdatedTransactionId_scrn   = "UpdatedTransactionId" // UpdatedTransactionId is a String
	Portfolio_UpdatedUserId_scrn   = "UpdatedUserId" // UpdatedUserId is a String
	Portfolio_UpdatedDateTime_scrn   = "UpdatedDateTime" // UpdatedDateTime is a Time
	Portfolio_DeletedTransactionId_scrn   = "DeletedTransactionId" // DeletedTransactionId is a String
	Portfolio_DeletedUserId_scrn   = "DeletedUserId" // DeletedUserId is a String
	Portfolio_ChangeType_scrn   = "ChangeType" // ChangeType is a String

	/// Definitions End
	///
)

//portfolio_PageList provides the information for the template for a list of Portfolios
type Portfolio_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Portfolio
	Context	 appContext
}

//portfolio_Page provides the information for the template for an individual Portfolio
type Portfolio_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	Code         string
	Code_props     FieldProperties
	Description1         string
	Description1_props     FieldProperties
	Description2         string
	Description2_props     FieldProperties
	IsDefault         string
	IsDefault_props     FieldProperties
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