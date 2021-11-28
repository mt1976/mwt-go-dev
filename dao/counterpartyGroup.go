package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/counterpartygroup.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CounterpartyGroup (counterpartygroup)
// Endpoint 	        : CounterpartyGroup (Group)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 28/11/2021 at 22:54:54
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

// CounterpartyGroup_GetList() returns a list of all CounterpartyGroup records
func CounterpartyGroup_GetList() (int, []dm.CounterpartyGroup, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyGroup_SQLTable)
	count, counterpartygroupList, _, _ := counterpartygroup_Fetch(tsql)
	return count, counterpartygroupList, nil
}



// CounterpartyGroup_GetByID() returns a single CounterpartyGroup record
func CounterpartyGroup_GetByID(id string) (int, dm.CounterpartyGroup, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyGroup_SQLTable)
	tsql = tsql + " WHERE " + dm.CounterpartyGroup_SQLSearchID + "='" + id + "'"

	_, _, counterpartygroupItem, _ := counterpartygroup_Fetch(tsql)
	return 1, counterpartygroupItem, nil
}



// CounterpartyGroup_DeleteByID() deletes a single CounterpartyGroup record
func CounterpartyGroup_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.CounterpartyGroup_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.CounterpartyGroup_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}


// CounterpartyGroup_Store() saves/stores a CounterpartyGroup record to the database
func CounterpartyGroup_Store(r dm.CounterpartyGroup) error {

	logs.Storing("CounterpartyGroup",fmt.Sprintf("%s", r))

	if len(r.Name) == 0 {
		r.Name = CounterpartyGroup_NewID(r)
	}



	adaptor.CounterpartyGroup_Delete(r.Name)
	adaptor.CounterpartyGroup_Update(r)



	return nil

}

// counterpartygroup_Fetch read all employees
func counterpartygroup_Fetch(tsql string) (int, []dm.CounterpartyGroup, dm.CounterpartyGroup, error) {

	var recItem dm.CounterpartyGroup
	var recList []dm.CounterpartyGroup

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
   recItem.Name  = get_String(rec, dm.CounterpartyGroup_Name, "")
   recItem.CountryCode  = get_String(rec, dm.CounterpartyGroup_CountryCode, "")
   recItem.SuperGroup  = get_String(rec, dm.CounterpartyGroup_SuperGroup, "")


// Automatically generated 28/11/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func CounterpartyGroup_NewID(r dm.CounterpartyGroup) string {
	
	
			id := uuid.New().String()

	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

