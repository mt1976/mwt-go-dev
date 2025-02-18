package dao

import (
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// Account_GetListByCounterparty returns a list of accounts for a counterparty.
func Account_GetListByCounterparty(idFirm string, idCentre string) (int, []dm.Account, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	tsql = tsql + " WHERE " + dm.Account_Firm_sql + "='" + idFirm + "' AND " + dm.Account_Centre_sql + "='" + idCentre + "'"

	count, sienaAccountList, _, _ := account_Fetch(tsql)
	return count, sienaAccountList, nil
}

// Account_GetListByCounterparty returns a list of accounts for a counterparty.
func Account_GetListByCounterpartyID(id string) (int, []dm.Account, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	tsql = tsql + " WHERE " + dm.Account_CompID_sql + "='" + id + "'"

	count, sienaAccountList, _, _ := account_Fetch(tsql)
	return count, sienaAccountList, nil
}

func account_DealtCA_Extra(recItem dm.Account) string {

	return core.Financial_FormatAmountToDPS(recItem.DealtAmount, recItem.CCY, recItem.CCYDp)
}

func account_AgainstCA_Extra(recItem dm.Account) string {
	return account_DealtCA_Extra(recItem)
}

func account_LedgerCA_Extra(recItem dm.Account) string {
	return core.Financial_FormatAmountToDPS(recItem.LedgerBalance, recItem.CCY, recItem.CCYDp)
}

func account_CashBalanceCA_Extra(recItem dm.Account) string {
	return core.Financial_FormatAmountToDPS(recItem.CashBalance, recItem.CCY, recItem.CCYDp)
}

func account_DealtCA_Extra_Store(interface{}, interface{}) error {
	// Account_GetListByCounterparty returns a list of accounts for a counterparty.
	return nil
}

func account_CashBalanceCA_Extra_Store(interface{}, interface{}) error {
	// Account_GetListByCounterparty returns a list of accounts for a counterparty.
	return nil
}

func account_AgainstCA_Extra_Store(interface{}, interface{}) error {
	// Account_GetListByCounterparty returns a list of accounts for a counterparty.
	return nil
}

func account_LedgerCA_Extra_Store(interface{}, interface{}) error {
	// Account_GetListByCounterparty returns a list of accounts for a counterparty.
	return nil
}
