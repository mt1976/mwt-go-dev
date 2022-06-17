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
// Date & Time		    : 17/06/2022 at 18:38:03
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
	CustomerOnboarding_TemplateList = "/CustomerOnboarding/CustomerOnboarding_List"
	CustomerOnboarding_TemplateView = "/CustomerOnboarding/CustomerOnboarding_View"
	CustomerOnboarding_TemplateEdit = "/CustomerOnboarding/CustomerOnboarding_Edit"
	CustomerOnboarding_TemplateNew  = "/CustomerOnboarding/CustomerOnboarding_New"
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
	///
	/// SQL Field Definitions
	///
CustomerOnboarding_ID_sql   = "ID" // ID is a String
CustomerOnboarding_CustomerName_sql   = "CustomerName" // CustomerName is a String
CustomerOnboarding_CustomerAddress_sql   = "CustomerAddress" // CustomerAddress is a String
CustomerOnboarding_CustomerBOLID_sql   = "CustomerBOLID" // CustomerBOLID is a String
CustomerOnboarding_CustomerFirmName_sql   = "CustomerFirmName" // CustomerFirmName is a String
CustomerOnboarding_CustomerType_sql   = "CustomerType" // CustomerType is a String
CustomerOnboarding_CustomerRDC_sql   = "CustomerRDC" // CustomerRDC is a String
CustomerOnboarding_CustomerSortCode_sql   = "CustomerSortCode" // CustomerSortCode is a String
CustomerOnboarding_CustomerGMClientNo_sql   = "CustomerGMClientNo" // CustomerGMClientNo is a String
CustomerOnboarding_CustomerDefaultBook_sql   = "CustomerDefaultBook" // CustomerDefaultBook is a String
CustomerOnboarding_CustomerRegion_sql   = "CustomerRegion" // CustomerRegion is a String
CustomerOnboarding_CustomerCategory_sql   = "CustomerCategory" // CustomerCategory is a String
CustomerOnboarding_CustomerTelephoneNo_sql   = "CustomerTelephoneNo" // CustomerTelephoneNo is a String

	/// Definitions End

	/// Application Field Definitions
	///
CustomerOnboarding_ID_scrn   = "ID" // ID is a String
CustomerOnboarding_CustomerName_scrn   = "CustomerName" // CustomerName is a String
CustomerOnboarding_CustomerAddress_scrn   = "CustomerAddress" // CustomerAddress is a String
CustomerOnboarding_CustomerBOLID_scrn   = "CustomerBOLID" // CustomerBOLID is a String
CustomerOnboarding_CustomerFirmName_scrn   = "CustomerFirmName" // CustomerFirmName is a String
CustomerOnboarding_CustomerType_scrn   = "CustomerType" // CustomerType is a String
CustomerOnboarding_CustomerRDC_scrn   = "CustomerRDC" // CustomerRDC is a String
CustomerOnboarding_CustomerSortCode_scrn   = "CustomerSortCode" // CustomerSortCode is a String
CustomerOnboarding_CustomerGMClientNo_scrn   = "CustomerGMClientNo" // CustomerGMClientNo is a String
CustomerOnboarding_CustomerDefaultBook_scrn   = "CustomerDefaultBook" // CustomerDefaultBook is a String
CustomerOnboarding_CustomerRegion_scrn   = "CustomerRegion" // CustomerRegion is a String
CustomerOnboarding_CustomerCategory_scrn   = "CustomerCategory" // CustomerCategory is a String
CustomerOnboarding_CustomerTelephoneNo_scrn   = "CustomerTelephoneNo" // CustomerTelephoneNo is a String

	/// Definitions End
)
