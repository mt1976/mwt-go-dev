package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/contactdeals.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : ContactDeals (contactdeals)
// Endpoint 	        : ContactDeals (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 24/06/2022 at 16:01:55
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

// ContactDeals_GetList() returns a list of all ContactDeals records
func ContactDeals_GetList() (int, []dm.ContactDeals, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.ContactDeals_SQLTable)
	count, contactdealsList, _, _ := contactdeals_Fetch(tsql)
	
	return count, contactdealsList, nil
}


// ContactDeals_GetLookup() returns a lookup list of all ContactDeals items in lookup format
func ContactDeals_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, contactdealsList, _ := ContactDeals_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: contactdealsList[i].NoteId, Name: contactdealsList[i].RefNo})
	}
	return returnList
}


// ContactDeals_GetByID() returns a single ContactDeals record
func ContactDeals_GetByID(id string) (int, dm.ContactDeals, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.ContactDeals_SQLTable)
	tsql = tsql + " WHERE " + dm.ContactDeals_SQLSearchID + "='" + id + "'"
	_, _, contactdealsItem, _ := contactdeals_Fetch(tsql)

	return 1, contactdealsItem, nil
}



// ContactDeals_DeleteByID() deletes a single ContactDeals record
func ContactDeals_Delete(id string) {


	adaptor.ContactDeals_Delete_impl(id)
	
}


// ContactDeals_Store() saves/stores a ContactDeals record to the database
func ContactDeals_Store(r dm.ContactDeals,req *http.Request) error {

	err := contactdeals_Save(r,Audit_User(req))

	return err
}

// ContactDeals_StoreSystem() saves/stores a ContactDeals record to the database
func ContactDeals_StoreSystem(r dm.ContactDeals) error {
	
	err := contactdeals_Save(r,Audit_Host())

	return err
}

// contactdeals_Save() saves/stores a ContactDeals record to the database
func contactdeals_Save(r dm.ContactDeals,usr string) error {

    var err error



	

	if len(r.NoteId) == 0 {
		r.NoteId = ContactDeals_NewID(r)
	}

// If there are fields below, create the methods in dao\contactdeals_impl.go





	
logs.Storing("ContactDeals",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/ContactDeals_impl.go file
	err1 := adaptor.ContactDeals_Delete_impl(r.NoteId)
	err2 := adaptor.ContactDeals_Update_impl(r.NoteId,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// contactdeals_Fetch read all ContactDeals's
func contactdeals_Fetch(tsql string) (int, []dm.ContactDeals, dm.ContactDeals, error) {

	var recItem dm.ContactDeals
	var recList []dm.ContactDeals

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 24/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.RefNo  = get_String(rec, dm.ContactDeals_RefNo_sql, "")
	   recItem.NoteId  = get_Int(rec, dm.ContactDeals_NoteId_sql, "0")
	   recItem.RecordState  = get_Int(rec, dm.ContactDeals_RecordState_sql, "0")
	
	// If there are fields below, create the methods in adaptor\ContactDeals_impl.go
	
	
	
	
	// 
	// Dynamically generated 24/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func ContactDeals_NewID(r dm.ContactDeals) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

