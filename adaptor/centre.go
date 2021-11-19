package adaptor

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func Centre_Delete(id string) error {
	var er error

	message:= "Implement Centre_Delete: " + id

	// Implement Centre_Delete_Impl in centre_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Centre_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Centre_Update(item dm.Centre) error {
	var er error

	message:= "Implement Centre_Update: " + item.Code

	// Implement Centre_Update_Impl in centre_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Centre_Update_Impl(item)
	//

	logs.Success(message)
	return er
}