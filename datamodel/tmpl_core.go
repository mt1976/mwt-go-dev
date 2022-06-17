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
// Date & Time		    : 17/06/2022 at 18:38:14
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
	Tmpl_TemplateList = "/Tmpl/Tmpl_List"
	Tmpl_TemplateView = "/Tmpl/Tmpl_View"
	Tmpl_TemplateEdit = "/Tmpl/Tmpl_Edit"
	Tmpl_TemplateNew  = "/Tmpl/Tmpl_New"
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
	///
	/// SQL Field Definitions
	///
Tmpl_SYSId_sql   = "_id" // SYSId is a Int
Tmpl_FIELD1_sql   = "FIELD1" // FIELD1 is a String
Tmpl_FIELD2_sql   = "FIELD2" // FIELD2 is a String
Tmpl_SYSCreated_sql   = "_created" // SYSCreated is a String
Tmpl_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
Tmpl_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
Tmpl_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
Tmpl_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
Tmpl_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
Tmpl_ID_sql   = "ID" // ID is a String
Tmpl_ExtraField_sql   = "ExtraField" // ExtraField is a String
Tmpl_ExtraField2_sql   = "ExtraField2" // ExtraField2 is a String
Tmpl_ExtraField3_sql   = "ExtraField3" // ExtraField3 is a String

	/// Definitions End

	/// Application Field Definitions
	///
Tmpl_SYSId_scrn   = "SYSId" // SYSId is a Int
Tmpl_FIELD1_scrn   = "FIELD1" // FIELD1 is a String
Tmpl_FIELD2_scrn   = "FIELD2" // FIELD2 is a String
Tmpl_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
Tmpl_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
Tmpl_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
Tmpl_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
Tmpl_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
Tmpl_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
Tmpl_ID_scrn   = "ID" // ID is a String
Tmpl_ExtraField_scrn   = "ExtraField" // ExtraField is a String
Tmpl_ExtraField2_scrn   = "ExtraField2" // ExtraField2 is a String
Tmpl_ExtraField3_scrn   = "ExtraField3" // ExtraField3 is a String

	/// Definitions End
)
