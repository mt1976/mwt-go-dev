package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/dealtypefundamental.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : DealTypeFundamental (dealtypefundamental)
// Endpoint 	        : DealTypeFundamental (DealTypeKey)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:51
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

// DealTypeFundamental_GetList() returns a list of all DealTypeFundamental records
func DealTypeFundamental_GetList() (int, []dm.DealTypeFundamental, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.DealTypeFundamental_SQLTable)
	count, dealtypefundamentalList, _, _ := dealtypefundamental_Fetch(tsql)
	
	return count, dealtypefundamentalList, nil
}



// DealTypeFundamental_GetByID() returns a single DealTypeFundamental record
func DealTypeFundamental_GetByID(id string) (int, dm.DealTypeFundamental, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.DealTypeFundamental_SQLTable)
	tsql = tsql + " WHERE " + dm.DealTypeFundamental_SQLSearchID + "='" + id + "'"
	_, _, dealtypefundamentalItem, _ := dealtypefundamental_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, dealtypefundamentalItem, nil
}



// DealTypeFundamental_DeleteByID() deletes a single DealTypeFundamental record
func DealTypeFundamental_Delete(id string) {


	adaptor.DealTypeFundamental_Delete_impl(id)
	
}


// DealTypeFundamental_Store() saves/stores a DealTypeFundamental record to the database
func DealTypeFundamental_Store(r dm.DealTypeFundamental,req *http.Request) error {

	err, r := DealTypeFundamental_Validate(r)
	if err == nil {
		err = dealtypefundamental_Save(r, Audit_User(req))
	} else {
		logs.Information("DealTypeFundamental_Store()", err.Error())
	}

	return err
}

// DealTypeFundamental_StoreSystem() saves/stores a DealTypeFundamental record to the database
func DealTypeFundamental_StoreSystem(r dm.DealTypeFundamental) error {
	
	err, r := DealTypeFundamental_Validate(r)
	if err == nil {
		err = dealtypefundamental_Save(r, Audit_Host())
	} else {
		logs.Information("DealTypeFundamental_Store()", err.Error())
	}

	return err
}

