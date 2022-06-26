package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/broker.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Broker (broker)
// Endpoint 	        : Broker (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:18
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

// Broker_GetList() returns a list of all Broker records
func Broker_GetList() (int, []dm.Broker, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Broker_SQLTable)
	count, brokerList, _, _ := broker_Fetch(tsql)
	
	return count, brokerList, nil
}



// Broker_GetByID() returns a single Broker record
func Broker_GetByID(id string) (int, dm.Broker, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Broker_SQLTable)
	tsql = tsql + " WHERE " + dm.Broker_SQLSearchID + "='" + id + "'"
	_, _, brokerItem, _ := broker_Fetch(tsql)

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, brokerItem, nil
}

// Broker_GetByReverseLookup(id string) returns a single Broker record using the Name field as key.
func Broker_GetByReverseLookup(id string) (int, dm.Broker, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Broker_SQLTable)
	tsql = tsql + " WHERE Name = '" + id + "'"

	_, _, brokerItem, _ := broker_Fetch(tsql)
	
	return 1, brokerItem, nil
}

// Broker_DeleteByID() deletes a single Broker record
func Broker_Delete(id string) {


	adaptor.Broker_Delete_impl(id)
	
}


// Broker_Store() saves/stores a Broker record to the database
func Broker_Store(r dm.Broker,req *http.Request) error {

	err, r := Broker_Validate(r)
	if err == nil {
		err = broker_Save(r, Audit_User(req))
	} else {
		logs.Information("Broker_Store()", err.Error())
	}

	return err
}

// Broker_StoreSystem() saves/stores a Broker record to the database
func Broker_StoreSystem(r dm.Broker) error {
	
	err, r := Broker_Validate(r)
	if err == nil {
		err = broker_Save(r, Audit_Host())
	} else {
		logs.Information("Broker_Store()", err.Error())
	}

	return err
}

// Broker_Validate() validates for saves/stores a Broker record to the database
func Broker_Validate(r dm.Broker) (error,dm.Broker) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// broker_Save() saves/stores a Broker record to the database
func broker_Save(r dm.Broker,usr string) error {

    var err error



	

	if len(r.Code) == 0 {
		r.Code = Broker_NewID(r)
	}

// If there are fields below, create the methods in dao\broker_impl.go








	
logs.Storing("Broker",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Broker_impl.go file
	err1 := adaptor.Broker_Delete_impl(r.Code)
	err2 := adaptor.Broker_Update_impl(r.Code,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// broker_Fetch read all Broker's
func broker_Fetch(tsql string) (int, []dm.Broker, dm.Broker, error) {

	var recItem dm.Broker
	var recList []dm.Broker

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.Code  = get_String(rec, dm.Broker_Code_sql, "")
	   recItem.Name  = get_String(rec, dm.Broker_Name_sql, "")
	   recItem.FullName  = get_String(rec, dm.Broker_FullName_sql, "")
	   recItem.Contact  = get_String(rec, dm.Broker_Contact_sql, "")
	   recItem.Address  = get_String(rec, dm.Broker_Address_sql, "")
	   recItem.LEI  = get_String(rec, dm.Broker_LEI_sql, "")
	
	// If there are fields below, create the methods in adaptor\Broker_impl.go
	
	
	
	
	
	
	
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
	


func Broker_NewID(r dm.Broker) string {
	
			id := uuid.New().String()
	
	return id
}



// broker_Fetch read all Broker's
func Broker_New() (int, []dm.Broker, dm.Broker, error) {

	var r = dm.Broker{}
	var rList []dm.Broker
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}