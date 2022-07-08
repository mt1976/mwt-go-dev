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
// Date & Time		    : 28/06/2022 at 16:10:56
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

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, salesdeskItem, nil
}

// SalesDesk_GetByReverseLookup(id string) returns a single SalesDesk record using the Name field as key.
func SalesDesk_GetByReverseLookup(id string) (int, dm.SalesDesk, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.SalesDesk_SQLTable)
	tsql = tsql + " WHERE Name = '" + id + "'"

	_, _, salesdeskItem, _ := salesdesk_Fetch(tsql)
	
	return 1, salesdeskItem, nil
}

// SalesDesk_DeleteByID() deletes a single SalesDesk record
func SalesDesk_Delete(id string) {


	adaptor.SalesDesk_Delete_impl(id)
	
}


// SalesDesk_Store() saves/stores a SalesDesk record to the database
func SalesDesk_Store(r dm.SalesDesk,req *http.Request) error {

	err, r := SalesDesk_Validate(r)
	if err == nil {
		err = salesdesk_Save(r, Audit_User(req))
	} else {
		logs.Information("SalesDesk_Store()", err.Error())
	}

	return err
}

// SalesDesk_StoreSystem() saves/stores a SalesDesk record to the database
func SalesDesk_StoreSystem(r dm.SalesDesk) error {
	
	err, r := SalesDesk_Validate(r)
	if err == nil {
		err = salesdesk_Save(r, Audit_Host())
	} else {
		logs.Information("SalesDesk_Store()", err.Error())
	}

	return err
}

// SalesDesk_Validate() validates for saves/stores a SalesDesk record to the database
func SalesDesk_Validate(r dm.SalesDesk) (error,dm.SalesDesk) {
	var err error
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return err,r
}
//

// salesdesk_Save() saves/stores a SalesDesk record to the database
func salesdesk_Save(r dm.SalesDesk,usr string) error {

    var err error



	

	if len(r.Name) == 0 {
		r.Name = SalesDesk_NewID(r)
	}

// If there are fields below, create the methods in dao\salesdesk_impl.go








	
logs.Storing("SalesDesk",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/SalesDesk_impl.go file
	err1 := adaptor.SalesDesk_Delete_impl(r.Name)
	err2 := adaptor.SalesDesk_Update_impl(r.Name,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// salesdesk_Fetch read all SalesDesk's
func salesdesk_Fetch(tsql string) (int, []dm.SalesDesk, dm.SalesDesk, error) {

	var recItem dm.SalesDesk
	var recList []dm.SalesDesk

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.Name  = get_String(rec, dm.SalesDesk_Name_sql, "")
	   recItem.ReportDealsOver  = get_String(rec, dm.SalesDesk_ReportDealsOver_sql, "")
	   recItem.ReportDealsOverCCY  = get_String(rec, dm.SalesDesk_ReportDealsOverCCY_sql, "")
	   recItem.AccountTransferCutOffTime  = get_Time(rec, dm.SalesDesk_AccountTransferCutOffTime_sql, "")
	   recItem.AccountTransferCutOffTimeTimeZone  = get_String(rec, dm.SalesDesk_AccountTransferCutOffTimeTimeZone_sql, "")
	   recItem.AccountTransferCutOffTimeCutOffPeriod  = get_String(rec, dm.SalesDesk_AccountTransferCutOffTimeCutOffPeriod_sql, "")
	
	// If there are fields below, create the methods in adaptor\SalesDesk_impl.go
	
	
	
	
	
	
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func SalesDesk_NewID(r dm.SalesDesk) string {
	
			id := uuid.New().String()
	
	return id
}



// salesdesk_Fetch read all SalesDesk's
func SalesDesk_New() (int, []dm.SalesDesk, dm.SalesDesk, error) {

	var r = dm.SalesDesk{}
	var rList []dm.SalesDesk
	

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}