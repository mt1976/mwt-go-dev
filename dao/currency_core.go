package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/currency.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Currency (currency)
// Endpoint 	        : Currency (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:08
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

// Currency_GetList() returns a list of all Currency records
func Currency_GetList() (int, []dm.Currency, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Currency_SQLTable)
	count, currencyList, _, _ := currency_Fetch(tsql)
	
	return count, currencyList, nil
}


// Currency_GetLookup() returns a lookup list of all Currency items in lookup format
func Currency_GetLookup() []dm.Lookup_Item {
	var returnList []dm.Lookup_Item
	count, currencyList, _ := Currency_GetList()
	for i := 0; i < count; i++ {
		returnList = append(returnList, dm.Lookup_Item{ID: currencyList[i].Code, Name: currencyList[i].Name})
	}
	return returnList
}


// Currency_GetByID() returns a single Currency record
func Currency_GetByID(id string) (int, dm.Currency, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Currency_SQLTable)
	tsql = tsql + " WHERE " + dm.Currency_SQLSearchID + "='" + id + "'"
	_, _, currencyItem, _ := currency_Fetch(tsql)

	return 1, currencyItem, nil
}

// Currency_GetByReverseLookup(id string) returns a single Currency record using the Name field as key.
func Currency_GetByReverseLookup(id string) (int, dm.Currency, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Currency_SQLTable)
	tsql = tsql + " WHERE Name = '" + id + "'"

	_, _, currencyItem, _ := currency_Fetch(tsql)
	
	return 1, currencyItem, nil
}

// Currency_DeleteByID() deletes a single Currency record
func Currency_Delete(id string) {


	adaptor.Currency_Delete_impl(id)
	
}


// Currency_Store() saves/stores a Currency record to the database
func Currency_Store(r dm.Currency,req *http.Request) error {

	err := currency_Save(r,Audit_User(req))

	return err
}

// Currency_StoreSystem() saves/stores a Currency record to the database
func Currency_StoreSystem(r dm.Currency) error {
	
	err := currency_Save(r,Audit_Host())

	return err
}

// currency_Save() saves/stores a Currency record to the database
func currency_Save(r dm.Currency,usr string) error {

    var err error



	

	if len(r.Code) == 0 {
		r.Code = Currency_NewID(r)
	}

// If there are fields below, create the methods in dao\currency_impl.go


























	
logs.Storing("Currency",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Currency_impl.go file
	err1 := adaptor.Currency_Delete_impl(r.Code)
	err2 := adaptor.Currency_Update_impl(r.Code,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// currency_Fetch read all Currency's
func currency_Fetch(tsql string) (int, []dm.Currency, dm.Currency, error) {

	var recItem dm.Currency
	var recList []dm.Currency

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.Code  = get_String(rec, dm.Currency_Code_sql, "")
	   recItem.Name  = get_String(rec, dm.Currency_Name_sql, "")
	   recItem.AmountDp  = get_Int(rec, dm.Currency_AmountDp_sql, "0")
	   recItem.Country  = get_String(rec, dm.Currency_Country_sql, "")
	   recItem.CountryName  = get_String(rec, dm.Currency_CountryName_sql, "")
	   recItem.IntBase  = get_String(rec, dm.Currency_IntBase_sql, "")
	   recItem.KeydateBase  = get_String(rec, dm.Currency_KeydateBase_sql, "")
	   recItem.InterestRateTolerance  = get_Float(rec, dm.Currency_InterestRateTolerance_sql, "0.00")
	   recItem.CheckPayTo  = get_Bool(rec, dm.Currency_CheckPayTo_sql, "True")
	   recItem.LatinAmericanSettlement  = get_Bool(rec, dm.Currency_LatinAmericanSettlement_sql, "True")
	   recItem.DefaultLayOffBookKey  = get_String(rec, dm.Currency_DefaultLayOffBookKey_sql, "")
	   recItem.CutOffTimeCutOffTime  = get_Time(rec, dm.Currency_CutOffTimeCutOffTime_sql, "")
	   recItem.CutOffTimeTimeZone  = get_String(rec, dm.Currency_CutOffTimeTimeZone_sql, "")
	   recItem.CutOffTimeDerivedDataUTCOffset  = get_String(rec, dm.Currency_CutOffTimeDerivedDataUTCOffset_sql, "")
	   recItem.CutOffTimeDerivedDataHasDaylightSaving  = get_Bool(rec, dm.Currency_CutOffTimeDerivedDataHasDaylightSaving_sql, "True")
	   recItem.CutOffTimeDerivedDataDaylightStart  = get_Time(rec, dm.Currency_CutOffTimeDerivedDataDaylightStart_sql, "")
	   recItem.CutOffTimeDerivedDataDaylightEnd  = get_Time(rec, dm.Currency_CutOffTimeDerivedDataDaylightEnd_sql, "")
	   recItem.DealerInterventionQuoteTimeout  = get_Int(rec, dm.Currency_DealerInterventionQuoteTimeout_sql, "0")
	   recItem.CutOffTimeCutOffPeriod  = get_String(rec, dm.Currency_CutOffTimeCutOffPeriod_sql, "")
	   recItem.StripRateFutureExchangeCode  = get_String(rec, dm.Currency_StripRateFutureExchangeCode_sql, "")
	   recItem.StripRateFutureCurrencyContractCurrencyIsoCode  = get_String(rec, dm.Currency_StripRateFutureCurrencyContractCurrencyIsoCode_sql, "")
	   recItem.StripRateFutureCurrencyContractFutureContractCode  = get_String(rec, dm.Currency_StripRateFutureCurrencyContractFutureContractCode_sql, "")
	   recItem.OvernightFundingSpreadBid  = get_Float(rec, dm.Currency_OvernightFundingSpreadBid_sql, "0.00")
	   recItem.OvernightFundingSpreadOffer  = get_Float(rec, dm.Currency_OvernightFundingSpreadOffer_sql, "0.00")
	
	// If there are fields below, create the methods in adaptor\Currency_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func Currency_NewID(r dm.Currency) string {
	
			id := uuid.New().String()
	
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

