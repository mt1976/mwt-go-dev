# **ContactStreamType** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**ContactStreamType** (contactstreamtype) |
|Endpoint 	    |**/ContactStreamType...** [^1]|
|Endpoint Query |**ID**|
Glyph|**far fa-comment** ("")
Friendly Name|**Conversations Types**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/ContactStreamType/ContactStreamTypeList) [Exportable]
* **View** (/ContactStreamType/ContactStreamTypeView)











##  Provides
 * Lookup (TypeId Description)






##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **contactManagerStreamType**
SQL Table Key | **typeId**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**TypeId**|Int|true|true|false|false|||||Y|typeId|0|false|false|false|text||
|**Description**|String|false|true|false|false|||||Y|description||false|false|false|text||
|**RecordState**|Int|false|true|false|false|LL|cmrecordstate|||Y|recordState|0|false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/contactStreamType_core.go |
| code | **adaptor** | /adaptor/contactStreamType_impl.go_template |
| code | **dao** | /dao/contactStreamType_core.go |
| code | **datamodel** | /datamodel/contactStreamType_core.go |
| code | **menu** | /design/menu/contactStreamType.json |
| html | **list** | /ContactStreamType_List.html |
| html | **view** | /ContactStreamType_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **28/06/2022** at **16:10:47**
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