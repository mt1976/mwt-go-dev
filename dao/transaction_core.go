package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/transaction.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Transaction (transaction)
// Endpoint 	        : Transaction (Ref)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:03
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

// Transaction_GetList() returns a list of all Transaction records
func Transaction_GetList() (int, []dm.Transaction, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Transaction_SQLTable)
	count, transactionList, _, _ := transaction_Fetch(tsql)
	
	return count, transactionList, nil
}



// Transaction_GetByID() returns a single Transaction record
func Transaction_GetByID(id string) (int, dm.Transaction, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Transaction_SQLTable)
	tsql = tsql + " WHERE " + dm.Transaction_SQLSearchID + "='" + id + "'"
	_, _, transactionItem, _ := transaction_Fetch(tsql)

	return 1, transactionItem, nil
}



// Transaction_DeleteByID() deletes a single Transaction record
func Transaction_Delete(id string) {


	adaptor.Transaction_Delete_impl(id)
	
}


// Transaction_Store() saves/stores a Transaction record to the database
func Transaction_Store(r dm.Transaction,req *http.Request) error {

	err := transaction_Save(r,Audit_User(req))

	return err
}

// Transaction_StoreSystem() saves/stores a Transaction record to the database
func Transaction_StoreSystem(r dm.Transaction) error {
	
	err := transaction_Save(r,Audit_Host())

	return err
}

