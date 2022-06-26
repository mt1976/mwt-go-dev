package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/dealauditeventconversations.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : DealAuditEventConversations (dealauditeventconversations)
// Endpoint 	        : DealAuditEventConversations (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:27
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

// DealAuditEventConversations_GetList() returns a list of all DealAuditEventConversations records
func DealAuditEventConversations_GetList() (int, []dm.DealAuditEventConversations, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.DealAuditEventConversations_SQLTable)
	count, dealauditeventconversationsList, _, _ := dealauditeventconversations_Fetch(tsql)
	
	return count, dealauditeventconversationsList, nil
}



// DealAuditEventConversations_GetByID() returns a single DealAuditEventConversations record
func DealAuditEventConversations_GetByID(id string) (int, dm.DealAuditEventConversations, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.DealAuditEventConversations_SQLTable)
	tsql = tsql + " WHERE " + dm.DealAuditEventConversations_SQLSearchID + "='" + id + "'"
	_, _, dealauditeventconversationsItem, _ := dealauditeventconversations_Fetch(tsql)

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, dealauditeventconversationsItem, nil
}



// DealAuditEventConversations_DeleteByID() deletes a single DealAuditEventConversations record
func DealAuditEventConversations_Delete(id string) {


	adaptor.DealAuditEventConversations_Delete_impl(id)
	
}


// DealAuditEventConversations_Store() saves/stores a DealAuditEventConversations record to the database
func DealAuditEventConversations_Store(r dm.DealAuditEventConversations,req *http.Request) error {

	err, r := DealAuditEventConversations_Validate(r)
	if err == nil {
		err = dealauditeventconversations_Save(r, Audit_User(req))
	} else {
		logs.Information("DealAuditEventConversations_Store()", err.Error())
	}

	return err
}

// DealAuditEventConversations_StoreSystem() saves/stores a DealAuditEventConversations record to the database
func DealAuditEventConversations_StoreSystem(r dm.DealAuditEventConversations) error {
	
	err, r := DealAuditEventConversations_Validate(r)
	if err == nil {
		err = dealauditeventconversations_Save(r, Audit_Host())
	} else {
		logs.Information("DealAuditEventConversations_Store()", err.Error())
	}

	return err
}

// DealAuditEventConversations_Validate() validates for saves/stores a DealAuditEventConversations record to the database
func DealAuditEventConversations_Validate(r dm.DealAuditEventConversations) (error,dm.DealAuditEventConversations) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// dealauditeventconversations_Save() saves/stores a DealAuditEventConversations record to the database
func dealauditeventconversations_Save(r dm.DealAuditEventConversations,usr string) error {

    var err error



	

	if len(r.InternalId) == 0 {
		r.InternalId = DealAuditEventConversations_NewID(r)
	}

// If there are fields below, create the methods in dao\dealauditeventconversations_impl.go














	
logs.Storing("DealAuditEventConversations",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/DealAuditEventConversations_impl.go file
	err1 := adaptor.DealAuditEventConversations_Delete_impl(r.InternalId)
	err2 := adaptor.DealAuditEventConversations_Update_impl(r.InternalId,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// dealauditeventconversations_Fetch read all DealAuditEventConversations's
func dealauditeventconversations_Fetch(tsql string) (int, []dm.DealAuditEventConversations, dm.DealAuditEventConversations, error) {

	var recItem dm.DealAuditEventConversations
	var recList []dm.DealAuditEventConversations

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.DealRefNo  = get_String(rec, dm.DealAuditEventConversations_DealRefNo_sql, "")
	   recItem.EventIndex  = get_Int(rec, dm.DealAuditEventConversations_EventIndex_sql, "0")
	   recItem.MessageIndex  = get_Int(rec, dm.DealAuditEventConversations_MessageIndex_sql, "0")
	   recItem.Message  = get_String(rec, dm.DealAuditEventConversations_Message_sql, "")
	   recItem.InternalId  = get_Int(rec, dm.DealAuditEventConversations_InternalId_sql, "0")
	   recItem.InternalDeleted  = get_Time(rec, dm.DealAuditEventConversations_InternalDeleted_sql, "")
	   recItem.UpdatedTransactionId  = get_String(rec, dm.DealAuditEventConversations_UpdatedTransactionId_sql, "")
	   recItem.UpdatedUserId  = get_String(rec, dm.DealAuditEventConversations_UpdatedUserId_sql, "")
	   recItem.UpdatedDateTime  = get_Time(rec, dm.DealAuditEventConversations_UpdatedDateTime_sql, "")
	   recItem.DeletedTransactionId  = get_String(rec, dm.DealAuditEventConversations_DeletedTransactionId_sql, "")
	   recItem.DeletedUserId  = get_String(rec, dm.DealAuditEventConversations_DeletedUserId_sql, "")
	   recItem.ChangeType  = get_String(rec, dm.DealAuditEventConversations_ChangeType_sql, "")
	
	// If there are fields below, create the methods in adaptor\DealAuditEventConversations_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func DealAuditEventConversations_NewID(r dm.DealAuditEventConversations) string {
	
			id := uuid.New().String()
	
	return id
}



// dealauditeventconversations_Fetch read all DealAuditEventConversations's
func DealAuditEventConversations_New() (int, []dm.DealAuditEventConversations, dm.DealAuditEventConversations, error) {

	var r = dm.DealAuditEventConversations{}
	var rList []dm.DealAuditEventConversations
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}