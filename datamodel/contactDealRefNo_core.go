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
// Date & Time		    : 28/06/2022 at 16:10:45
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
	//ContactDealRefNo_Redirect provides a page to return to aftern an action
	ContactDealRefNo_Redirect = ContactDealRefNo_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	ContactDealRefNo_RefNo_sql   = "refNo" // RefNo is a String
	ContactDealRefNo_NoteId_sql   = "noteId" // NoteId is a Int
	ContactDealRefNo_RecordState_sql   = "recordState" // RecordState is a Int

	/// Definitions End
	///
	/// Application Field Definitions
	///
	ContactDealRefNo_RefNo_scrn   = "RefNo" // RefNo is a String
	ContactDealRefNo_NoteId_scrn   = "NoteId" // NoteId is a Int
	ContactDealRefNo_RecordState_scrn   = "RecordState" // RecordState is a Int

	/// Definitions End
	///
)

//contactdealrefno_PageList provides the information for the template for a list of ContactDealRefNos
type ContactDealRefNo_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []ContactDealRefNo
	Context	 appContext
}

//contactdealrefno_Page provides the information for the template for an individual ContactDealRefNo
type ContactDealRefNo_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	RefNo         string
	RefNo_lookup    []Lookup_Item
	RefNo_props     FieldProperties
	NoteId         string
	NoteId_lookup    []Lookup_Item
	NoteId_props     FieldProperties
	RecordState         string
	RecordState_lookup    []Lookup_Item
	RecordState_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}