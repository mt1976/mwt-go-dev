package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/accounttransaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : AccountTransaction (accounttransaction)
// Endpoint 	        : AccountTransaction (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:42
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//AccountTransaction defines the datamolde for the AccountTransaction object
type AccountTransaction struct {


SienaReference       string
SienaReference_props FieldProperties
LegNo       string
LegNo_props FieldProperties
MMLegNo       string
MMLegNo_props FieldProperties
Narrative       string
Narrative_props FieldProperties
Amount       string
Amount_props FieldProperties
StartInterestDate       string
StartInterestDate_props FieldProperties
EndInterestDate       string
EndInterestDate_props FieldProperties
Amortisation       string
Amortisation_props FieldProperties
InterestAmount       string
InterestAmount_props FieldProperties
InterestAction       string
InterestAction_props FieldProperties
FixingDate       string
FixingDate_props FieldProperties
InterestCalculationDate       string
InterestCalculationDate_props FieldProperties
AmendmentAmount       string
AmendmentAmount_props FieldProperties
DealtCcy       string
DealtCcy_props FieldProperties
AmountDp       string
AmountDp_props FieldProperties
 // Any lookups will be added below















}

const (
	AccountTransaction_Title       = "Account Transaction"
	AccountTransaction_SQLTable    = "sienaAccountTransactions"
	AccountTransaction_SQLSearchID = "SienaReference"
	AccountTransaction_QueryString = "AccountNo"
	///
	/// Handler Defintions
	///
	AccountTransaction_Template     = "AccountTransaction"
	AccountTransaction_TemplateList = "/AccountTransaction/AccountTransaction_List"
	AccountTransaction_TemplateView = "/AccountTransaction/AccountTransaction_View"
	AccountTransaction_TemplateEdit = "/AccountTransaction/AccountTransaction_Edit"
	AccountTransaction_TemplateNew  = "/AccountTransaction/AccountTransaction_New"
	///
	/// Handler Monitor Paths
	///
	AccountTransaction_Path       = "/API/AccountTransaction/"
	AccountTransaction_PathList   = "/AccountTransactionList/"
	AccountTransaction_PathView   = "/AccountTransactionView/"
	AccountTransaction_PathEdit   = "/AccountTransactionEdit/"
	AccountTransaction_PathNew    = "/AccountTransactionNew/"
	AccountTransaction_PathSave   = "/AccountTransactionSave/"
	AccountTransaction_PathDelete = "/AccountTransactionDelete/"
	///
	//AccountTransaction_Redirect provides a page to return to aftern an action
	AccountTransaction_Redirect = AccountTransaction_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	AccountTransaction_SienaReference_sql   = "SienaReference" // SienaReference is a String
	AccountTransaction_LegNo_sql   = "LegNo" // LegNo is a Int
	AccountTransaction_MMLegNo_sql   = "MMLegNo" // MMLegNo is a Int
	AccountTransaction_Narrative_sql   = "Narrative" // Narrative is a String
	AccountTransaction_Amount_sql   = "Amount" // Amount is a Float
	AccountTransaction_StartInterestDate_sql   = "StartInterestDate" // StartInterestDate is a Time
	AccountTransaction_EndInterestDate_sql   = "EndInterestDate" // EndInterestDate is a Time
	AccountTransaction_Amortisation_sql   = "Amortisation" // Amortisation is a Float
	AccountTransaction_InterestAmount_sql   = "InterestAmount" // InterestAmount is a Float
	AccountTransaction_InterestAction_sql   = "InterestAction" // InterestAction is a String
	AccountTransaction_FixingDate_sql   = "FixingDate" // FixingDate is a Time
	AccountTransaction_InterestCalculationDate_sql   = "InterestCalculationDate" // InterestCalculationDate is a Time
	AccountTransaction_AmendmentAmount_sql   = "AmendmentAmount" // AmendmentAmount is a Float
	AccountTransaction_DealtCcy_sql   = "DealtCcy" // DealtCcy is a String
	AccountTransaction_AmountDp_sql   = "AmountDp" // AmountDp is a Int

	/// Definitions End
	///
	/// Application Field Definitions
	///
	AccountTransaction_SienaReference_scrn   = "SienaReference" // SienaReference is a String
	AccountTransaction_LegNo_scrn   = "LegNo" // LegNo is a Int
	AccountTransaction_MMLegNo_scrn   = "MMLegNo" // MMLegNo is a Int
	AccountTransaction_Narrative_scrn   = "Narrative" // Narrative is a String
	AccountTransaction_Amount_scrn   = "Amount" // Amount is a Float
	AccountTransaction_StartInterestDate_scrn   = "StartInterestDate" // StartInterestDate is a Time
	AccountTransaction_EndInterestDate_scrn   = "EndInterestDate" // EndInterestDate is a Time
	AccountTransaction_Amortisation_scrn   = "Amortisation" // Amortisation is a Float
	AccountTransaction_InterestAmount_scrn   = "InterestAmount" // InterestAmount is a Float
	AccountTransaction_InterestAction_scrn   = "InterestAction" // InterestAction is a String
	AccountTransaction_FixingDate_scrn   = "FixingDate" // FixingDate is a Time
	AccountTransaction_InterestCalculationDate_scrn   = "InterestCalculationDate" // InterestCalculationDate is a Time
	AccountTransaction_AmendmentAmount_scrn   = "AmendmentAmount" // AmendmentAmount is a Float
	AccountTransaction_DealtCcy_scrn   = "DealtCcy" // DealtCcy is a String
	AccountTransaction_AmountDp_scrn   = "AmountDp" // AmountDp is a Int

	/// Definitions End
	///
)

//accounttransaction_PageList provides the information for the template for a list of AccountTransactions
type AccountTransaction_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []AccountTransaction
	Context	 appContext
}

//accounttransaction_Page provides the information for the template for an individual AccountTransaction
type AccountTransaction_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SienaReference         string
	SienaReference_props     FieldProperties
	LegNo         string
	LegNo_props     FieldProperties
	MMLegNo         string
	MMLegNo_props     FieldProperties
	Narrative         string
	Narrative_props     FieldProperties
	Amount         string
	Amount_props     FieldProperties
	StartInterestDate         string
	StartInterestDate_props     FieldProperties
	EndInterestDate         string
	EndInterestDate_props     FieldProperties
	Amortisation         string
	Amortisation_props     FieldProperties
	InterestAmount         string
	InterestAmount_props     FieldProperties
	InterestAction         string
	InterestAction_props     FieldProperties
	FixingDate         string
	FixingDate_props     FieldProperties
	InterestCalculationDate         string
	InterestCalculationDate_props     FieldProperties
	AmendmentAmount         string
	AmendmentAmount_props     FieldProperties
	DealtCcy         string
	DealtCcy_props     FieldProperties
	AmountDp         string
	AmountDp_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}