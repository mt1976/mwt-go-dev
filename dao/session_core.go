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
// Date & Time		    : 12/12/2021 at 16:13:19
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

// Session_GetList() returns a list of all Session records
func Session_GetList() (int, []dm.Session, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Session_SQLTable)
	count, sessionList, _, _ := session_Fetch(tsql)
	
	return count, sessionList, nil
}



// Session_GetByID() returns a single Session record
func Session_GetByID(id string) (int, dm.Session, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Session_SQLTable)
	tsql = tsql + " WHERE " + dm.Session_SQLSearchID + "='" + id + "'"
	_, _, sessionItem, _ := session_Fetch(tsql)

	return 1, sessionItem, nil
}



// Session_DeleteByID() deletes a single Session record
func Session_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Session_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Session_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Session_Store() saves/stores a Session record to the database
func Session_Store(r dm.Session,req *http.Request) error {

	err := session_Save(r,Audit_User(req))

	return err
}

// Session_StoreSystem() saves/stores a Session record to the database
func Session_StoreSystem(r dm.Session) error {
	
	err := session_Save(r,Audit_Host())

	return err
}

// session_Save() saves/stores a Session record to the database
func session_Save(r dm.Session,usr string) error {

    var err error

	logs.Storing("Session",fmt.Sprintf("%s", r))

	if len(r.Id) == 0 {
		r.Id = Session_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.Session_SYSId, r.SYSId)
ts = addData(ts, dm.Session_Apptoken, r.Apptoken)
ts = addData(ts, dm.Session_Createdate, r.Createdate)
ts = addData(ts, dm.Session_Createtime, r.Createtime)
ts = addData(ts, dm.Session_Uniqueid, r.Uniqueid)
ts = addData(ts, dm.Session_Sessiontoken, r.Sessiontoken)
ts = addData(ts, dm.Session_Username, r.Username)
ts = addData(ts, dm.Session_Password, r.Password)
ts = addData(ts, dm.Session_Userip, r.Userip)
ts = addData(ts, dm.Session_Userhost, r.Userhost)
ts = addData(ts, dm.Session_Appip, r.Appip)
ts = addData(ts, dm.Session_Apphost, r.Apphost)
ts = addData(ts, dm.Session_Issued, r.Issued)
ts = addData(ts, dm.Session_Expiry, r.Expiry)
ts = addData(ts, dm.Session_Expiryraw, r.Expiryraw)
ts = addData(ts, dm.Session_Brand, r.Brand)
ts = addData(ts, dm.Session_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.Session_SYSWho, r.SYSWho)
ts = addData(ts, dm.Session_SYSHost, r.SYSHost)
ts = addData(ts, dm.Session_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.Session_Id, r.Id)
ts = addData(ts, dm.Session_Expires, r.Expires)
ts = addData(ts, dm.Session_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.Session_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.Session_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.Session_SYSUpdatedHost, r.SYSUpdatedHost)
ts = addData(ts, dm.Session_SessionRole, r.SessionRole)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Session_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Session_Delete(r.Id)
	das.Execute(tsql)



	return err

}


// session_Fetch read all employees
func session_Fetch(tsql string) (int, []dm.Session, dm.Session, error) {

	var recItem dm.Session
	var recList []dm.Session

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.Session_SYSId, "0")
   recItem.Apptoken  = get_String(rec, dm.Session_Apptoken, "")
   recItem.Createdate  = get_String(rec, dm.Session_Createdate, "")
   recItem.Createtime  = get_String(rec, dm.Session_Createtime, "")
   recItem.Uniqueid  = get_String(rec, dm.Session_Uniqueid, "")
   recItem.Sessiontoken  = get_String(rec, dm.Session_Sessiontoken, "")
   recItem.Username  = get_String(rec, dm.Session_Username, "")
   recItem.Password  = get_String(rec, dm.Session_Password, "")
   recItem.Userip  = get_String(rec, dm.Session_Userip, "")
   recItem.Userhost  = get_String(rec, dm.Session_Userhost, "")
   recItem.Appip  = get_String(rec, dm.Session_Appip, "")
   recItem.Apphost  = get_String(rec, dm.Session_Apphost, "")
   recItem.Issued  = get_String(rec, dm.Session_Issued, "")
   recItem.Expiry  = get_String(rec, dm.Session_Expiry, "")
   recItem.Expiryraw  = get_String(rec, dm.Session_Expiryraw, "")
   recItem.Brand  = get_String(rec, dm.Session_Brand, "")
   recItem.SYSCreated  = get_String(rec, dm.Session_SYSCreated, "")
   recItem.SYSWho  = get_String(rec, dm.Session_SYSWho, "")
   recItem.SYSHost  = get_String(rec, dm.Session_SYSHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.Session_SYSUpdated, "")
   recItem.Id  = get_String(rec, dm.Session_Id, "")
   recItem.Expires  = get_Time(rec, dm.Session_Expires, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.Session_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.Session_SYSCreatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.Session_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.Session_SYSUpdatedHost, "")
   recItem.SessionRole  = get_String(rec, dm.Session_SessionRole, "")
// If there are fields below, create the methods in dao\Session_Impl.go























































	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Session_NewID(r dm.Session) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

