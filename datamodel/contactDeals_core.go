package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/contactdeals.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ContactDeals (contactdeals)
// Endpoint 	        : ContactDeals (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 24/06/2022 at 16:01:55
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//ContactDeals defines the datamolde for the ContactDeals object
type ContactDeals struct {

RefNo       string
NoteId       string
RecordState       string

}

const (
	ContactDeals_Title       = "Conversation Deals"
	ContactDeals_SQLTable    = "contactManagerDealRefNo"
	ContactDeals_SQLSearchID = "noteId"
	ContactDeals_QueryString = "ID"
	///
	/// Handler Defintions
	///
	ContactDeals_Template     = "ContactDeals"
	ContactDeals_TemplateList = "/ContactDeals/ContactDeals_List"
	ContactDeals_TemplateView = "/ContactDeals/ContactDeals_View"
	ContactDeals_TemplateEdit = "/ContactDeals/ContactDeals_Edit"
	ContactDeals_TemplateNew  = "/ContactDeals/ContactDeals_New"
	///
	/// Handler Monitor Paths
	///
	ContactDeals_Path       = "/API/ContactDeals/"
	ContactDeals_PathList   = "/ContactDealsList/"
	ContactDeals_PathView   = "/ContactDealsView/"
	ContactDeals_PathEdit   = "/ContactDealsEdit/"
	ContactDeals_PathNew    = "/ContactDealsNew/"
	ContactDeals_PathSave   = "/ContactDealsSave/"
	ContactDeals_PathDelete = "/ContactDealsDelete/"
	///
	///
	/// SQL Field Definitions
	///
ContactDeals_RefNo_sql   = "refNo" // RefNo is a String
ContactDeals_NoteId_sql   = "noteId" // NoteId is a Int
ContactDeals_RecordState_sql   = "recordState" // RecordState is a Int

	/// Definitions End

	/// Application Field Definitions
	///
ContactDeals_RefNo_scrn   = "RefNo" // RefNo is a String
ContactDeals_NoteId_scrn   = "NoteId" // NoteId is a Int
ContactDeals_RecordState_scrn   = "RecordState" // RecordState is a Int

	/// Definitions End
)
