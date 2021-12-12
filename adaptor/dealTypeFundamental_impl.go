package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/dealtypefundamental.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : DealTypeFundamental (dealtypefundamental)
// Endpoint 	        : DealTypeFundamental (DealTypeKey)
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

func DealTypeFundamental_Delete_Impl(id string) error {
	var er error

	message := "Implement DealTypeFundamental_Delete: " + id

	// Implement DealTypeFundamental_Delete_Impl in dealtypefundamental_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := DealTypeFundamental_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func DealTypeFundamental_Update_Impl(item dm.DealTypeFundamental, usr string) error {
	var er error

	message := "Implement DealTypeFundamental_Update: " + item.DealTypeKey

	// Implement DealTypeFundamental_Update_Impl in dealtypefundamental_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := DealTypeFundamental_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
