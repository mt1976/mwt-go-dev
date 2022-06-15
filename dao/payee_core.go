package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/payee.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Payee (payee)
// Endpoint 	        : Payee (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:32:07
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (

	"fmt"
	"net/http"
core "github.com/mt1976/mwt-go-dev/core"
"github.com/google/uuid"
das  "github.com/mt1976/mwt-go-dev/das"
	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
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


	adaptor.Payee_Delete_impl(id)
	
}


// Payee_Store() saves/stores a Payee record to the database
func Payee_Store(r dm.Payee,req *http.Request) error {

	err := payee_Save(r,Audit_User(req))

	return err
}

// Payee_StoreSystem() saves/stores a Payee record to the database
func Payee_StoreSystem(r dm.Payee) error {
	
	err := payee_Save(r,Audit_Host())

	return err
}

// payee_Save() saves/stores a Payee record to the database
func payee_Save(r dm.Payee,usr string) error {

    var err error



	

	if len(r.KeyCounterpartyFirm) == 0 {
		r.KeyCounterpartyFirm = Payee_NewID(r)
	}

// If there are fields below, create the methods in dao\payee_impl.go























  r.Status,err = adaptor.Payee_Status_OnStore_impl (r.Status,r,usr)


	
logs.Storing("Payee",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Payee_impl.go file
	err1 := adaptor.Payee_Delete_impl(r.KeyCounterpartyFirm)
	err2 := adaptor.Payee_Update_impl(r.KeyCounterpartyFirm,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// payee_Fetch read all Payee's
func payee_Fetch(tsql string) (int, []dm.Payee, dm.Payee, error) {

	var recItem dm.Payee
	var recList []dm.Payee

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - START
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

// If there are fields below, create the methods in adaptor\Payee_impl.go























   recItem.Status  = adaptor.Payee_Status_OnFetch_impl (recItem)

	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Payee_NewID(r dm.Payee) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

