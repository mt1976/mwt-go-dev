package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/counterpartygroup.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyGroup (counterpartygroup)
// Endpoint 	        : CounterpartyGroup (Group)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:48
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CounterpartyGroup defines the datamolde for the CounterpartyGroup object
type CounterpartyGroup struct {


Name       string
Name_props FieldProperties
CountryCode       string
CountryCode_props FieldProperties
SuperGroup       string
SuperGroup_props FieldProperties
 // Any lookups will be added below



}

const (
	CounterpartyGroup_Title       = "Counterparty Group"
	CounterpartyGroup_SQLTable    = "sienaCounterpartyGroup"
	CounterpartyGroup_SQLSearchID = "Name"
	CounterpartyGroup_QueryString = "Group"
	///
	/// Handler Defintions
	///
	CounterpartyGroup_Template     = "CounterpartyGroup"
	CounterpartyGroup_TemplateList = "/CounterpartyGroup/CounterpartyGroup_List"
	CounterpartyGroup_TemplateView = "/CounterpartyGroup/CounterpartyGroup_View"
	CounterpartyGroup_TemplateEdit = "/CounterpartyGroup/CounterpartyGroup_Edit"
	CounterpartyGroup_TemplateNew  = "/CounterpartyGroup/CounterpartyGroup_New"
	///
	/// Handler Monitor Paths
	///
	CounterpartyGroup_Path       = "/API/CounterpartyGroup/"
	CounterpartyGroup_PathList   = "/CounterpartyGroupList/"
	CounterpartyGroup_PathView   = "/CounterpartyGroupView/"
	CounterpartyGroup_PathEdit   = "/CounterpartyGroupEdit/"
	CounterpartyGroup_PathNew    = "/CounterpartyGroupNew/"
	CounterpartyGroup_PathSave   = "/CounterpartyGroupSave/"
	CounterpartyGroup_PathDelete = "/CounterpartyGroupDelete/"
	///
	//CounterpartyGroup_Redirect provides a page to return to aftern an action
	CounterpartyGroup_Redirect = CounterpartyGroup_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	CounterpartyGroup_Name_sql   = "Name" // Name is a String
	CounterpartyGroup_CountryCode_sql   = "CountryCode" // CountryCode is a String
	CounterpartyGroup_SuperGroup_sql   = "SuperGroup" // SuperGroup is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	CounterpartyGroup_Name_scrn   = "Name" // Name is a String
	CounterpartyGroup_CountryCode_scrn   = "CountryCode" // CountryCode is a String
	CounterpartyGroup_SuperGroup_scrn   = "SuperGroup" // SuperGroup is a String

	/// Definitions End
	///
)

//counterpartygroup_PageList provides the information for the template for a list of CounterpartyGroups
type CounterpartyGroup_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []CounterpartyGroup
	Context	 appContext
}

//counterpartygroup_Page provides the information for the template for an individual CounterpartyGroup
type CounterpartyGroup_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	Name         string
	Name_props     FieldProperties
	CountryCode         string
	CountryCode_props     FieldProperties
	SuperGroup         string
	SuperGroup_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}