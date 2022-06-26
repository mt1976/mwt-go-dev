package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/contactstreamtype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ContactStreamType (contactstreamtype)
// Endpoint 	        : ContactStreamType (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:21
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//ContactStreamType defines the datamolde for the ContactStreamType object
type ContactStreamType struct {


TypeId       string
TypeId_props FieldProperties
Description       string
Description_props FieldProperties
RecordState       string
RecordState_props FieldProperties
 // Any lookups will be added below

RecordState_lookup []Lookup_Item

}

const (
	ContactStreamType_Title       = "Conversations Types"
	ContactStreamType_SQLTable    = "contactManagerStreamType"
	ContactStreamType_SQLSearchID = "typeId"
	ContactStreamType_QueryString = "ID"
	///
	/// Handler Defintions
	///
	ContactStreamType_Template     = "ContactStreamType"
	ContactStreamType_TemplateList = "/ContactStreamType/ContactStreamType_List"
	ContactStreamType_TemplateView = "/ContactStreamType/ContactStreamType_View"
	ContactStreamType_TemplateEdit = "/ContactStreamType/ContactStreamType_Edit"
	ContactStreamType_TemplateNew  = "/ContactStreamType/ContactStreamType_New"
	///
	/// Handler Monitor Paths
	///
	ContactStreamType_Path       = "/API/ContactStreamType/"
	ContactStreamType_PathList   = "/ContactStreamTypeList/"
	ContactStreamType_PathView   = "/ContactStreamTypeView/"
	ContactStreamType_PathEdit   = "/ContactStreamTypeEdit/"
	ContactStreamType_PathNew    = "/ContactStreamTypeNew/"
	ContactStreamType_PathSave   = "/ContactStreamTypeSave/"
	ContactStreamType_PathDelete = "/ContactStreamTypeDelete/"
	///
	///
	/// SQL Field Definitions
	///
ContactStreamType_TypeId_sql   = "typeId" // TypeId is a Int
ContactStreamType_Description_sql   = "description" // Description is a String
ContactStreamType_RecordState_sql   = "recordState" // RecordState is a Int

	/// Definitions End

	/// Application Field Definitions
	///
ContactStreamType_TypeId_scrn   = "TypeId" // TypeId is a Int
ContactStreamType_Description_scrn   = "Description" // Description is a String
ContactStreamType_RecordState_scrn   = "RecordState" // RecordState is a Int

	/// Definitions End
)
