package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/customeronboarding.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CustomerOnboarding (customeronboarding)
// Endpoint 	        : CustomerOnboarding (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:31:45
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

// CustomerOnboarding_GetList() returns a list of all CustomerOnboarding records
func CustomerOnboarding_GetList() (int, []dm.CustomerOnboarding, error) {
	
	count, customeronboardingList, _ := adaptor.CustomerOnboarding_GetList_impl()
	
	return count, customeronboardingList, nil
}



// CustomerOnboarding_GetByID() returns a single CustomerOnboarding record
func CustomerOnboarding_GetByID(id string) (int, dm.CustomerOnboarding, error) {


	 _, customeronboardingItem, _ := adaptor.CustomerOnboarding_GetByID_impl(id)
	
	return 1, customeronboardingItem, nil
}



// CustomerOnboarding_DeleteByID() deletes a single CustomerOnboarding record
func CustomerOnboarding_Delete(id string) {


	adaptor.CustomerOnboarding_Delete_impl(id)
	
}


// CustomerOnboarding_Store() saves/stores a CustomerOnboarding record to the database
func CustomerOnboarding_Store(r dm.CustomerOnboarding,req *http.Request) error {

	err := customeronboarding_Save(r,Audit_User(req))

	return err
}

// CustomerOnboarding_StoreSystem() saves/stores a CustomerOnboarding record to the database
func CustomerOnboarding_StoreSystem(r dm.CustomerOnboarding) error {
	
	err := customeronboarding_Save(r,Audit_Host())

	return err
}

// customeronboarding_Save() saves/stores a CustomerOnboarding record to the database
func customeronboarding_Save(r dm.CustomerOnboarding,usr string) error {

    var err error



	

	if len(r.ID) == 0 {
		r.ID = CustomerOnboarding_NewID(r)
	}

// If there are fields below, create the methods in dao\customeronboarding_impl.go















	
logs.Storing("CustomerOnboarding",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/CustomerOnboarding_impl.go file
	err1 := adaptor.CustomerOnboarding_Delete_impl(r.ID)
	err2 := adaptor.CustomerOnboarding_Update_impl(r.ID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// customeronboarding_Fetch is not required as GetByID, GetAll etc... have been diverted to _impl
	


func CustomerOnboarding_NewID(r dm.CustomerOnboarding) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

