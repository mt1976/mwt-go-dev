package adaptor

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func Broker_Delete(id string) error {
	var er error

	message:= "Implement Broker_Delete: " + id

	// Implement Broker_Delete_Impl in broker_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Broker_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Broker_Update(item dm.Broker) error {
	var er error

	message:= "Implement Broker_Update: " + item.Code

	// Implement Broker_Update_Impl in broker_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Broker_Update_Impl(item)
	//

	logs.Success(message)
	return er
}