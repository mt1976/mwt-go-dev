package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/accountladder.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : AccountLadder (accountladder)
// Endpoint 	        : AccountLadder (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:06
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type AccountLadder struct {

SienaReference        string
BusinessDate        string
ContractNumber        string
Balance        string
DealtCcy        string
AmountDp        string

}

const (
	AccountLadder_Title       = "Account Ladder"
	AccountLadder_SQLTable    = "sienaAccountLadder"
	AccountLadder_SQLSearchID = "SienaReference"
	AccountLadder_QueryString = "AccountNo"
	///
	/// Handler Defintions
	///
	AccountLadder_Template     = "AccountLadder"
	AccountLadder_TemplateList = "AccountLadder_List"
	AccountLadder_TemplateView = "AccountLadder_View"
	AccountLadder_TemplateEdit = "AccountLadder_Edit"
	AccountLadder_TemplateNew  = "AccountLadder_New"
	///
	/// Handler Monitor Paths
	///
	AccountLadder_Path       = "/API/AccountLadder/"
	AccountLadder_PathList   = "/AccountLadderList/"
	AccountLadder_PathView   = "/AccountLadderView/"
	AccountLadder_PathEdit   = "/AccountLadderEdit/"
	AccountLadder_PathNew    = "/AccountLadderNew/"
	AccountLadder_PathSave   = "/AccountLadderSave/"
	AccountLadder_PathDelete = "/AccountLadderDelete/"
	///
	/// SQL Field Definitions
	///
	AccountLadder_SienaReference   = "SienaReference" // SienaReference is a String
	AccountLadder_BusinessDate   = "BusinessDate" // BusinessDate is a Time
	AccountLadder_ContractNumber   = "ContractNumber" // ContractNumber is a String
	AccountLadder_Balance   = "Balance" // Balance is a Float
	AccountLadder_DealtCcy   = "DealtCcy" // DealtCcy is a String
	AccountLadder_AmountDp   = "AmountDp" // AmountDp is a Int

	/// Definitions End
)
