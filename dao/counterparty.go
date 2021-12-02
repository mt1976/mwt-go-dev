package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/counterparty.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Counterparty (counterparty)
// Endpoint 	        : Counterparty (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 02/12/2021 at 19:40:00
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

// Counterparty_GetList() returns a list of all Counterparty records
func Counterparty_GetList() (int, []dm.Counterparty, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Counterparty_SQLTable)
	count, counterpartyList, _, _ := counterparty_Fetch(tsql)
	return count, counterpartyList, nil
}



// Counterparty_GetByID() returns a single Counterparty record
func Counterparty_GetByID(id string) (int, dm.Counterparty, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Counterparty_SQLTable)
	tsql = tsql + " WHERE " + dm.Counterparty_SQLSearchID + "='" + id + "'"

	_, _, counterpartyItem, _ := counterparty_Fetch(tsql)
	return 1, counterpartyItem, nil
}



// Counterparty_DeleteByID() deletes a single Counterparty record
func Counterparty_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Counterparty_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Counterparty_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}


// Counterparty_Store() saves/stores a Counterparty record to the database
func Counterparty_Store(r dm.Counterparty) error {

	logs.Storing("Counterparty",fmt.Sprintf("%s", r))

	if len(r.CompID) == 0 {
		r.CompID = Counterparty_NewID(r)
	}



	adaptor.Counterparty_Delete(r.CompID)
	adaptor.Counterparty_Update(r)



	return nil

}

// counterparty_Fetch read all employees
func counterparty_Fetch(tsql string) (int, []dm.Counterparty, dm.Counterparty, error) {

	var recItem dm.Counterparty
	var recList []dm.Counterparty

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 02/12/2021 by matttownsend on silicon.local - START
   recItem.NameCentre  = get_String(rec, dm.Counterparty_NameCentre, "")
   recItem.NameFirm  = get_String(rec, dm.Counterparty_NameFirm, "")
   recItem.FullName  = get_String(rec, dm.Counterparty_FullName, "")
   recItem.TelephoneNumber  = get_String(rec, dm.Counterparty_TelephoneNumber, "")
   recItem.EmailAddress  = get_String(rec, dm.Counterparty_EmailAddress, "")
   recItem.CustomerType  = get_String(rec, dm.Counterparty_CustomerType, "")
   recItem.AccountOfficer  = get_String(rec, dm.Counterparty_AccountOfficer, "")
   recItem.CountryCode  = get_String(rec, dm.Counterparty_CountryCode, "")
   recItem.SectorCode  = get_String(rec, dm.Counterparty_SectorCode, "")
   recItem.CpartyGroupName  = get_String(rec, dm.Counterparty_CpartyGroupName, "")
   recItem.Notes  = get_String(rec, dm.Counterparty_Notes, "")
   recItem.Owner  = get_String(rec, dm.Counterparty_Owner, "")
   recItem.Authorised  = get_Bool(rec, dm.Counterparty_Authorised, "True")
   recItem.NameFirmName  = get_String(rec, dm.Counterparty_NameFirmName, "")
   recItem.NameCentreName  = get_String(rec, dm.Counterparty_NameCentreName, "")
   recItem.CountryCodeName  = get_String(rec, dm.Counterparty_CountryCodeName, "")
   recItem.SectorCodeName  = get_String(rec, dm.Counterparty_SectorCodeName, "")
   recItem.CompID  = get_String(rec, dm.Counterparty_CompID, "")
// Automatically generated 02/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Counterparty_NewID(r dm.Counterparty) string {
	
	
			id := uuid.New().String()

	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

