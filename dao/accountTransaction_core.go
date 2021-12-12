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
// Date & Time		    : 12/12/2021 at 16:13:06
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

	return 1, accounttransactionItem, nil
}



// AccountTransaction_DeleteByID() deletes a single AccountTransaction record
func AccountTransaction_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.AccountTransaction_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.AccountTransaction_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// AccountTransaction_Store() saves/stores a AccountTransaction record to the database
func AccountTransaction_Store(r dm.AccountTransaction,req *http.Request) error {

	err := accounttransaction_Save(r,Audit_User(req))

	return err
}

// AccountTransaction_StoreSystem() saves/stores a AccountTransaction record to the database
func AccountTransaction_StoreSystem(r dm.AccountTransaction) error {
	
	err := accounttransaction_Save(r,Audit_Host())

	return err
}

// accounttransaction_Save() saves/stores a AccountTransaction record to the database
func accounttransaction_Save(r dm.AccountTransaction,usr string) error {

    var err error

	logs.Storing("AccountTransaction",fmt.Sprintf("%s", r))

	if len(r.SienaReference) == 0 {
		r.SienaReference = AccountTransaction_NewID(r)
	}


// Please Create Functions Below in the adaptor/AccountTransaction_impl.go file
	err1 := adaptor.AccountTransaction_Delete_Impl(r.SienaReference)
	err2 := adaptor.AccountTransaction_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// accounttransaction_Fetch read all employees
func accounttransaction_Fetch(tsql string) (int, []dm.AccountTransaction, dm.AccountTransaction, error) {

	var recItem dm.AccountTransaction
	var recList []dm.AccountTransaction

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.SienaReference  = get_String(rec, dm.AccountTransaction_SienaReference, "")
   recItem.LegNo  = get_Int(rec, dm.AccountTransaction_LegNo, "0")
   recItem.MMLegNo  = get_Int(rec, dm.AccountTransaction_MMLegNo, "0")
   recItem.Narrative  = get_String(rec, dm.AccountTransaction_Narrative, "")
   recItem.Amount  = get_Float(rec, dm.AccountTransaction_Amount, "0.00")
   recItem.StartInterestDate  = get_Time(rec, dm.AccountTransaction_StartInterestDate, "")
   recItem.EndInterestDate  = get_Time(rec, dm.AccountTransaction_EndInterestDate, "")
   recItem.Amortisation  = get_Float(rec, dm.AccountTransaction_Amortisation, "0.00")
   recItem.InterestAmount  = get_Float(rec, dm.AccountTransaction_InterestAmount, "0.00")
   recItem.InterestAction  = get_String(rec, dm.AccountTransaction_InterestAction, "")
   recItem.FixingDate  = get_Time(rec, dm.AccountTransaction_FixingDate, "")
   recItem.InterestCalculationDate  = get_Time(rec, dm.AccountTransaction_InterestCalculationDate, "")
   recItem.AmendmentAmount  = get_Float(rec, dm.AccountTransaction_AmendmentAmount, "0.00")
   recItem.DealtCcy  = get_String(rec, dm.AccountTransaction_DealtCcy, "")
   recItem.AmountDp  = get_Int(rec, dm.AccountTransaction_AmountDp, "0")
// If there are fields below, create the methods in dao\AccountTransaction_Impl.go































	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func AccountTransaction_NewID(r dm.AccountTransaction) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

