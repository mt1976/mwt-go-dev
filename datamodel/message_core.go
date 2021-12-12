package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/message.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Message (message)
// Endpoint 	        : Message (Message)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:17
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Message struct {

SYSId        string
Id        string
Message        string
SYSCreated        string
SYSWho        string
SYSHost        string
SYSUpdated        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdatedBy        string
SYSUpdatedHost        string

}

const (
	Message_Title       = "System Message"
	Message_SQLTable    = "messageStore"
	Message_SQLSearchID = "id"
	Message_QueryString = "Message"
	///
	/// Handler Defintions
	///
	Message_Template     = "Message"
	Message_TemplateList = "Message_List"
	Message_TemplateView = "Message_View"
	Message_TemplateEdit = "Message_Edit"
	Message_TemplateNew  = "Message_New"
	///
	/// Handler Monitor Paths
	///
	Message_Path       = "/API/Message/"
	Message_PathList   = "/MessageList/"
	Message_PathView   = "/MessageView/"
	Message_PathEdit   = "/MessageEdit/"
	Message_PathNew    = "/MessageNew/"
	Message_PathSave   = "/MessageSave/"
	Message_PathDelete = "/MessageDelete/"
	///
	/// SQL Field Definitions
	///
	Message_SYSId   = "_id" // SYSId is a Int
	Message_Id   = "Id" // Id is a String
	Message_Message   = "Message" // Message is a String
	Message_SYSCreated   = "_created" // SYSCreated is a String
	Message_SYSWho   = "_who" // SYSWho is a String
	Message_SYSHost   = "_host" // SYSHost is a String
	Message_SYSUpdated   = "_updated" // SYSUpdated is a String
	Message_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	Message_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	Message_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	Message_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
