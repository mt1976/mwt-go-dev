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

// Portfolio_GetList() returns a list of all Portfolio records
func Portfolio_GetList() (int, []dm.Portfolio, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Portfolio_SQLTable)
	count, portfolioList, _, _ := portfolio_Fetch(tsql)
	
	return count, portfolioList, nil
}



// Portfolio_GetByID() returns a single Portfolio record
func Portfolio_GetByID(id string) (int, dm.Portfolio, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Portfolio_SQLTable)
	tsql = tsql + " WHERE " + dm.Portfolio_SQLSearchID + "='" + id + "'"
	_, _, portfolioItem, _ := portfolio_Fetch(tsql)

	return 1, portfolioItem, nil
}

// Portfolio_GetByReverseLookup(id string) returns a single Portfolio record
func Portfolio_GetByReverseLookup(id string) (int, dm.Portfolio, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Portfolio_SQLTable)
	tsql = tsql + " WHERE Description1 = '" + id + "'"

	_, _, portfolioItem, _ := portfolio_Fetch(tsql)
	
	return 1, portfolioItem, nil
}

// Portfolio_DeleteByID() deletes a single Portfolio record
func Portfolio_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Portfolio_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Portfolio_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

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

	logs.Storing("Portfolio",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = Portfolio_NewID(r)
	}


// Please Create Functions Below in the adaptor/Portfolio_impl.go file
	err1 := adaptor.Portfolio_Delete_Impl(r.Code)
	err2 := adaptor.Portfolio_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// portfolio_Fetch read all employees
func portfolio_Fetch(tsql string) (int, []dm.Portfolio, dm.Portfolio, error) {

	var recItem dm.Portfolio
	var recList []dm.Portfolio

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.Code  = get_String(rec, dm.Portfolio_Code, "")
   recItem.Description1  = get_String(rec, dm.Portfolio_Description1, "")
   recItem.Description2  = get_String(rec, dm.Portfolio_Description2, "")
   recItem.IsDefault  = get_Bool(rec, dm.Portfolio_IsDefault, "True")
   recItem.InternalId  = get_Int(rec, dm.Portfolio_InternalId, "0")
   recItem.InternalDeleted  = get_Time(rec, dm.Portfolio_InternalDeleted, "")
   recItem.UpdatedTransactionId  = get_String(rec, dm.Portfolio_UpdatedTransactionId, "")
   recItem.UpdatedUserId  = get_String(rec, dm.Portfolio_UpdatedUserId, "")
   recItem.UpdatedDateTime  = get_Time(rec, dm.Portfolio_UpdatedDateTime, "")
   recItem.DeletedTransactionId  = get_String(rec, dm.Portfolio_DeletedTransactionId, "")
   recItem.DeletedUserId  = get_String(rec, dm.Portfolio_DeletedUserId, "")
   recItem.ChangeType  = get_String(rec, dm.Portfolio_ChangeType, "")
// If there are fields below, create the methods in dao\Portfolio_Impl.go

























	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
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

