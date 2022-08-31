package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/dataloader.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : DataLoader (dataloader)
// Endpoint 	        : DataLoader (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:49
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

// DataLoader_GetList() returns a list of all DataLoader records
func DataLoader_GetList() (int, []dm.DataLoader, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.DataLoader_SQLTable)
	count, dataloaderList, _, _ := dataloader_Fetch(tsql)

	return count, dataloaderList, nil
}

// DataLoader_GetLookup() returns a lookup list of all DataLoader items in lookup format
func DataLoader_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, dataloaderList, _ := DataLoader_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: dataloaderList[i].Id, Name: dataloaderList[i].Name})
	}
	return returnList
}

// DataLoader_GetByID() returns a single DataLoader record
func DataLoader_GetByID(id string) (int, dm.DataLoader, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.DataLoader_SQLTable)
	tsql = tsql + " WHERE " + dm.DataLoader_SQLSearchID + "='" + id + "'"
	_, _, dataloaderItem, _ := dataloader_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, dataloaderItem, nil
}

// DataLoader_DeleteByID() deletes a single DataLoader record
func DataLoader_Delete(id string) {

	object_Table := core.ApplicationSQLSchema() + "." + dm.DataLoader_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.DataLoader_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}

// DataLoader_Store() saves/stores a DataLoader record to the database
func DataLoader_Store(r dm.DataLoader, req *http.Request) error {

	err, r := DataLoader_Validate(r)
	if err == nil {
		err = dataloader_Save(r, Audit_User(req))
	} else {
		logs.Information("DataLoader_Store()", err.Error())
	}

	return err
}

// DataLoader_StoreSystem() saves/stores a DataLoader record to the database
func DataLoader_StoreSystem(r dm.DataLoader) error {

	err, r := DataLoader_Validate(r)
	if err == nil {
		err = dataloader_Save(r, Audit_Host())
	} else {
		logs.Information("DataLoader_Store()", err.Error())
	}

	return err
}

// DataLoader_Validate() validates for saves/stores a DataLoader record to the database
func DataLoader_Validate(r dm.DataLoader) (error, dm.DataLoader) {
	var err error
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//

	return err, r
}

//

// dataloader_Save() saves/stores a DataLoader record to the database
func dataloader_Save(r dm.DataLoader, usr string) error {

	var err error

	if len(r.Id) == 0 {
		r.Id = DataLoader_NewID(r)
	}

	// If there are fields below, create the methods in dao\dataloader_impl.go

	logs.Storing("DataLoader", fmt.Sprintf("%s", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.DataLoader_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.DataLoader_Id_sql, r.Id)
	ts = addData(ts, dm.DataLoader_Name_sql, r.Name)
	ts = addData(ts, dm.DataLoader_Description_sql, r.Description)
	ts = addData(ts, dm.DataLoader_Filename_sql, r.Filename)
	ts = addData(ts, dm.DataLoader_Lastrun_sql, r.Lastrun)
	ts = addData(ts, dm.DataLoader_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.DataLoader_SYSWho_sql, r.SYSWho)
	ts = addData(ts, dm.DataLoader_SYSHost_sql, r.SYSHost)
	ts = addData(ts, dm.DataLoader_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.DataLoader_Type_sql, r.Type)
	ts = addData(ts, dm.DataLoader_Instance_sql, r.Instance)
	ts = addData(ts, dm.DataLoader_Extension_sql, r.Extension)
	ts = addData(ts, dm.DataLoader_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.DataLoader_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.DataLoader_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.DataLoader_SYSCreatedHost_sql, r.SYSCreatedHost)

	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationSQLSchema(), dm.DataLoader_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	DataLoader_Delete(r.Id)
	das.Execute(tsql)

	return err

}

// dataloader_Fetch read all DataLoader's
func dataloader_Fetch(tsql string) (int, []dm.DataLoader, dm.DataLoader, error) {

	var recItem dm.DataLoader
	var recList []dm.DataLoader

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.DataLoader_SYSId_sql, "0")
		recItem.Id = get_String(rec, dm.DataLoader_Id_sql, "")
		recItem.Name = get_String(rec, dm.DataLoader_Name_sql, "")
		recItem.Description = get_String(rec, dm.DataLoader_Description_sql, "")
		recItem.Filename = get_String(rec, dm.DataLoader_Filename_sql, "")
		recItem.Lastrun = get_String(rec, dm.DataLoader_Lastrun_sql, "")
		recItem.SYSCreated = get_String(rec, dm.DataLoader_SYSCreated_sql, "")
		recItem.SYSWho = get_String(rec, dm.DataLoader_SYSWho_sql, "")
		recItem.SYSHost = get_String(rec, dm.DataLoader_SYSHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.DataLoader_SYSUpdated_sql, "")
		recItem.Type = get_String(rec, dm.DataLoader_Type_sql, "")
		recItem.Instance = get_String(rec, dm.DataLoader_Instance_sql, "")
		recItem.Extension = get_String(rec, dm.DataLoader_Extension_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.DataLoader_SYSCreatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.DataLoader_SYSUpdatedHost_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.DataLoader_SYSUpdatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.DataLoader_SYSCreatedHost_sql, "")

		// If there are fields below, create the methods in adaptor\DataLoader_impl.go

		//
		// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
		// END
		///
		//Add to the list
		//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}

func DataLoader_NewID(r dm.DataLoader) string {

	id := uuid.New().String()

	return id
}

// dataloader_Fetch read all DataLoader's
func DataLoader_New() (int, []dm.DataLoader, dm.DataLoader, error) {

	var r = dm.DataLoader{}
	var rList []dm.DataLoader

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	rList = append(rList, r)

	return 1, rList, r, nil
}
