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
// Date & Time		    : 28/06/2022 at 16:10:54
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

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, mandateItem, nil
}



// Mandate_DeleteByID() deletes a single Mandate record
func Mandate_Delete(id string) {


	adaptor.Mandate_Delete_impl(id)
	
}


// Mandate_Store() saves/stores a Mandate record to the database
func Mandate_Store(r dm.Mandate,req *http.Request) error {

	err, r := Mandate_Validate(r)
	if err == nil {
		err = mandate_Save(r, Audit_User(req))
	} else {
		logs.Information("Mandate_Store()", err.Error())
	}

	return err
}

// Mandate_StoreSystem() saves/stores a Mandate record to the database
func Mandate_StoreSystem(r dm.Mandate) error {
	
	err, r := Mandate_Validate(r)
	if err == nil {
		err = mandate_Save(r, Audit_Host())
	} else {
		logs.Information("Mandate_Store()", err.Error())
	}

	return err
}

// Mandate_Validate() validates for saves/stores a Mandate record to the database
func Mandate_Validate(r dm.Mandate) (error,dm.Mandate) {
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

// mandate_Save() saves/stores a Mandate record to the database
func mandate_Save(r dm.Mandate,usr string) error {

    var err error



	

	if len(r.CompID) == 0 {
		r.CompID = Mandate_NewID(r)
	}

// If there are fields below, create the methods in dao\mandate_impl.go





















	
logs.Storing("Mandate",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Mandate_impl.go file
	err1 := adaptor.Mandate_Delete_impl(r.CompID)
	err2 := adaptor.Mandate_Update_impl(r.CompID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// mandate_Fetch read all Mandate's
func mandate_Fetch(tsql string) (int, []dm.Mandate, dm.Mandate, error) {

	var recItem dm.Mandate
	var recList []dm.Mandate

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.MandatedUserKeyCounterpartyFirm  = get_String(rec, dm.Mandate_MandatedUserKeyCounterpartyFirm_sql, "")
	   recItem.MandatedUserKeyCounterpartyCentre  = get_String(rec, dm.Mandate_MandatedUserKeyCounterpartyCentre_sql, "")
	   recItem.MandatedUserKeyUserName  = get_String(rec, dm.Mandate_MandatedUserKeyUserName_sql, "")
	   recItem.TelephoneNumber  = get_String(rec, dm.Mandate_TelephoneNumber_sql, "")
	   recItem.EmailAddress  = get_String(rec, dm.Mandate_EmailAddress_sql, "")
	   recItem.Active  = get_Bool(rec, dm.Mandate_Active_sql, "True")
	   recItem.FirstName  = get_String(rec, dm.Mandate_FirstName_sql, "")
	   recItem.Surname  = get_String(rec, dm.Mandate_Surname_sql, "")
	   recItem.DateOfBirth  = get_Time(rec, dm.Mandate_DateOfBirth_sql, "")
	   recItem.Postcode  = get_String(rec, dm.Mandate_Postcode_sql, "")
	   recItem.NationalIDNo  = get_String(rec, dm.Mandate_NationalIDNo_sql, "")
	   recItem.PassportNo  = get_String(rec, dm.Mandate_PassportNo_sql, "")
	   recItem.Country  = get_String(rec, dm.Mandate_Country_sql, "")
	   recItem.CountryName  = get_String(rec, dm.Mandate_CountryName_sql, "")
	   recItem.FirmName  = get_String(rec, dm.Mandate_FirmName_sql, "")
	   recItem.CentreName  = get_String(rec, dm.Mandate_CentreName_sql, "")
	   recItem.Notify  = get_Bool(rec, dm.Mandate_Notify_sql, "True")
	   recItem.SystemUser  = get_String(rec, dm.Mandate_SystemUser_sql, "")
	   recItem.CompID  = get_String(rec, dm.Mandate_CompID_sql, "")
	
	// If there are fields below, create the methods in adaptor\Mandate_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func Mandate_NewID(r dm.Mandate) string {
	
			id := uuid.New().String()
	
	return id
}



// mandate_Fetch read all Mandate's
func Mandate_New() (int, []dm.Mandate, dm.Mandate, error) {

	var r = dm.Mandate{}
	var rList []dm.Mandate
	

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}