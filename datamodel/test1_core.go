package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/test1.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Test1 (test1)
// Endpoint 	        : Test1 (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:32:10
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Test1 defines the datamolde for the Test1 object
type Test1 struct {

ID       string
Endpoint       string
Descr       string
Query       string
Source       string
Firm       string
Firm_lookup []Lookup_Item
YN       string
YN_lookup []Lookup_Item
User       string
User_lookup []Lookup_Item
Cheese       string
Onion       string

}

const (
	Test1_Title       = "Test1Template"
	Test1_SQLTable    = ""
	Test1_SQLSearchID = "ID"
	Test1_QueryString = "ID"
	///
	/// Handler Defintions
	///
	Test1_Template     = "Test1"
	Test1_TemplateList = "Test1_List"
	Test1_TemplateView = "Test1_View"
	Test1_TemplateEdit = "Test1_Edit"
	Test1_TemplateNew  = "Test1_New"
	///
	/// Handler Monitor Paths
	///
	Test1_Path       = "/API/Test1/"
	Test1_PathList   = "/Test1List/"
	Test1_PathView   = "/Test1View/"
	Test1_PathEdit   = "/Test1Edit/"
	Test1_PathNew    = "/Test1New/"
	Test1_PathSave   = "/Test1Save/"
	Test1_PathDelete = "/Test1Delete/"
	///
	/// SQL Field Definitions
	///
Test1_ID   = "ID" // ID is a String
Test1_Endpoint   = "Endpoint" // Endpoint is a String
Test1_Descr   = "Descr" // Descr is a String
Test1_Query   = "Query" // Query is a String
Test1_Source   = "Source" // Source is a String
Test1_Firm   = "Firm" // Firm is a String
Test1_YN   = "YN" // YN is a String
Test1_User   = "User" // User is a String
Test1_Cheese   = "Cheese" // Cheese is a String
Test1_Onion   = "Onion" // Onion is a String

	/// Definitions End
)
