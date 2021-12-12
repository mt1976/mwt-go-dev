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

	return 1, dealinginterfaceItem, nil
}



// DealingInterface_DeleteByID() deletes a single DealingInterface record
func DealingInterface_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.DealingInterface_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.DealingInterface_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// DealingInterface_Store() saves/stores a DealingInterface record to the database
func DealingInterface_Store(r dm.DealingInterface,req *http.Request) error {

	err := dealinginterface_Save(r,Audit_User(req))

	return err
}

// DealingInterface_StoreSystem() saves/stores a DealingInterface record to the database
func DealingInterface_StoreSystem(r dm.DealingInterface) error {
	
	err := dealinginterface_Save(r,Audit_Host())

	return err
}

// dealinginterface_Save() saves/stores a DealingInterface record to the database
func dealinginterface_Save(r dm.DealingInterface,usr string) error {

    var err error

	logs.Storing("DealingInterface",fmt.Sprintf("%s", r))

	if len(r.Name) == 0 {
		r.Name = DealingInterface_NewID(r)
	}


// Please Create Functions Below in the adaptor/DealingInterface_impl.go file
	err1 := adaptor.DealingInterface_Delete_Impl(r.Name)
	err2 := adaptor.DealingInterface_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// dealinginterface_Fetch read all employees
func dealinginterface_Fetch(tsql string) (int, []dm.DealingInterface, dm.DealingInterface, error) {

	var recItem dm.DealingInterface
	var recList []dm.DealingInterface

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.Name  = get_String(rec, dm.DealingInterface_Name, "")
   recItem.AcceptReducedAmount  = get_Bool(rec, dm.DealingInterface_AcceptReducedAmount, "True")
   recItem.QuoteAsIndicative  = get_Bool(rec, dm.DealingInterface_QuoteAsIndicative, "True")
   recItem.RateTimeOut  = get_Int(rec, dm.DealingInterface_RateTimeOut, "0")
   recItem.PropagationDelay  = get_Int(rec, dm.DealingInterface_PropagationDelay, "0")
   recItem.CheckLiquidity  = get_Bool(rec, dm.DealingInterface_CheckLiquidity, "True")
   recItem.ChangeQuoteDirection  = get_Bool(rec, dm.DealingInterface_ChangeQuoteDirection, "True")
   recItem.GenerateRejectedDeals  = get_Bool(rec, dm.DealingInterface_GenerateRejectedDeals, "True")
   recItem.SpotUpdatesForForwardQuotes  = get_Bool(rec, dm.DealingInterface_SpotUpdatesForForwardQuotes, "True")
   recItem.SettlementInstructionStyle  = get_String(rec, dm.DealingInterface_SettlementInstructionStyle, "")
   recItem.CanRetractQuotes  = get_Bool(rec, dm.DealingInterface_CanRetractQuotes, "True")
   recItem.CancelESPifNotPriced  = get_Bool(rec, dm.DealingInterface_CancelESPifNotPriced, "True")
   recItem.CancelRFQSifNotPriced  = get_Bool(rec, dm.DealingInterface_CancelRFQSifNotPriced, "True")
   recItem.CancelonDealingSuspended  = get_Bool(rec, dm.DealingInterface_CancelonDealingSuspended, "True")
   recItem.CreditCheckedatDI  = get_Bool(rec, dm.DealingInterface_CreditCheckedatDI, "True")
   recItem.DuplicateCheckonExternalRef  = get_Bool(rec, dm.DealingInterface_DuplicateCheckonExternalRef, "True")
   recItem.LimitCheckQuote  = get_Bool(rec, dm.DealingInterface_LimitCheckQuote, "True")
   recItem.LimitCheckonRFQDealSubmission  = get_Bool(rec, dm.DealingInterface_LimitCheckonRFQDealSubmission, "True")
   recItem.ListenonLimits  = get_Bool(rec, dm.DealingInterface_ListenonLimits, "True")
   recItem.MarginStyle  = get_String(rec, dm.DealingInterface_MarginStyle, "")
   recItem.UseRerouteDefinitionOnly  = get_Bool(rec, dm.DealingInterface_UseRerouteDefinitionOnly, "True")
   recItem.BypassConfirmation  = get_Bool(rec, dm.DealingInterface_BypassConfirmation, "True")
   recItem.DIOnAcceptance  = get_Bool(rec, dm.DealingInterface_DIOnAcceptance, "True")
   recItem.IgnoreESPAmountRules  = get_Bool(rec, dm.DealingInterface_IgnoreESPAmountRules, "True")
// If there are fields below, create the methods in dao\DealingInterface_Impl.go

















































	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func DealingInterface_NewID(r dm.DealingInterface) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

