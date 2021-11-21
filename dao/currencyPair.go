package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/currencypair.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CurrencyPair (currencypair)
// Endpoint 	        : CurrencyPair (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 15:44:02
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"log"
	"fmt"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	das  "github.com/mt1976/mwt-go-dev/das"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
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
func CurrencyPair_Store(r dm.CurrencyPair) error {

	logs.Warning(fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = CurrencyPair_NewID(r)
	}



	adaptor.CurrencyPair_Delete(r.Code)
	adaptor.CurrencyPair_Update(r)


	return nil
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
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
    recItem.AppInternalID = get_String(rec, dm.CurrencyPair_Code,"")
   recItem.CodeMajorCurrencyIsoCode  = get_String(rec, dm.CurrencyPair_CodeMajorCurrencyIsoCode, "")
   recItem.CodeMinorCurrencyIsoCode  = get_String(rec, dm.CurrencyPair_CodeMinorCurrencyIsoCode, "")
   recItem.ReciprocalActive  = get_Bool(rec, dm.CurrencyPair_ReciprocalActive, "True")
   recItem.Code  = get_String(rec, dm.CurrencyPair_Code, "")
   recItem.MajorName  = get_String(rec, dm.CurrencyPair_MajorName, "")
   recItem.MinorName  = get_String(rec, dm.CurrencyPair_MinorName, "")


// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
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

