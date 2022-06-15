package dao

import (
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// CounterpartyType_GetLookup() returns a lookup list of all CounterpartyType items in lookup format
func CounterpartyType_GetLookup() []dm.Lookup_Item {
	return StubLists_Get("counterpartytypes")
}

// Counterparty_MiFIDCategory_GetLookup() returns a lookup list of all CounterpartyType items in lookup format
func Counterparty_MiFIDCategory_GetLookup() []dm.Lookup_Item {
	return StubLists_Get("mifidcategory")
}

// Counterparty_MiFIDCategory_GetLookup() returns a lookup list of all CounterpartyType items in lookup format
func Counterparty_Origin_GetLookup() []dm.Lookup_Item {
	return StubLists_Get("originlist")
}

// Counterparty_MiFIDCategory_GetLookup() returns a lookup list of all CounterpartyType items in lookup format
func Counterparty_Owner_GetLookup() []dm.Lookup_Item {
	return StubLists_Get("users")
}

// Counterparty_MiFIDCategory_GetLookup() returns a lookup list of all CounterpartyType items in lookup format
func Counterparty_SystemUser_GetLookup() []dm.Lookup_Item {
	return StubLists_Get("users")
}
