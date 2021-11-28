package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/broker.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Broker (broker)
// Endpoint 	        : Broker (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 28/11/2021 at 22:54:52
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"log"
	"fmt"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	das  "github.com/mt1976/mwt-go-dev/das"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
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
	return 1, brokerItem, nil
}

// Broker_GetByReverseLookup(id string) returns a single Broker record
func Broker_GetByReverseLookup(id string) (int, dm.Broker, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Broker_SQLTable)
	tsql = tsql + " WHERE Name = '" + id + "'"

	_, _, brokerItem, _ := broker_Fetch(tsql)
	return 1, brokerItem, nil
}

// Broker_DeleteByID() deletes a single Broker record
func Broker_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Broker_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Broker_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}


// Broker_Store() saves/stores a Broker record to the database
func Broker_Store(r dm.Broker) error {

	logs.Storing("Broker",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = Broker_NewID(r)
	}



	adaptor.Broker_Delete(r.Code)
	adaptor.Broker_Update(r)



	return nil

}

// broker_Fetch read all employees
func broker_Fetch(tsql string) (int, []dm.Broker, dm.Broker, error) {

	var recItem dm.Broker
	var recList []dm.Broker

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 28/11/2021 by matttownsend on silicon.local - START
   recItem.Code  = get_String(rec, dm.Broker_Code, "")
   recItem.Name  = get_String(rec, dm.Broker_Name, "")
   recItem.FullName  = get_String(rec, dm.Broker_FullName, "")
   recItem.Contact  = get_String(rec, dm.Broker_Contact, "")
   recItem.Address  = get_String(rec, dm.Broker_Address, "")
   recItem.LEI  = get_String(rec, dm.Broker_LEI, "")
// Automatically generated 28/11/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Broker_NewID(r dm.Broker) string {
	
	
			id := uuid.New().String()

	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

