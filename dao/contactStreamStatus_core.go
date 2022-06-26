package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/contactstreamstatus.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : ContactStreamStatus (contactstreamstatus)
// Endpoint 	        : ContactStreamStatus (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:21
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

// ContactStreamStatus_GetList() returns a list of all ContactStreamStatus records
func ContactStreamStatus_GetList() (int, []dm.ContactStreamStatus, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.ContactStreamStatus_SQLTable)
	count, contactstreamstatusList, _, _ := contactstreamstatus_Fetch(tsql)
	
	return count, contactstreamstatusList, nil
}


// ContactStreamStatus_GetLookup() returns a lookup list of all ContactStreamStatus items in lookup format
func ContactStreamStatus_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, contactstreamstatusList, _ := ContactStreamStatus_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: contactstreamstatusList[i].StatusId, Name: contactstreamstatusList[i].Description})
	}
	return returnList
}


// ContactStreamStatus_GetByID() returns a single ContactStreamStatus record
func ContactStreamStatus_GetByID(id string) (int, dm.ContactStreamStatus, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.ContactStreamStatus_SQLTable)
	tsql = tsql + " WHERE " + dm.ContactStreamStatus_SQLSearchID + "='" + id + "'"
	_, _, contactstreamstatusItem, _ := contactstreamstatus_Fetch(tsql)

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, contactstreamstatusItem, nil
}



// ContactStreamStatus_DeleteByID() deletes a single ContactStreamStatus record
func ContactStreamStatus_Delete(id string) {


	adaptor.ContactStreamStatus_Delete_impl(id)
	
}


// ContactStreamStatus_Store() saves/stores a ContactStreamStatus record to the database
func ContactStreamStatus_Store(r dm.ContactStreamStatus,req *http.Request) error {

	err, r := ContactStreamStatus_Validate(r)
	if err == nil {
		err = contactstreamstatus_Save(r, Audit_User(req))
	} else {
		logs.Information("ContactStreamStatus_Store()", err.Error())
	}

	return err
}

// ContactStreamStatus_StoreSystem() saves/stores a ContactStreamStatus record to the database
func ContactStreamStatus_StoreSystem(r dm.ContactStreamStatus) error {
	
	err, r := ContactStreamStatus_Validate(r)
	if err == nil {
		err = contactstreamstatus_Save(r, Audit_Host())
	} else {
		logs.Information("ContactStreamStatus_Store()", err.Error())
	}

	return err
}

// ContactStreamStatus_Validate() validates for saves/stores a ContactStreamStatus record to the database
func ContactStreamStatus_Validate(r dm.ContactStreamStatus) (error,dm.ContactStreamStatus) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// contactstreamstatus_Save() saves/stores a ContactStreamStatus record to the database
func contactstreamstatus_Save(r dm.ContactStreamStatus,usr string) error {

    var err error



	

	if len(r.StatusId) == 0 {
		r.StatusId = ContactStreamStatus_NewID(r)
	}

// If there are fields below, create the methods in dao\contactstreamstatus_impl.go





	
logs.Storing("ContactStreamStatus",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/ContactStreamStatus_impl.go file
	err1 := adaptor.ContactStreamStatus_Delete_impl(r.StatusId)
	err2 := adaptor.ContactStreamStatus_Update_impl(r.StatusId,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// contactstreamstatus_Fetch read all ContactStreamStatus's
func contactstreamstatus_Fetch(tsql string) (int, []dm.ContactStreamStatus, dm.ContactStreamStatus, error) {

	var recItem dm.ContactStreamStatus
	var recList []dm.ContactStreamStatus

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.StatusId  = get_Int(rec, dm.ContactStreamStatus_StatusId_sql, "0")
	   recItem.Description  = get_String(rec, dm.ContactStreamStatus_Description_sql, "")
	   recItem.RecordState  = get_Int(rec, dm.ContactStreamStatus_RecordState_sql, "0")
	
	// If there are fields below, create the methods in adaptor\ContactStreamStatus_impl.go
	
	
	
	
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
	


func ContactStreamStatus_NewID(r dm.ContactStreamStatus) string {
	
			id := uuid.New().String()
	
	return id
}



// contactstreamstatus_Fetch read all ContactStreamStatus's
func ContactStreamStatus_New() (int, []dm.ContactStreamStatus, dm.ContactStreamStatus, error) {

	var r = dm.ContactStreamStatus{}
	var rList []dm.ContactStreamStatus
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}