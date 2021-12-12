package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/account.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Account (account)
// Endpoint 	        : Account (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2021 at 15:59:12
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Account_Delete_Impl(id string) error {
	var er error

	message := "Implement Account_Delete: " + id

	// Implement Account_Delete_Impl in account_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Account_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Account_Update_Impl(item dm.Account, usr string) error {
	var er error

	message := "Implement Account_Update: " + item.SienaReference + "" + usr

	// Implement Account_Update_Impl in account_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Account_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
