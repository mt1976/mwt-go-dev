package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/counterpartyaddress.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CounterpartyAddress (counterpartyaddress)
// Endpoint 	        : CounterpartyAddress (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 04/12/2021 at 17:36:43
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
	return 1, counterpartyaddressItem, nil
}



// CounterpartyAddress_DeleteByID() deletes a single CounterpartyAddress record
func CounterpartyAddress_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.CounterpartyAddress_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.CounterpartyAddress_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}


// CounterpartyAddress_Store() saves/stores a CounterpartyAddress record to the database
func CounterpartyAddress_Store(r dm.CounterpartyAddress) error {

	logs.Storing("CounterpartyAddress",fmt.Sprintf("%s", r))

	if len(r.CompID) == 0 {
		r.CompID = CounterpartyAddress_NewID(r)
	}



	adaptor.CounterpartyAddress_Delete(r.CompID)
	adaptor.CounterpartyAddress_Update(r)



	return nil

}

// counterpartyaddress_Fetch read all employees
func counterpartyaddress_Fetch(tsql string) (int, []dm.CounterpartyAddress, dm.CounterpartyAddress, error) {

	var recItem dm.CounterpartyAddress
	var recList []dm.CounterpartyAddress

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 04/12/2021 by matttownsend on silicon.local - START
   recItem.NameFirm  = get_String(rec, dm.CounterpartyAddress_NameFirm, "")
   recItem.NameCentre  = get_String(rec, dm.CounterpartyAddress_NameCentre, "")
   recItem.Address1  = get_String(rec, dm.CounterpartyAddress_Address1, "")
   recItem.Address2  = get_String(rec, dm.CounterpartyAddress_Address2, "")
   recItem.CityTown  = get_String(rec, dm.CounterpartyAddress_CityTown, "")
   recItem.County  = get_String(rec, dm.CounterpartyAddress_County, "")
   recItem.Postcode  = get_String(rec, dm.CounterpartyAddress_Postcode, "")
   recItem.CompID  = get_String(rec, dm.CounterpartyAddress_CompID, "")
// Automatically generated 04/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func CounterpartyAddress_NewID(r dm.CounterpartyAddress) string {
	
	
			id := uuid.New().String()

	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

