package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/dealtype.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : DealType (dealtype)
// Endpoint 	        : DealType (DealTypeKey)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:14
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

// DealType_GetList() returns a list of all DealType records
func DealType_GetList() (int, []dm.DealType, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.DealType_SQLTable)
	count, dealtypeList, _, _ := dealtype_Fetch(tsql)
	
	return count, dealtypeList, nil
}



// DealType_GetByID() returns a single DealType record
func DealType_GetByID(id string) (int, dm.DealType, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.DealType_SQLTable)
	tsql = tsql + " WHERE " + dm.DealType_SQLSearchID + "='" + id + "'"
	_, _, dealtypeItem, _ := dealtype_Fetch(tsql)

	return 1, dealtypeItem, nil
}



// DealType_DeleteByID() deletes a single DealType record
func DealType_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.DealType_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.DealType_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// DealType_Store() saves/stores a DealType record to the database
func DealType_Store(r dm.DealType,req *http.Request) error {

	err := dealtype_Save(r,Audit_User(req))

	return err
}

// DealType_StoreSystem() saves/stores a DealType record to the database
func DealType_StoreSystem(r dm.DealType) error {
	
	err := dealtype_Save(r,Audit_Host())

	return err
}

// dealtype_Save() saves/stores a DealType record to the database
func dealtype_Save(r dm.DealType,usr string) error {

    var err error

	logs.Storing("DealType",fmt.Sprintf("%s", r))

	if len(r.DealTypeKey) == 0 {
		r.DealTypeKey = DealType_NewID(r)
	}


// Please Create Functions Below in the adaptor/DealType_impl.go file
	err1 := adaptor.DealType_Delete_Impl(r.DealTypeKey)
	err2 := adaptor.DealType_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// dealtype_Fetch read all employees
func dealtype_Fetch(tsql string) (int, []dm.DealType, dm.DealType, error) {

	var recItem dm.DealType
	var recList []dm.DealType

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.DealTypeKey  = get_String(rec, dm.DealType_DealTypeKey, "")
   recItem.DealTypeShortName  = get_String(rec, dm.DealType_DealTypeShortName, "")
   recItem.HostKey  = get_String(rec, dm.DealType_HostKey, "")
   recItem.IsActive  = get_Bool(rec, dm.DealType_IsActive, "True")
   recItem.Interbook  = get_Bool(rec, dm.DealType_Interbook, "True")
   recItem.BackOfficeLink  = get_Bool(rec, dm.DealType_BackOfficeLink, "True")
   recItem.HasTicket  = get_Bool(rec, dm.DealType_HasTicket, "True")
   recItem.CurrencyOverride  = get_Bool(rec, dm.DealType_CurrencyOverride, "True")
   recItem.CurrencyHolderCurrency  = get_String(rec, dm.DealType_CurrencyHolderCurrency, "")
   recItem.AllBooks  = get_Bool(rec, dm.DealType_AllBooks, "True")
   recItem.FundamentalDealTypeKey  = get_String(rec, dm.DealType_FundamentalDealTypeKey, "")
   recItem.RelatedDealType  = get_String(rec, dm.DealType_RelatedDealType, "")
   recItem.BookName  = get_String(rec, dm.DealType_BookName, "")
   recItem.ExportMethod  = get_String(rec, dm.DealType_ExportMethod, "")
   recItem.DefaultUserLayoffBooks  = get_Bool(rec, dm.DealType_DefaultUserLayoffBooks, "True")
   recItem.RFQ  = get_Bool(rec, dm.DealType_RFQ, "True")
   recItem.OBS  = get_Bool(rec, dm.DealType_OBS, "True")
   recItem.KID  = get_Bool(rec, dm.DealType_KID, "True")
   recItem.InternalId  = get_Int(rec, dm.DealType_InternalId, "0")
   recItem.InternalDeleted  = get_Time(rec, dm.DealType_InternalDeleted, "")
   recItem.UpdatedTransactionId  = get_String(rec, dm.DealType_UpdatedTransactionId, "")
   recItem.UpdatedUserId  = get_String(rec, dm.DealType_UpdatedUserId, "")
   recItem.UpdatedDateTime  = get_Time(rec, dm.DealType_UpdatedDateTime, "")
   recItem.DeletedTransactionId  = get_String(rec, dm.DealType_DeletedTransactionId, "")
   recItem.DeletedUserId  = get_String(rec, dm.DealType_DeletedUserId, "")
   recItem.ChangeType  = get_String(rec, dm.DealType_ChangeType, "")
// If there are fields below, create the methods in dao\DealType_Impl.go





















































	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func DealType_NewID(r dm.DealType) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

