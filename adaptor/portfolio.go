package adaptor

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func Portfolio_Delete(id string) error {
	var er error

	message:= "Implement Portfolio_Delete: " + id

	// Implement Portfolio_Delete_Impl in portfolio_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Portfolio_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Portfolio_Update(item dm.Portfolio) error {
	var er error

	message:= "Implement Portfolio_Update: " + item.Code

	// Implement Portfolio_Update_Impl in portfolio_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Portfolio_Update_Impl(item)
	//

	logs.Success(message)
	return er
}