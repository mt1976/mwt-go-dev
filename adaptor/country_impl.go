package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/country.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Country (country)
// Endpoint 	        : Country (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2021 at 15:59:15
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Country_Delete_Impl(id string) error {
	var er error

	message := "Implement Country_Delete: " + id

	// Implement Country_Delete_Impl in country_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Country_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Country_Update_Impl(item dm.Country, usr string) error {
	var er error

	message := "Implement Country_Update: " + item.Code

	// Implement Country_Update_Impl in country_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Country_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
