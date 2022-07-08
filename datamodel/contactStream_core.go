package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/contactstream.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ContactStream (contactstream)
// Endpoint 	        : ContactStream (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:46
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//ContactStream defines the datamolde for the ContactStream object
type ContactStream struct {


StreamId       string
StreamId_props FieldProperties
ContactId       string
ContactId_props FieldProperties
Usr       string
Usr_props FieldProperties
Description       string
Description_props FieldProperties
StreamTypeId       string
StreamTypeId_props FieldProperties
StreamStatusId       string
StreamStatusId_props FieldProperties
RecordState       string
RecordState_props FieldProperties
CreatedDateTime       string
CreatedDateTime_props FieldProperties
OpenedDateTime       string
OpenedDateTime_props FieldProperties
ClosedDateTime       string
ClosedDateTime_props FieldProperties
 // Any lookups will be added below



StreamTypeId_lookup []Lookup_Item
StreamStatusId_lookup []Lookup_Item
RecordState_lookup []Lookup_Item




}

const (
	ContactStream_Title       = "Conversations"
	ContactStream_SQLTable    = "contactManagerStream"
	ContactStream_SQLSearchID = "streamId"
	ContactStream_QueryString = "ID"
	///
	/// Handler Defintions
	///
	ContactStream_Template     = "ContactStream"
	ContactStream_TemplateList = "/ContactStream/ContactStream_List"
	ContactStream_TemplateView = "/ContactStream/ContactStream_View"
	ContactStream_TemplateEdit = "/ContactStream/ContactStream_Edit"
	ContactStream_TemplateNew  = "/ContactStream/ContactStream_New"
	///
	/// Handler Monitor Paths
	///
	ContactStream_Path       = "/API/ContactStream/"
	ContactStream_PathList   = "/ContactStreamList/"
	ContactStream_PathView   = "/ContactStreamView/"
	ContactStream_PathEdit   = "/ContactStreamEdit/"
	ContactStream_PathNew    = "/ContactStreamNew/"
	ContactStream_PathSave   = "/ContactStreamSave/"
	ContactStream_PathDelete = "/ContactStreamDelete/"
	///
	//ContactStream_Redirect provides a page to return to aftern an action
	ContactStream_Redirect = ContactStream_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	ContactStream_StreamId_sql   = "streamId" // StreamId is a Int
	ContactStream_ContactId_sql   = "contactId" // ContactId is a Int
	ContactStream_Usr_sql   = "usr" // Usr is a String
	ContactStream_Description_sql   = "description" // Description is a String
	ContactStream_StreamTypeId_sql   = "streamTypeId" // StreamTypeId is a Int
	ContactStream_StreamStatusId_sql   = "streamStatusId" // StreamStatusId is a Int
	ContactStream_RecordState_sql   = "recordState" // RecordState is a Int
	ContactStream_CreatedDateTime_sql   = "createdDateTime" // CreatedDateTime is a Time
	ContactStream_OpenedDateTime_sql   = "openedDateTime" // OpenedDateTime is a Time
	ContactStream_ClosedDateTime_sql   = "closedDateTime" // ClosedDateTime is a Time

	/// Definitions End
	///
	/// Application Field Definitions
	///
	ContactStream_StreamId_scrn   = "StreamId" // StreamId is a Int
	ContactStream_ContactId_scrn   = "ContactId" // ContactId is a Int
	ContactStream_Usr_scrn   = "Usr" // Usr is a String
	ContactStream_Description_scrn   = "Description" // Description is a String
	ContactStream_StreamTypeId_scrn   = "StreamTypeId" // StreamTypeId is a Int
	ContactStream_StreamStatusId_scrn   = "StreamStatusId" // StreamStatusId is a Int
	ContactStream_RecordState_scrn   = "RecordState" // RecordState is a Int
	ContactStream_CreatedDateTime_scrn   = "CreatedDateTime" // CreatedDateTime is a Time
	ContactStream_OpenedDateTime_scrn   = "OpenedDateTime" // OpenedDateTime is a Time
	ContactStream_ClosedDateTime_scrn   = "ClosedDateTime" // ClosedDateTime is a Time

	/// Definitions End
	///
)

//contactstream_PageList provides the information for the template for a list of ContactStreams
type ContactStream_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []ContactStream
	Context	 appContext
}

//contactstream_Page provides the information for the template for an individual ContactStream
type ContactStream_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	StreamId         string
	StreamId_props     FieldProperties
	ContactId         string
	ContactId_props     FieldProperties
	Usr         string
	Usr_props     FieldProperties
	Description         string
	Description_props     FieldProperties
	StreamTypeId         string
	StreamTypeId_lookup    []Lookup_Item
	StreamTypeId_props     FieldProperties
	StreamStatusId         string
	StreamStatusId_lookup    []Lookup_Item
	StreamStatusId_props     FieldProperties
	RecordState         string
	RecordState_lookup    []Lookup_Item
	RecordState_props     FieldProperties
	CreatedDateTime         string
	CreatedDateTime_props     FieldProperties
	OpenedDateTime         string
	OpenedDateTime_props     FieldProperties
	ClosedDateTime         string
	ClosedDateTime_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}