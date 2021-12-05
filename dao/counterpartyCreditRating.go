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
// Date & Time		    : 05/12/2021 at 17:15:59
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
	return 1, counterpartycreditratingItem, nil
}



// CounterpartyCreditRating_DeleteByID() deletes a single CounterpartyCreditRating record
func CounterpartyCreditRating_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.CounterpartyCreditRating_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.CounterpartyCreditRating_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}


// CounterpartyCreditRating_Store() saves/stores a CounterpartyCreditRating record to the database
func CounterpartyCreditRating_Store(r dm.CounterpartyCreditRating) error {

	logs.Storing("CounterpartyCreditRating",fmt.Sprintf("%s", r))

	if len(r.CompID) == 0 {
		r.CompID = CounterpartyCreditRating_NewID(r)
	}



	adaptor.CounterpartyCreditRating_Delete(r.CompID)
	adaptor.CounterpartyCreditRating_Update(r)



	return nil

}

// counterpartycreditrating_Fetch read all employees
func counterpartycreditrating_Fetch(tsql string) (int, []dm.CounterpartyCreditRating, dm.CounterpartyCreditRating, error) {

	var recItem dm.CounterpartyCreditRating
	var recList []dm.CounterpartyCreditRating

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 05/12/2021 by matttownsend on silicon.local - START
   recItem.NameFirm  = get_String(rec, dm.CounterpartyCreditRating_NameFirm, "")
   recItem.NameCentre  = get_String(rec, dm.CounterpartyCreditRating_NameCentre, "")
   recItem.CreditRatingUsage  = get_String(rec, dm.CounterpartyCreditRating_CreditRatingUsage, "")
   recItem.CreditRatingAgency  = get_String(rec, dm.CounterpartyCreditRating_CreditRatingAgency, "")
   recItem.CreditRatingName  = get_String(rec, dm.CounterpartyCreditRating_CreditRatingName, "")
   recItem.CompID  = get_String(rec, dm.CounterpartyCreditRating_CompID, "")
// Automatically generated 05/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func CounterpartyCreditRating_NewID(r dm.CounterpartyCreditRating) string {
	
	
			id := uuid.New().String()

	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

