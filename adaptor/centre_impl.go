package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/centre.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Centre (centre)
// Endpoint 	        : Centre (Code)
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

func Centre_Delete_Impl(id string) error {
	var er error

	message := "Implement Centre_Delete: " + id

	// Implement Centre_Delete_Impl in centre_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Centre_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Centre_Update_Impl(item dm.Centre, usr string) error {
	var er error

	message := "Implement Centre_Update: " + item.Code

	// Implement Centre_Update_Impl in centre_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Centre_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
