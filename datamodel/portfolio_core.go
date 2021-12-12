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
// Date & Time		    : 12/12/2021 at 16:13:17
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Portfolio struct {

Code        string
Description1        string
Description2        string
IsDefault        string
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
	Portfolio_Title       = "Bank Porfolio"
	Portfolio_SQLTable    = "sienaPortfolio"
	Portfolio_SQLSearchID = "Code"
	Portfolio_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Portfolio_Template     = "Portfolio"
	Portfolio_TemplateList = "Portfolio_List"
	Portfolio_TemplateView = "Portfolio_View"
	Portfolio_TemplateEdit = "Portfolio_Edit"
	Portfolio_TemplateNew  = "Portfolio_New"
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
	/// SQL Field Definitions
	///
	Portfolio_Code   = "Code" // Code is a String
	Portfolio_Description1   = "Description1" // Description1 is a String
	Portfolio_Description2   = "Description2" // Description2 is a String
	Portfolio_IsDefault   = "isDefault" // IsDefault is a Bool
	Portfolio_InternalId   = "InternalId" // InternalId is a Int
	Portfolio_InternalDeleted   = "InternalDeleted" // InternalDeleted is a Time
	Portfolio_UpdatedTransactionId   = "UpdatedTransactionId" // UpdatedTransactionId is a String
	Portfolio_UpdatedUserId   = "UpdatedUserId" // UpdatedUserId is a String
	Portfolio_UpdatedDateTime   = "UpdatedDateTime" // UpdatedDateTime is a Time
	Portfolio_DeletedTransactionId   = "DeletedTransactionId" // DeletedTransactionId is a String
	Portfolio_DeletedUserId   = "DeletedUserId" // DeletedUserId is a String
	Portfolio_ChangeType   = "ChangeType" // ChangeType is a String

	/// Definitions End
)
