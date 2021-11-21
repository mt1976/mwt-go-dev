package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/counterpartygroup.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyGroup (counterpartygroup)
// Endpoint 	        : CounterpartyGroup (Group)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 15:44:01
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type CounterpartyGroup struct {
	AppInternalID string  // Special field for internal use only
Name        string
CountryCode        string
SuperGroup        string
Country_Impl        string
Parent_Impl        string

}

const (
	CounterpartyGroup_Title       = "Counterparty Groups"
	CounterpartyGroup_SQLTable    = "sienaCounterpartyGroup"
	CounterpartyGroup_SQLSearchID = "Name"
	CounterpartyGroup_QueryString = "Group"
	///
	/// Handler Defintions
	///
	CounterpartyGroup_TemplateList = "CounterpartyGroup_List"
	CounterpartyGroup_TemplateView = "CounterpartyGroup_View"
	CounterpartyGroup_TemplateEdit = "CounterpartyGroup_Edit"
	CounterpartyGroup_TemplateNew  = "CounterpartyGroup_New"
	///
	/// Handler Monitor Paths
	///
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
	CounterpartyGroup_Country_Impl   = "Country_Impl" // Country_Impl is a String
	CounterpartyGroup_Parent_Impl   = "Parent_Impl" // Parent_Impl is a String

	/// Definitions End
)
