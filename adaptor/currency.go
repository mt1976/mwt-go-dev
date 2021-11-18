package adaptor

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func Currency_Delete(id string) error {
	var er error

	message:= "Implement Currency_Delete: " + id

	// Implement Currency_Delete_Impl in currency_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Currency_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Currency_Update(item dm.Currency) error {
	var er error

	message:= "Implement Currency_Update: " + item.Code

	// Implement Currency_Update_Impl in currency_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Currency_Update_Impl(item)
	//

	logs.Success(message)
	return er
}