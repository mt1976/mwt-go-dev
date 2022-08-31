package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/dealauditevent.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DealAuditEvent (dealauditevent)
// Endpoint 	        : DealAuditEvent (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:50
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//DealAuditEvent defines the datamolde for the DealAuditEvent object
type DealAuditEvent struct {


DealRefNo       string
DealRefNo_props FieldProperties
EventIndex       string
EventIndex_props FieldProperties
CommonRefNo       string
CommonRefNo_props FieldProperties
Timestamp       string
Timestamp_props FieldProperties
UTCTimestamp       string
UTCTimestamp_props FieldProperties
EventType       string
EventType_props FieldProperties
Status       string
Status_props FieldProperties
LimitOrderStatus       string
LimitOrderStatus_props FieldProperties
Usr       string
Usr_props FieldProperties
DealingInterface       string
DealingInterface_props FieldProperties
SourceIP       string
SourceIP_props FieldProperties
MessageID       string
MessageID_props FieldProperties
Details       string
Details_props FieldProperties
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
	DealAuditEvent_Title       = "Deal Audit History"
	DealAuditEvent_SQLTable    = "sienaDealAuditEvent"
	DealAuditEvent_SQLSearchID = "InternalId"
	DealAuditEvent_QueryString = "ID"
	///
	/// Handler Defintions
	///
	DealAuditEvent_Template     = "DealAuditEvent"
	DealAuditEvent_TemplateList = "/DealAuditEvent/DealAuditEvent_List"
	DealAuditEvent_TemplateView = "/DealAuditEvent/DealAuditEvent_View"
	DealAuditEvent_TemplateEdit = "/DealAuditEvent/DealAuditEvent_Edit"
	DealAuditEvent_TemplateNew  = "/DealAuditEvent/DealAuditEvent_New"
	///
	/// Handler Monitor Paths
	///
	DealAuditEvent_Path       = "/API/DealAuditEvent/"
	DealAuditEvent_PathList   = "/DealAuditEventList/"
	DealAuditEvent_PathView   = "/DealAuditEventView/"
	DealAuditEvent_PathEdit   = "/DealAuditEventEdit/"
	DealAuditEvent_PathNew    = "/DealAuditEventNew/"
	DealAuditEvent_PathSave   = "/DealAuditEventSave/"
	DealAuditEvent_PathDelete = "/DealAuditEventDelete/"
	///
	//DealAuditEvent_Redirect provides a page to return to aftern an action
	DealAuditEvent_Redirect = DealAuditEvent_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	DealAuditEvent_DealRefNo_sql   = "DealRefNo" // DealRefNo is a String
	DealAuditEvent_EventIndex_sql   = "EventIndex" // EventIndex is a Int
	DealAuditEvent_CommonRefNo_sql   = "CommonRefNo" // CommonRefNo is a String
	DealAuditEvent_Timestamp_sql   = "Timestamp" // Timestamp is a Time
	DealAuditEvent_UTCTimestamp_sql   = "UTCTimestamp" // UTCTimestamp is a String
	DealAuditEvent_EventType_sql   = "EventType" // EventType is a String
	DealAuditEvent_Status_sql   = "Status" // Status is a String
	DealAuditEvent_LimitOrderStatus_sql   = "LimitOrderStatus" // LimitOrderStatus is a String
	DealAuditEvent_Usr_sql   = "Usr" // Usr is a String
	DealAuditEvent_DealingInterface_sql   = "DealingInterface" // DealingInterface is a String
	DealAuditEvent_SourceIP_sql   = "SourceIP" // SourceIP is a String
	DealAuditEvent_MessageID_sql   = "MessageID" // MessageID is a String
	DealAuditEvent_Details_sql   = "Details" // Details is a String
	DealAuditEvent_InternalId_sql   = "InternalId" // InternalId is a Int
	DealAuditEvent_InternalDeleted_sql   = "InternalDeleted" // InternalDeleted is a Time
	DealAuditEvent_UpdatedTransactionId_sql   = "UpdatedTransactionId" // UpdatedTransactionId is a String
	DealAuditEvent_UpdatedUserId_sql   = "UpdatedUserId" // UpdatedUserId is a String
	DealAuditEvent_UpdatedDateTime_sql   = "UpdatedDateTime" // UpdatedDateTime is a Time
	DealAuditEvent_DeletedTransactionId_sql   = "DeletedTransactionId" // DeletedTransactionId is a String
	DealAuditEvent_DeletedUserId_sql   = "DeletedUserId" // DeletedUserId is a String
	DealAuditEvent_ChangeType_sql   = "ChangeType" // ChangeType is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	DealAuditEvent_DealRefNo_scrn   = "DealRefNo" // DealRefNo is a String
	DealAuditEvent_EventIndex_scrn   = "EventIndex" // EventIndex is a Int
	DealAuditEvent_CommonRefNo_scrn   = "CommonRefNo" // CommonRefNo is a String
	DealAuditEvent_Timestamp_scrn   = "Timestamp" // Timestamp is a Time
	DealAuditEvent_UTCTimestamp_scrn   = "UTCTimestamp" // UTCTimestamp is a String
	DealAuditEvent_EventType_scrn   = "EventType" // EventType is a String
	DealAuditEvent_Status_scrn   = "Status" // Status is a String
	DealAuditEvent_LimitOrderStatus_scrn   = "LimitOrderStatus" // LimitOrderStatus is a String
	DealAuditEvent_Usr_scrn   = "Usr" // Usr is a String
	DealAuditEvent_DealingInterface_scrn   = "DealingInterface" // DealingInterface is a String
	DealAuditEvent_SourceIP_scrn   = "SourceIP" // SourceIP is a String
	DealAuditEvent_MessageID_scrn   = "MessageID" // MessageID is a String
	DealAuditEvent_Details_scrn   = "Details" // Details is a String
	DealAuditEvent_InternalId_scrn   = "InternalId" // InternalId is a Int
	DealAuditEvent_InternalDeleted_scrn   = "InternalDeleted" // InternalDeleted is a Time
	DealAuditEvent_UpdatedTransactionId_scrn   = "UpdatedTransactionId" // UpdatedTransactionId is a String
	DealAuditEvent_UpdatedUserId_scrn   = "UpdatedUserId" // UpdatedUserId is a String
	DealAuditEvent_UpdatedDateTime_scrn   = "UpdatedDateTime" // UpdatedDateTime is a Time
	DealAuditEvent_DeletedTransactionId_scrn   = "DeletedTransactionId" // DeletedTransactionId is a String
	DealAuditEvent_DeletedUserId_scrn   = "DeletedUserId" // DeletedUserId is a String
	DealAuditEvent_ChangeType_scrn   = "ChangeType" // ChangeType is a String

	/// Definitions End
	///
)

//dealauditevent_PageList provides the information for the template for a list of DealAuditEvents
type DealAuditEvent_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []DealAuditEvent
	Context	 appContext
}

//dealauditevent_Page provides the information for the template for an individual DealAuditEvent
type DealAuditEvent_Page struct {
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
	CommonRefNo         string
	CommonRefNo_props     FieldProperties
	Timestamp         string
	Timestamp_props     FieldProperties
	UTCTimestamp         string
	UTCTimestamp_props     FieldProperties
	EventType         string
	EventType_props     FieldProperties
	Status         string
	Status_props     FieldProperties
	LimitOrderStatus         string
	LimitOrderStatus_props     FieldProperties
	Usr         string
	Usr_props     FieldProperties
	DealingInterface         string
	DealingInterface_props     FieldProperties
	SourceIP         string
	SourceIP_props     FieldProperties
	MessageID         string
	MessageID_props     FieldProperties
	Details         string
	Details_props     FieldProperties
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