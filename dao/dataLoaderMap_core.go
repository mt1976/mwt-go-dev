package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/dataloadermap.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : DataLoaderMap (dataloadermap)
// Endpoint 	        : DataLoaderMap (Id)
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

// DataLoaderMap_GetList() returns a list of all DataLoaderMap records
func DataLoaderMap_GetList() (int, []dm.DataLoaderMap, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.DataLoaderMap_SQLTable)
	count, dataloadermapList, _, _ := dataloadermap_Fetch(tsql)
	
	return count, dataloadermapList, nil
}



// DataLoaderMap_GetByID() returns a single DataLoaderMap record
func DataLoaderMap_GetByID(id string) (int, dm.DataLoaderMap, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.DataLoaderMap_SQLTable)
	tsql = tsql + " WHERE " + dm.DataLoaderMap_SQLSearchID + "='" + id + "'"
	_, _, dataloadermapItem, _ := dataloadermap_Fetch(tsql)

	return 1, dataloadermapItem, nil
}



// DataLoaderMap_DeleteByID() deletes a single DataLoaderMap record
func DataLoaderMap_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.DataLoaderMap_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.DataLoaderMap_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// DataLoaderMap_Store() saves/stores a DataLoaderMap record to the database
func DataLoaderMap_Store(r dm.DataLoaderMap,req *http.Request) error {

	err := dataloadermap_Save(r,Audit_User(req))

	return err
}

// DataLoaderMap_StoreSystem() saves/stores a DataLoaderMap record to the database
func DataLoaderMap_StoreSystem(r dm.DataLoaderMap) error {
	
	err := dataloadermap_Save(r,Audit_Host())

	return err
}

// dataloadermap_Save() saves/stores a DataLoaderMap record to the database
func dataloadermap_Save(r dm.DataLoaderMap,usr string) error {

    var err error

	logs.Storing("DataLoaderMap",fmt.Sprintf("%s", r))

	if len(r.Id) == 0 {
		r.Id = DataLoaderMap_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.DataLoaderMap_SYSId, r.SYSId)
ts = addData(ts, dm.DataLoaderMap_Id, r.Id)
ts = addData(ts, dm.DataLoaderMap_Name, r.Name)
ts = addData(ts, dm.DataLoaderMap_Position, r.Position)
ts = addData(ts, dm.DataLoaderMap_Loader, r.Loader)
ts = addData(ts, dm.DataLoaderMap_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.DataLoaderMap_SYSWho, r.SYSWho)
ts = addData(ts, dm.DataLoaderMap_SYSHost, r.SYSHost)
ts = addData(ts, dm.DataLoaderMap_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.DataLoaderMap_Int_position, r.Int_position)
ts = addData(ts, dm.DataLoaderMap_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.DataLoaderMap_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.DataLoaderMap_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.DataLoaderMap_SYSUpdatedHost, r.SYSUpdatedHost)
ts = addData(ts, dm.DataLoaderMap_Loader_Lookup, r.Loader_Lookup)
ts = addData(ts, dm.DataLoaderMap_LoaderDescription_Lookup, r.LoaderDescription_Lookup)
ts = addData(ts, dm.DataLoaderMap_LoaderType_Lookup, r.LoaderType_Lookup)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.DataLoaderMap_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	DataLoaderMap_Delete(r.Id)
	das.Execute(tsql)



	return err

}


// dataloadermap_Fetch read all employees
func dataloadermap_Fetch(tsql string) (int, []dm.DataLoaderMap, dm.DataLoaderMap, error) {

	var recItem dm.DataLoaderMap
	var recList []dm.DataLoaderMap

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.DataLoaderMap_SYSId, "0")
   recItem.Id  = get_String(rec, dm.DataLoaderMap_Id, "")
   recItem.Name  = get_String(rec, dm.DataLoaderMap_Name, "")
   recItem.Position  = get_String(rec, dm.DataLoaderMap_Position, "")
   recItem.Loader  = get_String(rec, dm.DataLoaderMap_Loader, "")
   recItem.SYSCreated  = get_String(rec, dm.DataLoaderMap_SYSCreated, "")
   recItem.SYSWho  = get_String(rec, dm.DataLoaderMap_SYSWho, "")
   recItem.SYSHost  = get_String(rec, dm.DataLoaderMap_SYSHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.DataLoaderMap_SYSUpdated, "")
   recItem.Int_position  = get_Int(rec, dm.DataLoaderMap_Int_position, "0")
   recItem.SYSCreatedBy  = get_String(rec, dm.DataLoaderMap_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.DataLoaderMap_SYSCreatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.DataLoaderMap_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.DataLoaderMap_SYSUpdatedHost, "")



// If there are fields below, create the methods in dao\DataLoaderMap_Impl.go



































	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func DataLoaderMap_NewID(r dm.DataLoaderMap) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

