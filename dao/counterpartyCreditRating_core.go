package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/counterpartycreditrating.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CounterpartyCreditRating (counterpartycreditrating)
// Endpoint 	        : CounterpartyCreditRating (ID)
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

// CounterpartyCreditRating_GetList() returns a list of all CounterpartyCreditRating records
func CounterpartyCreditRating_GetList() (int, []dm.CounterpartyCreditRating, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyCreditRating_SQLTable)
	count, counterpartycreditratingList, _, _ := counterpartycreditrating_Fetch(tsql)
	
	return count, counterpartycreditratingList, nil
}



// CounterpartyCreditRating_GetByID() returns a single CounterpartyCreditRating record
func CounterpartyCreditRating_GetByID(id string) (int, dm.CounterpartyCreditRating, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyCreditRating_SQLTable)
	tsql = tsql + " WHERE " + dm.CounterpartyCreditRating_SQLSearchID + "='" + id + "'"
	_, _, counterpartycreditratingItem, _ := counterpartycreditrating_Fetch(tsql)

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, counterpartycreditratingItem, nil
}



// CounterpartyCreditRating_DeleteByID() deletes a single CounterpartyCreditRating record
func CounterpartyCreditRating_Delete(id string) {


	adaptor.CounterpartyCreditRating_Delete_impl(id)
	
}


// CounterpartyCreditRating_Store() saves/stores a CounterpartyCreditRating record to the database
func CounterpartyCreditRating_Store(r dm.CounterpartyCreditRating,req *http.Request) error {

	err, r := CounterpartyCreditRating_Validate(r)
	if err == nil {
		err = counterpartycreditrating_Save(r, Audit_User(req))
	} else {
		logs.Information("CounterpartyCreditRating_Store()", err.Error())
	}

	return err
}

// CounterpartyCreditRating_StoreSystem() saves/stores a CounterpartyCreditRating record to the database
func CounterpartyCreditRating_StoreSystem(r dm.CounterpartyCreditRating) error {
	
	err, r := CounterpartyCreditRating_Validate(r)
	if err == nil {
		err = counterpartycreditrating_Save(r, Audit_Host())
	} else {
		logs.Information("CounterpartyCreditRating_Store()", err.Error())
	}

	return err
}

// CounterpartyCreditRating_Validate() validates for saves/stores a CounterpartyCreditRating record to the database
func CounterpartyCreditRating_Validate(r dm.CounterpartyCreditRating) (error,dm.CounterpartyCreditRating) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// counterpartycreditrating_Save() saves/stores a CounterpartyCreditRating record to the database
func counterpartycreditrating_Save(r dm.CounterpartyCreditRating,usr string) error {

    var err error



	

	if len(r.CompID) == 0 {
		r.CompID = CounterpartyCreditRating_NewID(r)
	}

// If there are fields below, create the methods in dao\counterpartycreditrating_impl.go








	
logs.Storing("CounterpartyCreditRating",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/CounterpartyCreditRating_impl.go file
	err1 := adaptor.CounterpartyCreditRating_Delete_impl(r.CompID)
	err2 := adaptor.CounterpartyCreditRating_Update_impl(r.CompID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// counterpartycreditrating_Fetch read all CounterpartyCreditRating's
func counterpartycreditrating_Fetch(tsql string) (int, []dm.CounterpartyCreditRating, dm.CounterpartyCreditRating, error) {

	var recItem dm.CounterpartyCreditRating
	var recList []dm.CounterpartyCreditRating

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.NameFirm  = get_String(rec, dm.CounterpartyCreditRating_NameFirm_sql, "")
	   recItem.NameCentre  = get_String(rec, dm.CounterpartyCreditRating_NameCentre_sql, "")
	   recItem.CreditRatingUsage  = get_String(rec, dm.CounterpartyCreditRating_CreditRatingUsage_sql, "")
	   recItem.CreditRatingAgency  = get_String(rec, dm.CounterpartyCreditRating_CreditRatingAgency_sql, "")
	   recItem.CreditRatingName  = get_String(rec, dm.CounterpartyCreditRating_CreditRatingName_sql, "")
	   recItem.CompID  = get_String(rec, dm.CounterpartyCreditRating_CompID_sql, "")
	
	// If there are fields below, create the methods in adaptor\CounterpartyCreditRating_impl.go
	
	
	
	
	
	
	
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
	


func CounterpartyCreditRating_NewID(r dm.CounterpartyCreditRating) string {
	
			id := uuid.New().String()
	
	return id
}



// counterpartycreditrating_Fetch read all CounterpartyCreditRating's
func CounterpartyCreditRating_New() (int, []dm.CounterpartyCreditRating, dm.CounterpartyCreditRating, error) {

	var r = dm.CounterpartyCreditRating{}
	var rList []dm.CounterpartyCreditRating
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}