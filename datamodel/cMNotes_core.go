package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/cmnotes.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CMNotes (cmnotes)
// Endpoint 	        : CMNotes (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:08
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type CMNotes struct {

NoteId        string
StreamId        string
Summary        string
Details        string
RecordState        string
CreatedBy        string
CreatedDateTime        string

}

const (
	CMNotes_Title       = "Customer Note"
	CMNotes_SQLTable    = "contactManagerNote"
	CMNotes_SQLSearchID = "noteId"
	CMNotes_QueryString = "ID"
	///
	/// Handler Defintions
	///
	CMNotes_Template     = "CMNotes"
	CMNotes_TemplateList = "CMNotes_List"
	CMNotes_TemplateView = "CMNotes_View"
	CMNotes_TemplateEdit = "CMNotes_Edit"
	CMNotes_TemplateNew  = "CMNotes_New"
	///
	/// Handler Monitor Paths
	///
	CMNotes_Path       = "/API/CMNotes/"
	CMNotes_PathList   = "/CMNotesList/"
	CMNotes_PathView   = "/CMNotesView/"
	CMNotes_PathEdit   = "/CMNotesEdit/"
	CMNotes_PathNew    = "/CMNotesNew/"
	CMNotes_PathSave   = "/CMNotesSave/"
	CMNotes_PathDelete = "/CMNotesDelete/"
	///
	/// SQL Field Definitions
	///
	CMNotes_NoteId   = "noteId" // NoteId is a Int
	CMNotes_StreamId   = "streamId" // StreamId is a Int
	CMNotes_Summary   = "summary" // Summary is a String
	CMNotes_Details   = "details" // Details is a String
	CMNotes_RecordState   = "recordState" // RecordState is a Int
	CMNotes_CreatedBy   = "createdBy" // CreatedBy is a String
	CMNotes_CreatedDateTime   = "createdDateTime" // CreatedDateTime is a Time

	/// Definitions End
)
