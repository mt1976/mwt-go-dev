package adaptor

import (
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Counterparty_Update_Impl(item dm.Counterparty, usr string) error {

	message := "Implement Counterparty_Update: " + item.CompID

	// IGNORE
	var sienaKeyFields []StaticImport_KeyField
	var sienaFields []StaticImport_Field
	// ADD THE FIELDS TO THE LISTS ABOVE

	sienaKeyFields = StaticImport_AddKeyField(sienaKeyFields, dm.Counterparty_NameCentre, item.NameCentre)
	sienaKeyFields = StaticImport_AddKeyField(sienaKeyFields, dm.Counterparty_NameFirm, item.NameFirm)

	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_FullName, item.FullName)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_TelephoneNumber, item.TelephoneNumber)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_EmailAddress, item.EmailAddress)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_CustomerType, item.CustomerType)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_AccountOfficer, item.AccountOfficer)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_CountryCode, item.CountryCode)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_SectorCode, item.SectorCode)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_CpartyGroupName, item.CpartyGroupName)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_Notes, item.Notes)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_Owner, item.Owner)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_Authorised, item.Authorised)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_NameFirmName, item.NameFirmName)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_NameCentreName, item.NameCentreName)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_CountryCodeName, item.CountryCodeName)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_SectorCodeName, item.SectorCodeName)

	XMLmessage := StaticImport_Create(StaticImport_UpdateAction, "Counterparty", sienaKeyFields, sienaFields)

	err := StaticImport_DispatchXML(XMLmessage, StaticImport_UpdateAction)
	if err != nil {
		logs.Error("Error in StaticImport_Dispatch: ", err)
	} else {
		logs.Success(message)
	}
	return err
}

func Counterparty_Delete_Impl(id string) error {
	var er error

	message := "Implement Counterparty_Delete: " + id

	// Implement Counterparty_Delete_Impl in counterparty_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Counterparty_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}
