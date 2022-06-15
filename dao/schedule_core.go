package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/schedule.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Schedule (schedule)
// Endpoint 	        : Schedule (Schedule)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:32:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (

	"fmt"
	"net/http"
core "github.com/mt1976/mwt-go-dev/core"
"github.com/google/uuid"
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
func Schedule_Store(r dm.Schedule,req *http.Request) error {

	err := schedule_Save(r,Audit_User(req))

	return err
}

// Schedule_StoreSystem() saves/stores a Schedule record to the database
func Schedule_StoreSystem(r dm.Schedule) error {
	
	err := schedule_Save(r,Audit_Host())

	return err
}

// schedule_Save() saves/stores a Schedule record to the database
func schedule_Save(r dm.Schedule,usr string) error {

    var err error



	

	if len(r.Id) == 0 {
		r.Id = Schedule_NewID(r)
	}

// If there are fields below, create the methods in dao\schedule_impl.go




















	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	
logs.Storing("Schedule",fmt.Sprintf("%s", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.Schedule_SYSId, r.SYSId)
	ts = addData(ts, dm.Schedule_Id, r.Id)
	ts = addData(ts, dm.Schedule_Name, r.Name)
	ts = addData(ts, dm.Schedule_Description, r.Description)
	ts = addData(ts, dm.Schedule_Schedule, r.Schedule)
	ts = addData(ts, dm.Schedule_Started, r.Started)
	ts = addData(ts, dm.Schedule_Lastrun, r.Lastrun)
	ts = addData(ts, dm.Schedule_Message, r.Message)
	ts = addData(ts, dm.Schedule_SYSCreated, r.SYSCreated)
	ts = addData(ts, dm.Schedule_SYSWho, r.SYSWho)
	ts = addData(ts, dm.Schedule_SYSHost, r.SYSHost)
	ts = addData(ts, dm.Schedule_SYSUpdated, r.SYSUpdated)
	ts = addData(ts, dm.Schedule_Type, r.Type)
	ts = addData(ts, dm.Schedule_SYSCreatedBy, r.SYSCreatedBy)
	ts = addData(ts, dm.Schedule_SYSCreatedHost, r.SYSCreatedHost)
	ts = addData(ts, dm.Schedule_SYSUpdatedBy, r.SYSUpdatedBy)
	ts = addData(ts, dm.Schedule_SYSUpdatedHost, r.SYSUpdatedHost)
	ts = addData(ts, dm.Schedule_Human, r.Human)
		
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Schedule_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Schedule_Delete(r.Id)
	das.Execute(tsql)



	return err

}



// schedule_Fetch read all Schedule's
func schedule_Fetch(tsql string) (int, []dm.Schedule, dm.Schedule, error) {

	var recItem dm.Schedule
	var recList []dm.Schedule

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - START
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
// If there are fields below, create the methods in adaptor\Schedule_impl.go



















	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Schedule_NewID(r dm.Schedule) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

