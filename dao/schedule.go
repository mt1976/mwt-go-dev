package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/schedule.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Schedule (schedule)
// Endpoint 	        : Schedule (Schedule)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 03/12/2021 at 13:17:00
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"log"
	"fmt"

	
	core "github.com/mt1976/mwt-go-dev/core"
	das  "github.com/mt1976/mwt-go-dev/das"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
	
)

// Schedule_GetList() returns a list of all Schedule records
func Schedule_GetList() (int, []dm.Schedule, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Schedule_SQLTable)
	count, scheduleList, _, _ := schedule_Fetch(tsql)
	return count, scheduleList, nil
}



// Schedule_GetByID() returns a single Schedule record
func Schedule_GetByID(id string) (int, dm.Schedule, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Schedule_SQLTable)
	tsql = tsql + " WHERE " + dm.Schedule_SQLSearchID + "='" + id + "'"

	_, _, scheduleItem, _ := schedule_Fetch(tsql)
	return 1, scheduleItem, nil
}



// Schedule_DeleteByID() deletes a single Schedule record
func Schedule_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Schedule_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Schedule_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}


// Schedule_Store() saves/stores a Schedule record to the database
func Schedule_Store(r dm.Schedule) error {

	logs.Storing("Schedule",fmt.Sprintf("%s", r))

	return nil

}

// schedule_Fetch read all employees
func schedule_Fetch(tsql string) (int, []dm.Schedule, dm.Schedule, error) {

	var recItem dm.Schedule
	var recList []dm.Schedule

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.Schedule_SYSId, "0")
   recItem.Id  = get_String(rec, dm.Schedule_Id, "")
   recItem.Name  = get_String(rec, dm.Schedule_Name, "")
   recItem.Description  = get_String(rec, dm.Schedule_Description, "")
   recItem.Schedule  = get_String(rec, dm.Schedule_Schedule, "")
   recItem.Started  = get_String(rec, dm.Schedule_Started, "")
   recItem.Lastrun  = get_String(rec, dm.Schedule_Lastrun, "")
   recItem.Message  = get_String(rec, dm.Schedule_Message, "")
   recItem.SYSCreated  = get_String(rec, dm.Schedule_SYSCreated, "")
   recItem.SYSWho  = get_String(rec, dm.Schedule_SYSWho, "")
   recItem.SYSHost  = get_String(rec, dm.Schedule_SYSHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.Schedule_SYSUpdated, "")
   recItem.Type  = get_String(rec, dm.Schedule_Type, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.Schedule_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.Schedule_SYSCreatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.Schedule_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.Schedule_SYSUpdatedHost, "")
   recItem.Human  = get_String(rec, dm.Schedule_Human, "")
// Automatically generated 03/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Schedule_NewID(r dm.Schedule) string {
	
	

		// schedule_NewIDImpl should be specified in dao/Schedule_Impl.go
		// to provide the implementation for the special case.
		// override should return id - override function should be defined as
		// schedule_NewIDImpl(r dm.Schedule) string {...}
		//
		id := schedule_NewIDImpl(r)
	
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

