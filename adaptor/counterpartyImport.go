package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/counterpartyimport.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : CounterpartyImport (counterpartyimport)
// Endpoint 	        : CounterpartyImport (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 02/12/2021 at 19:40:02
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func CounterpartyImport_Delete(id string) error {
	var er error

	message:= "Implement CounterpartyImport_Delete: " + id

	// Implement CounterpartyImport_Delete_Impl in counterpartyimport_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CounterpartyImport_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func CounterpartyImport_Update(item dm.CounterpartyImport) error {
	var er error

	message:= "Implement CounterpartyImport_Update: " + item.CompID

	// Implement CounterpartyImport_Update_Impl in counterpartyimport_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CounterpartyImport_Update_Impl(item)
	//

	logs.Success(message)
	return er
}