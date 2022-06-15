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
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|
|**Id**|String|true|true|false|false|||||Y|id||false|false|false|
|**Object**|String|false|true|false|false|||||Y|object||false|false|false|
|**Field**|String|false|true|false|false|||||Y|field||false|false|false|
|**Value**|String|false|true|false|false|||||Y|value||false|false|false|
|**Expiry**|String|false|true|false|false|||||Y|expiry||false|false|false|
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|
|**Source**|String|false|true|false|false|||||Y|source||false|false|false|
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/cache_core.go |
| code | **dao** | /dao/cache_core.go |
| code | **datamodel** | /datamodel/cache_core.go |
| code | **menu** | /design/menu/cache.json |
| html | **list** | /html/Cache_List.html |
| html | **view** | /html/Cache_View.html |
| html | **edit** | /html/Cache_Edit.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:31:50**
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