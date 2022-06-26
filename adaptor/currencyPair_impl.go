package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/currencypair.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : CurrencyPair (currencypair)
// Endpoint 	        : CurrencyPair (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2021 at 15:59:15
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

func CurrencyPair_Delete_impl(id string) error {
	var er error

	message := "Implement CurrencyPair_Delete: " + id

	// Implement CurrencyPair_Delete_impl in currencypair_impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CurrencyPair_Delete_impl(item)
	//

	logs.Success(message)
	return er
}

func CurrencyPair_Update_impl(id string, item dm.CurrencyPair, usr string) error {
	var er error

	message := "Implement CurrencyPair_Update: " + item.Code

	// Implement CurrencyPair_Update_impl in currencypair_impl.go
	// Uncomment the line below to use the implementation
	//
	// er := CurrencyPair_Update_impl(item)
	//

	logs.Success(message)
	return er
}

func CurrencyPair_Code_OnFetch_impl(rec dm.CurrencyPair) string { return rec.Code }

func CurrencyPair_Code_OnStore_impl(fieldval string, rec dm.CurrencyPair, usr string) (string, error) {
	return rec.CodeMajorCurrencyIsoCode + rec.CodeMinorCurrencyIsoCode, nil
}

func CurrencyPair_NewID_impl(rec dm.CurrencyPair) string { return rec.Code }

//
// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
//
// CurrencyPair_Code_Impl provides validation/actions for Code
func CurrencyPair_Code_Impl(iAction string, iId string, iValue string, iRec dm.CurrencyPair, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("CurrencyPair", "Code", VAL+"-"+iAction, iId)
	switch iAction {

	case VAL:

	case NEW:

	case PUT:
		// Validate Code
		if iRec.CodeMajorCurrencyIsoCode == iRec.CodeMinorCurrencyIsoCode {
			fP = dm.AddFieldMessage(fP, core.FieldMessage_ERROR, "Invalide Currency Pair - Currencies Match", false, "far fa-thumbs-down")
		}
	case GET:

	default:

		logs.Warning("Tmpl_TDate_Impl" + " - Invalid Action")

	}

	return iId, fP
}

//
// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout
