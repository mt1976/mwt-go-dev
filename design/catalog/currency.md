# **Currency** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Currency** (currency) |
|Endpoint 	    |**/Currency...** [^1]|
|Endpoint Query |**Code**|
|REST API|**/API/Currency/**|
Glyph|**fas fa-dollar-sign** ()
Friendly Name|**Currency**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Currency/CurrencyList) [Exportable]
* **View** (/Currency/CurrencyView)











##  Provides
 * Lookup (Code Name)
 * Reverse Lookup (Name)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCurrency**
SQL Table Key | **Code**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**Code**|String|true|true|false|false|||||Y|Code||false|false|false|
|**Name**|String|false|true|false|false|||||Y|Name||false|false|false|
|**AmountDp**|Int|false|true|false|false|||||Y|AmountDp|0|false|false|false|
|**Country**|String|false|true|false|false|||||Y|Country||false|false|false|
|**CountryName**|String|false|true|false|false|||||Y|CountryName||false|false|false|
|**IntBase**|String|false|true|false|false|||||Y|IntBase||false|false|false|
|**KeydateBase**|String|false|true|false|false|||||Y|KeydateBase||false|false|false|
|**InterestRateTolerance**|Float|false|true|false|false|||||Y|InterestRateTolerance|0.00|false|false|false|
|**CheckPayTo**|Bool|false|true|false|false|||||Y|CheckPayTo|True|false|false|false|
|**LatinAmericanSettlement**|Bool|false|true|false|false|||||Y|LatinAmericanSettlement|True|false|false|false|
|**DefaultLayOffBookKey**|String|false|true|false|false|||||Y|DefaultLayOffBookKey||false|false|false|
|**CutOffTimeCutOffTime**|Time|false|true|false|false|||||Y|CutOffTimeCutOffTime||false|false|false|
|**CutOffTimeTimeZone**|String|false|true|false|false|||||Y|CutOffTimeTimeZone||false|false|false|
|**CutOffTimeDerivedDataUTCOffset**|String|false|true|false|false|||||Y|CutOffTimeDerivedDataUTCOffset||false|false|false|
|**CutOffTimeDerivedDataHasDaylightSaving**|Bool|false|true|false|false|||||Y|CutOffTimeDerivedDataHasDaylightSaving|True|false|false|false|
|**CutOffTimeDerivedDataDaylightStart**|Time|false|true|false|false|||||Y|CutOffTimeDerivedDataDaylightStart||false|false|false|
|**CutOffTimeDerivedDataDaylightEnd**|Time|false|true|false|false|||||Y|CutOffTimeDerivedDataDaylightEnd||false|false|false|
|**DealerInterventionQuoteTimeout**|Int|false|true|false|false|||||Y|DealerInterventionQuoteTimeout|0|false|false|false|
|**CutOffTimeCutOffPeriod**|String|false|true|false|false|||||Y|CutOffTimeCutOffPeriod||false|false|false|
|**StripRateFutureExchangeCode**|String|false|true|false|false|||||Y|StripRateFutureExchangeCode||false|false|false|
|**StripRateFutureCurrencyContractCurrencyIsoCode**|String|false|true|false|false|||||Y|StripRateFutureCurrencyContractCurrencyIsoCode||false|false|false|
|**StripRateFutureCurrencyContractFutureContractCode**|String|false|true|false|false|||||Y|StripRateFutureCurrencyContractFutureContractCode||false|false|false|
|**OvernightFundingSpreadBid**|Float|false|true|false|false|||||Y|OvernightFundingSpreadBid|0.00|false|false|false|
|**OvernightFundingSpreadOffer**|Float|false|true|false|false|||||Y|OvernightFundingSpreadOffer|0.00|false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/currency_core.go |
| code | **adaptor** | /adaptor/currency_impl.go_template |
| code | **api** | /application/currency_api.go |
| code | **dao** | /dao/currency_core.go |
| code | **datamodel** | /datamodel/currency_core.go |
| code | **menu** | /design/menu/currency.json |
| html | **list** | /html/Currency_List.html |
| html | **view** | /html/Currency_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:32:00**
Who & Where		     | **matttownsend (Matt Townsend)** on **silicon.local**

### Footnotes
[^1]: **Endpoint**
    * The full list of endpoints can be found in the [Actions](#action-id) section
[^2]: **Lookup Fields**
    * LL = A List Lookup. Define list in lits.cfg
    * OL = An Object Lookup. Get a list of values from an Object
    * FL = Fetches 1 value from an object based on the content of the field. 
[^3]: **Inputtable**   
    * H = Hidden Field
    * N = No Input Field
    * Y = Inputtable Field