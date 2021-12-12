package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/counterpartyextensions.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : CounterpartyExtensions (counterpartyextensions)
// Endpoint 	        : CounterpartyExtensions (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2021 at 15:59:14
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func CounterpartyExtensions_Delete_Impl(id string) error {
	var er error

	message := "Implement CounterpartyExtensions_Delete: " + id

	// Implement CounterpartyExtensions_Delete_Impl in counterpartyextensions_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CounterpartyExtensions_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func CounterpartyExtensions_Update_Impl(item dm.CounterpartyExtensions, usr string) error {
	var er error

	message := "Implement CounterpartyExtensions_Update: " + item.CompID

	// Implement CounterpartyExtensions_Update_Impl in counterpartyextensions_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CounterpartyExtensions_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
