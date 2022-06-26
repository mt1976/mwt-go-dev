# **DataLoaderMap** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**DataLoaderMap** (dataloadermap) |
|Endpoint 	    |**/DataLoaderMap...** [^1]|
|Endpoint Query |**Id**|
Glyph|**fas fa-city** ()
Friendly Name|**Data Loader Data Map**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/DataLoaderMap/DataLoaderMapList) [Exportable]
* **View** (/DataLoaderMap/DataLoaderMapView)
* **Edit** (/DataLoaderMap/DataLoaderMapEdit)
* **Save** (/DataLoaderMap/DataLoaderMapSave)









##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **loaderMapStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**Id**|String|true|true|false|false|||||Y|id||false|false|false|text||
|**Name**|String|false|true|false|false|||||Y|name||false|false|false|text||
|**Position**|String|false|true|false|false|||||Y|position||false|false|false|text||
|**Loader**|String|false|true|false|false|OL|DataLoader|Loader|Name|N|loader||false|false|false|||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|text||
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**Int_position**|Int|false|true|false|false|||||Y|int_position|0|false|false|false|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/dataLoaderMap_core.go |
| code | **dao** | /dao/dataLoaderMap_core.go |
| code | **datamodel** | /datamodel/dataLoaderMap_core.go |
| code | **menu** | /design/menu/dataLoaderMap.json |
| html | **list** | /DataLoaderMap_List.html |
| html | **view** | /DataLoaderMap_View.html |
| html | **edit** | /DataLoaderMap_Edit.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **26/06/2022** at **18:48:26**
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