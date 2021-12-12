package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/session.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Session (session)
// Endpoint 	        : Session (SessionID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:19
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Session struct {

SYSId        string
Apptoken        string
Createdate        string
Createtime        string
Uniqueid        string
Sessiontoken        string
Username        string
Password        string
Userip        string
Userhost        string
Appip        string
Apphost        string
Issued        string
Expiry        string
Expiryraw        string
Brand        string
SYSCreated        string
SYSWho        string
SYSHost        string
SYSUpdated        string
Id        string
Expires        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdatedBy        string
SYSUpdatedHost        string
SessionRole        string

}

const (
	Session_Title       = "Session"
	Session_SQLTable    = "sessionStore"
	Session_SQLSearchID = "Id"
	Session_QueryString = "SessionID"
	///
	/// Handler Defintions
	///
	Session_Template     = "Session"
	Session_TemplateList = "Session_List"
	Session_TemplateView = "Session_View"
	Session_TemplateEdit = "Session_Edit"
	Session_TemplateNew  = "Session_New"
	///
	/// Handler Monitor Paths
	///
	Session_Path       = "/API/Session/"
	Session_PathList   = "/SessionList/"
	Session_PathView   = "/SessionView/"
	Session_PathEdit   = "/SessionEdit/"
	Session_PathNew    = "/SessionNew/"
	Session_PathSave   = "/SessionSave/"
	Session_PathDelete = "/SessionDelete/"
	///
	/// SQL Field Definitions
	///
	Session_SYSId   = "_id" // SYSId is a Int
	Session_Apptoken   = "Apptoken" // Apptoken is a String
	Session_Createdate   = "Createdate" // Createdate is a String
	Session_Createtime   = "Createtime" // Createtime is a String
	Session_Uniqueid   = "Uniqueid" // Uniqueid is a String
	Session_Sessiontoken   = "Sessiontoken" // Sessiontoken is a String
	Session_Username   = "Username" // Username is a String
	Session_Password   = "Password" // Password is a String
	Session_Userip   = "Userip" // Userip is a String
	Session_Userhost   = "Userhost" // Userhost is a String
	Session_Appip   = "Appip" // Appip is a String
	Session_Apphost   = "Apphost" // Apphost is a String
	Session_Issued   = "Issued" // Issued is a String
	Session_Expiry   = "Expiry" // Expiry is a String
	Session_Expiryraw   = "Expiryraw" // Expiryraw is a String
	Session_Brand   = "Brand" // Brand is a String
	Session_SYSCreated   = "_created" // SYSCreated is a String
	Session_SYSWho   = "_who" // SYSWho is a String
	Session_SYSHost   = "_host" // SYSHost is a String
	Session_SYSUpdated   = "_updated" // SYSUpdated is a String
	Session_Id   = "Id" // Id is a String
	Session_Expires   = "Expires" // Expires is a Time
	Session_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	Session_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	Session_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	Session_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	Session_SessionRole   = "SessionRole" // SessionRole is a String

	/// Definitions End
)
