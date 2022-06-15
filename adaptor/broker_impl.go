package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/broker.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Broker (broker)
// Endpoint 	        : Broker (Code)
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

func Broker_Delete_impl(id string) error {
	var er error

	message := "Implement Broker_Delete: " + id

	// Implement Broker_Delete_impl in broker_impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Broker_Delete_impl(item)
	//

	logs.Success(message)
	return er
}

func Broker_Update_impl(id string, item dm.Broker, usr string) error {

	message := "Implement Broker_Update: " + item.Code + id

	// Implement Broker_Update_impl in broker_impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Broker_Update_impl(item)
	//

	var sienaKeys []StaticImport_KeyField
	var sienaFlds []StaticImport_Field

	// POPULATE THE XML FIELDS

	sienaKeys = StaticImport_AddKeyField(sienaKeys, "Code", item.Code)
	sienaFlds = StaticImport_AddField(sienaFlds, "Name", item.Name)
	sienaFlds = StaticImport_AddField(sienaFlds, "Fullname", item.FullName)
	sienaFlds = StaticImport_AddField(sienaFlds, "Contact", item.Contact)
	sienaFlds = StaticImport_AddField(sienaFlds, "Address", item.Address)
	sienaFlds = StaticImport_AddField(sienaFlds, "LEI", item.LEI)

	XMLmessage := StaticImport_Create(StaticImport_UpdateAction, "Broker", sienaKeys, sienaFlds)

	err := StaticImport_DispatchXML(XMLmessage, StaticImport_UpdateAction)
	if err != nil {
		logs.Error("Error in StaticImport_Dispatch: ", err)
	} else {
		logs.Success(message)
	}
	return nil
}

func Broker_NewID_impl(rec dm.Broker) string { return rec.Code }
