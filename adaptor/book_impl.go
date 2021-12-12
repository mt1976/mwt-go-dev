package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/book.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Book (book)
// Endpoint 	        : Book (Book)
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

func Book_Delete_Impl(id string) error {
	var er error

	message := "Implement Book_Delete: " + id

	// Implement Book_Delete_Impl in book_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Book_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Book_Update_Impl(item dm.Book, use string) error {
	var er error

	message := "Implement Book_Update: " + item.BookName

	// Implement Book_Update_Impl in book_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Book_Update_Impl(item)
	//

	logs.Success(message)
	return er
}
