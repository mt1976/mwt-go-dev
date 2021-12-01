package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/cmnotes.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : CMNotes (cmnotes)
// Endpoint 	        : CMNotes (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 01/12/2021 at 20:36:38
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func CMNotes_Delete(id string) error {
	var er error

	message:= "Implement CMNotes_Delete: " + id

	// Implement CMNotes_Delete_Impl in cmnotes_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CMNotes_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func CMNotes_Update(item dm.CMNotes) error {
	var er error

	message:= "Implement CMNotes_Update: " + item.NoteId

	// Implement CMNotes_Update_Impl in cmnotes_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CMNotes_Update_Impl(item)
	//

	logs.Success(message)
	return er
}