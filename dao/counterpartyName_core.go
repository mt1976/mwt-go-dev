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
// Date & Time		    : 12/12/2021 at 16:13:10
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


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.CounterpartyName_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.CounterpartyName_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

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

	logs.Storing("CounterpartyName",fmt.Sprintf("%s", r))

	if len(r.CompID) == 0 {
		r.CompID = CounterpartyName_NewID(r)
	}


// Please Create Functions Below in the adaptor/CounterpartyName_impl.go file
	err1 := adaptor.CounterpartyName_Delete_Impl(r.CompID)
	err2 := adaptor.CounterpartyName_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// counterpartyname_Fetch read all employees
func counterpartyname_Fetch(tsql string) (int, []dm.CounterpartyName, dm.CounterpartyName, error) {

	var recItem dm.CounterpartyName
	var recList []dm.CounterpartyName

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.NameFirm  = get_String(rec, dm.CounterpartyName_NameFirm, "")
   recItem.NameCentre  = get_String(rec, dm.CounterpartyName_NameCentre, "")
   recItem.FullName  = get_String(rec, dm.CounterpartyName_FullName, "")
   recItem.CompID  = get_String(rec, dm.CounterpartyName_CompID, "")
// If there are fields below, create the methods in dao\CounterpartyName_Impl.go









	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
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

