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
// Date & Time		    : 12/12/2021 at 16:13:09
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"log"
	
	"fmt"
	"net/http"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
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
func CounterpartyAddress_Store(r dm.CounterpartyAddress,req *http.Request) error {

	err := counterpartyaddress_Save(r,Audit_User(req))

	return err
}

// CounterpartyAddress_StoreSystem() saves/stores a CounterpartyAddress record to the database
func CounterpartyAddress_StoreSystem(r dm.CounterpartyAddress) error {
	
	err := counterpartyaddress_Save(r,Audit_Host())

	return err
}

// counterpartyaddress_Save() saves/stores a CounterpartyAddress record to the database
func counterpartyaddress_Save(r dm.CounterpartyAddress,usr string) error {

    var err error

	logs.Storing("CounterpartyAddress",fmt.Sprintf("%s", r))

	if len(r.CompID) == 0 {
		r.CompID = CounterpartyAddress_NewID(r)
	}


// Please Create Functions Below in the adaptor/CounterpartyAddress_impl.go file
	err1 := adaptor.CounterpartyAddress_Delete_Impl(r.CompID)
	err2 := adaptor.CounterpartyAddress_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

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
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.NameFirm  = get_String(rec, dm.CounterpartyAddress_NameFirm, "")
   recItem.NameCentre  = get_String(rec, dm.CounterpartyAddress_NameCentre, "")
   recItem.Address1  = get_String(rec, dm.CounterpartyAddress_Address1, "")
   recItem.Address2  = get_String(rec, dm.CounterpartyAddress_Address2, "")
   recItem.CityTown  = get_String(rec, dm.CounterpartyAddress_CityTown, "")
   recItem.County  = get_String(rec, dm.CounterpartyAddress_County, "")
   recItem.Postcode  = get_String(rec, dm.CounterpartyAddress_Postcode, "")
   recItem.CompID  = get_String(rec, dm.CounterpartyAddress_CompID, "")
// If there are fields below, create the methods in dao\CounterpartyAddress_Impl.go

















	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
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

