package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/externalmessage.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : ExternalMessage (externalmessage)
// Endpoint 	        : ExternalMessage (MessageID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/12/2021 at 09:31:50
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

// ExternalMessage_GetList() returns a list of all ExternalMessage records
func ExternalMessage_GetList() (int, []dm.ExternalMessage, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.ExternalMessage_SQLTable)
	count, externalmessageList, _, _ := externalmessage_Fetch(tsql)
	
	return count, externalmessageList, nil
}



// ExternalMessage_GetByID() returns a single ExternalMessage record
func ExternalMessage_GetByID(id string) (int, dm.ExternalMessage, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.ExternalMessage_SQLTable)
	tsql = tsql + " WHERE " + dm.ExternalMessage_SQLSearchID + "='" + id + "'"
	_, _, externalmessageItem, _ := externalmessage_Fetch(tsql)

	return 1, externalmessageItem, nil
}



// ExternalMessage_DeleteByID() deletes a single ExternalMessage record
func ExternalMessage_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.ExternalMessage_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.ExternalMessage_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// ExternalMessage_Store() saves/stores a ExternalMessage record to the database
func ExternalMessage_Store(r dm.ExternalMessage,req *http.Request) error {

	err := externalmessage_Save(r,Audit_User(req))

	return err
}

// ExternalMessage_StoreSystem() saves/stores a ExternalMessage record to the database
func ExternalMessage_StoreSystem(r dm.ExternalMessage) error {
	
	err := externalmessage_Save(r,Audit_Host())

	return err
}

// externalmessage_Save() saves/stores a ExternalMessage record to the database
func externalmessage_Save(r dm.ExternalMessage,usr string) error {

    var err error

	logs.Storing("ExternalMessage",fmt.Sprintf("%s", r))

	if len(r.MessageID) == 0 {
		r.MessageID = ExternalMessage_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.ExternalMessage_SYSId, r.SYSId)
ts = addData(ts, dm.ExternalMessage_MessageID, r.MessageID)
ts = addData(ts, dm.ExternalMessage_MessageFormat, r.MessageFormat)
ts = addData(ts, dm.ExternalMessage_MessageDeliveredTo, r.MessageDeliveredTo)
ts = addData(ts, dm.ExternalMessage_MessageBody, r.MessageBody)
ts = addData(ts, dm.ExternalMessage_MessageFilename, r.MessageFilename)
ts = addData(ts, dm.ExternalMessage_MessageLife, r.MessageLife)
ts = addData(ts, dm.ExternalMessage_MessageDate, r.MessageDate)
ts = addData(ts, dm.ExternalMessage_MessageTime, r.MessageTime)
ts = addData(ts, dm.ExternalMessage_MessageTimeoutAction, r.MessageTimeoutAction)
ts = addData(ts, dm.ExternalMessage_MessageACKNAK, r.MessageACKNAK)
ts = addData(ts, dm.ExternalMessage_ResponseID, r.ResponseID)
ts = addData(ts, dm.ExternalMessage_ResponseFilename, r.ResponseFilename)
ts = addData(ts, dm.ExternalMessage_ResponseBody, r.ResponseBody)
ts = addData(ts, dm.ExternalMessage_ResponseDate, r.ResponseDate)
ts = addData(ts, dm.ExternalMessage_ResponseTime, r.ResponseTime)
ts = addData(ts, dm.ExternalMessage_ResponseAdditionalInfo, r.ResponseAdditionalInfo)
ts = addData(ts, dm.ExternalMessage_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.ExternalMessage_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.ExternalMessage_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.ExternalMessage_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.ExternalMessage_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.ExternalMessage_SYSUpdatedHost, r.SYSUpdatedHost)
ts = addData(ts, dm.ExternalMessage_MessageTimeout, r.MessageTimeout)
ts = addData(ts, dm.ExternalMessage_MessageEmitted, r.MessageEmitted)
ts = addData(ts, dm.ExternalMessage_ResponseRecieved, r.ResponseRecieved)
ts = addData(ts, dm.ExternalMessage_MessageClass, r.MessageClass)
ts = addData(ts, dm.ExternalMessage_AppID, r.AppID)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.ExternalMessage_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	ExternalMessage_Delete(r.MessageID)
	das.Execute(tsql)



	return err

}


// externalmessage_Fetch read all employees
func externalmessage_Fetch(tsql string) (int, []dm.ExternalMessage, dm.ExternalMessage, error) {

	var recItem dm.ExternalMessage
	var recList []dm.ExternalMessage

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 14/12/2021 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.ExternalMessage_SYSId, "0")
   recItem.MessageID  = get_String(rec, dm.ExternalMessage_MessageID, "")
   recItem.MessageFormat  = get_String(rec, dm.ExternalMessage_MessageFormat, "")
   recItem.MessageDeliveredTo  = get_String(rec, dm.ExternalMessage_MessageDeliveredTo, "")
   recItem.MessageBody  = get_String(rec, dm.ExternalMessage_MessageBody, "")
   recItem.MessageFilename  = get_String(rec, dm.ExternalMessage_MessageFilename, "")
   recItem.MessageLife  = get_String(rec, dm.ExternalMessage_MessageLife, "")
   recItem.MessageDate  = get_String(rec, dm.ExternalMessage_MessageDate, "")
   recItem.MessageTime  = get_String(rec, dm.ExternalMessage_MessageTime, "")
   recItem.MessageTimeoutAction  = get_String(rec, dm.ExternalMessage_MessageTimeoutAction, "")
   recItem.MessageACKNAK  = get_String(rec, dm.ExternalMessage_MessageACKNAK, "")
   recItem.ResponseID  = get_String(rec, dm.ExternalMessage_ResponseID, "")
   recItem.ResponseFilename  = get_String(rec, dm.ExternalMessage_ResponseFilename, "")
   recItem.ResponseBody  = get_String(rec, dm.ExternalMessage_ResponseBody, "")
   recItem.ResponseDate  = get_String(rec, dm.ExternalMessage_ResponseDate, "")
   recItem.ResponseTime  = get_String(rec, dm.ExternalMessage_ResponseTime, "")
   recItem.ResponseAdditionalInfo  = get_String(rec, dm.ExternalMessage_ResponseAdditionalInfo, "")
   recItem.SYSCreated  = get_String(rec, dm.ExternalMessage_SYSCreated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.ExternalMessage_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.ExternalMessage_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.ExternalMessage_SYSUpdated, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.ExternalMessage_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.ExternalMessage_SYSUpdatedHost, "")
   recItem.MessageTimeout  = get_String(rec, dm.ExternalMessage_MessageTimeout, "")
   recItem.MessageEmitted  = get_String(rec, dm.ExternalMessage_MessageEmitted, "")
   recItem.ResponseRecieved  = get_String(rec, dm.ExternalMessage_ResponseRecieved, "")
   recItem.MessageClass  = get_String(rec, dm.ExternalMessage_MessageClass, "")
   recItem.AppID  = get_String(rec, dm.ExternalMessage_AppID, "")
// If there are fields below, create the methods in dao\ExternalMessage_Impl.go

























































	// Automatically generated 14/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func ExternalMessage_NewID(r dm.ExternalMessage) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

