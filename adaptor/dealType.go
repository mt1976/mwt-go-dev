package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/dealtype.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : DealType (dealtype)
// Endpoint 	        : DealType (DealTypeKey)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 04/12/2021 at 17:36:46
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func DealType_Delete(id string) error {
	var er error

	message:= "Implement DealType_Delete: " + id

	// Implement DealType_Delete_Impl in dealtype_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := DealType_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func DealType_Update(item dm.DealType) error {
	var er error

	message:= "Implement DealType_Update: " + item.DealTypeKey

	// Implement DealType_Update_Impl in dealtype_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := DealType_Update_Impl(item)
	//

	logs.Success(message)
	return er
}