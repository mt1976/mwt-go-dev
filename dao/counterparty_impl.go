package dao

import (
	"fmt"
	"strings"

	"github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	"github.com/mt1976/mwt-go-dev/logs"
)

// CounterpartyType_GetLookup() returns a lookup list of all CounterpartyType items in lookup format
func CounterpartyType_GetLookup() []dm.Lookup_Item {
	return buildLookupList("counterpartytypes")
}

// Counterparty_MiFIDCategory_GetLookup() returns a lookup list of all CounterpartyType items in lookup format
func Counterparty_MiFIDCategory_GetLookup() []dm.Lookup_Item {
	return buildLookupList("mifidcategory")
}

// Counterparty_MiFIDCategory_GetLookup() returns a lookup list of all CounterpartyType items in lookup format
func Counterparty_Origin_GetLookup() []dm.Lookup_Item {
	return buildLookupList("originlist")
}

// Counterparty_MiFIDCategory_GetLookup() returns a lookup list of all CounterpartyType items in lookup format
func Counterparty_Owner_GetLookup() []dm.Lookup_Item {
	return buildLookupList("users")
}

// Counterparty_MiFIDCategory_GetLookup() returns a lookup list of all CounterpartyType items in lookup format
func Counterparty_SystemUser_GetLookup() []dm.Lookup_Item {
	return buildLookupList("users")
}

func buildLookupList(listName string) []dm.Lookup_Item {
	ctList := core.SienaProperties[listName]
	logs.Success("listName=" + listName)
	logs.Success("List=" + ctList)
	// Turn comma-separated string into a slice of strings
	ctSlice := strings.Split(ctList, ",")
	fmt.Printf("ctSlice: %v\n", ctSlice)
	// Create a lookup list of all CounterpartyType items
	ctLookup := make([]dm.Lookup_Item, len(ctSlice))
	for i := 0; i < len(ctSlice); i++ {
		ctLookup[i] = dm.Lookup_Item{ID: ctSlice[i], Name: ctSlice[i]}
	}
	fmt.Printf("ctLookup: %v\n", ctLookup)
	return ctLookup
}
