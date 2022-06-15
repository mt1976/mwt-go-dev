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
// Date & Time		    : 14/06/2022 at 21:31:51
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (

	"fmt"
	"net/http"

"github.com/google/uuid"

	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
)

// Catalog_GetList() returns a list of all Catalog records
func Catalog_GetList() (int, []dm.Catalog, error) {
	
	count, catalogList, _ := adaptor.Catalog_GetList_impl()
	
	return count, catalogList, nil
}



// Catalog_GetByID() returns a single Catalog record
func Catalog_GetByID(id string) (int, dm.Catalog, error) {


	 _, catalogItem, _ := adaptor.Catalog_GetByID_impl(id)
	
	return 1, catalogItem, nil
}



// Catalog_DeleteByID() deletes a single Catalog record
func Catalog_Delete(id string) {


	adaptor.Catalog_Delete_impl(id)
	
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



	

	if len(r.ID) == 0 {
		r.ID = Catalog_NewID(r)
	}

// If there are fields below, create the methods in dao\catalog_impl.go







	
logs.Storing("Catalog",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Catalog_impl.go file
	err1 := adaptor.Catalog_Delete_impl(r.ID)
	err2 := adaptor.Catalog_Update_impl(r.ID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// catalog_Fetch is not required as GetByID, GetAll etc... have been diverted to _impl
	


func Catalog_NewID(r dm.Catalog) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

