package application

// ----------------------------------------------------------------
// Automatically generated  "/application/owner.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Owner (owner)
// Endpoint 	        : Owner (Owner)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:12
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"net/http"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

//owner_PageList provides the information for the template for a list of Owners
type Owner_PageList struct {
	SessionInfo dm.SessionInfo
	UserMenu    dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	ItemsOnPage int
	ItemList    []dm.Owner
}

//Owner_Redirect provides a page to return to aftern an action
const (
	Owner_Redirect = dm.Owner_PathList
)

//owner_Page provides the information for the template for an individual Owner
type Owner_Page struct {
	SessionInfo dm.SessionInfo
	UserMenu    dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	UserName         string
	FullName         string
	Type             string
	TradingEntity    string
	DefaultEnterBook string
	EmailAddress     string
	Enabled          string
	ExternalUserIds  string
	Language         string
	LocalCurrency    string
	Role             string
	TelephoneNumber  string
	TokenId          string
	Entity           string
	UserCode         string
	//
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
}

//Owner_Publish annouces the endpoints available for this object
func Owner_Publish(mux http.ServeMux) {
	//No API
	//Cannot List via GUI
	//Cannot View via GUI
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Owner_Title)
	//No API
}

// Builds/Popuplates the Owner Page
func owner_PopulatePage(rD dm.Owner, pageDetail Owner_Page) Owner_Page {
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	pageDetail.UserName = rD.UserName
	pageDetail.FullName = rD.FullName
	pageDetail.Type = rD.Type
	pageDetail.TradingEntity = rD.TradingEntity
	pageDetail.DefaultEnterBook = rD.DefaultEnterBook
	pageDetail.EmailAddress = rD.EmailAddress
	pageDetail.Enabled = rD.Enabled
	pageDetail.ExternalUserIds = rD.ExternalUserIds
	pageDetail.Language = rD.Language
	pageDetail.LocalCurrency = rD.LocalCurrency
	pageDetail.Role = rD.Role
	pageDetail.TelephoneNumber = rD.TelephoneNumber
	pageDetail.TokenId = rD.TokenId
	pageDetail.Entity = rD.Entity
	pageDetail.UserCode = rD.UserCode

	//
	// Automatically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//

	//
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return pageDetail
}
