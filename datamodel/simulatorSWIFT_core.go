package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/simulatorswift.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : SimulatorSWIFT (simulatorswift)
// Endpoint 	        : SimulatorSWIFT (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:14
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//SimulatorSWIFT defines the datamolde for the SimulatorSWIFT object
type SimulatorSWIFT struct {

ID       string
FileName       string
MessageRaw       string
MessageFmt       string
Action       string

}

const (
	SimulatorSWIFT_Title       = "SWIFT Simulator"
	SimulatorSWIFT_SQLTable    = ""
	SimulatorSWIFT_SQLSearchID = "ID"
	SimulatorSWIFT_QueryString = "ID"
	///
	/// Handler Defintions
	///
	SimulatorSWIFT_Template     = "SimulatorSWIFT"
	SimulatorSWIFT_TemplateList = "/SimulatorSWIFT/SimulatorSWIFT_List"
	SimulatorSWIFT_TemplateView = "/SimulatorSWIFT/SimulatorSWIFT_View"
	SimulatorSWIFT_TemplateEdit = "/SimulatorSWIFT/SimulatorSWIFT_Edit"
	SimulatorSWIFT_TemplateNew  = "/SimulatorSWIFT/SimulatorSWIFT_New"
	///
	/// Handler Monitor Paths
	///
	SimulatorSWIFT_Path       = "/API/SimulatorSWIFT/"
	SimulatorSWIFT_PathList   = "/SimulatorSWIFTList/"
	SimulatorSWIFT_PathView   = "/SimulatorSWIFTView/"
	SimulatorSWIFT_PathEdit   = "/SimulatorSWIFTEdit/"
	SimulatorSWIFT_PathNew    = "/SimulatorSWIFTNew/"
	SimulatorSWIFT_PathSave   = "/SimulatorSWIFTSave/"
	SimulatorSWIFT_PathDelete = "/SimulatorSWIFTDelete/"
	///
	///
	/// SQL Field Definitions
	///
SimulatorSWIFT_ID_sql   = "ID" // ID is a String
SimulatorSWIFT_FileName_sql   = "FileName" // FileName is a String
SimulatorSWIFT_MessageRaw_sql   = "MessageRaw" // MessageRaw is a String
SimulatorSWIFT_MessageFmt_sql   = "MessageFmt" // MessageFmt is a String
SimulatorSWIFT_Action_sql   = "Action" // Action is a String

	/// Definitions End

	/// Application Field Definitions
	///
SimulatorSWIFT_ID_scrn   = "ID" // ID is a String
SimulatorSWIFT_FileName_scrn   = "FileName" // FileName is a String
SimulatorSWIFT_MessageRaw_scrn   = "MessageRaw" // MessageRaw is a String
SimulatorSWIFT_MessageFmt_scrn   = "MessageFmt" // MessageFmt is a String
SimulatorSWIFT_Action_scrn   = "Action" // Action is a String

	/// Definitions End
)
