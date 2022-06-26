package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/contactdealrefno.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ContactDealRefNo (contactdealrefno)
// Endpoint 	        : ContactDealRefNo (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:20
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//ContactDealRefNo defines the datamolde for the ContactDealRefNo object
type ContactDealRefNo struct {


RefNo       string
RefNo_props FieldProperties
NoteId       string
NoteId_props FieldProperties
RecordState       string
RecordState_props FieldProperties
 // Any lookups will be added belowRefNo_lookup []Lookup_Item
NoteId_lookup []Lookup_Item
RecordState_lookup []Lookup_Item

}

const (
	ContactDealRefNo_Title       = "Conversation Deals"
	ContactDealRefNo_SQLTable    = "contactManagerDealRefNo"
	ContactDealRefNo_SQLSearchID = "noteId"
	ContactDealRefNo_QueryString = "ID"
	///
	/// Handler Defintions
	///
	ContactDealRefNo_Template     = "ContactDealRefNo"
	ContactDealRefNo_TemplateList = "/ContactDealRefNo/ContactDealRefNo_List"
	ContactDealRefNo_TemplateView = "/ContactDealRefNo/ContactDealRefNo_View"
	ContactDealRefNo_TemplateEdit = "/ContactDealRefNo/ContactDealRefNo_Edit"
	ContactDealRefNo_TemplateNew  = "/ContactDealRefNo/ContactDealRefNo_New"
	///
	/// Handler Monitor Paths
	///
	ContactDealRefNo_Path       = "/API/ContactDealRefNo/"
	ContactDealRefNo_PathList   = "/ContactDealRefNoList/"
	ContactDealRefNo_PathView   = "/ContactDealRefNoView/"
	ContactDealRefNo_PathEdit   = "/ContactDealRefNoEdit/"
	ContactDealRefNo_PathNew    = "/ContactDealRefNoNew/"
	ContactDealRefNo_PathSave   = "/ContactDealRefNoSave/"
	ContactDealRefNo_PathDelete = "/ContactDealRefNoDelete/"
	///
	///
	/// SQL Field Definitions
	///
ContactDealRefNo_RefNo_sql   = "refNo" // RefNo is a String
ContactDealRefNo_NoteId_sql   = "noteId" // NoteId is a Int
ContactDealRefNo_RecordState_sql   = "recordState" // RecordState is a Int

	/// Definitions End

	/// Application Field Definitions
	///
ContactDealRefNo_RefNo_scrn   = "RefNo" // RefNo is a String
ContactDealRefNo_NoteId_scrn   = "NoteId" // NoteId is a Int
ContactDealRefNo_RecordState_scrn   = "RecordState" // RecordState is a Int

	/// Definitions End
)
