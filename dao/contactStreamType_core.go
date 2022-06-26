package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/contactstreamtype.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : ContactStreamType (contactstreamtype)
// Endpoint 	        : ContactStreamType (ID)
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

// ContactStreamType_GetList() returns a list of all ContactStreamType records
func ContactStreamType_GetList() (int, []dm.ContactStreamType, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.ContactStreamType_SQLTable)
	count, contactstreamtypeList, _, _ := contactstreamtype_Fetch(tsql)
	
	return count, contactstreamtypeList, nil
}


// ContactStreamType_GetLookup() returns a lookup list of all ContactStreamType items in lookup format
func ContactStreamType_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, contactstreamtypeList, _ := ContactStreamType_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: contactstreamtypeList[i].TypeId, Name: contactstreamtypeList[i].Description})
	}
	return returnList
}


// ContactStreamType_GetByID() returns a single ContactStreamType record
func ContactStreamType_GetByID(id string) (int, dm.ContactStreamType, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.ContactStreamType_SQLTable)
	tsql = tsql + " WHERE " + dm.ContactStreamType_SQLSearchID + "='" + id + "'"
	_, _, contactstreamtypeItem, _ := contactstreamtype_Fetch(tsql)

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, contactstreamtypeItem, nil
}



// ContactStreamType_DeleteByID() deletes a single ContactStreamType record
func ContactStreamType_Delete(id string) {


	adaptor.ContactStreamType_Delete_impl(id)
	
}


// ContactStreamType_Store() saves/stores a ContactStreamType record to the database
func ContactStreamType_Store(r dm.ContactStreamType,req *http.Request) error {

	err, r := ContactStreamType_Validate(r)
	if err == nil {
		err = contactstreamtype_Save(r, Audit_User(req))
	} else {
		logs.Information("ContactStreamType_Store()", err.Error())
	}

	return err
}

// ContactStreamType_StoreSystem() saves/stores a ContactStreamType record to the database
func ContactStreamType_StoreSystem(r dm.ContactStreamType) error {
	
	err, r := ContactStreamType_Validate(r)
	if err == nil {
		err = contactstreamtype_Save(r, Audit_Host())
	} else {
		logs.Information("ContactStreamType_Store()", err.Error())
	}

	return err
}

// ContactStreamType_Validate() validates for saves/stores a ContactStreamType record to the database
func ContactStreamType_Validate(r dm.ContactStreamType) (error,dm.ContactStreamType) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// contactstreamtype_Save() saves/stores a ContactStreamType record to the database
func contactstreamtype_Save(r dm.ContactStreamType,usr string) error {

    var err error



	

	if len(r.TypeId) == 0 {
		r.TypeId = ContactStreamType_NewID(r)
	}

// If there are fields below, create the methods in dao\contactstreamtype_impl.go





	
logs.Storing("ContactStreamType",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/ContactStreamType_impl.go file
	err1 := adaptor.ContactStreamType_Delete_impl(r.TypeId)
	err2 := adaptor.ContactStreamType_Update_impl(r.TypeId,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// contactstreamtype_Fetch read all ContactStreamType's
func contactstreamtype_Fetch(tsql string) (int, []dm.ContactStreamType, dm.ContactStreamType, error) {

	var recItem dm.ContactStreamType
	var recList []dm.ContactStreamType

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.TypeId  = get_Int(rec, dm.ContactStreamType_TypeId_sql, "0")
	   recItem.Description  = get_String(rec, dm.ContactStreamType_Description_sql, "")
	   recItem.RecordState  = get_Int(rec, dm.ContactStreamType_RecordState_sql, "0")
	
	// If there are fields below, create the methods in adaptor\ContactStreamType_impl.go
	
	
	
	
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
	


func ContactStreamType_NewID(r dm.ContactStreamType) string {
	
			id := uuid.New().String()
	
	return id
}



// contactstreamtype_Fetch read all ContactStreamType's
func ContactStreamType_New() (int, []dm.ContactStreamType, dm.ContactStreamType, error) {

	var r = dm.ContactStreamType{}
	var rList []dm.ContactStreamType
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}