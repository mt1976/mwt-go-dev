package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/counterpartyextensions.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : CounterpartyExtensions (counterpartyextensions)
// Endpoint 	        : CounterpartyExtensions (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 04/12/2021 at 17:36:44
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func CounterpartyExtensions_Delete(id string) error {
	var er error

	message:= "Implement CounterpartyExtensions_Delete: " + id

	// Implement CounterpartyExtensions_Delete_Impl in counterpartyextensions_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CounterpartyExtensions_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func CounterpartyExtensions_Update(item dm.CounterpartyExtensions) error {
	var er error

	message:= "Implement CounterpartyExtensions_Update: " + item.CompID

	// Implement CounterpartyExtensions_Update_Impl in counterpartyextensions_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CounterpartyExtensions_Update_Impl(item)
	//

	logs.Success(message)
	return er
}