package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/contactstreamstatus.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ContactStreamStatus (contactstreamstatus)
// Endpoint 	        : ContactStreamStatus (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:21
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//ContactStreamStatus defines the datamolde for the ContactStreamStatus object
type ContactStreamStatus struct {


StatusId       string
StatusId_props FieldProperties
Description       string
Description_props FieldProperties
RecordState       string
RecordState_props FieldProperties
 // Any lookups will be added below



}

const (
	ContactStreamStatus_Title       = "Conversation Status"
	ContactStreamStatus_SQLTable    = "contactManagerStreamStatus"
	ContactStreamStatus_SQLSearchID = "statusId"
	ContactStreamStatus_QueryString = "ID"
	///
	/// Handler Defintions
	///
	ContactStreamStatus_Template     = "ContactStreamStatus"
	ContactStreamStatus_TemplateList = "/ContactStreamStatus/ContactStreamStatus_List"
	ContactStreamStatus_TemplateView = "/ContactStreamStatus/ContactStreamStatus_View"
	ContactStreamStatus_TemplateEdit = "/ContactStreamStatus/ContactStreamStatus_Edit"
	ContactStreamStatus_TemplateNew  = "/ContactStreamStatus/ContactStreamStatus_New"
	///
	/// Handler Monitor Paths
	///
	ContactStreamStatus_Path       = "/API/ContactStreamStatus/"
	ContactStreamStatus_PathList   = "/ContactStreamStatusList/"
	ContactStreamStatus_PathView   = "/ContactStreamStatusView/"
	ContactStreamStatus_PathEdit   = "/ContactStreamStatusEdit/"
	ContactStreamStatus_PathNew    = "/ContactStreamStatusNew/"
	ContactStreamStatus_PathSave   = "/ContactStreamStatusSave/"
	ContactStreamStatus_PathDelete = "/ContactStreamStatusDelete/"
	///
	///
	/// SQL Field Definitions
	///
ContactStreamStatus_StatusId_sql   = "statusId" // StatusId is a Int
ContactStreamStatus_Description_sql   = "description" // Description is a String
ContactStreamStatus_RecordState_sql   = "recordState" // RecordState is a Int

	/// Definitions End

	/// Application Field Definitions
	///
ContactStreamStatus_StatusId_scrn   = "StatusId" // StatusId is a Int
ContactStreamStatus_Description_scrn   = "Description" // Description is a String
ContactStreamStatus_RecordState_scrn   = "RecordState" // RecordState is a Int

	/// Definitions End
)
