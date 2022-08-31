package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/counterpartyname.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyName (counterpartyname)
// Endpoint 	        : CounterpartyName (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:48
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CounterpartyName defines the datamolde for the CounterpartyName object
type CounterpartyName struct {


NameFirm       string
NameFirm_props FieldProperties
NameCentre       string
NameCentre_props FieldProperties
FullName       string
FullName_props FieldProperties
CompID       string
CompID_props FieldProperties
 // Any lookups will be added below




}

const (
	CounterpartyName_Title       = "Counterparty Name"
	CounterpartyName_SQLTable    = "sienaCounterpartyNameLookup"
	CounterpartyName_SQLSearchID = "CompID"
	CounterpartyName_QueryString = "ID"
	///
	/// Handler Defintions
	///
	CounterpartyName_Template     = "CounterpartyName"
	CounterpartyName_TemplateList = "/CounterpartyName/CounterpartyName_List"
	CounterpartyName_TemplateView = "/CounterpartyName/CounterpartyName_View"
	CounterpartyName_TemplateEdit = "/CounterpartyName/CounterpartyName_Edit"
	CounterpartyName_TemplateNew  = "/CounterpartyName/CounterpartyName_New"
	///
	/// Handler Monitor Paths
	///
	CounterpartyName_Path       = "/API/CounterpartyName/"
	CounterpartyName_PathList   = "/CounterpartyNameList/"
	CounterpartyName_PathView   = "/CounterpartyNameView/"
	CounterpartyName_PathEdit   = "/CounterpartyNameEdit/"
	CounterpartyName_PathNew    = "/CounterpartyNameNew/"
	CounterpartyName_PathSave   = "/CounterpartyNameSave/"
	CounterpartyName_PathDelete = "/CounterpartyNameDelete/"
	///
	//CounterpartyName_Redirect provides a page to return to aftern an action
	CounterpartyName_Redirect = CounterpartyName_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	CounterpartyName_NameFirm_sql   = "NameFirm" // NameFirm is a String
	CounterpartyName_NameCentre_sql   = "NameCentre" // NameCentre is a String
	CounterpartyName_FullName_sql   = "FullName" // FullName is a String
	CounterpartyName_CompID_sql   = "CompID" // CompID is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	CounterpartyName_NameFirm_scrn   = "NameFirm" // NameFirm is a String
	CounterpartyName_NameCentre_scrn   = "NameCentre" // NameCentre is a String
	CounterpartyName_FullName_scrn   = "FullName" // FullName is a String
	CounterpartyName_CompID_scrn   = "CompID" // CompID is a String

	/// Definitions End
	///
)

//counterpartyname_PageList provides the information for the template for a list of CounterpartyNames
type CounterpartyName_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []CounterpartyName
	Context	 appContext
}

//counterpartyname_Page provides the information for the template for an individual CounterpartyName
type CounterpartyName_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	NameFirm         string
	NameFirm_props     FieldProperties
	NameCentre         string
	NameCentre_props     FieldProperties
	FullName         string
	FullName_props     FieldProperties
	CompID         string
	CompID_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}