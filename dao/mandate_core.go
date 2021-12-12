package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/mandate.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Mandate (mandate)
// Endpoint 	        : Mandate (Mandate)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:16
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
func Mandate_Store(r dm.Mandate,req *http.Request) error {

	err := mandate_Save(r,Audit_User(req))

	return err
}

// Mandate_StoreSystem() saves/stores a Mandate record to the database
func Mandate_StoreSystem(r dm.Mandate) error {
	
	err := mandate_Save(r,Audit_Host())

	return err
}

// mandate_Save() saves/stores a Mandate record to the database
func mandate_Save(r dm.Mandate,usr string) error {

    var err error

	logs.Storing("Mandate",fmt.Sprintf("%s", r))

	if len(r.CompID) == 0 {
		r.CompID = Mandate_NewID(r)
	}


// Please Create Functions Below in the adaptor/Mandate_impl.go file
	err1 := adaptor.Mandate_Delete_Impl(r.CompID)
	err2 := adaptor.Mandate_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

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
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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



// If there are fields below, create the methods in dao\Mandate_Impl.go













































	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Mandate_NewID(r dm.Mandate) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

