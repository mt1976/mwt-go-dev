package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/firm.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Firm (firm)
// Endpoint 	        : Firm (FirmName)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:15
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

// Firm_GetList() returns a list of all Firm records
func Firm_GetList() (int, []dm.Firm, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Firm_SQLTable)
	count, firmList, _, _ := firm_Fetch(tsql)
	
	return count, firmList, nil
}



// Firm_GetByID() returns a single Firm record
func Firm_GetByID(id string) (int, dm.Firm, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Firm_SQLTable)
	tsql = tsql + " WHERE " + dm.Firm_SQLSearchID + "='" + id + "'"
	_, _, firmItem, _ := firm_Fetch(tsql)

	return 1, firmItem, nil
}

// Firm_GetByReverseLookup(id string) returns a single Firm record
func Firm_GetByReverseLookup(id string) (int, dm.Firm, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Firm_SQLTable)
	tsql = tsql + " WHERE FullName = '" + id + "'"

	_, _, firmItem, _ := firm_Fetch(tsql)
	
	return 1, firmItem, nil
}

// Firm_DeleteByID() deletes a single Firm record
func Firm_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Firm_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Firm_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Firm_Store() saves/stores a Firm record to the database
func Firm_Store(r dm.Firm,req *http.Request) error {

	err := firm_Save(r,Audit_User(req))

	return err
}

// Firm_StoreSystem() saves/stores a Firm record to the database
func Firm_StoreSystem(r dm.Firm) error {
	
	err := firm_Save(r,Audit_Host())

	return err
}

// firm_Save() saves/stores a Firm record to the database
func firm_Save(r dm.Firm,usr string) error {

    var err error

	logs.Storing("Firm",fmt.Sprintf("%s", r))

	if len(r.FirmName) == 0 {
		r.FirmName = Firm_NewID(r)
	}


// Please Create Functions Below in the adaptor/Firm_impl.go file
	err1 := adaptor.Firm_Delete_Impl(r.FirmName)
	err2 := adaptor.Firm_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// firm_Fetch read all employees
func firm_Fetch(tsql string) (int, []dm.Firm, dm.Firm, error) {

	var recItem dm.Firm
	var recList []dm.Firm

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.FirmName  = get_String(rec, dm.Firm_FirmName, "")
   recItem.FullName  = get_String(rec, dm.Firm_FullName, "")
   recItem.Country  = get_String(rec, dm.Firm_Country, "")
   recItem.Sector  = get_String(rec, dm.Firm_Sector, "")


// If there are fields below, create the methods in dao\Firm_Impl.go













	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Firm_NewID(r dm.Firm) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

