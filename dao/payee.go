package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/payee.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Payee (payee)
// Endpoint 	        : Payee (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 22/11/2021 at 21:11:41
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"log"
	"fmt"

	
	core "github.com/mt1976/mwt-go-dev/core"
	das  "github.com/mt1976/mwt-go-dev/das"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
)

// Payee_GetList() returns a list of all Payee records
func Payee_GetList() (int, []dm.Payee, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Payee_SQLTable)
	count, payeeList, _, _ := payee_Fetch(tsql)
	return count, payeeList, nil
}

// Payee_GetByID() returns a single Payee record
func Payee_GetByID(id string) (int, dm.Payee, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Payee_SQLTable)
	tsql = tsql + " WHERE " + dm.Payee_SQLSearchID + "='" + id + "'"

	_, _, payeeItem, _ := payee_Fetch(tsql)
	return 1, payeeItem, nil
}



// Payee_DeleteByID() deletes a single Payee record
func Payee_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Payee_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Payee_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}

// Payee_Store() saves/stores a Payee record to the database
func Payee_Store(r dm.Payee) error {

	logs.Storing("Payee",fmt.Sprintf("%s", r))

	if len(r.KeyCounterpartyFirm) == 0 {
		r.KeyCounterpartyFirm = Payee_NewID(r)
	}



	adaptor.Payee_Delete(r.KeyCounterpartyFirm)
	adaptor.Payee_Update(r)


	return nil
}

// payee_Fetch read all employees
func payee_Fetch(tsql string) (int, []dm.Payee, dm.Payee, error) {

	var recItem dm.Payee
	var recList []dm.Payee

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
    recItem.AppInternalID = get_String(rec, dm.Payee_KeyCounterpartyFirm,"")
   recItem.SourceTable  = get_String(rec, dm.Payee_SourceTable, "")
   recItem.KeyCounterpartyFirm  = get_String(rec, dm.Payee_KeyCounterpartyFirm, "")
   recItem.KeyCounterpartyCentre  = get_String(rec, dm.Payee_KeyCounterpartyCentre, "")
   recItem.KeyCurrency  = get_String(rec, dm.Payee_KeyCurrency, "")
   recItem.KeyName  = get_String(rec, dm.Payee_KeyName, "")
   recItem.KeyNumber  = get_String(rec, dm.Payee_KeyNumber, "")
   recItem.KeyDirection  = get_String(rec, dm.Payee_KeyDirection, "")
   recItem.KeyType  = get_String(rec, dm.Payee_KeyType, "")
   recItem.FullName  = get_String(rec, dm.Payee_FullName, "")
   recItem.Address  = get_String(rec, dm.Payee_Address, "")
   recItem.PhoneNo  = get_String(rec, dm.Payee_PhoneNo, "")
   recItem.Country  = get_String(rec, dm.Payee_Country, "")
   recItem.Bic  = get_String(rec, dm.Payee_Bic, "")
   recItem.Iban  = get_String(rec, dm.Payee_Iban, "")
   recItem.AccountNo  = get_String(rec, dm.Payee_AccountNo, "")
   recItem.FedWireNo  = get_String(rec, dm.Payee_FedWireNo, "")
   recItem.SortCode  = get_String(rec, dm.Payee_SortCode, "")
   recItem.BankName  = get_String(rec, dm.Payee_BankName, "")
   recItem.BankPinCode  = get_String(rec, dm.Payee_BankPinCode, "")
   recItem.BankAddress  = get_String(rec, dm.Payee_BankAddress, "")
   recItem.Reason  = get_String(rec, dm.Payee_Reason, "")
   recItem.BankSettlementAcct  = get_Bool(rec, dm.Payee_BankSettlementAcct, "True")
   recItem.UpdatedUserId  = get_String(rec, dm.Payee_UpdatedUserId, "")




// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Payee_NewID(r dm.Payee) string {
	
	

		// payee_NewIDImpl should be specified in dao/Payee_Impl.go
		// to provide the implementation for the special case.
		// override should return id - override function should be defined as
		// payee_NewIDImpl(r dm.Payee) string {...}
		//
		id := payee_NewIDImpl(r)
	
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

