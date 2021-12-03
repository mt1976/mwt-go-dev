package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/template.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Template (template)
// Endpoint 	        : Template (TemplateID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 03/12/2021 at 13:17:01
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

	logs.Storing("Template",fmt.Sprintf("%s", r))

	if len(r.ID) == 0 {
		r.ID = Template_NewID(r)
	}




//Deal with the if its Application or null add this bit, otherwise dont.
		//fmt.Println(credentialStore)

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, Audit_User())
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",Audit_User())
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

	ts = addData(ts, dm.Template_SYSId, r.SYSId)
	ts = addData(ts, dm.Template_FIELD1, r.FIELD1)
	ts = addData(ts, dm.Template_FIELD2, r.FIELD2)
	ts = addData(ts, dm.Template_SYSCreated, r.SYSCreated)
	ts = addData(ts, dm.Template_SYSCreatedBy, r.SYSCreatedBy)
	ts = addData(ts, dm.Template_SYSCreatedHost, r.SYSCreatedHost)
	ts = addData(ts, dm.Template_SYSUpdated, r.SYSUpdated)
	ts = addData(ts, dm.Template_SYSUpdatedHost, r.SYSUpdatedHost)
	ts = addData(ts, dm.Template_SYSUpdatedBy, r.SYSUpdatedBy)
	ts = addData(ts, dm.Template_ID, r.ID)
	

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
	// Automatically generated 03/12/2021 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.Template_SYSId, "0")
   recItem.FIELD1  = get_String(rec, dm.Template_FIELD1, "")
   recItem.FIELD2  = get_String(rec, dm.Template_FIELD2, "")
   recItem.SYSCreated  = get_String(rec, dm.Template_SYSCreated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.Template_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.Template_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.Template_SYSUpdated, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.Template_SYSUpdatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.Template_SYSUpdatedBy, "")
   recItem.ID  = get_String(rec, dm.Template_ID, "")
// Automatically generated 03/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Template_NewID(r dm.Template) string {
	
	
			id := uuid.New().String()

	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

