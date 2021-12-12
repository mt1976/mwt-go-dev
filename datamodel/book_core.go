package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/book.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Book (book)
// Endpoint 	        : Book (Book)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:07
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Book struct {

BookName        string
FullName        string
PLManage        string
PLTransfer        string
DerivePL        string
CostOfCarry        string
CostOfFunding        string
LotAllocationMethod        string
InternalId        string

}

const (
	Book_Title       = "Book"
	Book_SQLTable    = "sienaBook"
	Book_SQLSearchID = "BookName"
	Book_QueryString = "Book"
	///
	/// Handler Defintions
	///
	Book_Template     = "Book"
	Book_TemplateList = "Book_List"
	Book_TemplateView = "Book_View"
	Book_TemplateEdit = "Book_Edit"
	Book_TemplateNew  = "Book_New"
	///
	/// Handler Monitor Paths
	///
	Book_Path       = "/API/Book/"
	Book_PathList   = "/BookList/"
	Book_PathView   = "/BookView/"
	Book_PathEdit   = "/BookEdit/"
	Book_PathNew    = "/BookNew/"
	Book_PathSave   = "/BookSave/"
	Book_PathDelete = "/BookDelete/"
	///
	/// SQL Field Definitions
	///
	Book_BookName   = "BookName" // BookName is a String
	Book_FullName   = "FullName" // FullName is a String
	Book_PLManage   = "PLManage" // PLManage is a String
	Book_PLTransfer   = "PLTransfer" // PLTransfer is a String
	Book_DerivePL   = "DerivePL" // DerivePL is a Bool
	Book_CostOfCarry   = "CostOfCarry" // CostOfCarry is a Bool
	Book_CostOfFunding   = "CostOfFunding" // CostOfFunding is a Bool
	Book_LotAllocationMethod   = "LotAllocationMethod" // LotAllocationMethod is a String
	Book_InternalId   = "InternalId" // InternalId is a Int

	/// Definitions End
)