// DealTypeFundamental_Validate() validates for saves/stores a DealTypeFundamental record to the database
func DealTypeFundamental_Validate(r dm.DealTypeFundamental) (error,dm.DealTypeFundamental) {
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

// dealtypefundamental_Save() saves/stores a DealTypeFundamental record to the database
func dealtypefundamental_Save(r dm.DealTypeFundamental,usr string) error {

    var err error



	

	if len(r.DealTypeKey) == 0 {
		r.DealTypeKey = DealTypeFundamental_NewID(r)
	}

// If there are fields below, create the methods in dao\dealtypefundamental_impl.go




























































































	
logs.Storing("DealTypeFundamental",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/DealTypeFundamental_impl.go file
	err1 := adaptor.DealTypeFundamental_Delete_impl(r.DealTypeKey)
	err2 := adaptor.DealTypeFundamental_Update_impl(r.DealTypeKey,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// dealtypefundamental_Fetch read all DealTypeFundamental's
func dealtypefundamental_Fetch(tsql string) (int, []dm.DealTypeFundamental, dm.DealTypeFundamental, error) {

	var recItem dm.DealTypeFundamental
	var recList []dm.DealTypeFundamental

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.DealTypeKey  = get_String(rec, dm.DealTypeFundamental_DealTypeKey_sql, "")
	   recItem.Amendment  = get_Bool(rec, dm.DealTypeFundamental_Amendment_sql, "True")
	   recItem.DefaultFrequency  = get_Int(rec, dm.DealTypeFundamental_DefaultFrequency_sql, "0")
	   recItem.Directions  = get_String(rec, dm.DealTypeFundamental_Directions_sql, "")
	   recItem.SettledTermDealType  = get_String(rec, dm.DealTypeFundamental_SettledTermDealType_sql, "")
	   recItem.Optn  = get_Bool(rec, dm.DealTypeFundamental_Optn_sql, "True")
	   recItem.AllowPledge  = get_Bool(rec, dm.DealTypeFundamental_AllowPledge_sql, "True")
	   recItem.Takeup  = get_Bool(rec, dm.DealTypeFundamental_Takeup_sql, "True")
	   recItem.MismatchDealType  = get_String(rec, dm.DealTypeFundamental_MismatchDealType_sql, "")
	   recItem.SettledHedgeTermDealType  = get_String(rec, dm.DealTypeFundamental_SettledHedgeTermDealType_sql, "")
	   recItem.SettlementCode  = get_String(rec, dm.DealTypeFundamental_SettlementCode_sql, "")
	   recItem.TermSubType  = get_String(rec, dm.DealTypeFundamental_TermSubType_sql, "")
	   recItem.FundingDealType  = get_String(rec, dm.DealTypeFundamental_FundingDealType_sql, "")
	   recItem.TransferType  = get_String(rec, dm.DealTypeFundamental_TransferType_sql, "")
	   recItem.TermDealType  = get_String(rec, dm.DealTypeFundamental_TermDealType_sql, "")
	   recItem.NegotiableInstrumentType  = get_String(rec, dm.DealTypeFundamental_NegotiableInstrumentType_sql, "")
	   recItem.Mismatch  = get_Bool(rec, dm.DealTypeFundamental_Mismatch_sql, "True")
	   recItem.ComplexTransferSubType  = get_String(rec, dm.DealTypeFundamental_ComplexTransferSubType_sql, "")
	   recItem.LayOffDealType  = get_String(rec, dm.DealTypeFundamental_LayOffDealType_sql, "")
	   recItem.NIAccount  = get_Int(rec, dm.DealTypeFundamental_NIAccount_sql, "0")
	   recItem.SimpleMMsubtype  = get_Int(rec, dm.DealTypeFundamental_SimpleMMsubtype_sql, "0")
	   recItem.SwapDealType  = get_String(rec, dm.DealTypeFundamental_SwapDealType_sql, "")
	   recItem.Positions  = get_String(rec, dm.DealTypeFundamental_Positions_sql, "")
	   recItem.OptionOutright  = get_String(rec, dm.DealTypeFundamental_OptionOutright_sql, "")
	   recItem.SettledHedgeSpotDealType  = get_String(rec, dm.DealTypeFundamental_SettledHedgeSpotDealType_sql, "")
	   recItem.StraightThroughInterestMethod  = get_Bool(rec, dm.DealTypeFundamental_StraightThroughInterestMethod_sql, "True")
	   recItem.SubType  = get_String(rec, dm.DealTypeFundamental_SubType_sql, "")
	   recItem.Rollover  = get_Bool(rec, dm.DealTypeFundamental_Rollover_sql, "True")
	   recItem.DefaultIssuer  = get_String(rec, dm.DealTypeFundamental_DefaultIssuer_sql, "")
	   recItem.DefaultStartDate  = get_Int(rec, dm.DealTypeFundamental_DefaultStartDate_sql, "0")
	   recItem.Fee  = get_String(rec, dm.DealTypeFundamental_Fee_sql, "")
	   recItem.NDF  = get_Bool(rec, dm.DealTypeFundamental_NDF_sql, "True")
	   recItem.FXFX  = get_Bool(rec, dm.DealTypeFundamental_FXFX_sql, "True")
	   recItem.ONIA  = get_Bool(rec, dm.DealTypeFundamental_ONIA_sql, "True")
	   recItem.MarginSubType  = get_Int(rec, dm.DealTypeFundamental_MarginSubType_sql, "0")
	   recItem.TransferDealType  = get_String(rec, dm.DealTypeFundamental_TransferDealType_sql, "")
	   recItem.IsFX  = get_Bool(rec, dm.DealTypeFundamental_IsFX_sql, "True")
	   recItem.Ordr  = get_String(rec, dm.DealTypeFundamental_Ordr_sql, "")
	   recItem.OptionStyle  = get_String(rec, dm.DealTypeFundamental_OptionStyle_sql, "")
	   recItem.SpotDealType  = get_String(rec, dm.DealTypeFundamental_SpotDealType_sql, "")
	   recItem.CanIssue  = get_Bool(rec, dm.DealTypeFundamental_CanIssue_sql, "True")
	   recItem.CanShort  = get_Bool(rec, dm.DealTypeFundamental_CanShort_sql, "True")
	   recItem.FXMarginTradingType  = get_Int(rec, dm.DealTypeFundamental_FXMarginTradingType_sql, "0")
	   recItem.Internal  = get_Bool(rec, dm.DealTypeFundamental_Internal_sql, "True")
	   recItem.TicketBasename  = get_String(rec, dm.DealTypeFundamental_TicketBasename_sql, "")
	   recItem.InterestRateFutureType  = get_String(rec, dm.DealTypeFundamental_InterestRateFutureType_sql, "")
	   recItem.TradingLimitProductCode  = get_String(rec, dm.DealTypeFundamental_TradingLimitProductCode_sql, "")
	   recItem.Forward  = get_Bool(rec, dm.DealTypeFundamental_Forward_sql, "True")
	   recItem.MaturityNotificationPeriod  = get_String(rec, dm.DealTypeFundamental_MaturityNotificationPeriod_sql, "")
	   recItem.NotificationEvents  = get_String(rec, dm.DealTypeFundamental_NotificationEvents_sql, "")
	   recItem.SwapSubType  = get_String(rec, dm.DealTypeFundamental_SwapSubType_sql, "")
	   recItem.ProductCode  = get_String(rec, dm.DealTypeFundamental_ProductCode_sql, "")
	   recItem.Funding  = get_Bool(rec, dm.DealTypeFundamental_Funding_sql, "True")
	   recItem.AllocationPricing  = get_String(rec, dm.DealTypeFundamental_AllocationPricing_sql, "")
	   recItem.CancelPeriod  = get_String(rec, dm.DealTypeFundamental_CancelPeriod_sql, "")
	   recItem.MMMarginTradingType  = get_Int(rec, dm.DealTypeFundamental_MMMarginTradingType_sql, "0")
	   recItem.OptionSpot  = get_String(rec, dm.DealTypeFundamental_OptionSpot_sql, "")
	   recItem.Transfer  = get_Bool(rec, dm.DealTypeFundamental_Transfer_sql, "True")
	   recItem.NotificationPeriod  = get_String(rec, dm.DealTypeFundamental_NotificationPeriod_sql, "")
	   recItem.Paymentdateshift  = get_Int(rec, dm.DealTypeFundamental_Paymentdateshift_sql, "0")
	   recItem.CloseOut  = get_Bool(rec, dm.DealTypeFundamental_CloseOut_sql, "True")
	   recItem.FXOptionPricing  = get_String(rec, dm.DealTypeFundamental_FXOptionPricing_sql, "")
	   recItem.SettledHedgeOutrightDealType  = get_String(rec, dm.DealTypeFundamental_SettledHedgeOutrightDealType_sql, "")
	   recItem.RepoBond  = get_String(rec, dm.DealTypeFundamental_RepoBond_sql, "")
	   recItem.RepoTerm  = get_String(rec, dm.DealTypeFundamental_RepoTerm_sql, "")
	   recItem.RepoType  = get_Int(rec, dm.DealTypeFundamental_RepoType_sql, "0")
	   recItem.DateRule  = get_String(rec, dm.DealTypeFundamental_DateRule_sql, "")
	   recItem.CorpTransferDealType  = get_String(rec, dm.DealTypeFundamental_CorpTransferDealType_sql, "")
	   recItem.GenerateFXImage  = get_Bool(rec, dm.DealTypeFundamental_GenerateFXImage_sql, "True")
	   recItem.Variant  = get_String(rec, dm.DealTypeFundamental_Variant_sql, "")
	   recItem.HedgeTermDealType  = get_String(rec, dm.DealTypeFundamental_HedgeTermDealType_sql, "")
	   recItem.PricingModel  = get_String(rec, dm.DealTypeFundamental_PricingModel_sql, "")
	   recItem.HedgeOutrightDealType  = get_String(rec, dm.DealTypeFundamental_HedgeOutrightDealType_sql, "")
	   recItem.Fixing  = get_Bool(rec, dm.DealTypeFundamental_Fixing_sql, "True")
	   recItem.Payment  = get_Bool(rec, dm.DealTypeFundamental_Payment_sql, "True")
	   recItem.MT  = get_Bool(rec, dm.DealTypeFundamental_MT_sql, "True")
	   recItem.SettlementInstructionStyle  = get_String(rec, dm.DealTypeFundamental_SettlementInstructionStyle_sql, "")
	   recItem.QuoteHistoryRequired  = get_Bool(rec, dm.DealTypeFundamental_QuoteHistoryRequired_sql, "True")
	   recItem.Brokerage  = get_Bool(rec, dm.DealTypeFundamental_Brokerage_sql, "True")
	   recItem.ExposureDisabled  = get_Bool(rec, dm.DealTypeFundamental_ExposureDisabled_sql, "True")
	   recItem.CreditLine  = get_String(rec, dm.DealTypeFundamental_CreditLine_sql, "")
	   recItem.Encumbered  = get_Bool(rec, dm.DealTypeFundamental_Encumbered_sql, "True")
	   recItem.InternalId  = get_Int(rec, dm.DealTypeFundamental_InternalId_sql, "0")
	   recItem.InternalDeleted  = get_Time(rec, dm.DealTypeFundamental_InternalDeleted_sql, "")
	   recItem.UpdatedTransactionId  = get_String(rec, dm.DealTypeFundamental_UpdatedTransactionId_sql, "")
	   recItem.UpdatedUserId  = get_String(rec, dm.DealTypeFundamental_UpdatedUserId_sql, "")
	   recItem.UpdatedDateTime  = get_Time(rec, dm.DealTypeFundamental_UpdatedDateTime_sql, "")
	   recItem.DeletedTransactionId  = get_String(rec, dm.DealTypeFundamental_DeletedTransactionId_sql, "")
	   recItem.DeletedUserId  = get_String(rec, dm.DealTypeFundamental_DeletedUserId_sql, "")
	   recItem.ChangeType  = get_String(rec, dm.DealTypeFundamental_ChangeType_sql, "")
	
	// If there are fields below, create the methods in adaptor\DealTypeFundamental_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func DealTypeFundamental_NewID(r dm.DealTypeFundamental) string {
	
			id := uuid.New().String()
	
	return id
}



// dealtypefundamental_Fetch read all DealTypeFundamental's
func DealTypeFundamental_New() (int, []dm.DealTypeFundamental, dm.DealTypeFundamental, error) {

	var r = dm.DealTypeFundamental{}
	var rList []dm.DealTypeFundamental
	

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}