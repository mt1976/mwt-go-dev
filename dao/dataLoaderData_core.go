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
// Date & Time		    : 26/06/2022 at 18:48:26
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

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
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

	err, r := DataLoaderData_Validate(r)
	if err == nil {
		err = dataloaderdata_Save(r, Audit_User(req))
	} else {
		logs.Information("DataLoaderData_Store()", err.Error())
	}

	return err
}

// DataLoaderData_StoreSystem() saves/stores a DataLoaderData record to the database
func DataLoaderData_StoreSystem(r dm.DataLoaderData) error {
	
	err, r := DataLoaderData_Validate(r)
	if err == nil {
		err = dataloaderdata_Save(r, Audit_Host())
	} else {
		logs.Information("DataLoaderData_Store()", err.Error())
	}

	return err
}

// DataLoaderData_Validate() validates for saves/stores a DataLoaderData record to the database
func DataLoaderData_Validate(r dm.DataLoaderData) (error,dm.DataLoaderData) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// dataloaderdata_Save() saves/stores a DataLoaderData record to the database
func dataloaderdata_Save(r dm.DataLoaderData,usr string) error {

    var err error



	

	if len(r.Id) == 0 {
		r.Id = DataLoaderData_NewID(r)
	}

// If there are fields below, create the methods in dao\dataloaderdata_impl.go

















	
logs.Storing("DataLoaderData",fmt.Sprintf("%s", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.DataLoaderData_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.DataLoaderData_Id_sql, r.Id)
	ts = addData(ts, dm.DataLoaderData_Row_sql, r.Row)
	ts = addData(ts, dm.DataLoaderData_Position_sql, r.Position)
	ts = addData(ts, dm.DataLoaderData_Value_sql, r.Value)
	ts = addData(ts, dm.DataLoaderData_Loader_sql, r.Loader)
	ts = addData(ts, dm.DataLoaderData_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.DataLoaderData_SYSWho_sql, r.SYSWho)
	ts = addData(ts, dm.DataLoaderData_SYSHost_sql, r.SYSHost)
	ts = addData(ts, dm.DataLoaderData_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.DataLoaderData_Map_sql, r.Map)
	ts = addData(ts, dm.DataLoaderData_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.DataLoaderData_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.DataLoaderData_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.DataLoaderData_SYSUpdatedHost_sql, r.SYSUpdatedHost)
		
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.DataLoaderData_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	DataLoaderData_Delete(r.Id)
	das.Execute(tsql)



	return err

}



// dataloaderdata_Fetch read all DataLoaderData's
func dataloaderdata_Fetch(tsql string) (int, []dm.DataLoaderData, dm.DataLoaderData, error) {

	var recItem dm.DataLoaderData
	var recList []dm.DataLoaderData

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.DataLoaderData_SYSId_sql, "0")
	   recItem.Id  = get_String(rec, dm.DataLoaderData_Id_sql, "")
	   recItem.Row  = get_String(rec, dm.DataLoaderData_Row_sql, "")
	   recItem.Position  = get_String(rec, dm.DataLoaderData_Position_sql, "")
	   recItem.Value  = get_String(rec, dm.DataLoaderData_Value_sql, "")
	   recItem.Loader  = get_String(rec, dm.DataLoaderData_Loader_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.DataLoaderData_SYSCreated_sql, "")
	   recItem.SYSWho  = get_String(rec, dm.DataLoaderData_SYSWho_sql, "")
	   recItem.SYSHost  = get_String(rec, dm.DataLoaderData_SYSHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.DataLoaderData_SYSUpdated_sql, "")
	   recItem.Map  = get_String(rec, dm.DataLoaderData_Map_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.DataLoaderData_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.DataLoaderData_SYSCreatedHost_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.DataLoaderData_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.DataLoaderData_SYSUpdatedHost_sql, "")
	
	// If there are fields below, create the methods in adaptor\DataLoaderData_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func DataLoaderData_NewID(r dm.DataLoaderData) string {
	
			id := uuid.New().String()
	
	return id
}



// dataloaderdata_Fetch read all DataLoaderData's
func DataLoaderData_New() (int, []dm.DataLoaderData, dm.DataLoaderData, error) {

	var r = dm.DataLoaderData{}
	var rList []dm.DataLoaderData
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}