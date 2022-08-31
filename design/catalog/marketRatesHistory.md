# **MarketRatesHistory** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**MarketRatesHistory** (marketrateshistory) |
|Endpoint 	    |**/MarketRatesHistory...** [^1]|
|Endpoint Query |**ID**|
|REST API|**/API/MarketRatesHistory/**|
Glyph|**fas fa-file-invoice-dollar** (text-info)
Friendly Name|**Rates Store History**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/MarketRatesHistory/MarketRatesHistoryList) [Exportable]
* **View** (/MarketRatesHistory/MarketRatesHistoryView)











##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **rateDataStoreHistory**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**Id**|String|true|true|false|false|||||Y|id||false|false|false|text||
|**Bid**|String|false|true|false|false|||||Y|bid||false|false|false|text||
|**Mid**|String|false|true|false|false|||||Y|mid||false|false|false|text||
|**Offer**|String|false|true|false|false|||||Y|offer||false|false|false|text||
|**Market**|String|false|true|false|false|||||Y|market||false|false|false|text||
|**Tenor**|String|false|true|false|false|||||Y|tenor||false|false|false|text||
|**Series**|String|false|true|false|false|||||Y|series||false|false|false|text||
|**Name**|String|false|true|false|false|||||Y|name||false|false|false|text||
|**Source**|String|false|true|false|false|||||Y|source||false|false|false|text||
|**Destination**|String|false|true|false|false|||||Y|destination||false|false|false|text||
|**Class**|String|false|true|false|false|||||Y|class||false|false|false|text||
|**Date**|String|false|true|false|false|||||Y|date||false|false|false|text||
|**Time**|String|false|true|false|false|||||Y|time||false|false|false|text||
|**Datetime**|String|false|true|false|false|||||Y|datetime||false|false|false|text||
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|text||
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/marketRatesHistory_core.go |
| code | **api** | /application/marketRatesHistory_api.go |
| code | **dao** | /dao/marketRatesHistory_core.go |
| code | **datamodel** | /datamodel/marketRatesHistory_core.go |
| code | **menu** | /design/menu/marketRatesHistory.json |
| html | **list** | /MarketRatesHistory_List.html |
| html | **view** | /MarketRatesHistory_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **Dysprosium [r4-21.12.31]**
Date & Time		     | **31/08/2022** at **14:19:48**
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