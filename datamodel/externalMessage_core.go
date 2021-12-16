package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/externalmessage.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ExternalMessage (externalmessage)
// Endpoint 	        : ExternalMessage (MessageID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/12/2021 at 09:31:50
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type ExternalMessage struct {

SYSId        string
MessageID        string
MessageFormat        string
MessageDeliveredTo        string
MessageBody        string
MessageFilename        string
MessageLife        string
MessageDate        string
MessageTime        string
MessageTimeoutAction        string
MessageACKNAK        string
ResponseID        string
ResponseFilename        string
ResponseBody        string
ResponseDate        string
ResponseTime        string
ResponseAdditionalInfo        string
SYSCreated        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdated        string
SYSUpdatedBy        string
SYSUpdatedHost        string
MessageTimeout        string
MessageEmitted        string
ResponseRecieved        string
MessageClass        string
AppID        string

}

const (
	ExternalMessage_Title       = "ExternalMessage"
	ExternalMessage_SQLTable    = "externalMessageStore"
	ExternalMessage_SQLSearchID = "MessageID"
	ExternalMessage_QueryString = "MessageID"
	///
	/// Handler Defintions
	///
	ExternalMessage_Template     = "ExternalMessage"
	ExternalMessage_TemplateList = "ExternalMessage_List"
	ExternalMessage_TemplateView = "ExternalMessage_View"
	ExternalMessage_TemplateEdit = "ExternalMessage_Edit"
	ExternalMessage_TemplateNew  = "ExternalMessage_New"
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
	/// SQL Field Definitions
	///
	ExternalMessage_SYSId   = "_id" // SYSId is a Int
	ExternalMessage_MessageID   = "messageID" // MessageID is a String
	ExternalMessage_MessageFormat   = "messageFormat" // MessageFormat is a String
	ExternalMessage_MessageDeliveredTo   = "messageDeliveredTo" // MessageDeliveredTo is a String
	ExternalMessage_MessageBody   = "messageBody" // MessageBody is a String
	ExternalMessage_MessageFilename   = "messageFilename" // MessageFilename is a String
	ExternalMessage_MessageLife   = "messageLife" // MessageLife is a String
	ExternalMessage_MessageDate   = "messageDate" // MessageDate is a String
	ExternalMessage_MessageTime   = "messageTime" // MessageTime is a String
	ExternalMessage_MessageTimeoutAction   = "messageTimeoutAction" // MessageTimeoutAction is a String
	ExternalMessage_MessageACKNAK   = "messageACKNAK" // MessageACKNAK is a String
	ExternalMessage_ResponseID   = "responseID" // ResponseID is a String
	ExternalMessage_ResponseFilename   = "responseFilename" // ResponseFilename is a String
	ExternalMessage_ResponseBody   = "responseBody" // ResponseBody is a String
	ExternalMessage_ResponseDate   = "responseDate" // ResponseDate is a String
	ExternalMessage_ResponseTime   = "responseTime" // ResponseTime is a String
	ExternalMessage_ResponseAdditionalInfo   = "responseAdditionalInfo" // ResponseAdditionalInfo is a String
	ExternalMessage_SYSCreated   = "_created" // SYSCreated is a String
	ExternalMessage_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	ExternalMessage_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	ExternalMessage_SYSUpdated   = "_updated" // SYSUpdated is a String
	ExternalMessage_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	ExternalMessage_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	ExternalMessage_MessageTimeout   = "messageTimeout" // MessageTimeout is a String
	ExternalMessage_MessageEmitted   = "messageEmitted" // MessageEmitted is a String
	ExternalMessage_ResponseRecieved   = "responseRecieved" // ResponseRecieved is a String
	ExternalMessage_MessageClass   = "messageClass" // MessageClass is a String
	ExternalMessage_AppID   = "appID" // AppID is a String

	/// Definitions End
)
