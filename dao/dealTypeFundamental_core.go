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
// Date & Time		    : 14/06/2022 at 21:32:04
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

	return 1, dealtypefundamentalItem, nil
}



// DealTypeFundamental_DeleteByID() deletes a single DealTypeFundamental record
func DealTypeFundamental_Delete(id string) {


	adaptor.DealTypeFundamental_Delete_impl(id)
	
}


// DealTypeFundamental_Store() saves/stores a DealTypeFundamental record to the database
func DealTypeFundamental_Store(r dm.DealTypeFundamental,req *http.Request) error {

	err := dealtypefundamental_Save(r,Audit_User(req))

	return err
}

// DealTypeFundamental_StoreSystem() saves/stores a DealTypeFundamental record to the database
func DealTypeFundamental_StoreSystem(r dm.DealTypeFundamental) error {
	
	err := dealtypefundamental_Save(r,Audit_Host())

	return err
}

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
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - START
   recItem.DealTypeKey  = get_String(rec, dm.DealTypeFundamental_DealTypeKey, "")
   recItem.Amendment  = get_Bool(rec, dm.DealTypeFundamental_Amendment, "True")
   recItem.DefaultFrequency  = get_Int(rec, dm.DealTypeFundamental_DefaultFrequency, "0")
   recItem.Directions  = get_String(rec, dm.DealTypeFundamental_Directions, "")
   recItem.SettledTermDealType  = get_String(rec, dm.DealTypeFundamental_SettledTermDealType, "")
   recItem.Optn  = get_Bool(rec, dm.DealTypeFundamental_Optn, "True")
   recItem.AllowPledge  = get_Bool(rec, dm.DealTypeFundamental_AllowPledge, "True")
   recItem.Takeup  = get_Bool(rec, dm.DealTypeFundamental_Takeup, "True")
   recItem.MismatchDealType  = get_String(rec, dm.DealTypeFundamental_MismatchDealType, "")
   recItem.SettledHedgeTermDealType  = get_String(rec, dm.DealTypeFundamental_SettledHedgeTermDealType, "")
   recItem.SettlementCode  = get_String(rec, dm.DealTypeFundamental_SettlementCode, "")
   recItem.TermSubType  = get_String(rec, dm.DealTypeFundamental_TermSubType, "")
   recItem.FundingDealType  = get_String(rec, dm.DealTypeFundamental_FundingDealType, "")
   recItem.TransferType  = get_String(rec, dm.DealTypeFundamental_TransferType, "")
   recItem.TermDealType  = get_String(rec, dm.DealTypeFundamental_TermDealType, "")
   recItem.NegotiableInstrumentType  = get_String(rec, dm.DealTypeFundamental_NegotiableInstrumentType, "")
   recItem.Mismatch  = get_Bool(rec, dm.DealTypeFundamental_Mismatch, "True")
   recItem.ComplexTransferSubType  = get_String(rec, dm.DealTypeFundamental_ComplexTransferSubType, "")
   recItem.LayOffDealType  = get_String(rec, dm.DealTypeFundamental_LayOffDealType, "")
   recItem.NIAccount  = get_Int(rec, dm.DealTypeFundamental_NIAccount, "0")
   recItem.SimpleMMsubtype  = get_Int(rec, dm.DealTypeFundamental_SimpleMMsubtype, "0")
   recItem.SwapDealType  = get_String(rec, dm.DealTypeFundamental_SwapDealType, "")
   recItem.Positions  = get_String(rec, dm.DealTypeFundamental_Positions, "")
   recItem.OptionOutright  = get_String(rec, dm.DealTypeFundamental_OptionOutright, "")
   recItem.SettledHedgeSpotDealType  = get_String(rec, dm.DealTypeFundamental_SettledHedgeSpotDealType, "")
   recItem.StraightThroughInterestMethod  = get_Bool(rec, dm.DealTypeFundamental_StraightThroughInterestMethod, "True")
   recItem.SubType  = get_String(rec, dm.DealTypeFundamental_SubType, "")
   recItem.Rollover  = get_Bool(rec, dm.DealTypeFundamental_Rollover, "True")
   recItem.DefaultIssuer  = get_String(rec, dm.DealTypeFundamental_DefaultIssuer, "")
   recItem.DefaultStartDate  = get_Int(rec, dm.DealTypeFundamental_DefaultStartDate, "0")
   recItem.Fee  = get_String(rec, dm.DealTypeFundamental_Fee, "")
   recItem.NDF  = get_Bool(rec, dm.DealTypeFundamental_NDF, "True")
   recItem.FXFX  = get_Bool(rec, dm.DealTypeFundamental_FXFX, "True")
   recItem.ONIA  = get_Bool(rec, dm.DealTypeFundamental_ONIA, "True")
   recItem.MarginSubType  = get_Int(rec, dm.DealTypeFundamental_MarginSubType, "0")
   recItem.TransferDealType  = get_String(rec, dm.DealTypeFundamental_TransferDealType, "")
   recItem.IsFX  = get_Bool(rec, dm.DealTypeFundamental_IsFX, "True")
   recItem.Ordr  = get_String(rec, dm.DealTypeFundamental_Ordr, "")
   recItem.OptionStyle  = get_String(rec, dm.DealTypeFundamental_OptionStyle, "")
   recItem.SpotDealType  = get_String(rec, dm.DealTypeFundamental_SpotDealType, "")
   recItem.CanIssue  = get_Bool(rec, dm.DealTypeFundamental_CanIssue, "True")
   recItem.CanShort  = get_Bool(rec, dm.DealTypeFundamental_CanShort, "True")
   recItem.FXMarginTradingType  = get_Int(rec, dm.DealTypeFundamental_FXMarginTradingType, "0")
   recItem.Internal  = get_Bool(rec, dm.DealTypeFundamental_Internal, "True")
   recItem.TicketBasename  = get_String(rec, dm.DealTypeFundamental_TicketBasename, "")
   recItem.InterestRateFutureType  = get_String(rec, dm.DealTypeFundamental_InterestRateFutureType, "")
   recItem.TradingLimitProductCode  = get_String(rec, dm.DealTypeFundamental_TradingLimitProductCode, "")
   recItem.Forward  = get_Bool(rec, dm.DealTypeFundamental_Forward, "True")
   recItem.MaturityNotificationPeriod  = get_String(rec, dm.DealTypeFundamental_MaturityNotificationPeriod, "")
   recItem.NotificationEvents  = get_String(rec, dm.DealTypeFundamental_NotificationEvents, "")
   recItem.SwapSubType  = get_String(rec, dm.DealTypeFundamental_SwapSubType, "")
   recItem.ProductCode  = get_String(rec, dm.DealTypeFundamental_ProductCode, "")
   recItem.Funding  = get_Bool(rec, dm.DealTypeFundamental_Funding, "True")
   recItem.AllocationPricing  = get_String(rec, dm.DealTypeFundamental_AllocationPricing, "")
   recItem.CancelPeriod  = get_String(rec, dm.DealTypeFundamental_CancelPeriod, "")
   recItem.MMMarginTradingType  = get_Int(rec, dm.DealTypeFundamental_MMMarginTradingType, "0")
   recItem.OptionSpot  = get_String(rec, dm.DealTypeFundamental_OptionSpot, "")
   recItem.Transfer  = get_Bool(rec, dm.DealTypeFundamental_Transfer, "True")
   recItem.NotificationPeriod  = get_String(rec, dm.DealTypeFundamental_NotificationPeriod, "")
   recItem.Paymentdateshift  = get_Int(rec, dm.DealTypeFundamental_Paymentdateshift, "0")
   recItem.CloseOut  = get_Bool(rec, dm.DealTypeFundamental_CloseOut, "True")
   recItem.FXOptionPricing  = get_String(rec, dm.DealTypeFundamental_FXOptionPricing, "")
   recItem.SettledHedgeOutrightDealType  = get_String(rec, dm.DealTypeFundamental_SettledHedgeOutrightDealType, "")
   recItem.RepoBond  = get_String(rec, dm.DealTypeFundamental_RepoBond, "")
   recItem.RepoTerm  = get_String(rec, dm.DealTypeFundamental_RepoTerm, "")
   recItem.RepoType  = get_Int(rec, dm.DealTypeFundamental_RepoType, "0")
   recItem.DateRule  = get_String(rec, dm.DealTypeFundamental_DateRule, "")
   recItem.CorpTransferDealType  = get_String(rec, dm.DealTypeFundamental_CorpTransferDealType, "")
   recItem.GenerateFXImage  = get_Bool(rec, dm.DealTypeFundamental_GenerateFXImage, "True")
   recItem.Variant  = get_String(rec, dm.DealTypeFundamental_Variant, "")
   recItem.HedgeTermDealType  = get_String(rec, dm.DealTypeFundamental_HedgeTermDealType, "")
   recItem.PricingModel  = get_String(rec, dm.DealTypeFundamental_PricingModel, "")
   recItem.HedgeOutrightDealType  = get_String(rec, dm.DealTypeFundamental_HedgeOutrightDealType, "")
   recItem.Fixing  = get_Bool(rec, dm.DealTypeFundamental_Fixing, "True")
   recItem.Payment  = get_Bool(rec, dm.DealTypeFundamental_Payment, "True")
   recItem.MT  = get_Bool(rec, dm.DealTypeFundamental_MT, "True")
   recItem.SettlementInstructionStyle  = get_String(rec, dm.DealTypeFundamental_SettlementInstructionStyle, "")
   recItem.QuoteHistoryRequired  = get_Bool(rec, dm.DealTypeFundamental_QuoteHistoryRequired, "True")
   recItem.Brokerage  = get_Bool(rec, dm.DealTypeFundamental_Brokerage, "True")
   recItem.ExposureDisabled  = get_Bool(rec, dm.DealTypeFundamental_ExposureDisabled, "True")
   recItem.CreditLine  = get_String(rec, dm.DealTypeFundamental_CreditLine, "")
   recItem.Encumbered  = get_Bool(rec, dm.DealTypeFundamental_Encumbered, "True")
   recItem.InternalId  = get_Int(rec, dm.DealTypeFundamental_InternalId, "0")
   recItem.InternalDeleted  = get_Time(rec, dm.DealTypeFundamental_InternalDeleted, "")
   recItem.UpdatedTransactionId  = get_String(rec, dm.DealTypeFundamental_UpdatedTransactionId, "")
   recItem.UpdatedUserId  = get_String(rec, dm.DealTypeFundamental_UpdatedUserId, "")
   recItem.UpdatedDateTime  = get_Time(rec, dm.DealTypeFundamental_UpdatedDateTime, "")
   recItem.DeletedTransactionId  = get_String(rec, dm.DealTypeFundamental_DeletedTransactionId, "")
   recItem.DeletedUserId  = get_String(rec, dm.DealTypeFundamental_DeletedUserId, "")
   recItem.ChangeType  = get_String(rec, dm.DealTypeFundamental_ChangeType, "")
// If there are fields below, create the methods in adaptor\DealTypeFundamental_impl.go



























































































	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func DealTypeFundamental_NewID(r dm.DealTypeFundamental) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

