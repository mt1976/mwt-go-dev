package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/dealconversation.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : DealConversation (dealconversation)
// Endpoint 	        : DealConversation (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:26
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

// DealConversation_GetList() returns a list of all DealConversation records
func DealConversation_GetList() (int, []dm.DealConversation, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.DealConversation_SQLTable)
	count, dealconversationList, _, _ := dealconversation_Fetch(tsql)
	
	return count, dealconversationList, nil
}



// DealConversation_GetByID() returns a single DealConversation record
func DealConversation_GetByID(id string) (int, dm.DealConversation, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.DealConversation_SQLTable)
	tsql = tsql + " WHERE " + dm.DealConversation_SQLSearchID + "='" + id + "'"
	_, _, dealconversationItem, _ := dealconversation_Fetch(tsql)

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, dealconversationItem, nil
}



// DealConversation_DeleteByID() deletes a single DealConversation record
func DealConversation_Delete(id string) {


	adaptor.DealConversation_Delete_impl(id)
	
}


// DealConversation_Store() saves/stores a DealConversation record to the database
func DealConversation_Store(r dm.DealConversation,req *http.Request) error {

	err, r := DealConversation_Validate(r)
	if err == nil {
		err = dealconversation_Save(r, Audit_User(req))
	} else {
		logs.Information("DealConversation_Store()", err.Error())
	}

	return err
}

// DealConversation_StoreSystem() saves/stores a DealConversation record to the database
func DealConversation_StoreSystem(r dm.DealConversation) error {
	
	err, r := DealConversation_Validate(r)
	if err == nil {
		err = dealconversation_Save(r, Audit_Host())
	} else {
		logs.Information("DealConversation_Store()", err.Error())
	}

	return err
}

// DealConversation_Validate() validates for saves/stores a DealConversation record to the database
func DealConversation_Validate(r dm.DealConversation) (error,dm.DealConversation) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// dealconversation_Save() saves/stores a DealConversation record to the database
func dealconversation_Save(r dm.DealConversation,usr string) error {

    var err error



	

	if len(r.MessageLogReference) == 0 {
		r.MessageLogReference = DealConversation_NewID(r)
	}

// If there are fields below, create the methods in dao\dealconversation_impl.go














	
logs.Storing("DealConversation",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/DealConversation_impl.go file
	err1 := adaptor.DealConversation_Delete_impl(r.MessageLogReference)
	err2 := adaptor.DealConversation_Update_impl(r.MessageLogReference,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// dealconversation_Fetch read all DealConversation's
func dealconversation_Fetch(tsql string) (int, []dm.DealConversation, dm.DealConversation, error) {

	var recItem dm.DealConversation
	var recList []dm.DealConversation

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SienaReference  = get_String(rec, dm.DealConversation_SienaReference_sql, "")
	   recItem.Status  = get_String(rec, dm.DealConversation_Status_sql, "")
	   recItem.MessageType  = get_String(rec, dm.DealConversation_MessageType_sql, "")
	   recItem.ContractNumber  = get_String(rec, dm.DealConversation_ContractNumber_sql, "")
	   recItem.AckReference  = get_String(rec, dm.DealConversation_AckReference_sql, "")
	   recItem.NewTX  = get_Bool(rec, dm.DealConversation_NewTX_sql, "True")
	   recItem.LegNo  = get_Int(rec, dm.DealConversation_LegNo_sql, "0")
	   recItem.Summary  = get_String(rec, dm.DealConversation_Summary_sql, "")
	   recItem.BusinessDate  = get_Time(rec, dm.DealConversation_BusinessDate_sql, "")
	   recItem.TXNo  = get_Int(rec, dm.DealConversation_TXNo_sql, "0")
	   recItem.ExternalSystem  = get_String(rec, dm.DealConversation_ExternalSystem_sql, "")
	   recItem.MessageLogReference  = get_String(rec, dm.DealConversation_MessageLogReference_sql, "")
	
	// If there are fields below, create the methods in adaptor\DealConversation_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func DealConversation_NewID(r dm.DealConversation) string {
	
			id := uuid.New().String()
	
	return id
}



// dealconversation_Fetch read all DealConversation's
func DealConversation_New() (int, []dm.DealConversation, dm.DealConversation, error) {

	var r = dm.DealConversation{}
	var rList []dm.DealConversation
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}