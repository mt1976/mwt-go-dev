package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/tmpl.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Tmpl (tmpl)
// Endpoint 	        : Tmpl (TemplateID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:33
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	adaptor "github.com/mt1976/mwt-go-dev/adaptor"
	core "github.com/mt1976/mwt-go-dev/core"
	das "github.com/mt1976/mwt-go-dev/das"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

// Tmpl_GetList() returns a list of all Tmpl records
func Tmpl_GetList() (int, []dm.Tmpl, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Tmpl_SQLTable)
	count, tmplList, _, _ := tmpl_Fetch(tsql)

	return count, tmplList, nil
}

// Tmpl_GetByID() returns a single Tmpl record
func Tmpl_GetByID(id string) (int, dm.Tmpl, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Tmpl_SQLTable)
	tsql = tsql + " WHERE " + dm.Tmpl_SQLSearchID + "='" + id + "'"
	_, _, tmplItem, _ := tmpl_Fetch(tsql)

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	tmplItem.ID, tmplItem.ID_props = adaptor.Tmpl_ID_Impl(adaptor.GET, id, tmplItem.ID, tmplItem, tmplItem.ID_props)
	tmplItem.ExtraField, tmplItem.ExtraField_props = adaptor.Tmpl_ExtraField_Impl(adaptor.GET, id, tmplItem.ExtraField, tmplItem, tmplItem.ExtraField_props)
	tmplItem.ExtraField3, tmplItem.ExtraField3_props = adaptor.Tmpl_ExtraField3_Impl(adaptor.GET, id, tmplItem.ExtraField3, tmplItem, tmplItem.ExtraField3_props)
	tmplItem.TDate, tmplItem.TDate_props = adaptor.Tmpl_TDate_Impl(adaptor.GET, id, tmplItem.TDate, tmplItem, tmplItem.TDate_props)
	//
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, tmplItem, nil
}

// Tmpl_DeleteByID() deletes a single Tmpl record
func Tmpl_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Tmpl_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Tmpl_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}

// Tmpl_Store() saves/stores a Tmpl record to the database
func Tmpl_Store(r dm.Tmpl, req *http.Request) error {

	err, r := Tmpl_Validate(r)
	if err == nil {
		err = tmpl_Save(r, Audit_User(req))
	} else {
		logs.Information("Tmpl_Store()", err.Error())
	}

	return err
}

// Tmpl_StoreSystem() saves/stores a Tmpl record to the database
func Tmpl_StoreSystem(r dm.Tmpl) error {

	err, r := Tmpl_Validate(r)
	if err == nil {
		err = tmpl_Save(r, Audit_Host())
	} else {
		logs.Information("Tmpl_Store()", err.Error())
	}

	return err
}

// Tmpl_Validate() validates for saves/stores a Tmpl record to the database
func Tmpl_Validate(r dm.Tmpl) (error, dm.Tmpl) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	r.ID, r.ID_props = adaptor.Tmpl_ID_Impl(adaptor.PUT, r.ID, r.ID, r, r.ID_props)
	r.ExtraField, r.ExtraField_props = adaptor.Tmpl_ExtraField_Impl(adaptor.PUT, r.ID, r.ExtraField, r, r.ExtraField_props)
	r.ExtraField3, r.ExtraField3_props = adaptor.Tmpl_ExtraField3_Impl(adaptor.PUT, r.ID, r.ExtraField3, r, r.ExtraField3_props)
	r.TDate, r.TDate_props = adaptor.Tmpl_TDate_Impl(adaptor.PUT, r.ID, r.TDate, r, r.TDate_props)
	//
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return nil, r
}

//

// tmpl_Save() saves/stores a Tmpl record to the database
func tmpl_Save(r dm.Tmpl, usr string) error {

	var err error

	if len(r.ID) == 0 {
		r.ID = Tmpl_NewID(r)
	}

	// If there are fields below, create the methods in dao\tmpl_impl.go

	r.ID, err = adaptor.Tmpl_ID_OnStore_impl(r.ID, r, usr)
	r.ExtraField, err = adaptor.Tmpl_ExtraField_OnStore_impl(r.ExtraField, r, usr)

	r.ExtraField3, err = adaptor.Tmpl_ExtraField3_OnStore_impl(r.ExtraField3, r, usr)
	r.TDate, err = adaptor.Tmpl_TDate_OnStore_impl(r.TDate, r, usr)

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost, Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("", usr)
	r.SYSUpdatedHost = Audit_Update("", Audit_Host())

	logs.Storing("Tmpl", fmt.Sprintf("%s", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.Tmpl_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Tmpl_FIELD1_sql, r.FIELD1)
	ts = addData(ts, dm.Tmpl_FIELD2_sql, r.FIELD2)
	ts = addData(ts, dm.Tmpl_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Tmpl_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Tmpl_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Tmpl_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Tmpl_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Tmpl_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Tmpl_ID_sql, r.ID)

	//
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Tmpl_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Tmpl_Delete(r.ID)
	das.Execute(tsql)

	return err

}

// tmpl_Fetch read all Tmpl's
func tmpl_Fetch(tsql string) (int, []dm.Tmpl, dm.Tmpl, error) {

	var recItem dm.Tmpl
	var recList []dm.Tmpl

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.Tmpl_SYSId_sql, "0")
		recItem.FIELD1 = get_String(rec, dm.Tmpl_FIELD1_sql, "N")
		recItem.FIELD2 = get_String(rec, dm.Tmpl_FIELD2_sql, "")
		recItem.SYSCreated = get_String(rec, dm.Tmpl_SYSCreated_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.Tmpl_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.Tmpl_SYSCreatedHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.Tmpl_SYSUpdated_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.Tmpl_SYSUpdatedHost_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.Tmpl_SYSUpdatedBy_sql, "")
		recItem.ID = get_String(rec, dm.Tmpl_ID_sql, "")

		// If there are fields below, create the methods in adaptor\Tmpl_impl.go

		recItem.ID = adaptor.Tmpl_ID_OnFetch_impl(recItem)
		recItem.ExtraField = adaptor.Tmpl_ExtraField_OnFetch_impl(recItem)

		recItem.ExtraField3 = adaptor.Tmpl_ExtraField3_OnFetch_impl(recItem)
		recItem.TDate = adaptor.Tmpl_TDate_OnFetch_impl(recItem)

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

func Tmpl_NewID(r dm.Tmpl) string {

	id := uuid.New().String()

	return id
}

// tmpl_Fetch read all Tmpl's
func Tmpl_New() (int, []dm.Tmpl, dm.Tmpl, error) {

	var r = dm.Tmpl{}
	var rList []dm.Tmpl

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	r.ID, r.ID_props = adaptor.Tmpl_ID_Impl(adaptor.NEW, r.ID, r.ID, r, r.ID_props)
	r.ExtraField, r.ExtraField_props = adaptor.Tmpl_ExtraField_Impl(adaptor.NEW, r.ID, r.ExtraField, r, r.ExtraField_props)
	r.ExtraField3, r.ExtraField3_props = adaptor.Tmpl_ExtraField3_Impl(adaptor.NEW, r.ID, r.ExtraField3, r, r.ExtraField3_props)
	r.TDate, r.TDate_props = adaptor.Tmpl_TDate_Impl(adaptor.NEW, r.ID, r.TDate, r, r.TDate_props)
	//
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	rList = append(rList, r)

	return 1, rList, r, nil
}
