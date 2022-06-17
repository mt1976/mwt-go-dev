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
// Date & Time		    : 17/06/2022 at 18:38:09
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	das "github.com/mt1976/mwt-go-dev/das"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
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
func DataLoaderMap_Store(r dm.DataLoaderMap, req *http.Request) error {

	err := dataloadermap_Save(r, Audit_User(req))

	return err
}

// DataLoaderMap_StoreSystem() saves/stores a DataLoaderMap record to the database
func DataLoaderMap_StoreSystem(r dm.DataLoaderMap) error {

	err := dataloadermap_Save(r, Audit_Host())

	return err
}

// dataloadermap_Save() saves/stores a DataLoaderMap record to the database
func dataloadermap_Save(r dm.DataLoaderMap, usr string) error {

	var err error

	if len(r.Id) == 0 {
		r.Id = DataLoaderMap_NewID(r)
	}

	// If there are fields below, create the methods in dao\dataloadermap_impl.go

	logs.Storing("DataLoaderMap", fmt.Sprintf("%s", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.DataLoaderMap_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.DataLoaderMap_Id_sql, r.Id)
	ts = addData(ts, dm.DataLoaderMap_Name_sql, r.Name)
	ts = addData(ts, dm.DataLoaderMap_Position_sql, r.Position)
	ts = addData(ts, dm.DataLoaderMap_Loader_sql, r.Loader)
	ts = addData(ts, dm.DataLoaderMap_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.DataLoaderMap_SYSWho_sql, r.SYSWho)
	ts = addData(ts, dm.DataLoaderMap_SYSHost_sql, r.SYSHost)
	ts = addData(ts, dm.DataLoaderMap_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.DataLoaderMap_Int_position_sql, r.Int_position)
	ts = addData(ts, dm.DataLoaderMap_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.DataLoaderMap_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.DataLoaderMap_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.DataLoaderMap_SYSUpdatedHost_sql, r.SYSUpdatedHost)

	//
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.DataLoaderMap_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	DataLoaderMap_Delete(r.Id)
	das.Execute(tsql)

	return err

}

// dataloadermap_Fetch read all DataLoaderMap's
func dataloadermap_Fetch(tsql string) (int, []dm.DataLoaderMap, dm.DataLoaderMap, error) {

	var recItem dm.DataLoaderMap
	var recList []dm.DataLoaderMap

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.DataLoaderMap_SYSId_sql, "0")
		recItem.Id = get_String(rec, dm.DataLoaderMap_Id_sql, "")
		recItem.Name = get_String(rec, dm.DataLoaderMap_Name_sql, "")
		recItem.Position = get_String(rec, dm.DataLoaderMap_Position_sql, "")
		recItem.Loader = get_String(rec, dm.DataLoaderMap_Loader_sql, "")
		recItem.SYSCreated = get_String(rec, dm.DataLoaderMap_SYSCreated_sql, "")
		recItem.SYSWho = get_String(rec, dm.DataLoaderMap_SYSWho_sql, "")
		recItem.SYSHost = get_String(rec, dm.DataLoaderMap_SYSHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.DataLoaderMap_SYSUpdated_sql, "")
		recItem.Int_position = get_Int(rec, dm.DataLoaderMap_Int_position_sql, "0")
		recItem.SYSCreatedBy = get_String(rec, dm.DataLoaderMap_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.DataLoaderMap_SYSCreatedHost_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.DataLoaderMap_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.DataLoaderMap_SYSUpdatedHost_sql, "")

		// If there are fields below, create the methods in adaptor\DataLoaderMap_impl.go

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

func DataLoaderMap_NewID(r dm.DataLoaderMap) string {

	id := uuid.New().String()

	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------
