package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/counterpartygroup.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : CounterpartyGroup (counterpartygroup)
// Endpoint 	        : CounterpartyGroup (Group)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 15:44:01
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func CounterpartyGroup_Delete(id string) error {
	var er error

	message:= "Implement CounterpartyGroup_Delete: " + id

	// Implement CounterpartyGroup_Delete_Impl in counterpartygroup_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CounterpartyGroup_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func CounterpartyGroup_Update(item dm.CounterpartyGroup) error {
	var er error

	message:= "Implement CounterpartyGroup_Update: " + item.Name

	// Implement CounterpartyGroup_Update_Impl in counterpartygroup_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CounterpartyGroup_Update_Impl(item)
	//

	logs.Success(message)
	return er
}