package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/dealconversation.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : DealConversation (dealconversation)
// Endpoint 	        : DealConversation (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 22/11/2021 at 11:25:55
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
	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
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
	return 1, dealconversationItem, nil
}



// DealConversation_DeleteByID() deletes a single DealConversation record
func DealConversation_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.DealConversation_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.DealConversation_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}

// DealConversation_Store() saves/stores a DealConversation record to the database
func DealConversation_Store(r dm.DealConversation) error {

	logs.Storing("DealConversation",fmt.Sprintf("%s", r))

	if len(r.MessageLogReference) == 0 {
		r.MessageLogReference = DealConversation_NewID(r)
	}



	adaptor.DealConversation_Delete(r.MessageLogReference)
	adaptor.DealConversation_Update(r)


	return nil
}

// dealconversation_Fetch read all employees
func dealconversation_Fetch(tsql string) (int, []dm.DealConversation, dm.DealConversation, error) {

	var recItem dm.DealConversation
	var recList []dm.DealConversation

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
    recItem.AppInternalID = get_String(rec, dm.DealConversation_MessageLogReference,"")
   recItem.SienaReference  = get_String(rec, dm.DealConversation_SienaReference, "")
   recItem.Status  = get_String(rec, dm.DealConversation_Status, "")
   recItem.MessageType  = get_String(rec, dm.DealConversation_MessageType, "")
   recItem.ContractNumber  = get_String(rec, dm.DealConversation_ContractNumber, "")
   recItem.AckReference  = get_String(rec, dm.DealConversation_AckReference, "")
   recItem.NewTX  = get_Bool(rec, dm.DealConversation_NewTX, "True")
   recItem.LegNo  = get_Int(rec, dm.DealConversation_LegNo, "0")
   recItem.Summary  = get_String(rec, dm.DealConversation_Summary, "")
   recItem.BusinessDate  = get_Time(rec, dm.DealConversation_BusinessDate, "")
   recItem.TXNo  = get_Int(rec, dm.DealConversation_TXNo, "0")
   recItem.ExternalSystem  = get_String(rec, dm.DealConversation_ExternalSystem, "")
   recItem.MessageLogReference  = get_String(rec, dm.DealConversation_MessageLogReference, "")
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func DealConversation_NewID(r dm.DealConversation) string {
	
	
			id := uuid.New().String()

	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

