package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/centre.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Centre (centre)
// Endpoint 	        : Centre (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:20
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

// Centre_GetList() returns a list of all Centre records
func Centre_GetList() (int, []dm.Centre, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Centre_SQLTable)
	count, centreList, _, _ := centre_Fetch(tsql)
	
	return count, centreList, nil
}


// Centre_GetLookup() returns a lookup list of all Centre items in lookup format
func Centre_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, centreList, _ := Centre_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: centreList[i].Code, Name: centreList[i].Name})
	}
	return returnList
}


// Centre_GetByID() returns a single Centre record
func Centre_GetByID(id string) (int, dm.Centre, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Centre_SQLTable)
	tsql = tsql + " WHERE " + dm.Centre_SQLSearchID + "='" + id + "'"
	_, _, centreItem, _ := centre_Fetch(tsql)

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, centreItem, nil
}

// Centre_GetByReverseLookup(id string) returns a single Centre record using the Name field as key.
func Centre_GetByReverseLookup(id string) (int, dm.Centre, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Centre_SQLTable)
	tsql = tsql + " WHERE Name = '" + id + "'"

	_, _, centreItem, _ := centre_Fetch(tsql)
	
	return 1, centreItem, nil
}

// Centre_DeleteByID() deletes a single Centre record
func Centre_Delete(id string) {


	adaptor.Centre_Delete_impl(id)
	
}


// Centre_Store() saves/stores a Centre record to the database
func Centre_Store(r dm.Centre,req *http.Request) error {

	err, r := Centre_Validate(r)
	if err == nil {
		err = centre_Save(r, Audit_User(req))
	} else {
		logs.Information("Centre_Store()", err.Error())
	}

	return err
}

// Centre_StoreSystem() saves/stores a Centre record to the database
func Centre_StoreSystem(r dm.Centre) error {
	
	err, r := Centre_Validate(r)
	if err == nil {
		err = centre_Save(r, Audit_Host())
	} else {
		logs.Information("Centre_Store()", err.Error())
	}

	return err
}

// Centre_Validate() validates for saves/stores a Centre record to the database
func Centre_Validate(r dm.Centre) (error,dm.Centre) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// centre_Save() saves/stores a Centre record to the database
func centre_Save(r dm.Centre,usr string) error {

    var err error



	

	if len(r.Code) == 0 {
		r.Code = Centre_NewID(r)
	}

// If there are fields below, create the methods in dao\centre_impl.go





	
logs.Storing("Centre",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Centre_impl.go file
	err1 := adaptor.Centre_Delete_impl(r.Code)
	err2 := adaptor.Centre_Update_impl(r.Code,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// centre_Fetch read all Centre's
func centre_Fetch(tsql string) (int, []dm.Centre, dm.Centre, error) {

	var recItem dm.Centre
	var recList []dm.Centre

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.Code  = get_String(rec, dm.Centre_Code_sql, "")
	   recItem.Name  = get_String(rec, dm.Centre_Name_sql, "")
	   recItem.Country  = get_String(rec, dm.Centre_Country_sql, "")
	
	// If there are fields below, create the methods in adaptor\Centre_impl.go
	
	
	
	
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
	


func Centre_NewID(r dm.Centre) string {
	
			id := uuid.New().String()
	
	return id
}



// centre_Fetch read all Centre's
func Centre_New() (int, []dm.Centre, dm.Centre, error) {

	var r = dm.Centre{}
	var rList []dm.Centre
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}