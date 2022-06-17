package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/currency.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Currency (currency)
// Endpoint 	        : Currency (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Currency defines the datamolde for the Currency object
type Currency struct {

Code       string
Name       string
AmountDp       string
Country       string
CountryName       string
IntBase       string
KeydateBase       string
InterestRateTolerance       string
CheckPayTo       string
LatinAmericanSettlement       string
DefaultLayOffBookKey       string
CutOffTimeCutOffTime       string
CutOffTimeTimeZone       string
CutOffTimeDerivedDataUTCOffset       string
CutOffTimeDerivedDataHasDaylightSaving       string
CutOffTimeDerivedDataDaylightStart       string
CutOffTimeDerivedDataDaylightEnd       string
DealerInterventionQuoteTimeout       string
CutOffTimeCutOffPeriod       string
StripRateFutureExchangeCode       string
StripRateFutureCurrencyContractCurrencyIsoCode       string
StripRateFutureCurrencyContractFutureContractCode       string
OvernightFundingSpreadBid       string
OvernightFundingSpreadOffer       string

}

const (
	Currency_Title       = "Currency"
	Currency_SQLTable    = "sienaCurrency"
	Currency_SQLSearchID = "Code"
	Currency_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Currency_Template     = "Currency"
	Currency_TemplateList = "/Currency/Currency_List"
	Currency_TemplateView = "/Currency/Currency_View"
	Currency_TemplateEdit = "/Currency/Currency_Edit"
	Currency_TemplateNew  = "/Currency/Currency_New"
	///
	/// Handler Monitor Paths
	///
	Currency_Path       = "/API/Currency/"
	Currency_PathList   = "/CurrencyList/"
	Currency_PathView   = "/CurrencyView/"
	Currency_PathEdit   = "/CurrencyEdit/"
	Currency_PathNew    = "/CurrencyNew/"
	Currency_PathSave   = "/CurrencySave/"
	Currency_PathDelete = "/CurrencyDelete/"
	///
	///
	/// SQL Field Definitions
	///
Currency_Code_sql   = "Code" // Code is a String
Currency_Name_sql   = "Name" // Name is a String
Currency_AmountDp_sql   = "AmountDp" // AmountDp is a Int
Currency_Country_sql   = "Country" // Country is a String
Currency_CountryName_sql   = "CountryName" // CountryName is a String
Currency_IntBase_sql   = "IntBase" // IntBase is a String
Currency_KeydateBase_sql   = "KeydateBase" // KeydateBase is a String
Currency_InterestRateTolerance_sql   = "InterestRateTolerance" // InterestRateTolerance is a Float
Currency_CheckPayTo_sql   = "CheckPayTo" // CheckPayTo is a Bool
Currency_LatinAmericanSettlement_sql   = "LatinAmericanSettlement" // LatinAmericanSettlement is a Bool
Currency_DefaultLayOffBookKey_sql   = "DefaultLayOffBookKey" // DefaultLayOffBookKey is a String
Currency_CutOffTimeCutOffTime_sql   = "CutOffTimeCutOffTime" // CutOffTimeCutOffTime is a Time
Currency_CutOffTimeTimeZone_sql   = "CutOffTimeTimeZone" // CutOffTimeTimeZone is a String
Currency_CutOffTimeDerivedDataUTCOffset_sql   = "CutOffTimeDerivedDataUTCOffset" // CutOffTimeDerivedDataUTCOffset is a String
Currency_CutOffTimeDerivedDataHasDaylightSaving_sql   = "CutOffTimeDerivedDataHasDaylightSaving" // CutOffTimeDerivedDataHasDaylightSaving is a Bool
Currency_CutOffTimeDerivedDataDaylightStart_sql   = "CutOffTimeDerivedDataDaylightStart" // CutOffTimeDerivedDataDaylightStart is a Time
Currency_CutOffTimeDerivedDataDaylightEnd_sql   = "CutOffTimeDerivedDataDaylightEnd" // CutOffTimeDerivedDataDaylightEnd is a Time
Currency_DealerInterventionQuoteTimeout_sql   = "DealerInterventionQuoteTimeout" // DealerInterventionQuoteTimeout is a Int
Currency_CutOffTimeCutOffPeriod_sql   = "CutOffTimeCutOffPeriod" // CutOffTimeCutOffPeriod is a String
Currency_StripRateFutureExchangeCode_sql   = "StripRateFutureExchangeCode" // StripRateFutureExchangeCode is a String
Currency_StripRateFutureCurrencyContractCurrencyIsoCode_sql   = "StripRateFutureCurrencyContractCurrencyIsoCode" // StripRateFutureCurrencyContractCurrencyIsoCode is a String
Currency_StripRateFutureCurrencyContractFutureContractCode_sql   = "StripRateFutureCurrencyContractFutureContractCode" // StripRateFutureCurrencyContractFutureContractCode is a String
Currency_OvernightFundingSpreadBid_sql   = "OvernightFundingSpreadBid" // OvernightFundingSpreadBid is a Float
Currency_OvernightFundingSpreadOffer_sql   = "OvernightFundingSpreadOffer" // OvernightFundingSpreadOffer is a Float

	/// Definitions End

	/// Application Field Definitions
	///
Currency_Code_scrn   = "Code" // Code is a String
Currency_Name_scrn   = "Name" // Name is a String
Currency_AmountDp_scrn   = "AmountDp" // AmountDp is a Int
Currency_Country_scrn   = "Country" // Country is a String
Currency_CountryName_scrn   = "CountryName" // CountryName is a String
Currency_IntBase_scrn   = "IntBase" // IntBase is a String
Currency_KeydateBase_scrn   = "KeydateBase" // KeydateBase is a String
Currency_InterestRateTolerance_scrn   = "InterestRateTolerance" // InterestRateTolerance is a Float
Currency_CheckPayTo_scrn   = "CheckPayTo" // CheckPayTo is a Bool
Currency_LatinAmericanSettlement_scrn   = "LatinAmericanSettlement" // LatinAmericanSettlement is a Bool
Currency_DefaultLayOffBookKey_scrn   = "DefaultLayOffBookKey" // DefaultLayOffBookKey is a String
Currency_CutOffTimeCutOffTime_scrn   = "CutOffTimeCutOffTime" // CutOffTimeCutOffTime is a Time
Currency_CutOffTimeTimeZone_scrn   = "CutOffTimeTimeZone" // CutOffTimeTimeZone is a String
Currency_CutOffTimeDerivedDataUTCOffset_scrn   = "CutOffTimeDerivedDataUTCOffset" // CutOffTimeDerivedDataUTCOffset is a String
Currency_CutOffTimeDerivedDataHasDaylightSaving_scrn   = "CutOffTimeDerivedDataHasDaylightSaving" // CutOffTimeDerivedDataHasDaylightSaving is a Bool
Currency_CutOffTimeDerivedDataDaylightStart_scrn   = "CutOffTimeDerivedDataDaylightStart" // CutOffTimeDerivedDataDaylightStart is a Time
Currency_CutOffTimeDerivedDataDaylightEnd_scrn   = "CutOffTimeDerivedDataDaylightEnd" // CutOffTimeDerivedDataDaylightEnd is a Time
Currency_DealerInterventionQuoteTimeout_scrn   = "DealerInterventionQuoteTimeout" // DealerInterventionQuoteTimeout is a Int
Currency_CutOffTimeCutOffPeriod_scrn   = "CutOffTimeCutOffPeriod" // CutOffTimeCutOffPeriod is a String
Currency_StripRateFutureExchangeCode_scrn   = "StripRateFutureExchangeCode" // StripRateFutureExchangeCode is a String
Currency_StripRateFutureCurrencyContractCurrencyIsoCode_scrn   = "StripRateFutureCurrencyContractCurrencyIsoCode" // StripRateFutureCurrencyContractCurrencyIsoCode is a String
Currency_StripRateFutureCurrencyContractFutureContractCode_scrn   = "StripRateFutureCurrencyContractFutureContractCode" // StripRateFutureCurrencyContractFutureContractCode is a String
Currency_OvernightFundingSpreadBid_scrn   = "OvernightFundingSpreadBid" // OvernightFundingSpreadBid is a Float
Currency_OvernightFundingSpreadOffer_scrn   = "OvernightFundingSpreadOffer" // OvernightFundingSpreadOffer is a Float

	/// Definitions End
)
