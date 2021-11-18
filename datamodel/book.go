package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/book.go"
// ----------------------------------------------------------------
// Package            : datamodel
// Object 			  : Book
// Endpoint Root 	  : Book
// Search QueryString : Book
// From   			  : 
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 18/11/2021 at 17:07:04
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Book struct {
	AppInternalID string  // Special field for internal use only
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
	Book_Title       = "Books"
	Book_SQLTable    = "sienaBook"
	Book_SQLSearchID = "BookName"
	Book_QueryString = "Book"
	///
	/// Handler Defintions
	///
	Book_TemplateList = "Book_List"
	Book_TemplateView = "Book_View"
	Book_TemplateEdit = "Book_Edit"
	Book_TemplateNew  = "Book_New"
	///
	/// Handler Monitor Paths
	///
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
