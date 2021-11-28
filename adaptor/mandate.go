package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/mandate.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Mandate (mandate)
// Endpoint 	        : Mandate (Mandate)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 28/11/2021 at 22:54:58
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func Mandate_Delete(id string) error {
	var er error

	message:= "Implement Mandate_Delete: " + id

	// Implement Mandate_Delete_Impl in mandate_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Mandate_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Mandate_Update(item dm.Mandate) error {
	var er error

	message:= "Implement Mandate_Update: " + item.FirmName

	// Implement Mandate_Update_Impl in mandate_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Mandate_Update_Impl(item)
	//

	logs.Success(message)
	return er
}