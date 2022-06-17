package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/portfolio.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Portfolio (portfolio)
// Endpoint 	        : Portfolio (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:13
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

// Portfolio_GetList() returns a list of all Portfolio records
func Portfolio_GetList() (int, []dm.Portfolio, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Portfolio_SQLTable)
	count, portfolioList, _, _ := portfolio_Fetch(tsql)
	
	return count, portfolioList, nil
}


// Portfolio_GetLookup() returns a lookup list of all Portfolio items in lookup format
func Portfolio_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, portfolioList, _ := Portfolio_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: portfolioList[i].Code, Name: portfolioList[i].Description1})
	}
	return returnList
}


// Portfolio_GetByID() returns a single Portfolio record
func Portfolio_GetByID(id string) (int, dm.Portfolio, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Portfolio_SQLTable)
	tsql = tsql + " WHERE " + dm.Portfolio_SQLSearchID + "='" + id + "'"
	_, _, portfolioItem, _ := portfolio_Fetch(tsql)

	return 1, portfolioItem, nil
}

// Portfolio_GetByReverseLookup(id string) returns a single Portfolio record using the Description1 field as key.
func Portfolio_GetByReverseLookup(id string) (int, dm.Portfolio, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Portfolio_SQLTable)
	tsql = tsql + " WHERE Description1 = '" + id + "'"

	_, _, portfolioItem, _ := portfolio_Fetch(tsql)
	
	return 1, portfolioItem, nil
}

// Portfolio_DeleteByID() deletes a single Portfolio record
func Portfolio_Delete(id string) {


	adaptor.Portfolio_Delete_impl(id)
	
}


// Portfolio_Store() saves/stores a Portfolio record to the database
func Portfolio_Store(r dm.Portfolio,req *http.Request) error {

	err := portfolio_Save(r,Audit_User(req))

	return err
}

// Portfolio_StoreSystem() saves/stores a Portfolio record to the database
func Portfolio_StoreSystem(r dm.Portfolio) error {
	
	err := portfolio_Save(r,Audit_Host())

	return err
}

// portfolio_Save() saves/stores a Portfolio record to the database
func portfolio_Save(r dm.Portfolio,usr string) error {

    var err error



	

	if len(r.Code) == 0 {
		r.Code = Portfolio_NewID(r)
	}

// If there are fields below, create the methods in dao\portfolio_impl.go














	
logs.Storing("Portfolio",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Portfolio_impl.go file
	err1 := adaptor.Portfolio_Delete_impl(r.Code)
	err2 := adaptor.Portfolio_Update_impl(r.Code,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// portfolio_Fetch read all Portfolio's
func portfolio_Fetch(tsql string) (int, []dm.Portfolio, dm.Portfolio, error) {

	var recItem dm.Portfolio
	var recList []dm.Portfolio

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.Code  = get_String(rec, dm.Portfolio_Code_sql, "")
	   recItem.Description1  = get_String(rec, dm.Portfolio_Description1_sql, "")
	   recItem.Description2  = get_String(rec, dm.Portfolio_Description2_sql, "")
	   recItem.IsDefault  = get_Bool(rec, dm.Portfolio_IsDefault_sql, "True")
	   recItem.InternalId  = get_Int(rec, dm.Portfolio_InternalId_sql, "0")
	   recItem.InternalDeleted  = get_Time(rec, dm.Portfolio_InternalDeleted_sql, "")
	   recItem.UpdatedTransactionId  = get_String(rec, dm.Portfolio_UpdatedTransactionId_sql, "")
	   recItem.UpdatedUserId  = get_String(rec, dm.Portfolio_UpdatedUserId_sql, "")
	   recItem.UpdatedDateTime  = get_Time(rec, dm.Portfolio_UpdatedDateTime_sql, "")
	   recItem.DeletedTransactionId  = get_String(rec, dm.Portfolio_DeletedTransactionId_sql, "")
	   recItem.DeletedUserId  = get_String(rec, dm.Portfolio_DeletedUserId_sql, "")
	   recItem.ChangeType  = get_String(rec, dm.Portfolio_ChangeType_sql, "")
	
	// If there are fields below, create the methods in adaptor\Portfolio_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func Portfolio_NewID(r dm.Portfolio) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

