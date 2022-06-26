package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/account.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Account (account)
// Endpoint 	        : Account (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2021 at 15:59:12
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Account_Delete_impl(id string) error {
	var er error

	message := "Implement Account_Delete: " + id

	// Implement Account_Delete_impl in account_impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Account_Delete_impl(item)
	//

	logs.Success(message)
	return er
}

func Account_Update_impl(id string, item dm.Account, usr string) error {
	var er error

	message := "Implement Account_Update: " + item.SienaReference + "" + usr

	// Implement Account_Update_impl in account_impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Account_Update_impl(item)
	//

	logs.Success(message)
	return er
}

func Account_NewID_impl(rec dm.Account) string { return rec.SienaReference }

// If there are fields below, create the methods in adaptor\account_impl.go

func Account_DealtCA_OnStore_impl(fieldval string, rec dm.Account, usr string) (string, error) {
	return fieldval, nil
}
func Account_AgainstCA_OnStore_impl(fieldval string, rec dm.Account, usr string) (string, error) {
	return fieldval, nil
}
func Account_LedgerCA_OnStore_impl(fieldval string, rec dm.Account, usr string) (string, error) {
	return fieldval, nil
}
func Account_CashBalanceCA_OnStore_impl(fieldval string, rec dm.Account, usr string) (string, error) {
	return fieldval, nil
}

func Account_DealtCA_OnFetch_impl(rec dm.Account) string {
	return core.Financial_FormatAmountToDPS(rec.DealtAmount, rec.CCY, rec.CCYDp)
}
func Account_AgainstCA_OnFetch_impl(rec dm.Account) string {
	return core.Financial_FormatAmountToDPS(rec.DealtAmount, rec.CCY, rec.CCYDp)
}
func Account_LedgerCA_OnFetch_impl(rec dm.Account) string {
	return core.Financial_FormatAmountToDPS(rec.LedgerBalance, rec.CCY, rec.CCYDp)
}
func Account_CashBalanceCA_OnFetch_impl(rec dm.Account) string {
	return core.Financial_FormatAmountToDPS(rec.CashBalance, rec.CCY, rec.CCYDp)
}

//
// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
//
// Account_DealtCA_Impl provides validation/actions for DealtCA
func Account_DealtCA_Impl(iAction string, iId string, iValue string, iRec dm.Account, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Account", "DealtCA", VAL+"-"+iAction, iId)
	return "", fP
}

// Account_AgainstCA_Impl provides validation/actions for AgainstCA
func Account_AgainstCA_Impl(iAction string, iId string, iValue string, iRec dm.Account, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Account", "AgainstCA", VAL+"-"+iAction, iId)
	return "", fP
}

// Account_LedgerCA_Impl provides validation/actions for LedgerCA
func Account_LedgerCA_Impl(iAction string, iId string, iValue string, iRec dm.Account, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Account", "LedgerCA", VAL+"-"+iAction, iId)
	return "", fP
}

// Account_CashBalanceCA_Impl provides validation/actions for CashBalanceCA
func Account_CashBalanceCA_Impl(iAction string, iId string, iValue string, iRec dm.Account, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Account", "CashBalanceCA", VAL+"-"+iAction, iId)
	return "", fP
}

//
// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout
