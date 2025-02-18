package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/currencypair.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CurrencyPair (currencypair)
// Endpoint 	        : CurrencyPair (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:49
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

// CurrencyPair_GetList() returns a list of all CurrencyPair records
func CurrencyPair_GetList() (int, []dm.CurrencyPair, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CurrencyPair_SQLTable)
	count, currencypairList, _, _ := currencypair_Fetch(tsql)
	
	return count, currencypairList, nil
}


// CurrencyPair_GetLookup() returns a lookup list of all CurrencyPair items in lookup format
func CurrencyPair_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, currencypairList, _ := CurrencyPair_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: currencypairList[i].Code, Name: currencypairList[i].Code})
	}
	return returnList
}


// CurrencyPair_GetByID() returns a single CurrencyPair record
func CurrencyPair_GetByID(id string) (int, dm.CurrencyPair, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CurrencyPair_SQLTable)
	tsql = tsql + " WHERE " + dm.CurrencyPair_SQLSearchID + "='" + id + "'"
	_, _, currencypairItem, _ := currencypair_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	currencypairItem.Code,currencypairItem.Code_props = adaptor.CurrencyPair_Code_impl (adaptor.GET,id,currencypairItem.Code,currencypairItem,currencypairItem.Code_props)
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, currencypairItem, nil
}



// CurrencyPair_DeleteByID() deletes a single CurrencyPair record
func CurrencyPair_Delete(id string) {


	adaptor.CurrencyPair_Delete_impl(id)
	
}


// CurrencyPair_Store() saves/stores a CurrencyPair record to the database
func CurrencyPair_Store(r dm.CurrencyPair,req *http.Request) error {

	err, r := CurrencyPair_Validate(r)
	if err == nil {
		err = currencypair_Save(r, Audit_User(req))
	} else {
		logs.Information("CurrencyPair_Store()", err.Error())
	}

	return err
}

// CurrencyPair_StoreSystem() saves/stores a CurrencyPair record to the database
func CurrencyPair_StoreSystem(r dm.CurrencyPair) error {
	
	err, r := CurrencyPair_Validate(r)
	if err == nil {
		err = currencypair_Save(r, Audit_Host())
	} else {
		logs.Information("CurrencyPair_Store()", err.Error())
	}

	return err
}

// CurrencyPair_Validate() validates for saves/stores a CurrencyPair record to the database
func CurrencyPair_Validate(r dm.CurrencyPair) (error,dm.CurrencyPair) {
	var err error
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.Code,r.Code_props = adaptor.CurrencyPair_Code_impl (adaptor.PUT,r.Code,r.Code,r,r.Code_props)
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return err,r
}
//

// currencypair_Save() saves/stores a CurrencyPair record to the database
func currencypair_Save(r dm.CurrencyPair,usr string) error {

    var err error



	

	if len(r.Code) == 0 {
		r.Code = CurrencyPair_NewID(r)
	}

// If there are fields below, create the methods in dao\currencypair_impl.go



  r.Code,err = adaptor.CurrencyPair_Code_OnStore_impl (r.Code,r,usr)


	
logs.Storing("CurrencyPair",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/CurrencyPair_impl.go file
	err1 := adaptor.CurrencyPair_Delete_impl(r.Code)
	err2 := adaptor.CurrencyPair_Update_impl(r.Code,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// currencypair_Fetch read all CurrencyPair's
func currencypair_Fetch(tsql string) (int, []dm.CurrencyPair, dm.CurrencyPair, error) {

	var recItem dm.CurrencyPair
	var recList []dm.CurrencyPair

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.CodeMajorCurrencyIsoCode  = get_String(rec, dm.CurrencyPair_CodeMajorCurrencyIsoCode_sql, "")
	   recItem.CodeMinorCurrencyIsoCode  = get_String(rec, dm.CurrencyPair_CodeMinorCurrencyIsoCode_sql, "")
	   recItem.ReciprocalActive  = get_Bool(rec, dm.CurrencyPair_ReciprocalActive_sql, "True")
	   recItem.Code  = get_String(rec, dm.CurrencyPair_Code_sql, "")
	
	// If there are fields below, create the methods in adaptor\CurrencyPair_impl.go
	
	
	
	   recItem.Code  = adaptor.CurrencyPair_Code_OnFetch_impl (recItem)
	
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
	


func CurrencyPair_NewID(r dm.CurrencyPair) string {
	
			id := uuid.New().String()
	
	return id
}



// currencypair_Fetch read all CurrencyPair's
func CurrencyPair_New() (int, []dm.CurrencyPair, dm.CurrencyPair, error) {

	var r = dm.CurrencyPair{}
	var rList []dm.CurrencyPair
	

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	r.Code,r.Code_props = adaptor.CurrencyPair_Code_impl (adaptor.NEW,r.Code,r.Code,r,r.Code_props)
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}