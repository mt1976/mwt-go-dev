# **Cache** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Cache** (cache) |
|Endpoint 	    |**/Cache...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-database** (text-info)
Friendly Name|**Cache Content**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Cache/CacheList) [Exportable]
* **View** (/Cache/CacheView)
* **Edit** (/Cache/CacheEdit)
* **Save** (/Cache/CacheSave)









##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **cacheStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**Id**|String|true|true|false|false|||||Y|id||false|false|false|text||
|**Object**|String|false|true|false|false|||||Y|object||false|false|false|text||
|**Field**|String|false|true|false|false|||||Y|field||false|false|false|text||
|**Value**|String|false|true|false|false|||||Y|value||false|false|false|text||
|**Expiry**|String|false|true|false|false|||||Y|expiry||false|false|false|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|text||
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**Source**|String|false|true|false|false|||||Y|source||false|false|false|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/cache_core.go |
| code | **dao** | /dao/cache_core.go |
| code | **datamodel** | /datamodel/cache_core.go |
| code | **menu** | /design/menu/cache.json |
| html | **list** | /Cache_List.html |
| html | **view** | /Cache_View.html |
| html | **edit** | /Cache_Edit.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **28/06/2022** at **16:10:44**
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