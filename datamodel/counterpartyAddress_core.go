package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/counterpartyaddress.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyAddress (counterpartyaddress)
// Endpoint 	        : CounterpartyAddress (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:47
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CounterpartyAddress defines the datamolde for the CounterpartyAddress object
type CounterpartyAddress struct {


NameFirm       string
NameFirm_props FieldProperties
NameCentre       string
NameCentre_props FieldProperties
Address1       string
Address1_props FieldProperties
Address2       string
Address2_props FieldProperties
CityTown       string
CityTown_props FieldProperties
County       string
County_props FieldProperties
Postcode       string
Postcode_props FieldProperties
CompID       string
CompID_props FieldProperties
 // Any lookups will be added below








}

const (
	CounterpartyAddress_Title       = "Counterparty Address"
	CounterpartyAddress_SQLTable    = "sienaCounterpartyAddress"
	CounterpartyAddress_SQLSearchID = "CompID"
	CounterpartyAddress_QueryString = "ID"
	///
	/// Handler Defintions
	///
	CounterpartyAddress_Template     = "CounterpartyAddress"
	CounterpartyAddress_TemplateList = "/CounterpartyAddress/CounterpartyAddress_List"
	CounterpartyAddress_TemplateView = "/CounterpartyAddress/CounterpartyAddress_View"
	CounterpartyAddress_TemplateEdit = "/CounterpartyAddress/CounterpartyAddress_Edit"
	CounterpartyAddress_TemplateNew  = "/CounterpartyAddress/CounterpartyAddress_New"
	///
	/// Handler Monitor Paths
	///
	CounterpartyAddress_Path       = "/API/CounterpartyAddress/"
	CounterpartyAddress_PathList   = "/CounterpartyAddressList/"
	CounterpartyAddress_PathView   = "/CounterpartyAddressView/"
	CounterpartyAddress_PathEdit   = "/CounterpartyAddressEdit/"
	CounterpartyAddress_PathNew    = "/CounterpartyAddressNew/"
	CounterpartyAddress_PathSave   = "/CounterpartyAddressSave/"
	CounterpartyAddress_PathDelete = "/CounterpartyAddressDelete/"
	///
	//CounterpartyAddress_Redirect provides a page to return to aftern an action
	CounterpartyAddress_Redirect = CounterpartyAddress_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	CounterpartyAddress_NameFirm_sql   = "NameFirm" // NameFirm is a String
	CounterpartyAddress_NameCentre_sql   = "NameCentre" // NameCentre is a String
	CounterpartyAddress_Address1_sql   = "Address1" // Address1 is a String
	CounterpartyAddress_Address2_sql   = "Address2" // Address2 is a String
	CounterpartyAddress_CityTown_sql   = "CityTown" // CityTown is a String
	CounterpartyAddress_County_sql   = "County" // County is a String
	CounterpartyAddress_Postcode_sql   = "Postcode" // Postcode is a String
	CounterpartyAddress_CompID_sql   = "CompID" // CompID is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	CounterpartyAddress_NameFirm_scrn   = "NameFirm" // NameFirm is a String
	CounterpartyAddress_NameCentre_scrn   = "NameCentre" // NameCentre is a String
	CounterpartyAddress_Address1_scrn   = "Address1" // Address1 is a String
	CounterpartyAddress_Address2_scrn   = "Address2" // Address2 is a String
	CounterpartyAddress_CityTown_scrn   = "CityTown" // CityTown is a String
	CounterpartyAddress_County_scrn   = "County" // County is a String
	CounterpartyAddress_Postcode_scrn   = "Postcode" // Postcode is a String
	CounterpartyAddress_CompID_scrn   = "CompID" // CompID is a String

	/// Definitions End
	///
)

//counterpartyaddress_PageList provides the information for the template for a list of CounterpartyAddresss
type CounterpartyAddress_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []CounterpartyAddress
	Context	 appContext
}

//counterpartyaddress_Page provides the information for the template for an individual CounterpartyAddress
type CounterpartyAddress_Page struct {
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
	Address1         string
	Address1_props     FieldProperties
	Address2         string
	Address2_props     FieldProperties
	CityTown         string
	CityTown_props     FieldProperties
	County         string
	County_props     FieldProperties
	Postcode         string
	Postcode_props     FieldProperties
	CompID         string
	CompID_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}