package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/accountladder.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : AccountLadder (accountladder)
// Endpoint 	        : AccountLadder (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:41
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//AccountLadder defines the datamolde for the AccountLadder object
type AccountLadder struct {


SienaReference       string
SienaReference_props FieldProperties
BusinessDate       string
BusinessDate_props FieldProperties
ContractNumber       string
ContractNumber_props FieldProperties
Balance       string
Balance_props FieldProperties
DealtCcy       string
DealtCcy_props FieldProperties
AmountDp       string
AmountDp_props FieldProperties
 // Any lookups will be added below






}

const (
	AccountLadder_Title       = "Account Ladder"
	AccountLadder_SQLTable    = "sienaAccountLadder"
	AccountLadder_SQLSearchID = "SienaReference"
	AccountLadder_QueryString = "AccountNo"
	///
	/// Handler Defintions
	///
	AccountLadder_Template     = "AccountLadder"
	AccountLadder_TemplateList = "/AccountLadder/AccountLadder_List"
	AccountLadder_TemplateView = "/AccountLadder/AccountLadder_View"
	AccountLadder_TemplateEdit = "/AccountLadder/AccountLadder_Edit"
	AccountLadder_TemplateNew  = "/AccountLadder/AccountLadder_New"
	///
	/// Handler Monitor Paths
	///
	AccountLadder_Path       = "/API/AccountLadder/"
	AccountLadder_PathList   = "/AccountLadderList/"
	AccountLadder_PathView   = "/AccountLadderView/"
	AccountLadder_PathEdit   = "/AccountLadderEdit/"
	AccountLadder_PathNew    = "/AccountLadderNew/"
	AccountLadder_PathSave   = "/AccountLadderSave/"
	AccountLadder_PathDelete = "/AccountLadderDelete/"
	///
	//AccountLadder_Redirect provides a page to return to aftern an action
	AccountLadder_Redirect = AccountLadder_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	AccountLadder_SienaReference_sql   = "SienaReference" // SienaReference is a String
	AccountLadder_BusinessDate_sql   = "BusinessDate" // BusinessDate is a Time
	AccountLadder_ContractNumber_sql   = "ContractNumber" // ContractNumber is a String
	AccountLadder_Balance_sql   = "Balance" // Balance is a Float
	AccountLadder_DealtCcy_sql   = "DealtCcy" // DealtCcy is a String
	AccountLadder_AmountDp_sql   = "AmountDp" // AmountDp is a Int

	/// Definitions End
	///
	/// Application Field Definitions
	///
	AccountLadder_SienaReference_scrn   = "SienaReference" // SienaReference is a String
	AccountLadder_BusinessDate_scrn   = "BusinessDate" // BusinessDate is a Time
	AccountLadder_ContractNumber_scrn   = "ContractNumber" // ContractNumber is a String
	AccountLadder_Balance_scrn   = "Balance" // Balance is a Float
	AccountLadder_DealtCcy_scrn   = "DealtCcy" // DealtCcy is a String
	AccountLadder_AmountDp_scrn   = "AmountDp" // AmountDp is a Int

	/// Definitions End
	///
)

//accountladder_PageList provides the information for the template for a list of AccountLadders
type AccountLadder_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []AccountLadder
	Context	 appContext
}

//accountladder_Page provides the information for the template for an individual AccountLadder
type AccountLadder_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SienaReference         string
	SienaReference_props     FieldProperties
	BusinessDate         string
	BusinessDate_props     FieldProperties
	ContractNumber         string
	ContractNumber_props     FieldProperties
	Balance         string
	Balance_props     FieldProperties
	DealtCcy         string
	DealtCcy_props     FieldProperties
	AmountDp         string
	AmountDp_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}