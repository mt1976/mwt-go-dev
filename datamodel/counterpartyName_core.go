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
// Date & Time		    : 17/06/2022 at 18:38:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CounterpartyName defines the datamolde for the CounterpartyName object
type CounterpartyName struct {

NameFirm       string
NameCentre       string
FullName       string
CompID       string

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
	CounterpartyName_TemplateList = "/CounterpartyName/CounterpartyName_List"
	CounterpartyName_TemplateView = "/CounterpartyName/CounterpartyName_View"
	CounterpartyName_TemplateEdit = "/CounterpartyName/CounterpartyName_Edit"
	CounterpartyName_TemplateNew  = "/CounterpartyName/CounterpartyName_New"
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
	///
	/// SQL Field Definitions
	///
CounterpartyName_NameFirm_sql   = "NameFirm" // NameFirm is a String
CounterpartyName_NameCentre_sql   = "NameCentre" // NameCentre is a String
CounterpartyName_FullName_sql   = "FullName" // FullName is a String
CounterpartyName_CompID_sql   = "CompID" // CompID is a String

	/// Definitions End

	/// Application Field Definitions
	///
CounterpartyName_NameFirm_scrn   = "NameFirm" // NameFirm is a String
CounterpartyName_NameCentre_scrn   = "NameCentre" // NameCentre is a String
CounterpartyName_FullName_scrn   = "FullName" // FullName is a String
CounterpartyName_CompID_scrn   = "CompID" // CompID is a String

	/// Definitions End
)
