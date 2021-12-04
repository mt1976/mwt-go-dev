package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/accountladder.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : AccountLadder (accountladder)
// Endpoint 	        : AccountLadder (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 04/12/2021 at 17:36:42
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
	
)

// AccountLadder_GetList() returns a list of all AccountLadder records
func AccountLadder_GetList() (int, []dm.AccountLadder, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.AccountLadder_SQLTable)
	count, accountladderList, _, _ := accountladder_Fetch(tsql)
	return count, accountladderList, nil
}



// AccountLadder_GetByID() returns a single AccountLadder record
func AccountLadder_GetByID(id string) (int, dm.AccountLadder, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.AccountLadder_SQLTable)
	tsql = tsql + " WHERE " + dm.AccountLadder_SQLSearchID + "='" + id + "'"

	_, _, accountladderItem, _ := accountladder_Fetch(tsql)
	return 1, accountladderItem, nil
}



// AccountLadder_DeleteByID() deletes a single AccountLadder record
func AccountLadder_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.AccountLadder_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.AccountLadder_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}


// AccountLadder_Store() saves/stores a AccountLadder record to the database
func AccountLadder_Store(r dm.AccountLadder) error {

	logs.Storing("AccountLadder",fmt.Sprintf("%s", r))

	return nil

}

// accountladder_Fetch read all employees
func accountladder_Fetch(tsql string) (int, []dm.AccountLadder, dm.AccountLadder, error) {

	var recItem dm.AccountLadder
	var recList []dm.AccountLadder

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 04/12/2021 by matttownsend on silicon.local - START
   recItem.SienaReference  = get_String(rec, dm.AccountLadder_SienaReference, "")
   recItem.BusinessDate  = get_Time(rec, dm.AccountLadder_BusinessDate, "")
   recItem.ContractNumber  = get_String(rec, dm.AccountLadder_ContractNumber, "")
   recItem.Balance  = get_Float(rec, dm.AccountLadder_Balance, "0.00")
   recItem.DealtCcy  = get_String(rec, dm.AccountLadder_DealtCcy, "")
   recItem.AmountDp  = get_Int(rec, dm.AccountLadder_AmountDp, "0")
// Automatically generated 04/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func AccountLadder_NewID(r dm.AccountLadder) string {
	
	
			id := uuid.New().String()

	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

