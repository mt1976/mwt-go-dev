package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/counterparty.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Counterparty (counterparty)
// Endpoint 	        : Counterparty (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:21
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

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, counterpartyItem, nil
}



// Counterparty_DeleteByID() deletes a single Counterparty record
func Counterparty_Delete(id string) {


	adaptor.Counterparty_Delete_impl(id)
	
}


// Counterparty_Store() saves/stores a Counterparty record to the database
func Counterparty_Store(r dm.Counterparty,req *http.Request) error {

	err, r := Counterparty_Validate(r)
	if err == nil {
		err = counterparty_Save(r, Audit_User(req))
	} else {
		logs.Information("Counterparty_Store()", err.Error())
	}

	return err
}

// Counterparty_StoreSystem() saves/stores a Counterparty record to the database
func Counterparty_StoreSystem(r dm.Counterparty) error {
	
	err, r := Counterparty_Validate(r)
	if err == nil {
		err = counterparty_Save(r, Audit_Host())
	} else {
		logs.Information("Counterparty_Store()", err.Error())
	}

	return err
}

// Counterparty_Validate() validates for saves/stores a Counterparty record to the database
func Counterparty_Validate(r dm.Counterparty) (error,dm.Counterparty) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// counterparty_Save() saves/stores a Counterparty record to the database
func counterparty_Save(r dm.Counterparty,usr string) error {

    var err error



	

	if len(r.CompID) == 0 {
		r.CompID = Counterparty_NewID(r)
	}

// If there are fields below, create the methods in dao\counterparty_impl.go




















	
logs.Storing("Counterparty",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Counterparty_impl.go file
	err1 := adaptor.Counterparty_Delete_impl(r.CompID)
	err2 := adaptor.Counterparty_Update_impl(r.CompID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// counterparty_Fetch read all Counterparty's
func counterparty_Fetch(tsql string) (int, []dm.Counterparty, dm.Counterparty, error) {

	var recItem dm.Counterparty
	var recList []dm.Counterparty

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.NameCentre  = get_String(rec, dm.Counterparty_NameCentre_sql, "")
	   recItem.NameFirm  = get_String(rec, dm.Counterparty_NameFirm_sql, "")
	   recItem.FullName  = get_String(rec, dm.Counterparty_FullName_sql, "")
	   recItem.TelephoneNumber  = get_String(rec, dm.Counterparty_TelephoneNumber_sql, "")
	   recItem.EmailAddress  = get_String(rec, dm.Counterparty_EmailAddress_sql, "")
	   recItem.CustomerType  = get_String(rec, dm.Counterparty_CustomerType_sql, "")
	   recItem.AccountOfficer  = get_String(rec, dm.Counterparty_AccountOfficer_sql, "")
	   recItem.CountryCode  = get_String(rec, dm.Counterparty_CountryCode_sql, "")
	   recItem.SectorCode  = get_String(rec, dm.Counterparty_SectorCode_sql, "")
	   recItem.CpartyGroupName  = get_String(rec, dm.Counterparty_CpartyGroupName_sql, "")
	   recItem.Notes  = get_String(rec, dm.Counterparty_Notes_sql, "")
	   recItem.Owner  = get_String(rec, dm.Counterparty_Owner_sql, "")
	   recItem.Authorised  = get_Bool(rec, dm.Counterparty_Authorised_sql, "True")
	   recItem.NameFirmName  = get_String(rec, dm.Counterparty_NameFirmName_sql, "")
	   recItem.NameCentreName  = get_String(rec, dm.Counterparty_NameCentreName_sql, "")
	   recItem.CountryCodeName  = get_String(rec, dm.Counterparty_CountryCodeName_sql, "")
	   recItem.SectorCodeName  = get_String(rec, dm.Counterparty_SectorCodeName_sql, "")
	   recItem.CompID  = get_String(rec, dm.Counterparty_CompID_sql, "")
	
	// If there are fields below, create the methods in adaptor\Counterparty_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func Counterparty_NewID(r dm.Counterparty) string {
	
			id := uuid.New().String()
	
	return id
}



// counterparty_Fetch read all Counterparty's
func Counterparty_New() (int, []dm.Counterparty, dm.Counterparty, error) {

	var r = dm.Counterparty{}
	var rList []dm.Counterparty
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}