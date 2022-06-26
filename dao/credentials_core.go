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
// Date & Time		    : 26/06/2022 at 18:48:24
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (

	"fmt"
	"net/http"
core "github.com/mt1976/mwt-go-dev/core"
"github.com/google/uuid"
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


// Credentials_GetLookup() returns a lookup list of all Credentials items in lookup format
func Credentials_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, credentialsList, _ := Credentials_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: credentialsList[i].Id, Name: credentialsList[i].Username})
	}
	return returnList
}


// Credentials_GetByID() returns a single Credentials record
func Credentials_GetByID(id string) (int, dm.Credentials, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Credentials_SQLTable)
	tsql = tsql + " WHERE " + dm.Credentials_SQLSearchID + "='" + id + "'"
	_, _, credentialsItem, _ := credentials_Fetch(tsql)

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
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

	err, r := Credentials_Validate(r)
	if err == nil {
		err = credentials_Save(r, Audit_User(req))
	} else {
		logs.Information("Credentials_Store()", err.Error())
	}

	return err
}

// Credentials_StoreSystem() saves/stores a Credentials record to the database
func Credentials_StoreSystem(r dm.Credentials) error {
	
	err, r := Credentials_Validate(r)
	if err == nil {
		err = credentials_Save(r, Audit_Host())
	} else {
		logs.Information("Credentials_Store()", err.Error())
	}

	return err
}

// Credentials_Validate() validates for saves/stores a Credentials record to the database
func Credentials_Validate(r dm.Credentials) (error,dm.Credentials) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// credentials_Save() saves/stores a Credentials record to the database
func credentials_Save(r dm.Credentials,usr string) error {

    var err error



	

	if len(r.Id) == 0 {
		r.Id = Credentials_NewID(r)
	}

// If there are fields below, create the methods in dao\credentials_impl.go
























	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	
logs.Storing("Credentials",fmt.Sprintf("%s", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.Credentials_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.Credentials_Id_sql, r.Id)
	ts = addData(ts, dm.Credentials_Username_sql, r.Username)
	ts = addData(ts, dm.Credentials_Password_sql, r.Password)
	ts = addData(ts, dm.Credentials_Firstname_sql, r.Firstname)
	ts = addData(ts, dm.Credentials_Lastname_sql, r.Lastname)
	ts = addData(ts, dm.Credentials_Knownas_sql, r.Knownas)
	ts = addData(ts, dm.Credentials_Email_sql, r.Email)
	ts = addData(ts, dm.Credentials_Issued_sql, r.Issued)
	ts = addData(ts, dm.Credentials_Expiry_sql, r.Expiry)
	ts = addData(ts, dm.Credentials_RoleType_sql, r.RoleType)
	ts = addData(ts, dm.Credentials_Brand_sql, r.Brand)
	ts = addData(ts, dm.Credentials_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.Credentials_SYSWho_sql, r.SYSWho)
	ts = addData(ts, dm.Credentials_SYSHost_sql, r.SYSHost)
	ts = addData(ts, dm.Credentials_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.Credentials_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.Credentials_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.Credentials_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.Credentials_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.Credentials_State_sql, r.State)
	ts = addData(ts, dm.Credentials_Notes_sql, r.Notes)
		
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Credentials_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Credentials_Delete(r.Id)
	das.Execute(tsql)



	return err

}



// credentials_Fetch read all Credentials's
func credentials_Fetch(tsql string) (int, []dm.Credentials, dm.Credentials, error) {

	var recItem dm.Credentials
	var recList []dm.Credentials

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.Credentials_SYSId_sql, "0")
	   recItem.Id  = get_String(rec, dm.Credentials_Id_sql, "")
	   recItem.Username  = get_String(rec, dm.Credentials_Username_sql, "")
	   recItem.Password  = get_String(rec, dm.Credentials_Password_sql, "")
	   recItem.Firstname  = get_String(rec, dm.Credentials_Firstname_sql, "")
	   recItem.Lastname  = get_String(rec, dm.Credentials_Lastname_sql, "")
	   recItem.Knownas  = get_String(rec, dm.Credentials_Knownas_sql, "")
	   recItem.Email  = get_String(rec, dm.Credentials_Email_sql, "")
	   recItem.Issued  = get_String(rec, dm.Credentials_Issued_sql, "")
	   recItem.Expiry  = get_String(rec, dm.Credentials_Expiry_sql, "")
	   recItem.RoleType  = get_String(rec, dm.Credentials_RoleType_sql, "")
	   recItem.Brand  = get_String(rec, dm.Credentials_Brand_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.Credentials_SYSCreated_sql, "")
	   recItem.SYSWho  = get_String(rec, dm.Credentials_SYSWho_sql, "")
	   recItem.SYSHost  = get_String(rec, dm.Credentials_SYSHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.Credentials_SYSUpdated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.Credentials_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.Credentials_SYSCreatedHost_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.Credentials_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.Credentials_SYSUpdatedHost_sql, "")
	   recItem.State  = get_String(rec, dm.Credentials_State_sql, "")
	   recItem.Notes  = get_String(rec, dm.Credentials_Notes_sql, "")
	
	// If there are fields below, create the methods in adaptor\Credentials_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func Credentials_NewID(r dm.Credentials) string {
	
			id := uuid.New().String()
	
	return id
}



// credentials_Fetch read all Credentials's
func Credentials_New() (int, []dm.Credentials, dm.Credentials, error) {

	var r = dm.Credentials{}
	var rList []dm.Credentials
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}