package dao

import (
	"fmt"
	"log"

	core "github.com/mt1976/mwt-go-dev/core"
	"github.com/mt1976/mwt-go-dev/das"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// getSienaCentreList read all employees
func Centre_GetList() (int, []dm.Centre, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Centre_SQLTable)
	count, sienaCentreList, _, _ := fetchSienaCentreData(tsql)
	return count, sienaCentreList, nil
}

// getSienaCentreList read all employees
func Centre_GetByID(id string) (int, dm.Centre, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Centre_SQLTable)
	tsql = tsql + " WHERE " + dm.Centre_Code + "='" + id + "'"

	_, _, sienaCentre, _ := fetchSienaCentreData(tsql)
	return 1, sienaCentre, nil
}

// getSienaCentreList read all employees
func Centre_Store(updateItem dm.Centre) error {

	//fmt.Println(db.Stats().OpenConnections)
	//fmt.Println(core.SienaPropertiesDB["schema"])
	// TODO Implement Store Function for Centres
	fmt.Println(updateItem)
	return nil
}

// fetchSienaCentreData read all employees
func fetchSienaCentreData(tsql string) (int, []dm.Centre, dm.Centre, error) {

	var sienaCentre dm.Centre
	var sienaCentreList []dm.Centre

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	//	spew.Dump(returnList)

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// Translate Siena data into Centre data
		sienaCentre.Code = get_String(rec, dm.Centre_Code, "")
		sienaCentre.Name = get_String(rec, dm.Centre_Name, "")
		sienaCentre.Country = get_String(rec, dm.Centre_Country, "")
		sienaCentre.CountryName = get_String(rec, dm.Centre_CountryName, "")

		//Post Import Actions

		//Add to the list
		sienaCentreList = append(sienaCentreList, sienaCentre)

	}
	return noitems, sienaCentreList, sienaCentre, nil
}
