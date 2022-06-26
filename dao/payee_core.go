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
// Date & Time		    : 26/06/2022 at 18:48:31
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	adaptor "github.com/mt1976/mwt-go-dev/adaptor"
	core "github.com/mt1976/mwt-go-dev/core"
	das "github.com/mt1976/mwt-go-dev/das"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
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

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	payeeItem.Status, payeeItem.Status_props = adaptor.Payee_Status_Impl(adaptor.GET, id, payeeItem.Status, payeeItem, payeeItem.Status_props)
	//
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, payeeItem, nil
}

// Payee_DeleteByID() deletes a single Payee record
func Payee_Delete(id string) {

	adaptor.Payee_Delete_impl(id)

}

// Payee_Store() saves/stores a Payee record to the database
func Payee_Store(r dm.Payee, req *http.Request) error {

	err, r := Payee_Validate(r)
	if err == nil {
		err = payee_Save(r, Audit_User(req))
	} else {
		logs.Information("Payee_Store()", err.Error())
	}

	return err
}

// Payee_StoreSystem() saves/stores a Payee record to the database
func Payee_StoreSystem(r dm.Payee) error {

	err, r := Payee_Validate(r)
	if err == nil {
		err = payee_Save(r, Audit_Host())
	} else {
		logs.Information("Payee_Store()", err.Error())
	}

	return err
}

// Payee_Validate() validates for saves/stores a Payee record to the database
func Payee_Validate(r dm.Payee) (error, dm.Payee) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	r.Status, r.Status_props = adaptor.Payee_Status_Impl(adaptor.PUT, r.KeyCounterpartyFirm, r.Status, r, r.Status_props)
	//
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return nil, r
}

//

// payee_Save() saves/stores a Payee record to the database
func payee_Save(r dm.Payee, usr string) error {

	var err error

	if len(r.KeyCounterpartyFirm) == 0 {
		r.KeyCounterpartyFirm = Payee_NewID(r)
	}

	// If there are fields below, create the methods in dao\payee_impl.go

	r.Status, err = adaptor.Payee_Status_OnStore_impl(r.Status, r, usr)

	logs.Storing("Payee", fmt.Sprintf("%s", r))

	// Please Create Functions Below in the adaptor/Payee_impl.go file
	err1 := adaptor.Payee_Delete_impl(r.KeyCounterpartyFirm)
	err2 := adaptor.Payee_Update_impl(r.KeyCounterpartyFirm, r, usr)
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
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SourceTable = get_String(rec, dm.Payee_SourceTable_sql, "")
		recItem.KeyCounterpartyFirm = get_String(rec, dm.Payee_KeyCounterpartyFirm_sql, "")
		recItem.KeyCounterpartyCentre = get_String(rec, dm.Payee_KeyCounterpartyCentre_sql, "")
		recItem.KeyCurrency = get_String(rec, dm.Payee_KeyCurrency_sql, "")
		recItem.KeyName = get_String(rec, dm.Payee_KeyName_sql, "")
		recItem.KeyNumber = get_String(rec, dm.Payee_KeyNumber_sql, "")
		recItem.KeyDirection = get_String(rec, dm.Payee_KeyDirection_sql, "")
		recItem.KeyType = get_String(rec, dm.Payee_KeyType_sql, "")
		recItem.FullName = get_String(rec, dm.Payee_FullName_sql, "")
		recItem.Address = get_String(rec, dm.Payee_Address_sql, "")
		recItem.PhoneNo = get_String(rec, dm.Payee_PhoneNo_sql, "")
		recItem.Country = get_String(rec, dm.Payee_Country_sql, "")
		recItem.Bic = get_String(rec, dm.Payee_Bic_sql, "")
		recItem.Iban = get_String(rec, dm.Payee_Iban_sql, "")
		recItem.AccountNo = get_String(rec, dm.Payee_AccountNo_sql, "")
		recItem.FedWireNo = get_String(rec, dm.Payee_FedWireNo_sql, "")
		recItem.SortCode = get_String(rec, dm.Payee_SortCode_sql, "")
		recItem.BankName = get_String(rec, dm.Payee_BankName_sql, "")
		recItem.BankPinCode = get_String(rec, dm.Payee_BankPinCode_sql, "")
		recItem.BankAddress = get_String(rec, dm.Payee_BankAddress_sql, "")
		recItem.Reason = get_String(rec, dm.Payee_Reason_sql, "")
		recItem.BankSettlementAcct = get_Bool(rec, dm.Payee_BankSettlementAcct_sql, "True")
		recItem.UpdatedUserId = get_String(rec, dm.Payee_UpdatedUserId_sql, "")

		// If there are fields below, create the methods in adaptor\Payee_impl.go

		recItem.Status = adaptor.Payee_Status_OnFetch_impl(recItem)

		//
		// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
		// END
		///
		//Add to the list
		//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}

func Payee_NewID(r dm.Payee) string {

	id := uuid.New().String()

	return id
}

// payee_Fetch read all Payee's
func Payee_New() (int, []dm.Payee, dm.Payee, error) {

	var r = dm.Payee{}
	var rList []dm.Payee

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	r.Status, r.Status_props = adaptor.Payee_Status_Impl(adaptor.NEW, r.KeyCounterpartyFirm, r.Status, r, r.Status_props)
	//
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	rList = append(rList, r)

	return 1, rList, r, nil
}
