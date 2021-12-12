package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/portfolio.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Portfolio (portfolio)
// Endpoint 	        : Portfolio (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2021 at 15:59:18
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Portfolio_Delete_Impl(id string) error {
	var er error

	message := "Implement Portfolio_Delete: " + id

	// Implement Portfolio_Delete_Impl in portfolio_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Portfolio_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Portfolio_Update_Impl(item dm.Portfolio, usr string) error {
	var er error

	message := "Implement Portfolio_Update: " + item.Code

	// Implement Portfolio_Update_Impl in portfolio_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Portfolio_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
