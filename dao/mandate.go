package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/mandate.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Mandate (mandate)
// Endpoint 	        : Mandate (Mandate)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 03/12/2021 at 13:16:59
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"log"
	"fmt"

	
	core "github.com/mt1976/mwt-go-dev/core"
	das  "github.com/mt1976/mwt-go-dev/das"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
)

// Mandate_GetList() returns a list of all Mandate records
func Mandate_GetList() (int, []dm.Mandate, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Mandate_SQLTable)
	count, mandateList, _, _ := mandate_Fetch(tsql)
	return count, mandateList, nil
}



// Mandate_GetByID() returns a single Mandate record
func Mandate_GetByID(id string) (int, dm.Mandate, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Mandate_SQLTable)
	tsql = tsql + " WHERE " + dm.Mandate_SQLSearchID + "='" + id + "'"

	_, _, mandateItem, _ := mandate_Fetch(tsql)
	return 1, mandateItem, nil
}



// Mandate_DeleteByID() deletes a single Mandate record
func Mandate_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Mandate_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Mandate_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}


// Mandate_Store() saves/stores a Mandate record to the database
func Mandate_Store(r dm.Mandate) error {

	logs.Storing("Mandate",fmt.Sprintf("%s", r))

	if len(r.CompID) == 0 {
		r.CompID = Mandate_NewID(r)
	}



	adaptor.Mandate_Delete(r.CompID)
	adaptor.Mandate_Update(r)



	return nil

}

// mandate_Fetch read all employees
func mandate_Fetch(tsql string) (int, []dm.Mandate, dm.Mandate, error) {

	var recItem dm.Mandate
	var recList []dm.Mandate

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - START
   recItem.MandatedUserKeyCounterpartyFirm  = get_String(rec, dm.Mandate_MandatedUserKeyCounterpartyFirm, "")
   recItem.MandatedUserKeyCounterpartyCentre  = get_String(rec, dm.Mandate_MandatedUserKeyCounterpartyCentre, "")
   recItem.MandatedUserKeyUserName  = get_String(rec, dm.Mandate_MandatedUserKeyUserName, "")
   recItem.TelephoneNumber  = get_String(rec, dm.Mandate_TelephoneNumber, "")
   recItem.EmailAddress  = get_String(rec, dm.Mandate_EmailAddress, "")
   recItem.Active  = get_Bool(rec, dm.Mandate_Active, "True")
   recItem.FirstName  = get_String(rec, dm.Mandate_FirstName, "")
   recItem.Surname  = get_String(rec, dm.Mandate_Surname, "")
   recItem.DateOfBirth  = get_Time(rec, dm.Mandate_DateOfBirth, "")
   recItem.Postcode  = get_String(rec, dm.Mandate_Postcode, "")
   recItem.NationalIDNo  = get_String(rec, dm.Mandate_NationalIDNo, "")
   recItem.PassportNo  = get_String(rec, dm.Mandate_PassportNo, "")
   recItem.Country  = get_String(rec, dm.Mandate_Country, "")
   recItem.CountryName  = get_String(rec, dm.Mandate_CountryName, "")
   recItem.FirmName  = get_String(rec, dm.Mandate_FirmName, "")
   recItem.CentreName  = get_String(rec, dm.Mandate_CentreName, "")
   recItem.Notify  = get_Bool(rec, dm.Mandate_Notify, "True")
   recItem.SystemUser  = get_String(rec, dm.Mandate_SystemUser, "")
   recItem.CompID  = get_String(rec, dm.Mandate_CompID, "")



// Automatically generated 03/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Mandate_NewID(r dm.Mandate) string {
	
	

		// mandate_NewIDImpl should be specified in dao/Mandate_Impl.go
		// to provide the implementation for the special case.
		// override should return id - override function should be defined as
		// mandate_NewIDImpl(r dm.Mandate) string {...}
		//
		id := mandate_NewIDImpl(r)
	
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

