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
// Date & Time		    : 14/06/2022 at 21:31:54
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//CounterpartyCreditRating defines the datamolde for the CounterpartyCreditRating object
type CounterpartyCreditRating struct {

NameFirm       string
NameCentre       string
CreditRatingUsage       string
CreditRatingAgency       string
CreditRatingName       string
CompID       string

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
	CounterpartyCreditRating_TemplateList = "CounterpartyCreditRating_List"
	CounterpartyCreditRating_TemplateView = "CounterpartyCreditRating_View"
	CounterpartyCreditRating_TemplateEdit = "CounterpartyCreditRating_Edit"
	CounterpartyCreditRating_TemplateNew  = "CounterpartyCreditRating_New"
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
	/// SQL Field Definitions
	///
CounterpartyCreditRating_NameFirm   = "NameFirm" // NameFirm is a String
CounterpartyCreditRating_NameCentre   = "NameCentre" // NameCentre is a String
CounterpartyCreditRating_CreditRatingUsage   = "CreditRatingUsage" // CreditRatingUsage is a String
CounterpartyCreditRating_CreditRatingAgency   = "CreditRatingAgency" // CreditRatingAgency is a String
CounterpartyCreditRating_CreditRatingName   = "CreditRatingName" // CreditRatingName is a String
CounterpartyCreditRating_CompID   = "CompID" // CompID is a String

	/// Definitions End
)
