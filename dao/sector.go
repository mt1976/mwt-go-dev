package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/sector.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Sector (sector)
// Endpoint 	        : Sector (Sector)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 22/11/2021 at 11:25:57
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

// Sector_GetList() returns a list of all Sector records
func Sector_GetList() (int, []dm.Sector, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Sector_SQLTable)
	count, sectorList, _, _ := sector_Fetch(tsql)
	return count, sectorList, nil
}

// Sector_GetByID() returns a single Sector record
func Sector_GetByID(id string) (int, dm.Sector, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Sector_SQLTable)
	tsql = tsql + " WHERE " + dm.Sector_SQLSearchID + "='" + id + "'"

	_, _, sectorItem, _ := sector_Fetch(tsql)
	return 1, sectorItem, nil
}

// Sector_GetByReverseLookup(id string) returns a single Sector record
func Sector_GetByReverseLookup(id string) (int, dm.Sector, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Sector_SQLTable)
	tsql = tsql + " WHERE Name = '" + id + "'"

	_, _, sectorItem, _ := sector_Fetch(tsql)
	return 1, sectorItem, nil
}

// Sector_DeleteByID() deletes a single Sector record
func Sector_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Sector_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Sector_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}

// Sector_Store() saves/stores a Sector record to the database
func Sector_Store(r dm.Sector) error {

	logs.Storing("Sector",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = Sector_NewID(r)
	}



	adaptor.Sector_Delete(r.Code)
	adaptor.Sector_Update(r)


	return nil
}

// sector_Fetch read all employees
func sector_Fetch(tsql string) (int, []dm.Sector, dm.Sector, error) {

	var recItem dm.Sector
	var recList []dm.Sector

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
    recItem.AppInternalID = get_String(rec, dm.Sector_Code,"")
   recItem.Code  = get_String(rec, dm.Sector_Code, "")
   recItem.Name  = get_String(rec, dm.Sector_Name, "")
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Sector_NewID(r dm.Sector) string {
	
	
			id := uuid.New().String()

	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

