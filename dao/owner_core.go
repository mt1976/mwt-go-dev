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
// Date & Time		    : 12/12/2021 at 16:13:17
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

	return 1, ownerItem, nil
}

// Owner_GetByReverseLookup(id string) returns a single Owner record
func Owner_GetByReverseLookup(id string) (int, dm.Owner, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Owner_SQLTable)
	tsql = tsql + " WHERE FullName = '" + id + "'"

	_, _, ownerItem, _ := owner_Fetch(tsql)
	
	return 1, ownerItem, nil
}

// Owner_DeleteByID() deletes a single Owner record
func Owner_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Owner_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Owner_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Owner_Store() saves/stores a Owner record to the database
func Owner_Store(r dm.Owner,req *http.Request) error {

	err := owner_Save(r,Audit_User(req))

	return err
}

// Owner_StoreSystem() saves/stores a Owner record to the database
func Owner_StoreSystem(r dm.Owner) error {
	
	err := owner_Save(r,Audit_Host())

	return err
}

// owner_Save() saves/stores a Owner record to the database
func owner_Save(r dm.Owner,usr string) error {

    var err error

	logs.Storing("Owner",fmt.Sprintf("%s", r))

	if len(r.UserName) == 0 {
		r.UserName = Owner_NewID(r)
	}


// Please Create Functions Below in the adaptor/Owner_impl.go file
	err1 := adaptor.Owner_Delete_Impl(r.UserName)
	err2 := adaptor.Owner_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// owner_Fetch read all employees
func owner_Fetch(tsql string) (int, []dm.Owner, dm.Owner, error) {

	var recItem dm.Owner
	var recList []dm.Owner

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.UserName  = get_String(rec, dm.Owner_UserName, "")
   recItem.FullName  = get_String(rec, dm.Owner_FullName, "")
   recItem.Type  = get_String(rec, dm.Owner_Type, "")
   recItem.TradingEntity  = get_String(rec, dm.Owner_TradingEntity, "")
   recItem.DefaultEnterBook  = get_String(rec, dm.Owner_DefaultEnterBook, "")
   recItem.EmailAddress  = get_String(rec, dm.Owner_EmailAddress, "")
   recItem.Enabled  = get_String(rec, dm.Owner_Enabled, "")
   recItem.ExternalUserIds  = get_String(rec, dm.Owner_ExternalUserIds, "")
   recItem.Language  = get_String(rec, dm.Owner_Language, "")
   recItem.LocalCurrency  = get_String(rec, dm.Owner_LocalCurrency, "")
   recItem.Role  = get_String(rec, dm.Owner_Role, "")
   recItem.TelephoneNumber  = get_String(rec, dm.Owner_TelephoneNumber, "")
   recItem.TokenId  = get_String(rec, dm.Owner_TokenId, "")
   recItem.Entity  = get_String(rec, dm.Owner_Entity, "")
   recItem.UserCode  = get_String(rec, dm.Owner_UserCode, "")
// If there are fields below, create the methods in dao\Owner_Impl.go































	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Owner_NewID(r dm.Owner) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

