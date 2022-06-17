package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/accountladder.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : AccountLadder (accountladder)
// Endpoint 	        : AccountLadder (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:05
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


	adaptor.AccountLadder_Delete_impl(id)
	
}


// AccountLadder_Store() saves/stores a AccountLadder record to the database
func AccountLadder_Store(r dm.AccountLadder,req *http.Request) error {

	err := accountladder_Save(r,Audit_User(req))

	return err
}

// AccountLadder_StoreSystem() saves/stores a AccountLadder record to the database
func AccountLadder_StoreSystem(r dm.AccountLadder) error {
	
	err := accountladder_Save(r,Audit_Host())

	return err
}

// accountladder_Save() saves/stores a AccountLadder record to the database
func accountladder_Save(r dm.AccountLadder,usr string) error {

    var err error



	

	if len(r.SienaReference) == 0 {
		r.SienaReference = AccountLadder_NewID(r)
	}

// If there are fields below, create the methods in dao\accountladder_impl.go








	
logs.Storing("AccountLadder",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/AccountLadder_impl.go file
	err1 := adaptor.AccountLadder_Delete_impl(r.SienaReference)
	err2 := adaptor.AccountLadder_Update_impl(r.SienaReference,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// accountladder_Fetch read all AccountLadder's
func accountladder_Fetch(tsql string) (int, []dm.AccountLadder, dm.AccountLadder, error) {

	var recItem dm.AccountLadder
	var recList []dm.AccountLadder

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SienaReference  = get_String(rec, dm.AccountLadder_SienaReference_sql, "")
	   recItem.BusinessDate  = get_Time(rec, dm.AccountLadder_BusinessDate_sql, "")
	   recItem.ContractNumber  = get_String(rec, dm.AccountLadder_ContractNumber_sql, "")
	   recItem.Balance  = get_Float(rec, dm.AccountLadder_Balance_sql, "0.00")
	   recItem.DealtCcy  = get_String(rec, dm.AccountLadder_DealtCcy_sql, "")
	   recItem.AmountDp  = get_Int(rec, dm.AccountLadder_AmountDp_sql, "0")
	
	// If there are fields below, create the methods in adaptor\AccountLadder_impl.go
	
	
	
	
	
	
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
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

