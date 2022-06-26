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
// Date & Time		    : 26/06/2022 at 18:48:16
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	adaptor "github.com/mt1976/mwt-go-dev/adaptor"
	core "github.com/mt1976/mwt-go-dev/core"
	das "github.com/mt1976/mwt-go-dev/das"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

// Account_GetList() returns a list of all Account records
func Account_GetList() (int, []dm.Account, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	count, accountList, _, _ := account_Fetch(tsql)

	return count, accountList, nil
}

// Account_GetLookup() returns a lookup list of all Account items in lookup format
func Account_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, accountList, _ := Account_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: accountList[i].SienaReference, Name: accountList[i].AccountName})
	}
	return returnList
}

// Account_GetByID() returns a single Account record
func Account_GetByID(id string) (int, dm.Account, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	tsql = tsql + " WHERE " + dm.Account_SQLSearchID + "='" + id + "'"
	_, _, accountItem, _ := account_Fetch(tsql)

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	accountItem.DealtCA, accountItem.DealtCA_props = adaptor.Account_DealtCA_Impl(adaptor.GET, id, accountItem.DealtCA, accountItem, accountItem.DealtCA_props)
	accountItem.AgainstCA, accountItem.AgainstCA_props = adaptor.Account_AgainstCA_Impl(adaptor.GET, id, accountItem.AgainstCA, accountItem, accountItem.AgainstCA_props)
	accountItem.LedgerCA, accountItem.LedgerCA_props = adaptor.Account_LedgerCA_Impl(adaptor.GET, id, accountItem.LedgerCA, accountItem, accountItem.LedgerCA_props)
	accountItem.CashBalanceCA, accountItem.CashBalanceCA_props = adaptor.Account_CashBalanceCA_Impl(adaptor.GET, id, accountItem.CashBalanceCA, accountItem, accountItem.CashBalanceCA_props)
	//
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return 1, accountItem, nil
}

// Account_GetByReverseLookup(id string) returns a single Account record using the CashBalance field as key.
func Account_GetByReverseLookup(id string) (int, dm.Account, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	tsql = tsql + " WHERE CashBalance = '" + id + "'"

	_, _, accountItem, _ := account_Fetch(tsql)

	return 1, accountItem, nil
}

// Account_DeleteByID() deletes a single Account record
func Account_Delete(id string) {

	adaptor.Account_Delete_impl(id)

}

// Account_Store() saves/stores a Account record to the database
func Account_Store(r dm.Account, req *http.Request) error {

	err, r := Account_Validate(r)
	if err == nil {
		err = account_Save(r, Audit_User(req))
	} else {
		logs.Information("Account_Store()", err.Error())
	}

	return err
}

// Account_StoreSystem() saves/stores a Account record to the database
func Account_StoreSystem(r dm.Account) error {

	err, r := Account_Validate(r)
	if err == nil {
		err = account_Save(r, Audit_Host())
	} else {
		logs.Information("Account_Store()", err.Error())
	}

	return err
}

// Account_Validate() validates for saves/stores a Account record to the database
func Account_Validate(r dm.Account) (error, dm.Account) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	r.DealtCA, r.DealtCA_props = adaptor.Account_DealtCA_Impl(adaptor.PUT, r.SienaReference, r.DealtCA, r, r.DealtCA_props)
	r.AgainstCA, r.AgainstCA_props = adaptor.Account_AgainstCA_Impl(adaptor.PUT, r.SienaReference, r.AgainstCA, r, r.AgainstCA_props)
	r.LedgerCA, r.LedgerCA_props = adaptor.Account_LedgerCA_Impl(adaptor.PUT, r.SienaReference, r.LedgerCA, r, r.LedgerCA_props)
	r.CashBalanceCA, r.CashBalanceCA_props = adaptor.Account_CashBalanceCA_Impl(adaptor.PUT, r.SienaReference, r.CashBalanceCA, r, r.CashBalanceCA_props)
	//
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	return nil, r
}

//

