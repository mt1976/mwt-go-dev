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
// Date & Time		    : 12/12/2021 at 16:13:09
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type CounterpartyAddress struct {

NameFirm        string
NameCentre        string
Address1        string
Address2        string
CityTown        string
County        string
Postcode        string
CompID        string

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
	CounterpartyAddress_TemplateList = "CounterpartyAddress_List"
	CounterpartyAddress_TemplateView = "CounterpartyAddress_View"
	CounterpartyAddress_TemplateEdit = "CounterpartyAddress_Edit"
	CounterpartyAddress_TemplateNew  = "CounterpartyAddress_New"
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
	/// SQL Field Definitions
	///
	CounterpartyAddress_NameFirm   = "NameFirm" // NameFirm is a String
	CounterpartyAddress_NameCentre   = "NameCentre" // NameCentre is a String
	CounterpartyAddress_Address1   = "Address1" // Address1 is a String
	CounterpartyAddress_Address2   = "Address2" // Address2 is a String
	CounterpartyAddress_CityTown   = "CityTown" // CityTown is a String
	CounterpartyAddress_County   = "County" // County is a String
	CounterpartyAddress_Postcode   = "Postcode" // Postcode is a String
	CounterpartyAddress_CompID   = "CompID" // CompID is a String

	/// Definitions End
)
