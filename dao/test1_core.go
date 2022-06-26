package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/test1.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Test1 (test1)
// Endpoint 	        : Test1 (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:33
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"

	adaptor "github.com/mt1976/mwt-go-dev/adaptor"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

// Test1_GetList() returns a list of all Test1 records
func Test1_GetList() (int, []dm.Test1, error) {

	count, test1List, _ := adaptor.Test1_GetList_impl()

	return count, test1List, nil
}

// Test1_GetLookup() returns a lookup list of all Test1 items in lookup format
func Test1_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, test1List, _ := Test1_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: test1List[i].ID, Name: test1List[i].Descr})
	}
	return returnList
}

// Test1_GetByID() returns a single Test1 record
func Test1_GetByID(id string) (int, dm.Test1, error) {

	_, test1Item, _ := adaptor.Test1_GetByID_impl(id)

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	test1Item.Cheese, test1Item.Cheese_props = adaptor.Test1_Cheese_Impl(adaptor.GET, id, test1Item.Cheese, test1Item, test1Item.Cheese_props)
	test1Item.Onion, test1Item.Onion_props = adaptor.Test1_Onion_Impl(adaptor.GET, id, test1Item.Onion, test1Item, test1Item.Onion_props)
	//
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, test1Item, nil
}

// Test1_GetByReverseLookup(id string) returns a single Test1 record using the name_fisherRLU field as key.
func Test1_GetByReverseLookup(id string) (int, dm.Test1, error) {

	_, test1Item, _ := adaptor.Test1_GetByReverseLookup_impl(id, "name_fisherRLU")

	return 1, test1Item, nil
}

// Test1_DeleteByID() deletes a single Test1 record
func Test1_Delete(id string) {

	adaptor.Test1_Delete_impl(id)

}

// Test1_Store() saves/stores a Test1 record to the database
func Test1_Store(r dm.Test1, req *http.Request) error {

	err, r := Test1_Validate(r)
	if err == nil {
		err = test1_Save(r, Audit_User(req))
	} else {
		logs.Information("Test1_Store()", err.Error())
	}

	return err
}

// Test1_StoreSystem() saves/stores a Test1 record to the database
func Test1_StoreSystem(r dm.Test1) error {

	err, r := Test1_Validate(r)
	if err == nil {
		err = test1_Save(r, Audit_Host())
	} else {
		logs.Information("Test1_Store()", err.Error())
	}

	return err
}

// Test1_Validate() validates for saves/stores a Test1 record to the database
func Test1_Validate(r dm.Test1) (error, dm.Test1) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	r.Cheese, r.Cheese_props = adaptor.Test1_Cheese_Impl(adaptor.PUT, r.ID, r.Cheese, r, r.Cheese_props)
	r.Onion, r.Onion_props = adaptor.Test1_Onion_Impl(adaptor.PUT, r.ID, r.Onion, r, r.Onion_props)
	//
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return nil, r
}

//

// test1_Save() saves/stores a Test1 record to the database
func test1_Save(r dm.Test1, usr string) error {

	var err error

	if len(r.ID) == 0 {
		r.ID = Test1_NewID(r)
	}

	// If there are fields below, create the methods in dao\test1_impl.go

	r.Cheese, err = adaptor.Test1_Cheese_OnStore_impl(r.Cheese, r, usr)
	r.Onion, err = adaptor.Test1_Onion_OnStore_impl(r.Onion, r, usr)

	logs.Storing("Test1", fmt.Sprintf("%s", r))

	// Please Create Functions Below in the adaptor/Test1_impl.go file
	err1 := adaptor.Test1_Delete_impl(r.ID)
	err2 := adaptor.Test1_Update_impl(r.ID, r, usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}

	return err

}

// test1_Fetch is not required as GetByID, GetAll etc... have been diverted to _impl

func Test1_NewID(r dm.Test1) string {

	id := uuid.New().String()

	return id
}

// test1_Fetch read all Test1's
func Test1_New() (int, []dm.Test1, dm.Test1, error) {

	var r = dm.Test1{}
	var rList []dm.Test1

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	r.Cheese, r.Cheese_props = adaptor.Test1_Cheese_Impl(adaptor.NEW, r.ID, r.Cheese, r, r.Cheese_props)
	r.Onion, r.Onion_props = adaptor.Test1_Onion_Impl(adaptor.NEW, r.ID, r.Onion, r, r.Onion_props)
	//
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	rList = append(rList, r)

	return 1, rList, r, nil
}
