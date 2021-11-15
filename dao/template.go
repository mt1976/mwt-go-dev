package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/template.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			  : Template
// Endpoint Root 	  : Template
// Search QueryString : TemplateID
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 15/11/2021 at 19:03:10
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------
import (
	"fmt"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	das  "github.com/mt1976/mwt-go-dev/das"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
)

// Template_GetList() returns a list of all Template records
func Template_GetList() (int, []dm.Template, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Template_SQLTable)
	count, templateList, _, _ := template_Fetch(tsql)
	return count, templateList, nil
}

// Template_GetByID() returns a single Template record
func Template_GetByID(id string) (int, dm.Template, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Template_SQLTable)
	tsql = tsql + " WHERE " + dm.Template_SQLSearchID + "='" + id + "'"

	_, _, templateItem, _ := template_Fetch(tsql)
	return 1, templateItem, nil
}

// Template_DeleteByID() deletes a single Template record
func Template_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Template_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Template_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}

// Template_Store() saves/stores a Template record to the database
func Template_Store(r dm.Template) error {

	// TODO Implement Store Function for Template
	fmt.Println(r)


//TODO Deal with the if its Application or null add this bit, otherwise dont.
		//fmt.Println(credentialStore)
	createDate := time.Now().Format(core.DATETIMEFORMATUSER)
	if len(r.SYSCreated) == 0 {
		r.SYSCreated = createDate
	}

	currentUserID, _ := user.Current()
	userID := currentUserID.Name
	host, _ := os.Hostname()

	if len(r.AppInternalID) == 0 {
		r.AppInternalID = newtemplateStoreID()
		r.SYSCreated = createDate
		r.SYSWho = userID
		r.SYSHost = host
	}

	r.SYSUpdated = createDate
//TODO Deal with the if its Application or null add this bit, otherwise dont.


	ts := SQLData{}

	ts = addData(ts, dm.Template_SYSId, r.SYSId)
	ts = addData(ts, dm.Template_FIELD1, r.FIELD1)
	ts = addData(ts, dm.Template_FIELD2, r.FIELD2)
	ts = addData(ts, dm.Template_SYSCreated, r.SYSCreated)
	ts = addData(ts, dm.Template_SYSWho, r.SYSWho)
	ts = addData(ts, dm.Template_SYSHost, r.SYSHost)
	ts = addData(ts, dm.Template_SYSCreatedBy, r.SYSCreatedBy)
	ts = addData(ts, dm.Template_SYSCreatedHost, r.SYSCreatedHost)
	ts = addData(ts, dm.Template_SYSUpdated, r.SYSUpdated)
	ts = addData(ts, dm.Template_SYSUpdatedHost, r.SYSUpdatedHost)
	ts = addData(ts, dm.Template_SYSUpdatedBy, r.SYSUpdatedBy)
	ts = addData(ts, dm.Template_ID, r.ID)
	ts = addData(ts, dm.Template_SYSVersion, r.SYSVersion)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Template_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Template_Delete(r.ID)

	das.Execute(tsql)

	return nil
}

// template_Fetch read all employees
func template_Fetch(tsql string) (int, []dm.Template, dm.Template, error) {

	var recItem dm.Template
	var recList []dm.Template

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
    recItem.AppInternalID = get_String(rec, dm.Template_ID,"")
   recItem.SYSId  = get_Int(rec, dm.Template_SYSId, "0")
   recItem.FIELD1  = get_String(rec, dm.Template_FIELD1, "")
   recItem.FIELD2  = get_String(rec, dm.Template_FIELD2, "")
   recItem.SYSCreated  = get_String(rec, dm.Template_SYSCreated, "")
   recItem.SYSWho  = get_String(rec, dm.Template_SYSWho, "")
   recItem.SYSHost  = get_String(rec, dm.Template_SYSHost, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.Template_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.Template_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.Template_SYSUpdated, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.Template_SYSUpdatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.Template_SYSUpdatedBy, "")
   recItem.ID  = get_String(rec, dm.Template_ID, "")
   recItem.SYSVersion  = get_String(rec, dm.Template_SYSVersion, "")
// Automatically generated 15/11/2021 by matttownsend on silicon.local - END
		//Post Import Actions

		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func newtemplateStoreID() string {
	id := uuid.New().String()
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

