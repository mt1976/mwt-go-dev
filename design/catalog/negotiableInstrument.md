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
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**Id**|String|false|true|false|false|||||Y|id||false|false|false|text||
|**LongName**|String|false|true|false|false|||||Y|longName||false|false|false|text||
|**Isin**|String|false|true|false|false|||||Y|isin||false|false|false|text||
|**Tidm**|String|false|true|false|false|||||Y|tidm||false|false|false|text||
|**Sedol**|String|false|true|false|false|||||Y|sedol||false|false|false|text||
|**IssueDate**|String|false|true|false|false|||||Y|issueDate||false|false|false|text||
|**MaturityDate**|String|false|true|false|false|||||Y|maturityDate||false|false|false|text||
|**CouponValue**|String|false|true|false|false|||||Y|couponValue||false|false|false|text||
|**CouponType**|String|false|true|false|false|||||Y|couponType||false|false|false|text||
|**Segment**|String|false|true|false|false|||||Y|segment||false|false|false|text||
|**Sector**|String|false|true|false|false|||||Y|sector||false|false|false|text||
|**CodeConventionCalculateAccrual**|String|false|true|false|false|||||Y|codeConventionCalculateAccrual||false|false|false|text||
|**MinimumDenomination**|String|false|true|false|false|||||Y|minimumDenomination||false|false|false|text||
|**DenominationCurrency**|String|false|true|false|false|||||Y|denominationCurrency||false|false|false|text||
|**TradingCurrency**|String|false|true|false|false|||||Y|tradingCurrency||false|false|false|text||
|**Type**|String|false|true|false|false|||||Y|type||false|false|false|text||
|**FlatYield**|String|false|true|false|false|||||Y|flatYield||false|false|false|text||
|**PaymentCouponDate**|String|false|true|false|false|||||Y|paymentCouponDate||false|false|false|text||
|**PeriodOfCoupon**|String|false|true|false|false|||||Y|periodOfCoupon||false|false|false|text||
|**ExCouponDate**|String|false|true|false|false|||||Y|exCouponDate||false|false|false|text||
|**DateOfIndexInflation**|String|false|true|false|false|||||Y|dateOfIndexInflation||false|false|false|text||
|**UnitOfQuotation**|String|false|true|false|false|||||Y|unitOfQuotation||false|false|false|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|text||
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**Issuer**|String|false|true|false|false|||||Y|issuer||false|false|false|text||
|**IssueAmount**|String|false|true|false|false|||||Y|issueAmount||false|false|false|text||
|**RunningYield**|String|false|true|false|false|||||Y|runningYield||false|false|false|text||
|**LEI**|String|false|true|false|false|||||Y|LEI||false|false|false|text||
|**CUSIP**|String|false|true|false|false|||||Y|CUSIP||false|false|false|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/negotiableInstrument_core.go |
| code | **api** | /application/negotiableInstrument_api.go |
| code | **dao** | /dao/negotiableInstrument_core.go |
| code | **datamodel** | /datamodel/negotiableInstrument_core.go |
| code | **menu** | /design/menu/negotiableInstrument.json |
| html | **list** | /NegotiableInstrument_List.html |
| html | **view** | /NegotiableInstrument_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **17/06/2022** at **18:38:12**
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