package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/currency.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Currency (currency)
// Endpoint 	        : Currency (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 04/12/2021 at 17:36:45
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func Currency_Delete(id string) error {
	var er error

	message:= "Implement Currency_Delete: " + id

	// Implement Currency_Delete_Impl in currency_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Currency_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Currency_Update(item dm.Currency) error {
	var er error

	message:= "Implement Currency_Update: " + item.Code

	// Implement Currency_Update_Impl in currency_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Currency_Update_Impl(item)
	//

	logs.Success(message)
	return er
}