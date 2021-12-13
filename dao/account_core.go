package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/account.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Account (account)
// Endpoint 	        : Account (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 13/12/2021 at 17:02:22
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

// Account_GetList() returns a list of all Account records
func Account_GetList() (int, []dm.Account, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	count, accountList, _, _ := account_Fetch(tsql)
	
	return count, accountList, nil
}



// Account_GetByID() returns a single Account record
func Account_GetByID(id string) (int, dm.Account, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	tsql = tsql + " WHERE " + dm.Account_SQLSearchID + "='" + id + "'"
	_, _, accountItem, _ := account_Fetch(tsql)

	return 1, accountItem, nil
}

// Account_GetByReverseLookup(id string) returns a single Account record
func Account_GetByReverseLookup(id string) (int, dm.Account, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	tsql = tsql + " WHERE CashBalance = '" + id + "'"

	_, _, accountItem, _ := account_Fetch(tsql)
	
	return 1, accountItem, nil
}

// Account_DeleteByID() deletes a single Account record
func Account_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Account_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Account_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Account_Store() saves/stores a Account record to the database
func Account_Store(r dm.Account,req *http.Request) error {

	err := account_Save(r,Audit_User(req))

	return err
}

// Account_StoreSystem() saves/stores a Account record to the database
func Account_StoreSystem(r dm.Account) error {
	
	err := account_Save(r,Audit_Host())

	return err
}

