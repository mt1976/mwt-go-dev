package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/dealauditeventconversations.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DealAuditEventConversations (dealauditeventconversations)
// Endpoint 	        : DealAuditEventConversations (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:50
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//DealAuditEventConversations defines the datamolde for the DealAuditEventConversations object
type DealAuditEventConversations struct {


DealRefNo       string
DealRefNo_props FieldProperties
EventIndex       string
EventIndex_props FieldProperties
MessageIndex       string
MessageIndex_props FieldProperties
Message       string
Message_props FieldProperties
InternalId       string
InternalId_props FieldProperties
InternalDeleted       string
InternalDeleted_props FieldProperties
UpdatedTransactionId       string
UpdatedTransactionId_props FieldProperties
UpdatedUserId       string
UpdatedUserId_props FieldProperties
UpdatedDateTime       string
UpdatedDateTime_props FieldProperties
DeletedTransactionId       string
DeletedTransactionId_props FieldProperties
DeletedUserId       string
DeletedUserId_props FieldProperties
ChangeType       string
ChangeType_props FieldProperties
 // Any lookups will be added below












}

const (
	DealAuditEventConversations_Title       = "Deal Audit Conversation"
	DealAuditEventConversations_SQLTable    = "sienaDealAuditEventConversation"
	DealAuditEventConversations_SQLSearchID = "InternalId"
	DealAuditEventConversations_QueryString = "ID"
	///
	/// Handler Defintions
	///
	DealAuditEventConversations_Template     = "DealAuditEventConversations"
	DealAuditEventConversations_TemplateList = "/DealAuditEventConversations/DealAuditEventConversations_List"
	DealAuditEventConversations_TemplateView = "/DealAuditEventConversations/DealAuditEventConversations_View"
	DealAuditEventConversations_TemplateEdit = "/DealAuditEventConversations/DealAuditEventConversations_Edit"
	DealAuditEventConversations_TemplateNew  = "/DealAuditEventConversations/DealAuditEventConversations_New"
	///
	/// Handler Monitor Paths
	///
	DealAuditEventConversations_Path       = "/API/DealAuditEventConversations/"
	DealAuditEventConversations_PathList   = "/DealAuditEventConversationsList/"
	DealAuditEventConversations_PathView   = "/DealAuditEventConversationsView/"
	DealAuditEventConversations_PathEdit   = "/DealAuditEventConversationsEdit/"
	DealAuditEventConversations_PathNew    = "/DealAuditEventConversationsNew/"
	DealAuditEventConversations_PathSave   = "/DealAuditEventConversationsSave/"
	DealAuditEventConversations_PathDelete = "/DealAuditEventConversationsDelete/"
	///
	//DealAuditEventConversations_Redirect provides a page to return to aftern an action
	DealAuditEventConversations_Redirect = DealAuditEventConversations_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	DealAuditEventConversations_DealRefNo_sql   = "DealRefNo" // DealRefNo is a String
	DealAuditEventConversations_EventIndex_sql   = "EventIndex" // EventIndex is a Int
	DealAuditEventConversations_MessageIndex_sql   = "MessageIndex" // MessageIndex is a Int
	DealAuditEventConversations_Message_sql   = "Message" // Message is a String
	DealAuditEventConversations_InternalId_sql   = "InternalId" // InternalId is a Int
	DealAuditEventConversations_InternalDeleted_sql   = "InternalDeleted" // InternalDeleted is a Time
	DealAuditEventConversations_UpdatedTransactionId_sql   = "UpdatedTransactionId" // UpdatedTransactionId is a String
	DealAuditEventConversations_UpdatedUserId_sql   = "UpdatedUserId" // UpdatedUserId is a String
	DealAuditEventConversations_UpdatedDateTime_sql   = "UpdatedDateTime" // UpdatedDateTime is a Time
	DealAuditEventConversations_DeletedTransactionId_sql   = "DeletedTransactionId" // DeletedTransactionId is a String
	DealAuditEventConversations_DeletedUserId_sql   = "DeletedUserId" // DeletedUserId is a String
	DealAuditEventConversations_ChangeType_sql   = "ChangeType" // ChangeType is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	DealAuditEventConversations_DealRefNo_scrn   = "DealRefNo" // DealRefNo is a String
	DealAuditEventConversations_EventIndex_scrn   = "EventIndex" // EventIndex is a Int
	DealAuditEventConversations_MessageIndex_scrn   = "MessageIndex" // MessageIndex is a Int
	DealAuditEventConversations_Message_scrn   = "Message" // Message is a String
	DealAuditEventConversations_InternalId_scrn   = "InternalId" // InternalId is a Int
	DealAuditEventConversations_InternalDeleted_scrn   = "InternalDeleted" // InternalDeleted is a Time
	DealAuditEventConversations_UpdatedTransactionId_scrn   = "UpdatedTransactionId" // UpdatedTransactionId is a String
	DealAuditEventConversations_UpdatedUserId_scrn   = "UpdatedUserId" // UpdatedUserId is a String
	DealAuditEventConversations_UpdatedDateTime_scrn   = "UpdatedDateTime" // UpdatedDateTime is a Time
	DealAuditEventConversations_DeletedTransactionId_scrn   = "DeletedTransactionId" // DeletedTransactionId is a String
	DealAuditEventConversations_DeletedUserId_scrn   = "DeletedUserId" // DeletedUserId is a String
	DealAuditEventConversations_ChangeType_scrn   = "ChangeType" // ChangeType is a String

	/// Definitions End
	///
)

//dealauditeventconversations_PageList provides the information for the template for a list of DealAuditEventConversationss
type DealAuditEventConversations_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []DealAuditEventConversations
	Context	 appContext
}

//dealauditeventconversations_Page provides the information for the template for an individual DealAuditEventConversations
type DealAuditEventConversations_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	DealRefNo         string
	DealRefNo_props     FieldProperties
	EventIndex         string
	EventIndex_props     FieldProperties
	MessageIndex         string
	MessageIndex_props     FieldProperties
	Message         string
	Message_props     FieldProperties
	InternalId         string
	InternalId_props     FieldProperties
	InternalDeleted         string
	InternalDeleted_props     FieldProperties
	UpdatedTransactionId         string
	UpdatedTransactionId_props     FieldProperties
	UpdatedUserId         string
	UpdatedUserId_props     FieldProperties
	UpdatedDateTime         string
	UpdatedDateTime_props     FieldProperties
	DeletedTransactionId         string
	DeletedTransactionId_props     FieldProperties
	DeletedUserId         string
	DeletedUserId_props     FieldProperties
	ChangeType         string
	ChangeType_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}