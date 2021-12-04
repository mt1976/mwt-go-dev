package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/counterpartycreditrating.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : CounterpartyCreditRating (counterpartycreditrating)
// Endpoint 	        : CounterpartyCreditRating (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 04/12/2021 at 17:36:43
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func CounterpartyCreditRating_Delete(id string) error {
	var er error

	message:= "Implement CounterpartyCreditRating_Delete: " + id

	// Implement CounterpartyCreditRating_Delete_Impl in counterpartycreditrating_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CounterpartyCreditRating_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func CounterpartyCreditRating_Update(item dm.CounterpartyCreditRating) error {
	var er error

	message:= "Implement CounterpartyCreditRating_Update: " + item.CompID

	// Implement CounterpartyCreditRating_Update_Impl in counterpartycreditrating_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CounterpartyCreditRating_Update_Impl(item)
	//

	logs.Success(message)
	return er
}