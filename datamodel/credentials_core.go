package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/credentials.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Credentials (credentials)
// Endpoint 	        : Credentials (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:11
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Credentials struct {

SYSId        string
Id        string
Username        string
Password        string
Firstname        string
Lastname        string
Knownas        string
Email        string
Issued        string
Expiry        string
RoleType        string
Brand        string
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
	Credentials_Title       = "Credential"
	Credentials_SQLTable    = "credentialsStore"
	Credentials_SQLSearchID = "id"
	Credentials_QueryString = "Id"
	///
	/// Handler Defintions
	///
	Credentials_Template     = "Credentials"
	Credentials_TemplateList = "Credentials_List"
	Credentials_TemplateView = "Credentials_View"
	Credentials_TemplateEdit = "Credentials_Edit"
	Credentials_TemplateNew  = "Credentials_New"
	///
	/// Handler Monitor Paths
	///
	Credentials_Path       = "/API/Credentials/"
	Credentials_PathList   = "/CredentialsList/"
	Credentials_PathView   = "/CredentialsView/"
	Credentials_PathEdit   = "/CredentialsEdit/"
	Credentials_PathNew    = "/CredentialsNew/"
	Credentials_PathSave   = "/CredentialsSave/"
	Credentials_PathDelete = "/CredentialsDelete/"
	///
	/// SQL Field Definitions
	///
	Credentials_SYSId   = "_id" // SYSId is a Int
	Credentials_Id   = "Id" // Id is a String
	Credentials_Username   = "Username" // Username is a String
	Credentials_Password   = "Password" // Password is a String
	Credentials_Firstname   = "Firstname" // Firstname is a String
	Credentials_Lastname   = "Lastname" // Lastname is a String
	Credentials_Knownas   = "Knownas" // Knownas is a String
	Credentials_Email   = "Email" // Email is a String
	Credentials_Issued   = "Issued" // Issued is a String
	Credentials_Expiry   = "Expiry" // Expiry is a String
	Credentials_RoleType   = "RoleType" // RoleType is a String
	Credentials_Brand   = "Brand" // Brand is a String
	Credentials_SYSCreated   = "_created" // SYSCreated is a String
	Credentials_SYSWho   = "_who" // SYSWho is a String
	Credentials_SYSHost   = "_host" // SYSHost is a String
	Credentials_SYSUpdated   = "_updated" // SYSUpdated is a String
	Credentials_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	Credentials_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	Credentials_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	Credentials_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
