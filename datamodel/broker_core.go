package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/broker.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Broker (broker)
// Endpoint 	        : Broker (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:43
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Broker defines the datamolde for the Broker object
type Broker struct {


Code       string
Code_props FieldProperties
Name       string
Name_props FieldProperties
FullName       string
FullName_props FieldProperties
Contact       string
Contact_props FieldProperties
Address       string
Address_props FieldProperties
LEI       string
LEI_props FieldProperties
 // Any lookups will be added below






}

const (
	Broker_Title       = "Broker"
	Broker_SQLTable    = "sienaBroker"
	Broker_SQLSearchID = "Code"
	Broker_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Broker_Template     = "Broker"
	Broker_TemplateList = "/Broker/Broker_List"
	Broker_TemplateView = "/Broker/Broker_View"
	Broker_TemplateEdit = "/Broker/Broker_Edit"
	Broker_TemplateNew  = "/Broker/Broker_New"
	///
	/// Handler Monitor Paths
	///
	Broker_Path       = "/API/Broker/"
	Broker_PathList   = "/BrokerList/"
	Broker_PathView   = "/BrokerView/"
	Broker_PathEdit   = "/BrokerEdit/"
	Broker_PathNew    = "/BrokerNew/"
	Broker_PathSave   = "/BrokerSave/"
	Broker_PathDelete = "/BrokerDelete/"
	///
	//Broker_Redirect provides a page to return to aftern an action
	Broker_Redirect = Broker_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	Broker_Code_sql   = "Code" // Code is a String
	Broker_Name_sql   = "Name" // Name is a String
	Broker_FullName_sql   = "FullName" // FullName is a String
	Broker_Contact_sql   = "Contact" // Contact is a String
	Broker_Address_sql   = "Address" // Address is a String
	Broker_LEI_sql   = "LEI" // LEI is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Broker_Code_scrn   = "Code" // Code is a String
	Broker_Name_scrn   = "Name" // Name is a String
	Broker_FullName_scrn   = "FullName" // FullName is a String
	Broker_Contact_scrn   = "Contact" // Contact is a String
	Broker_Address_scrn   = "Address" // Address is a String
	Broker_LEI_scrn   = "LEI" // LEI is a String

	/// Definitions End
	///
)

//broker_PageList provides the information for the template for a list of Brokers
type Broker_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Broker
	Context	 appContext
}

//broker_Page provides the information for the template for an individual Broker
type Broker_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	Code         string
	Code_props     FieldProperties
	Name         string
	Name_props     FieldProperties
	FullName         string
	FullName_props     FieldProperties
	Contact         string
	Contact_props     FieldProperties
	Address         string
	Address_props     FieldProperties
	LEI         string
	LEI_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}