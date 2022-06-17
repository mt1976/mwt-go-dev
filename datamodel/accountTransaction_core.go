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
// Date & Time		    : 17/06/2022 at 18:38:05
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//AccountTransaction defines the datamolde for the AccountTransaction object
type AccountTransaction struct {

SienaReference       string
LegNo       string
MMLegNo       string
Narrative       string
Amount       string
StartInterestDate       string
EndInterestDate       string
Amortisation       string
InterestAmount       string
InterestAction       string
FixingDate       string
InterestCalculationDate       string
AmendmentAmount       string
DealtCcy       string
AmountDp       string

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
)
