package adaptor

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func Book_Delete(id string) error {
	var er error

	message:= "Implement Book_Delete: " + id

	// Implement Book_Delete_Impl in book_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Book_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Book_Update(item dm.Book) error {
	var er error

	message:= "Implement Book_Update: " + item.BookName

	// Implement Book_Update_Impl in book_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Book_Update_Impl(item)
	//

	logs.Success(message)
	return er
}