package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/payee.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Payee (payee)
// Endpoint 	        : Payee (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 21:30:57
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func Payee_Delete(id string) error {
	var er error

	message:= "Implement Payee_Delete: " + id

	// Implement Payee_Delete_Impl in payee_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Payee_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Payee_Update(item dm.Payee) error {
	var er error

	message:= "Implement Payee_Update: " + item.KeyCounterpartyFirm

	// Implement Payee_Update_Impl in payee_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Payee_Update_Impl(item)
	//

	logs.Success(message)
	return er
}