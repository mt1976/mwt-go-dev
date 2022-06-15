package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/tmpl.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Tmpl (tmpl)
// Endpoint 	        : Tmpl (TemplateID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:32:11
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Tmpl defines the datamolde for the Tmpl object
type Tmpl struct {

SYSId       string
FIELD1       string
FIELD1_lookup []Lookup_Item
FIELD2       string
FIELD2_lookup []Lookup_Item
SYSCreated       string
SYSCreatedBy       string
SYSCreatedHost       string
SYSUpdated       string
SYSUpdatedHost       string
SYSUpdatedBy       string
ID       string
ExtraField       string
ExtraField2       string
ExtraField3       string

}

const (
	Tmpl_Title       = "Template Contents"
	Tmpl_SQLTable    = "template"
	Tmpl_SQLSearchID = "ID"
	Tmpl_QueryString = "TemplateID"
	///
	/// Handler Defintions
	///
	Tmpl_Template     = "Tmpl"
	Tmpl_TemplateList = "Tmpl_List"
	Tmpl_TemplateView = "Tmpl_View"
	Tmpl_TemplateEdit = "Tmpl_Edit"
	Tmpl_TemplateNew  = "Tmpl_New"
	///
	/// Handler Monitor Paths
	///
	Tmpl_Path       = "/API/Tmpl/"
	Tmpl_PathList   = "/TmplList/"
	Tmpl_PathView   = "/TmplView/"
	Tmpl_PathEdit   = "/TmplEdit/"
	Tmpl_PathNew    = "/TmplNew/"
	Tmpl_PathSave   = "/TmplSave/"
	Tmpl_PathDelete = "/TmplDelete/"
	///
	/// SQL Field Definitions
	///
Tmpl_SYSId   = "_id" // SYSId is a Int
Tmpl_FIELD1   = "FIELD1" // FIELD1 is a String
Tmpl_FIELD2   = "FIELD2" // FIELD2 is a String
Tmpl_SYSCreated   = "_created" // SYSCreated is a String
Tmpl_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
Tmpl_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
Tmpl_SYSUpdated   = "_updated" // SYSUpdated is a String
Tmpl_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
Tmpl_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
Tmpl_ID   = "ID" // ID is a String
Tmpl_ExtraField   = "ExtraField" // ExtraField is a String
Tmpl_ExtraField2   = "ExtraField2" // ExtraField2 is a String
Tmpl_ExtraField3   = "ExtraField3" // ExtraField3 is a String

	/// Definitions End
)
