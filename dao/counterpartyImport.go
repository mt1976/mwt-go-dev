package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/counterpartyimport.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CounterpartyImport (counterpartyimport)
// Endpoint 	        : CounterpartyImport (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 02/12/2021 at 19:40:02
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

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.CounterpartyImport_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.CounterpartyImport_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}


// CounterpartyImport_Store() saves/stores a CounterpartyImport record to the database
func CounterpartyImport_Store(r dm.CounterpartyImport) error {

	logs.Storing("CounterpartyImport",fmt.Sprintf("%s", r))

	if len(r.CompID) == 0 {
		r.CompID = CounterpartyImport_NewID(r)
	}



	adaptor.CounterpartyImport_Delete(r.CompID)
	adaptor.CounterpartyImport_Update(r)



	return nil

}

// counterpartyimport_Fetch read all employees
func counterpartyimport_Fetch(tsql string) (int, []dm.CounterpartyImport, dm.CounterpartyImport, error) {

	var recItem dm.CounterpartyImport
	var recList []dm.CounterpartyImport

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 02/12/2021 by matttownsend on silicon.local - START
   recItem.KeyImportID  = get_String(rec, dm.CounterpartyImport_KeyImportID, "")
   recItem.Firm  = get_String(rec, dm.CounterpartyImport_Firm, "")
   recItem.Centre  = get_String(rec, dm.CounterpartyImport_Centre, "")
   recItem.FirmName  = get_String(rec, dm.CounterpartyImport_FirmName, "")
   recItem.CentreName  = get_String(rec, dm.CounterpartyImport_CentreName, "")
   recItem.KeyOriginID  = get_String(rec, dm.CounterpartyImport_KeyOriginID, "")
   recItem.FullName  = get_String(rec, dm.CounterpartyImport_FullName, "")
   recItem.CompID  = get_String(rec, dm.CounterpartyImport_CompID, "")
// Automatically generated 02/12/2021 by matttownsend on silicon.local - END
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

