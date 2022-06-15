# **NegotiableInstrument** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**NegotiableInstrument** (negotiableinstrument) |
|Endpoint 	    |**/NegotiableInstrument...** [^1]|
|Endpoint Query |**Id**|
|REST API|**/API/NegotiableInstrument/**|
Glyph|**fas fa-file-signature** (text-info)
Friendly Name|**Negotiable Instrument**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/NegotiableInstrument/NegotiableInstrumentList) [Exportable]
* **View** (/NegotiableInstrument/NegotiableInstrumentView)

* **Save** (/NegotiableInstrument/NegotiableInstrumentSave)









##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **lseGiltsDataStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|
|**Id**|String|false|true|false|false|||||Y|id||false|false|false|
|**LongName**|String|false|true|false|false|||||Y|longName||false|false|false|
|**Isin**|String|false|true|false|false|||||Y|isin||false|false|false|
|**Tidm**|String|false|true|false|false|||||Y|tidm||false|false|false|
|**Sedol**|String|false|true|false|false|||||Y|sedol||false|false|false|
|**IssueDate**|String|false|true|false|false|||||Y|issueDate||false|false|false|
|**MaturityDate**|String|false|true|false|false|||||Y|maturityDate||false|false|false|
|**CouponValue**|String|false|true|false|false|||||Y|couponValue||false|false|false|
|**CouponType**|String|false|true|false|false|||||Y|couponType||false|false|false|
|**Segment**|String|false|true|false|false|||||Y|segment||false|false|false|
|**Sector**|String|false|true|false|false|||||Y|sector||false|false|false|
|**CodeConventionCalculateAccrual**|String|false|true|false|false|||||Y|codeConventionCalculateAccrual||false|false|false|
|**MinimumDenomination**|String|false|true|false|false|||||Y|minimumDenomination||false|false|false|
|**DenominationCurrency**|String|false|true|false|false|||||Y|denominationCurrency||false|false|false|
|**TradingCurrency**|String|false|true|false|false|||||Y|tradingCurrency||false|false|false|
|**Type**|String|false|true|false|false|||||Y|type||false|false|false|
|**FlatYield**|String|false|true|false|false|||||Y|flatYield||false|false|false|
|**PaymentCouponDate**|String|false|true|false|false|||||Y|paymentCouponDate||false|false|false|
|**PeriodOfCoupon**|String|false|true|false|false|||||Y|periodOfCoupon||false|false|false|
|**ExCouponDate**|String|false|true|false|false|||||Y|exCouponDate||false|false|false|
|**DateOfIndexInflation**|String|false|true|false|false|||||Y|dateOfIndexInflation||false|false|false|
|**UnitOfQuotation**|String|false|true|false|false|||||Y|unitOfQuotation||false|false|false|
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|
|**Issuer**|String|false|true|false|false|||||Y|issuer||false|false|false|
|**IssueAmount**|String|false|true|false|false|||||Y|issueAmount||false|false|false|
|**RunningYield**|String|false|true|false|false|||||Y|runningYield||false|false|false|
|**LEI**|String|false|true|false|false|||||Y|LEI||false|false|false|
|**CUSIP**|String|false|true|false|false|||||Y|CUSIP||false|false|false|
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/negotiableInstrument_core.go |
| code | **api** | /application/negotiableInstrument_api.go |
| code | **dao** | /dao/negotiableInstrument_core.go |
| code | **datamodel** | /datamodel/negotiableInstrument_core.go |
| code | **menu** | /design/menu/negotiableInstrument.json |
| html | **list** | /html/NegotiableInstrument_List.html |
| html | **view** | /html/NegotiableInstrument_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:32:05**
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