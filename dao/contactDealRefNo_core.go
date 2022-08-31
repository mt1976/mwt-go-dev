package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/contactdealrefno.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : ContactDealRefNo (contactdealrefno)
// Endpoint 	        : ContactDealRefNo (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:45
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

// ContactDealRefNo_GetList() returns a list of all ContactDealRefNo records
func ContactDealRefNo_GetList() (int, []dm.ContactDealRefNo, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.ContactDealRefNo_SQLTable)
	count, contactdealrefnoList, _, _ := contactdealrefno_Fetch(tsql)
	
	return count, contactdealrefnoList, nil
}


// ContactDealRefNo_GetLookup() returns a lookup list of all ContactDealRefNo items in lookup format
func ContactDealRefNo_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, contactdealrefnoList, _ := ContactDealRefNo_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: contactdealrefnoList[i].NoteId, Name: contactdealrefnoList[i].RefNo})
	}
	return returnList
}


// ContactDealRefNo_GetByID() returns a single ContactDealRefNo record
func ContactDealRefNo_GetByID(id string) (int, dm.ContactDealRefNo, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.ContactDealRefNo_SQLTable)
	tsql = tsql + " WHERE " + dm.ContactDealRefNo_SQLSearchID + "='" + id + "'"
	_, _, contactdealrefnoItem, _ := contactdealrefno_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, contactdealrefnoItem, nil
}



// ContactDealRefNo_DeleteByID() deletes a single ContactDealRefNo record
func ContactDealRefNo_Delete(id string) {


	adaptor.ContactDealRefNo_Delete_impl(id)
	
}


// ContactDealRefNo_Store() saves/stores a ContactDealRefNo record to the database
func ContactDealRefNo_Store(r dm.ContactDealRefNo,req *http.Request) error {

	err, r := ContactDealRefNo_Validate(r)
	if err == nil {
		err = contactdealrefno_Save(r, Audit_User(req))
	} else {
		logs.Information("ContactDealRefNo_Store()", err.Error())
	}

	return err
}

// ContactDealRefNo_StoreSystem() saves/stores a ContactDealRefNo record to the database
func ContactDealRefNo_StoreSystem(r dm.ContactDealRefNo) error {
	
	err, r := ContactDealRefNo_Validate(r)
	if err == nil {
		err = contactdealrefno_Save(r, Audit_Host())
	} else {
		logs.Information("ContactDealRefNo_Store()", err.Error())
	}

	return err
}

// ContactDealRefNo_Validate() validates for saves/stores a ContactDealRefNo record to the database
func ContactDealRefNo_Validate(r dm.ContactDealRefNo) (error,dm.ContactDealRefNo) {
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

// contactdealrefno_Save() saves/stores a ContactDealRefNo record to the database
func contactdealrefno_Save(r dm.ContactDealRefNo,usr string) error {

    var err error



	

	if len(r.NoteId) == 0 {
		r.NoteId = ContactDealRefNo_NewID(r)
	}

// If there are fields below, create the methods in dao\contactdealrefno_impl.go





	
logs.Storing("ContactDealRefNo",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/ContactDealRefNo_impl.go file
	err1 := adaptor.ContactDealRefNo_Delete_impl(r.NoteId)
	err2 := adaptor.ContactDealRefNo_Update_impl(r.NoteId,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// contactdealrefno_Fetch read all ContactDealRefNo's
func contactdealrefno_Fetch(tsql string) (int, []dm.ContactDealRefNo, dm.ContactDealRefNo, error) {

	var recItem dm.ContactDealRefNo
	var recList []dm.ContactDealRefNo

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.RefNo  = get_String(rec, dm.ContactDealRefNo_RefNo_sql, "")
	   recItem.NoteId  = get_Int(rec, dm.ContactDealRefNo_NoteId_sql, "0")
	   recItem.RecordState  = get_Int(rec, dm.ContactDealRefNo_RecordState_sql, "0")
	
	// If there are fields below, create the methods in adaptor\ContactDealRefNo_impl.go
	
	
	
	
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
	


func ContactDealRefNo_NewID(r dm.ContactDealRefNo) string {
	
			id := uuid.New().String()
	
	return id
}



// contactdealrefno_Fetch read all ContactDealRefNo's
func ContactDealRefNo_New() (int, []dm.ContactDealRefNo, dm.ContactDealRefNo, error) {

	var r = dm.ContactDealRefNo{}
	var rList []dm.ContactDealRefNo
	

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}