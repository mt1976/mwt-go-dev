package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/counterpartygroup.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : CounterpartyGroup (counterpartygroup)
// Endpoint 	        : CounterpartyGroup (Group)
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

func CounterpartyGroup_Delete_Impl(id string) error {
	var er error

	message := "Implement CounterpartyGroup_Delete: " + id

	// Implement CounterpartyGroup_Delete_Impl in counterpartygroup_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CounterpartyGroup_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func CounterpartyGroup_Update_Impl(item dm.CounterpartyGroup, usr string) error {
	var er error

	message := "Implement CounterpartyGroup_Update: " + item.Name

	// Implement CounterpartyGroup_Update_Impl in counterpartygroup_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CounterpartyGroup_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
