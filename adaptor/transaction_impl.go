package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/transaction.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Transaction (transaction)
// Endpoint 	        : Transaction (Ref)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2021 at 15:59:16
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Transaction_Delete_impl(id string) error {
	var er error

	message := "Implement Transaction_Delete: " + id

	// Implement Transaction_Delete_impl in transaction_impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Transaction_Delete_impl(item)
	//

	logs.Success(message)
	return er
}

func Transaction_Update_impl(id string, item dm.Transaction, usr string) error {
	var er error

	message := "Implement Transaction_Update: " + item.SienaReference

	// Implement Transaction_Update_impl in transaction_impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Transaction_Update_impl(item)
	//

	logs.Success(message)
	return er
}

func Transaction_Dealt_OnStore_impl(fieldval string, rec dm.Transaction, usr string) (string, error) {
	return fieldval, nil
}
func Transaction_Against_OnStore_impl(fieldval string, rec dm.Transaction, usr string) (string, error) {
	return fieldval, nil
}

func Transaction_Dealt_OnFetch_impl(rec dm.Transaction) string   { return "" }
func Transaction_Against_OnFetch_impl(rec dm.Transaction) string { return "" }

func Transaction_NewID_impl(rec dm.Transaction) string { return rec.SienaReference }

func Transaction_Dealt_Impl(iAction string, iId string, iValue string, iRec dm.Transaction, fP dm.FieldProperties) (string, dm.FieldProperties) {

	logs.Callout("Transaction", "Dealt", VAL+"-"+iAction, iId)

	return "", fP

}

// Transaction_Against_Impl provides validation/actions for Against

func Transaction_Against_Impl(iAction string, iId string, iValue string, iRec dm.Transaction, fP dm.FieldProperties) (string, dm.FieldProperties) {

	logs.Callout("Transaction", "Against", VAL+"-"+iAction, iId)

	return "", fP

}

//

// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local

// END - Validation API/Callout
