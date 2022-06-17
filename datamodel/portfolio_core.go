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
// Date & Time		    : 17/06/2022 at 18:38:13
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Portfolio defines the datamolde for the Portfolio object
type Portfolio struct {

Code       string
Description1       string
Description2       string
IsDefault       string
InternalId       string
InternalDeleted       string
UpdatedTransactionId       string
UpdatedUserId       string
UpdatedDateTime       string
DeletedTransactionId       string
DeletedUserId       string
ChangeType       string

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
)
