package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/counterpartygroup.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyGroup (counterpartygroup)
// Endpoint 	        : CounterpartyGroup (Group)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 03/12/2021 at 13:16:56
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type CounterpartyGroup struct {

Name        string
CountryCode        string
SuperGroup        string
Country_Impl        string
Parent_Impl        string

}

const (
	CounterpartyGroup_Title       = "Counterparty Group"
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
