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
// Date & Time		    : 12/12/2021 at 16:13:06
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type AccountTransaction struct {

SienaReference        string
LegNo        string
MMLegNo        string
Narrative        string
Amount        string
StartInterestDate        string
EndInterestDate        string
Amortisation        string
InterestAmount        string
InterestAction        string
FixingDate        string
InterestCalculationDate        string
AmendmentAmount        string
DealtCcy        string
AmountDp        string

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
	AccountTransaction_TemplateList = "AccountTransaction_List"
	AccountTransaction_TemplateView = "AccountTransaction_View"
	AccountTransaction_TemplateEdit = "AccountTransaction_Edit"
	AccountTransaction_TemplateNew  = "AccountTransaction_New"
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
	/// SQL Field Definitions
	///
	AccountTransaction_SienaReference   = "SienaReference" // SienaReference is a String
	AccountTransaction_LegNo   = "LegNo" // LegNo is a Int
	AccountTransaction_MMLegNo   = "MMLegNo" // MMLegNo is a Int
	AccountTransaction_Narrative   = "Narrative" // Narrative is a String
	AccountTransaction_Amount   = "Amount" // Amount is a Float
	AccountTransaction_StartInterestDate   = "StartInterestDate" // StartInterestDate is a Time
	AccountTransaction_EndInterestDate   = "EndInterestDate" // EndInterestDate is a Time
	AccountTransaction_Amortisation   = "Amortisation" // Amortisation is a Float
	AccountTransaction_InterestAmount   = "InterestAmount" // InterestAmount is a Float
	AccountTransaction_InterestAction   = "InterestAction" // InterestAction is a String
	AccountTransaction_FixingDate   = "FixingDate" // FixingDate is a Time
	AccountTransaction_InterestCalculationDate   = "InterestCalculationDate" // InterestCalculationDate is a Time
	AccountTransaction_AmendmentAmount   = "AmendmentAmount" // AmendmentAmount is a Float
	AccountTransaction_DealtCcy   = "DealtCcy" // DealtCcy is a String
	AccountTransaction_AmountDp   = "AmountDp" // AmountDp is a Int

	/// Definitions End
)
