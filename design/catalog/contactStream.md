# **ContactStream** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**ContactStream** (contactstream) |
|Endpoint 	    |**/ContactStream...** [^1]|
|Endpoint Query |**ID**|
Glyph|**far fa-comment** ("")
Friendly Name|**Conversations**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/ContactStream/ContactStreamList) [Exportable]
* **View** (/ContactStream/ContactStreamView)











##  Provides
 * Lookup (StreamId Description)






##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **contactManagerStream**
SQL Table Key | **streamId**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**StreamId**|Int|true|true|false|false|||||Y|streamId|0|false|false|false|text||
|**ContactId**|Int|false|true|false|false|||||Y|contactId|0|false|false|false|text||
|**Usr**|String|false|true|false|false|||||Y|usr||false|false|false|text||
|**Description**|String|false|true|false|false|||||Y|description||false|false|false|text||
|**StreamTypeId**|Int|false|true|false|false|OL|ContactStreamType|||Y|streamTypeId|0|false|false|false|text||
|**StreamStatusId**|Int|false|true|false|false|OL|ContactStreamStatus|||Y|streamStatusId|0|false|false|false|text||
|**RecordState**|Int|false|true|false|false|LL|cmrecordstatus|||Y|recordState|0|false|false|false|text||
|**CreatedDateTime**|Time|false|true|false|true|||||Y|createdDateTime||false|false|false|datetime||
|**OpenedDateTime**|Time|false|true|false|true|||||Y|openedDateTime||false|false|false|datetime||
|**ClosedDateTime**|Time|false|true|false|true|||||Y|closedDateTime||false|false|false|datetime||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/contactStream_core.go |
| code | **adaptor** | /adaptor/contactStream_impl.go_template |
| code | **dao** | /dao/contactStream_core.go |
| code | **datamodel** | /datamodel/contactStream_core.go |
| code | **menu** | /design/menu/contactStream.json |
| html | **list** | /ContactStream_List.html |
| html | **view** | /ContactStream_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **28/06/2022** at **16:10:46**
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