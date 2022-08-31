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
// Date & Time		    : 28/06/2022 at 16:10:42
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Book defines the datamolde for the Book object
type Book struct {


BookName       string
BookName_props FieldProperties
FullName       string
FullName_props FieldProperties
PLManage       string
PLManage_props FieldProperties
PLTransfer       string
PLTransfer_props FieldProperties
DerivePL       string
DerivePL_props FieldProperties
CostOfCarry       string
CostOfCarry_props FieldProperties
CostOfFunding       string
CostOfFunding_props FieldProperties
LotAllocationMethod       string
LotAllocationMethod_props FieldProperties
InternalId       string
InternalId_props FieldProperties
 // Any lookups will be added below









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
	Book_TemplateList = "/Book/Book_List"
	Book_TemplateView = "/Book/Book_View"
	Book_TemplateEdit = "/Book/Book_Edit"
	Book_TemplateNew  = "/Book/Book_New"
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
	//Book_Redirect provides a page to return to aftern an action
	Book_Redirect = Book_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	Book_BookName_sql   = "BookName" // BookName is a String
	Book_FullName_sql   = "FullName" // FullName is a String
	Book_PLManage_sql   = "PLManage" // PLManage is a String
	Book_PLTransfer_sql   = "PLTransfer" // PLTransfer is a String
	Book_DerivePL_sql   = "DerivePL" // DerivePL is a Bool
	Book_CostOfCarry_sql   = "CostOfCarry" // CostOfCarry is a Bool
	Book_CostOfFunding_sql   = "CostOfFunding" // CostOfFunding is a Bool
	Book_LotAllocationMethod_sql   = "LotAllocationMethod" // LotAllocationMethod is a String
	Book_InternalId_sql   = "InternalId" // InternalId is a Int

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Book_BookName_scrn   = "BookName" // BookName is a String
	Book_FullName_scrn   = "FullName" // FullName is a String
	Book_PLManage_scrn   = "PLManage" // PLManage is a String
	Book_PLTransfer_scrn   = "PLTransfer" // PLTransfer is a String
	Book_DerivePL_scrn   = "DerivePL" // DerivePL is a Bool
	Book_CostOfCarry_scrn   = "CostOfCarry" // CostOfCarry is a Bool
	Book_CostOfFunding_scrn   = "CostOfFunding" // CostOfFunding is a Bool
	Book_LotAllocationMethod_scrn   = "LotAllocationMethod" // LotAllocationMethod is a String
	Book_InternalId_scrn   = "InternalId" // InternalId is a Int

	/// Definitions End
	///
)

//book_PageList provides the information for the template for a list of Books
type Book_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Book
	Context	 appContext
}

//book_Page provides the information for the template for an individual Book
type Book_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	BookName         string
	BookName_props     FieldProperties
	FullName         string
	FullName_props     FieldProperties
	PLManage         string
	PLManage_props     FieldProperties
	PLTransfer         string
	PLTransfer_props     FieldProperties
	DerivePL         string
	DerivePL_props     FieldProperties
	CostOfCarry         string
	CostOfCarry_props     FieldProperties
	CostOfFunding         string
	CostOfFunding_props     FieldProperties
	LotAllocationMethod         string
	LotAllocationMethod_props     FieldProperties
	InternalId         string
	InternalId_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}