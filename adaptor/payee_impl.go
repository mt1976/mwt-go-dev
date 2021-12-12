package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/payee.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Payee (payee)
// Endpoint 	        : Payee (ID)
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

func Payee_Delete_Impl(id string) error {
	var er error

	message := "Implement Payee_Delete: " + id

	// Implement Payee_Delete_Impl in payee_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Payee_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Payee_Update_Impl(item dm.Payee, usr string) error {
	var er error

	message := "Implement Payee_Update: " + item.KeyCounterpartyFirm

	// Implement Payee_Update_Impl in payee_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Payee_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
