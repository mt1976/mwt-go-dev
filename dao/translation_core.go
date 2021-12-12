package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/translation.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Translation (translation)
// Endpoint 	        : Translation (Message)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:20
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

// Translation_GetList() returns a list of all Translation records
func Translation_GetList() (int, []dm.Translation, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Translation_SQLTable)
	count, translationList, _, _ := translation_Fetch(tsql)
	
	return count, translationList, nil
}



// Translation_GetByID() returns a single Translation record
func Translation_GetByID(id string) (int, dm.Translation, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Translation_SQLTable)
	tsql = tsql + " WHERE " + dm.Translation_SQLSearchID + "='" + id + "'"
	_, _, translationItem, _ := translation_Fetch(tsql)

	return 1, translationItem, nil
}



// Translation_DeleteByID() deletes a single Translation record
func Translation_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Translation_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Translation_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Translation_Store() saves/stores a Translation record to the database
func Translation_Store(r dm.Translation,req *http.Request) error {

	err := translation_Save(r,Audit_User(req))

	return err
}

// Translation_StoreSystem() saves/stores a Translation record to the database
func Translation_StoreSystem(r dm.Translation) error {
	
	err := translation_Save(r,Audit_Host())

	return err
}

// translation_Save() saves/stores a Translation record to the database
func translation_Save(r dm.Translation,usr string) error {

    var err error

	logs.Storing("Translation",fmt.Sprintf("%s", r))

	if len(r.Id) == 0 {
		r.Id = Translation_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.Translation_SYSId, r.SYSId)
ts = addData(ts, dm.Translation_Id, r.Id)
ts = addData(ts, dm.Translation_Class, r.Class)
ts = addData(ts, dm.Translation_Message, r.Message)
ts = addData(ts, dm.Translation_Translation, r.Translation)
ts = addData(ts, dm.Translation_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.Translation_SYSWho, r.SYSWho)
ts = addData(ts, dm.Translation_SYSHost, r.SYSHost)
ts = addData(ts, dm.Translation_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.Translation_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.Translation_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.Translation_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.Translation_SYSUpdatedHost, r.SYSUpdatedHost)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Translation_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Translation_Delete(r.Id)
	das.Execute(tsql)



	return err

}


// translation_Fetch read all employees
func translation_Fetch(tsql string) (int, []dm.Translation, dm.Translation, error) {

	var recItem dm.Translation
	var recList []dm.Translation

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.Translation_SYSId, "0")
   recItem.Id  = get_String(rec, dm.Translation_Id, "")
   recItem.Class  = get_String(rec, dm.Translation_Class, "")
   recItem.Message  = get_String(rec, dm.Translation_Message, "")
   recItem.Translation  = get_String(rec, dm.Translation_Translation, "")
   recItem.SYSCreated  = get_String(rec, dm.Translation_SYSCreated, "")
   recItem.SYSWho  = get_String(rec, dm.Translation_SYSWho, "")
   recItem.SYSHost  = get_String(rec, dm.Translation_SYSHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.Translation_SYSUpdated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.Translation_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.Translation_SYSCreatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.Translation_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.Translation_SYSUpdatedHost, "")
// If there are fields below, create the methods in dao\Translation_Impl.go



























	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Translation_NewID(r dm.Translation) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

