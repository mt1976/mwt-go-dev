package dao

import (
	"strings"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	"github.com/mt1976/mwt-go-dev/logs"
)

func payee_NewIDImpl(r dm.Payee) string { return "" }

// Payee_GetByID() returns a list of payees for a couterparty
func Payee_GetByCounterparty(firm string, centre string) (int, []dm.Payee, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Payee_SQLTable)
	tsql = tsql + " WHERE " + dm.Payee_KeyCounterpartyFirm + "='" + firm + "' AND " + dm.Payee_KeyCounterpartyCentre + "='" + centre + "'"

	noItems, payeeList, _, _ := payee_Fetch(tsql)
	return noItems, payeeList, nil
}

// Payee_GetByFullKey retrives a record from the payee table using the rediculous Siena key.
func Payee_GetByFullKey(ID_source string, ID_firm string, ID_centre string, ID_currency string, UD_name string, ID_number string, ID_direction string, ID_type string) (int, dm.Payee, error) {

	//tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterpartyPayee WHERE SourceTable='%s' AND KeyCounterpartyFirm='%s' AND KeyCounterpartyCentre='%s' AND KeyCurrency='%s' AND KeyName='%s' AND KeyNumber='%s' AND KeyDirection='%s' AND KeyType='%s' ORDER BY KeyCounterpartyFirm, KeyCounterpartyCentre, KeyCurrency DESC;", sienaCounterpartyPayeeSQL, core.SienaPropertiesDB["schema"], idSource, idFirm, idCentre, idCCY, idName, idNumber, idDirection, idType)

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Payee_SQLTable)
	tsql = tsql + " WHERE SourceTable=" + sq(ID_source)

	tsql = tsql_AND(tsql, dm.Payee_KeyCounterpartyFirm, ID_firm)
	tsql = tsql_AND(tsql, dm.Payee_KeyCounterpartyCentre, ID_centre)
	tsql = tsql_AND(tsql, dm.Payee_KeyCurrency, ID_currency)
	tsql = tsql_AND(tsql, dm.Payee_KeyName, UD_name)
	tsql = tsql_AND(tsql, dm.Payee_KeyNumber, ID_number)
	tsql = tsql_AND(tsql, dm.Payee_KeyDirection, ID_direction)
	//tsql = tsql_AND(tsql, dm.Payee_KeyType, ID_type)

	_, _, payeeItem, _ := payee_Fetch(tsql)
	return 1, payeeItem, nil
}

// Payee_GetListByCounterparty returns a list of accounts for a counterparty.
func Payee_GetListByCounterpartyID(id string) (int, []dm.Payee, error) {

	// tokenise the id to get firm and centre
	parts := strings.Split(id, "|")
	if len(parts) != 2 {
		logs.Warning("Invalid counterparty id:" + id)
		return 0, nil, nil
	}

	count, sienaPayeeList, er := Payee_GetByCounterparty(parts[0], parts[1])
	return count, sienaPayeeList, er
}

func payee_Status_Extra(r dm.Payee) string {

	val := ""
	if r.SourceTable == "" {
		logs.Warning("SourceTable is empty for " + r.FullName)
		return ""
	}

	splitSource := strings.Split(r.SourceTable, ".")

	if splitSource[1] == "Payee" {
		val = "Authorised"
	} else {
		val = "Pending"
	}

	return val
}

func payee_Reason_Override(r dm.Payee) string {
	return r.Reason
}
