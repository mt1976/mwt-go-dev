package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/catalog.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Catalog (catalog)
// Endpoint 	        : Catalog (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:16:04
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Catalog struct {

ID        string
Endpoint        string
Descr        string
Query        string
Source        string

}

const (
	Catalog_Title       = "API Catalog"
	Catalog_SQLTable    = ""
	Catalog_SQLSearchID = "ID"
	Catalog_QueryString = "ID"
	///
	/// Handler Defintions
	///
	Catalog_Template     = "Catalog"
	Catalog_TemplateList = "Catalog_List"
	Catalog_TemplateView = "Catalog_View"
	Catalog_TemplateEdit = "Catalog_Edit"
	Catalog_TemplateNew  = "Catalog_New"
	///
	/// Handler Monitor Paths
	///
	Catalog_Path       = "/API/Catalog/"
	Catalog_PathList   = "/CatalogList/"
	Catalog_PathView   = "/CatalogView/"
	Catalog_PathEdit   = "/CatalogEdit/"
	Catalog_PathNew    = "/CatalogNew/"
	Catalog_PathSave   = "/CatalogSave/"
	Catalog_PathDelete = "/CatalogDelete/"
	///
	/// SQL Field Definitions
	///
	Catalog_ID   = "ID" // ID is a String
	Catalog_Endpoint   = "Endpoint" // Endpoint is a String
	Catalog_Descr   = "Descr" // Descr is a String
	Catalog_Query   = "Query" // Query is a String
	Catalog_Source   = "Source" // Source is a String

	/// Definitions End
)
