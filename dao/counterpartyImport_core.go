package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/counterpartyimport.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CounterpartyImport (counterpartyimport)
// Endpoint 	        : CounterpartyImport (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:31:57
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

// CounterpartyImport_GetList() returns a list of all CounterpartyImport records
func CounterpartyImport_GetList() (int, []dm.CounterpartyImport, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyImport_SQLTable)
	count, counterpartyimportList, _, _ := counterpartyimport_Fetch(tsql)
	
	return count, counterpartyimportList, nil
}



// CounterpartyImport_GetByID() returns a single CounterpartyImport record
func CounterpartyImport_GetByID(id string) (int, dm.CounterpartyImport, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyImport_SQLTable)
	tsql = tsql + " WHERE " + dm.CounterpartyImport_SQLSearchID + "='" + id + "'"
	_, _, counterpartyimportItem, _ := counterpartyimport_Fetch(tsql)

	return 1, counterpartyimportItem, nil
}



// CounterpartyImport_DeleteByID() deletes a single CounterpartyImport record
func CounterpartyImport_Delete(id string) {


	adaptor.CounterpartyImport_Delete_impl(id)
	
}


// CounterpartyImport_Store() saves/stores a CounterpartyImport record to the database
func CounterpartyImport_Store(r dm.CounterpartyImport,req *http.Request) error {

	err := counterpartyimport_Save(r,Audit_User(req))

	return err
}

// CounterpartyImport_StoreSystem() saves/stores a CounterpartyImport record to the database
func CounterpartyImport_StoreSystem(r dm.CounterpartyImport) error {
	
	err := counterpartyimport_Save(r,Audit_Host())

	return err
}

// counterpartyimport_Save() saves/stores a CounterpartyImport record to the database
func counterpartyimport_Save(r dm.CounterpartyImport,usr string) error {

    var err error



	

	if len(r.CompID) == 0 {
		r.CompID = CounterpartyImport_NewID(r)
	}

// If there are fields below, create the methods in dao\counterpartyimport_impl.go










	
logs.Storing("CounterpartyImport",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/CounterpartyImport_impl.go file
	err1 := adaptor.CounterpartyImport_Delete_impl(r.CompID)
	err2 := adaptor.CounterpartyImport_Update_impl(r.CompID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// counterpartyimport_Fetch read all CounterpartyImport's
func counterpartyimport_Fetch(tsql string) (int, []dm.CounterpartyImport, dm.CounterpartyImport, error) {

	var recItem dm.CounterpartyImport
	var recList []dm.CounterpartyImport

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - START
   recItem.KeyImportID  = get_String(rec, dm.CounterpartyImport_KeyImportID, "")
   recItem.Firm  = get_String(rec, dm.CounterpartyImport_Firm, "")
   recItem.Centre  = get_String(rec, dm.CounterpartyImport_Centre, "")
   recItem.FirmName  = get_String(rec, dm.CounterpartyImport_FirmName, "")
   recItem.CentreName  = get_String(rec, dm.CounterpartyImport_CentreName, "")
   recItem.KeyOriginID  = get_String(rec, dm.CounterpartyImport_KeyOriginID, "")
   recItem.FullName  = get_String(rec, dm.CounterpartyImport_FullName, "")
   recItem.CompID  = get_String(rec, dm.CounterpartyImport_CompID, "")
// If there are fields below, create the methods in adaptor\CounterpartyImport_impl.go









	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func CounterpartyImport_NewID(r dm.CounterpartyImport) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

