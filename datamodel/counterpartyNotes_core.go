package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/counterpartynotes.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyNotes (counterpartynotes)
// Endpoint 	        : CounterpartyNotes (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:48
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CounterpartyNotes defines the datamolde for the CounterpartyNotes object
type CounterpartyNotes struct {


NoteId       string
NoteId_props FieldProperties
StreamId       string
StreamId_props FieldProperties
Summary       string
Summary_props FieldProperties
Details       string
Details_props FieldProperties
RecordState       string
RecordState_props FieldProperties
CreatedBy       string
CreatedBy_props FieldProperties
CreatedDateTime       string
CreatedDateTime_props FieldProperties
 // Any lookups will be added below
StreamId_lookup []Lookup_Item


RecordState_lookup []Lookup_Item



}

const (
	CounterpartyNotes_Title       = "Counterparty Notes"
	CounterpartyNotes_SQLTable    = "contactManagerNote"
	CounterpartyNotes_SQLSearchID = "noteId"
	CounterpartyNotes_QueryString = "ID"
	///
	/// Handler Defintions
	///
	CounterpartyNotes_Template     = "CounterpartyNotes"
	CounterpartyNotes_TemplateList = "/CounterpartyNotes/CounterpartyNotes_List"
	CounterpartyNotes_TemplateView = "/CounterpartyNotes/CounterpartyNotes_View"
	CounterpartyNotes_TemplateEdit = "/CounterpartyNotes/CounterpartyNotes_Edit"
	CounterpartyNotes_TemplateNew  = "/CounterpartyNotes/CounterpartyNotes_New"
	///
	/// Handler Monitor Paths
	///
	CounterpartyNotes_Path       = "/API/CounterpartyNotes/"
	CounterpartyNotes_PathList   = "/CounterpartyNotesList/"
	CounterpartyNotes_PathView   = "/CounterpartyNotesView/"
	CounterpartyNotes_PathEdit   = "/CounterpartyNotesEdit/"
	CounterpartyNotes_PathNew    = "/CounterpartyNotesNew/"
	CounterpartyNotes_PathSave   = "/CounterpartyNotesSave/"
	CounterpartyNotes_PathDelete = "/CounterpartyNotesDelete/"
	///
	//CounterpartyNotes_Redirect provides a page to return to aftern an action
	CounterpartyNotes_Redirect = CounterpartyNotes_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	CounterpartyNotes_NoteId_sql   = "noteId" // NoteId is a Int
	CounterpartyNotes_StreamId_sql   = "streamId" // StreamId is a Int
	CounterpartyNotes_Summary_sql   = "summary" // Summary is a String
	CounterpartyNotes_Details_sql   = "details" // Details is a String
	CounterpartyNotes_RecordState_sql   = "recordState" // RecordState is a Int
	CounterpartyNotes_CreatedBy_sql   = "createdBy" // CreatedBy is a String
	CounterpartyNotes_CreatedDateTime_sql   = "createdDateTime" // CreatedDateTime is a Time

	/// Definitions End
	///
	/// Application Field Definitions
	///
	CounterpartyNotes_NoteId_scrn   = "NoteId" // NoteId is a Int
	CounterpartyNotes_StreamId_scrn   = "StreamId" // StreamId is a Int
	CounterpartyNotes_Summary_scrn   = "Summary" // Summary is a String
	CounterpartyNotes_Details_scrn   = "Details" // Details is a String
	CounterpartyNotes_RecordState_scrn   = "RecordState" // RecordState is a Int
	CounterpartyNotes_CreatedBy_scrn   = "CreatedBy" // CreatedBy is a String
	CounterpartyNotes_CreatedDateTime_scrn   = "CreatedDateTime" // CreatedDateTime is a Time

	/// Definitions End
	///
)

//counterpartynotes_PageList provides the information for the template for a list of CounterpartyNotess
type CounterpartyNotes_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []CounterpartyNotes
	Context	 appContext
}

//counterpartynotes_Page provides the information for the template for an individual CounterpartyNotes
type CounterpartyNotes_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	NoteId         string
	NoteId_props     FieldProperties
	StreamId         string
	StreamId_lookup    []Lookup_Item
	StreamId_props     FieldProperties
	Summary         string
	Summary_props     FieldProperties
	Details         string
	Details_props     FieldProperties
	RecordState         string
	RecordState_lookup    []Lookup_Item
	RecordState_props     FieldProperties
	CreatedBy         string
	CreatedBy_props     FieldProperties
	CreatedDateTime         string
	CreatedDateTime_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}