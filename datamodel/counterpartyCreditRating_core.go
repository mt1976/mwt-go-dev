package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/counterpartycreditrating.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyCreditRating (counterpartycreditrating)
// Endpoint 	        : CounterpartyCreditRating (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:47
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CounterpartyCreditRating defines the datamolde for the CounterpartyCreditRating object
type CounterpartyCreditRating struct {


NameFirm       string
NameFirm_props FieldProperties
NameCentre       string
NameCentre_props FieldProperties
CreditRatingUsage       string
CreditRatingUsage_props FieldProperties
CreditRatingAgency       string
CreditRatingAgency_props FieldProperties
CreditRatingName       string
CreditRatingName_props FieldProperties
CompID       string
CompID_props FieldProperties
 // Any lookups will be added below






}

const (
	CounterpartyCreditRating_Title       = "Counterparty Credit Rating"
	CounterpartyCreditRating_SQLTable    = "sienaCounterpartyCreditRating"
	CounterpartyCreditRating_SQLSearchID = "CompID"
	CounterpartyCreditRating_QueryString = "ID"
	///
	/// Handler Defintions
	///
	CounterpartyCreditRating_Template     = "CounterpartyCreditRating"
	CounterpartyCreditRating_TemplateList = "/CounterpartyCreditRating/CounterpartyCreditRating_List"
	CounterpartyCreditRating_TemplateView = "/CounterpartyCreditRating/CounterpartyCreditRating_View"
	CounterpartyCreditRating_TemplateEdit = "/CounterpartyCreditRating/CounterpartyCreditRating_Edit"
	CounterpartyCreditRating_TemplateNew  = "/CounterpartyCreditRating/CounterpartyCreditRating_New"
	///
	/// Handler Monitor Paths
	///
	CounterpartyCreditRating_Path       = "/API/CounterpartyCreditRating/"
	CounterpartyCreditRating_PathList   = "/CounterpartyCreditRatingList/"
	CounterpartyCreditRating_PathView   = "/CounterpartyCreditRatingView/"
	CounterpartyCreditRating_PathEdit   = "/CounterpartyCreditRatingEdit/"
	CounterpartyCreditRating_PathNew    = "/CounterpartyCreditRatingNew/"
	CounterpartyCreditRating_PathSave   = "/CounterpartyCreditRatingSave/"
	CounterpartyCreditRating_PathDelete = "/CounterpartyCreditRatingDelete/"
	///
	//CounterpartyCreditRating_Redirect provides a page to return to aftern an action
	CounterpartyCreditRating_Redirect = CounterpartyCreditRating_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	CounterpartyCreditRating_NameFirm_sql   = "NameFirm" // NameFirm is a String
	CounterpartyCreditRating_NameCentre_sql   = "NameCentre" // NameCentre is a String
	CounterpartyCreditRating_CreditRatingUsage_sql   = "CreditRatingUsage" // CreditRatingUsage is a String
	CounterpartyCreditRating_CreditRatingAgency_sql   = "CreditRatingAgency" // CreditRatingAgency is a String
	CounterpartyCreditRating_CreditRatingName_sql   = "CreditRatingName" // CreditRatingName is a String
	CounterpartyCreditRating_CompID_sql   = "CompID" // CompID is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	CounterpartyCreditRating_NameFirm_scrn   = "NameFirm" // NameFirm is a String
	CounterpartyCreditRating_NameCentre_scrn   = "NameCentre" // NameCentre is a String
	CounterpartyCreditRating_CreditRatingUsage_scrn   = "CreditRatingUsage" // CreditRatingUsage is a String
	CounterpartyCreditRating_CreditRatingAgency_scrn   = "CreditRatingAgency" // CreditRatingAgency is a String
	CounterpartyCreditRating_CreditRatingName_scrn   = "CreditRatingName" // CreditRatingName is a String
	CounterpartyCreditRating_CompID_scrn   = "CompID" // CompID is a String

	/// Definitions End
	///
)

//counterpartycreditrating_PageList provides the information for the template for a list of CounterpartyCreditRatings
type CounterpartyCreditRating_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []CounterpartyCreditRating
	Context	 appContext
}

//counterpartycreditrating_Page provides the information for the template for an individual CounterpartyCreditRating
type CounterpartyCreditRating_Page struct {
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
	CreditRatingUsage         string
	CreditRatingUsage_props     FieldProperties
	CreditRatingAgency         string
	CreditRatingAgency_props     FieldProperties
	CreditRatingName         string
	CreditRatingName_props     FieldProperties
	CompID         string
	CompID_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}