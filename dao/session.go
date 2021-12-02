package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/session.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Session (session)
// Endpoint 	        : Session (SessionID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 02/12/2021 at 19:40:08
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"log"
	"fmt"

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
func Session_Store(r dm.Session) error {

	logs.Storing("Session",fmt.Sprintf("%s", r))

	return nil

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
	// Automatically generated 02/12/2021 by matttownsend on silicon.local - START
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
// Automatically generated 02/12/2021 by matttownsend on silicon.local - END
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

