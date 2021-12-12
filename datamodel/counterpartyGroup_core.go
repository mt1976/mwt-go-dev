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
// Date & Time		    : 12/12/2021 at 16:13:09
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type CounterpartyGroup struct {

Name        string
CountryCode        string
SuperGroup        string
Country_Lookup        string
Parent_Lookup        string

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
	CounterpartyGroup_TemplateList = "CounterpartyGroup_List"
	CounterpartyGroup_TemplateView = "CounterpartyGroup_View"
	CounterpartyGroup_TemplateEdit = "CounterpartyGroup_Edit"
	CounterpartyGroup_TemplateNew  = "CounterpartyGroup_New"
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
	/// SQL Field Definitions
	///
	CounterpartyGroup_Name   = "Name" // Name is a String
	CounterpartyGroup_CountryCode   = "CountryCode" // CountryCode is a String
	CounterpartyGroup_SuperGroup   = "SuperGroup" // SuperGroup is a String
	CounterpartyGroup_Country_Lookup   = "Country_Lookup" // Country_Lookup is a String
	CounterpartyGroup_Parent_Lookup   = "Parent_Lookup" // Parent_Lookup is a String

	/// Definitions End
)
