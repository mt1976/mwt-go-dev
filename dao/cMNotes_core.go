package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/cmnotes.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CMNotes (cmnotes)
// Endpoint 	        : CMNotes (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:08
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

// CMNotes_GetList() returns a list of all CMNotes records
func CMNotes_GetList() (int, []dm.CMNotes, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CMNotes_SQLTable)
	count, cmnotesList, _, _ := cmnotes_Fetch(tsql)
	
	return count, cmnotesList, nil
}



// CMNotes_GetByID() returns a single CMNotes record
func CMNotes_GetByID(id string) (int, dm.CMNotes, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CMNotes_SQLTable)
	tsql = tsql + " WHERE " + dm.CMNotes_SQLSearchID + "='" + id + "'"
	_, _, cmnotesItem, _ := cmnotes_Fetch(tsql)

	return 1, cmnotesItem, nil
}



// CMNotes_DeleteByID() deletes a single CMNotes record
func CMNotes_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.CMNotes_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.CMNotes_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// CMNotes_Store() saves/stores a CMNotes record to the database
func CMNotes_Store(r dm.CMNotes,req *http.Request) error {

	err := cmnotes_Save(r,Audit_User(req))

	return err
}

// CMNotes_StoreSystem() saves/stores a CMNotes record to the database
func CMNotes_StoreSystem(r dm.CMNotes) error {
	
	err := cmnotes_Save(r,Audit_Host())

	return err
}

// cmnotes_Save() saves/stores a CMNotes record to the database
func cmnotes_Save(r dm.CMNotes,usr string) error {

    var err error

	logs.Storing("CMNotes",fmt.Sprintf("%s", r))

	if len(r.NoteId) == 0 {
		r.NoteId = CMNotes_NewID(r)
	}


// Please Create Functions Below in the adaptor/CMNotes_impl.go file
	err1 := adaptor.CMNotes_Delete_Impl(r.NoteId)
	err2 := adaptor.CMNotes_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// cmnotes_Fetch read all employees
func cmnotes_Fetch(tsql string) (int, []dm.CMNotes, dm.CMNotes, error) {

	var recItem dm.CMNotes
	var recList []dm.CMNotes

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.NoteId  = get_Int(rec, dm.CMNotes_NoteId, "0")
   recItem.StreamId  = get_Int(rec, dm.CMNotes_StreamId, "0")
   recItem.Summary  = get_String(rec, dm.CMNotes_Summary, "")
   recItem.Details  = get_String(rec, dm.CMNotes_Details, "")
   recItem.RecordState  = get_Int(rec, dm.CMNotes_RecordState, "0")
   recItem.CreatedBy  = get_String(rec, dm.CMNotes_CreatedBy, "")
   recItem.CreatedDateTime  = get_Time(rec, dm.CMNotes_CreatedDateTime, "")
// If there are fields below, create the methods in dao\CMNotes_Impl.go















	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func CMNotes_NewID(r dm.CMNotes) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

