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
// Date & Time		    : 17/06/2022 at 18:38:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Credentials defines the datamolde for the Credentials object
type Credentials struct {

SYSId       string
Id       string
Username       string
Password       string
Firstname       string
Lastname       string
Knownas       string
Email       string
Issued       string
Expiry       string
RoleType       string
Brand       string
SYSCreated       string
SYSWho       string
SYSHost       string
SYSUpdated       string
SYSCreatedBy       string
SYSCreatedHost       string
SYSUpdatedBy       string
SYSUpdatedHost       string

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
	Credentials_TemplateList = "/Credentials/Credentials_List"
	Credentials_TemplateView = "/Credentials/Credentials_View"
	Credentials_TemplateEdit = "/Credentials/Credentials_Edit"
	Credentials_TemplateNew  = "/Credentials/Credentials_New"
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
	///
	/// SQL Field Definitions
	///
Credentials_SYSId_sql   = "_id" // SYSId is a Int
Credentials_Id_sql   = "Id" // Id is a String
Credentials_Username_sql   = "Username" // Username is a String
Credentials_Password_sql   = "Password" // Password is a String
Credentials_Firstname_sql   = "Firstname" // Firstname is a String
Credentials_Lastname_sql   = "Lastname" // Lastname is a String
Credentials_Knownas_sql   = "Knownas" // Knownas is a String
Credentials_Email_sql   = "Email" // Email is a String
Credentials_Issued_sql   = "Issued" // Issued is a String
Credentials_Expiry_sql   = "Expiry" // Expiry is a String
Credentials_RoleType_sql   = "RoleType" // RoleType is a String
Credentials_Brand_sql   = "Brand" // Brand is a String
Credentials_SYSCreated_sql   = "_created" // SYSCreated is a String
Credentials_SYSWho_sql   = "_who" // SYSWho is a String
Credentials_SYSHost_sql   = "_host" // SYSHost is a String
Credentials_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
Credentials_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
Credentials_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
Credentials_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
Credentials_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End

	/// Application Field Definitions
	///
Credentials_SYSId_scrn   = "SYSId" // SYSId is a Int
Credentials_Id_scrn   = "Id" // Id is a String
Credentials_Username_scrn   = "Username" // Username is a String
Credentials_Password_scrn   = "Password" // Password is a String
Credentials_Firstname_scrn   = "Firstname" // Firstname is a String
Credentials_Lastname_scrn   = "Lastname" // Lastname is a String
Credentials_Knownas_scrn   = "Knownas" // Knownas is a String
Credentials_Email_scrn   = "Email" // Email is a String
Credentials_Issued_scrn   = "Issued" // Issued is a String
Credentials_Expiry_scrn   = "Expiry" // Expiry is a String
Credentials_RoleType_scrn   = "RoleType" // RoleType is a String
Credentials_Brand_scrn   = "Brand" // Brand is a String
Credentials_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
Credentials_SYSWho_scrn   = "SYSWho" // SYSWho is a String
Credentials_SYSHost_scrn   = "SYSHost" // SYSHost is a String
Credentials_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
Credentials_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
Credentials_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
Credentials_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
Credentials_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
