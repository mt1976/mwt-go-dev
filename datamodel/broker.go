package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/broker.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Broker (broker)
// Endpoint 	        : Broker (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 02/12/2021 at 19:39:59
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Broker struct {

Code        string
Name        string
FullName        string
Contact        string
Address        string
LEI        string

}

const (
	Broker_Title       = "Brokers"
	Broker_SQLTable    = "sienaBroker"
	Broker_SQLSearchID = "Code"
	Broker_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Broker_TemplateList = "Broker_List"
	Broker_TemplateView = "Broker_View"
	Broker_TemplateEdit = "Broker_Edit"
	Broker_TemplateNew  = "Broker_New"
	///
	/// Handler Monitor Paths
	///
	Broker_PathList   = "/BrokerList/"
	Broker_PathView   = "/BrokerView/"
	Broker_PathEdit   = "/BrokerEdit/"
	Broker_PathNew    = "/BrokerNew/"
	Broker_PathSave   = "/BrokerSave/"
	Broker_PathDelete = "/BrokerDelete/"
	///
	/// SQL Field Definitions
	///
	Broker_Code   = "Code" // Code is a String
	Broker_Name   = "Name" // Name is a String
	Broker_FullName   = "FullName" // FullName is a String
	Broker_Contact   = "Contact" // Contact is a String
	Broker_Address   = "Address" // Address is a String
	Broker_LEI   = "LEI" // LEI is a String

	/// Definitions End
)
