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
// Date & Time		    : 14/06/2022 at 21:32:03
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//DealAuditEvent defines the datamolde for the DealAuditEvent object
type DealAuditEvent struct {

DealRefNo       string
EventIndex       string
CommonRefNo       string
Timestamp       string
UTCTimestamp       string
EventType       string
Status       string
LimitOrderStatus       string
Usr       string
DealingInterface       string
SourceIP       string
MessageID       string
Details       string
InternalId       string
InternalDeleted       string
UpdatedTransactionId       string
UpdatedUserId       string
UpdatedDateTime       string
DeletedTransactionId       string
DeletedUserId       string
ChangeType       string

}

const (
	DealAuditEvent_Title       = "Deal Audit History"
	DealAuditEvent_SQLTable    = "DealAuditEvent"
	DealAuditEvent_SQLSearchID = "InternalId"
	DealAuditEvent_QueryString = "ID"
	///
	/// Handler Defintions
	///
	DealAuditEvent_Template     = "DealAuditEvent"
	DealAuditEvent_TemplateList = "DealAuditEvent_List"
	DealAuditEvent_TemplateView = "DealAuditEvent_View"
	DealAuditEvent_TemplateEdit = "DealAuditEvent_Edit"
	DealAuditEvent_TemplateNew  = "DealAuditEvent_New"
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
	/// SQL Field Definitions
	///
DealAuditEvent_DealRefNo   = "DealRefNo" // DealRefNo is a String
DealAuditEvent_EventIndex   = "EventIndex" // EventIndex is a Int
DealAuditEvent_CommonRefNo   = "CommonRefNo" // CommonRefNo is a String
DealAuditEvent_Timestamp   = "Timestamp" // Timestamp is a Time
DealAuditEvent_UTCTimestamp   = "UTCTimestamp" // UTCTimestamp is a String
DealAuditEvent_EventType   = "EventType" // EventType is a String
DealAuditEvent_Status   = "Status" // Status is a String
DealAuditEvent_LimitOrderStatus   = "LimitOrderStatus" // LimitOrderStatus is a String
DealAuditEvent_Usr   = "Usr" // Usr is a String
DealAuditEvent_DealingInterface   = "DealingInterface" // DealingInterface is a String
DealAuditEvent_SourceIP   = "SourceIP" // SourceIP is a String
DealAuditEvent_MessageID   = "MessageID" // MessageID is a String
DealAuditEvent_Details   = "Details" // Details is a String
DealAuditEvent_InternalId   = "InternalId" // InternalId is a Int
DealAuditEvent_InternalDeleted   = "InternalDeleted" // InternalDeleted is a Time
DealAuditEvent_UpdatedTransactionId   = "UpdatedTransactionId" // UpdatedTransactionId is a String
DealAuditEvent_UpdatedUserId   = "UpdatedUserId" // UpdatedUserId is a String
DealAuditEvent_UpdatedDateTime   = "UpdatedDateTime" // UpdatedDateTime is a Time
DealAuditEvent_DeletedTransactionId   = "DeletedTransactionId" // DeletedTransactionId is a String
DealAuditEvent_DeletedUserId   = "DeletedUserId" // DeletedUserId is a String
DealAuditEvent_ChangeType   = "ChangeType" // ChangeType is a String

	/// Definitions End
)
