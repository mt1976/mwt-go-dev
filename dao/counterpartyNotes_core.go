package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/counterpartynotes.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CounterpartyNotes (counterpartynotes)
// Endpoint 	        : CounterpartyNotes (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:48
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

// CounterpartyNotes_GetList() returns a list of all CounterpartyNotes records
func CounterpartyNotes_GetList() (int, []dm.CounterpartyNotes, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyNotes_SQLTable)
	count, counterpartynotesList, _, _ := counterpartynotes_Fetch(tsql)
	
	return count, counterpartynotesList, nil
}


// CounterpartyNotes_GetLookup() returns a lookup list of all CounterpartyNotes items in lookup format
func CounterpartyNotes_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, counterpartynotesList, _ := CounterpartyNotes_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: counterpartynotesList[i].NoteId, Name: counterpartynotesList[i].Summary})
	}
	return returnList
}


// CounterpartyNotes_GetByID() returns a single CounterpartyNotes record
func CounterpartyNotes_GetByID(id string) (int, dm.CounterpartyNotes, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyNotes_SQLTable)
	tsql = tsql + " WHERE " + dm.CounterpartyNotes_SQLSearchID + "='" + id + "'"
	_, _, counterpartynotesItem, _ := counterpartynotes_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, counterpartynotesItem, nil
}



// CounterpartyNotes_DeleteByID() deletes a single CounterpartyNotes record
func CounterpartyNotes_Delete(id string) {


	adaptor.CounterpartyNotes_Delete_impl(id)
	
}


// CounterpartyNotes_Store() saves/stores a CounterpartyNotes record to the database
func CounterpartyNotes_Store(r dm.CounterpartyNotes,req *http.Request) error {

	err, r := CounterpartyNotes_Validate(r)
	if err == nil {
		err = counterpartynotes_Save(r, Audit_User(req))
	} else {
		logs.Information("CounterpartyNotes_Store()", err.Error())
	}

	return err
}

// CounterpartyNotes_StoreSystem() saves/stores a CounterpartyNotes record to the database
func CounterpartyNotes_StoreSystem(r dm.CounterpartyNotes) error {
	
	err, r := CounterpartyNotes_Validate(r)
	if err == nil {
		err = counterpartynotes_Save(r, Audit_Host())
	} else {
		logs.Information("CounterpartyNotes_Store()", err.Error())
	}

	return err
}

// CounterpartyNotes_Validate() validates for saves/stores a CounterpartyNotes record to the database
func CounterpartyNotes_Validate(r dm.CounterpartyNotes) (error,dm.CounterpartyNotes) {
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

// counterpartynotes_Save() saves/stores a CounterpartyNotes record to the database
func counterpartynotes_Save(r dm.CounterpartyNotes,usr string) error {

    var err error



	

	if len(r.NoteId) == 0 {
		r.NoteId = CounterpartyNotes_NewID(r)
	}

// If there are fields below, create the methods in dao\counterpartynotes_impl.go









	
logs.Storing("CounterpartyNotes",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/CounterpartyNotes_impl.go file
	err1 := adaptor.CounterpartyNotes_Delete_impl(r.NoteId)
	err2 := adaptor.CounterpartyNotes_Update_impl(r.NoteId,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// counterpartynotes_Fetch read all CounterpartyNotes's
func counterpartynotes_Fetch(tsql string) (int, []dm.CounterpartyNotes, dm.CounterpartyNotes, error) {

	var recItem dm.CounterpartyNotes
	var recList []dm.CounterpartyNotes

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.NoteId  = get_Int(rec, dm.CounterpartyNotes_NoteId_sql, "0")
	   recItem.StreamId  = get_Int(rec, dm.CounterpartyNotes_StreamId_sql, "0")
	   recItem.Summary  = get_String(rec, dm.CounterpartyNotes_Summary_sql, "")
	   recItem.Details  = get_String(rec, dm.CounterpartyNotes_Details_sql, "")
	   recItem.RecordState  = get_Int(rec, dm.CounterpartyNotes_RecordState_sql, "0")
	   recItem.CreatedBy  = get_String(rec, dm.CounterpartyNotes_CreatedBy_sql, "")
	   recItem.CreatedDateTime  = get_Time(rec, dm.CounterpartyNotes_CreatedDateTime_sql, "")
	
	// If there are fields below, create the methods in adaptor\CounterpartyNotes_impl.go
	
	
	
	
	
	
	
	
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
	


func CounterpartyNotes_NewID(r dm.CounterpartyNotes) string {
	
			id := uuid.New().String()
	
	return id
}



// counterpartynotes_Fetch read all CounterpartyNotes's
func CounterpartyNotes_New() (int, []dm.CounterpartyNotes, dm.CounterpartyNotes, error) {

	var r = dm.CounterpartyNotes{}
	var rList []dm.CounterpartyNotes
	

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}