// transaction_Save() saves/stores a Transaction record to the database
func transaction_Save(r dm.Transaction,usr string) error {

    var err error



	

	if len(r.SienaReference) == 0 {
		r.SienaReference = Transaction_NewID(r)
	}

// If there are fields below, create the methods in dao\transaction_impl.go
















































































































  r.Dealt,err = adaptor.Transaction_Dealt_OnStore_impl (r.Dealt,r,usr)
  r.Against,err = adaptor.Transaction_Against_OnStore_impl (r.Against,r,usr)


	
logs.Storing("Transaction",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Transaction_impl.go file
	err1 := adaptor.Transaction_Delete_impl(r.SienaReference)
	err2 := adaptor.Transaction_Update_impl(r.SienaReference,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// transaction_Fetch read all Transaction's
func transaction_Fetch(tsql string) (int, []dm.Transaction, dm.Transaction, error) {

	var recItem dm.Transaction
	var recList []dm.Transaction

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SienaReference  = get_String(rec, dm.Transaction_SienaReference_sql, "")
	   recItem.Status  = get_String(rec, dm.Transaction_Status_sql, "")
	   recItem.ValueDate  = get_Time(rec, dm.Transaction_ValueDate_sql, "")
	   recItem.MaturityDate  = get_Time(rec, dm.Transaction_MaturityDate_sql, "")
	   recItem.ContractNumber  = get_String(rec, dm.Transaction_ContractNumber_sql, "")
	   recItem.ExternalReference  = get_String(rec, dm.Transaction_ExternalReference_sql, "")
	   recItem.Book  = get_String(rec, dm.Transaction_Book_sql, "")
	   recItem.MandatedUser  = get_String(rec, dm.Transaction_MandatedUser_sql, "")
	   recItem.Portfolio  = get_String(rec, dm.Transaction_Portfolio_sql, "")
	   recItem.AgreementId  = get_Int(rec, dm.Transaction_AgreementId_sql, "0")
	   recItem.BackOfficeRefNo  = get_String(rec, dm.Transaction_BackOfficeRefNo_sql, "")
	   recItem.ISIN  = get_String(rec, dm.Transaction_ISIN_sql, "")
	   recItem.UTI  = get_String(rec, dm.Transaction_UTI_sql, "")
	   recItem.BookName  = get_String(rec, dm.Transaction_BookName_sql, "")
	   recItem.Centre  = get_String(rec, dm.Transaction_Centre_sql, "")
	   recItem.Firm  = get_String(rec, dm.Transaction_Firm_sql, "")
	   recItem.DealTypeShortName  = get_String(rec, dm.Transaction_DealTypeShortName_sql, "")
	   recItem.FullDealType  = get_String(rec, dm.Transaction_FullDealType_sql, "")
	   recItem.TradeDate  = get_Time(rec, dm.Transaction_TradeDate_sql, "")
	   recItem.DealtCcy  = get_String(rec, dm.Transaction_DealtCcy_sql, "")
	   recItem.DealtAmount  = get_Float(rec, dm.Transaction_DealtAmount_sql, "0.00")
	   recItem.AgainstAmount  = get_Float(rec, dm.Transaction_AgainstAmount_sql, "0.00")
	   recItem.AgainstCcy  = get_String(rec, dm.Transaction_AgainstCcy_sql, "")
	   recItem.AllInRate  = get_Float(rec, dm.Transaction_AllInRate_sql, "0.00")
	   recItem.MktRate  = get_Float(rec, dm.Transaction_MktRate_sql, "0.00")
	   recItem.SettleCcy  = get_String(rec, dm.Transaction_SettleCcy_sql, "")
	   recItem.Direction  = get_String(rec, dm.Transaction_Direction_sql, "")
	   recItem.NpvRate  = get_Float(rec, dm.Transaction_NpvRate_sql, "0.00")
	   recItem.OriginUser  = get_String(rec, dm.Transaction_OriginUser_sql, "")
	   recItem.PayInstruction  = get_String(rec, dm.Transaction_PayInstruction_sql, "")
	   recItem.ReceiptInstruction  = get_String(rec, dm.Transaction_ReceiptInstruction_sql, "")
	   recItem.NIName  = get_String(rec, dm.Transaction_NIName_sql, "")
	   recItem.CCYPair  = get_String(rec, dm.Transaction_CCYPair_sql, "")
	   recItem.Instrument  = get_String(rec, dm.Transaction_Instrument_sql, "")
	   recItem.PortfolioName  = get_String(rec, dm.Transaction_PortfolioName_sql, "")
	   recItem.RVDate  = get_Time(rec, dm.Transaction_RVDate_sql, "")
	   recItem.RVMTM  = get_Float(rec, dm.Transaction_RVMTM_sql, "0.00")
	   recItem.CounterBook  = get_String(rec, dm.Transaction_CounterBook_sql, "")
	   recItem.CounterBookName  = get_String(rec, dm.Transaction_CounterBookName_sql, "")
	   recItem.Party  = get_String(rec, dm.Transaction_Party_sql, "")
	   recItem.PartyName  = get_String(rec, dm.Transaction_PartyName_sql, "")
	   recItem.NameCentre  = get_String(rec, dm.Transaction_NameCentre_sql, "")
	   recItem.NameFirm  = get_String(rec, dm.Transaction_NameFirm_sql, "")
	   recItem.CustomerExternalView  = get_String(rec, dm.Transaction_CustomerExternalView_sql, "")
	   recItem.CustomerSienaView  = get_String(rec, dm.Transaction_CustomerSienaView_sql, "")
	   recItem.CompID  = get_String(rec, dm.Transaction_CompID_sql, "")
	   recItem.SienaDealer  = get_String(rec, dm.Transaction_SienaDealer_sql, "")
	   recItem.DealOwner  = get_String(rec, dm.Transaction_DealOwner_sql, "")
	   recItem.DealOwnerMnemonic  = get_String(rec, dm.Transaction_DealOwnerMnemonic_sql, "")
	   recItem.EditedByUser  = get_String(rec, dm.Transaction_EditedByUser_sql, "")
	   recItem.UTCOriginTime  = get_String(rec, dm.Transaction_UTCOriginTime_sql, "")
	   recItem.UTCUpdateTime  = get_String(rec, dm.Transaction_UTCUpdateTime_sql, "")
	   recItem.MarginTrading  = get_Bool(rec, dm.Transaction_MarginTrading_sql, "True")
	   recItem.SwapPoints  = get_Float(rec, dm.Transaction_SwapPoints_sql, "0.00")
	   recItem.SpotDate  = get_Time(rec, dm.Transaction_SpotDate_sql, "")
	   recItem.SpotRate  = get_Float(rec, dm.Transaction_SpotRate_sql, "0.00")
	   recItem.MktSpotRate  = get_Float(rec, dm.Transaction_MktSpotRate_sql, "0.00")
	   recItem.SpotSalesMargin  = get_Float(rec, dm.Transaction_SpotSalesMargin_sql, "0.00")
	   recItem.SpotChlMargin  = get_Float(rec, dm.Transaction_SpotChlMargin_sql, "0.00")
	   recItem.RerouteCcy  = get_String(rec, dm.Transaction_RerouteCcy_sql, "")
	   recItem.CustomerPayInstruction  = get_String(rec, dm.Transaction_CustomerPayInstruction_sql, "")
	   recItem.CustomerReceiptInstruction  = get_String(rec, dm.Transaction_CustomerReceiptInstruction_sql, "")
	   recItem.BackOfficeNotes  = get_String(rec, dm.Transaction_BackOfficeNotes_sql, "")
	   recItem.CustomerStatementNotes  = get_String(rec, dm.Transaction_CustomerStatementNotes_sql, "")
	   recItem.NotesMargin  = get_Float(rec, dm.Transaction_NotesMargin_sql, "0.00")
	   recItem.RequestedBy  = get_String(rec, dm.Transaction_RequestedBy_sql, "")
	   recItem.EditReason  = get_String(rec, dm.Transaction_EditReason_sql, "")
	   recItem.EditOtherReason  = get_String(rec, dm.Transaction_EditOtherReason_sql, "")
	   recItem.NICleanPrice  = get_Float(rec, dm.Transaction_NICleanPrice_sql, "0.00")
	   recItem.NIDirtyPrice  = get_Float(rec, dm.Transaction_NIDirtyPrice_sql, "0.00")
	   recItem.NIYield  = get_Float(rec, dm.Transaction_NIYield_sql, "0.00")
	   recItem.NIClearingSystem  = get_String(rec, dm.Transaction_NIClearingSystem_sql, "")
	   recItem.Acceptor  = get_String(rec, dm.Transaction_Acceptor_sql, "")
	   recItem.NIDiscount  = get_Float(rec, dm.Transaction_NIDiscount_sql, "0.00")
	   recItem.FastPay  = get_Bool(rec, dm.Transaction_FastPay_sql, "True")
	   recItem.PaymentFee  = get_Float(rec, dm.Transaction_PaymentFee_sql, "0.00")
	   recItem.PaymentFeePolicy  = get_String(rec, dm.Transaction_PaymentFeePolicy_sql, "")
	   recItem.PaymentReason  = get_String(rec, dm.Transaction_PaymentReason_sql, "")
	   recItem.PaymentDate  = get_Time(rec, dm.Transaction_PaymentDate_sql, "")
	   recItem.SettlementDate  = get_Time(rec, dm.Transaction_SettlementDate_sql, "")
	   recItem.FixingDate  = get_Time(rec, dm.Transaction_FixingDate_sql, "")
	   recItem.VenueUTI  = get_String(rec, dm.Transaction_VenueUTI_sql, "")
	   recItem.EditVersion  = get_Int(rec, dm.Transaction_EditVersion_sql, "0")
	   recItem.BrokeragePercentage  = get_Float(rec, dm.Transaction_BrokeragePercentage_sql, "0.00")
	   recItem.BrokerageAmount  = get_Float(rec, dm.Transaction_BrokerageAmount_sql, "0.00")
	   recItem.BrokerageCurrency  = get_String(rec, dm.Transaction_BrokerageCurrency_sql, "")
	   recItem.BrokerageDate  = get_Time(rec, dm.Transaction_BrokerageDate_sql, "")
	   recItem.AccountName  = get_String(rec, dm.Transaction_AccountName_sql, "")
	   recItem.AccountNumber  = get_String(rec, dm.Transaction_AccountNumber_sql, "")
	   recItem.CashBalance  = get_Float(rec, dm.Transaction_CashBalance_sql, "0.00")
	   recItem.DebitFrequency  = get_String(rec, dm.Transaction_DebitFrequency_sql, "")
	   recItem.CreditFrequency  = get_String(rec, dm.Transaction_CreditFrequency_sql, "")
	   recItem.ManuallyQuoted  = get_String(rec, dm.Transaction_ManuallyQuoted_sql, "")
	   recItem.LedgerBalance  = get_Float(rec, dm.Transaction_LedgerBalance_sql, "0.00")
	   recItem.SettAmtOutstanding  = get_Float(rec, dm.Transaction_SettAmtOutstanding_sql, "0.00")
	   recItem.FeePercentage  = get_Float(rec, dm.Transaction_FeePercentage_sql, "0.00")
	   recItem.FeeAmount  = get_Float(rec, dm.Transaction_FeeAmount_sql, "0.00")
	   recItem.Venue  = get_String(rec, dm.Transaction_Venue_sql, "")
	   recItem.EURAmount  = get_Float(rec, dm.Transaction_EURAmount_sql, "0.00")
	   recItem.EUROtherAmount  = get_Float(rec, dm.Transaction_EUROtherAmount_sql, "0.00")
	   recItem.LEI  = get_String(rec, dm.Transaction_LEI_sql, "")
	   recItem.Equity  = get_String(rec, dm.Transaction_Equity_sql, "")
	   recItem.Shares  = get_Int(rec, dm.Transaction_Shares_sql, "0")
	   recItem.QuoteExpiryDate  = get_Time(rec, dm.Transaction_QuoteExpiryDate_sql, "")
	   recItem.Commodity  = get_String(rec, dm.Transaction_Commodity_sql, "")
	   recItem.PaymentSystemSienaView  = get_String(rec, dm.Transaction_PaymentSystemSienaView_sql, "")
	   recItem.PaymentSystemExternalView  = get_String(rec, dm.Transaction_PaymentSystemExternalView_sql, "")
	   recItem.SalesProfit  = get_Float(rec, dm.Transaction_SalesProfit_sql, "0.00")
	   recItem.RejectReason  = get_String(rec, dm.Transaction_RejectReason_sql, "")
	   recItem.PaymentError  = get_String(rec, dm.Transaction_PaymentError_sql, "")
	   recItem.RepoPrincipal  = get_Float(rec, dm.Transaction_RepoPrincipal_sql, "0.00")
	   recItem.FixingFrequency  = get_String(rec, dm.Transaction_FixingFrequency_sql, "")
	
	
	
	// If there are fields below, create the methods in adaptor\Transaction_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	   recItem.Dealt  = adaptor.Transaction_Dealt_OnFetch_impl (recItem)
	   recItem.Against  = adaptor.Transaction_Against_OnFetch_impl (recItem)
	
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
	


func Transaction_NewID(r dm.Transaction) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

