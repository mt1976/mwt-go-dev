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
// Date & Time		    : 17/06/2022 at 18:38:06
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Broker defines the datamolde for the Broker object
type Broker struct {

Code       string
Name       string
FullName       string
Contact       string
Address       string
LEI       string

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

	/// Application Field Definitions
	///
Broker_Code_scrn   = "Code" // Code is a String
Broker_Name_scrn   = "Name" // Name is a String
Broker_FullName_scrn   = "FullName" // FullName is a String
Broker_Contact_scrn   = "Contact" // Contact is a String
Broker_Address_scrn   = "Address" // Address is a String
Broker_LEI_scrn   = "LEI" // LEI is a String

	/// Definitions End
)
