package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/credentialsaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CredentialsAction (credentialsaction)
// Endpoint 	        : CredentialsAction (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:25
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CredentialsAction defines the datamolde for the CredentialsAction object
type CredentialsAction struct {


ID       string
ID_props FieldProperties
User       string
User_props FieldProperties
Action       string
Action_props FieldProperties
Notes       string
Notes_props FieldProperties
 // Any lookups will be added below
User_lookup []Lookup_Item
Action_lookup []Lookup_Item


}

const (
	CredentialsAction_Title       = "Credentials Update"
	CredentialsAction_SQLTable    = "inboxMessages"
	CredentialsAction_SQLSearchID = "MailId"
	CredentialsAction_QueryString = "ID"
	///
	/// Handler Defintions
	///
	CredentialsAction_Template     = "CredentialsAction"
	CredentialsAction_TemplateList = "/CredentialsAction/CredentialsAction_List"
	CredentialsAction_TemplateView = "/CredentialsAction/CredentialsAction_View"
	CredentialsAction_TemplateEdit = "/CredentialsAction/CredentialsAction_Edit"
	CredentialsAction_TemplateNew  = "/CredentialsAction/CredentialsAction_New"
	///
	/// Handler Monitor Paths
	///
	CredentialsAction_Path       = "/API/CredentialsAction/"
	CredentialsAction_PathList   = "/CredentialsActionList/"
	CredentialsAction_PathView   = "/CredentialsActionView/"
	CredentialsAction_PathEdit   = "/CredentialsActionEdit/"
	CredentialsAction_PathNew    = "/CredentialsActionNew/"
	CredentialsAction_PathSave   = "/CredentialsActionSave/"
	CredentialsAction_PathDelete = "/CredentialsActionDelete/"
	///
	///
	/// SQL Field Definitions
	///
CredentialsAction_ID_sql   = "ID" // ID is a String
CredentialsAction_User_sql   = "User" // User is a String
CredentialsAction_Action_sql   = "Action" // Action is a String
CredentialsAction_Notes_sql   = "Notes" // Notes is a String

	/// Definitions End

	/// Application Field Definitions
	///
CredentialsAction_ID_scrn   = "ID" // ID is a String
CredentialsAction_User_scrn   = "User" // User is a String
CredentialsAction_Action_scrn   = "Action" // Action is a String
CredentialsAction_Notes_scrn   = "Notes" // Notes is a String

	/// Definitions End
)
