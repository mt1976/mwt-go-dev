package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/systems.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Systems (systems)
// Endpoint 	        : Systems (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:19
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"log"
	
	"fmt"
	"net/http"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	das  "github.com/mt1976/mwt-go-dev/das"
	
	
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
)

// Systems_GetList() returns a list of all Systems records
func Systems_GetList() (int, []dm.Systems, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Systems_SQLTable)
	count, systemsList, _, _ := systems_Fetch(tsql)
	
	return count, systemsList, nil
}



// Systems_GetByID() returns a single Systems record
func Systems_GetByID(id string) (int, dm.Systems, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Systems_SQLTable)
	tsql = tsql + " WHERE " + dm.Systems_SQLSearchID + "='" + id + "'"
	_, _, systemsItem, _ := systems_Fetch(tsql)

	return 1, systemsItem, nil
}



// Systems_DeleteByID() deletes a single Systems record
func Systems_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Systems_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Systems_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Systems_Store() saves/stores a Systems record to the database
func Systems_Store(r dm.Systems,req *http.Request) error {

	err := systems_Save(r,Audit_User(req))

	return err
}

// Systems_StoreSystem() saves/stores a Systems record to the database
func Systems_StoreSystem(r dm.Systems) error {
	
	err := systems_Save(r,Audit_Host())

	return err
}

// systems_Save() saves/stores a Systems record to the database
func systems_Save(r dm.Systems,usr string) error {

    var err error

	logs.Storing("Systems",fmt.Sprintf("%s", r))

	if len(r.Id) == 0 {
		r.Id = Systems_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.Systems_SYSId, r.SYSId)
ts = addData(ts, dm.Systems_Id, r.Id)
ts = addData(ts, dm.Systems_Name, r.Name)
ts = addData(ts, dm.Systems_Staticin, r.Staticin)
ts = addData(ts, dm.Systems_Staticout, r.Staticout)
ts = addData(ts, dm.Systems_Txnin, r.Txnin)
ts = addData(ts, dm.Systems_Txnout, r.Txnout)
ts = addData(ts, dm.Systems_Fundscheckin, r.Fundscheckin)
ts = addData(ts, dm.Systems_Fundscheckout, r.Fundscheckout)
ts = addData(ts, dm.Systems_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.Systems_SYSWho, r.SYSWho)
ts = addData(ts, dm.Systems_SYSHost, r.SYSHost)
ts = addData(ts, dm.Systems_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.Systems_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.Systems_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.Systems_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.Systems_SYSUpdatedHost, r.SYSUpdatedHost)
ts = addData(ts, dm.Systems_SWIFTin, r.SWIFTin)
ts = addData(ts, dm.Systems_SWIFTout, r.SWIFTout)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Systems_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Systems_Delete(r.Id)
	das.Execute(tsql)



	return err

}


// systems_Fetch read all employees
func systems_Fetch(tsql string) (int, []dm.Systems, dm.Systems, error) {

	var recItem dm.Systems
	var recList []dm.Systems

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.Systems_SYSId, "0")
   recItem.Id  = get_String(rec, dm.Systems_Id, "")
   recItem.Name  = get_String(rec, dm.Systems_Name, "")
   recItem.Staticin  = get_String(rec, dm.Systems_Staticin, "")
   recItem.Staticout  = get_String(rec, dm.Systems_Staticout, "")
   recItem.Txnin  = get_String(rec, dm.Systems_Txnin, "")
   recItem.Txnout  = get_String(rec, dm.Systems_Txnout, "")
   recItem.Fundscheckin  = get_String(rec, dm.Systems_Fundscheckin, "")
   recItem.Fundscheckout  = get_String(rec, dm.Systems_Fundscheckout, "")
   recItem.SYSCreated  = get_String(rec, dm.Systems_SYSCreated, "")
   recItem.SYSWho  = get_String(rec, dm.Systems_SYSWho, "")
   recItem.SYSHost  = get_String(rec, dm.Systems_SYSHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.Systems_SYSUpdated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.Systems_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.Systems_SYSCreatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.Systems_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.Systems_SYSUpdatedHost, "")
   recItem.SWIFTin  = get_String(rec, dm.Systems_SWIFTin, "")
   recItem.SWIFTout  = get_String(rec, dm.Systems_SWIFTout, "")
// If there are fields below, create the methods in dao\Systems_Impl.go







































	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Systems_NewID(r dm.Systems) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

