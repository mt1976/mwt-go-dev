package dao

import (
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// AccountLadder_GetList() returns a list of all AccountLadder records
func AccountLadder_GetListByID(id string) (int, []dm.AccountLadder, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.AccountLadder_SQLTable)
	tsql += " WHERE " + dm.AccountLadder_SienaReference + " = '" + id + "'"
	count, accountladderList, _, _ := accountladder_Fetch(tsql)
	return count, accountladderList, nil
}
