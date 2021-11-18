package adaptor

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