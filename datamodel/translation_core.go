package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/translation.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Translation (translation)
// Endpoint 	        : Translation (Message)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:20
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Translation struct {

SYSId        string
Id        string
Class        string
Message        string
Translation        string
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
	Translation_Title       = "Message Translation"
	Translation_SQLTable    = "translationStore"
	Translation_SQLSearchID = "Id"
	Translation_QueryString = "Message"
	///
	/// Handler Defintions
	///
	Translation_Template     = "Translation"
	Translation_TemplateList = "Translation_List"
	Translation_TemplateView = "Translation_View"
	Translation_TemplateEdit = "Translation_Edit"
	Translation_TemplateNew  = "Translation_New"
	///
	/// Handler Monitor Paths
	///
	Translation_Path       = "/API/Translation/"
	Translation_PathList   = "/TranslationList/"
	Translation_PathView   = "/TranslationView/"
	Translation_PathEdit   = "/TranslationEdit/"
	Translation_PathNew    = "/TranslationNew/"
	Translation_PathSave   = "/TranslationSave/"
	Translation_PathDelete = "/TranslationDelete/"
	///
	/// SQL Field Definitions
	///
	Translation_SYSId   = "_id" // SYSId is a Int
	Translation_Id   = "Id" // Id is a String
	Translation_Class   = "Class" // Class is a String
	Translation_Message   = "Message" // Message is a String
	Translation_Translation   = "Translation" // Translation is a String
	Translation_SYSCreated   = "_created" // SYSCreated is a String
	Translation_SYSWho   = "_who" // SYSWho is a String
	Translation_SYSHost   = "_host" // SYSHost is a String
	Translation_SYSUpdated   = "_updated" // SYSUpdated is a String
	Translation_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	Translation_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	Translation_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	Translation_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
