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
// Date & Time		    : 12/12/2021 at 16:13:18
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type SalesDesk struct {

Name        string
ReportDealsOver        string
ReportDealsOverCCY        string
AccountTransferCutOffTime        string
AccountTransferCutOffTimeTimeZone        string
AccountTransferCutOffTimeCutOffPeriod        string

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
	SalesDesk_TemplateList = "SalesDesk_List"
	SalesDesk_TemplateView = "SalesDesk_View"
	SalesDesk_TemplateEdit = "SalesDesk_Edit"
	SalesDesk_TemplateNew  = "SalesDesk_New"
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
	/// SQL Field Definitions
	///
	SalesDesk_Name   = "Name" // Name is a String
	SalesDesk_ReportDealsOver   = "ReportDealsOver" // ReportDealsOver is a String
	SalesDesk_ReportDealsOverCCY   = "ReportDealsOverCCY" // ReportDealsOverCCY is a String
	SalesDesk_AccountTransferCutOffTime   = "AccountTransferCutOffTime" // AccountTransferCutOffTime is a Time
	SalesDesk_AccountTransferCutOffTimeTimeZone   = "AccountTransferCutOffTimeTimeZone" // AccountTransferCutOffTimeTimeZone is a String
	SalesDesk_AccountTransferCutOffTimeCutOffPeriod   = "AccountTransferCutOffTimeCutOffPeriod" // AccountTransferCutOffTimeCutOffPeriod is a String

	/// Definitions End
)
