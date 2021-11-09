package dao

import (
	"fmt"
	"log"

	core "github.com/mt1976/mwt-go-dev/core"
	das "github.com/mt1976/mwt-go-dev/das"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// Country_GetList read all employees
func Country_GetList() (int, []dm.Country, error) {

	Country_SQLTable := core.SienaPropertiesDB["schema"] + "." + dm.Country_Table

	tsql := "SELECT * FROM " + Country_SQLTable

	count, sienaCountryList, _, _ := fetchSienaCountryData(tsql)
	return count, sienaCountryList, nil
}

// getSienaCountryList read all employees
func Country_GetByID(id string) (int, dm.Country, error) {

	Country_SQLTable := core.SienaPropertiesDB["schema"] + "." + dm.Country_Table

	tsql := "SELECT * FROM " + Country_SQLTable + " WHERE " + dm.Country_Code + " = '" + id + "'"

	_, _, sienaCountry, _ := fetchSienaCountryData(tsql)
	return 1, sienaCountry, nil
}

// getSienaCountryList read all employees
func Country_Store(updateItem dm.Country) error {

	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(core.SienaPropertiesDB["schema"])
	fmt.Println(updateItem)

	return nil
}

// fetchSienaCountryData read all employees
func fetchSienaCountryData(tsql string) (int, []dm.Country, dm.Country, error) {

	var sienaCountry dm.Country
	var sienaCountryList []dm.Country

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	//log.Println("No of items: ", noitems)
	//log.Println("No of items: ", len(countryList))

	for i := 0; i < noitems; i++ {

		tc := returnList[i]
		//log.Println("countryList", tc)
		//log.Println("tc:", tc[dm.Country_Code])

		sienaCountry.Code = tc[dm.Country_Code].(string)
		sienaCountry.Name = tc[dm.Country_Name].(string)
		sienaCountry.ShortCode = tc[dm.Country_ShortCode].(string)
		sienaCountry.EU_EEA = sienaDBBoolYN(tc[dm.Country_EU_EEA].(bool))

		sienaCountryList = append(sienaCountryList, sienaCountry)

	}

	return noitems, sienaCountryList, sienaCountry, nil
}
