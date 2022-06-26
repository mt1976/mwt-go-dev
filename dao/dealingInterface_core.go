package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/dealinginterface.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : DealingInterface (dealinginterface)
// Endpoint 	        : DealingInterface (Name)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:27
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

// DealingInterface_GetList() returns a list of all DealingInterface records
func DealingInterface_GetList() (int, []dm.DealingInterface, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.DealingInterface_SQLTable)
	count, dealinginterfaceList, _, _ := dealinginterface_Fetch(tsql)
	
	return count, dealinginterfaceList, nil
}



// DealingInterface_GetByID() returns a single DealingInterface record
func DealingInterface_GetByID(id string) (int, dm.DealingInterface, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.DealingInterface_SQLTable)
	tsql = tsql + " WHERE " + dm.DealingInterface_SQLSearchID + "='" + id + "'"
	_, _, dealinginterfaceItem, _ := dealinginterface_Fetch(tsql)

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, dealinginterfaceItem, nil
}



// DealingInterface_DeleteByID() deletes a single DealingInterface record
func DealingInterface_Delete(id string) {


	adaptor.DealingInterface_Delete_impl(id)
	
}


// DealingInterface_Store() saves/stores a DealingInterface record to the database
func DealingInterface_Store(r dm.DealingInterface,req *http.Request) error {

	err, r := DealingInterface_Validate(r)
	if err == nil {
		err = dealinginterface_Save(r, Audit_User(req))
	} else {
		logs.Information("DealingInterface_Store()", err.Error())
	}

	return err
}

// DealingInterface_StoreSystem() saves/stores a DealingInterface record to the database
func DealingInterface_StoreSystem(r dm.DealingInterface) error {
	
	err, r := DealingInterface_Validate(r)
	if err == nil {
		err = dealinginterface_Save(r, Audit_Host())
	} else {
		logs.Information("DealingInterface_Store()", err.Error())
	}

	return err
}

// DealingInterface_Validate() validates for saves/stores a DealingInterface record to the database
func DealingInterface_Validate(r dm.DealingInterface) (error,dm.DealingInterface) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// dealinginterface_Save() saves/stores a DealingInterface record to the database
func dealinginterface_Save(r dm.DealingInterface,usr string) error {

    var err error



	

	if len(r.Name) == 0 {
		r.Name = DealingInterface_NewID(r)
	}

// If there are fields below, create the methods in dao\dealinginterface_impl.go


























	
logs.Storing("DealingInterface",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/DealingInterface_impl.go file
	err1 := adaptor.DealingInterface_Delete_impl(r.Name)
	err2 := adaptor.DealingInterface_Update_impl(r.Name,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// dealinginterface_Fetch read all DealingInterface's
func dealinginterface_Fetch(tsql string) (int, []dm.DealingInterface, dm.DealingInterface, error) {

	var recItem dm.DealingInterface
	var recList []dm.DealingInterface

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.Name  = get_String(rec, dm.DealingInterface_Name_sql, "")
	   recItem.AcceptReducedAmount  = get_Bool(rec, dm.DealingInterface_AcceptReducedAmount_sql, "True")
	   recItem.QuoteAsIndicative  = get_Bool(rec, dm.DealingInterface_QuoteAsIndicative_sql, "True")
	   recItem.RateTimeOut  = get_Int(rec, dm.DealingInterface_RateTimeOut_sql, "0")
	   recItem.PropagationDelay  = get_Int(rec, dm.DealingInterface_PropagationDelay_sql, "0")
	   recItem.CheckLiquidity  = get_Bool(rec, dm.DealingInterface_CheckLiquidity_sql, "True")
	   recItem.ChangeQuoteDirection  = get_Bool(rec, dm.DealingInterface_ChangeQuoteDirection_sql, "True")
	   recItem.GenerateRejectedDeals  = get_Bool(rec, dm.DealingInterface_GenerateRejectedDeals_sql, "True")
	   recItem.SpotUpdatesForForwardQuotes  = get_Bool(rec, dm.DealingInterface_SpotUpdatesForForwardQuotes_sql, "True")
	   recItem.SettlementInstructionStyle  = get_String(rec, dm.DealingInterface_SettlementInstructionStyle_sql, "")
	   recItem.CanRetractQuotes  = get_Bool(rec, dm.DealingInterface_CanRetractQuotes_sql, "True")
	   recItem.CancelESPifNotPriced  = get_Bool(rec, dm.DealingInterface_CancelESPifNotPriced_sql, "True")
	   recItem.CancelRFQSifNotPriced  = get_Bool(rec, dm.DealingInterface_CancelRFQSifNotPriced_sql, "True")
	   recItem.CancelonDealingSuspended  = get_Bool(rec, dm.DealingInterface_CancelonDealingSuspended_sql, "True")
	   recItem.CreditCheckedatDI  = get_Bool(rec, dm.DealingInterface_CreditCheckedatDI_sql, "True")
	   recItem.DuplicateCheckonExternalRef  = get_Bool(rec, dm.DealingInterface_DuplicateCheckonExternalRef_sql, "True")
	   recItem.LimitCheckQuote  = get_Bool(rec, dm.DealingInterface_LimitCheckQuote_sql, "True")
	   recItem.LimitCheckonRFQDealSubmission  = get_Bool(rec, dm.DealingInterface_LimitCheckonRFQDealSubmission_sql, "True")
	   recItem.ListenonLimits  = get_Bool(rec, dm.DealingInterface_ListenonLimits_sql, "True")
	   recItem.MarginStyle  = get_String(rec, dm.DealingInterface_MarginStyle_sql, "")
	   recItem.UseRerouteDefinitionOnly  = get_Bool(rec, dm.DealingInterface_UseRerouteDefinitionOnly_sql, "True")
	   recItem.BypassConfirmation  = get_Bool(rec, dm.DealingInterface_BypassConfirmation_sql, "True")
	   recItem.DIOnAcceptance  = get_Bool(rec, dm.DealingInterface_DIOnAcceptance_sql, "True")
	   recItem.IgnoreESPAmountRules  = get_Bool(rec, dm.DealingInterface_IgnoreESPAmountRules_sql, "True")
	
	// If there are fields below, create the methods in adaptor\DealingInterface_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func DealingInterface_NewID(r dm.DealingInterface) string {
	
			id := uuid.New().String()
	
	return id
}



// dealinginterface_Fetch read all DealingInterface's
func DealingInterface_New() (int, []dm.DealingInterface, dm.DealingInterface, error) {

	var r = dm.DealingInterface{}
	var rList []dm.DealingInterface
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}