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
// Date & Time		    : 17/06/2022 at 18:38:06
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Catalog defines the datamolde for the Catalog object
type Catalog struct {

ID       string
Endpoint       string
Descr       string
Query       string
Source       string

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
	Catalog_TemplateList = "/Catalog/Catalog_List"
	Catalog_TemplateView = "/Catalog/Catalog_View"
	Catalog_TemplateEdit = "/Catalog/Catalog_Edit"
	Catalog_TemplateNew  = "/Catalog/Catalog_New"
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
	///
	/// SQL Field Definitions
	///
Catalog_ID_sql   = "ID" // ID is a String
Catalog_Endpoint_sql   = "Endpoint" // Endpoint is a String
Catalog_Descr_sql   = "Descr" // Descr is a String
Catalog_Query_sql   = "Query" // Query is a String
Catalog_Source_sql   = "Source" // Source is a String

	/// Definitions End

	/// Application Field Definitions
	///
Catalog_ID_scrn   = "ID" // ID is a String
Catalog_Endpoint_scrn   = "Endpoint" // Endpoint is a String
Catalog_Descr_scrn   = "Descr" // Descr is a String
Catalog_Query_scrn   = "Query" // Query is a String
Catalog_Source_scrn   = "Source" // Source is a String

	/// Definitions End
)
