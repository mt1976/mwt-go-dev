package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/currencypair.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : CurrencyPair (currencypair)
// Endpoint 	        : CurrencyPair (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2021 at 15:59:15
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func CurrencyPair_Delete_Impl(id string) error {
	var er error

	message := "Implement CurrencyPair_Delete: " + id

	// Implement CurrencyPair_Delete_Impl in currencypair_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CurrencyPair_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func CurrencyPair_Update_Impl(item dm.CurrencyPair, usr string) error {
	var er error

	message := "Implement CurrencyPair_Update: " + item.Code

	// Implement CurrencyPair_Update_Impl in currencypair_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CurrencyPair_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
