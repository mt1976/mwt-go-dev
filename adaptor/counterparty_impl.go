package adaptor

import (
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Counterparty_Update_Impl(item dm.Counterparty, usr string) error {

	//fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sienaFieldNameCentre StaticImport_KeyField
	var sienaFieldNameFirm StaticImport_KeyField
	var sienaFieldFullName StaticImport_Field
	var sienaFieldTelephoneNumber StaticImport_Field
	var sienaFieldEmailAddress StaticImport_Field
	var sienaFieldCustomerType StaticImport_Field
	var sienaFieldAccountOfficer StaticImport_Field
	var sienaFieldCountryCode StaticImport_Field
	var sienaFieldSectorCode StaticImport_Field
	var sienaFieldCpartyGroupName StaticImport_Field
	var sienaFieldNotes StaticImport_Field
	var sienaFieldOwner StaticImport_Field
	var sienaFieldAuthorised StaticImport_Field
	var sienaFieldNameFirmName StaticImport_Field
	var sienaFieldNameCentreName StaticImport_Field
	var sienaFieldCountryCodeName StaticImport_Field
	var sienaFieldSectorCodeName StaticImport_Field

	// POPULATE THE XML FIELDS
	sienaFieldNameCentre.Name = "NameCentre"
	sienaFieldNameFirm.Name = "NameFirm"
	sienaFieldFullName.Name = "FullName"
	sienaFieldTelephoneNumber.Name = "TelephoneNumber"
	sienaFieldEmailAddress.Name = "EmailAddress"
	sienaFieldCustomerType.Name = "CustomerType"
	sienaFieldAccountOfficer.Name = "AccountOfficer"
	sienaFieldCountryCode.Name = "CountryCode"
	sienaFieldSectorCode.Name = "SectorCode"
	sienaFieldCpartyGroupName.Name = "CpartyGroupName"
	sienaFieldNotes.Name = "Notes"
	sienaFieldOwner.Name = "Owner"
	sienaFieldAuthorised.Name = "Authorised"
	sienaFieldNameFirmName.Name = "NameFirmName"
	sienaFieldNameCentreName.Name = "NameCentreName"
	sienaFieldCountryCodeName.Name = "CountryCodeName"
	sienaFieldSectorCodeName.Name = "SectorCodeName"

	// POPULATE THE XML values

	sienaFieldNameCentre.Text = item.NameCentre
	sienaFieldNameFirm.Text = item.NameFirm
	sienaFieldFullName.Text = item.FullName
	sienaFieldTelephoneNumber.Text = item.TelephoneNumber
	sienaFieldEmailAddress.Text = item.EmailAddress
	sienaFieldCustomerType.Text = item.CustomerType
	sienaFieldAccountOfficer.Text = item.AccountOfficer
	sienaFieldCountryCode.Text = item.CountryCode
	sienaFieldSectorCode.Text = item.SectorCode
	sienaFieldCpartyGroupName.Text = item.CpartyGroupName
	sienaFieldNotes.Text = item.Notes
	sienaFieldOwner.Text = item.Owner
	sienaFieldAuthorised.Text = item.Authorised
	sienaFieldNameFirmName.Text = item.NameFirmName
	sienaFieldNameCentreName.Text = item.NameCentreName
	sienaFieldCountryCodeName.Text = item.CountryCodeName
	sienaFieldSectorCodeName.Text = item.SectorCodeName

	// IGNORE
	var sienaKeyFields []StaticImport_KeyField
	var sienaFields []StaticImport_Field
	// ADD THE FIELDS TO THE LISTS ABOVE

	sienaKeyFields = append(sienaKeyFields, sienaFieldNameCentre)
	sienaKeyFields = append(sienaKeyFields, sienaFieldNameFirm)
	sienaFields = append(sienaFields, sienaFieldFullName)
	sienaFields = append(sienaFields, sienaFieldTelephoneNumber)
	sienaFields = append(sienaFields, sienaFieldEmailAddress)
	sienaFields = append(sienaFields, sienaFieldCustomerType)
	sienaFields = append(sienaFields, sienaFieldAccountOfficer)
	sienaFields = append(sienaFields, sienaFieldCountryCode)
	sienaFields = append(sienaFields, sienaFieldSectorCode)
	sienaFields = append(sienaFields, sienaFieldCpartyGroupName)
	sienaFields = append(sienaFields, sienaFieldNotes)
	sienaFields = append(sienaFields, sienaFieldOwner)
	sienaFields = append(sienaFields, sienaFieldAuthorised)
	sienaFields = append(sienaFields, sienaFieldNameFirmName)
	sienaFields = append(sienaFields, sienaFieldNameCentreName)
	sienaFields = append(sienaFields, sienaFieldCountryCodeName)
	sienaFields = append(sienaFields, sienaFieldSectorCodeName)

	// IGNORE
	sienaRecord := &StaticImport_Record{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []StaticImport_Record
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable StaticImport_Table
	sienaTable.Name = "Counterparty"
	sienaTable.Classname = "com.eurobase.siena.data.Counterpartys.Counterparty"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction StaticImport_Transaction
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var StaticImport_XMLContent StaticImport_XML
	StaticImport_XMLContent.TRANSACTIONS = sienaTransaction
	return nil
}

func Counterparty_Delete_Impl(id string, usr string) error {
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