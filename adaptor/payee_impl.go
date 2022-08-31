package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/payee.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Payee (payee)
// Endpoint 	        : Payee (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2021 at 15:59:17
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	"strings"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Payee_Delete_impl(id string) error {
	var er error

	message := "Implement Payee_Delete: " + id

	// Implement Payee_Delete_impl in payee_impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Payee_Delete_impl(item)
	//

	logs.Success(message)
	return er
}

func Payee_Update_impl(id string, item dm.Payee, usr string) error {
	var er error

	message := "Implement Payee_Update: " + item.KeyCounterpartyFirm

	// Implement Payee_Update_impl in payee_impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Payee_Update_impl(item)
	//

	logs.Success(message)
	return er
}

func Payee_NewID_impl(rec dm.Payee) string { return rec.KeyCounterpartyFirm }

// If there are fields below, create the methods in adaptor\payee_impl.go

func Payee_Status_OnStore_impl(fieldval string, rec dm.Payee, usr string) (string, error) {
	return fieldval, nil
}

func Payee_Status_OnFetch_impl(rec dm.Payee) string {
	val := ""

	if rec.SourceTable == "" {

		logs.Warning("SourceTable is empty for " + rec.FullName)

		return ""

	}

	splitSource := strings.Split(rec.SourceTable, ".")

	if splitSource[1] == "Payee" {

		val = "Authorised"

	} else {

		val = "Pending"

	}

	return val
}

// Payee_Status_Impl provides validation/actions for Status

func Payee_Status_Impl(iAction string, iId string, iValue string, iRec dm.Payee, fP dm.FieldProperties) (string, dm.FieldProperties) {

	logs.Callout("Payee", "Status", VAL+"-"+iAction, iId)

	return "", fP

}

//

// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local

// END - Validation API/Callout

// If there are fields below, create the methods in adaptor\payee_impl.go
// START - GET API/Callout
// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
//

//
// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
//
// Payee_Status_impl provides validation/actions for Status
func Payee_Status_impl(iAction string, iId string, iValue string, iRec dm.Payee, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Payee", "Status", VAL+"-"+iAction, iId)
	return iValue, fP
}

//
// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func Payee_ObjectValidation_impl(iAction string, iId string, iRec dm.Payee) (dm.Payee, error, string) {
	logs.Callout("Payee", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("Tmpl_TDate_impl" + " - Invalid Action")
	}
	return iRec, nil, ""
}
