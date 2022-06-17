package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/externalmessage.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ExternalMessage (externalmessage)
// Endpoint 	        : ExternalMessage (Message)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:11
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//ExternalMessage defines the datamolde for the ExternalMessage object
type ExternalMessage struct {

SYSId       string
MessageID       string
MessageFormat       string
MessageDeliveredTo       string
MessageBody       string
MessageFilename       string
MessageLife       string
MessageDate       string
MessageTime       string
MessageTimeoutAction       string
MessageACKNAK       string
MessageACKNAK_lookup []Lookup_Item
ResponseID       string
ResponseFilename       string
ResponseBody       string
ResponseDate       string
ResponseTime       string
ResponseAdditionalInfo       string
SYSCreated       string
SYSCreatedBy       string
SYSCreatedHost       string
SYSUpdated       string
SYSUpdatedBy       string
SYSUpdatedHost       string
MessageTimeout       string
MessageEmitted       string
ResponseRecieved       string
MessageClass       string
AppID       string

}

const (
	ExternalMessage_Title       = "ExternalMessage"
	ExternalMessage_SQLTable    = "externalMessageStore"
	ExternalMessage_SQLSearchID = "MessageID"
	ExternalMessage_QueryString = "Message"
	///
	/// Handler Defintions
	///
	ExternalMessage_Template     = "ExternalMessage"
	ExternalMessage_TemplateList = "/ExternalMessage/ExternalMessage_List"
	ExternalMessage_TemplateView = "/ExternalMessage/ExternalMessage_View"
	ExternalMessage_TemplateEdit = "/ExternalMessage/ExternalMessage_Edit"
	ExternalMessage_TemplateNew  = "/ExternalMessage/ExternalMessage_New"
	///
	/// Handler Monitor Paths
	///
	ExternalMessage_Path       = "/API/ExternalMessage/"
	ExternalMessage_PathList   = "/ExternalMessageList/"
	ExternalMessage_PathView   = "/ExternalMessageView/"
	ExternalMessage_PathEdit   = "/ExternalMessageEdit/"
	ExternalMessage_PathNew    = "/ExternalMessageNew/"
	ExternalMessage_PathSave   = "/ExternalMessageSave/"
	ExternalMessage_PathDelete = "/ExternalMessageDelete/"
	///
	///
	/// SQL Field Definitions
	///
ExternalMessage_SYSId_sql   = "_id" // SYSId is a Int
ExternalMessage_MessageID_sql   = "messageID" // MessageID is a String
ExternalMessage_MessageFormat_sql   = "messageFormat" // MessageFormat is a String
ExternalMessage_MessageDeliveredTo_sql   = "messageDeliveredTo" // MessageDeliveredTo is a String
ExternalMessage_MessageBody_sql   = "messageBody" // MessageBody is a String
ExternalMessage_MessageFilename_sql   = "messageFilename" // MessageFilename is a String
ExternalMessage_MessageLife_sql   = "messageLife" // MessageLife is a String
ExternalMessage_MessageDate_sql   = "messageDate" // MessageDate is a String
ExternalMessage_MessageTime_sql   = "messageTime" // MessageTime is a String
ExternalMessage_MessageTimeoutAction_sql   = "messageTimeoutAction" // MessageTimeoutAction is a String
ExternalMessage_MessageACKNAK_sql   = "messageACKNAK" // MessageACKNAK is a String
ExternalMessage_ResponseID_sql   = "responseID" // ResponseID is a String
ExternalMessage_ResponseFilename_sql   = "responseFilename" // ResponseFilename is a String
ExternalMessage_ResponseBody_sql   = "responseBody" // ResponseBody is a String
ExternalMessage_ResponseDate_sql   = "responseDate" // ResponseDate is a String
ExternalMessage_ResponseTime_sql   = "responseTime" // ResponseTime is a String
ExternalMessage_ResponseAdditionalInfo_sql   = "responseAdditionalInfo" // ResponseAdditionalInfo is a String
ExternalMessage_SYSCreated_sql   = "_created" // SYSCreated is a String
ExternalMessage_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
ExternalMessage_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
ExternalMessage_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
ExternalMessage_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
ExternalMessage_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
ExternalMessage_MessageTimeout_sql   = "messageTimeout" // MessageTimeout is a String
ExternalMessage_MessageEmitted_sql   = "messageEmitted" // MessageEmitted is a String
ExternalMessage_ResponseRecieved_sql   = "responseRecieved" // ResponseRecieved is a String
ExternalMessage_MessageClass_sql   = "messageClass" // MessageClass is a String
ExternalMessage_AppID_sql   = "appID" // AppID is a String

	/// Definitions End

	/// Application Field Definitions
	///
ExternalMessage_SYSId_scrn   = "SYSId" // SYSId is a Int
ExternalMessage_MessageID_scrn   = "MessageID" // MessageID is a String
ExternalMessage_MessageFormat_scrn   = "MessageFormat" // MessageFormat is a String
ExternalMessage_MessageDeliveredTo_scrn   = "MessageDeliveredTo" // MessageDeliveredTo is a String
ExternalMessage_MessageBody_scrn   = "MessageBody" // MessageBody is a String
ExternalMessage_MessageFilename_scrn   = "MessageFilename" // MessageFilename is a String
ExternalMessage_MessageLife_scrn   = "MessageLife" // MessageLife is a String
ExternalMessage_MessageDate_scrn   = "MessageDate" // MessageDate is a String
ExternalMessage_MessageTime_scrn   = "MessageTime" // MessageTime is a String
ExternalMessage_MessageTimeoutAction_scrn   = "MessageTimeoutAction" // MessageTimeoutAction is a String
ExternalMessage_MessageACKNAK_scrn   = "MessageACKNAK" // MessageACKNAK is a String
ExternalMessage_ResponseID_scrn   = "ResponseID" // ResponseID is a String
ExternalMessage_ResponseFilename_scrn   = "ResponseFilename" // ResponseFilename is a String
ExternalMessage_ResponseBody_scrn   = "ResponseBody" // ResponseBody is a String
ExternalMessage_ResponseDate_scrn   = "ResponseDate" // ResponseDate is a String
ExternalMessage_ResponseTime_scrn   = "ResponseTime" // ResponseTime is a String
ExternalMessage_ResponseAdditionalInfo_scrn   = "ResponseAdditionalInfo" // ResponseAdditionalInfo is a String
ExternalMessage_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
ExternalMessage_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
ExternalMessage_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
ExternalMessage_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
ExternalMessage_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
ExternalMessage_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
ExternalMessage_MessageTimeout_scrn   = "MessageTimeout" // MessageTimeout is a String
ExternalMessage_MessageEmitted_scrn   = "MessageEmitted" // MessageEmitted is a String
ExternalMessage_ResponseRecieved_scrn   = "ResponseRecieved" // ResponseRecieved is a String
ExternalMessage_MessageClass_scrn   = "MessageClass" // MessageClass is a String
ExternalMessage_AppID_scrn   = "AppID" // AppID is a String

	/// Definitions End
)
