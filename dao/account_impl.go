package dao

import (
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// Account_GetListByCounterparty returns a list of accounts for a counterparty.
func Account_GetListByCounterparty(idFirm string, idCentre string) (int, []dm.Account, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	tsql = tsql + " WHERE " + dm.Account_Firm + "='" + idFirm + "' AND " + dm.Account_Centre + "='" + idCentre + "'"

	count, sienaAccountList, _, _ := account_Fetch(tsql)
	return count, sienaAccountList, nil
}

// Account_GetListByCounterparty returns a list of accounts for a counterparty.
func Account_GetListByCounterpartyID(id string) (int, []dm.Account, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	tsql = tsql + " WHERE " + dm.Account_CompID + "='" + id + "'"

	count, sienaAccountList, _, _ := account_Fetch(tsql)
	return count, sienaAccountList, nil
}

func account_DealtCA_Extra(recItem dm.Account) string {

	return core.FormatCurrencyDps(recItem.DealtAmount, recItem.CCY, recItem.CCYDp)
}

func account_AgainstCA_Extra(recItem dm.Account) string {
	return account_DealtCA_Extra(recItem)
}

func account_LedgerCA_Extra(recItem dm.Account) string {
	return core.FormatCurrencyDps(recItem.LedgerBalance, recItem.CCY, recItem.CCYDp)
}

func account_CashBalanceCA_Extra(recItem dm.Account) string {
	return core.FormatCurrencyDps(recItem.CashBalance, recItem.CCY, recItem.CCYDp)
}
