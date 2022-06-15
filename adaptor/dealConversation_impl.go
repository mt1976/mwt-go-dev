package adaptor

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func DealConversation_Delete_impl(id string) error {

	message := "Implement DealConversation_Update_impl: " + id

	// Implement CurrencyPair_Delete_impl in currencypair_impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CurrencyPair_Delete_impl(item)
	//

	logs.Success(message)
	return nil
}

func DealConversation_Update_impl(id string, item dm.DealConversation, usr string) error {

	message := "Implement DealConversation_Update_impl: " + item.SienaReference

	// Implement CurrencyPair_Delete_impl in currencypair_impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CurrencyPair_Delete_impl(item)
	//

	logs.Success(message)
	return nil
}
