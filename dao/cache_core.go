package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/cache.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Cache (cache)
// Endpoint 	        : Cache (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:44
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

// Cache_GetList() returns a list of all Cache records
func Cache_GetList() (int, []dm.Cache, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Cache_SQLTable)
	count, cacheList, _, _ := cache_Fetch(tsql)

	return count, cacheList, nil
}

// Cache_GetByID() returns a single Cache record
func Cache_GetByID(id string) (int, dm.Cache, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Cache_SQLTable)
	tsql = tsql + " WHERE " + dm.Cache_SQLSearchID + "='" + id + "'"
	_, _, cacheItem, _ := cache_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, cacheItem, nil
}

// Cache_DeleteByID() deletes a single Cache record
func Cache_Delete(id string) {

	object_Table := core.ApplicationSQLSchema() + "." + dm.Cache_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Cache_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}

// Cache_Store() saves/stores a Cache record to the database
func Cache_Store(r dm.Cache, req *http.Request) error {

	err, r := Cache_Validate(r)
	if err == nil {
		err = cache_Save(r, Audit_User(req))
	} else {
		logs.Information("Cache_Store()", err.Error())
	}

	return err
}

// Cache_StoreSystem() saves/stores a Cache record to the database
func Cache_StoreSystem(r dm.Cache) error {

	err, r := Cache_Validate(r)
	if err == nil {
		err = cache_Save(r, Audit_Host())
	} else {
		logs.Information("Cache_Store()", err.Error())
	}

	return err
}

// Cache_Validate() validates for saves/stores a Cache record to the database
func Cache_Validate(r dm.Cache) (error, dm.Cache) {
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

// cache_Save() saves/stores a Cache record to the database
func cache_Save(r dm.Cache, usr string) error {

	var err error

	if len(r.Id) == 0 {
		r.Id = Cache_NewID(r)
	}

	// If there are fields below, create the methods in dao\cache_impl.go

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost, Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("", usr)
	r.SYSUpdatedHost = Audit_Update("", Audit_Host())

	logs.Storing("Cache", fmt.Sprintf("%s", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.Cache_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Cache_Id_sql, r.Id)
	ts = addData(ts, dm.Cache_Object_sql, r.Object)
	ts = addData(ts, dm.Cache_Field_sql, r.Field)
	ts = addData(ts, dm.Cache_Value_sql, r.Value)
	ts = addData(ts, dm.Cache_Expiry_sql, r.Expiry)
	ts = addData(ts, dm.Cache_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Cache_SYSWho_sql, r.SYSWho)
	ts = addData(ts, dm.Cache_SYSHost_sql, r.SYSHost)
	ts = addData(ts, dm.Cache_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Cache_Source_sql, r.Source)
	ts = addData(ts, dm.Cache_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Cache_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Cache_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Cache_SYSUpdatedHost_sql, r.SYSUpdatedHost)

	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationSQLSchema(), dm.Cache_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Cache_Delete(r.Id)
	das.Execute(tsql)

	return err

}

// cache_Fetch read all Cache's
func cache_Fetch(tsql string) (int, []dm.Cache, dm.Cache, error) {

	var recItem dm.Cache
	var recList []dm.Cache

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.Cache_SYSId_sql, "0")
		recItem.Id = get_String(rec, dm.Cache_Id_sql, "")
		recItem.Object = get_String(rec, dm.Cache_Object_sql, "")
		recItem.Field = get_String(rec, dm.Cache_Field_sql, "")
		recItem.Value = get_String(rec, dm.Cache_Value_sql, "")
		recItem.Expiry = get_String(rec, dm.Cache_Expiry_sql, "")
		recItem.SYSCreated = get_String(rec, dm.Cache_SYSCreated_sql, "")
		recItem.SYSWho = get_String(rec, dm.Cache_SYSWho_sql, "")
		recItem.SYSHost = get_String(rec, dm.Cache_SYSHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.Cache_SYSUpdated_sql, "")
		recItem.Source = get_String(rec, dm.Cache_Source_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.Cache_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.Cache_SYSCreatedHost_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.Cache_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.Cache_SYSUpdatedHost_sql, "")

		// If there are fields below, create the methods in adaptor\Cache_impl.go

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

func Cache_NewID(r dm.Cache) string {

	id := uuid.New().String()

	return id
}

// cache_Fetch read all Cache's
func Cache_New() (int, []dm.Cache, dm.Cache, error) {

	var r = dm.Cache{}
	var rList []dm.Cache

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	rList = append(rList, r)

	return 1, rList, r, nil
}
