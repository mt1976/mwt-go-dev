package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/counterpartyname.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CounterpartyName (counterpartyname)
// Endpoint 	        : CounterpartyName (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:08
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

// CounterpartyName_GetList() returns a list of all CounterpartyName records
func CounterpartyName_GetList() (int, []dm.CounterpartyName, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyName_SQLTable)
	count, counterpartynameList, _, _ := counterpartyname_Fetch(tsql)
	
	return count, counterpartynameList, nil
}



// CounterpartyName_GetByID() returns a single CounterpartyName record
func CounterpartyName_GetByID(id string) (int, dm.CounterpartyName, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyName_SQLTable)
	tsql = tsql + " WHERE " + dm.CounterpartyName_SQLSearchID + "='" + id + "'"
	_, _, counterpartynameItem, _ := counterpartyname_Fetch(tsql)

	return 1, counterpartynameItem, nil
}



// CounterpartyName_DeleteByID() deletes a single CounterpartyName record
func CounterpartyName_Delete(id string) {


	adaptor.CounterpartyName_Delete_impl(id)
	
}


// CounterpartyName_Store() saves/stores a CounterpartyName record to the database
func CounterpartyName_Store(r dm.CounterpartyName,req *http.Request) error {

	err := counterpartyname_Save(r,Audit_User(req))

	return err
}

// CounterpartyName_StoreSystem() saves/stores a CounterpartyName record to the database
func CounterpartyName_StoreSystem(r dm.CounterpartyName) error {
	
	err := counterpartyname_Save(r,Audit_Host())

	return err
}

// counterpartyname_Save() saves/stores a CounterpartyName record to the database
func counterpartyname_Save(r dm.CounterpartyName,usr string) error {

    var err error



	

	if len(r.CompID) == 0 {
		r.CompID = CounterpartyName_NewID(r)
	}

// If there are fields below, create the methods in dao\counterpartyname_impl.go






	
logs.Storing("CounterpartyName",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/CounterpartyName_impl.go file
	err1 := adaptor.CounterpartyName_Delete_impl(r.CompID)
	err2 := adaptor.CounterpartyName_Update_impl(r.CompID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// counterpartyname_Fetch read all CounterpartyName's
func counterpartyname_Fetch(tsql string) (int, []dm.CounterpartyName, dm.CounterpartyName, error) {

	var recItem dm.CounterpartyName
	var recList []dm.CounterpartyName

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.NameFirm  = get_String(rec, dm.CounterpartyName_NameFirm_sql, "")
	   recItem.NameCentre  = get_String(rec, dm.CounterpartyName_NameCentre_sql, "")
	   recItem.FullName  = get_String(rec, dm.CounterpartyName_FullName_sql, "")
	   recItem.CompID  = get_String(rec, dm.CounterpartyName_CompID_sql, "")
	
	// If there are fields below, create the methods in adaptor\CounterpartyName_impl.go
	
	
	
	
	
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
	


func CounterpartyName_NewID(r dm.CounterpartyName) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

