package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/firm.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Firm (firm)
// Endpoint 	        : Firm (FirmName)
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

func Firm_Delete_Impl(id string) error {
	var er error

	message := "Implement Firm_Delete: " + id

	// Implement Firm_Delete_Impl in firm_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Firm_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Firm_Update_Impl(item dm.Firm, usr string) error {
	var er error

	message := "Implement Firm_Update: " + item.FirmName

	// Implement Firm_Update_Impl in firm_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Firm_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
