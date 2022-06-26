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
// Date & Time		    : 26/06/2022 at 18:48:33
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

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
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

	err, r := Systems_Validate(r)
	if err == nil {
		err = systems_Save(r, Audit_User(req))
	} else {
		logs.Information("Systems_Store()", err.Error())
	}

	return err
}

// Systems_StoreSystem() saves/stores a Systems record to the database
func Systems_StoreSystem(r dm.Systems) error {
	
	err, r := Systems_Validate(r)
	if err == nil {
		err = systems_Save(r, Audit_Host())
	} else {
		logs.Information("Systems_Store()", err.Error())
	}

	return err
}

// Systems_Validate() validates for saves/stores a Systems record to the database
func Systems_Validate(r dm.Systems) (error,dm.Systems) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// systems_Save() saves/stores a Systems record to the database
func systems_Save(r dm.Systems,usr string) error {

    var err error



	

	if len(r.Id) == 0 {
		r.Id = Systems_NewID(r)
	}

// If there are fields below, create the methods in dao\systems_impl.go





















	
logs.Storing("Systems",fmt.Sprintf("%s", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.Systems_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Systems_Id_sql, r.Id)
	ts = addData(ts, dm.Systems_Name_sql, r.Name)
	ts = addData(ts, dm.Systems_Staticin_sql, r.Staticin)
	ts = addData(ts, dm.Systems_Staticout_sql, r.Staticout)
	ts = addData(ts, dm.Systems_Txnin_sql, r.Txnin)
	ts = addData(ts, dm.Systems_Txnout_sql, r.Txnout)
	ts = addData(ts, dm.Systems_Fundscheckin_sql, r.Fundscheckin)
	ts = addData(ts, dm.Systems_Fundscheckout_sql, r.Fundscheckout)
	ts = addData(ts, dm.Systems_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Systems_SYSWho_sql, r.SYSWho)
	ts = addData(ts, dm.Systems_SYSHost_sql, r.SYSHost)
	ts = addData(ts, dm.Systems_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Systems_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Systems_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Systems_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Systems_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Systems_SWIFTin_sql, r.SWIFTin)
	ts = addData(ts, dm.Systems_SWIFTout_sql, r.SWIFTout)
		
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Systems_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Systems_Delete(r.Id)
	das.Execute(tsql)



	return err

}



// systems_Fetch read all Systems's
func systems_Fetch(tsql string) (int, []dm.Systems, dm.Systems, error) {

	var recItem dm.Systems
	var recList []dm.Systems

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.Systems_SYSId_sql, "0")
	   recItem.Id  = get_String(rec, dm.Systems_Id_sql, "")
	   recItem.Name  = get_String(rec, dm.Systems_Name_sql, "")
	   recItem.Staticin  = get_String(rec, dm.Systems_Staticin_sql, "")
	   recItem.Staticout  = get_String(rec, dm.Systems_Staticout_sql, "")
	   recItem.Txnin  = get_String(rec, dm.Systems_Txnin_sql, "")
	   recItem.Txnout  = get_String(rec, dm.Systems_Txnout_sql, "")
	   recItem.Fundscheckin  = get_String(rec, dm.Systems_Fundscheckin_sql, "")
	   recItem.Fundscheckout  = get_String(rec, dm.Systems_Fundscheckout_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.Systems_SYSCreated_sql, "")
	   recItem.SYSWho  = get_String(rec, dm.Systems_SYSWho_sql, "")
	   recItem.SYSHost  = get_String(rec, dm.Systems_SYSHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.Systems_SYSUpdated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.Systems_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.Systems_SYSCreatedHost_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.Systems_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.Systems_SYSUpdatedHost_sql, "")
	   recItem.SWIFTin  = get_String(rec, dm.Systems_SWIFTin_sql, "")
	   recItem.SWIFTout  = get_String(rec, dm.Systems_SWIFTout_sql, "")
	
	// If there are fields below, create the methods in adaptor\Systems_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func Systems_NewID(r dm.Systems) string {
	
			id := uuid.New().String()
	
	return id
}



// systems_Fetch read all Systems's
func Systems_New() (int, []dm.Systems, dm.Systems, error) {

	var r = dm.Systems{}
	var rList []dm.Systems
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}