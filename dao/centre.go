package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/centre.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Centre (centre)
// Endpoint 	        : Centre (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 15:44:00
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"log"
	"fmt"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	das  "github.com/mt1976/mwt-go-dev/das"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
)

// Centre_GetList() returns a list of all Centre records
func Centre_GetList() (int, []dm.Centre, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Centre_SQLTable)
	count, centreList, _, _ := centre_Fetch(tsql)
	return count, centreList, nil
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
func Centre_Store(r dm.Centre) error {

	logs.Warning(fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = Centre_NewID(r)
	}



	adaptor.Centre_Delete(r.Code)
	adaptor.Centre_Update(r)


	return nil
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
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
    recItem.AppInternalID = get_String(rec, dm.Centre_Code,"")
   recItem.Code  = get_String(rec, dm.Centre_Code, "")
   recItem.Name  = get_String(rec, dm.Centre_Name, "")
   recItem.Country  = get_String(rec, dm.Centre_Country, "")

// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
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

