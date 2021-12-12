package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/accountladder.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : AccountLadder (accountladder)
// Endpoint 	        : AccountLadder (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2021 at 15:59:13
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func AccountLadder_Delete_Impl(id string) error {
	var er error

	message := "Implement AccountLadder_Delete: " + id

	// Implement AccountLadder_Delete_Impl in accountladder_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := AccountLadder_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func AccountLadder_Update_Impl(item dm.AccountLadder, usr string) error {
	var er error

	message := "Implement AccountLadder_Update: " + item.SienaReference

	// Implement AccountLadder_Update_Impl in accountladder_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := AccountLadder_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
