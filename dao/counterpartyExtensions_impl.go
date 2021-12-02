package dao

import (
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// CounterpartyExtensions_GetByID() returns a single CounterpartyExtensions record
func CounterpartyExtensions_GetListByID(id string) (int, []dm.CounterpartyExtensions, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyExtensions_SQLTable)
	tsql = tsql + " WHERE " + dm.CounterpartyExtensions_SQLSearchID + "='" + id + "'"

	_, itemList, _, _ := counterpartyextensions_Fetch(tsql)
	return 1, itemList, nil
}
