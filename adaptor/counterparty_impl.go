package adaptor

import (
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Counterparty_Update_impl(id string, item dm.Counterparty, usr string) error {

	message := "Implement Counterparty_Update: " + item.CompID

	// IGNORE
	var sienaKeyFields []StaticImport_KeyField
	var sienaFields []StaticImport_Field
	// ADD THE FIELDS TO THE LISTS ABOVE

	sienaKeyFields = StaticImport_AddKeyField(sienaKeyFields, dm.Counterparty_NameCentre_sql, item.NameCentre)
	sienaKeyFields = StaticImport_AddKeyField(sienaKeyFields, dm.Counterparty_NameFirm_sql, item.NameFirm)

	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_FullName_sql, item.FullName)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_TelephoneNumber_sql, item.TelephoneNumber)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_EmailAddress_sql, item.EmailAddress)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_CustomerType_sql, item.CustomerType)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_AccountOfficer_sql, item.AccountOfficer)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_CountryCode_sql, item.CountryCode)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_SectorCode_sql, item.SectorCode)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_CpartyGroupName_sql, item.CpartyGroupName)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_Notes_sql, item.Notes)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_Owner_sql, item.Owner)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_Authorised_sql, item.Authorised)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_NameFirmName_sql, item.NameFirmName)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_NameCentreName_sql, item.NameCentreName)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_CountryCodeName_sql, item.CountryCodeName)
	sienaFields = StaticImport_AddField(sienaFields, dm.Counterparty_SectorCodeName_sql, item.SectorCodeName)

	XMLmessage := StaticImport_Create(StaticImport_UpdateAction, "Counterparty", sienaKeyFields, sienaFields)

	err := StaticImport_DispatchXML(XMLmessage, StaticImport_UpdateAction)
	if err != nil {
		logs.Error("Error in StaticImport_Dispatch: ", err)
	} else {
		logs.Success(message)
	}
	return err
}

func Counterparty_Delete_impl(id string) error {
	var er error

	message := "Implement Counterparty_Delete: " + id

	// Implement Counterparty_Delete_impl in counterparty_impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Counterparty_Delete_impl(item)
	//

	logs.Success(message)
	return er
}
