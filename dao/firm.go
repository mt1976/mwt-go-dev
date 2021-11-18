package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/firm.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			  : Firm
// Endpoint Root 	  : Firm
// Search QueryString : FirmName
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 18/11/2021 at 21:34:20
// Who & Where		  : matttownsend on silicon.local
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

// Firm_GetList() returns a list of all Firm records
func Firm_GetList() (int, []dm.Firm, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Firm_SQLTable)
	count, firmList, _, _ := firm_Fetch(tsql)
	return count, firmList, nil
}

// Firm_GetByID() returns a single Firm record
func Firm_GetByID(id string) (int, dm.Firm, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Firm_SQLTable)
	tsql = tsql + " WHERE " + dm.Firm_SQLSearchID + "='" + id + "'"

	_, _, firmItem, _ := firm_Fetch(tsql)
	return 1, firmItem, nil
}

// Firm_DeleteByID() deletes a single Firm record
func Firm_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Firm_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Firm_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}

// Firm_Store() saves/stores a Firm record to the database
func Firm_Store(r dm.Firm) error {

	logs.Warning(fmt.Sprintf("%s", r))

	if len(r.FirmName) == 0 {
		r.FirmName= firm_NewID(r)
	}



	adaptor.Firm_Delete(r.FirmName)
	adaptor.Firm_Update(r)


	return nil
}

// firm_Fetch read all employees
func firm_Fetch(tsql string) (int, []dm.Firm, dm.Firm, error) {

	var recItem dm.Firm
	var recList []dm.Firm

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
    recItem.AppInternalID = get_String(rec, dm.Firm_FirmName,"")
   recItem.FirmName  = get_String(rec, dm.Firm_FirmName, "")
   recItem.FullName  = get_String(rec, dm.Firm_FullName, "")
   recItem.Country  = get_String(rec, dm.Firm_Country, "")
   recItem.Sector  = get_String(rec, dm.Firm_Sector, "")
   recItem.SectorName  = get_String(rec, dm.Firm_SectorName, "")
   recItem.CountryName  = get_String(rec, dm.Firm_CountryName, "")
// Automatically generated 18/11/2021 by matttownsend on silicon.local - END
		//Post Import Actions

		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func firm_NewID(r dm.Firm) string {
	id := uuid.New().String()
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

