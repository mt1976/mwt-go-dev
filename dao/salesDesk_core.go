package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/salesdesk.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : SalesDesk (salesdesk)
// Endpoint 	        : SalesDesk (Desk)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:18
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

// SalesDesk_GetList() returns a list of all SalesDesk records
func SalesDesk_GetList() (int, []dm.SalesDesk, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.SalesDesk_SQLTable)
	count, salesdeskList, _, _ := salesdesk_Fetch(tsql)
	
	return count, salesdeskList, nil
}


// SalesDesk_GetLookup() returns a lookup list of all SalesDesk items in lookup format
func SalesDesk_GetLookup() []dm.Lookup_Item {

	var returnList []dm.Lookup_Item

	
	    count, salesdeskList, _ := SalesDesk_GetList()
	
	
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: salesdeskList[i].Name, Name: salesdeskList[i].Name})
	}
	return returnList
}


// SalesDesk_GetByID() returns a single SalesDesk record
func SalesDesk_GetByID(id string) (int, dm.SalesDesk, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.SalesDesk_SQLTable)
	tsql = tsql + " WHERE " + dm.SalesDesk_SQLSearchID + "='" + id + "'"
	_, _, salesdeskItem, _ := salesdesk_Fetch(tsql)

	return 1, salesdeskItem, nil
}

// SalesDesk_GetByReverseLookup(id string) returns a single SalesDesk record
func SalesDesk_GetByReverseLookup(id string) (int, dm.SalesDesk, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.SalesDesk_SQLTable)
	tsql = tsql + " WHERE Name = '" + id + "'"

	_, _, salesdeskItem, _ := salesdesk_Fetch(tsql)
	
	return 1, salesdeskItem, nil
}

// SalesDesk_DeleteByID() deletes a single SalesDesk record
func SalesDesk_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.SalesDesk_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.SalesDesk_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// SalesDesk_Store() saves/stores a SalesDesk record to the database
func SalesDesk_Store(r dm.SalesDesk,req *http.Request) error {

	err := salesdesk_Save(r,Audit_User(req))

	return err
}

// SalesDesk_StoreSystem() saves/stores a SalesDesk record to the database
func SalesDesk_StoreSystem(r dm.SalesDesk) error {
	
	err := salesdesk_Save(r,Audit_Host())

	return err
}

// salesdesk_Save() saves/stores a SalesDesk record to the database
func salesdesk_Save(r dm.SalesDesk,usr string) error {

    var err error

	logs.Storing("SalesDesk",fmt.Sprintf("%s", r))

	if len(r.Name) == 0 {
		r.Name = SalesDesk_NewID(r)
	}


// Please Create Functions Below in the adaptor/SalesDesk_impl.go file
	err1 := adaptor.SalesDesk_Delete_Impl(r.Name)
	err2 := adaptor.SalesDesk_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// salesdesk_Fetch read all employees
func salesdesk_Fetch(tsql string) (int, []dm.SalesDesk, dm.SalesDesk, error) {

	var recItem dm.SalesDesk
	var recList []dm.SalesDesk

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.Name  = get_String(rec, dm.SalesDesk_Name, "")
   recItem.ReportDealsOver  = get_String(rec, dm.SalesDesk_ReportDealsOver, "")
   recItem.ReportDealsOverCCY  = get_String(rec, dm.SalesDesk_ReportDealsOverCCY, "")
   recItem.AccountTransferCutOffTime  = get_Time(rec, dm.SalesDesk_AccountTransferCutOffTime, "")
   recItem.AccountTransferCutOffTimeTimeZone  = get_String(rec, dm.SalesDesk_AccountTransferCutOffTimeTimeZone, "")
   recItem.AccountTransferCutOffTimeCutOffPeriod  = get_String(rec, dm.SalesDesk_AccountTransferCutOffTimeCutOffPeriod, "")
// If there are fields below, create the methods in dao\SalesDesk_Impl.go













	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func SalesDesk_NewID(r dm.SalesDesk) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

