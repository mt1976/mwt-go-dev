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
// Date & Time		    : 12/12/2021 at 16:13:08
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"log"
	
	"fmt"
	"net/http"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
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

	return 1, centreItem, nil
}

// Centre_GetByReverseLookup(id string) returns a single Centre record
func Centre_GetByReverseLookup(id string) (int, dm.Centre, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Centre_SQLTable)
	tsql = tsql + " WHERE Name = '" + id + "'"

	_, _, centreItem, _ := centre_Fetch(tsql)
	
	return 1, centreItem, nil
}

// Centre_DeleteByID() deletes a single Centre record
func Centre_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Centre_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Centre_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Centre_Store() saves/stores a Centre record to the database
func Centre_Store(r dm.Centre,req *http.Request) error {

	err := centre_Save(r,Audit_User(req))

	return err
}

// Centre_StoreSystem() saves/stores a Centre record to the database
func Centre_StoreSystem(r dm.Centre) error {
	
	err := centre_Save(r,Audit_Host())

	return err
}

// centre_Save() saves/stores a Centre record to the database
func centre_Save(r dm.Centre,usr string) error {

    var err error

	logs.Storing("Centre",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = Centre_NewID(r)
	}


// Please Create Functions Below in the adaptor/Centre_impl.go file
	err1 := adaptor.Centre_Delete_Impl(r.Code)
	err2 := adaptor.Centre_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// centre_Fetch read all employees
func centre_Fetch(tsql string) (int, []dm.Centre, dm.Centre, error) {

	var recItem dm.Centre
	var recList []dm.Centre

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.Code  = get_String(rec, dm.Centre_Code, "")
   recItem.Name  = get_String(rec, dm.Centre_Name, "")
   recItem.Country  = get_String(rec, dm.Centre_Country, "")

// If there are fields below, create the methods in dao\Centre_Impl.go









	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Centre_NewID(r dm.Centre) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

