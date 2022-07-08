package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/translation.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Translation (translation)
// Endpoint 	        : Translation (Message)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:58
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

// Translation_GetList() returns a list of all Translation records
func Translation_GetList() (int, []dm.Translation, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Translation_SQLTable)
	count, translationList, _, _ := translation_Fetch(tsql)

	return count, translationList, nil
}

// Translation_GetByID() returns a single Translation record
func Translation_GetByID(id string) (int, dm.Translation, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Translation_SQLTable)
	tsql = tsql + " WHERE " + dm.Translation_SQLSearchID + "='" + id + "'"
	_, _, translationItem, _ := translation_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, translationItem, nil
}

// Translation_DeleteByID() deletes a single Translation record
func Translation_Delete(id string) {

	object_Table := core.ApplicationSQLSchema() + "." + dm.Translation_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Translation_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}

// Translation_Store() saves/stores a Translation record to the database
func Translation_Store(r dm.Translation, req *http.Request) error {

	err, r := Translation_Validate(r)
	if err == nil {
		err = translation_Save(r, Audit_User(req))
	} else {
		logs.Information("Translation_Store()", err.Error())
	}

	return err
}

// Translation_StoreSystem() saves/stores a Translation record to the database
func Translation_StoreSystem(r dm.Translation) error {

	err, r := Translation_Validate(r)
	if err == nil {
		err = translation_Save(r, Audit_Host())
	} else {
		logs.Information("Translation_Store()", err.Error())
	}

	return err
}

// Translation_Validate() validates for saves/stores a Translation record to the database
func Translation_Validate(r dm.Translation) (error, dm.Translation) {
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

// translation_Save() saves/stores a Translation record to the database
func translation_Save(r dm.Translation, usr string) error {

	var err error

	if len(r.Id) == 0 {
		r.Id = Translation_NewID(r)
	}

	// If there are fields below, create the methods in dao\translation_impl.go

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost, Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("", usr)
	r.SYSUpdatedHost = Audit_Update("", Audit_Host())

	logs.Storing("Translation", fmt.Sprintf("%s", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.Translation_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Translation_Id_sql, r.Id)
	ts = addData(ts, dm.Translation_Class_sql, r.Class)
	ts = addData(ts, dm.Translation_Message_sql, r.Message)
	ts = addData(ts, dm.Translation_Translation_sql, r.Translation)
	ts = addData(ts, dm.Translation_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Translation_SYSWho_sql, r.SYSWho)
	ts = addData(ts, dm.Translation_SYSHost_sql, r.SYSHost)
	ts = addData(ts, dm.Translation_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Translation_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Translation_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Translation_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Translation_SYSUpdatedHost_sql, r.SYSUpdatedHost)

	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationSQLSchema(), dm.Translation_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Translation_Delete(r.Id)
	das.Execute(tsql)

	return err

}

// translation_Fetch read all Translation's
func translation_Fetch(tsql string) (int, []dm.Translation, dm.Translation, error) {

	var recItem dm.Translation
	var recList []dm.Translation

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.Translation_SYSId_sql, "0")
		recItem.Id = get_String(rec, dm.Translation_Id_sql, "")
		recItem.Class = get_String(rec, dm.Translation_Class_sql, "")
		recItem.Message = get_String(rec, dm.Translation_Message_sql, "")
		recItem.Translation = get_String(rec, dm.Translation_Translation_sql, "")
		recItem.SYSCreated = get_String(rec, dm.Translation_SYSCreated_sql, "")
		recItem.SYSWho = get_String(rec, dm.Translation_SYSWho_sql, "")
		recItem.SYSHost = get_String(rec, dm.Translation_SYSHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.Translation_SYSUpdated_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.Translation_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.Translation_SYSCreatedHost_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.Translation_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.Translation_SYSUpdatedHost_sql, "")

		// If there are fields below, create the methods in adaptor\Translation_impl.go

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

func Translation_NewID(r dm.Translation) string {

	id := uuid.New().String()

	return id
}

// translation_Fetch read all Translation's
func Translation_New() (int, []dm.Translation, dm.Translation, error) {

	var r = dm.Translation{}
	var rList []dm.Translation

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	rList = append(rList, r)

	return 1, rList, r, nil
}
