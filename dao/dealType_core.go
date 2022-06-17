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
// Date & Time		    : 17/06/2022 at 18:38:10
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


	adaptor.DealType_Delete_impl(id)
	
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



	

	if len(r.DealTypeKey) == 0 {
		r.DealTypeKey = DealType_NewID(r)
	}

// If there are fields below, create the methods in dao\dealtype_impl.go




























	
logs.Storing("DealType",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/DealType_impl.go file
	err1 := adaptor.DealType_Delete_impl(r.DealTypeKey)
	err2 := adaptor.DealType_Update_impl(r.DealTypeKey,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// dealtype_Fetch read all DealType's
func dealtype_Fetch(tsql string) (int, []dm.DealType, dm.DealType, error) {

	var recItem dm.DealType
	var recList []dm.DealType

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.DealTypeKey  = get_String(rec, dm.DealType_DealTypeKey_sql, "")
	   recItem.DealTypeShortName  = get_String(rec, dm.DealType_DealTypeShortName_sql, "")
	   recItem.HostKey  = get_String(rec, dm.DealType_HostKey_sql, "")
	   recItem.IsActive  = get_Bool(rec, dm.DealType_IsActive_sql, "True")
	   recItem.Interbook  = get_Bool(rec, dm.DealType_Interbook_sql, "True")
	   recItem.BackOfficeLink  = get_Bool(rec, dm.DealType_BackOfficeLink_sql, "True")
	   recItem.HasTicket  = get_Bool(rec, dm.DealType_HasTicket_sql, "True")
	   recItem.CurrencyOverride  = get_Bool(rec, dm.DealType_CurrencyOverride_sql, "True")
	   recItem.CurrencyHolderCurrency  = get_String(rec, dm.DealType_CurrencyHolderCurrency_sql, "")
	   recItem.AllBooks  = get_Bool(rec, dm.DealType_AllBooks_sql, "True")
	   recItem.FundamentalDealTypeKey  = get_String(rec, dm.DealType_FundamentalDealTypeKey_sql, "")
	   recItem.RelatedDealType  = get_String(rec, dm.DealType_RelatedDealType_sql, "")
	   recItem.BookName  = get_String(rec, dm.DealType_BookName_sql, "")
	   recItem.ExportMethod  = get_String(rec, dm.DealType_ExportMethod_sql, "")
	   recItem.DefaultUserLayoffBooks  = get_Bool(rec, dm.DealType_DefaultUserLayoffBooks_sql, "True")
	   recItem.RFQ  = get_Bool(rec, dm.DealType_RFQ_sql, "True")
	   recItem.OBS  = get_Bool(rec, dm.DealType_OBS_sql, "True")
	   recItem.KID  = get_Bool(rec, dm.DealType_KID_sql, "True")
	   recItem.InternalId  = get_Int(rec, dm.DealType_InternalId_sql, "0")
	   recItem.InternalDeleted  = get_Time(rec, dm.DealType_InternalDeleted_sql, "")
	   recItem.UpdatedTransactionId  = get_String(rec, dm.DealType_UpdatedTransactionId_sql, "")
	   recItem.UpdatedUserId  = get_String(rec, dm.DealType_UpdatedUserId_sql, "")
	   recItem.UpdatedDateTime  = get_Time(rec, dm.DealType_UpdatedDateTime_sql, "")
	   recItem.DeletedTransactionId  = get_String(rec, dm.DealType_DeletedTransactionId_sql, "")
	   recItem.DeletedUserId  = get_String(rec, dm.DealType_DeletedUserId_sql, "")
	   recItem.ChangeType  = get_String(rec, dm.DealType_ChangeType_sql, "")
	
	// If there are fields below, create the methods in adaptor\DealType_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
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