// account_Save() saves/stores a Account record to the database
func account_Save(r dm.Account, usr string) error {

	var err error

	if len(r.SienaReference) == 0 {
		r.SienaReference = Account_NewID(r)
	}

	// If there are fields below, create the methods in dao\account_impl.go

	r.DealtCA, err = adaptor.Account_DealtCA_OnStore_impl(r.DealtCA, r, usr)
	r.AgainstCA, err = adaptor.Account_AgainstCA_OnStore_impl(r.AgainstCA, r, usr)
	r.LedgerCA, err = adaptor.Account_LedgerCA_OnStore_impl(r.LedgerCA, r, usr)
	r.CashBalanceCA, err = adaptor.Account_CashBalanceCA_OnStore_impl(r.CashBalanceCA, r, usr)

	logs.Storing("Account", fmt.Sprintf("%s", r))

	// Please Create Functions Below in the adaptor/Account_impl.go file
	err1 := adaptor.Account_Delete_impl(r.SienaReference)
	err2 := adaptor.Account_Update_impl(r.SienaReference, r, usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}

	return err

}

// account_Fetch read all Account's
func account_Fetch(tsql string) (int, []dm.Account, dm.Account, error) {

	var recItem dm.Account
	var recList []dm.Account

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(), err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// START
		// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
		//
		recItem.SienaReference = get_String(rec, dm.Account_SienaReference_sql, "")
		recItem.CustomerSienaView = get_String(rec, dm.Account_CustomerSienaView_sql, "")
		recItem.SienaCommonRef = get_String(rec, dm.Account_SienaCommonRef_sql, "")
		recItem.Status = get_String(rec, dm.Account_Status_sql, "")
		recItem.StartDate = get_Time(rec, dm.Account_StartDate_sql, "")
		recItem.MaturityDate = get_Time(rec, dm.Account_MaturityDate_sql, "")
		recItem.ContractNumber = get_String(rec, dm.Account_ContractNumber_sql, "")
		recItem.ExternalReference = get_String(rec, dm.Account_ExternalReference_sql, "")
		recItem.CCY = get_String(rec, dm.Account_CCY_sql, "")
		recItem.Book = get_String(rec, dm.Account_Book_sql, "")
		recItem.MandatedUser = get_String(rec, dm.Account_MandatedUser_sql, "")
		recItem.BackOfficeNotes = get_String(rec, dm.Account_BackOfficeNotes_sql, "")
		recItem.CashBalance = get_Float(rec, dm.Account_CashBalance_sql, "0.00")
		recItem.AccountNumber = get_String(rec, dm.Account_AccountNumber_sql, "")
		recItem.AccountName = get_String(rec, dm.Account_AccountName_sql, "")
		recItem.LedgerBalance = get_Float(rec, dm.Account_LedgerBalance_sql, "0.00")
		recItem.Portfolio = get_String(rec, dm.Account_Portfolio_sql, "")
		recItem.AgreementId = get_Int(rec, dm.Account_AgreementId_sql, "0")
		recItem.BackOfficeRefNo = get_String(rec, dm.Account_BackOfficeRefNo_sql, "")
		recItem.ISIN = get_String(rec, dm.Account_ISIN_sql, "")
		recItem.UTI = get_String(rec, dm.Account_UTI_sql, "")
		recItem.CCYName = get_String(rec, dm.Account_CCYName_sql, "")
		recItem.BookName = get_String(rec, dm.Account_BookName_sql, "")
		recItem.PortfolioName = get_String(rec, dm.Account_PortfolioName_sql, "")
		recItem.Centre = get_String(rec, dm.Account_Centre_sql, "")
		recItem.DealTypeKey = get_String(rec, dm.Account_DealTypeKey_sql, "")
		recItem.DealTypeShortName = get_String(rec, dm.Account_DealTypeShortName_sql, "")
		recItem.InternalId = get_Int(rec, dm.Account_InternalId_sql, "0")
		recItem.InternalDeleted = get_Time(rec, dm.Account_InternalDeleted_sql, "")
		recItem.UpdatedTransactionId = get_String(rec, dm.Account_UpdatedTransactionId_sql, "")
		recItem.UpdatedUserId = get_String(rec, dm.Account_UpdatedUserId_sql, "")
		recItem.UpdatedDateTime = get_Time(rec, dm.Account_UpdatedDateTime_sql, "")
		recItem.DeletedTransactionId = get_String(rec, dm.Account_DeletedTransactionId_sql, "")
		recItem.DeletedUserId = get_String(rec, dm.Account_DeletedUserId_sql, "")
		recItem.ChangeType = get_String(rec, dm.Account_ChangeType_sql, "")
		recItem.CCYDp = get_Int(rec, dm.Account_CCYDp_sql, "0")
		recItem.CompID = get_String(rec, dm.Account_CompID_sql, "")
		recItem.Firm = get_String(rec, dm.Account_Firm_sql, "")
		recItem.DealType = get_String(rec, dm.Account_DealType_sql, "")
		recItem.FullDealType = get_String(rec, dm.Account_FullDealType_sql, "")
		recItem.DealingInterface = get_String(rec, dm.Account_DealingInterface_sql, "")
		recItem.DealtAmount = get_Float(rec, dm.Account_DealtAmount_sql, "0.00")
		recItem.ParentContractNumber = get_String(rec, dm.Account_ParentContractNumber_sql, "")
		recItem.InterestFrequency = get_String(rec, dm.Account_InterestFrequency_sql, "")
		recItem.InterestAction = get_String(rec, dm.Account_InterestAction_sql, "")
		recItem.InterestStrategy = get_String(rec, dm.Account_InterestStrategy_sql, "")
		recItem.InterestBasis = get_String(rec, dm.Account_InterestBasis_sql, "")
		recItem.SienaDealer = get_String(rec, dm.Account_SienaDealer_sql, "")
		recItem.DealOwner = get_String(rec, dm.Account_DealOwner_sql, "")
		recItem.OriginUser = get_String(rec, dm.Account_OriginUser_sql, "")
		recItem.EditedByUser = get_String(rec, dm.Account_EditedByUser_sql, "")
		recItem.DealOwnerMnemonic = get_String(rec, dm.Account_DealOwnerMnemonic_sql, "")
		recItem.UTCOriginTime = get_String(rec, dm.Account_UTCOriginTime_sql, "")
		recItem.UTCUpdateTime = get_String(rec, dm.Account_UTCUpdateTime_sql, "")
		recItem.CustomerStatementNotes = get_String(rec, dm.Account_CustomerStatementNotes_sql, "")
		recItem.NotesMargin = get_Float(rec, dm.Account_NotesMargin_sql, "0.00")
		recItem.RequestedBy = get_String(rec, dm.Account_RequestedBy_sql, "")
		recItem.EditReason = get_String(rec, dm.Account_EditReason_sql, "")
		recItem.EditOtherReason = get_String(rec, dm.Account_EditOtherReason_sql, "")
		recItem.NoticeDays = get_Int(rec, dm.Account_NoticeDays_sql, "0")
		recItem.DebitFrequency = get_String(rec, dm.Account_DebitFrequency_sql, "")
		recItem.CreditFrequency = get_String(rec, dm.Account_CreditFrequency_sql, "")
		recItem.EURAmount = get_Float(rec, dm.Account_EURAmount_sql, "0.00")
		recItem.EUROtherAmount = get_Float(rec, dm.Account_EUROtherAmount_sql, "0.00")
		recItem.PaymentSystemSienaView = get_String(rec, dm.Account_PaymentSystemSienaView_sql, "")
		recItem.PaymentSystemExternalView = get_String(rec, dm.Account_PaymentSystemExternalView_sql, "")

		// If there are fields below, create the methods in adaptor\Account_impl.go

		recItem.DealtCA = adaptor.Account_DealtCA_OnFetch_impl(recItem)
		recItem.AgainstCA = adaptor.Account_AgainstCA_OnFetch_impl(recItem)
		recItem.LedgerCA = adaptor.Account_LedgerCA_OnFetch_impl(recItem)
		recItem.CashBalanceCA = adaptor.Account_CashBalanceCA_OnFetch_impl(recItem)

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

func Account_NewID(r dm.Account) string {

	id := uuid.New().String()

	return id
}

// account_Fetch read all Account's
func Account_New() (int, []dm.Account, dm.Account, error) {

	var r = dm.Account{}
	var rList []dm.Account

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	r.DealtCA, r.DealtCA_props = adaptor.Account_DealtCA_Impl(adaptor.NEW, r.SienaReference, r.DealtCA, r, r.DealtCA_props)
	r.AgainstCA, r.AgainstCA_props = adaptor.Account_AgainstCA_Impl(adaptor.NEW, r.SienaReference, r.AgainstCA, r, r.AgainstCA_props)
	r.LedgerCA, r.LedgerCA_props = adaptor.Account_LedgerCA_Impl(adaptor.NEW, r.SienaReference, r.LedgerCA, r, r.LedgerCA_props)
	r.CashBalanceCA, r.CashBalanceCA_props = adaptor.Account_CashBalanceCA_Impl(adaptor.NEW, r.SienaReference, r.CashBalanceCA, r, r.CashBalanceCA_props)
	//
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END

	rList = append(rList, r)

	return 1, rList, r, nil
}
