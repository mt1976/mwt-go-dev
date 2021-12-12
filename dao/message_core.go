package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/message.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Message (message)
// Endpoint 	        : Message (Message)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:17
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

// Message_GetList() returns a list of all Message records
func Message_GetList() (int, []dm.Message, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Message_SQLTable)
	count, messageList, _, _ := message_Fetch(tsql)
	
	return count, messageList, nil
}



// Message_GetByID() returns a single Message record
func Message_GetByID(id string) (int, dm.Message, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Message_SQLTable)
	tsql = tsql + " WHERE " + dm.Message_SQLSearchID + "='" + id + "'"
	_, _, messageItem, _ := message_Fetch(tsql)

	return 1, messageItem, nil
}



// Message_DeleteByID() deletes a single Message record
func Message_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Message_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Message_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Message_Store() saves/stores a Message record to the database
func Message_Store(r dm.Message,req *http.Request) error {

	err := message_Save(r,Audit_User(req))

	return err
}

// Message_StoreSystem() saves/stores a Message record to the database
func Message_StoreSystem(r dm.Message) error {
	
	err := message_Save(r,Audit_Host())

	return err
}

// message_Save() saves/stores a Message record to the database
func message_Save(r dm.Message,usr string) error {

    var err error

	logs.Storing("Message",fmt.Sprintf("%s", r))

	if len(r.Id) == 0 {
		r.Id = Message_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.Message_SYSId, r.SYSId)
ts = addData(ts, dm.Message_Id, r.Id)
ts = addData(ts, dm.Message_Message, r.Message)
ts = addData(ts, dm.Message_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.Message_SYSWho, r.SYSWho)
ts = addData(ts, dm.Message_SYSHost, r.SYSHost)
ts = addData(ts, dm.Message_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.Message_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.Message_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.Message_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.Message_SYSUpdatedHost, r.SYSUpdatedHost)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Message_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Message_Delete(r.Id)
	das.Execute(tsql)



	return err

}


// message_Fetch read all employees
func message_Fetch(tsql string) (int, []dm.Message, dm.Message, error) {

	var recItem dm.Message
	var recList []dm.Message

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.Message_SYSId, "0")
   recItem.Id  = get_String(rec, dm.Message_Id, "")
   recItem.Message  = get_String(rec, dm.Message_Message, "")
   recItem.SYSCreated  = get_String(rec, dm.Message_SYSCreated, "")
   recItem.SYSWho  = get_String(rec, dm.Message_SYSWho, "")
   recItem.SYSHost  = get_String(rec, dm.Message_SYSHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.Message_SYSUpdated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.Message_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.Message_SYSCreatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.Message_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.Message_SYSUpdatedHost, "")
// If there are fields below, create the methods in dao\Message_Impl.go























	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Message_NewID(r dm.Message) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

