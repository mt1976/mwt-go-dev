# **DataLoader** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**DataLoader** (dataloader) |
|Endpoint 	    |**/DataLoader...** [^1]|
|Endpoint Query |**Id**|
Glyph|**fas fa-city** ()
Friendly Name|**Data Loader**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/DataLoader/DataLoaderList) 
* **View** (/DataLoader/DataLoaderView)
* **Edit** (/DataLoader/DataLoaderEdit)
* **Save** (/DataLoader/DataLoaderSave)









##  Provides
 * Lookup (Id Name)






##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **loaderStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**Id**|String|true|true|false|false|||||Y|id||false|false|false|text||
|**Name**|String|false|true|false|false|||||Y|name||false|false|false|text||
|**Description**|String|false|true|false|false|||||Y|description||false|false|false|text||
|**Filename**|String|false|true|false|false|||||Y|filename||false|false|false|text||
|**Lastrun**|String|false|true|false|false|||||Y|lastrun||false|false|false|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|text||
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**Type**|String|false|true|false|false|||||Y|type||false|false|false|text||
|**Instance**|String|false|true|false|false|||||Y|instance||false|false|false|text||
|**Extension**|String|false|true|false|false|||||Y|extension||false|false|false|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/dataLoader_core.go |
| code | **dao** | /dao/dataLoader_core.go |
| code | **datamodel** | /datamodel/dataLoader_core.go |
| code | **menu** | /design/menu/dataLoader.json |
| html | **list** | /DataLoader_List.html |
| html | **view** | /DataLoader_View.html |
| html | **edit** | /DataLoader_Edit.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **28/06/2022** at **16:10:49**
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