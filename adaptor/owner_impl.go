package adaptor

import (
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Owner_Update_Impl(item dm.Owner, usr string) error {

	//fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE

	return nil
}

func Owner_Delete_Impl(id string) error {
	var er error

	message := "Implement Owner_Delete: " + id

	// Implement Owner_Delete_Impl in counterparty_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Owner_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}
