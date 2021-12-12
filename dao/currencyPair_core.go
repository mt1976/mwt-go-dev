package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/currencypair.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CurrencyPair (currencypair)
// Endpoint 	        : CurrencyPair (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:11
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"log"
	
	"fmt"
	"net/http"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	das  "github.com/mt1976/mwt-go-dev/das"
	
	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
)

// CurrencyPair_GetList() returns a list of all CurrencyPair records
func CurrencyPair_GetList() (int, []dm.CurrencyPair, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CurrencyPair_SQLTable)
	count, currencypairList, _, _ := currencypair_Fetch(tsql)
	
	return count, currencypairList, nil
}



// CurrencyPair_GetByID() returns a single CurrencyPair record
func CurrencyPair_GetByID(id string) (int, dm.CurrencyPair, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CurrencyPair_SQLTable)
	tsql = tsql + " WHERE " + dm.CurrencyPair_SQLSearchID + "='" + id + "'"
	_, _, currencypairItem, _ := currencypair_Fetch(tsql)

	return 1, currencypairItem, nil
}



// CurrencyPair_DeleteByID() deletes a single CurrencyPair record
func CurrencyPair_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.CurrencyPair_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.CurrencyPair_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// CurrencyPair_Store() saves/stores a CurrencyPair record to the database
func CurrencyPair_Store(r dm.CurrencyPair,req *http.Request) error {

	err := currencypair_Save(r,Audit_User(req))

	return err
}

// CurrencyPair_StoreSystem() saves/stores a CurrencyPair record to the database
func CurrencyPair_StoreSystem(r dm.CurrencyPair) error {
	
	err := currencypair_Save(r,Audit_Host())

	return err
}

// currencypair_Save() saves/stores a CurrencyPair record to the database
func currencypair_Save(r dm.CurrencyPair,usr string) error {

    var err error

	logs.Storing("CurrencyPair",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = CurrencyPair_NewID(r)
	}


// Please Create Functions Below in the adaptor/CurrencyPair_impl.go file
	err1 := adaptor.CurrencyPair_Delete_Impl(r.Code)
	err2 := adaptor.CurrencyPair_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// currencypair_Fetch read all employees
func currencypair_Fetch(tsql string) (int, []dm.CurrencyPair, dm.CurrencyPair, error) {

	var recItem dm.CurrencyPair
	var recList []dm.CurrencyPair

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.CodeMajorCurrencyIsoCode  = get_String(rec, dm.CurrencyPair_CodeMajorCurrencyIsoCode, "")
   recItem.CodeMinorCurrencyIsoCode  = get_String(rec, dm.CurrencyPair_CodeMinorCurrencyIsoCode, "")
   recItem.ReciprocalActive  = get_Bool(rec, dm.CurrencyPair_ReciprocalActive, "True")
   recItem.Code  = get_String(rec, dm.CurrencyPair_Code, "")
   recItem.MajorName  = get_String(rec, dm.CurrencyPair_MajorName, "")
   recItem.MinorName  = get_String(rec, dm.CurrencyPair_MinorName, "")


// If there are fields below, create the methods in dao\CurrencyPair_Impl.go

















	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func CurrencyPair_NewID(r dm.CurrencyPair) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

