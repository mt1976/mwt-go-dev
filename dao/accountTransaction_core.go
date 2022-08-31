package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/accounttransaction.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : AccountTransaction (accounttransaction)
// Endpoint 	        : AccountTransaction (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:42
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

// AccountTransaction_GetList() returns a list of all AccountTransaction records
func AccountTransaction_GetList() (int, []dm.AccountTransaction, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.AccountTransaction_SQLTable)
	count, accounttransactionList, _, _ := accounttransaction_Fetch(tsql)
	
	return count, accounttransactionList, nil
}



// AccountTransaction_GetByID() returns a single AccountTransaction record
func AccountTransaction_GetByID(id string) (int, dm.AccountTransaction, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.AccountTransaction_SQLTable)
	tsql = tsql + " WHERE " + dm.AccountTransaction_SQLSearchID + "='" + id + "'"
	_, _, accounttransactionItem, _ := accounttransaction_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, accounttransactionItem, nil
}



// AccountTransaction_DeleteByID() deletes a single AccountTransaction record
func AccountTransaction_Delete(id string) {


	adaptor.AccountTransaction_Delete_impl(id)
	
}


// AccountTransaction_Store() saves/stores a AccountTransaction record to the database
func AccountTransaction_Store(r dm.AccountTransaction,req *http.Request) error {

	err, r := AccountTransaction_Validate(r)
	if err == nil {
		err = accounttransaction_Save(r, Audit_User(req))
	} else {
		logs.Information("AccountTransaction_Store()", err.Error())
	}

	return err
}

// AccountTransaction_StoreSystem() saves/stores a AccountTransaction record to the database
func AccountTransaction_StoreSystem(r dm.AccountTransaction) error {
	
	err, r := AccountTransaction_Validate(r)
	if err == nil {
		err = accounttransaction_Save(r, Audit_Host())
	} else {
		logs.Information("AccountTransaction_Store()", err.Error())
	}

	return err
}

// AccountTransaction_Validate() validates for saves/stores a AccountTransaction record to the database
func AccountTransaction_Validate(r dm.AccountTransaction) (error,dm.AccountTransaction) {
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

// accounttransaction_Save() saves/stores a AccountTransaction record to the database
func accounttransaction_Save(r dm.AccountTransaction,usr string) error {

    var err error



	

	if len(r.SienaReference) == 0 {
		r.SienaReference = AccountTransaction_NewID(r)
	}

// If there are fields below, create the methods in dao\accounttransaction_impl.go

















	
logs.Storing("AccountTransaction",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/AccountTransaction_impl.go file
	err1 := adaptor.AccountTransaction_Delete_impl(r.SienaReference)
	err2 := adaptor.AccountTransaction_Update_impl(r.SienaReference,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// accounttransaction_Fetch read all AccountTransaction's
func accounttransaction_Fetch(tsql string) (int, []dm.AccountTransaction, dm.AccountTransaction, error) {

	var recItem dm.AccountTransaction
	var recList []dm.AccountTransaction

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SienaReference  = get_String(rec, dm.AccountTransaction_SienaReference_sql, "")
	   recItem.LegNo  = get_Int(rec, dm.AccountTransaction_LegNo_sql, "0")
	   recItem.MMLegNo  = get_Int(rec, dm.AccountTransaction_MMLegNo_sql, "0")
	   recItem.Narrative  = get_String(rec, dm.AccountTransaction_Narrative_sql, "")
	   recItem.Amount  = get_Float(rec, dm.AccountTransaction_Amount_sql, "0.00")
	   recItem.StartInterestDate  = get_Time(rec, dm.AccountTransaction_StartInterestDate_sql, "")
	   recItem.EndInterestDate  = get_Time(rec, dm.AccountTransaction_EndInterestDate_sql, "")
	   recItem.Amortisation  = get_Float(rec, dm.AccountTransaction_Amortisation_sql, "0.00")
	   recItem.InterestAmount  = get_Float(rec, dm.AccountTransaction_InterestAmount_sql, "0.00")
	   recItem.InterestAction  = get_String(rec, dm.AccountTransaction_InterestAction_sql, "")
	   recItem.FixingDate  = get_Time(rec, dm.AccountTransaction_FixingDate_sql, "")
	   recItem.InterestCalculationDate  = get_Time(rec, dm.AccountTransaction_InterestCalculationDate_sql, "")
	   recItem.AmendmentAmount  = get_Float(rec, dm.AccountTransaction_AmendmentAmount_sql, "0.00")
	   recItem.DealtCcy  = get_String(rec, dm.AccountTransaction_DealtCcy_sql, "")
	   recItem.AmountDp  = get_Int(rec, dm.AccountTransaction_AmountDp_sql, "0")
	
	// If there are fields below, create the methods in adaptor\AccountTransaction_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func AccountTransaction_NewID(r dm.AccountTransaction) string {
	
			id := uuid.New().String()
	
	return id
}



// accounttransaction_Fetch read all AccountTransaction's
func AccountTransaction_New() (int, []dm.AccountTransaction, dm.AccountTransaction, error) {

	var r = dm.AccountTransaction{}
	var rList []dm.AccountTransaction
	

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}