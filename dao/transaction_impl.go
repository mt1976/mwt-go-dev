package dao

import (
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// Transaction_GetListByCounterparty returns a list of accounts for a counterparty.
func Transaction_GetListByCounterpartyID(id string) (int, []dm.Transaction, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Transaction_SQLTable)
	tsql = tsql + " WHERE " + dm.Transaction_CompID + "='" + id + "'"

	count, sienaTransactionList, _, _ := transaction_Fetch(tsql)
	return count, sienaTransactionList, nil
}
