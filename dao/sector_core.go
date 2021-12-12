package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/sector.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Sector (sector)
// Endpoint 	        : Sector (Sector)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:19
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

// Sector_GetList() returns a list of all Sector records
func Sector_GetList() (int, []dm.Sector, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Sector_SQLTable)
	count, sectorList, _, _ := sector_Fetch(tsql)
	
	return count, sectorList, nil
}


// Sector_GetLookup() returns a lookup list of all Sector items in lookup format
func Sector_GetLookup() []dm.Lookup_Item {

	var returnList []dm.Lookup_Item

	
	    count, sectorList, _ := Sector_GetList()
	
	
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: sectorList[i].Code, Name: sectorList[i].Name})
	}
	return returnList
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
func Sector_Store(r dm.Sector,req *http.Request) error {

	err := sector_Save(r,Audit_User(req))

	return err
}

// Sector_StoreSystem() saves/stores a Sector record to the database
func Sector_StoreSystem(r dm.Sector) error {
	
	err := sector_Save(r,Audit_Host())

	return err
}

// sector_Save() saves/stores a Sector record to the database
func sector_Save(r dm.Sector,usr string) error {

    var err error

	logs.Storing("Sector",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = Sector_NewID(r)
	}


// Please Create Functions Below in the adaptor/Sector_impl.go file
	err1 := adaptor.Sector_Delete_Impl(r.Code)
	err2 := adaptor.Sector_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

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
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.Code  = get_String(rec, dm.Sector_Code, "")
   recItem.Name  = get_String(rec, dm.Sector_Name, "")
// If there are fields below, create the methods in dao\Sector_Impl.go





	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
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

