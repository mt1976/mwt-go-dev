package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/session.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Session (session)
// Endpoint 	        : Session (SessionID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:57
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

// Session_GetList() returns a list of all Session records
func Session_GetList() (int, []dm.Session, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Session_SQLTable)
	count, sessionList, _, _ := session_Fetch(tsql)

	return count, sessionList, nil
}

// Session_GetByID() returns a single Session record
func Session_GetByID(id string) (int, dm.Session, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Session_SQLTable)
	tsql = tsql + " WHERE " + dm.Session_SQLSearchID + "='" + id + "'"
	_, _, sessionItem, _ := session_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, sessionItem, nil
}

// Session_DeleteByID() deletes a single Session record
func Session_Delete(id string) {

	object_Table := core.ApplicationSQLSchema() + "." + dm.Session_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Session_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}

// Session_Store() saves/stores a Session record to the database
func Session_Store(r dm.Session, req *http.Request) error {

	err, r := Session_Validate(r)
	if err == nil {
		err = session_Save(r, Audit_User(req))
	} else {
		logs.Information("Session_Store()", err.Error())
	}

	return err
}

// Session_StoreSystem() saves/stores a Session record to the database
func Session_StoreSystem(r dm.Session) error {

	err, r := Session_Validate(r)
	if err == nil {
		err = session_Save(r, Audit_Host())
	} else {
		logs.Information("Session_Store()", err.Error())
	}

	return err
}

// Session_Validate() validates for saves/stores a Session record to the database
func Session_Validate(r dm.Session) (error, dm.Session) {
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

// session_Save() saves/stores a Session record to the database
func session_Save(r dm.Session, usr string) error {

	var err error

	if len(r.Id) == 0 {
		r.Id = Session_NewID(r)
	}

	// If there are fields below, create the methods in dao\session_impl.go

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost, Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("", usr)
	r.SYSUpdatedHost = Audit_Update("", Audit_Host())

	logs.Storing("Session", fmt.Sprintf("%s", r))

	//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	ts = addData(ts, dm.Session_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Session_Apptoken_sql, r.Apptoken)
	ts = addData(ts, dm.Session_Createdate_sql, r.Createdate)
	ts = addData(ts, dm.Session_Createtime_sql, r.Createtime)
	ts = addData(ts, dm.Session_Uniqueid_sql, r.Uniqueid)
	ts = addData(ts, dm.Session_Sessiontoken_sql, r.Sessiontoken)
	ts = addData(ts, dm.Session_Username_sql, r.Username)
	ts = addData(ts, dm.Session_Password_sql, r.Password)
	ts = addData(ts, dm.Session_Userip_sql, r.Userip)
	ts = addData(ts, dm.Session_Userhost_sql, r.Userhost)
	ts = addData(ts, dm.Session_Appip_sql, r.Appip)
	ts = addData(ts, dm.Session_Apphost_sql, r.Apphost)
	ts = addData(ts, dm.Session_Issued_sql, r.Issued)
	ts = addData(ts, dm.Session_Expiry_sql, r.Expiry)
	ts = addData(ts, dm.Session_Expiryraw_sql, r.Expiryraw)
	ts = addData(ts, dm.Session_Brand_sql, r.Brand)
	ts = addData(ts, dm.Session_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Session_SYSWho_sql, r.SYSWho)
	ts = addData(ts, dm.Session_SYSHost_sql, r.SYSHost)
	ts = addData(ts, dm.Session_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Session_Id_sql, r.Id)
	ts = addData(ts, dm.Session_Expires_sql, r.Expires)
	ts = addData(ts, dm.Session_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Session_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Session_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Session_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Session_SessionRole_sql, r.SessionRole)

	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationSQLSchema(), dm.Session_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Session_Delete(r.Id)
	das.Execute(tsql)

	return err

}

// session_Fetch read all Session's
func session_Fetch(tsql string) (int, []dm.Session, dm.Session, error) {

	var recItem dm.Session
	var recList []dm.Session

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SYSId = get_Int(rec, dm.Session_SYSId_sql, "0")
		recItem.Apptoken = get_String(rec, dm.Session_Apptoken_sql, "")
		recItem.Createdate = get_String(rec, dm.Session_Createdate_sql, "")
		recItem.Createtime = get_String(rec, dm.Session_Createtime_sql, "")
		recItem.Uniqueid = get_String(rec, dm.Session_Uniqueid_sql, "")
		recItem.Sessiontoken = get_String(rec, dm.Session_Sessiontoken_sql, "")
		recItem.Username = get_String(rec, dm.Session_Username_sql, "")
		recItem.Password = get_String(rec, dm.Session_Password_sql, "")
		recItem.Userip = get_String(rec, dm.Session_Userip_sql, "")
		recItem.Userhost = get_String(rec, dm.Session_Userhost_sql, "")
		recItem.Appip = get_String(rec, dm.Session_Appip_sql, "")
		recItem.Apphost = get_String(rec, dm.Session_Apphost_sql, "")
		recItem.Issued = get_String(rec, dm.Session_Issued_sql, "")
		recItem.Expiry = get_String(rec, dm.Session_Expiry_sql, "")
		recItem.Expiryraw = get_String(rec, dm.Session_Expiryraw_sql, "")
		recItem.Brand = get_String(rec, dm.Session_Brand_sql, "")
		recItem.SYSCreated = get_String(rec, dm.Session_SYSCreated_sql, "")
		recItem.SYSWho = get_String(rec, dm.Session_SYSWho_sql, "")
		recItem.SYSHost = get_String(rec, dm.Session_SYSHost_sql, "")
		recItem.SYSUpdated = get_String(rec, dm.Session_SYSUpdated_sql, "")
		recItem.Id = get_String(rec, dm.Session_Id_sql, "")
		recItem.Expires = get_Time(rec, dm.Session_Expires_sql, "")
		recItem.SYSCreatedBy = get_String(rec, dm.Session_SYSCreatedBy_sql, "")
		recItem.SYSCreatedHost = get_String(rec, dm.Session_SYSCreatedHost_sql, "")
		recItem.SYSUpdatedBy = get_String(rec, dm.Session_SYSUpdatedBy_sql, "")
		recItem.SYSUpdatedHost = get_String(rec, dm.Session_SYSUpdatedHost_sql, "")
		recItem.SessionRole = get_String(rec, dm.Session_SessionRole_sql, "")

		// If there are fields below, create the methods in adaptor\Session_impl.go

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

func Session_NewID(r dm.Session) string {

	id := uuid.New().String()

	return id
}

// session_Fetch read all Session's
func Session_New() (int, []dm.Session, dm.Session, error) {

	var r = dm.Session{}
	var rList []dm.Session

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	rList = append(rList, r)

	return 1, rList, r, nil
}
