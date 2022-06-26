# **ContactDealRefNo** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**ContactDealRefNo** (contactdealrefno) |
|Endpoint 	    |**/ContactDealRefNo...** [^1]|
|Endpoint Query |**ID**|
Glyph|**far fa-comment** ("")
Friendly Name|**Conversation Deals**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/ContactDealRefNo/ContactDealRefNoList) [Exportable]
* **View** (/ContactDealRefNo/ContactDealRefNoView)











##  Provides
 * Lookup (NoteId RefNo)






##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **contactManagerDealRefNo**
SQL Table Key | **noteId**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**RefNo**|String|false|true|false|false|OL|Transaction|||Y|refNo||false|false|false|text||
|**NoteId**|Int|false|true|false|false|OL|CounterpartyNotes|||Y|noteId|0|false|false|false|text||
|**RecordState**|Int|false|true|false|false|LL|cmrecordstate|||Y|recordState|0|false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/contactDealRefNo_core.go |
| code | **adaptor** | /adaptor/contactDealRefNo_impl.go_template |
| code | **dao** | /dao/contactDealRefNo_core.go |
| code | **datamodel** | /datamodel/contactDealRefNo_core.go |
| code | **menu** | /design/menu/contactDealRefNo.json |
| html | **list** | /ContactDealRefNo_List.html |
| html | **view** | /ContactDealRefNo_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **26/06/2022** at **18:48:20**
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