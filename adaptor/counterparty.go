package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/counterparty.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Counterparty (counterparty)
// Endpoint 	        : Counterparty (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 02/12/2021 at 19:40:00
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func Counterparty_Delete(id string) error {
	var er error

	message:= "Implement Counterparty_Delete: " + id

	// Implement Counterparty_Delete_Impl in counterparty_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Counterparty_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Counterparty_Update(item dm.Counterparty) error {
	var er error

	message:= "Implement Counterparty_Update: " + item.CompID

	// Implement Counterparty_Update_Impl in counterparty_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Counterparty_Update_Impl(item)
	//

	logs.Success(message)
	return er
}