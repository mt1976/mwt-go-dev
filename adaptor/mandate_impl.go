package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/mandate.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Mandate (mandate)
// Endpoint 	        : Mandate (Mandate)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2021 at 15:59:17
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Mandate_Delete_Impl(id string) error {
	var er error

	message := "Implement Mandate_Delete: " + id

	// Implement Mandate_Delete_Impl in mandate_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Mandate_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Mandate_Update_Impl(item dm.Mandate, usr string) error {
	var er error

	message := "Implement Mandate_Update: " + item.CompID

	// Implement Mandate_Update_Impl in mandate_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Mandate_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
