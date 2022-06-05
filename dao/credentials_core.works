package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/credentials.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Credentials (credentials)
// Endpoint 	        : Credentials (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:11
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"log"
	
	"fmt"
	"net/http"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	das  "github.com/mt1976/mwt-go-dev/das"
	
	
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
)

// Credentials_GetList() returns a list of all Credentials records
func Credentials_GetList() (int, []dm.Credentials, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Credentials_SQLTable)
	count, credentialsList, _, _ := credentials_Fetch(tsql)
	
	return count, credentialsList, nil
}



// Credentials_GetByID() returns a single Credentials record
func Credentials_GetByID(id string) (int, dm.Credentials, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Credentials_SQLTable)
	tsql = tsql + " WHERE " + dm.Credentials_SQLSearchID + "='" + id + "'"
	_, _, credentialsItem, _ := credentials_Fetch(tsql)

	return 1, credentialsItem, nil
}



// Credentials_DeleteByID() deletes a single Credentials record
func Credentials_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Credentials_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Credentials_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Credentials_Store() saves/stores a Credentials record to the database
func Credentials_Store(r dm.Credentials,req *http.Request) error {

	err := credentials_Save(r,Audit_User(req))

	return err
}

// Credentials_StoreSystem() saves/stores a Credentials record to the database
func Credentials_StoreSystem(r dm.Credentials) error {
	
	err := credentials_Save(r,Audit_Host())

	return err
}

// credentials_Save() saves/stores a Credentials record to the database
func credentials_Save(r dm.Credentials,usr string) error {

    var err error

	logs.Storing("Credentials",fmt.Sprintf("%s", r))

	if len(r.Id) == 0 {
		r.Id = Credentials_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.Credentials_SYSId, r.SYSId)
ts = addData(ts, dm.Credentials_Id, r.Id)
ts = addData(ts, dm.Credentials_Username, r.Username)
ts = addData(ts, dm.Credentials_Password, r.Password)
ts = addData(ts, dm.Credentials_Firstname, r.Firstname)
ts = addData(ts, dm.Credentials_Lastname, r.Lastname)
ts = addData(ts, dm.Credentials_Knownas, r.Knownas)
ts = addData(ts, dm.Credentials_Email, r.Email)
ts = addData(ts, dm.Credentials_Issued, r.Issued)
ts = addData(ts, dm.Credentials_Expiry, r.Expiry)
ts = addData(ts, dm.Credentials_RoleType, r.RoleType)
ts = addData(ts, dm.Credentials_Brand, r.Brand)
ts = addData(ts, dm.Credentials_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.Credentials_SYSWho, r.SYSWho)
ts = addData(ts, dm.Credentials_SYSHost, r.SYSHost)
ts = addData(ts, dm.Credentials_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.Credentials_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.Credentials_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.Credentials_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.Credentials_SYSUpdatedHost, r.SYSUpdatedHost)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Credentials_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Credentials_Delete(r.Id)
	das.Execute(tsql)



	return err

}


// credentials_Fetch read all employees
func credentials_Fetch(tsql string) (int, []dm.Credentials, dm.Credentials, error) {

	var recItem dm.Credentials
	var recList []dm.Credentials

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.Credentials_SYSId, "0")
   recItem.Id  = get_String(rec, dm.Credentials_Id, "")
   recItem.Username  = get_String(rec, dm.Credentials_Username, "")
   recItem.Password  = get_String(rec, dm.Credentials_Password, "")
   recItem.Firstname  = get_String(rec, dm.Credentials_Firstname, "")
   recItem.Lastname  = get_String(rec, dm.Credentials_Lastname, "")
   recItem.Knownas  = get_String(rec, dm.Credentials_Knownas, "")
   recItem.Email  = get_String(rec, dm.Credentials_Email, "")
   recItem.Issued  = get_String(rec, dm.Credentials_Issued, "")
   recItem.Expiry  = get_String(rec, dm.Credentials_Expiry, "")
   recItem.RoleType  = get_String(rec, dm.Credentials_RoleType, "")
   recItem.Brand  = get_String(rec, dm.Credentials_Brand, "")
   recItem.SYSCreated  = get_String(rec, dm.Credentials_SYSCreated, "")
   recItem.SYSWho  = get_String(rec, dm.Credentials_SYSWho, "")
   recItem.SYSHost  = get_String(rec, dm.Credentials_SYSHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.Credentials_SYSUpdated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.Credentials_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.Credentials_SYSCreatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.Credentials_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.Credentials_SYSUpdatedHost, "")
// If there are fields below, create the methods in dao\Credentials_Impl.go









































	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Credentials_NewID(r dm.Credentials) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

