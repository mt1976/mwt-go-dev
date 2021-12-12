package adaptor

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func DealConversation_Delete_Impl(id string) error {

	message := "Implement DealConversation_Update_Impl: " + id

	// Implement CurrencyPair_Delete_Impl in currencypair_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CurrencyPair_Delete_Impl(item)
	//

	logs.Success(message)
	return nil
}

func DealConversation_Update_Impl(item dm.DealConversation, usr string) error {

	message := "Implement DealConversation_Update_Impl: " + item.SienaReference

	// Implement CurrencyPair_Delete_Impl in currencypair_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CurrencyPair_Delete_Impl(item)
	//

	logs.Success(message)
	return nil
}
