package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/catalog.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Catalog (catalog)
// Endpoint 	        : Catalog (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:16:04
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"fmt"
	"net/http"

	
	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
)

// Catalog_GetList() returns a list of all Catalog records
func Catalog_GetList() (int, []dm.Catalog, error) {
	
	count, catalogList, _ := adaptor.Catalog_GetList_Impl()
	
	return count, catalogList, nil
}



// Catalog_GetByID() returns a single Catalog record
func Catalog_GetByID(id string) (int, dm.Catalog, error) {


	 _, catalogItem, _ := adaptor.Catalog_GetByID_Impl(id)
	
	return 1, catalogItem, nil
}



// Catalog_DeleteByID() deletes a single Catalog record
func Catalog_Delete(id string) {


	adaptor.Catalog_Delete_Impl(id)
	
}


// Catalog_Store() saves/stores a Catalog record to the database
func Catalog_Store(r dm.Catalog,req *http.Request) error {

	err := catalog_Save(r,Audit_User(req))

	return err
}

// Catalog_StoreSystem() saves/stores a Catalog record to the database
func Catalog_StoreSystem(r dm.Catalog) error {
	
	err := catalog_Save(r,Audit_Host())

	return err
}

// catalog_Save() saves/stores a Catalog record to the database
func catalog_Save(r dm.Catalog,usr string) error {

    var err error

	logs.Storing("Catalog",fmt.Sprintf("%s", r))

	if len(r.ID) == 0 {
		r.ID = Catalog_NewID(r)
	}


// Please Create Functions Below in the adaptor/Catalog_impl.go file
	err1 := adaptor.Catalog_Delete_Impl(r.ID)
	err2 := adaptor.Catalog_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


func Catalog_NewID(r dm.Catalog) string {
	
		// catalog_NewIDImpl should be specified in dao/Catalog_Impl.go
		// to provide the implementation for the special case.
		// override should return id - override function should be defined as
		// catalog_NewID_Impl(r dm.Catalog) string {...}
		//
		id := adaptor.Catalog_NewID_Impl(r)
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

