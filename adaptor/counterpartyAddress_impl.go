package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/counterpartyaddress.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : CounterpartyAddress (counterpartyaddress)
// Endpoint 	        : CounterpartyAddress (ID)
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

func CounterpartyAddress_Delete_Impl(id string) error {
	var er error

	message := "Implement CounterpartyAddress_Delete: " + id

	// Implement CounterpartyAddress_Delete_Impl in counterpartyaddress_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CounterpartyAddress_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func CounterpartyAddress_Update_Impl(item dm.CounterpartyAddress, usr string) error {
	var er error

	message := "Implement CounterpartyAddress_Update: " + item.CompID

	// Implement CounterpartyAddress_Update_Impl in counterpartyaddress_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CounterpartyAddress_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
