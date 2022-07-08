package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/product.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Product (product)
// Endpoint 	        : Product (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:55
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Product defines the datamolde for the Product object
type Product struct {


Code       string
Code_props FieldProperties
Name       string
Name_props FieldProperties
Factor       string
Factor_props FieldProperties
MaxTermPrecedence       string
MaxTermPrecedence_props FieldProperties
InternalId       string
InternalId_props FieldProperties
InternalDeleted       string
InternalDeleted_props FieldProperties
UpdatedTransactionId       string
UpdatedTransactionId_props FieldProperties
UpdatedUserId       string
UpdatedUserId_props FieldProperties
UpdatedDateTime       string
UpdatedDateTime_props FieldProperties
DeletedTransactionId       string
DeletedTransactionId_props FieldProperties
DeletedUserId       string
DeletedUserId_props FieldProperties
ChangeType       string
ChangeType_props FieldProperties
 // Any lookups will be added below












}

const (
	Product_Title       = "Product Group"
	Product_SQLTable    = "sienaProduct"
	Product_SQLSearchID = "Code"
	Product_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Product_Template     = "Product"
	Product_TemplateList = "/Product/Product_List"
	Product_TemplateView = "/Product/Product_View"
	Product_TemplateEdit = "/Product/Product_Edit"
	Product_TemplateNew  = "/Product/Product_New"
	///
	/// Handler Monitor Paths
	///
	Product_Path       = "/API/Product/"
	Product_PathList   = "/ProductList/"
	Product_PathView   = "/ProductView/"
	Product_PathEdit   = "/ProductEdit/"
	Product_PathNew    = "/ProductNew/"
	Product_PathSave   = "/ProductSave/"
	Product_PathDelete = "/ProductDelete/"
	///
	//Product_Redirect provides a page to return to aftern an action
	Product_Redirect = Product_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	Product_Code_sql   = "Code" // Code is a String
	Product_Name_sql   = "Name" // Name is a String
	Product_Factor_sql   = "Factor" // Factor is a Float
	Product_MaxTermPrecedence_sql   = "MaxTermPrecedence" // MaxTermPrecedence is a Bool
	Product_InternalId_sql   = "InternalId" // InternalId is a Int
	Product_InternalDeleted_sql   = "InternalDeleted" // InternalDeleted is a Time
	Product_UpdatedTransactionId_sql   = "UpdatedTransactionId" // UpdatedTransactionId is a String
	Product_UpdatedUserId_sql   = "UpdatedUserId" // UpdatedUserId is a String
	Product_UpdatedDateTime_sql   = "UpdatedDateTime" // UpdatedDateTime is a Time
	Product_DeletedTransactionId_sql   = "DeletedTransactionId" // DeletedTransactionId is a String
	Product_DeletedUserId_sql   = "DeletedUserId" // DeletedUserId is a String
	Product_ChangeType_sql   = "ChangeType" // ChangeType is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	Product_Code_scrn   = "Code" // Code is a String
	Product_Name_scrn   = "Name" // Name is a String
	Product_Factor_scrn   = "Factor" // Factor is a Float
	Product_MaxTermPrecedence_scrn   = "MaxTermPrecedence" // MaxTermPrecedence is a Bool
	Product_InternalId_scrn   = "InternalId" // InternalId is a Int
	Product_InternalDeleted_scrn   = "InternalDeleted" // InternalDeleted is a Time
	Product_UpdatedTransactionId_scrn   = "UpdatedTransactionId" // UpdatedTransactionId is a String
	Product_UpdatedUserId_scrn   = "UpdatedUserId" // UpdatedUserId is a String
	Product_UpdatedDateTime_scrn   = "UpdatedDateTime" // UpdatedDateTime is a Time
	Product_DeletedTransactionId_scrn   = "DeletedTransactionId" // DeletedTransactionId is a String
	Product_DeletedUserId_scrn   = "DeletedUserId" // DeletedUserId is a String
	Product_ChangeType_scrn   = "ChangeType" // ChangeType is a String

	/// Definitions End
	///
)

//product_PageList provides the information for the template for a list of Products
type Product_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []Product
	Context	 appContext
}

//product_Page provides the information for the template for an individual Product
type Product_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	Code         string
	Code_props     FieldProperties
	Name         string
	Name_props     FieldProperties
	Factor         string
	Factor_props     FieldProperties
	MaxTermPrecedence         string
	MaxTermPrecedence_props     FieldProperties
	InternalId         string
	InternalId_props     FieldProperties
	InternalDeleted         string
	InternalDeleted_props     FieldProperties
	UpdatedTransactionId         string
	UpdatedTransactionId_props     FieldProperties
	UpdatedUserId         string
	UpdatedUserId_props     FieldProperties
	UpdatedDateTime         string
	UpdatedDateTime_props     FieldProperties
	DeletedTransactionId         string
	DeletedTransactionId_props     FieldProperties
	DeletedUserId         string
	DeletedUserId_props     FieldProperties
	ChangeType         string
	ChangeType_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}