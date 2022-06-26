package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/owner.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Owner (owner)
// Endpoint 	        : Owner (Owner)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:30
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

// Owner_GetList() returns a list of all Owner records
func Owner_GetList() (int, []dm.Owner, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Owner_SQLTable)
	count, ownerList, _, _ := owner_Fetch(tsql)
	
	return count, ownerList, nil
}


// Owner_GetLookup() returns a lookup list of all Owner items in lookup format
func Owner_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, ownerList, _ := Owner_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: ownerList[i].UserName, Name: ownerList[i].FullName})
	}
	return returnList
}


// Owner_GetByID() returns a single Owner record
func Owner_GetByID(id string) (int, dm.Owner, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Owner_SQLTable)
	tsql = tsql + " WHERE " + dm.Owner_SQLSearchID + "='" + id + "'"
	_, _, ownerItem, _ := owner_Fetch(tsql)

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, ownerItem, nil
}

// Owner_GetByReverseLookup(id string) returns a single Owner record using the FullName field as key.
func Owner_GetByReverseLookup(id string) (int, dm.Owner, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Owner_SQLTable)
	tsql = tsql + " WHERE FullName = '" + id + "'"

	_, _, ownerItem, _ := owner_Fetch(tsql)
	
	return 1, ownerItem, nil
}

// Owner_DeleteByID() deletes a single Owner record
func Owner_Delete(id string) {


	adaptor.Owner_Delete_impl(id)
	
}


// Owner_Store() saves/stores a Owner record to the database
func Owner_Store(r dm.Owner,req *http.Request) error {

	err, r := Owner_Validate(r)
	if err == nil {
		err = owner_Save(r, Audit_User(req))
	} else {
		logs.Information("Owner_Store()", err.Error())
	}

	return err
}

// Owner_StoreSystem() saves/stores a Owner record to the database
func Owner_StoreSystem(r dm.Owner) error {
	
	err, r := Owner_Validate(r)
	if err == nil {
		err = owner_Save(r, Audit_Host())
	} else {
		logs.Information("Owner_Store()", err.Error())
	}

	return err
}

// Owner_Validate() validates for saves/stores a Owner record to the database
func Owner_Validate(r dm.Owner) (error,dm.Owner) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// owner_Save() saves/stores a Owner record to the database
func owner_Save(r dm.Owner,usr string) error {

    var err error



	

	if len(r.UserName) == 0 {
		r.UserName = Owner_NewID(r)
	}

// If there are fields below, create the methods in dao\owner_impl.go

















	
logs.Storing("Owner",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Owner_impl.go file
	err1 := adaptor.Owner_Delete_impl(r.UserName)
	err2 := adaptor.Owner_Update_impl(r.UserName,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// owner_Fetch read all Owner's
func owner_Fetch(tsql string) (int, []dm.Owner, dm.Owner, error) {

	var recItem dm.Owner
	var recList []dm.Owner

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.UserName  = get_String(rec, dm.Owner_UserName_sql, "")
	   recItem.FullName  = get_String(rec, dm.Owner_FullName_sql, "")
	   recItem.Type  = get_String(rec, dm.Owner_Type_sql, "")
	   recItem.TradingEntity  = get_String(rec, dm.Owner_TradingEntity_sql, "")
	   recItem.DefaultEnterBook  = get_String(rec, dm.Owner_DefaultEnterBook_sql, "")
	   recItem.EmailAddress  = get_String(rec, dm.Owner_EmailAddress_sql, "")
	   recItem.Enabled  = get_String(rec, dm.Owner_Enabled_sql, "")
	   recItem.ExternalUserIds  = get_String(rec, dm.Owner_ExternalUserIds_sql, "")
	   recItem.Language  = get_String(rec, dm.Owner_Language_sql, "")
	   recItem.LocalCurrency  = get_String(rec, dm.Owner_LocalCurrency_sql, "")
	   recItem.Role  = get_String(rec, dm.Owner_Role_sql, "")
	   recItem.TelephoneNumber  = get_String(rec, dm.Owner_TelephoneNumber_sql, "")
	   recItem.TokenId  = get_String(rec, dm.Owner_TokenId_sql, "")
	   recItem.Entity  = get_String(rec, dm.Owner_Entity_sql, "")
	   recItem.UserCode  = get_String(rec, dm.Owner_UserCode_sql, "")
	
	// If there are fields below, create the methods in adaptor\Owner_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func Owner_NewID(r dm.Owner) string {
	
			id := uuid.New().String()
	
	return id
}



// owner_Fetch read all Owner's
func Owner_New() (int, []dm.Owner, dm.Owner, error) {

	var r = dm.Owner{}
	var rList []dm.Owner
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}