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
// Date & Time		    : 26/06/2022 at 18:48:23
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
	///
	/// SQL Field Definitions
	///
CounterpartyGroup_Name_sql   = "Name" // Name is a String
CounterpartyGroup_CountryCode_sql   = "CountryCode" // CountryCode is a String
CounterpartyGroup_SuperGroup_sql   = "SuperGroup" // SuperGroup is a String

	/// Definitions End

	/// Application Field Definitions
	///
CounterpartyGroup_Name_scrn   = "Name" // Name is a String
CounterpartyGroup_CountryCode_scrn   = "CountryCode" // CountryCode is a String
CounterpartyGroup_SuperGroup_scrn   = "SuperGroup" // SuperGroup is a String

	/// Definitions End
)
