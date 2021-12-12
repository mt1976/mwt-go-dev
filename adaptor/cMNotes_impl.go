package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/cmnotes.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : CMNotes (cmnotes)
// Endpoint 	        : CMNotes (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2021 at 15:59:14
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func CMNotes_Delete_Impl(id string) error {
	var er error

	message := "Implement CMNotes_Delete: " + id

	// Implement CMNotes_Delete_Impl in cmnotes_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CMNotes_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func CMNotes_Update_Impl(item dm.CMNotes, usr string) error {
	var er error

	message := "Implement CMNotes_Update: " + item.NoteId

	// Implement CMNotes_Update_Impl in cmnotes_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CMNotes_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
