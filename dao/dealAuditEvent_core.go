package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/dealauditevent.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : DealAuditEvent (dealauditevent)
// Endpoint 	        : DealAuditEvent (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:32:03
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (

	"fmt"
	"net/http"
core "github.com/mt1976/mwt-go-dev/core"
"github.com/google/uuid"
das  "github.com/mt1976/mwt-go-dev/das"
	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
)

// DealAuditEvent_GetList() returns a list of all DealAuditEvent records
func DealAuditEvent_GetList() (int, []dm.DealAuditEvent, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.DealAuditEvent_SQLTable)
	count, dealauditeventList, _, _ := dealauditevent_Fetch(tsql)
	
	return count, dealauditeventList, nil
}



// DealAuditEvent_GetByID() returns a single DealAuditEvent record
func DealAuditEvent_GetByID(id string) (int, dm.DealAuditEvent, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.DealAuditEvent_SQLTable)
	tsql = tsql + " WHERE " + dm.DealAuditEvent_SQLSearchID + "='" + id + "'"
	_, _, dealauditeventItem, _ := dealauditevent_Fetch(tsql)

	return 1, dealauditeventItem, nil
}



// DealAuditEvent_DeleteByID() deletes a single DealAuditEvent record
func DealAuditEvent_Delete(id string) {


	adaptor.DealAuditEvent_Delete_impl(id)
	
}


// DealAuditEvent_Store() saves/stores a DealAuditEvent record to the database
func DealAuditEvent_Store(r dm.DealAuditEvent,req *http.Request) error {

	err := dealauditevent_Save(r,Audit_User(req))

	return err
}

// DealAuditEvent_StoreSystem() saves/stores a DealAuditEvent record to the database
func DealAuditEvent_StoreSystem(r dm.DealAuditEvent) error {
	
	err := dealauditevent_Save(r,Audit_Host())

	return err
}

// dealauditevent_Save() saves/stores a DealAuditEvent record to the database
func dealauditevent_Save(r dm.DealAuditEvent,usr string) error {

    var err error



	

	if len(r.InternalId) == 0 {
		r.InternalId = DealAuditEvent_NewID(r)
	}

// If there are fields below, create the methods in dao\dealauditevent_impl.go























	
logs.Storing("DealAuditEvent",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/DealAuditEvent_impl.go file
	err1 := adaptor.DealAuditEvent_Delete_impl(r.InternalId)
	err2 := adaptor.DealAuditEvent_Update_impl(r.InternalId,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// dealauditevent_Fetch read all DealAuditEvent's
func dealauditevent_Fetch(tsql string) (int, []dm.DealAuditEvent, dm.DealAuditEvent, error) {

	var recItem dm.DealAuditEvent
	var recList []dm.DealAuditEvent

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - START
   recItem.DealRefNo  = get_String(rec, dm.DealAuditEvent_DealRefNo, "")
   recItem.EventIndex  = get_Int(rec, dm.DealAuditEvent_EventIndex, "0")
   recItem.CommonRefNo  = get_String(rec, dm.DealAuditEvent_CommonRefNo, "")
   recItem.Timestamp  = get_Time(rec, dm.DealAuditEvent_Timestamp, "")
   recItem.UTCTimestamp  = get_String(rec, dm.DealAuditEvent_UTCTimestamp, "")
   recItem.EventType  = get_String(rec, dm.DealAuditEvent_EventType, "")
   recItem.Status  = get_String(rec, dm.DealAuditEvent_Status, "")
   recItem.LimitOrderStatus  = get_String(rec, dm.DealAuditEvent_LimitOrderStatus, "")
   recItem.Usr  = get_String(rec, dm.DealAuditEvent_Usr, "")
   recItem.DealingInterface  = get_String(rec, dm.DealAuditEvent_DealingInterface, "")
   recItem.SourceIP  = get_String(rec, dm.DealAuditEvent_SourceIP, "")
   recItem.MessageID  = get_String(rec, dm.DealAuditEvent_MessageID, "")
   recItem.Details  = get_String(rec, dm.DealAuditEvent_Details, "")
   recItem.InternalId  = get_Int(rec, dm.DealAuditEvent_InternalId, "0")
   recItem.InternalDeleted  = get_Time(rec, dm.DealAuditEvent_InternalDeleted, "")
   recItem.UpdatedTransactionId  = get_String(rec, dm.DealAuditEvent_UpdatedTransactionId, "")
   recItem.UpdatedUserId  = get_String(rec, dm.DealAuditEvent_UpdatedUserId, "")
   recItem.UpdatedDateTime  = get_Time(rec, dm.DealAuditEvent_UpdatedDateTime, "")
   recItem.DeletedTransactionId  = get_String(rec, dm.DealAuditEvent_DeletedTransactionId, "")
   recItem.DeletedUserId  = get_String(rec, dm.DealAuditEvent_DeletedUserId, "")
   recItem.ChangeType  = get_String(rec, dm.DealAuditEvent_ChangeType, "")
// If there are fields below, create the methods in adaptor\DealAuditEvent_impl.go






















	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func DealAuditEvent_NewID(r dm.DealAuditEvent) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

