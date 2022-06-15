package application

// ----------------------------------------------------------------
// Automatically generated  "/application/salesdesk.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : SalesDesk (salesdesk)
// Endpoint 	        : SalesDesk (Desk)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:32:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"net/http"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

//salesdesk_PageList provides the information for the template for a list of SalesDesks
type SalesDesk_PageList struct {
	SessionInfo dm.SessionInfo
	UserMenu    dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	ItemsOnPage int
	ItemList    []dm.SalesDesk
}

//SalesDesk_Redirect provides a page to return to aftern an action
const (
	SalesDesk_Redirect = dm.SalesDesk_PathList
)

//salesdesk_Page provides the information for the template for an individual SalesDesk
type SalesDesk_Page struct {
	SessionInfo dm.SessionInfo
	UserMenu    dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	Name                                  string
	ReportDealsOver                       string
	ReportDealsOverCCY                    string
	AccountTransferCutOffTime             string
	AccountTransferCutOffTimeTimeZone     string
	AccountTransferCutOffTimeCutOffPeriod string
	//
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}

//SalesDesk_Publish annouces the endpoints available for this object
func SalesDesk_Publish(mux http.ServeMux) {
	//No API
	//Cannot List via GUI
	//Cannot View via GUI
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.SalesDesk_Title)
	//No API
}

// Builds/Popuplates the SalesDesk Page
func salesdesk_PopulatePage(rD dm.SalesDesk, pageDetail SalesDesk_Page) SalesDesk_Page {
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	pageDetail.Name = rD.Name
	pageDetail.ReportDealsOver = rD.ReportDealsOver
	pageDetail.ReportDealsOverCCY = rD.ReportDealsOverCCY
	pageDetail.AccountTransferCutOffTime = rD.AccountTransferCutOffTime
	pageDetail.AccountTransferCutOffTimeTimeZone = rD.AccountTransferCutOffTimeTimeZone
	pageDetail.AccountTransferCutOffTimeCutOffPeriod = rD.AccountTransferCutOffTimeCutOffPeriod

	//
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//

	//
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return pageDetail
}
