package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/country.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Country (country)
// Endpoint 	        : Country (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:48
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

// Country_GetList() returns a list of all Country records
func Country_GetList() (int, []dm.Country, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Country_SQLTable)
	count, countryList, _, _ := country_Fetch(tsql)
	
	return count, countryList, nil
}


// Country_GetLookup() returns a lookup list of all Country items in lookup format
func Country_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, countryList, _ := Country_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: countryList[i].Code, Name: countryList[i].Name})
	}
	return returnList
}


// Country_GetByID() returns a single Country record
func Country_GetByID(id string) (int, dm.Country, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Country_SQLTable)
	tsql = tsql + " WHERE " + dm.Country_SQLSearchID + "='" + id + "'"
	_, _, countryItem, _ := country_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, countryItem, nil
}

// Country_GetByReverseLookup(id string) returns a single Country record using the Name field as key.
func Country_GetByReverseLookup(id string) (int, dm.Country, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Country_SQLTable)
	tsql = tsql + " WHERE Name = '" + id + "'"

	_, _, countryItem, _ := country_Fetch(tsql)
	
	return 1, countryItem, nil
}

// Country_DeleteByID() deletes a single Country record
func Country_Delete(id string) {


	adaptor.Country_Delete_impl(id)
	
}


// Country_Store() saves/stores a Country record to the database
func Country_Store(r dm.Country,req *http.Request) error {

	err, r := Country_Validate(r)
	if err == nil {
		err = country_Save(r, Audit_User(req))
	} else {
		logs.Information("Country_Store()", err.Error())
	}

	return err
}

// Country_StoreSystem() saves/stores a Country record to the database
func Country_StoreSystem(r dm.Country) error {
	
	err, r := Country_Validate(r)
	if err == nil {
		err = country_Save(r, Audit_Host())
	} else {
		logs.Information("Country_Store()", err.Error())
	}

	return err
}

// Country_Validate() validates for saves/stores a Country record to the database
func Country_Validate(r dm.Country) (error,dm.Country) {
	var err error
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return err,r
}
//

// country_Save() saves/stores a Country record to the database
func country_Save(r dm.Country,usr string) error {

    var err error



	

	if len(r.Code) == 0 {
		r.Code = Country_NewID(r)
	}

// If there are fields below, create the methods in dao\country_impl.go







	
logs.Storing("Country",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Country_impl.go file
	err1 := adaptor.Country_Delete_impl(r.Code)
	err2 := adaptor.Country_Update_impl(r.Code,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// country_Fetch read all Country's
func country_Fetch(tsql string) (int, []dm.Country, dm.Country, error) {

	var recItem dm.Country
	var recList []dm.Country

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.Code  = get_String(rec, dm.Country_Code_sql, "")
	   recItem.Name  = get_String(rec, dm.Country_Name_sql, "")
	   recItem.ShortCode  = get_String(rec, dm.Country_ShortCode_sql, "")
	   recItem.EU_EEA  = get_Bool(rec, dm.Country_EU_EEA_sql, "True")
	   recItem.HolidaysWeekend  = get_String(rec, dm.Country_HolidaysWeekend_sql, "")
	
	// If there are fields below, create the methods in adaptor\Country_impl.go
	
	
	
	
	
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Country_NewID(r dm.Country) string {
	
			id := uuid.New().String()
	
	return id
}



// country_Fetch read all Country's
func Country_New() (int, []dm.Country, dm.Country, error) {

	var r = dm.Country{}
	var rList []dm.Country
	

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}