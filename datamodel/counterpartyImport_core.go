package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/counterpartyimport.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyImport (counterpartyimport)
// Endpoint 	        : CounterpartyImport (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:07
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CounterpartyImport defines the datamolde for the CounterpartyImport object
type CounterpartyImport struct {

KeyImportID       string
Firm       string
Centre       string
FirmName       string
CentreName       string
KeyOriginID       string
FullName       string
CompID       string

}

const (
	CounterpartyImport_Title       = "Counterparty Import"
	CounterpartyImport_SQLTable    = "sienaCounterpartyImportID"
	CounterpartyImport_SQLSearchID = "CompID"
	CounterpartyImport_QueryString = "ID"
	///
	/// Handler Defintions
	///
	CounterpartyImport_Template     = "CounterpartyImport"
	CounterpartyImport_TemplateList = "/CounterpartyImport/CounterpartyImport_List"
	CounterpartyImport_TemplateView = "/CounterpartyImport/CounterpartyImport_View"
	CounterpartyImport_TemplateEdit = "/CounterpartyImport/CounterpartyImport_Edit"
	CounterpartyImport_TemplateNew  = "/CounterpartyImport/CounterpartyImport_New"
	///
	/// Handler Monitor Paths
	///
	CounterpartyImport_Path       = "/API/CounterpartyImport/"
	CounterpartyImport_PathList   = "/CounterpartyImportList/"
	CounterpartyImport_PathView   = "/CounterpartyImportView/"
	CounterpartyImport_PathEdit   = "/CounterpartyImportEdit/"
	CounterpartyImport_PathNew    = "/CounterpartyImportNew/"
	CounterpartyImport_PathSave   = "/CounterpartyImportSave/"
	CounterpartyImport_PathDelete = "/CounterpartyImportDelete/"
	///
	///
	/// SQL Field Definitions
	///
CounterpartyImport_KeyImportID_sql   = "KeyImportID" // KeyImportID is a String
CounterpartyImport_Firm_sql   = "Firm" // Firm is a String
CounterpartyImport_Centre_sql   = "Centre" // Centre is a String
CounterpartyImport_FirmName_sql   = "FirmName" // FirmName is a String
CounterpartyImport_CentreName_sql   = "CentreName" // CentreName is a String
CounterpartyImport_KeyOriginID_sql   = "KeyOriginID" // KeyOriginID is a String
CounterpartyImport_FullName_sql   = "FullName" // FullName is a String
CounterpartyImport_CompID_sql   = "CompID" // CompID is a String

	/// Definitions End

	/// Application Field Definitions
	///
CounterpartyImport_KeyImportID_scrn   = "KeyImportID" // KeyImportID is a String
CounterpartyImport_Firm_scrn   = "Firm" // Firm is a String
CounterpartyImport_Centre_scrn   = "Centre" // Centre is a String
CounterpartyImport_FirmName_scrn   = "FirmName" // FirmName is a String
CounterpartyImport_CentreName_scrn   = "CentreName" // CentreName is a String
CounterpartyImport_KeyOriginID_scrn   = "KeyOriginID" // KeyOriginID is a String
CounterpartyImport_FullName_scrn   = "FullName" // FullName is a String
CounterpartyImport_CompID_scrn   = "CompID" // CompID is a String

	/// Definitions End
)
