package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/customeronboarding.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CustomerOnboarding (customeronboarding)
// Endpoint 	        : CustomerOnboarding (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:31:45
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CustomerOnboarding defines the datamolde for the CustomerOnboarding object
type CustomerOnboarding struct {

ID       string
CustomerName       string
CustomerAddress       string
CustomerBOLID       string
CustomerFirmName       string
CustomerType       string
CustomerRDC       string
CustomerSortCode       string
CustomerGMClientNo       string
CustomerDefaultBook       string
CustomerRegion       string
CustomerCategory       string
CustomerTelephoneNo       string

}

const (
	CustomerOnboarding_Title       = "Customer Onboarding Tool"
	CustomerOnboarding_SQLTable    = ""
	CustomerOnboarding_SQLSearchID = "ID"
	CustomerOnboarding_QueryString = "ID"
	///
	/// Handler Defintions
	///
	CustomerOnboarding_Template     = "CustomerOnboarding"
	CustomerOnboarding_TemplateList = "CustomerOnboarding_List"
	CustomerOnboarding_TemplateView = "CustomerOnboarding_View"
	CustomerOnboarding_TemplateEdit = "CustomerOnboarding_Edit"
	CustomerOnboarding_TemplateNew  = "CustomerOnboarding_New"
	///
	/// Handler Monitor Paths
	///
	CustomerOnboarding_Path       = "/API/CustomerOnboarding/"
	CustomerOnboarding_PathList   = "/CustomerOnboardingList/"
	CustomerOnboarding_PathView   = "/CustomerOnboardingView/"
	CustomerOnboarding_PathEdit   = "/CustomerOnboardingEdit/"
	CustomerOnboarding_PathNew    = "/CustomerOnboardingNew/"
	CustomerOnboarding_PathSave   = "/CustomerOnboardingSave/"
	CustomerOnboarding_PathDelete = "/CustomerOnboardingDelete/"
	///
	/// SQL Field Definitions
	///
CustomerOnboarding_ID   = "ID" // ID is a String
CustomerOnboarding_CustomerName   = "CustomerName" // CustomerName is a String
CustomerOnboarding_CustomerAddress   = "CustomerAddress" // CustomerAddress is a String
CustomerOnboarding_CustomerBOLID   = "CustomerBOLID" // CustomerBOLID is a String
CustomerOnboarding_CustomerFirmName   = "CustomerFirmName" // CustomerFirmName is a String
CustomerOnboarding_CustomerType   = "CustomerType" // CustomerType is a String
CustomerOnboarding_CustomerRDC   = "CustomerRDC" // CustomerRDC is a String
CustomerOnboarding_CustomerSortCode   = "CustomerSortCode" // CustomerSortCode is a String
CustomerOnboarding_CustomerGMClientNo   = "CustomerGMClientNo" // CustomerGMClientNo is a String
CustomerOnboarding_CustomerDefaultBook   = "CustomerDefaultBook" // CustomerDefaultBook is a String
CustomerOnboarding_CustomerRegion   = "CustomerRegion" // CustomerRegion is a String
CustomerOnboarding_CustomerCategory   = "CustomerCategory" // CustomerCategory is a String
CustomerOnboarding_CustomerTelephoneNo   = "CustomerTelephoneNo" // CustomerTelephoneNo is a String

	/// Definitions End
)
