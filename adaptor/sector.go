package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/sector.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Sector (sector)
// Endpoint 	        : Sector (Sector)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 22/11/2021 at 11:25:57
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func Sector_Delete(id string) error {
	var er error

	message:= "Implement Sector_Delete: " + id

	// Implement Sector_Delete_Impl in sector_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Sector_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Sector_Update(item dm.Sector) error {
	var er error

	message:= "Implement Sector_Update: " + item.Code

	// Implement Sector_Update_Impl in sector_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Sector_Update_Impl(item)
	//

	logs.Success(message)
	return er
}