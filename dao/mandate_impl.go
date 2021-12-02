package dao

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

func mandate_NewIDImpl(r dm.Mandate) string {
	return uuid.New().String()
}

func Mandate_GetListByID(id string) (int, []dm.Mandate, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Mandate_SQLTable)
	tsql = tsql + " WHERE " + dm.Mandate_SQLSearchID + "='" + id + "'"
	count, counterpartyimportList, _, _ := mandate_Fetch(tsql)
	return count, counterpartyimportList, nil
}

func Mandate_GetListByCounterparty(idFirm string) (int, []dm.Mandate, error) {

	parts := strings.Split(idFirm, "|")
	fmt.Printf("parts: %v\n", parts)
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Mandate_SQLTable)
	tsql = tsql + " WHERE " + dm.Mandate_MandatedUserKeyCounterpartyFirm + "='" + parts[0] + "'"
	tsql = tsql + " AND " + dm.Mandate_MandatedUserKeyCounterpartyCentre + "='" + parts[1] + "'"
	fmt.Printf("tsql: %v\n", tsql)
	count, returnList, _, _ := mandate_Fetch(tsql)
	return count, returnList, nil
}

func Mandate_GetByFullID(idMu string, idFirm string, idCentre string) (int, dm.Mandate, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Mandate_SQLTable)
	tsql = tsql + " WHERE " + dm.Mandate_MandatedUserKeyCounterpartyFirm + "='" + idFirm + "'"
	tsql = tsql + " AND " + dm.Mandate_MandatedUserKeyCounterpartyCentre + "='" + idCentre + "'"
	tsql = tsql + " AND " + dm.Mandate_MandatedUserKeyUserName + "='" + idMu + "'"

	fmt.Printf("tsql: %v\n", tsql)
	count, _, returnList, _ := mandate_Fetch(tsql)
	return count, returnList, nil
}
