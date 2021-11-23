package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/cache.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Cache (cache)
// Endpoint 	        : Cache (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 22/11/2021 at 21:11:39
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Cache struct {
	AppInternalID string  // Special field for internal use only
SYSId        string
Id        string
Object        string
Field        string
Value        string
Expiry        string
SYSCreated        string
SYSWho        string
SYSHost        string
SYSUpdated        string
Source        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdatedBy        string
SYSUpdatedHost        string

}

const (
	Cache_Title       = "Cache Contents"
	Cache_SQLTable    = "cacheStore"
	Cache_SQLSearchID = "Id"
	Cache_QueryString = "ID"
	///
	/// Handler Defintions
	///
	Cache_TemplateList = "Cache_List"
	Cache_TemplateView = "Cache_View"
	Cache_TemplateEdit = "Cache_Edit"
	Cache_TemplateNew  = "Cache_New"
	///
	/// Handler Monitor Paths
	///
	Cache_PathList   = "/CacheList/"
	Cache_PathView   = "/CacheView/"
	Cache_PathEdit   = "/CacheEdit/"
	Cache_PathNew    = "/CacheNew/"
	Cache_PathSave   = "/CacheSave/"
	Cache_PathDelete = "/CacheDelete/"
	///
	/// SQL Field Definitions
	///
	Cache_SYSId   = "_id" // SYSId is a Int
	Cache_Id   = "id" // Id is a String
	Cache_Object   = "object" // Object is a String
	Cache_Field   = "field" // Field is a String
	Cache_Value   = "value" // Value is a String
	Cache_Expiry   = "expiry" // Expiry is a String
	Cache_SYSCreated   = "_created" // SYSCreated is a String
	Cache_SYSWho   = "_who" // SYSWho is a String
	Cache_SYSHost   = "_host" // SYSHost is a String
	Cache_SYSUpdated   = "_updated" // SYSUpdated is a String
	Cache_Source   = "source" // Source is a String
	Cache_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	Cache_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	Cache_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	Cache_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
