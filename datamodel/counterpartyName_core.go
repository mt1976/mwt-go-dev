package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/counterpartyname.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyName (counterpartyname)
// Endpoint 	        : CounterpartyName (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:10
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type CounterpartyName struct {

NameFirm        string
NameCentre        string
FullName        string
CompID        string

}

const (
	CounterpartyName_Title       = "Counterparty Name"
	CounterpartyName_SQLTable    = "sienaCounterpartyNameLookup"
	CounterpartyName_SQLSearchID = "CompID"
	CounterpartyName_QueryString = "ID"
	///
	/// Handler Defintions
	///
	CounterpartyName_Template     = "CounterpartyName"
	CounterpartyName_TemplateList = "CounterpartyName_List"
	CounterpartyName_TemplateView = "CounterpartyName_View"
	CounterpartyName_TemplateEdit = "CounterpartyName_Edit"
	CounterpartyName_TemplateNew  = "CounterpartyName_New"
	///
	/// Handler Monitor Paths
	///
	CounterpartyName_Path       = "/API/CounterpartyName/"
	CounterpartyName_PathList   = "/CounterpartyNameList/"
	CounterpartyName_PathView   = "/CounterpartyNameView/"
	CounterpartyName_PathEdit   = "/CounterpartyNameEdit/"
	CounterpartyName_PathNew    = "/CounterpartyNameNew/"
	CounterpartyName_PathSave   = "/CounterpartyNameSave/"
	CounterpartyName_PathDelete = "/CounterpartyNameDelete/"
	///
	/// SQL Field Definitions
	///
	CounterpartyName_NameFirm   = "NameFirm" // NameFirm is a String
	CounterpartyName_NameCentre   = "NameCentre" // NameCentre is a String
	CounterpartyName_FullName   = "FullName" // FullName is a String
	CounterpartyName_CompID   = "CompID" // CompID is a String

	/// Definitions End
)
