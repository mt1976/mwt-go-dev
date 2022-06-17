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
// Date & Time		    : 17/06/2022 at 18:38:14
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Translation defines the datamolde for the Translation object
type Translation struct {

SYSId       string
Id       string
Class       string
Message       string
Translation       string
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
	Translation_Title       = "Message Translation"
	Translation_SQLTable    = "translationStore"
	Translation_SQLSearchID = "Id"
	Translation_QueryString = "Message"
	///
	/// Handler Defintions
	///
	Translation_Template     = "Translation"
	Translation_TemplateList = "/Translation/Translation_List"
	Translation_TemplateView = "/Translation/Translation_View"
	Translation_TemplateEdit = "/Translation/Translation_Edit"
	Translation_TemplateNew  = "/Translation/Translation_New"
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
	///
	/// SQL Field Definitions
	///
Translation_SYSId_sql   = "_id" // SYSId is a Int
Translation_Id_sql   = "Id" // Id is a String
Translation_Class_sql   = "Class" // Class is a String
Translation_Message_sql   = "Message" // Message is a String
Translation_Translation_sql   = "Translation" // Translation is a String
Translation_SYSCreated_sql   = "_created" // SYSCreated is a String
Translation_SYSWho_sql   = "_who" // SYSWho is a String
Translation_SYSHost_sql   = "_host" // SYSHost is a String
Translation_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
Translation_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
Translation_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
Translation_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String
Translation_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End

	/// Application Field Definitions
	///
Translation_SYSId_scrn   = "SYSId" // SYSId is a Int
Translation_Id_scrn   = "Id" // Id is a String
Translation_Class_scrn   = "Class" // Class is a String
Translation_Message_scrn   = "Message" // Message is a String
Translation_Translation_scrn   = "Translation" // Translation is a String
Translation_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
Translation_SYSWho_scrn   = "SYSWho" // SYSWho is a String
Translation_SYSHost_scrn   = "SYSHost" // SYSHost is a String
Translation_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
Translation_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
Translation_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
Translation_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String
Translation_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
