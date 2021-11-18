package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/country.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			  : Country
// Endpoint Root 	  : Country
// Search QueryString : Code
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 18/11/2021 at 21:34:18
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

// Country_GetList() returns a list of all Country records
func Country_GetList() (int, []dm.Country, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Country_SQLTable)
	count, countryList, _, _ := country_Fetch(tsql)
	return count, countryList, nil
}

// Country_GetByID() returns a single Country record
func Country_GetByID(id string) (int, dm.Country, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Country_SQLTable)
	tsql = tsql + " WHERE " + dm.Country_SQLSearchID + "='" + id + "'"

	_, _, countryItem, _ := country_Fetch(tsql)
	return 1, countryItem, nil
}

// Country_DeleteByID() deletes a single Country record
func Country_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Country_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Country_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}

// Country_Store() saves/stores a Country record to the database
func Country_Store(r dm.Country) error {

	logs.Warning(fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code= country_NewID(r)
	}



	adaptor.Country_Delete(r.Code)
	adaptor.Country_Update(r)


	return nil
}

// country_Fetch read all employees
func country_Fetch(tsql string) (int, []dm.Country, dm.Country, error) {

	var recItem dm.Country
	var recList []dm.Country

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
    recItem.AppInternalID = get_String(rec, dm.Country_Code,"")
   recItem.Code  = get_String(rec, dm.Country_Code, "")
   recItem.Name  = get_String(rec, dm.Country_Name, "")
   recItem.ShortCode  = get_String(rec, dm.Country_ShortCode, "")
   recItem.EU_EEA  = get_Bool(rec, dm.Country_EU_EEA, "True")
   recItem.HolidaysWeekend  = get_String(rec, dm.Country_HolidaysWeekend, "")
// Automatically generated 18/11/2021 by matttownsend on silicon.local - END
		//Post Import Actions

		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func country_NewID(r dm.Country) string {
	id := uuid.New().String()
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

