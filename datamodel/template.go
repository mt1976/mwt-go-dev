package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/template.go"
// ----------------------------------------------------------------
// Package            : datamodel
// Object 			  : Template
// Endpoint Root 	  : Template
// Search QueryString : TemplateID
// From   			  : 
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 15/11/2021 at 16:52:52
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Template struct {
	AppInternalID string  // Special field for internal use only
SYSId        string
FIELD1        string
FIELD2        string
SYSCreated        string
SYSWho        string
SYSHost        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdated        string
SYSUpdatedHost        string
SYSUpdatedBy        string
ID        string
SYSVersion        string

}

const (
	Template_Title       = "Template Contents"
	Template_SQLTable    = "template"
	Template_SQLSearchID = "ID"
	Template_QueryString = "TemplateID"
	///
	/// Handler Defintions
	///
	Template_TemplateList = "Template_List"
	Template_TemplateView = "Template_View"
	Template_TemplateEdit = "Template_Edit"
	Template_TemplateNew  = "Template_New"
	///
	/// Handler Monitor Paths
	///
	Template_PathList   = "/TemplateList/"
	Template_PathView   = "/TemplateView/"
	Template_PathEdit   = "/TemplateEdit/"
	Template_PathNew    = "/TemplateNew/"
	Template_PathSave   = "/TemplateSave/"
	Template_PathDelete = "/TemplateDelete/"
	///
	/// SQL Field Definitions
	///
	Template_SYSId   = "_id" // SYSId is a Int
	Template_FIELD1   = "FIELD1" // FIELD1 is a String
	Template_FIELD2   = "FIELD2" // FIELD2 is a String
	Template_SYSCreated   = "_created" // SYSCreated is a String
	Template_SYSWho   = "_who" // SYSWho is a String
	Template_SYSHost   = "_host" // SYSHost is a String
	Template_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	Template_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	Template_SYSUpdated   = "_updated" // SYSUpdated is a String
	Template_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	Template_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	Template_ID   = "ID" // ID is a String
	Template_SYSVersion   = "_version" // SYSVersion is a String

	/// Definitions End
)
