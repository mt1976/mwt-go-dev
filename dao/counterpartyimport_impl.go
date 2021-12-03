package dao

import (
	"strings"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

func CounterpartyImport_GetListByID(id string) (int, []dm.CounterpartyImport, error) {

	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyImport_SQLTable)
	tsql = tsql + " WHERE " + dm.CounterpartyImport_SQLSearchID + "='" + id + "'"
	count, counterpartyimportList, _, _ := counterpartyimport_Fetch(tsql)
	return count, counterpartyimportList, nil
}

func CounterpartyImport_GetListByCounterparty(id string) (int, []dm.CounterpartyImport, error) {
	
	parts := strings.Split(id,"|")
	idFirm := parts[0]
	idCentre := parts[1] 
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyImport_SQLTable)
	tsql = tsql + " WHERE " + dm.CounterpartyImport_Firm + "='" + idFirm + "'"
	tsql = tsql + " AND " + dm.CounterpartyImport_Centre + "='" + idCentre + "'"

	count, counterpartyimportList, _, _ := counterpartyimport_Fetch(tsql)
	return count, counterpartyimportList, nil
}
