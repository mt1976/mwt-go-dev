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
// Date & Time		    : 14/06/2022 at 21:32:01
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

// DataLoader_GetList() returns a list of all DataLoader records
func DataLoader_GetList() (int, []dm.DataLoader, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.DataLoader_SQLTable)
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


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.DataLoader_SQLTable)
	tsql = tsql + " WHERE " + dm.DataLoader_SQLSearchID + "='" + id + "'"
	_, _, dataloaderItem, _ := dataloader_Fetch(tsql)

	return 1, dataloaderItem, nil
}



// DataLoader_DeleteByID() deletes a single DataLoader record
func DataLoader_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.DataLoader_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.DataLoader_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// DataLoader_Store() saves/stores a DataLoader record to the database
func DataLoader_Store(r dm.DataLoader,req *http.Request) error {

	err := dataloader_Save(r,Audit_User(req))

	return err
}

// DataLoader_StoreSystem() saves/stores a DataLoader record to the database
func DataLoader_StoreSystem(r dm.DataLoader) error {
	
	err := dataloader_Save(r,Audit_Host())

	return err
}

// dataloader_Save() saves/stores a DataLoader record to the database
func dataloader_Save(r dm.DataLoader,usr string) error {

    var err error



	

	if len(r.Id) == 0 {
		r.Id = DataLoader_NewID(r)
	}

// If there are fields below, create the methods in dao\dataloader_impl.go



















	
logs.Storing("DataLoader",fmt.Sprintf("%s", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.DataLoader_SYSId, r.SYSId)
	ts = addData(ts, dm.DataLoader_Id, r.Id)
	ts = addData(ts, dm.DataLoader_Name, r.Name)
	ts = addData(ts, dm.DataLoader_Description, r.Description)
	ts = addData(ts, dm.DataLoader_Filename, r.Filename)
	ts = addData(ts, dm.DataLoader_Lastrun, r.Lastrun)
	ts = addData(ts, dm.DataLoader_SYSCreated, r.SYSCreated)
	ts = addData(ts, dm.DataLoader_SYSWho, r.SYSWho)
	ts = addData(ts, dm.DataLoader_SYSHost, r.SYSHost)
	ts = addData(ts, dm.DataLoader_SYSUpdated, r.SYSUpdated)
	ts = addData(ts, dm.DataLoader_Type, r.Type)
	ts = addData(ts, dm.DataLoader_Instance, r.Instance)
	ts = addData(ts, dm.DataLoader_Extension, r.Extension)
	ts = addData(ts, dm.DataLoader_SYSCreatedBy, r.SYSCreatedBy)
	ts = addData(ts, dm.DataLoader_SYSUpdatedHost, r.SYSUpdatedHost)
	ts = addData(ts, dm.DataLoader_SYSUpdatedBy, r.SYSUpdatedBy)
	ts = addData(ts, dm.DataLoader_SYSCreatedHost, r.SYSCreatedHost)
		
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.DataLoader_SQLTable)
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
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.DataLoader_SYSId, "0")
   recItem.Id  = get_String(rec, dm.DataLoader_Id, "")
   recItem.Name  = get_String(rec, dm.DataLoader_Name, "")
   recItem.Description  = get_String(rec, dm.DataLoader_Description, "")
   recItem.Filename  = get_String(rec, dm.DataLoader_Filename, "")
   recItem.Lastrun  = get_String(rec, dm.DataLoader_Lastrun, "")
   recItem.SYSCreated  = get_String(rec, dm.DataLoader_SYSCreated, "")
   recItem.SYSWho  = get_String(rec, dm.DataLoader_SYSWho, "")
   recItem.SYSHost  = get_String(rec, dm.DataLoader_SYSHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.DataLoader_SYSUpdated, "")
   recItem.Type  = get_String(rec, dm.DataLoader_Type, "")
   recItem.Instance  = get_String(rec, dm.DataLoader_Instance, "")
   recItem.Extension  = get_String(rec, dm.DataLoader_Extension, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.DataLoader_SYSCreatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.DataLoader_SYSUpdatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.DataLoader_SYSUpdatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.DataLoader_SYSCreatedHost, "")
// If there are fields below, create the methods in adaptor\DataLoader_impl.go


















	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func DataLoader_NewID(r dm.DataLoader) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

