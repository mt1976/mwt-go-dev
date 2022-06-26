package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/counterpartyaddress.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CounterpartyAddress (counterpartyaddress)
// Endpoint 	        : CounterpartyAddress (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:22
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

// CounterpartyAddress_GetList() returns a list of all CounterpartyAddress records
func CounterpartyAddress_GetList() (int, []dm.CounterpartyAddress, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyAddress_SQLTable)
	count, counterpartyaddressList, _, _ := counterpartyaddress_Fetch(tsql)
	
	return count, counterpartyaddressList, nil
}



// CounterpartyAddress_GetByID() returns a single CounterpartyAddress record
func CounterpartyAddress_GetByID(id string) (int, dm.CounterpartyAddress, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyAddress_SQLTable)
	tsql = tsql + " WHERE " + dm.CounterpartyAddress_SQLSearchID + "='" + id + "'"
	_, _, counterpartyaddressItem, _ := counterpartyaddress_Fetch(tsql)

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, counterpartyaddressItem, nil
}



// CounterpartyAddress_DeleteByID() deletes a single CounterpartyAddress record
func CounterpartyAddress_Delete(id string) {


	adaptor.CounterpartyAddress_Delete_impl(id)
	
}


// CounterpartyAddress_Store() saves/stores a CounterpartyAddress record to the database
func CounterpartyAddress_Store(r dm.CounterpartyAddress,req *http.Request) error {

	err, r := CounterpartyAddress_Validate(r)
	if err == nil {
		err = counterpartyaddress_Save(r, Audit_User(req))
	} else {
		logs.Information("CounterpartyAddress_Store()", err.Error())
	}

	return err
}

// CounterpartyAddress_StoreSystem() saves/stores a CounterpartyAddress record to the database
func CounterpartyAddress_StoreSystem(r dm.CounterpartyAddress) error {
	
	err, r := CounterpartyAddress_Validate(r)
	if err == nil {
		err = counterpartyaddress_Save(r, Audit_Host())
	} else {
		logs.Information("CounterpartyAddress_Store()", err.Error())
	}

	return err
}

// CounterpartyAddress_Validate() validates for saves/stores a CounterpartyAddress record to the database
func CounterpartyAddress_Validate(r dm.CounterpartyAddress) (error,dm.CounterpartyAddress) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// counterpartyaddress_Save() saves/stores a CounterpartyAddress record to the database
func counterpartyaddress_Save(r dm.CounterpartyAddress,usr string) error {

    var err error



	

	if len(r.CompID) == 0 {
		r.CompID = CounterpartyAddress_NewID(r)
	}

// If there are fields below, create the methods in dao\counterpartyaddress_impl.go










	
logs.Storing("CounterpartyAddress",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/CounterpartyAddress_impl.go file
	err1 := adaptor.CounterpartyAddress_Delete_impl(r.CompID)
	err2 := adaptor.CounterpartyAddress_Update_impl(r.CompID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// counterpartyaddress_Fetch read all CounterpartyAddress's
func counterpartyaddress_Fetch(tsql string) (int, []dm.CounterpartyAddress, dm.CounterpartyAddress, error) {

	var recItem dm.CounterpartyAddress
	var recList []dm.CounterpartyAddress

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.NameFirm  = get_String(rec, dm.CounterpartyAddress_NameFirm_sql, "")
	   recItem.NameCentre  = get_String(rec, dm.CounterpartyAddress_NameCentre_sql, "")
	   recItem.Address1  = get_String(rec, dm.CounterpartyAddress_Address1_sql, "")
	   recItem.Address2  = get_String(rec, dm.CounterpartyAddress_Address2_sql, "")
	   recItem.CityTown  = get_String(rec, dm.CounterpartyAddress_CityTown_sql, "")
	   recItem.County  = get_String(rec, dm.CounterpartyAddress_County_sql, "")
	   recItem.Postcode  = get_String(rec, dm.CounterpartyAddress_Postcode_sql, "")
	   recItem.CompID  = get_String(rec, dm.CounterpartyAddress_CompID_sql, "")
	
	// If there are fields below, create the methods in adaptor\CounterpartyAddress_impl.go
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func CounterpartyAddress_NewID(r dm.CounterpartyAddress) string {
	
			id := uuid.New().String()
	
	return id
}



// counterpartyaddress_Fetch read all CounterpartyAddress's
func CounterpartyAddress_New() (int, []dm.CounterpartyAddress, dm.CounterpartyAddress, error) {

	var r = dm.CounterpartyAddress{}
	var rList []dm.CounterpartyAddress
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}