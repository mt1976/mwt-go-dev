# **Systems** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Systems** (systems) |
|Endpoint 	    |**/Systems...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-plug** (text-info)
Friendly Name|**Connected System**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Systems/SystemsList) [Exportable]
* **View** (/Systems/SystemsView)
* **Edit** (/Systems/SystemsEdit)
* **Save** (/Systems/SystemsSave)
* **New** (/Systems/SystemsNew)
* **Delete** (/Systems/SystemsDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **systemStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|
|**Id**|String|true|true|false|false|||||Y|Id||false|false|false|
|**Name**|String|false|true|false|false|||||Y|Name||false|false|false|
|**Staticin**|String|false|true|false|false|||||Y|Staticin||false|false|false|
|**Staticout**|String|false|true|false|false|||||Y|Staticout||false|false|false|
|**Txnin**|String|false|true|false|false|||||Y|Txnin||false|false|false|
|**Txnout**|String|false|true|false|false|||||Y|Txnout||false|false|false|
|**Fundscheckin**|String|false|true|false|false|||||Y|Fundscheckin||false|false|false|
|**Fundscheckout**|String|false|true|false|false|||||Y|Fundscheckout||false|false|false|
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|
|**SWIFTin**|String|false|true|false|false|||||Y|SWIFTin||false|false|false|
|**SWIFTout**|String|false|true|false|false|||||Y|SWIFTout||false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/systems_core.go |
| code | **dao** | /dao/systems_core.go |
| code | **datamodel** | /datamodel/systems_core.go |
| code | **menu** | /design/menu/systems.json |
| html | **list** | /html/Systems_List.html |
| html | **view** | /html/Systems_View.html |
| html | **edit** | /html/Systems_Edit.html |
| html | **new** | /html/Systems_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:32:10**
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