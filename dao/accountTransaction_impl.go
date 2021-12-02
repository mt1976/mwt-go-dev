package dao

import (
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// AccountTransaction_GetList() returns a list of all AccountTransaction records
func AccountTransaction_GetListByID(id string) (int, []dm.AccountTransaction, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.AccountTransaction_SQLTable)
	tsql += " WHERE " + dm.AccountTransaction_SienaReference + " = '" + id + "'"
	count, accounttransactionList, _, _ := accounttransaction_Fetch(tsql)
	return count, accounttransactionList, nil
}
