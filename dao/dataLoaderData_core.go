package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/dataloaderdata.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : DataLoaderData (dataloaderdata)
// Endpoint 	        : DataLoaderData (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:12
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

// DataLoaderData_GetList() returns a list of all DataLoaderData records
func DataLoaderData_GetList() (int, []dm.DataLoaderData, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.DataLoaderData_SQLTable)
	count, dataloaderdataList, _, _ := dataloaderdata_Fetch(tsql)
	
	return count, dataloaderdataList, nil
}



// DataLoaderData_GetByID() returns a single DataLoaderData record
func DataLoaderData_GetByID(id string) (int, dm.DataLoaderData, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.DataLoaderData_SQLTable)
	tsql = tsql + " WHERE " + dm.DataLoaderData_SQLSearchID + "='" + id + "'"
	_, _, dataloaderdataItem, _ := dataloaderdata_Fetch(tsql)

	return 1, dataloaderdataItem, nil
}



// DataLoaderData_DeleteByID() deletes a single DataLoaderData record
func DataLoaderData_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.DataLoaderData_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.DataLoaderData_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// DataLoaderData_Store() saves/stores a DataLoaderData record to the database
func DataLoaderData_Store(r dm.DataLoaderData,req *http.Request) error {

	err := dataloaderdata_Save(r,Audit_User(req))

	return err
}

// DataLoaderData_StoreSystem() saves/stores a DataLoaderData record to the database
func DataLoaderData_StoreSystem(r dm.DataLoaderData) error {
	
	err := dataloaderdata_Save(r,Audit_Host())

	return err
}

// dataloaderdata_Save() saves/stores a DataLoaderData record to the database
func dataloaderdata_Save(r dm.DataLoaderData,usr string) error {

    var err error

	logs.Storing("DataLoaderData",fmt.Sprintf("%s", r))

	if len(r.Id) == 0 {
		r.Id = DataLoaderData_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.DataLoaderData_SYSId, r.SYSId)
ts = addData(ts, dm.DataLoaderData_Id, r.Id)
ts = addData(ts, dm.DataLoaderData_Row, r.Row)
ts = addData(ts, dm.DataLoaderData_Position, r.Position)
ts = addData(ts, dm.DataLoaderData_Value, r.Value)
ts = addData(ts, dm.DataLoaderData_Loader, r.Loader)
ts = addData(ts, dm.DataLoaderData_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.DataLoaderData_SYSWho, r.SYSWho)
ts = addData(ts, dm.DataLoaderData_SYSHost, r.SYSHost)
ts = addData(ts, dm.DataLoaderData_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.DataLoaderData_Map, r.Map)
ts = addData(ts, dm.DataLoaderData_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.DataLoaderData_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.DataLoaderData_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.DataLoaderData_SYSUpdatedHost, r.SYSUpdatedHost)
ts = addData(ts, dm.DataLoaderData_Loader_Lookup, r.Loader_Lookup)
ts = addData(ts, dm.DataLoaderData_LoaderDescription_Lookup, r.LoaderDescription_Lookup)
ts = addData(ts, dm.DataLoaderData_LoaderType_Lookup, r.LoaderType_Lookup)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.DataLoaderData_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	DataLoaderData_Delete(r.Id)
	das.Execute(tsql)



	return err

}


// dataloaderdata_Fetch read all employees
func dataloaderdata_Fetch(tsql string) (int, []dm.DataLoaderData, dm.DataLoaderData, error) {

	var recItem dm.DataLoaderData
	var recList []dm.DataLoaderData

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.DataLoaderData_SYSId, "0")
   recItem.Id  = get_String(rec, dm.DataLoaderData_Id, "")
   recItem.Row  = get_String(rec, dm.DataLoaderData_Row, "")
   recItem.Position  = get_String(rec, dm.DataLoaderData_Position, "")
   recItem.Value  = get_String(rec, dm.DataLoaderData_Value, "")
   recItem.Loader  = get_String(rec, dm.DataLoaderData_Loader, "")
   recItem.SYSCreated  = get_String(rec, dm.DataLoaderData_SYSCreated, "")
   recItem.SYSWho  = get_String(rec, dm.DataLoaderData_SYSWho, "")
   recItem.SYSHost  = get_String(rec, dm.DataLoaderData_SYSHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.DataLoaderData_SYSUpdated, "")
   recItem.Map  = get_String(rec, dm.DataLoaderData_Map, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.DataLoaderData_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.DataLoaderData_SYSCreatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.DataLoaderData_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.DataLoaderData_SYSUpdatedHost, "")



// If there are fields below, create the methods in dao\DataLoaderData_Impl.go





































	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func DataLoaderData_NewID(r dm.DataLoaderData) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

