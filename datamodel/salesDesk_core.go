package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/salesdesk.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : SalesDesk (salesdesk)
// Endpoint 	        : SalesDesk (Desk)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:56
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//SalesDesk defines the datamolde for the SalesDesk object
type SalesDesk struct {


Name       string
Name_props FieldProperties
ReportDealsOver       string
ReportDealsOver_props FieldProperties
ReportDealsOverCCY       string
ReportDealsOverCCY_props FieldProperties
AccountTransferCutOffTime       string
AccountTransferCutOffTime_props FieldProperties
AccountTransferCutOffTimeTimeZone       string
AccountTransferCutOffTimeTimeZone_props FieldProperties
AccountTransferCutOffTimeCutOffPeriod       string
AccountTransferCutOffTimeCutOffPeriod_props FieldProperties
 // Any lookups will be added below






}

const (
	SalesDesk_Title       = "Sales Desk"
	SalesDesk_SQLTable    = "sienaSalesDesk"
	SalesDesk_SQLSearchID = "Name"
	SalesDesk_QueryString = "Desk"
	///
	/// Handler Defintions
	///
	SalesDesk_Template     = "SalesDesk"
	SalesDesk_TemplateList = "/SalesDesk/SalesDesk_List"
	SalesDesk_TemplateView = "/SalesDesk/SalesDesk_View"
	SalesDesk_TemplateEdit = "/SalesDesk/SalesDesk_Edit"
	SalesDesk_TemplateNew  = "/SalesDesk/SalesDesk_New"
	///
	/// Handler Monitor Paths
	///
	SalesDesk_Path       = "/API/SalesDesk/"
	SalesDesk_PathList   = "/SalesDeskList/"
	SalesDesk_PathView   = "/SalesDeskView/"
	SalesDesk_PathEdit   = "/SalesDeskEdit/"
	SalesDesk_PathNew    = "/SalesDeskNew/"
	SalesDesk_PathSave   = "/SalesDeskSave/"
	SalesDesk_PathDelete = "/SalesDeskDelete/"
	///
	//SalesDesk_Redirect provides a page to return to aftern an action
	SalesDesk_Redirect = SalesDesk_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	SalesDesk_Name_sql   = "Name" // Name is a String
	SalesDesk_ReportDealsOver_sql   = "ReportDealsOver" // ReportDealsOver is a String
	SalesDesk_ReportDealsOverCCY_sql   = "ReportDealsOverCCY" // ReportDealsOverCCY is a String
	SalesDesk_AccountTransferCutOffTime_sql   = "AccountTransferCutOffTime" // AccountTransferCutOffTime is a Time
	SalesDesk_AccountTransferCutOffTimeTimeZone_sql   = "AccountTransferCutOffTimeTimeZone" // AccountTransferCutOffTimeTimeZone is a String
	SalesDesk_AccountTransferCutOffTimeCutOffPeriod_sql   = "AccountTransferCutOffTimeCutOffPeriod" // AccountTransferCutOffTimeCutOffPeriod is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	SalesDesk_Name_scrn   = "Name" // Name is a String
	SalesDesk_ReportDealsOver_scrn   = "ReportDealsOver" // ReportDealsOver is a String
	SalesDesk_ReportDealsOverCCY_scrn   = "ReportDealsOverCCY" // ReportDealsOverCCY is a String
	SalesDesk_AccountTransferCutOffTime_scrn   = "AccountTransferCutOffTime" // AccountTransferCutOffTime is a Time
	SalesDesk_AccountTransferCutOffTimeTimeZone_scrn   = "AccountTransferCutOffTimeTimeZone" // AccountTransferCutOffTimeTimeZone is a String
	SalesDesk_AccountTransferCutOffTimeCutOffPeriod_scrn   = "AccountTransferCutOffTimeCutOffPeriod" // AccountTransferCutOffTimeCutOffPeriod is a String

	/// Definitions End
	///
)

//salesdesk_PageList provides the information for the template for a list of SalesDesks
type SalesDesk_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []SalesDesk
	Context	 appContext
}

//salesdesk_Page provides the information for the template for an individual SalesDesk
type SalesDesk_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	Name         string
	Name_props     FieldProperties
	ReportDealsOver         string
	ReportDealsOver_props     FieldProperties
	ReportDealsOverCCY         string
	ReportDealsOverCCY_props     FieldProperties
	AccountTransferCutOffTime         string
	AccountTransferCutOffTime_props     FieldProperties
	AccountTransferCutOffTimeTimeZone         string
	AccountTransferCutOffTimeTimeZone_props     FieldProperties
	AccountTransferCutOffTimeCutOffPeriod         string
	AccountTransferCutOffTimeCutOffPeriod_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}