// account_Save() saves/stores a Account record to the database
func account_Save(r dm.Account,usr string) error {

    var err error

	logs.Storing("Account",fmt.Sprintf("%s", r))

	if len(r.SienaReference) == 0 {
		r.SienaReference = Account_NewID(r)
	}


// Please Create Functions Below in the adaptor/Account_impl.go file
	err1 := adaptor.Account_Delete_Impl(r.SienaReference)
	err2 := adaptor.Account_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// account_Fetch read all employees
func account_Fetch(tsql string) (int, []dm.Account, dm.Account, error) {

	var recItem dm.Account
	var recList []dm.Account

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 13/12/2021 by matttownsend on silicon.local - START
   recItem.SienaReference  = get_String(rec, dm.Account_SienaReference, "")
   recItem.CustomerSienaView  = get_String(rec, dm.Account_CustomerSienaView, "")
   recItem.SienaCommonRef  = get_String(rec, dm.Account_SienaCommonRef, "")
   recItem.Status  = get_String(rec, dm.Account_Status, "")
   recItem.StartDate  = get_Time(rec, dm.Account_StartDate, "")
   recItem.MaturityDate  = get_Time(rec, dm.Account_MaturityDate, "")
   recItem.ContractNumber  = get_String(rec, dm.Account_ContractNumber, "")
   recItem.ExternalReference  = get_String(rec, dm.Account_ExternalReference, "")
   recItem.CCY  = get_String(rec, dm.Account_CCY, "")
   recItem.Book  = get_String(rec, dm.Account_Book, "")
   recItem.MandatedUser  = get_String(rec, dm.Account_MandatedUser, "")
   recItem.BackOfficeNotes  = get_String(rec, dm.Account_BackOfficeNotes, "")
   recItem.CashBalance  = get_Float(rec, dm.Account_CashBalance, "0.00")
   recItem.AccountNumber  = get_String(rec, dm.Account_AccountNumber, "")
   recItem.AccountName  = get_String(rec, dm.Account_AccountName, "")
   recItem.LedgerBalance  = get_Float(rec, dm.Account_LedgerBalance, "0.00")
   recItem.Portfolio  = get_String(rec, dm.Account_Portfolio, "")
   recItem.AgreementId  = get_Int(rec, dm.Account_AgreementId, "0")
   recItem.BackOfficeRefNo  = get_String(rec, dm.Account_BackOfficeRefNo, "")
   recItem.ISIN  = get_String(rec, dm.Account_ISIN, "")
   recItem.UTI  = get_String(rec, dm.Account_UTI, "")
   recItem.CCYName  = get_String(rec, dm.Account_CCYName, "")
   recItem.BookName  = get_String(rec, dm.Account_BookName, "")
   recItem.PortfolioName  = get_String(rec, dm.Account_PortfolioName, "")
   recItem.Centre  = get_String(rec, dm.Account_Centre, "")
   recItem.DealTypeKey  = get_String(rec, dm.Account_DealTypeKey, "")
   recItem.DealTypeShortName  = get_String(rec, dm.Account_DealTypeShortName, "")
   recItem.InternalId  = get_Int(rec, dm.Account_InternalId, "0")
   recItem.InternalDeleted  = get_Time(rec, dm.Account_InternalDeleted, "")
   recItem.UpdatedTransactionId  = get_String(rec, dm.Account_UpdatedTransactionId, "")
   recItem.UpdatedUserId  = get_String(rec, dm.Account_UpdatedUserId, "")
   recItem.UpdatedDateTime  = get_Time(rec, dm.Account_UpdatedDateTime, "")
   recItem.DeletedTransactionId  = get_String(rec, dm.Account_DeletedTransactionId, "")
   recItem.DeletedUserId  = get_String(rec, dm.Account_DeletedUserId, "")
   recItem.ChangeType  = get_String(rec, dm.Account_ChangeType, "")
   recItem.CCYDp  = get_Int(rec, dm.Account_CCYDp, "0")
   recItem.CompID  = get_String(rec, dm.Account_CompID, "")
   recItem.Firm  = get_String(rec, dm.Account_Firm, "")
   recItem.DealType  = get_String(rec, dm.Account_DealType, "")
   recItem.FullDealType  = get_String(rec, dm.Account_FullDealType, "")
   recItem.DealingInterface  = get_String(rec, dm.Account_DealingInterface, "")
   recItem.DealtAmount  = get_Float(rec, dm.Account_DealtAmount, "0.00")
   recItem.ParentContractNumber  = get_String(rec, dm.Account_ParentContractNumber, "")
   recItem.InterestFrequency  = get_String(rec, dm.Account_InterestFrequency, "")
   recItem.InterestAction  = get_String(rec, dm.Account_InterestAction, "")
   recItem.InterestStrategy  = get_String(rec, dm.Account_InterestStrategy, "")
   recItem.InterestBasis  = get_String(rec, dm.Account_InterestBasis, "")
   recItem.SienaDealer  = get_String(rec, dm.Account_SienaDealer, "")
   recItem.DealOwner  = get_String(rec, dm.Account_DealOwner, "")
   recItem.OriginUser  = get_String(rec, dm.Account_OriginUser, "")
   recItem.EditedByUser  = get_String(rec, dm.Account_EditedByUser, "")
   recItem.DealOwnerMnemonic  = get_String(rec, dm.Account_DealOwnerMnemonic, "")
   recItem.UTCOriginTime  = get_String(rec, dm.Account_UTCOriginTime, "")
   recItem.UTCUpdateTime  = get_String(rec, dm.Account_UTCUpdateTime, "")
   recItem.CustomerStatementNotes  = get_String(rec, dm.Account_CustomerStatementNotes, "")
   recItem.NotesMargin  = get_Float(rec, dm.Account_NotesMargin, "0.00")
   recItem.RequestedBy  = get_String(rec, dm.Account_RequestedBy, "")
   recItem.EditReason  = get_String(rec, dm.Account_EditReason, "")
   recItem.EditOtherReason  = get_String(rec, dm.Account_EditOtherReason, "")
   recItem.NoticeDays  = get_Int(rec, dm.Account_NoticeDays, "0")
   recItem.DebitFrequency  = get_String(rec, dm.Account_DebitFrequency, "")
   recItem.CreditFrequency  = get_String(rec, dm.Account_CreditFrequency, "")
   recItem.EURAmount  = get_Float(rec, dm.Account_EURAmount, "0.00")
   recItem.EUROtherAmount  = get_Float(rec, dm.Account_EUROtherAmount, "0.00")
   recItem.PaymentSystemSienaView  = get_String(rec, dm.Account_PaymentSystemSienaView, "")
   recItem.PaymentSystemExternalView  = get_String(rec, dm.Account_PaymentSystemExternalView, "")









// If there are fields below, create the methods in dao\Account_Impl.go







































































   recItem.DealtCA_Extra  = account_DealtCA_Extra (recItem)
   recItem.AgainstCA_Extra  = account_AgainstCA_Extra (recItem)
   recItem.LedgerCA_Extra  = account_LedgerCA_Extra (recItem)
   recItem.CashBalanceCA_Extra  = account_CashBalanceCA_Extra (recItem)












































































	// Automatically generated 13/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Account_NewID(r dm.Account) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

