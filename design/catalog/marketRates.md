# **MarketRates** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**MarketRates** (marketrates) |
|Endpoint 	    |**/MarketRates...** [^1]|
|Endpoint Query |**ID**|
|REST API|**/API/MarketRates/**|
Glyph|**fas fa-file-invoice-dollar** (text-info)
Friendly Name|**Rates Store Content**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/MarketRates/MarketRatesList) [Exportable]
* **View** (/MarketRates/MarketRatesView)











##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **rateDataStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|
|**Id**|String|true|true|false|false|||||Y|id||false|false|false|
|**Bid**|String|false|true|false|false|||||Y|bid||false|false|false|
|**Mid**|String|false|true|false|false|||||Y|mid||false|false|false|
|**Offer**|String|false|true|false|false|||||Y|offer||false|false|false|
|**Market**|String|false|true|false|false|||||Y|market||false|false|false|
|**Tenor**|String|false|true|false|false|||||Y|tenor||false|false|false|
|**Series**|String|false|true|false|false|||||Y|series||false|false|false|
|**Name**|String|false|true|false|false|||||Y|name||false|false|false|
|**Source**|String|false|true|false|false|||||Y|source||false|false|false|
|**Destination**|String|false|true|false|false|||||Y|destination||false|false|false|
|**Class**|String|false|true|false|false|||||Y|class||false|false|false|
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|
|**Date**|String|false|true|false|false|||||Y|date||false|false|false|
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/marketRates_core.go |
| code | **api** | /application/marketRates_api.go |
| code | **dao** | /dao/marketRates_core.go |
| code | **datamodel** | /datamodel/marketRates_core.go |
| code | **menu** | /design/menu/marketRates.json |
| html | **list** | /html/MarketRates_List.html |
| html | **view** | /html/MarketRates_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:32:06**
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