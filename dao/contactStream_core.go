package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/contactstream.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : ContactStream (contactstream)
// Endpoint 	        : ContactStream (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:46
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

// ContactStream_GetList() returns a list of all ContactStream records
func ContactStream_GetList() (int, []dm.ContactStream, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.ContactStream_SQLTable)
	count, contactstreamList, _, _ := contactstream_Fetch(tsql)
	
	return count, contactstreamList, nil
}


// ContactStream_GetLookup() returns a lookup list of all ContactStream items in lookup format
func ContactStream_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, contactstreamList, _ := ContactStream_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: contactstreamList[i].StreamId, Name: contactstreamList[i].Description})
	}
	return returnList
}


// ContactStream_GetByID() returns a single ContactStream record
func ContactStream_GetByID(id string) (int, dm.ContactStream, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.ContactStream_SQLTable)
	tsql = tsql + " WHERE " + dm.ContactStream_SQLSearchID + "='" + id + "'"
	_, _, contactstreamItem, _ := contactstream_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, contactstreamItem, nil
}



// ContactStream_DeleteByID() deletes a single ContactStream record
func ContactStream_Delete(id string) {


	adaptor.ContactStream_Delete_impl(id)
	
}


// ContactStream_Store() saves/stores a ContactStream record to the database
func ContactStream_Store(r dm.ContactStream,req *http.Request) error {

	err, r := ContactStream_Validate(r)
	if err == nil {
		err = contactstream_Save(r, Audit_User(req))
	} else {
		logs.Information("ContactStream_Store()", err.Error())
	}

	return err
}

// ContactStream_StoreSystem() saves/stores a ContactStream record to the database
func ContactStream_StoreSystem(r dm.ContactStream) error {
	
	err, r := ContactStream_Validate(r)
	if err == nil {
		err = contactstream_Save(r, Audit_Host())
	} else {
		logs.Information("ContactStream_Store()", err.Error())
	}

	return err
}

// ContactStream_Validate() validates for saves/stores a ContactStream record to the database
func ContactStream_Validate(r dm.ContactStream) (error,dm.ContactStream) {
	var err error
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return err,r
}
//

// contactstream_Save() saves/stores a ContactStream record to the database
func contactstream_Save(r dm.ContactStream,usr string) error {

    var err error



	

	if len(r.StreamId) == 0 {
		r.StreamId = ContactStream_NewID(r)
	}

// If there are fields below, create the methods in dao\contactstream_impl.go












	
logs.Storing("ContactStream",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/ContactStream_impl.go file
	err1 := adaptor.ContactStream_Delete_impl(r.StreamId)
	err2 := adaptor.ContactStream_Update_impl(r.StreamId,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// contactstream_Fetch read all ContactStream's
func contactstream_Fetch(tsql string) (int, []dm.ContactStream, dm.ContactStream, error) {

	var recItem dm.ContactStream
	var recList []dm.ContactStream

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.StreamId  = get_Int(rec, dm.ContactStream_StreamId_sql, "0")
	   recItem.ContactId  = get_Int(rec, dm.ContactStream_ContactId_sql, "0")
	   recItem.Usr  = get_String(rec, dm.ContactStream_Usr_sql, "")
	   recItem.Description  = get_String(rec, dm.ContactStream_Description_sql, "")
	   recItem.StreamTypeId  = get_Int(rec, dm.ContactStream_StreamTypeId_sql, "0")
	   recItem.StreamStatusId  = get_Int(rec, dm.ContactStream_StreamStatusId_sql, "0")
	   recItem.RecordState  = get_Int(rec, dm.ContactStream_RecordState_sql, "0")
	   recItem.CreatedDateTime  = get_Time(rec, dm.ContactStream_CreatedDateTime_sql, "")
	   recItem.OpenedDateTime  = get_Time(rec, dm.ContactStream_OpenedDateTime_sql, "")
	   recItem.ClosedDateTime  = get_Time(rec, dm.ContactStream_ClosedDateTime_sql, "")
	
	// If there are fields below, create the methods in adaptor\ContactStream_impl.go
	
	
	
	
	
	
	
	
	
	
	
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
	


func ContactStream_NewID(r dm.ContactStream) string {
	
			id := uuid.New().String()
	
	return id
}



// contactstream_Fetch read all ContactStream's
func ContactStream_New() (int, []dm.ContactStream, dm.ContactStream, error) {

	var r = dm.ContactStream{}
	var rList []dm.ContactStream
	

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}