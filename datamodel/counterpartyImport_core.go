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
// Date & Time		    : 12/12/2021 at 16:13:10
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type CounterpartyImport struct {

KeyImportID        string
Firm        string
Centre        string
FirmName        string
CentreName        string
KeyOriginID        string
FullName        string
CompID        string

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
	CounterpartyImport_TemplateList = "CounterpartyImport_List"
	CounterpartyImport_TemplateView = "CounterpartyImport_View"
	CounterpartyImport_TemplateEdit = "CounterpartyImport_Edit"
	CounterpartyImport_TemplateNew  = "CounterpartyImport_New"
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
	/// SQL Field Definitions
	///
	CounterpartyImport_KeyImportID   = "KeyImportID" // KeyImportID is a String
	CounterpartyImport_Firm   = "Firm" // Firm is a String
	CounterpartyImport_Centre   = "Centre" // Centre is a String
	CounterpartyImport_FirmName   = "FirmName" // FirmName is a String
	CounterpartyImport_CentreName   = "CentreName" // CentreName is a String
	CounterpartyImport_KeyOriginID   = "KeyOriginID" // KeyOriginID is a String
	CounterpartyImport_FullName   = "FullName" // FullName is a String
	CounterpartyImport_CompID   = "CompID" // CompID is a String

	/// Definitions End
)
