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
// Date & Time		    : 05/06/2022 at 12:56:56
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type SimulatorSWIFT struct {

ID        string
FileName        string
MessageRaw        string
MessageFmt        string
Action        string


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
	SimulatorSWIFT_TemplateList = "SimulatorSWIFT_List"
	SimulatorSWIFT_TemplateView = "SimulatorSWIFT_View"
	SimulatorSWIFT_TemplateEdit = "SimulatorSWIFT_Edit"
	SimulatorSWIFT_TemplateNew  = "SimulatorSWIFT_New"
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
	/// SQL Field Definitions
	///
	SimulatorSWIFT_ID   = "ID" // ID is a String
	SimulatorSWIFT_FileName   = "FileName" // FileName is a String
	SimulatorSWIFT_MessageRaw   = "MessageRaw" // MessageRaw is a String
	SimulatorSWIFT_MessageFmt   = "MessageFmt" // MessageFmt is a String
	SimulatorSWIFT_Action   = "Action" // Action is a String


	/// Definitions End
)
