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


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Transaction_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Transaction_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

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

	logs.Storing("Transaction",fmt.Sprintf("%s", r))

	if len(r.SienaReference) == 0 {
		r.SienaReference = Transaction_NewID(r)
	}


// Please Create Functions Below in the adaptor/Transaction_impl.go file
	err1 := adaptor.Transaction_Delete_Impl(r.SienaReference)
	err2 := adaptor.Transaction_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// transaction_Fetch read all employees
func transaction_Fetch(tsql string) (int, []dm.Transaction, dm.Transaction, error) {

	var recItem dm.Transaction
	var recList []dm.Transaction

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.SienaReference  = get_String(rec, dm.Transaction_SienaReference, "")
   recItem.Status  = get_String(rec, dm.Transaction_Status, "")
   recItem.ValueDate  = get_Time(rec, dm.Transaction_ValueDate, "")
   recItem.MaturityDate  = get_Time(rec, dm.Transaction_MaturityDate, "")
   recItem.ContractNumber  = get_String(rec, dm.Transaction_ContractNumber, "")
   recItem.ExternalReference  = get_String(rec, dm.Transaction_ExternalReference, "")
   recItem.Book  = get_String(rec, dm.Transaction_Book, "")
   recItem.MandatedUser  = get_String(rec, dm.Transaction_MandatedUser, "")
   recItem.Portfolio  = get_String(rec, dm.Transaction_Portfolio, "")
   recItem.AgreementId  = get_Int(rec, dm.Transaction_AgreementId, "0")
   recItem.BackOfficeRefNo  = get_String(rec, dm.Transaction_BackOfficeRefNo, "")
   recItem.ISIN  = get_String(rec, dm.Transaction_ISIN, "")
   recItem.UTI  = get_String(rec, dm.Transaction_UTI, "")
   recItem.BookName  = get_String(rec, dm.Transaction_BookName, "")
   recItem.Centre  = get_String(rec, dm.Transaction_Centre, "")
   recItem.Firm  = get_String(rec, dm.Transaction_Firm, "")
   recItem.DealTypeShortName  = get_String(rec, dm.Transaction_DealTypeShortName, "")
   recItem.FullDealType  = get_String(rec, dm.Transaction_FullDealType, "")
   recItem.TradeDate  = get_Time(rec, dm.Transaction_TradeDate, "")
   recItem.DealtCcy  = get_String(rec, dm.Transaction_DealtCcy, "")
   recItem.DealtAmount  = get_Float(rec, dm.Transaction_DealtAmount, "0.00")
   recItem.AgainstAmount  = get_Float(rec, dm.Transaction_AgainstAmount, "0.00")
   recItem.AgainstCcy  = get_String(rec, dm.Transaction_AgainstCcy, "")
   recItem.AllInRate  = get_Float(rec, dm.Transaction_AllInRate, "0.00")
   recItem.MktRate  = get_Float(rec, dm.Transaction_MktRate, "0.00")
   recItem.SettleCcy  = get_String(rec, dm.Transaction_SettleCcy, "")
   recItem.Direction  = get_String(rec, dm.Transaction_Direction, "")
   recItem.NpvRate  = get_Float(rec, dm.Transaction_NpvRate, "0.00")
   recItem.OriginUser  = get_String(rec, dm.Transaction_OriginUser, "")
   recItem.PayInstruction  = get_String(rec, dm.Transaction_PayInstruction, "")
   recItem.ReceiptInstruction  = get_String(rec, dm.Transaction_ReceiptInstruction, "")
   recItem.NIName  = get_String(rec, dm.Transaction_NIName, "")
   recItem.CCYPair  = get_String(rec, dm.Transaction_CCYPair, "")
   recItem.Instrument  = get_String(rec, dm.Transaction_Instrument, "")
   recItem.PortfolioName  = get_String(rec, dm.Transaction_PortfolioName, "")
   recItem.RVDate  = get_Time(rec, dm.Transaction_RVDate, "")
   recItem.RVMTM  = get_Float(rec, dm.Transaction_RVMTM, "0.00")
   recItem.CounterBook  = get_String(rec, dm.Transaction_CounterBook, "")
   recItem.CounterBookName  = get_String(rec, dm.Transaction_CounterBookName, "")
   recItem.Party  = get_String(rec, dm.Transaction_Party, "")
   recItem.PartyName  = get_String(rec, dm.Transaction_PartyName, "")
   recItem.NameCentre  = get_String(rec, dm.Transaction_NameCentre, "")
   recItem.NameFirm  = get_String(rec, dm.Transaction_NameFirm, "")
   recItem.CustomerExternalView  = get_String(rec, dm.Transaction_CustomerExternalView, "")
   recItem.CustomerSienaView  = get_String(rec, dm.Transaction_CustomerSienaView, "")
   recItem.CompID  = get_String(rec, dm.Transaction_CompID, "")
   recItem.SienaDealer  = get_String(rec, dm.Transaction_SienaDealer, "")
   recItem.DealOwner  = get_String(rec, dm.Transaction_DealOwner, "")
   recItem.DealOwnerMnemonic  = get_String(rec, dm.Transaction_DealOwnerMnemonic, "")
   recItem.EditedByUser  = get_String(rec, dm.Transaction_EditedByUser, "")
   recItem.UTCOriginTime  = get_String(rec, dm.Transaction_UTCOriginTime, "")
   recItem.UTCUpdateTime  = get_String(rec, dm.Transaction_UTCUpdateTime, "")
   recItem.MarginTrading  = get_Bool(rec, dm.Transaction_MarginTrading, "True")
   recItem.SwapPoints  = get_Float(rec, dm.Transaction_SwapPoints, "0.00")
   recItem.SpotDate  = get_Time(rec, dm.Transaction_SpotDate, "")
   recItem.SpotRate  = get_Float(rec, dm.Transaction_SpotRate, "0.00")
   recItem.MktSpotRate  = get_Float(rec, dm.Transaction_MktSpotRate, "0.00")
   recItem.SpotSalesMargin  = get_Float(rec, dm.Transaction_SpotSalesMargin, "0.00")
   recItem.SpotChlMargin  = get_Float(rec, dm.Transaction_SpotChlMargin, "0.00")
   recItem.RerouteCcy  = get_String(rec, dm.Transaction_RerouteCcy, "")
   recItem.CustomerPayInstruction  = get_String(rec, dm.Transaction_CustomerPayInstruction, "")
   recItem.CustomerReceiptInstruction  = get_String(rec, dm.Transaction_CustomerReceiptInstruction, "")
   recItem.BackOfficeNotes  = get_String(rec, dm.Transaction_BackOfficeNotes, "")
   recItem.CustomerStatementNotes  = get_String(rec, dm.Transaction_CustomerStatementNotes, "")
   recItem.NotesMargin  = get_Float(rec, dm.Transaction_NotesMargin, "0.00")
   recItem.RequestedBy  = get_String(rec, dm.Transaction_RequestedBy, "")
   recItem.EditReason  = get_String(rec, dm.Transaction_EditReason, "")
   recItem.EditOtherReason  = get_String(rec, dm.Transaction_EditOtherReason, "")
   recItem.NICleanPrice  = get_Float(rec, dm.Transaction_NICleanPrice, "0.00")
   recItem.NIDirtyPrice  = get_Float(rec, dm.Transaction_NIDirtyPrice, "0.00")
   recItem.NIYield  = get_Float(rec, dm.Transaction_NIYield, "0.00")
   recItem.NIClearingSystem  = get_String(rec, dm.Transaction_NIClearingSystem, "")
   recItem.Acceptor  = get_String(rec, dm.Transaction_Acceptor, "")
   recItem.NIDiscount  = get_Float(rec, dm.Transaction_NIDiscount, "0.00")
   recItem.FastPay  = get_Bool(rec, dm.Transaction_FastPay, "True")
   recItem.PaymentFee  = get_Float(rec, dm.Transaction_PaymentFee, "0.00")
   recItem.PaymentFeePolicy  = get_String(rec, dm.Transaction_PaymentFeePolicy, "")
   recItem.PaymentReason  = get_String(rec, dm.Transaction_PaymentReason, "")
   recItem.PaymentDate  = get_Time(rec, dm.Transaction_PaymentDate, "")
   recItem.SettlementDate  = get_Time(rec, dm.Transaction_SettlementDate, "")
   recItem.FixingDate  = get_Time(rec, dm.Transaction_FixingDate, "")
   recItem.VenueUTI  = get_String(rec, dm.Transaction_VenueUTI, "")
   recItem.EditVersion  = get_Int(rec, dm.Transaction_EditVersion, "0")
   recItem.BrokeragePercentage  = get_Float(rec, dm.Transaction_BrokeragePercentage, "0.00")
   recItem.BrokerageAmount  = get_Float(rec, dm.Transaction_BrokerageAmount, "0.00")
   recItem.BrokerageCurrency  = get_String(rec, dm.Transaction_BrokerageCurrency, "")
   recItem.BrokerageDate  = get_Time(rec, dm.Transaction_BrokerageDate, "")
   recItem.AccountName  = get_String(rec, dm.Transaction_AccountName, "")
   recItem.AccountNumber  = get_String(rec, dm.Transaction_AccountNumber, "")
   recItem.CashBalance  = get_Float(rec, dm.Transaction_CashBalance, "0.00")
   recItem.DebitFrequency  = get_String(rec, dm.Transaction_DebitFrequency, "")
   recItem.CreditFrequency  = get_String(rec, dm.Transaction_CreditFrequency, "")
   recItem.ManuallyQuoted  = get_String(rec, dm.Transaction_ManuallyQuoted, "")
   recItem.LedgerBalance  = get_Float(rec, dm.Transaction_LedgerBalance, "0.00")
   recItem.SettAmtOutstanding  = get_Float(rec, dm.Transaction_SettAmtOutstanding, "0.00")
   recItem.FeePercentage  = get_Float(rec, dm.Transaction_FeePercentage, "0.00")
   recItem.FeeAmount  = get_Float(rec, dm.Transaction_FeeAmount, "0.00")
   recItem.Venue  = get_String(rec, dm.Transaction_Venue, "")
   recItem.EURAmount  = get_Float(rec, dm.Transaction_EURAmount, "0.00")
   recItem.EUROtherAmount  = get_Float(rec, dm.Transaction_EUROtherAmount, "0.00")
   recItem.LEI  = get_String(rec, dm.Transaction_LEI, "")
   recItem.Equity  = get_String(rec, dm.Transaction_Equity, "")
   recItem.Shares  = get_Int(rec, dm.Transaction_Shares, "0")
   recItem.QuoteExpiryDate  = get_Time(rec, dm.Transaction_QuoteExpiryDate, "")
   recItem.Commodity  = get_String(rec, dm.Transaction_Commodity, "")
   recItem.PaymentSystemSienaView  = get_String(rec, dm.Transaction_PaymentSystemSienaView, "")
   recItem.PaymentSystemExternalView  = get_String(rec, dm.Transaction_PaymentSystemExternalView, "")
   recItem.SalesProfit  = get_Float(rec, dm.Transaction_SalesProfit, "0.00")
   recItem.RejectReason  = get_String(rec, dm.Transaction_RejectReason, "")
   recItem.PaymentError  = get_String(rec, dm.Transaction_PaymentError, "")
   recItem.RepoPrincipal  = get_Float(rec, dm.Transaction_RepoPrincipal, "0.00")
   recItem.FixingFrequency  = get_String(rec, dm.Transaction_FixingFrequency, "")
// If there are fields below, create the methods in dao\Transaction_Impl.go

































































































































































































































	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
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

